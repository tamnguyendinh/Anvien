package resolution

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/tsstdlib"
)

const (
	typeScriptExternalReferenceEvidenceKind = "typescript-standard-library-reference"
	typeScriptAuthorityEvidenceKind         = "typescript-authority-result-v1"
	typeScriptExternalProofKind             = "typescript-standard-library-authority"
)

// typeScriptExternalSite carries only the transient graph endpoint data that
// is deliberately absent from the immutable P6-B authority result contract.
type typeScriptExternalSite struct {
	SourceSiteID  string
	SourceGraphID string
	FromScope     string
	ReferenceKind ReferenceKind
	TargetRole    string
	TargetText    string
}

type typeScriptExternalGroup struct {
	symbolID            string
	ownerID             string
	authorityKind       string
	typeScriptVersion   string
	catalogProofState   tsstdlib.CatalogProofState
	authorityHash       string
	catalogHash         string
	catalogArtifactHash string
	requestedNames      map[string]struct{}
	requestedTargetText map[string]struct{}
	meanings            map[string]struct{}
	profileHashes       map[string]struct{}
	configHashes        map[string]struct{}
	sourceSiteIDs       map[string]struct{}
	declarations        map[string]tsstdlib.Declaration
}

func emitTypeScriptExternalSymbols(e *emitter, records []TypeScriptAuthorityResult) error {
	sites, err := indexTypeScriptExternalSites(e.typeScriptExternalSites)
	if err != nil {
		return err
	}
	groups := make(map[string]*typeScriptExternalGroup)
	recordBySite := make(map[string]TypeScriptAuthorityResult, len(records))
	for _, record := range records {
		site, ok := sites[record.SourceSiteID]
		if !ok {
			return fmt.Errorf("missing TypeScript external endpoint for source site %q", record.SourceSiteID)
		}
		recordBySite[record.SourceSiteID] = record
		if record.Status != tsstdlib.LookupResolved {
			continue
		}
		group := groups[record.ResolvedSymbolID]
		if group == nil {
			group = newTypeScriptExternalGroup()
			groups[record.ResolvedSymbolID] = group
		}
		if err := group.add(record, site); err != nil {
			return err
		}
	}
	if len(recordBySite) != len(sites) {
		return fmt.Errorf("TypeScript external endpoint/result drift: endpoints=%d results=%d", len(sites), len(recordBySite))
	}

	symbolIDs := make([]string, 0, len(groups))
	for symbolID := range groups {
		symbolIDs = append(symbolIDs, symbolID)
	}
	sort.Strings(symbolIDs)
	for _, symbolID := range symbolIDs {
		node := groups[symbolID].node()
		if existing, ok := e.graph.GetNode(node.ID); ok {
			if !reflect.DeepEqual(existing, node) {
				return fmt.Errorf("external symbol identity collision for graph node %q", node.ID)
			}
			continue
		}
		e.emitNode(node)
	}

	for _, record := range records {
		if record.Status != tsstdlib.LookupResolved {
			continue
		}
		if err := emitTypeScriptExternalReference(e, sites[record.SourceSiteID], record); err != nil {
			return err
		}
	}
	return nil
}

func indexTypeScriptExternalSites(values []typeScriptExternalSite) (map[string]typeScriptExternalSite, error) {
	out := make(map[string]typeScriptExternalSite, len(values))
	for _, site := range values {
		if site.SourceSiteID == "" || site.SourceGraphID == "" || site.FromScope == "" ||
			site.ReferenceKind == "" || site.TargetRole == "" || strings.TrimSpace(site.TargetText) == "" {
			return nil, fmt.Errorf("incomplete TypeScript external endpoint for source site %q", site.SourceSiteID)
		}
		if previous, duplicate := out[site.SourceSiteID]; duplicate {
			if previous != site {
				return nil, fmt.Errorf("conflicting TypeScript external endpoint for source site %q", site.SourceSiteID)
			}
			continue
		}
		out[site.SourceSiteID] = site
	}
	return out, nil
}

func newTypeScriptExternalGroup() *typeScriptExternalGroup {
	return &typeScriptExternalGroup{
		requestedNames:      make(map[string]struct{}),
		requestedTargetText: make(map[string]struct{}),
		meanings:            make(map[string]struct{}),
		profileHashes:       make(map[string]struct{}),
		configHashes:        make(map[string]struct{}),
		sourceSiteIDs:       make(map[string]struct{}),
		declarations:        make(map[string]tsstdlib.Declaration),
	}
}

func (group *typeScriptExternalGroup) add(record TypeScriptAuthorityResult, site typeScriptExternalSite) error {
	if group.symbolID == "" {
		group.symbolID = record.ResolvedSymbolID
		group.ownerID = record.ResolvedOwnerID
		group.authorityKind = record.AuthorityKind
		group.typeScriptVersion = record.TypeScriptVersion
		group.catalogProofState = record.CatalogProofState
		group.authorityHash = record.AuthorityHash
		group.catalogHash = record.CatalogHash
		group.catalogArtifactHash = record.CatalogArtifactHash
	} else if group.symbolID != record.ResolvedSymbolID ||
		group.ownerID != record.ResolvedOwnerID ||
		group.authorityKind != record.AuthorityKind ||
		group.typeScriptVersion != record.TypeScriptVersion ||
		group.catalogProofState != record.CatalogProofState ||
		group.authorityHash != record.AuthorityHash ||
		group.catalogHash != record.CatalogHash ||
		group.catalogArtifactHash != record.CatalogArtifactHash {
		return fmt.Errorf("conflicting TypeScript external provenance for semantic symbol %q", record.ResolvedSymbolID)
	}
	group.requestedNames[record.RequestedName] = struct{}{}
	group.requestedTargetText[site.TargetText] = struct{}{}
	group.meanings[string(record.RequestedMeaning)] = struct{}{}
	group.profileHashes[record.ProfileHash] = struct{}{}
	group.configHashes[record.ConfigHash] = struct{}{}
	group.sourceSiteIDs[record.SourceSiteID] = struct{}{}
	for _, declaration := range record.DeclarationRanges {
		group.declarations[typeScriptDeclarationKey(declaration)] = declaration
	}
	return nil
}

func (group *typeScriptExternalGroup) node() graph.Node {
	names := sortedTypeScriptExternalStrings(group.requestedNames)
	targetTexts := sortedTypeScriptExternalStrings(group.requestedTargetText)
	meanings := sortedTypeScriptExternalStrings(group.meanings)
	profileHashes := sortedTypeScriptExternalStrings(group.profileHashes)
	configHashes := sortedTypeScriptExternalStrings(group.configHashes)
	sourceSiteIDs := sortedTypeScriptExternalStrings(group.sourceSiteIDs)
	declarations := sortedTypeScriptDeclarations(group.declarations)
	libraries := typeScriptDeclarationLibraries(declarations)
	properties := graph.NodeProperties{
		"name":                 names[0],
		"qualifiedName":        names[0],
		"requestedNames":       names,
		"requestedTargetTexts": targetTexts,
		"meaning":              meanings[0],
		"meanings":             meanings,
		"semanticSymbolId":     group.symbolID,
		"authorityKind":        group.authorityKind,
		"typeScriptVersion":    group.typeScriptVersion,
		"catalogProofState":    string(group.catalogProofState),
		"authorityHash":        group.authorityHash,
		"catalogHash":          group.catalogHash,
		"catalogArtifactHash":  group.catalogArtifactHash,
		"profileHashes":        profileHashes,
		"configHashes":         configHashes,
		"declarationLibraries": libraries,
		"declarationRanges":    declarations,
		"sourceSiteIds":        sourceSiteIDs,
		"sourceSiteCount":      len(sourceSiteIDs),
		"origin":               TypeScriptStandardLibraryStage,
		"external":             true,
		"editable":             false,
		"repositoryOwned":      false,
	}
	if group.ownerID != "" {
		properties["semanticOwnerId"] = group.ownerID
	}
	return graph.Node{
		ID:         externalSymbolGraphID(group.symbolID),
		Label:      scopeir.NodeExternalSymbol,
		Properties: properties,
	}
}

func emitTypeScriptExternalReference(e *emitter, site typeScriptExternalSite, record TypeScriptAuthorityResult) error {
	source, ok := e.graph.GetNode(site.SourceGraphID)
	if !ok {
		return fmt.Errorf("missing source graph node %q for TypeScript external site %q", site.SourceGraphID, record.SourceSiteID)
	}
	if source.Label == scopeir.NodeExternalSymbol {
		return fmt.Errorf("TypeScript external site %q cannot use an external source endpoint", record.SourceSiteID)
	}
	encoded, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("marshal TypeScript authority proof for source site %q: %w", record.SourceSiteID, err)
	}
	relationshipType := relationshipTypeForReference(site.ReferenceKind)
	relationship := graph.Relationship{
		ID:               "rel:" + string(relationshipType) + ":" + site.SourceGraphID + "->" + externalSymbolGraphID(record.ResolvedSymbolID) + ":" + record.SourceSiteID,
		SourceID:         site.SourceGraphID,
		TargetID:         externalSymbolGraphID(record.ResolvedSymbolID),
		Type:             relationshipType,
		Confidence:       1,
		Reason:           "resolved by embedded TypeScript standard-library authority",
		ResolutionSource: TypeScriptStandardLibraryStage,
		FileHash:         record.FileHash,
		SourceSiteID:     record.SourceSiteID,
		SourceSiteIDs:    []string{record.SourceSiteID},
		SourceSiteCount:  1,
		SourceSiteStatus: sourceSiteStatusResolved,
		ProofKind:        typeScriptExternalProofKind,
		TargetRole:       site.TargetRole,
		TargetText:       site.TargetText,
		FilePath:         record.FilePath,
		StartLine:        record.Range.StartLine,
		StartCol:         record.Range.StartCol,
		EndLine:          record.Range.EndLine,
		EndCol:           record.Range.EndCol,
		Evidence: []graph.Evidence{
			{Kind: typeScriptExternalReferenceEvidenceKind, Weight: 1, Note: record.RequestedName},
			{Kind: typeScriptAuthorityEvidenceKind, Weight: 1, Note: string(encoded)},
		},
	}
	if site.ReferenceKind == ReferenceRead {
		step := 1
		relationship.Step = &step
	}
	if site.ReferenceKind == ReferenceWrite {
		step := 2
		relationship.Step = &step
	}
	e.emitRelationship(relationship)
	e.referenceIndex.add(Reference{
		FromScope:        site.FromScope,
		ToDefID:          record.ResolvedSymbolID,
		FilePath:         record.FilePath,
		FileHash:         record.FileHash,
		Range:            record.Range,
		Kind:             site.ReferenceKind,
		Confidence:       1,
		SourceSiteID:     record.SourceSiteID,
		SourceSiteStatus: sourceSiteStatusResolved,
		ProofKind:        typeScriptExternalProofKind,
		TargetRole:       site.TargetRole,
		TargetText:       site.TargetText,
		Evidence:         append([]graph.Evidence(nil), relationship.Evidence...),
	})
	e.metrics.ResolvedReferences++
	switch site.ReferenceKind {
	case ReferenceCall:
		e.metrics.ResolvedCalls++
	case ReferenceRead, ReferenceWrite:
		e.metrics.ResolvedAccesses++
	case ReferenceTypeReference:
		e.metrics.ResolvedTypeReferences++
	}
	return nil
}

func externalSymbolGraphID(semanticSymbolID string) string {
	return graph.GenerateID(string(scopeir.NodeExternalSymbol), semanticSymbolID)
}

func sortedTypeScriptExternalStrings(values map[string]struct{}) []string {
	out := make([]string, 0, len(values))
	for value := range values {
		if value != "" {
			out = append(out, value)
		}
	}
	sort.Strings(out)
	return out
}

func typeScriptDeclarationKey(value tsstdlib.Declaration) string {
	return fmt.Sprintf("%s\x00%d\x00%d\x00%d\x00%d", value.Library, value.StartLine, value.StartCol, value.EndLine, value.EndCol)
}

func sortedTypeScriptDeclarations(values map[string]tsstdlib.Declaration) []tsstdlib.Declaration {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	out := make([]tsstdlib.Declaration, 0, len(keys))
	for _, key := range keys {
		out = append(out, values[key])
	}
	return out
}

func typeScriptDeclarationLibraries(values []tsstdlib.Declaration) []string {
	seen := make(map[string]struct{}, len(values))
	for _, value := range values {
		seen[value.Library] = struct{}{}
	}
	return sortedTypeScriptExternalStrings(seen)
}
