package tsstdlib

import (
	"bytes"
	"crypto/sha256"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
	"sync"
)

const (
	SchemaVersion              = "tsstdlib.catalog.v1"
	AuthorityKind              = "typescript_standard_library"
	IdentityVersion            = "tsstdlib.semantic.v1"
	TypeScriptVersion          = "5.9.3"
	TypeScriptPackageIntegrity = "sha512-jl1vZzPDinLr9eUt3J/t7V6FgNEw9QjvBPdysz9KfQDD41fQrC2Y4vKQdiaUpFT4bXlb1RHhLpp8wtm6M5TgSw=="
	GenerationCommand          = "node anvien-web/scripts/generate-tsstdlib-catalog.mjs"
)

//go:embed catalog.v1.json
var embeddedCatalog []byte

type Meaning string

const (
	MeaningValue     Meaning = "value"
	MeaningType      Meaning = "type"
	MeaningNamespace Meaning = "namespace"
)

type LookupStatus string

const (
	LookupResolved              LookupStatus = "resolved"
	LookupNotFound              LookupStatus = "not_found"
	LookupProfileExcluded       LookupStatus = "profile_excluded"
	LookupMeaningMismatch       LookupStatus = "meaning_mismatch"
	LookupCapabilityUnavailable LookupStatus = "capability_unavailable"
)

type Reason string

const (
	ReasonDisabledByNoLib      Reason = "disabled_by_no_lib"
	ReasonProfileExcludes      Reason = "profile_excludes_declaration"
	ReasonMeaningMismatch      Reason = "meaning_mismatch"
	ReasonConfigInvalid        Reason = "config_invalid"
	ReasonConfigTopology       Reason = "config_topology_unsupported"
	ReasonConfigUnreadable     Reason = "config_unreadable"
	ReasonCatalogMissing       Reason = "catalog_missing"
	ReasonCatalogSchema        Reason = "catalog_schema_unsupported"
	ReasonCatalogVersion       Reason = "catalog_version_mismatch"
	ReasonCatalogHash          Reason = "catalog_hash_mismatch"
	ReasonCatalogInputManifest Reason = "catalog_input_manifest_mismatch"
)

// CatalogProofState makes the provenance/absence contract explicit. Ready
// results carry validated authority, logical-catalog, and artifact hashes.
// Missing results carry none of those hashes. Rejected results carry only the
// hash of the attempted artifact; authority and logical-catalog identity are
// deliberately absent because validation did not establish them.
type CatalogProofState string

const (
	CatalogProofReady    CatalogProofState = "ready"
	CatalogProofMissing  CatalogProofState = "missing"
	CatalogProofRejected CatalogProofState = "rejected"
)

type ProfileStatus string

const (
	ProfileReady       ProfileStatus = "ready"
	ProfileUnavailable ProfileStatus = "unavailable"
)

type Declaration struct {
	Library   string `json:"library"`
	StartLine int    `json:"startLine"`
	StartCol  int    `json:"startCol"`
	EndLine   int    `json:"endLine"`
	EndCol    int    `json:"endCol"`
}

type LookupResult struct {
	Status              LookupStatus      `json:"status"`
	Reason              Reason            `json:"reason,omitempty"`
	SymbolID            string            `json:"symbolId,omitempty"`
	OwnerID             string            `json:"ownerId,omitempty"`
	Name                string            `json:"name"`
	Meaning             Meaning           `json:"meaning"`
	Declarations        []Declaration     `json:"declarations,omitempty"`
	AuthorityKind       string            `json:"authorityKind"`
	CatalogProofState   CatalogProofState `json:"catalogProofState"`
	AuthorityHash       string            `json:"authorityHash,omitempty"`
	TypeScriptVersion   string            `json:"typescriptVersion"`
	CatalogHash         string            `json:"catalogHash,omitempty"`
	CatalogArtifactHash string            `json:"catalogArtifactHash,omitempty"`
	ProfileHash         string            `json:"profileHash,omitempty"`
	ConfigHash          string            `json:"configHash,omitempty"`
}

func (result LookupResult) Handled() bool {
	return result.Status != LookupNotFound
}

type Profile struct {
	Status      ProfileStatus `json:"status"`
	Reason      Reason        `json:"reason,omitempty"`
	Target      string        `json:"target,omitempty"`
	Libraries   []string      `json:"libraries,omitempty"`
	ConfigPath  string        `json:"configPath,omitempty"`
	ConfigHash  string        `json:"configHash"`
	ProfileHash string        `json:"profileHash"`

	librarySet map[int]struct{}
}

func (profile Profile) clone() Profile {
	profile.Libraries = append([]string(nil), profile.Libraries...)
	profile.librarySet = nil
	return profile
}

type Metadata struct {
	SchemaVersion       string `json:"schemaVersion"`
	AuthorityKind       string `json:"authorityKind"`
	IdentityVersion     string `json:"identityVersion"`
	AuthorityHash       string `json:"authorityHash"`
	TypeScriptVersion   string `json:"typescriptVersion"`
	PackageIntegrity    string `json:"packageIntegrity"`
	GenerationCommand   string `json:"generationCommand"`
	CatalogHash         string `json:"catalogHash"`
	CatalogArtifactHash string `json:"catalogArtifactHash"`
	InputCount          int    `json:"inputCount"`
	InputBytes          int    `json:"inputBytes"`
	SymbolCount         int    `json:"symbolCount"`
	MemberCount         int    `json:"memberCount"`
}

type Authority struct {
	catalog             *catalogIndex
	profile             Profile
	catalogProofState   CatalogProofState
	catalogArtifactHash string
}

func NewAuthority(repoRoot string, inventory []string) *Authority {
	catalog, reason, err := loadDefaultCatalog()
	if err != nil {
		return unavailableCatalogAuthority(repoRoot, inventory, embeddedCatalog, reason)
	}
	return readyCatalogAuthority(repoRoot, inventory, catalog)
}

// NewAuthorityFromCatalog validates an offline compact catalog artifact before
// constructing an authority. It exists so callers at the resolver/analyzer
// boundary can exercise the same fail-closed carriage contract without
// replacing the checked-in embedded production artifact.
func NewAuthorityFromCatalog(repoRoot string, inventory []string, raw []byte) *Authority {
	catalog, err := loadCatalog(raw)
	if err != nil {
		return unavailableCatalogAuthority(repoRoot, inventory, raw, catalogReason(err))
	}
	return readyCatalogAuthority(repoRoot, inventory, catalog)
}

func readyCatalogAuthority(repoRoot string, inventory []string, catalog *catalogIndex) *Authority {
	return &Authority{
		catalog:             catalog,
		profile:             selectProfile(catalog, repoRoot, inventory),
		catalogProofState:   CatalogProofReady,
		catalogArtifactHash: catalog.artifactHash,
	}
}

func unavailableCatalogAuthority(_ string, inventory []string, raw []byte, reason Reason) *Authority {
	proofState := CatalogProofRejected
	artifactHash := ""
	if reason == ReasonCatalogMissing {
		proofState = CatalogProofMissing
	} else if len(raw) > 0 {
		artifactHash = sha256Hex(raw)
	}
	// Catalog validation precedes profile selection. The complete scanner
	// inventory is therefore the strongest safe configuration proof available;
	// profileDigest commits it together with the exact typed failure reason.
	return &Authority{
		profile:             unavailableProfile(reason, "", inventoryHash(inventory)),
		catalogProofState:   proofState,
		catalogArtifactHash: artifactHash,
	}
}

func (authority *Authority) Profile() Profile {
	if authority == nil {
		return Profile{Status: ProfileUnavailable, Reason: ReasonCatalogMissing}
	}
	return authority.profile.clone()
}

func (authority *Authority) Metadata() Metadata {
	if authority == nil || authority.catalog == nil {
		return Metadata{
			SchemaVersion:     SchemaVersion,
			AuthorityKind:     AuthorityKind,
			IdentityVersion:   IdentityVersion,
			TypeScriptVersion: TypeScriptVersion,
		}
	}
	return authority.catalog.metadata
}

func (authority *Authority) LookupGlobal(name string, meaning Meaning) LookupResult {
	name = strings.TrimSpace(name)
	result := authority.baseResult(name, meaning)
	if status, reason := authority.availability(); status != LookupResolved {
		result.Status = status
		result.Reason = reason
		return result
	}
	symbol, ok := authority.catalog.symbols[name]
	if !ok {
		result.Status = LookupNotFound
		return result
	}
	refs := symbol.lane(meaning)
	if len(refs) == 0 {
		result.Status = LookupMeaningMismatch
		result.Reason = ReasonMeaningMismatch
		return result
	}
	declarations := authority.activeDeclarations(refs)
	if len(declarations) == 0 {
		result.Status = LookupProfileExcluded
		result.Reason = ReasonProfileExcludes
		return result
	}
	result.Status = LookupResolved
	result.SymbolID = symbol.id(meaning)
	result.Declarations = declarations
	return result
}

func (authority *Authority) LookupMember(ownerName string, ownerMeaning Meaning, memberName string, memberMeaning Meaning) LookupResult {
	ownerName = strings.TrimSpace(ownerName)
	memberName = strings.TrimSpace(memberName)
	qualifiedName := memberName
	if ownerName != "" && memberName != "" {
		qualifiedName = ownerName + "." + memberName
	}
	result := authority.baseResult(qualifiedName, memberMeaning)
	owner := authority.LookupGlobal(ownerName, ownerMeaning)
	if owner.Status != LookupResolved {
		result.Status = owner.Status
		result.Reason = owner.Reason
		return result
	}
	result.OwnerID = owner.SymbolID
	resolved := authority.lookupMember(
		ownerName,
		ownerMeaning,
		memberName,
		memberMeaning,
		make(map[string]struct{}),
	)
	result.Status = resolved.status
	result.Reason = resolved.reason
	if resolved.ownerID != "" {
		result.OwnerID = resolved.ownerID
	}
	result.SymbolID = resolved.member.id(memberMeaning)
	result.Declarations = resolved.declarations
	return result
}

type memberLookup struct {
	status       LookupStatus
	reason       Reason
	ownerID      string
	member       memberDTO
	declarations []Declaration
}

func (authority *Authority) lookupMember(ownerName string, ownerMeaning Meaning, memberName string, memberMeaning Meaning, visited map[string]struct{}) memberLookup {
	key := ownerName + "\x00" + string(ownerMeaning)
	if _, seen := visited[key]; seen {
		return memberLookup{status: LookupNotFound}
	}
	visited[key] = struct{}{}

	symbol, ok := authority.catalog.symbols[ownerName]
	if !ok {
		return memberLookup{status: LookupNotFound}
	}
	best := memberLookup{status: LookupNotFound}
	if member, found := findMember(symbol.members(ownerMeaning), memberName); found {
		refs := member.lane(memberMeaning)
		switch {
		case len(refs) == 0:
			best = preferMemberLookup(best, memberLookup{status: LookupMeaningMismatch, reason: ReasonMeaningMismatch})
		default:
			declarations := authority.activeDeclarations(refs)
			if len(declarations) > 0 {
				return memberLookup{
					status:       LookupResolved,
					ownerID:      symbol.id(ownerMeaning),
					member:       member,
					declarations: declarations,
				}
			}
			best = preferMemberLookup(best, memberLookup{status: LookupProfileExcluded, reason: ReasonProfileExcludes})
		}
	}

	for _, parent := range symbol.parents(ownerMeaning) {
		global := authority.LookupGlobal(parent.name, parent.meaning)
		switch global.Status {
		case LookupResolved:
			candidate := authority.lookupMember(parent.name, parent.meaning, memberName, memberMeaning, visited)
			if candidate.status == LookupResolved {
				return candidate
			}
			best = preferMemberLookup(best, candidate)
		case LookupCapabilityUnavailable:
			return memberLookup{status: global.Status, reason: global.Reason}
		case LookupProfileExcluded:
			best = preferMemberLookup(best, memberLookup{status: global.Status, reason: global.Reason})
		case LookupMeaningMismatch:
			best = preferMemberLookup(best, memberLookup{status: global.Status, reason: global.Reason})
		}
	}

	if symbol.hasMemberInOtherLane(ownerMeaning, memberName) {
		best = preferMemberLookup(best, memberLookup{status: LookupMeaningMismatch, reason: ReasonMeaningMismatch})
	}
	return best
}

type memberParent struct {
	name    string
	meaning Meaning
}

func preferMemberLookup(current memberLookup, candidate memberLookup) memberLookup {
	priority := func(status LookupStatus) int {
		switch status {
		case LookupCapabilityUnavailable:
			return 4
		case LookupProfileExcluded:
			return 3
		case LookupMeaningMismatch:
			return 2
		default:
			return 1
		}
	}
	if priority(candidate.status) > priority(current.status) {
		return candidate
	}
	return current
}

func (authority *Authority) availability() (LookupStatus, Reason) {
	if authority == nil {
		return LookupCapabilityUnavailable, ReasonCatalogMissing
	}
	if authority.profile.Status != ProfileReady {
		if authority.profile.Reason == "" {
			return LookupCapabilityUnavailable, ReasonCatalogMissing
		}
		return LookupCapabilityUnavailable, authority.profile.Reason
	}
	if authority.catalog == nil {
		return LookupCapabilityUnavailable, ReasonCatalogMissing
	}
	return LookupResolved, ""
}

func (authority *Authority) baseResult(name string, meaning Meaning) LookupResult {
	result := LookupResult{
		Status:            LookupNotFound,
		Name:              name,
		Meaning:           meaning,
		AuthorityKind:     AuthorityKind,
		CatalogProofState: CatalogProofMissing,
		TypeScriptVersion: TypeScriptVersion,
	}
	if authority == nil {
		return result
	}
	result.ConfigHash = authority.profile.ConfigHash
	result.ProfileHash = authority.profile.ProfileHash
	result.CatalogProofState = authority.catalogProofState
	result.CatalogArtifactHash = authority.catalogArtifactHash
	if authority.catalog != nil {
		result.AuthorityHash = authority.catalog.authorityHash
		result.CatalogHash = authority.catalog.dto.Hash
	}
	return result
}

func (authority *Authority) activeDeclarations(refs []declarationDTO) []Declaration {
	out := make([]Declaration, 0, len(refs))
	for _, ref := range refs {
		if _, ok := authority.profile.librarySet[ref[0]]; !ok {
			continue
		}
		input := authority.catalog.dto.Inputs[ref[0]]
		out = append(out, Declaration{
			Library:   input.Path,
			StartLine: ref[1],
			StartCol:  ref[2],
			EndLine:   ref[3],
			EndCol:    ref[4],
		})
	}
	return out
}

type declarationDTO [5]int

type inputDTO struct {
	Path       string `json:"p"`
	Bytes      int    `json:"b"`
	SHA256     string `json:"h"`
	References []int  `json:"r,omitempty"`
}

type memberDTO struct {
	Name        string           `json:"n"`
	Value       []declarationDTO `json:"v,omitempty"`
	Type        []declarationDTO `json:"t,omitempty"`
	Namespace   []declarationDTO `json:"s,omitempty"`
	ValueID     string           `json:"iv,omitempty"`
	TypeID      string           `json:"it,omitempty"`
	NamespaceID string           `json:"is,omitempty"`
}

func (member memberDTO) id(meaning Meaning) string {
	switch meaning {
	case MeaningValue:
		return member.ValueID
	case MeaningType:
		return member.TypeID
	case MeaningNamespace:
		return member.NamespaceID
	default:
		return ""
	}
}

func (member memberDTO) lane(meaning Meaning) []declarationDTO {
	switch meaning {
	case MeaningValue:
		return member.Value
	case MeaningType:
		return member.Type
	case MeaningNamespace:
		return member.Namespace
	default:
		return nil
	}
}

type symbolDTO struct {
	Name             string           `json:"n"`
	Value            []declarationDTO `json:"v,omitempty"`
	Type             []declarationDTO `json:"t,omitempty"`
	Namespace        []declarationDTO `json:"s,omitempty"`
	BaseTypes        []string         `json:"bt,omitempty"`
	ValueOwners      []string         `json:"ov,omitempty"`
	ValueMembers     []memberDTO      `json:"mv,omitempty"`
	TypeMembers      []memberDTO      `json:"mt,omitempty"`
	NamespaceMembers []memberDTO      `json:"ms,omitempty"`
	ValueID          string           `json:"iv,omitempty"`
	TypeID           string           `json:"it,omitempty"`
	NamespaceID      string           `json:"is,omitempty"`
}

func (symbol symbolDTO) id(meaning Meaning) string {
	switch meaning {
	case MeaningValue:
		return symbol.ValueID
	case MeaningType:
		return symbol.TypeID
	case MeaningNamespace:
		return symbol.NamespaceID
	default:
		return ""
	}
}

func (symbol symbolDTO) lane(meaning Meaning) []declarationDTO {
	switch meaning {
	case MeaningValue:
		return symbol.Value
	case MeaningType:
		return symbol.Type
	case MeaningNamespace:
		return symbol.Namespace
	default:
		return nil
	}
}

func (symbol symbolDTO) members(meaning Meaning) []memberDTO {
	switch meaning {
	case MeaningValue:
		return symbol.ValueMembers
	case MeaningType:
		return symbol.TypeMembers
	case MeaningNamespace:
		return symbol.NamespaceMembers
	default:
		return nil
	}
}

func (symbol symbolDTO) parents(meaning Meaning) []memberParent {
	var names []string
	parentMeaning := MeaningType
	switch meaning {
	case MeaningValue:
		names = symbol.ValueOwners
	case MeaningType:
		names = symbol.BaseTypes
	default:
		return nil
	}
	out := make([]memberParent, 0, len(names))
	for _, name := range names {
		out = append(out, memberParent{name: name, meaning: parentMeaning})
	}
	return out
}

func (symbol symbolDTO) hasMemberInOtherLane(selected Meaning, name string) bool {
	for _, meaning := range []Meaning{MeaningValue, MeaningType, MeaningNamespace} {
		if meaning == selected {
			continue
		}
		if _, ok := findMember(symbol.members(meaning), name); ok {
			return true
		}
	}
	return false
}

func findMember(members []memberDTO, name string) (memberDTO, bool) {
	index := sort.Search(len(members), func(index int) bool { return members[index].Name >= name })
	if index >= len(members) || members[index].Name != name {
		return memberDTO{}, false
	}
	return members[index], true
}

type catalogDTO struct {
	Schema          string         `json:"schema"`
	AuthorityKind   string         `json:"authority"`
	IdentityVersion string         `json:"identity"`
	TypeScript      string         `json:"ts"`
	Integrity       string         `json:"integrity"`
	Generator       string         `json:"generator"`
	Hash            string         `json:"hash"`
	Inputs          []inputDTO     `json:"inputs"`
	Aliases         map[string]int `json:"aliases"`
	Targets         map[string]int `json:"targets"`
	Symbols         []symbolDTO    `json:"symbols"`
}

type catalogIndex struct {
	dto           catalogDTO
	symbols       map[string]symbolDTO
	artifactHash  string
	authorityHash string
	metadata      Metadata
}

type catalogError struct {
	reason Reason
	err    error
}

func (err *catalogError) Error() string {
	return err.err.Error()
}

func (err *catalogError) Unwrap() error {
	return err.err
}

var (
	defaultCatalogOnce sync.Once
	defaultCatalog     *catalogIndex
	defaultCatalogErr  error
)

func loadDefaultCatalog() (*catalogIndex, Reason, error) {
	defaultCatalogOnce.Do(func() {
		defaultCatalog, defaultCatalogErr = loadCatalog(embeddedCatalog)
	})
	if defaultCatalogErr == nil {
		return defaultCatalog, "", nil
	}
	return nil, catalogReason(defaultCatalogErr), defaultCatalogErr
}

func catalogReason(err error) Reason {
	var typed *catalogError
	if errors.As(err, &typed) {
		return typed.reason
	}
	return ReasonCatalogInputManifest
}

func loadCatalog(raw []byte) (*catalogIndex, error) {
	if len(raw) == 0 {
		return nil, newCatalogError(ReasonCatalogMissing, "embedded catalog is empty")
	}
	decoder := json.NewDecoder(strings.NewReader(string(raw)))
	decoder.DisallowUnknownFields()
	var dto catalogDTO
	if err := decoder.Decode(&dto); err != nil {
		return nil, newCatalogError(ReasonCatalogInputManifest, "decode catalog: %v", err)
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return nil, newCatalogError(ReasonCatalogInputManifest, "catalog contains trailing data")
	}
	if dto.Schema != SchemaVersion {
		return nil, newCatalogError(ReasonCatalogSchema, "catalog schema %q is unsupported", dto.Schema)
	}
	if dto.AuthorityKind != AuthorityKind ||
		dto.IdentityVersion != IdentityVersion ||
		dto.TypeScript != TypeScriptVersion ||
		dto.Integrity != TypeScriptPackageIntegrity ||
		dto.Generator != GenerationCommand {
		return nil, newCatalogError(ReasonCatalogVersion, "catalog compiler provenance does not match the runtime contract")
	}
	storedHash := dto.Hash
	dto.Hash = ""
	canonical, err := marshalCatalogCanonical(dto)
	if err != nil {
		return nil, newCatalogError(ReasonCatalogInputManifest, "marshal catalog for hash validation: %v", err)
	}
	computedHash := sha256Hex(canonical)
	if storedHash == "" || storedHash != computedHash {
		return nil, newCatalogError(ReasonCatalogHash, "catalog hash mismatch: got %q, want %q", storedHash, computedHash)
	}
	dto.Hash = storedHash
	if err := validateCatalogManifest(dto); err != nil {
		return nil, err
	}
	symbols := make(map[string]symbolDTO, len(dto.Symbols))
	memberCount := 0
	for _, symbol := range dto.Symbols {
		symbols[symbol.Name] = symbol
		memberCount += len(symbol.ValueMembers) + len(symbol.TypeMembers) + len(symbol.NamespaceMembers)
	}
	inputBytes := 0
	for _, input := range dto.Inputs {
		inputBytes += input.Bytes
	}
	artifactHash := sha256Hex(raw)
	authorityHash := authorityDigest(dto, artifactHash)
	return &catalogIndex{
		dto:           dto,
		symbols:       symbols,
		artifactHash:  artifactHash,
		authorityHash: authorityHash,
		metadata: Metadata{
			SchemaVersion:       dto.Schema,
			AuthorityKind:       dto.AuthorityKind,
			IdentityVersion:     dto.IdentityVersion,
			AuthorityHash:       authorityHash,
			TypeScriptVersion:   dto.TypeScript,
			PackageIntegrity:    dto.Integrity,
			GenerationCommand:   dto.Generator,
			CatalogHash:         dto.Hash,
			CatalogArtifactHash: artifactHash,
			InputCount:          len(dto.Inputs),
			InputBytes:          inputBytes,
			SymbolCount:         len(dto.Symbols),
			MemberCount:         memberCount,
		},
	}, nil
}

type logicalMemberDTO struct {
	Name      string           `json:"n"`
	Value     []declarationDTO `json:"v,omitempty"`
	Type      []declarationDTO `json:"t,omitempty"`
	Namespace []declarationDTO `json:"s,omitempty"`
}

type logicalSymbolDTO struct {
	Name             string             `json:"n"`
	Value            []declarationDTO   `json:"v,omitempty"`
	Type             []declarationDTO   `json:"t,omitempty"`
	Namespace        []declarationDTO   `json:"s,omitempty"`
	BaseTypes        []string           `json:"bt,omitempty"`
	ValueOwners      []string           `json:"ov,omitempty"`
	ValueMembers     []logicalMemberDTO `json:"mv,omitempty"`
	TypeMembers      []logicalMemberDTO `json:"mt,omitempty"`
	NamespaceMembers []logicalMemberDTO `json:"ms,omitempty"`
}

type logicalCatalogDTO struct {
	Schema          string             `json:"schema"`
	AuthorityKind   string             `json:"authority"`
	IdentityVersion string             `json:"identity"`
	TypeScript      string             `json:"ts"`
	Integrity       string             `json:"integrity"`
	Generator       string             `json:"generator"`
	Hash            string             `json:"hash"`
	Inputs          []inputDTO         `json:"inputs"`
	Aliases         map[string]int     `json:"aliases"`
	Targets         map[string]int     `json:"targets"`
	Symbols         []logicalSymbolDTO `json:"symbols"`
}

func marshalCatalogCanonical(dto catalogDTO) ([]byte, error) {
	logical := logicalCatalogDTO{
		Schema:          dto.Schema,
		AuthorityKind:   dto.AuthorityKind,
		IdentityVersion: dto.IdentityVersion,
		TypeScript:      dto.TypeScript,
		Integrity:       dto.Integrity,
		Generator:       dto.Generator,
		Hash:            dto.Hash,
		Inputs:          dto.Inputs,
		Aliases:         dto.Aliases,
		Targets:         dto.Targets,
		Symbols:         make([]logicalSymbolDTO, 0, len(dto.Symbols)),
	}
	for _, symbol := range dto.Symbols {
		logical.Symbols = append(logical.Symbols, logicalSymbol(symbol))
	}
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(logical); err != nil {
		return nil, err
	}
	return bytes.TrimSuffix(buffer.Bytes(), []byte{'\n'}), nil
}

func logicalSymbol(symbol symbolDTO) logicalSymbolDTO {
	return logicalSymbolDTO{
		Name:             symbol.Name,
		Value:            symbol.Value,
		Type:             symbol.Type,
		Namespace:        symbol.Namespace,
		BaseTypes:        symbol.BaseTypes,
		ValueOwners:      symbol.ValueOwners,
		ValueMembers:     logicalMembers(symbol.ValueMembers),
		TypeMembers:      logicalMembers(symbol.TypeMembers),
		NamespaceMembers: logicalMembers(symbol.NamespaceMembers),
	}
}

func logicalMembers(members []memberDTO) []logicalMemberDTO {
	out := make([]logicalMemberDTO, 0, len(members))
	for _, member := range members {
		out = append(out, logicalMemberDTO{
			Name:      member.Name,
			Value:     member.Value,
			Type:      member.Type,
			Namespace: member.Namespace,
		})
	}
	return out
}

func validateCatalogManifest(dto catalogDTO) error {
	if len(dto.Inputs) == 0 || len(dto.Symbols) == 0 {
		return newCatalogError(ReasonCatalogInputManifest, "catalog inputs and symbols must be non-empty")
	}
	for index, input := range dto.Inputs {
		if input.Path == "" || input.Bytes <= 0 || !validSHA256(input.SHA256) {
			return newCatalogError(ReasonCatalogInputManifest, "catalog input %d is invalid", index)
		}
		if index > 0 && dto.Inputs[index-1].Path >= input.Path {
			return newCatalogError(ReasonCatalogInputManifest, "catalog inputs are not strictly sorted")
		}
		if err := validateIndexes(input.References, len(dto.Inputs), "input references"); err != nil {
			return err
		}
	}
	for name, index := range dto.Aliases {
		if strings.TrimSpace(name) == "" || index < 0 || index >= len(dto.Inputs) {
			return newCatalogError(ReasonCatalogInputManifest, "catalog alias %q is invalid", name)
		}
	}
	for name, index := range dto.Targets {
		if strings.TrimSpace(name) == "" || index < 0 || index >= len(dto.Inputs) {
			return newCatalogError(ReasonCatalogInputManifest, "catalog target %q is invalid", name)
		}
	}
	knownSymbols := make(map[string]struct{}, len(dto.Symbols))
	for index, symbol := range dto.Symbols {
		if symbol.Name == "" || (index > 0 && dto.Symbols[index-1].Name >= symbol.Name) {
			return newCatalogError(ReasonCatalogInputManifest, "catalog symbol %d ordering is invalid", index)
		}
		knownSymbols[symbol.Name] = struct{}{}
	}
	semanticIDs := make(map[string]string)
	for _, symbol := range dto.Symbols {
		if err := validateDeclarationLanes(len(dto.Inputs), symbol.Value, symbol.Type, symbol.Namespace); err != nil {
			return fmt.Errorf("symbol %q: %w", symbol.Name, err)
		}
		if err := validateSemanticIDs(
			dto,
			[][]string{{"global"}},
			symbol.Name,
			symbol.Value,
			symbol.Type,
			symbol.Namespace,
			symbol.ValueID,
			symbol.TypeID,
			symbol.NamespaceID,
			semanticIDs,
		); err != nil {
			return fmt.Errorf("symbol %q: %w", symbol.Name, err)
		}
		for _, owners := range [][]string{symbol.BaseTypes, symbol.ValueOwners} {
			for index, owner := range owners {
				if owner == "" || (index > 0 && owners[index-1] >= owner) {
					return newCatalogError(ReasonCatalogInputManifest, "catalog owner links are invalid for %q", symbol.Name)
				}
				if _, ok := knownSymbols[owner]; !ok {
					return newCatalogError(ReasonCatalogInputManifest, "catalog owner %q for %q does not exist", owner, symbol.Name)
				}
			}
		}
		memberGroups := []struct {
			meaning Meaning
			members []memberDTO
		}{
			{meaning: MeaningValue, members: symbol.ValueMembers},
			{meaning: MeaningType, members: symbol.TypeMembers},
			{meaning: MeaningNamespace, members: symbol.NamespaceMembers},
		}
		for _, group := range memberGroups {
			members := group.members
			for memberIndex, member := range members {
				if member.Name == "" || (memberIndex > 0 && members[memberIndex-1].Name >= member.Name) {
					return newCatalogError(ReasonCatalogInputManifest, "catalog member ordering is invalid for %q", symbol.Name)
				}
				if err := validateDeclarationLanes(len(dto.Inputs), member.Value, member.Type, member.Namespace); err != nil {
					return fmt.Errorf("member %q.%q: %w", symbol.Name, member.Name, err)
				}
				if err := validateSemanticIDs(
					dto,
					[][]string{{"global", symbol.Name, string(group.meaning)}},
					member.Name,
					member.Value,
					member.Type,
					member.Namespace,
					member.ValueID,
					member.TypeID,
					member.NamespaceID,
					semanticIDs,
				); err != nil {
					return fmt.Errorf("member %q.%q: %w", symbol.Name, member.Name, err)
				}
			}
		}
	}
	return nil
}

func validateSemanticIDs(
	dto catalogDTO,
	ownerPath [][]string,
	name string,
	value []declarationDTO,
	typeLane []declarationDTO,
	namespace []declarationDTO,
	valueID string,
	typeID string,
	namespaceID string,
	seen map[string]string,
) error {
	lanes := []struct {
		meaning Meaning
		refs    []declarationDTO
		id      string
	}{
		{meaning: MeaningValue, refs: value, id: valueID},
		{meaning: MeaningType, refs: typeLane, id: typeID},
		{meaning: MeaningNamespace, refs: namespace, id: namespaceID},
	}
	for _, lane := range lanes {
		if len(lane.refs) == 0 {
			if lane.id != "" {
				return newCatalogError(ReasonCatalogInputManifest, "semantic ID exists without a %s declaration lane", lane.meaning)
			}
			continue
		}
		expected, err := catalogSemanticID(dto, ownerPath, name, lane.meaning, lane.refs)
		if err != nil {
			return err
		}
		identity := strings.Join([]string{strings.Join(ownerPath[len(ownerPath)-1], "/"), name, string(lane.meaning)}, "\x00")
		if previous, duplicate := seen[lane.id]; duplicate {
			return newCatalogError(ReasonCatalogInputManifest, "duplicate semantic ID %q for %s and %s", lane.id, previous, identity)
		}
		if lane.id != expected {
			return newCatalogError(ReasonCatalogInputManifest, "semantic ID component drift for %q/%s", name, lane.meaning)
		}
		seen[lane.id] = identity
	}
	return nil
}

type semanticIdentityPayload struct {
	Authority    string     `json:"a"`
	TypeScript   string     `json:"v"`
	Catalog      string     `json:"c"`
	OwnerPath    [][]string `json:"o"`
	Name         string     `json:"n"`
	Meaning      Meaning    `json:"m"`
	Declarations [][]any    `json:"d"`
}

func catalogSemanticID(dto catalogDTO, ownerPath [][]string, name string, meaning Meaning, refs []declarationDTO) (string, error) {
	declarations := make([][]any, 0, len(refs))
	for _, ref := range refs {
		if ref[0] < 0 || ref[0] >= len(dto.Inputs) {
			return "", newCatalogError(ReasonCatalogInputManifest, "semantic ID declaration input is invalid")
		}
		declarations = append(declarations, []any{
			dto.Inputs[ref[0]].Path,
			ref[1],
			ref[2],
			ref[3],
			ref[4],
		})
	}
	payload := semanticIdentityPayload{
		Authority:    dto.AuthorityKind,
		TypeScript:   dto.TypeScript,
		Catalog:      dto.Hash,
		OwnerPath:    ownerPath,
		Name:         name,
		Meaning:      meaning,
		Declarations: declarations,
	}
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(payload); err != nil {
		return "", newCatalogError(ReasonCatalogInputManifest, "marshal semantic identity: %v", err)
	}
	raw := bytes.TrimSuffix(buffer.Bytes(), []byte{'\n'})
	return "tsstdlib:" + sha256Hex(raw), nil
}

func validateDeclarationLanes(inputCount int, lanes ...[]declarationDTO) error {
	total := 0
	for _, lane := range lanes {
		total += len(lane)
		for index, ref := range lane {
			if ref[0] < 0 || ref[0] >= inputCount || ref[1] <= 0 || ref[2] <= 0 || ref[3] <= 0 || ref[4] <= 0 {
				return newCatalogError(ReasonCatalogInputManifest, "declaration reference is invalid")
			}
			if index > 0 && !declarationLess(lane[index-1], ref) {
				return newCatalogError(ReasonCatalogInputManifest, "declaration references are not strictly sorted")
			}
		}
	}
	if total == 0 {
		return newCatalogError(ReasonCatalogInputManifest, "declaration has no semantic lane")
	}
	return nil
}

func declarationLess(left declarationDTO, right declarationDTO) bool {
	for index := range left {
		if left[index] == right[index] {
			continue
		}
		return left[index] < right[index]
	}
	return false
}

func validateIndexes(values []int, upper int, field string) error {
	for index, value := range values {
		if value < 0 || value >= upper || (index > 0 && values[index-1] >= value) {
			return newCatalogError(ReasonCatalogInputManifest, "%s are invalid", field)
		}
	}
	return nil
}

func validSHA256(value string) bool {
	if len(value) != sha256.Size*2 {
		return false
	}
	_, err := hex.DecodeString(value)
	return err == nil
}

func newCatalogError(reason Reason, format string, args ...any) error {
	return &catalogError{reason: reason, err: fmt.Errorf(format, args...)}
}

func authorityDigest(dto catalogDTO, artifactHash string) string {
	return sha256Hex([]byte(strings.Join([]string{
		dto.AuthorityKind,
		dto.IdentityVersion,
		dto.Schema,
		dto.TypeScript,
		dto.Integrity,
		dto.Generator,
		dto.Hash,
		artifactHash,
	}, "\x00")))
}

func sha256Hex(raw []byte) string {
	digest := sha256.Sum256(raw)
	return hex.EncodeToString(digest[:])
}
