//go:build ladybugdb

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/tamnguyendinh/anvien/internal/filecontext"
	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/lbugnative"
)

type point struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type sourceRange struct {
	Start point `json:"start"`
	End   point `json:"end"`
}

type sourceRef struct {
	Path           string      `json:"path"`
	FileSHA256     string      `json:"fileSha256"`
	Range          sourceRange `json:"range"`
	SelectionRange sourceRange `json:"selectionRange"`
	LexicalOwner   string      `json:"lexicalOwner"`
}

type expectedAccess struct {
	State string `json:"state"`
	Value any    `json:"value"`
}

type expectedCompatibility struct {
	IsExported bool `json:"isExported"`
}

type expectedValues struct {
	ExportedName              string                `json:"exportedName"`
	ExportKind                string                `json:"exportKind"`
	Meanings                  []string              `json:"meanings"`
	TypeOnly                  bool                  `json:"typeOnly"`
	ExportFactCount           int                   `json:"exportFactCount"`
	FileToExportRelationCount int                   `json:"fileToExportRelationCount"`
	Access                    expectedAccess        `json:"access"`
	Compatibility             expectedCompatibility `json:"compatibility"`
	ExportDiagnosticCount     int                   `json:"exportDiagnosticCount"`
	Child05DerivedState       string                `json:"child05DerivedState"`
}

type oracleRow struct {
	RowID           string         `json:"rowId"`
	Polarity        string         `json:"polarity"`
	Source          sourceRef      `json:"source"`
	DeclarationKind string         `json:"declarationKind"`
	LocalName       string         `json:"localName"`
	Expected        expectedValues `json:"expected"`
}

type oracle struct {
	SchemaVersion    string      `json:"schemaVersion"`
	OracleID         string      `json:"oracleId"`
	EvidenceID       string      `json:"evidenceId"`
	PositiveCount    int         `json:"positiveCount"`
	NegativeCount    int         `json:"negativeCount"`
	PositiveRows     []oracleRow `json:"positiveRows"`
	NegativeControls []oracleRow `json:"negativeControls"`
}

type fieldCheck struct {
	Expected any  `json:"expected"`
	Actual   any  `json:"actual"`
	Pass     bool `json:"pass"`
}

type rowResult struct {
	RowID                  string                `json:"rowId"`
	Polarity               string                `json:"polarity"`
	Path                   string                `json:"path"`
	LocalName              string                `json:"localName"`
	DefinitionMatchCount   int                   `json:"definitionMatchCount"`
	DefinitionID           string                `json:"definitionId,omitempty"`
	DefinitionLabel        string                `json:"definitionLabel,omitempty"`
	ExportNodeIDs          []string              `json:"exportNodeIds"`
	GraphRelationCount     int                   `json:"graphRelationCount"`
	LadybugRelationCount   int                   `json:"ladybugRelationCount"`
	GraphLadybugDifference int                   `json:"graphLadybugDifferenceCount"`
	ReaderSymbolFound      bool                  `json:"readerSymbolFound"`
	ReaderExported         bool                  `json:"readerExported"`
	Fields                 map[string]fieldCheck `json:"fields"`
	Failures               []string              `json:"failures"`
	Pass                   bool                  `json:"pass"`
}

type graphData struct {
	NodeCount                   int
	RelationshipCount           int
	DuplicateNodeIDCount        int
	OrphanRelationshipEndpoints int
	OrphanLocalDefinitionRefs   int
	ExportDiagnosticCount       int
	ForbiddenStateCount         int
	ForbiddenStateSamples       []string
	GlobalExportCount           int
	TargetExports               []graph.Node
	RelevantNodes               []graph.Node
	RelevantRelationships       []graph.Relationship
	RelevantNodeIDs             map[string]struct{}
	AllNodeIDs                  map[string]struct{}
}

type ladybugParity struct {
	ExportRowCount            int            `json:"exportRowCount"`
	ExportFieldComparisons    int            `json:"exportFieldComparisons"`
	ExportFieldDifferences    int            `json:"exportFieldDifferences"`
	ExportIDsMissingInLadybug []string       `json:"exportIdsMissingInLadybug"`
	ExportIDsMissingInJSON    []string       `json:"exportIdsMissingInGraphJson"`
	DefinitionComparisons     int            `json:"definitionCompatibilityComparisons"`
	DefinitionDifferences     int            `json:"definitionCompatibilityDifferences"`
	RelationshipRowCount      int            `json:"relationshipRowCount"`
	RelationshipDifferences   int            `json:"relationshipDifferences"`
	PerExportDifferenceCount  map[string]int `json:"perExportDifferenceCount"`
	PerDefinitionDifference   map[string]int `json:"perDefinitionDifferenceCount"`
	LadybugRelationCountByID  map[string]int `json:"ladybugRelationCountByExportId"`
	GraphRelationCountByID    map[string]int `json:"graphRelationCountByExportId"`
	ReaderOpenedReadOnly      bool           `json:"readerOpenedReadOnly"`
}

type readerFileResult struct {
	Found               bool            `json:"found"`
	ExportedSymbolCount int             `json:"exportedSymbolCount"`
	SymbolExported      map[string]bool `json:"symbolExportedById"`
}

type result struct {
	SchemaVersion      string                      `json:"schemaVersion"`
	GeneratedAt        string                      `json:"generatedAt"`
	Status             string                      `json:"status"`
	OracleID           string                      `json:"oracleId"`
	EvidenceID         string                      `json:"evidenceId"`
	Graph              map[string]any              `json:"graph"`
	Ladybug            ladybugParity               `json:"ladybug"`
	AffectedReader     map[string]readerFileResult `json:"affectedReader"`
	PositiveRows       []rowResult                 `json:"positiveRows"`
	NegativeControls   []rowResult                 `json:"negativeControls"`
	PositivePassCount  int                         `json:"positivePassCount"`
	NegativePassCount  int                         `json:"negativePassCount"`
	ComparisonComplete bool                        `json:"comparisonComplete"`
	ContractPass       bool                        `json:"contractPass"`
	Findings           []string                    `json:"findings"`
	Error              string                      `json:"error,omitempty"`
}

func main() {
	oraclePath := flag.String("oracle", "", "sealed expected-values.json")
	graphPath := flag.String("graph", "", "fresh graph.json")
	ladybugPath := flag.String("ladybug", "", "fresh Ladybug database")
	outPath := flag.String("out", "", "result JSON")
	flag.Parse()

	res, err := validate(*oraclePath, *graphPath, *ladybugPath)
	if err != nil {
		res = result{
			SchemaVersion: "p4c2.qa-comparison.v1",
			GeneratedAt:   time.Now().Format(time.RFC3339Nano),
			Status:        "BLOCKED",
			Error:         err.Error(),
		}
	}
	data, marshalErr := json.MarshalIndent(res, "", "  ")
	if marshalErr != nil {
		fmt.Fprintln(os.Stderr, marshalErr)
		os.Exit(4)
	}
	data = append(data, '\n')
	if err := os.WriteFile(*outPath, data, 0o644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(4)
	}
	fmt.Printf("comparison_status=%s\npositive=%d/%d\nnegative=%d/%d\ncomplete=%t\ncontract_pass=%t\n",
		res.Status, res.PositivePassCount, len(res.PositiveRows), res.NegativePassCount, len(res.NegativeControls), res.ComparisonComplete, res.ContractPass)
	if res.Error != "" {
		fmt.Printf("error=%s\n", res.Error)
		os.Exit(3)
	}
	if !res.ContractPass {
		os.Exit(2)
	}
}

func validate(oraclePath, graphPath, ladybugPath string) (result, error) {
	var o oracle
	data, err := os.ReadFile(oraclePath)
	if err != nil {
		return result{}, fmt.Errorf("read oracle: %w", err)
	}
	if err := json.Unmarshal(data, &o); err != nil {
		return result{}, fmt.Errorf("parse oracle: %w", err)
	}
	paths := map[string]struct{}{}
	for _, row := range append(append([]oracleRow{}, o.PositiveRows...), o.NegativeControls...) {
		paths[row.Source.Path] = struct{}{}
	}

	gd, err := readGraph(graphPath, paths)
	if err != nil {
		return result{}, err
	}

	definitionMatches := map[string][]graph.Node{}
	for _, row := range append(append([]oracleRow{}, o.PositiveRows...), o.NegativeControls...) {
		definitionMatches[row.RowID] = matchingDefinitions(row, gd.RelevantNodes)
	}

	parity, err := compareLadybug(ladybugPath, paths, gd, definitionMatches)
	if err != nil {
		return result{}, fmt.Errorf("Ladybug read-only parity: %w", err)
	}

	reader := compareAffectedReader(paths, gd)
	positive := make([]rowResult, 0, len(o.PositiveRows))
	negative := make([]rowResult, 0, len(o.NegativeControls))
	positivePass := 0
	negativePass := 0
	for _, row := range o.PositiveRows {
		r := evaluateRow(row, definitionMatches[row.RowID], gd, parity, reader)
		positive = append(positive, r)
		if r.Pass {
			positivePass++
		}
	}
	for _, row := range o.NegativeControls {
		r := evaluateRow(row, definitionMatches[row.RowID], gd, parity, reader)
		negative = append(negative, r)
		if r.Pass {
			negativePass++
		}
	}

	targetDirect := 0
	for _, node := range gd.TargetExports {
		if normalizedExportKind(node) == "direct_declaration" {
			targetDirect++
		}
	}
	readerComplete := true
	for path := range paths {
		if !reader[path].Found {
			readerComplete = false
		}
	}
	parityPass := parity.ExportFieldDifferences == 0 && len(parity.ExportIDsMissingInLadybug) == 0 && len(parity.ExportIDsMissingInJSON) == 0 && parity.DefinitionDifferences == 0 && parity.RelationshipDifferences == 0
	integrityPass := gd.DuplicateNodeIDCount == 0 && gd.OrphanRelationshipEndpoints == 0 && gd.OrphanLocalDefinitionRefs == 0 && gd.ExportDiagnosticCount == 0 && gd.ForbiddenStateCount == 0
	contractPass := positivePass == o.PositiveCount && negativePass == o.NegativeCount && targetDirect == o.PositiveCount && parityPass && readerComplete && integrityPass
	findings := []string{}
	if positivePass != o.PositiveCount {
		findings = append(findings, fmt.Sprintf("positive rows %d/%d", positivePass, o.PositiveCount))
	}
	if negativePass != o.NegativeCount {
		findings = append(findings, fmt.Sprintf("negative controls %d/%d", negativePass, o.NegativeCount))
	}
	if targetDirect != o.PositiveCount {
		findings = append(findings, fmt.Sprintf("target direct exports %d, expected %d", targetDirect, o.PositiveCount))
	}
	if !parityPass {
		findings = append(findings, "Graph JSON/Ladybug record parity mismatch")
	}
	if !readerComplete {
		findings = append(findings, "affected FileContext reader did not open every target file")
	}
	if !integrityPass {
		findings = append(findings, "graph integrity/forbidden-state invariant mismatch")
	}
	status := "PASS"
	if !contractPass {
		status = "FAIL_COMPLETE"
	}
	return result{
		SchemaVersion: "p4c2.qa-comparison.v1",
		GeneratedAt:   time.Now().Format(time.RFC3339Nano),
		Status:        status,
		OracleID:      o.OracleID,
		EvidenceID:    o.EvidenceID,
		Graph: map[string]any{
			"nodeCount":                           gd.NodeCount,
			"relationshipCount":                   gd.RelationshipCount,
			"globalExportCount":                   gd.GlobalExportCount,
			"targetExportCount":                   len(gd.TargetExports),
			"targetDirectExportCount":             targetDirect,
			"duplicateNodeIdCount":                gd.DuplicateNodeIDCount,
			"orphanRelationshipEndpointCount":     gd.OrphanRelationshipEndpoints,
			"orphanLocalDefinitionReferenceCount": gd.OrphanLocalDefinitionRefs,
			"exportDiagnosticCount":               gd.ExportDiagnosticCount,
			"forbiddenChild05StateCount":          gd.ForbiddenStateCount,
			"forbiddenChild05StateSamples":        gd.ForbiddenStateSamples,
		},
		Ladybug:            parity,
		AffectedReader:     reader,
		PositiveRows:       positive,
		NegativeControls:   negative,
		PositivePassCount:  positivePass,
		NegativePassCount:  negativePass,
		ComparisonComplete: true,
		ContractPass:       contractPass,
		Findings:           findings,
	}, nil
}

func readGraph(path string, targetPaths map[string]struct{}) (*graphData, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open graph: %w", err)
	}
	defer f.Close()
	dec := json.NewDecoder(bufio.NewReaderSize(f, 4<<20))
	tok, err := dec.Token()
	if err != nil || tok != json.Delim('{') {
		return nil, fmt.Errorf("graph root is not an object: %v", err)
	}
	gd := &graphData{AllNodeIDs: map[string]struct{}{}, RelevantNodeIDs: map[string]struct{}{}}
	for dec.More() {
		keyToken, err := dec.Token()
		if err != nil {
			return nil, err
		}
		key := keyToken.(string)
		switch key {
		case "nodes":
			if tok, err := dec.Token(); err != nil || tok != json.Delim('[') {
				return nil, fmt.Errorf("nodes is not an array: %v", err)
			}
			for dec.More() {
				var node graph.Node
				if err := dec.Decode(&node); err != nil {
					return nil, fmt.Errorf("decode node: %w", err)
				}
				gd.NodeCount++
				if _, exists := gd.AllNodeIDs[node.ID]; exists {
					gd.DuplicateNodeIDCount++
				}
				gd.AllNodeIDs[node.ID] = struct{}{}
				label := string(node.Label)
				if strings.Contains(strings.ToLower(label), "exportdiagnostic") {
					gd.ExportDiagnosticCount++
				}
				if label == "Export" {
					gd.GlobalExportCount++
					localID := propString(node, "localDefinitionNodeId")
					_ = localID
					for key := range node.Properties {
						if forbiddenKey(key) {
							gd.ForbiddenStateCount++
							if len(gd.ForbiddenStateSamples) < 20 {
								gd.ForbiddenStateSamples = append(gd.ForbiddenStateSamples, node.ID+":"+key)
							}
						}
					}
				}
				filePath := propString(node, "filePath")
				_, relevant := targetPaths[filePath]
				if relevant {
					gd.RelevantNodes = append(gd.RelevantNodes, node)
					gd.RelevantNodeIDs[node.ID] = struct{}{}
					if label == "Export" {
						gd.TargetExports = append(gd.TargetExports, node)
					}
				}
			}
			if _, err := dec.Token(); err != nil {
				return nil, err
			}
		case "relationships":
			if tok, err := dec.Token(); err != nil || tok != json.Delim('[') {
				return nil, fmt.Errorf("relationships is not an array: %v", err)
			}
			for dec.More() {
				var rel graph.Relationship
				if err := dec.Decode(&rel); err != nil {
					return nil, fmt.Errorf("decode relationship: %w", err)
				}
				gd.RelationshipCount++
				_, sourceExists := gd.AllNodeIDs[rel.SourceID]
				_, targetExists := gd.AllNodeIDs[rel.TargetID]
				if !sourceExists {
					gd.OrphanRelationshipEndpoints++
				}
				if !targetExists {
					gd.OrphanRelationshipEndpoints++
				}
				_, sourceRelevant := gd.RelevantNodeIDs[rel.SourceID]
				_, targetRelevant := gd.RelevantNodeIDs[rel.TargetID]
				if sourceRelevant && targetRelevant {
					gd.RelevantRelationships = append(gd.RelevantRelationships, rel)
				}
			}
			if _, err := dec.Token(); err != nil {
				return nil, err
			}
		default:
			var discard any
			if err := dec.Decode(&discard); err != nil {
				return nil, err
			}
		}
	}
	if _, err := dec.Token(); err != nil {
		return nil, err
	}
	for _, node := range gd.TargetExports {
		if localID := propString(node, "localDefinitionNodeId"); localID != "" {
			if _, ok := gd.AllNodeIDs[localID]; !ok {
				gd.OrphanLocalDefinitionRefs++
			}
		}
	}
	return gd, nil
}

func compareLadybug(path string, targetPaths map[string]struct{}, gd *graphData, definitions map[string][]graph.Node) (ladybugParity, error) {
	p := ladybugParity{
		PerExportDifferenceCount: map[string]int{}, PerDefinitionDifference: map[string]int{},
		LadybugRelationCountByID: map[string]int{}, GraphRelationCountByID: map[string]int{}, ReaderOpenedReadOnly: true,
	}
	runner, err := lbugnative.OpenReadRunner(path)
	if err != nil {
		return p, err
	}
	defer runner.Close()
	where := pathWhere("e", targetPaths)
	exportFields := []string{"id", "name", "filePath", "fileHash", "kind", "exportedName", "localName", "localDefId", "localDefinitionNodeId", "targetRaw", "targetExportedName", "meanings", "typeOnly", "startLine", "startCol", "endLine", "endCol", "selectionStartLine", "selectionStartCol", "selectionEndLine", "selectionEndCol", "statementStartLine", "statementStartCol", "statementEndLine", "statementEndCol", "siteKind", "appLayer", "functionalArea"}
	returns := make([]string, 0, len(exportFields))
	for _, field := range exportFields {
		returns = append(returns, "e."+field+" AS "+field)
	}
	rows, err := runner.QueryRows("MATCH (e:Export) WHERE " + where + " RETURN " + strings.Join(returns, ", ") + " ORDER BY e.id")
	if err != nil {
		return p, err
	}
	p.ExportRowCount = len(rows)
	rowByID := map[string]map[string]any{}
	for _, row := range rows {
		rowByID[fmt.Sprint(row["id"])] = row
	}
	graphByID := map[string]graph.Node{}
	for _, node := range gd.TargetExports {
		graphByID[node.ID] = node
		row, ok := rowByID[node.ID]
		if !ok {
			p.ExportIDsMissingInLadybug = append(p.ExportIDsMissingInLadybug, node.ID)
			continue
		}
		for _, field := range exportFields {
			p.ExportFieldComparisons++
			if !ladybugValueMatches(field, graphField(node, field), fmt.Sprint(row[field])) {
				p.ExportFieldDifferences++
				p.PerExportDifferenceCount[node.ID]++
			}
		}
	}
	for id := range rowByID {
		if _, ok := graphByID[id]; !ok {
			p.ExportIDsMissingInJSON = append(p.ExportIDsMissingInJSON, id)
		}
	}
	sort.Strings(p.ExportIDsMissingInLadybug)
	sort.Strings(p.ExportIDsMissingInJSON)

	definitionByID := map[string]graph.Node{}
	for _, matches := range definitions {
		for _, node := range matches {
			definitionByID[node.ID] = node
		}
	}
	byLabel := map[string][]graph.Node{}
	for _, node := range definitionByID {
		byLabel[string(node.Label)] = append(byLabel[string(node.Label)], node)
	}
	for label, nodes := range byLabel {
		conditions := make([]string, 0, len(nodes))
		for _, node := range nodes {
			conditions = append(conditions, "n.id = "+cypherQuote(node.ID))
		}
		q := "MATCH (n:" + label + ") WHERE " + strings.Join(conditions, " OR ") + " RETURN n.id AS id, n.isExported AS isExported ORDER BY n.id"
		defRows, err := runner.QueryRows(q)
		if err != nil {
			return p, err
		}
		actual := map[string]string{}
		for _, row := range defRows {
			actual[fmt.Sprint(row["id"])] = fmt.Sprint(row["isExported"])
		}
		for _, node := range nodes {
			p.DefinitionComparisons++
			value, ok := actual[node.ID]
			if !ok || !ladybugValueMatches("isExported", propBool(node, "isExported"), value) {
				p.DefinitionDifferences++
				p.PerDefinitionDifference[node.ID]++
			}
		}
	}

	graphRelations := map[string]struct{}{}
	for _, rel := range gd.RelevantRelationships {
		if string(rel.Type) == "CONTAINS" {
			if node, ok := graphByID[rel.TargetID]; ok && propString(node, "filePath") != "" {
				key := rel.SourceID + "\x00" + rel.TargetID + "\x00" + string(rel.Type)
				graphRelations[key] = struct{}{}
				p.GraphRelationCountByID[rel.TargetID]++
			}
		}
	}
	relRows, err := runner.QueryRows("MATCH (f:File)-[r]->(e:Export) WHERE (" + where + ") AND r.type = 'CONTAINS' RETURN f.id AS sourceId, e.id AS targetId, r.type AS type ORDER BY e.id")
	if err != nil {
		return p, err
	}
	p.RelationshipRowCount = len(relRows)
	ladybugRelations := map[string]struct{}{}
	for _, row := range relRows {
		sourceID := fmt.Sprint(row["sourceId"])
		targetID := fmt.Sprint(row["targetId"])
		relType := fmt.Sprint(row["type"])
		key := sourceID + "\x00" + targetID + "\x00" + relType
		ladybugRelations[key] = struct{}{}
		p.LadybugRelationCountByID[targetID]++
	}
	for key := range graphRelations {
		if _, ok := ladybugRelations[key]; !ok {
			p.RelationshipDifferences++
		}
	}
	for key := range ladybugRelations {
		if _, ok := graphRelations[key]; !ok {
			p.RelationshipDifferences++
		}
	}
	return p, nil
}

func compareAffectedReader(paths map[string]struct{}, gd *graphData) map[string]readerFileResult {
	g := &graph.Graph{Nodes: gd.RelevantNodes, Relationships: gd.RelevantRelationships}
	builder := filecontext.NewBuilder(g)
	out := map[string]readerFileResult{}
	for path := range paths {
		ctx, ok := builder.BuildFileContext(path, filecontext.Options{})
		entry := readerFileResult{Found: ok, SymbolExported: map[string]bool{}}
		if ok {
			entry.ExportedSymbolCount = ctx.Summary.ExportedSymbolCount
			for _, node := range ctx.SymbolTree {
				flattenReader(node, entry.SymbolExported)
			}
		}
		out[path] = entry
	}
	return out
}

func flattenReader(node filecontext.SymbolTreeNode, out map[string]bool) {
	out[node.ID] = node.Exported
	for _, child := range node.Children {
		flattenReader(child, out)
	}
}

func evaluateRow(row oracleRow, definitions []graph.Node, gd *graphData, parity ladybugParity, reader map[string]readerFileResult) rowResult {
	r := rowResult{RowID: row.RowID, Polarity: row.Polarity, Path: row.Source.Path, LocalName: row.LocalName, DefinitionMatchCount: len(definitions), Fields: map[string]fieldCheck{}, Pass: true}
	check := func(name string, expected, actual any, pass bool) {
		r.Fields[name] = fieldCheck{Expected: expected, Actual: actual, Pass: pass}
		if !pass {
			r.Pass = false
			r.Failures = append(r.Failures, name)
		}
	}
	check("definitionMatchCount", 1, len(definitions), len(definitions) == 1)
	var def graph.Node
	if len(definitions) == 1 {
		def = definitions[0]
		r.DefinitionID = def.ID
		r.DefinitionLabel = string(def.Label)
		check("declarationKind", row.DeclarationKind, normalizedDeclarationKind(def), normalizedDeclarationKind(def) == row.DeclarationKind)
		_, accessPresent := def.Properties["access"]
		check("access", row.Expected.Access.State, map[string]any{"state": stateFromPresence(accessPresent), "value": def.Properties["access"]}, !accessPresent && row.Expected.Access.State == "absent")
		actualExported := propBool(def, "isExported")
		check("compatibility.isExported", row.Expected.Compatibility.IsExported, actualExported, actualExported == row.Expected.Compatibility.IsExported)
		readerEntry := reader[row.Source.Path]
		readerExported, found := readerEntry.SymbolExported[def.ID]
		r.ReaderSymbolFound = found
		r.ReaderExported = readerExported
		check("reader.symbolFound", true, found, found)
		check("reader.exported", row.Expected.Compatibility.IsExported, readerExported, found && readerExported == row.Expected.Compatibility.IsExported)
		check("ladybug.definitionCompatibilityParity", 0, parity.PerDefinitionDifference[def.ID], parity.PerDefinitionDifference[def.ID] == 0)
	}

	candidates := []graph.Node{}
	for _, export := range gd.TargetExports {
		linked := def.ID != "" && propString(export, "localDefinitionNodeId") == def.ID
		site := propString(export, "filePath") == row.Source.Path && propString(export, "localName") == row.LocalName && intProp(export, "selectionStartLine") == row.Source.SelectionRange.Start.Line && intProp(export, "selectionStartCol") == row.Source.SelectionRange.Start.Column-1 && intProp(export, "selectionEndLine") == row.Source.SelectionRange.End.Line && intProp(export, "selectionEndCol") == row.Source.SelectionRange.End.Column-1
		if linked || site {
			candidates = append(candidates, export)
		}
	}
	for _, export := range candidates {
		r.ExportNodeIDs = append(r.ExportNodeIDs, export.ID)
		r.GraphRelationCount += parity.GraphRelationCountByID[export.ID]
		r.LadybugRelationCount += parity.LadybugRelationCountByID[export.ID]
		r.GraphLadybugDifference += parity.PerExportDifferenceCount[export.ID]
	}
	sort.Strings(r.ExportNodeIDs)
	check("exportFactCount", row.Expected.ExportFactCount, len(candidates), len(candidates) == row.Expected.ExportFactCount)
	check("fileToExportRelationCount", row.Expected.FileToExportRelationCount, r.GraphRelationCount, r.GraphRelationCount == row.Expected.FileToExportRelationCount)
	check("ladybug.fileToExportRelationCount", row.Expected.FileToExportRelationCount, r.LadybugRelationCount, r.LadybugRelationCount == row.Expected.FileToExportRelationCount)
	check("graphLadybug.exportFieldDifferences", 0, r.GraphLadybugDifference, r.GraphLadybugDifference == 0)

	if row.Expected.ExportFactCount == 1 && len(candidates) == 1 {
		export := candidates[0]
		check("exportedName", row.Expected.ExportedName, propString(export, "exportedName"), propString(export, "exportedName") == row.Expected.ExportedName)
		check("name", row.Expected.ExportedName, propString(export, "name"), propString(export, "name") == row.Expected.ExportedName)
		check("localName", row.LocalName, propString(export, "localName"), propString(export, "localName") == row.LocalName)
		actualKind := map[string]string{"kind": propString(export, "kind"), "siteKind": propString(export, "siteKind"), "normalized": normalizedExportKind(export)}
		check("exportKind", row.Expected.ExportKind, actualKind, normalizedExportKind(export) == row.Expected.ExportKind)
		check("meanings", row.Expected.Meanings, propStrings(export, "meanings"), reflect.DeepEqual(row.Expected.Meanings, propStrings(export, "meanings")))
		check("typeOnly", row.Expected.TypeOnly, propBool(export, "typeOnly"), propBool(export, "typeOnly") == row.Expected.TypeOnly)
		check("fileHash", strings.ToUpper(row.Source.FileSHA256), strings.ToUpper(propString(export, "fileHash")), strings.EqualFold(row.Source.FileSHA256, propString(export, "fileHash")))
		statement := map[string]int{"startLine": intProp(export, "statementStartLine"), "startColumn": intProp(export, "statementStartCol") + 1, "endLine": intProp(export, "statementEndLine"), "endColumn": intProp(export, "statementEndCol") + 1}
		expectedStatement := map[string]int{"startLine": row.Source.Range.Start.Line, "startColumn": row.Source.Range.Start.Column, "endLine": row.Source.Range.End.Line, "endColumn": row.Source.Range.End.Column}
		check("sourceRange", expectedStatement, statement, reflect.DeepEqual(expectedStatement, statement))
		selection := map[string]int{"startLine": intProp(export, "selectionStartLine"), "startColumn": intProp(export, "selectionStartCol") + 1, "endLine": intProp(export, "selectionEndLine"), "endColumn": intProp(export, "selectionEndCol") + 1}
		expectedSelection := map[string]int{"startLine": row.Source.SelectionRange.Start.Line, "startColumn": row.Source.SelectionRange.Start.Column, "endLine": row.Source.SelectionRange.End.Line, "endColumn": row.Source.SelectionRange.End.Column}
		check("selectionRange", expectedSelection, selection, reflect.DeepEqual(expectedSelection, selection))
		check("localDefinitionNodeId", def.ID, propString(export, "localDefinitionNodeId"), def.ID != "" && propString(export, "localDefinitionNodeId") == def.ID)
		_, exportAccessPresent := export.Properties["access"]
		check("exportAccess", "absent", stateFromPresence(exportAccessPresent), !exportAccessPresent)
		forbidden := []string{}
		for key := range export.Properties {
			if forbiddenKey(key) {
				forbidden = append(forbidden, key)
			}
		}
		sort.Strings(forbidden)
		check("child05DerivedState", "empty", forbidden, len(forbidden) == 0 && propString(export, "targetRaw") == "" && propString(export, "targetExportedName") == "")
	}
	check("exportDiagnosticCount", row.Expected.ExportDiagnosticCount, gd.ExportDiagnosticCount, gd.ExportDiagnosticCount == row.Expected.ExportDiagnosticCount)
	return r
}

func matchingDefinitions(row oracleRow, nodes []graph.Node) []graph.Node {
	out := []graph.Node{}
	for _, node := range nodes {
		if propString(node, "filePath") != row.Source.Path || propString(node, "name") != row.LocalName {
			continue
		}
		if normalizedDeclarationKind(node) != row.DeclarationKind {
			continue
		}
		if intProp(node, "selectionStartLine") != row.Source.SelectionRange.Start.Line || intProp(node, "selectionStartCol") != row.Source.SelectionRange.Start.Column-1 || intProp(node, "selectionEndLine") != row.Source.SelectionRange.End.Line || intProp(node, "selectionEndCol") != row.Source.SelectionRange.End.Column-1 {
			continue
		}
		out = append(out, node)
	}
	return out
}

func normalizedDeclarationKind(node graph.Node) string {
	switch string(node.Label) {
	case "TypeAlias":
		return "type_alias"
	case "Function":
		return "function"
	case "Variable", "Const":
		return "const_binding"
	default:
		return strings.ToLower(string(node.Label))
	}
}

func normalizedExportKind(node graph.Node) string {
	if propString(node, "kind") == "direct" && propString(node, "siteKind") == "export_declaration" {
		return "direct_declaration"
	}
	return propString(node, "kind")
}

func forbiddenKey(key string) bool {
	normalized := strings.NewReplacer("_", "", "-", "").Replace(strings.ToLower(key))
	return strings.Contains(normalized, "terminal") || strings.Contains(normalized, "resolvedtarget") || strings.Contains(normalized, "publicapi") || strings.Contains(normalized, "packageapi")
}

func propString(node graph.Node, key string) string {
	if node.Properties == nil {
		return ""
	}
	value, _ := node.Properties[key].(string)
	return value
}

func propBool(node graph.Node, key string) bool {
	if node.Properties == nil {
		return false
	}
	value, _ := node.Properties[key].(bool)
	return value
}

func intProp(node graph.Node, key string) int {
	if node.Properties == nil {
		return 0
	}
	switch value := node.Properties[key].(type) {
	case float64:
		return int(value)
	case int:
		return value
	case json.Number:
		n, _ := value.Int64()
		return int(n)
	}
	return 0
}

func propStrings(node graph.Node, key string) []string {
	if node.Properties == nil {
		return nil
	}
	switch value := node.Properties[key].(type) {
	case []string:
		return append([]string(nil), value...)
	case []any:
		out := make([]string, 0, len(value))
		for _, item := range value {
			if text, ok := item.(string); ok {
				out = append(out, text)
			}
		}
		return out
	}
	return nil
}

func graphField(node graph.Node, field string) any {
	if field == "id" {
		return node.ID
	}
	return node.Properties[field]
}

func ladybugValueMatches(field string, expected any, actual string) bool {
	if field == "meanings" {
		return reflect.DeepEqual(anyStrings(expected), parseArray(actual))
	}
	switch value := expected.(type) {
	case bool:
		parsed, err := strconv.ParseBool(strings.ToLower(strings.TrimSpace(actual)))
		return err == nil && parsed == value
	case float64:
		parsed, err := strconv.ParseInt(strings.TrimSpace(actual), 10, 64)
		return err == nil && int64(value) == parsed
	case int:
		parsed, err := strconv.ParseInt(strings.TrimSpace(actual), 10, 64)
		return err == nil && int64(value) == parsed
	case nil:
		return actual == "" || strings.EqualFold(actual, "null")
	case string:
		return value == actual
	default:
		return fmt.Sprint(value) == actual
	}
}

func anyStrings(value any) []string {
	switch values := value.(type) {
	case []string:
		return values
	case []any:
		out := make([]string, 0, len(values))
		for _, item := range values {
			out = append(out, fmt.Sprint(item))
		}
		return out
	}
	return nil
}

func parseArray(value string) []string {
	value = strings.TrimSpace(value)
	if value == "" || value == "[]" {
		return []string{}
	}
	var decoded []string
	if json.Unmarshal([]byte(value), &decoded) == nil {
		return decoded
	}
	value = strings.TrimPrefix(strings.TrimSuffix(value, "]"), "[")
	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.Trim(strings.TrimSpace(part), "\"'")
		if part != "" {
			out = append(out, part)
		}
	}
	return out
}

func pathWhere(alias string, paths map[string]struct{}) string {
	ordered := make([]string, 0, len(paths))
	for path := range paths {
		ordered = append(ordered, path)
	}
	sort.Strings(ordered)
	conditions := make([]string, 0, len(ordered))
	for _, path := range ordered {
		conditions = append(conditions, alias+".filePath = "+cypherQuote(path))
	}
	return strings.Join(conditions, " OR ")
}

func cypherQuote(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "\\'") + "'"
}

func stateFromPresence(present bool) string {
	if present {
		return "present"
	}
	return "absent"
}
