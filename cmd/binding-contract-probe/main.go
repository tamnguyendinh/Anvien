package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tamnguyendinh/anvien/internal/parser"
	"github.com/tamnguyendinh/anvien/internal/providers/tsjs"
	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

const (
	probeSchemaVersion = "binding-contract-probe.v1"
	parseTimeout       = 30 * time.Second
)

type singleFileFlag struct {
	value string
	set   bool
}

func (value *singleFileFlag) String() string {
	return value.value
}

func (value *singleFileFlag) Set(raw string) error {
	if value.set {
		return errors.New("-file must be provided exactly once")
	}
	if strings.TrimSpace(raw) == "" {
		return errors.New("-file must not be empty")
	}
	value.value = raw
	value.set = true
	return nil
}

type nameListFlag []string

func (values *nameListFlag) String() string {
	return strings.Join(*values, ",")
}

func (values *nameListFlag) Set(raw string) error {
	if raw == "" {
		return errors.New("-name must not be empty")
	}
	if strings.TrimSpace(raw) != raw {
		return errors.New("-name must not have leading or trailing whitespace")
	}
	*values = append(*values, raw)
	return nil
}

type commandOptions struct {
	filePath string
	names    []string
}

type probeOutput struct {
	SchemaVersion string                    `json:"schemaVersion"`
	File          probeFileMetadata         `json:"file"`
	Results       []bindingContractMetadata `json:"results"`
}

type probeFileMetadata struct {
	Path     string           `json:"path"`
	Language scanner.Language `json:"language"`
	SHA256   string           `json:"sha256"`
	Bytes    int              `json:"bytes"`
	Grammar  string           `json:"grammar"`
	RootKind string           `json:"rootKind"`
}

type bindingContractMetadata struct {
	Name                    string               `json:"name"`
	BindingLeaf             bindingLeafMetadata  `json:"bindingLeaf"`
	Definition              definitionMetadata   `json:"definition"`
	LexicalOwner            lexicalOwnerMetadata `json:"lexicalOwner"`
	OwnerLocalBinding       scopeir.BindingFact  `json:"ownerLocalBinding"`
	OwnedDefIDOccurrences   int                  `json:"ownedDefIdOccurrences"`
	LocalBindingOccurrences int                  `json:"localBindingOccurrences"`
}

type bindingLeafMetadata struct {
	FilePath       string                           `json:"filePath"`
	FileHash       string                           `json:"fileHash"`
	Name           string                           `json:"name"`
	Range          scopeir.Range                    `json:"range"`
	SelectionRange scopeir.Range                    `json:"selectionRange"`
	Path           []scopeir.BindingPathSegment     `json:"path"`
	Rest           bool                             `json:"rest"`
	Default        bool                             `json:"default"`
	Provenance     scopeir.BindingPatternProvenance `json:"provenance"`
}

type definitionMetadata struct {
	DefID          string            `json:"defId"`
	FilePath       string            `json:"filePath"`
	FileHash       string            `json:"fileHash"`
	Name           string            `json:"name"`
	Label          scopeir.NodeLabel `json:"label"`
	Range          scopeir.Range     `json:"range"`
	SelectionRange scopeir.Range     `json:"selectionRange"`
}

type lexicalOwnerMetadata struct {
	ID          string            `json:"id"`
	Parent      *string           `json:"parent"`
	Kind        scopeir.ScopeKind `json:"kind"`
	Range       scopeir.Range     `json:"range"`
	FilePath    string            `json:"filePath"`
	FileHash    string            `json:"fileHash"`
	OwnedDefIDs []string          `json:"ownedDefIds"`
}

func main() {
	err := run(os.Args[1:], os.Stdout, os.Stderr)
	if err == nil || errors.Is(err, flag.ErrHelp) {
		return
	}
	fmt.Fprintf(os.Stderr, "binding-contract-probe: %v\n", err)
	os.Exit(1)
}

func run(args []string, stdout io.Writer, stderr io.Writer) error {
	options, err := parseCommandOptions(args, stderr)
	if err != nil {
		return err
	}

	output, err := probeFile(options.filePath, options.names)
	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(output); err != nil {
		return fmt.Errorf("encode metadata JSON: %w", err)
	}
	if _, err := io.Copy(stdout, &buffer); err != nil {
		return fmt.Errorf("write metadata JSON: %w", err)
	}
	return nil
}

func parseCommandOptions(args []string, stderr io.Writer) (commandOptions, error) {
	flags := flag.NewFlagSet("binding-contract-probe", flag.ContinueOnError)
	flags.SetOutput(stderr)

	var file singleFileFlag
	var names nameListFlag
	flags.Var(&file, "file", "one TypeScript source file to inspect")
	flags.Var(&names, "name", "local name to validate; repeat in the required output order")
	if err := flags.Parse(args); err != nil {
		return commandOptions{}, err
	}
	if flags.NArg() != 0 {
		return commandOptions{}, fmt.Errorf("unexpected positional arguments: %d", flags.NArg())
	}
	if !file.set {
		return commandOptions{}, errors.New("missing -file")
	}
	if len(names) == 0 {
		return commandOptions{}, errors.New("at least one -name is required")
	}

	seen := make(map[string]struct{}, len(names))
	for _, name := range names {
		if _, exists := seen[name]; exists {
			return commandOptions{}, fmt.Errorf("duplicate requested local name %q", name)
		}
		seen[name] = struct{}{}
	}

	return commandOptions{filePath: file.value, names: append([]string(nil), names...)}, nil
}

func probeFile(filePath string, names []string) (probeOutput, error) {
	absolutePath, err := filepath.Abs(filePath)
	if err != nil {
		return probeOutput{}, fmt.Errorf("resolve TypeScript file %q: %w", filePath, err)
	}
	absolutePath = filepath.Clean(absolutePath)
	canonicalPath := filepath.ToSlash(absolutePath)

	info, err := os.Stat(absolutePath)
	if err != nil {
		return probeOutput{}, fmt.Errorf("stat TypeScript file %q: %w", canonicalPath, err)
	}
	if !info.Mode().IsRegular() {
		return probeOutput{}, fmt.Errorf("TypeScript file %q is not a regular file", canonicalPath)
	}
	language, ok := scanner.DetectLanguage(canonicalPath)
	if !ok || language != scanner.TypeScript {
		return probeOutput{}, fmt.Errorf("file %q is not a recognized TypeScript source file", canonicalPath)
	}

	source, err := os.ReadFile(absolutePath)
	if err != nil {
		return probeOutput{}, fmt.Errorf("read TypeScript file %q: %w", canonicalPath, err)
	}
	sum := sha256.Sum256(source)
	fileHash := hex.EncodeToString(sum[:])

	ir, parseMetadata, err := extractScopeIR(canonicalPath, fileHash, source)
	if err != nil {
		return probeOutput{}, err
	}
	results, err := validateBindingContracts(ir, names)
	if err != nil {
		return probeOutput{}, err
	}

	return probeOutput{
		SchemaVersion: probeSchemaVersion,
		File: probeFileMetadata{
			Path:     canonicalPath,
			Language: scanner.TypeScript,
			SHA256:   fileHash,
			Bytes:    len(source),
			Grammar:  parseMetadata.Grammar,
			RootKind: parseMetadata.RootKind,
		},
		Results: results,
	}, nil
}

func extractScopeIR(filePath string, fileHash string, source []byte) (scopeir.ScopeIR, parser.Result, error) {
	pool := parser.NewPool(nil, parser.PoolOptions{ParseTimeout: parseTimeout})
	defer pool.Close()

	parsed, err := pool.Parse(context.Background(), parser.Request{
		FilePath: filePath,
		Language: scanner.TypeScript,
		Source:   source,
	})
	if err != nil {
		return scopeir.ScopeIR{}, parser.Result{}, fmt.Errorf("parse TypeScript file %q: %w", filePath, err)
	}
	defer parsed.Close()
	if parsed.HasError {
		return scopeir.ScopeIR{}, parser.Result{}, fmt.Errorf("parse TypeScript file %q: syntax tree contains errors", filePath)
	}

	ir, err := tsjs.Extract(tsjs.Request{
		FilePath: filePath,
		FileHash: fileHash,
		Language: scanner.TypeScript,
		Source:   source,
		Root:     parsed.Tree.RootNode(),
	})
	if err != nil {
		return scopeir.ScopeIR{}, parser.Result{}, fmt.Errorf("extract TypeScript ScopeIR for %q: %w", filePath, err)
	}
	metadata := *parsed
	metadata.Tree = nil
	return ir, metadata, nil
}

func validateBindingContracts(ir scopeir.ScopeIR, names []string) ([]bindingContractMetadata, error) {
	normalized := ir.Normalized()
	results := make([]bindingContractMetadata, 0, len(names))
	for _, name := range names {
		result, err := validateBindingContract(normalized, name)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func validateBindingContract(ir scopeir.ScopeIR, name string) (bindingContractMetadata, error) {
	leaves := make([]scopeir.BindingLeafFact, 0, 1)
	for _, leaf := range ir.BindingLeaves {
		if leaf.Name == name && leaf.Provenance.Context == scopeir.BindingContextVariable {
			leaves = append(leaves, leaf)
		}
	}
	if len(leaves) != 1 {
		return bindingContractMetadata{}, fmt.Errorf(
			"local %q: expected exactly one variable-context BindingLeafFact, found %d",
			name,
			len(leaves),
		)
	}
	leaf := leaves[0]
	if err := validateLeaf(ir, leaf); err != nil {
		return bindingContractMetadata{}, fmt.Errorf("local %q: BindingLeafFact: %w", name, err)
	}

	definitions := make([]scopeir.DefinitionFact, 0, 1)
	variableNameCandidates := 0
	for _, definition := range ir.Definitions {
		if definition.Name != name || definition.Label != scopeir.NodeVariable {
			continue
		}
		variableNameCandidates++
		if definition.FilePath == leaf.FilePath &&
			definition.FileHash == leaf.FileHash &&
			definition.Range == leaf.Range &&
			optionalRangeEqual(definition.SelectionRange, leaf.SelectionRange) {
			definitions = append(definitions, definition)
		}
	}
	if len(definitions) != 1 {
		return bindingContractMetadata{}, fmt.Errorf(
			"local %q: expected exactly one matching Variable DefinitionFact, found %d (same-name Variable candidates: %d)",
			name,
			len(definitions),
			variableNameCandidates,
		)
	}
	definition := definitions[0]
	if definition.ID == "" {
		return bindingContractMetadata{}, fmt.Errorf("local %q: matching Variable DefinitionFact has an empty DefID", name)
	}
	if definition.SelectionRange == nil {
		return bindingContractMetadata{}, fmt.Errorf("local %q: matching Variable DefinitionFact has no selection range", name)
	}
	definitionIDOccurrences := 0
	for _, candidate := range ir.Definitions {
		if candidate.ID == definition.ID {
			definitionIDOccurrences++
		}
	}
	if definitionIDOccurrences != 1 {
		return bindingContractMetadata{}, fmt.Errorf(
			"local %q: DefID %q occurs in %d DefinitionFacts; expected exactly one",
			name,
			definition.ID,
			definitionIDOccurrences,
		)
	}

	type ownerCandidate struct {
		scope       scopeir.ScopeFact
		occurrences int
	}
	owners := make([]ownerCandidate, 0, 1)
	for _, candidate := range ir.Scopes {
		occurrences := countString(candidate.OwnedDefIDs, definition.ID)
		if occurrences != 0 {
			owners = append(owners, ownerCandidate{scope: candidate, occurrences: occurrences})
		}
	}
	if len(owners) != 1 {
		return bindingContractMetadata{}, fmt.Errorf(
			"local %q: expected exactly one lexical owner ScopeFact for DefID %q, found %d",
			name,
			definition.ID,
			len(owners),
		)
	}
	owner := owners[0]
	if owner.occurrences != 1 {
		return bindingContractMetadata{}, fmt.Errorf(
			"local %q: lexical owner %q contains DefID %q %d times in OwnedDefIDs; expected exactly once",
			name,
			owner.scope.ID,
			definition.ID,
			owner.occurrences,
		)
	}
	if err := validateOwner(ir, owner.scope, leaf, definition); err != nil {
		return bindingContractMetadata{}, fmt.Errorf("local %q: lexical owner %q: %w", name, owner.scope.ID, err)
	}

	exactBindings := make([]scopeir.BindingFact, 0, 1)
	sameNameLocals := 0
	sameDefLocals := 0
	for _, binding := range owner.scope.Bindings {
		if binding.Origin != scopeir.BindingLocal {
			continue
		}
		if binding.Name == name {
			sameNameLocals++
		}
		if binding.DefID == definition.ID {
			sameDefLocals++
		}
		if binding.Name == name && binding.DefID == definition.ID {
			exactBindings = append(exactBindings, binding)
		}
	}
	if len(exactBindings) != 1 || sameNameLocals != 1 || sameDefLocals != 1 {
		return bindingContractMetadata{}, fmt.Errorf(
			"local %q: owner-local BindingLocal mismatch for DefID %q (exact=%d same-name=%d same-def=%d); expected 1/1/1",
			name,
			definition.ID,
			len(exactBindings),
			sameNameLocals,
			sameDefLocals,
		)
	}

	selectionRange := *leaf.SelectionRange
	definitionSelectionRange := *definition.SelectionRange
	return bindingContractMetadata{
		Name: name,
		BindingLeaf: bindingLeafMetadata{
			FilePath:       leaf.FilePath,
			FileHash:       leaf.FileHash,
			Name:           leaf.Name,
			Range:          leaf.Range,
			SelectionRange: selectionRange,
			Path:           cloneBindingPath(leaf.Path),
			Rest:           leaf.Rest,
			Default:        leaf.Default,
			Provenance:     leaf.Provenance,
		},
		Definition: definitionMetadata{
			DefID:          definition.ID,
			FilePath:       definition.FilePath,
			FileHash:       definition.FileHash,
			Name:           definition.Name,
			Label:          definition.Label,
			Range:          definition.Range,
			SelectionRange: definitionSelectionRange,
		},
		LexicalOwner: lexicalOwnerMetadata{
			ID:          owner.scope.ID,
			Parent:      cloneString(owner.scope.Parent),
			Kind:        owner.scope.Kind,
			Range:       owner.scope.Range,
			FilePath:    owner.scope.FilePath,
			FileHash:    owner.scope.FileHash,
			OwnedDefIDs: append([]string(nil), owner.scope.OwnedDefIDs...),
		},
		OwnerLocalBinding:       exactBindings[0],
		OwnedDefIDOccurrences:   owner.occurrences,
		LocalBindingOccurrences: len(exactBindings),
	}, nil
}

func validateLeaf(ir scopeir.ScopeIR, leaf scopeir.BindingLeafFact) error {
	if leaf.FilePath != ir.FilePath {
		return fmt.Errorf("file path %q does not match ScopeIR file path %q", leaf.FilePath, ir.FilePath)
	}
	if leaf.FileHash != ir.FileHash {
		return fmt.Errorf("file hash %q does not match ScopeIR file hash %q", leaf.FileHash, ir.FileHash)
	}
	if err := validateRange("range", leaf.Range); err != nil {
		return err
	}
	if leaf.SelectionRange == nil {
		return errors.New("selection range is missing")
	}
	if err := validateRange("selection range", *leaf.SelectionRange); err != nil {
		return err
	}
	if !rangeContains(leaf.Range, *leaf.SelectionRange) {
		return errors.New("range does not contain selection range")
	}
	if leaf.Provenance.Context != scopeir.BindingContextVariable {
		return fmt.Errorf("context is %q; expected %q", leaf.Provenance.Context, scopeir.BindingContextVariable)
	}
	if leaf.Provenance.PatternKind == "" {
		return errors.New("provenance pattern kind is empty")
	}
	if err := validateRange("provenance construct range", leaf.Provenance.ConstructRange); err != nil {
		return err
	}
	if err := validateRange("provenance pattern range", leaf.Provenance.PatternRange); err != nil {
		return err
	}
	if !rangeContains(leaf.Provenance.ConstructRange, leaf.Provenance.PatternRange) {
		return errors.New("provenance construct range does not contain pattern range")
	}
	if !rangeContains(leaf.Provenance.PatternRange, *leaf.SelectionRange) {
		return errors.New("provenance pattern range does not contain selection range")
	}
	for index, segment := range leaf.Path {
		if err := validatePathSegment(segment); err != nil {
			return fmt.Errorf("path segment %d: %w", index, err)
		}
	}
	return nil
}

func validatePathSegment(segment scopeir.BindingPathSegment) error {
	if err := validateRange("source range", segment.SourceRange); err != nil {
		return err
	}
	switch segment.Kind {
	case scopeir.BindingPathArrayIndex:
		if segment.ArrayIndex == nil || *segment.ArrayIndex < 0 {
			return errors.New("array-index segment has no non-negative array index")
		}
		if segment.PropertyName != "" || segment.ComputedExpression != "" {
			return errors.New("array-index segment contains property metadata")
		}
	case scopeir.BindingPathStaticProperty:
		if segment.ArrayIndex != nil || segment.PropertyName == "" || segment.ComputedExpression != "" {
			return errors.New("static-property segment has invalid kind-specific metadata")
		}
	case scopeir.BindingPathComputedProperty:
		if segment.ArrayIndex != nil || segment.PropertyName != "" || segment.ComputedExpression == "" {
			return errors.New("computed-property segment has invalid kind-specific metadata")
		}
	default:
		return fmt.Errorf("unsupported typed path kind %q", segment.Kind)
	}
	return nil
}

func validateOwner(
	ir scopeir.ScopeIR,
	owner scopeir.ScopeFact,
	leaf scopeir.BindingLeafFact,
	definition scopeir.DefinitionFact,
) error {
	if owner.ID == "" {
		return errors.New("scope ID is empty")
	}
	scopeIDOccurrences := 0
	for _, candidate := range ir.Scopes {
		if candidate.ID == owner.ID {
			scopeIDOccurrences++
		}
	}
	if scopeIDOccurrences != 1 {
		return fmt.Errorf("scope ID occurs %d times; expected exactly once", scopeIDOccurrences)
	}
	if owner.FilePath != ir.FilePath || owner.FilePath != definition.FilePath || owner.FilePath != leaf.FilePath {
		return errors.New("file path does not match the ScopeIR, BindingLeafFact, and DefinitionFact")
	}
	if owner.FileHash != ir.FileHash || owner.FileHash != definition.FileHash || owner.FileHash != leaf.FileHash {
		return errors.New("file hash does not match the ScopeIR, BindingLeafFact, and DefinitionFact")
	}
	if err := validateRange("scope range", owner.Range); err != nil {
		return err
	}
	if !rangeContains(owner.Range, leaf.Range) || !rangeContains(owner.Range, definition.Range) {
		return errors.New("scope range does not contain the leaf and definition ranges")
	}
	if leaf.SelectionRange == nil || definition.SelectionRange == nil {
		return errors.New("leaf or definition selection range is missing")
	}
	if !rangeContains(owner.Range, *leaf.SelectionRange) || !rangeContains(owner.Range, *definition.SelectionRange) {
		return errors.New("scope range does not contain the leaf and definition selection ranges")
	}
	return nil
}

func validateRange(label string, value scopeir.Range) error {
	if value.StartLine <= 0 || value.EndLine <= 0 || value.StartCol < 0 || value.EndCol < 0 {
		return fmt.Errorf("%s has invalid coordinates", label)
	}
	if value.StartLine > value.EndLine ||
		(value.StartLine == value.EndLine && value.StartCol > value.EndCol) {
		return fmt.Errorf("%s is reversed", label)
	}
	return nil
}

func rangeContains(outer scopeir.Range, inner scopeir.Range) bool {
	startsBefore := outer.StartLine < inner.StartLine ||
		(outer.StartLine == inner.StartLine && outer.StartCol <= inner.StartCol)
	endsAfter := outer.EndLine > inner.EndLine ||
		(outer.EndLine == inner.EndLine && outer.EndCol >= inner.EndCol)
	return startsBefore && endsAfter
}

func optionalRangeEqual(left *scopeir.Range, right *scopeir.Range) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	return *left == *right
}

func countString(values []string, expected string) int {
	count := 0
	for _, value := range values {
		if value == expected {
			count++
		}
	}
	return count
}

func cloneBindingPath(path []scopeir.BindingPathSegment) []scopeir.BindingPathSegment {
	cloned := make([]scopeir.BindingPathSegment, len(path))
	copy(cloned, path)
	for index := range cloned {
		if cloned[index].ArrayIndex == nil {
			continue
		}
		value := *cloned[index].ArrayIndex
		cloned[index].ArrayIndex = &value
	}
	return cloned
}

func cloneString(value *string) *string {
	if value == nil {
		return nil
	}
	cloned := *value
	return &cloned
}
