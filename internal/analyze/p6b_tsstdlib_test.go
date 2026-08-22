package analyze

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
	"github.com/tamnguyendinh/anvien/internal/resolution"
	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/tsstdlib"
)

func TestP6BRunInjectsEmbeddedTypeScriptStandardLibrary(t *testing.T) {
	repoPath, err := filepath.Abs(filepath.Join("testdata", "p6b-tsstdlib-runtime"))
	if err != nil {
		t.Fatalf("resolve fixture path: %v", err)
	}
	storagePath := filepath.Join(repoPath, ".anvien")
	if err := os.RemoveAll(storagePath); err != nil {
		t.Fatalf("clear fixture storage before analyze: %v", err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(storagePath); err != nil {
			t.Errorf("clear fixture storage after analyze: %v", err)
		}
	})

	result, err := Run(context.Background(), repoPath, Options{Force: true})
	if err != nil {
		t.Fatalf("run analyze with embedded TypeScript authority: %v", err)
	}
	metrics := result.Metrics.Resolution
	if metrics.ResolvedExternalDeclarations < 3 ||
		metrics.ExternalCapabilityUnavailable != 0 ||
		metrics.ExternalProfileExcluded != 0 ||
		metrics.ExternalMeaningMismatches != 0 {
		t.Fatalf("unexpected TypeScript authority metrics: %#v", metrics)
	}
	requireAnalyzeP6BSiteCounterEquality(t, result)
	want := map[string]bool{
		"call|Math.max":          false,
		"call|Promise":           false,
		"type-reference|Promise": false,
	}
	for _, record := range result.TypeScriptAuthorityResults {
		key := record.SiteKind + "|" + record.RequestedName
		if _, required := want[key]; required {
			want[key] = true
		}
	}
	for key, found := range want {
		if !found {
			t.Fatalf("built analyze result is missing %s: %#v", key, result.TypeScriptAuthorityResults)
		}
	}
	wantSymbols := make(map[string]struct{})
	wantSites := make(map[string]struct{})
	for _, record := range result.TypeScriptAuthorityResults {
		if record.Status != tsstdlib.LookupResolved {
			continue
		}
		wantSymbols[record.ResolvedSymbolID] = struct{}{}
		wantSites[record.SourceSiteID] = struct{}{}
	}
	externalNodes := make(map[string]struct{})
	for _, node := range result.Graph.Nodes {
		if node.Label != scopeir.NodeExternalSymbol {
			continue
		}
		semanticID, _ := node.Properties["semanticSymbolId"].(string)
		if _, expected := wantSymbols[semanticID]; !expected ||
			node.ID != graph.GenerateID(string(scopeir.NodeExternalSymbol), semanticID) ||
			node.Properties["external"] != true ||
			node.Properties["editable"] != false ||
			node.Properties["repositoryOwned"] != false {
			t.Fatalf("built analyze emitted invalid external node: %#v", node)
		}
		if _, repositoryFile := node.Properties["filePath"]; repositoryFile {
			t.Fatalf("built external node claimed repository file ownership: %#v", node)
		}
		externalNodes[node.ID] = struct{}{}
	}
	if len(externalNodes) != len(wantSymbols) {
		t.Fatalf("built external node count = %d, want %d", len(externalNodes), len(wantSymbols))
	}
	proofSites := make(map[string]int)
	for _, relationship := range result.Graph.Relationships {
		if _, external := externalNodes[relationship.TargetID]; !external {
			continue
		}
		if relationship.Type == graph.RelDefines || relationship.Type == graph.RelImports ||
			relationship.ResolutionSource != resolution.TypeScriptStandardLibraryStage {
			t.Fatalf("built external edge disguised repository ownership: %#v", relationship)
		}
		for _, evidence := range relationship.Evidence {
			if evidence.Kind != "typescript-authority-result-v1" {
				continue
			}
			var proof resolution.TypeScriptAuthorityResult
			if err := json.Unmarshal([]byte(evidence.Note), &proof); err != nil {
				t.Fatalf("decode built external authority proof: %v", err)
			}
			proofSites[proof.SourceSiteID]++
		}
	}
	for sourceSiteID := range wantSites {
		if proofSites[sourceSiteID] != 1 {
			t.Fatalf("built external source site %q proof count = %d, want 1", sourceSiteID, proofSites[sourceSiteID])
		}
	}
	for _, node := range result.Graph.Nodes {
		diagnostics, ok := node.Properties[graphhealth.DiagnosticPropertyKey].([]graphhealth.Diagnostic)
		if !ok {
			continue
		}
		for _, diagnostic := range diagnostics {
			if strings.Contains(diagnostic.TargetText, "Promise") || diagnostic.TargetText == "Math.max" {
				t.Fatalf("standard-library site remained a repository gap: %#v", diagnostic)
			}
		}
	}
}

func TestP6BRunExposesUnavailableAndMismatchSiteRecords(t *testing.T) {
	tests := []struct {
		name       string
		fixture    string
		wantStatus map[tsstdlib.LookupStatus]int
	}{
		{
			name:    "no-lib-capability-unavailable",
			fixture: "p6b-tsstdlib-runtime-no-lib",
			wantStatus: map[tsstdlib.LookupStatus]int{
				tsstdlib.LookupCapabilityUnavailable: 3,
			},
		},
		{
			name:    "default-profile-excluded-and-meaning-mismatch",
			fixture: "p6b-tsstdlib-runtime-default",
			wantStatus: map[tsstdlib.LookupStatus]int{
				tsstdlib.LookupProfileExcluded: 1,
				tsstdlib.LookupMeaningMismatch: 1,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := runP6BAnalyzeFixture(t, test.fixture)
			requireAnalyzeP6BSiteCounterEquality(t, result)
			counts := map[tsstdlib.LookupStatus]int{}
			for _, record := range result.TypeScriptAuthorityResults {
				counts[record.Status]++
			}
			for status, minimum := range test.wantStatus {
				if counts[status] < minimum {
					t.Fatalf("built authority status %s count=%d, want at least %d; records=%#v", status, counts[status], minimum, result.TypeScriptAuthorityResults)
				}
			}
		})
	}
}

func TestP6BRunCarriesCatalogValidationFailuresPerSourceSite(t *testing.T) {
	wantSiteIDs := map[string]struct{}{
		"SourceSite:main.ts#call#Math.max#2#18#2#44":                  {},
		"SourceSite:main.ts#call#Promise#3#9#3#59":                    {},
		"SourceSite:main.ts#call#resolve#3#42#3#58":                   {},
		"SourceSite:main.ts#type-reference#Math.max#2#8#2#44":         {},
		"SourceSite:main.ts#type-reference#Promise#1#27#1#34":         {},
		"SourceSite:main.ts#type-reference#Promise#1#45#1#52":         {},
		"SourceSite:main.ts#type-reference#Promise<number>#1#25#1#42": {},
	}

	for _, test := range analyzeP6BCatalogValidationFailureCases(t) {
		t.Run(test.name, func(t *testing.T) {
			result := runP6BAnalyzeFixtureWithCatalog(t, "p6b-tsstdlib-runtime", test.raw)
			requireAnalyzeP6BSiteCounterEquality(t, result)
			if len(result.TypeScriptAuthorityResults) != len(wantSiteIDs) ||
				result.Metrics.Resolution.ExternalCapabilityUnavailable != len(wantSiteIDs) ||
				result.Metrics.Resolution.UnresolvedReferences != 0 ||
				result.Metrics.Resolution.UnresolvedReferenceDiagnostics != len(wantSiteIDs) {
				t.Fatalf("built catalog-failure carriage records=%#v metrics=%#v want unique sites=%d", result.TypeScriptAuthorityResults, result.Metrics.Resolution, len(wantSiteIDs))
			}
			seen := make(map[string]struct{}, len(result.TypeScriptAuthorityResults))
			for _, record := range result.TypeScriptAuthorityResults {
				if _, expected := wantSiteIDs[record.SourceSiteID]; !expected {
					t.Fatalf("built catalog failure invented source site %q: %#v", record.SourceSiteID, record)
				}
				seen[record.SourceSiteID] = struct{}{}
				if record.Status != tsstdlib.LookupCapabilityUnavailable || record.Reason != test.reason || record.CatalogProofState != test.proofState {
					t.Fatalf("built catalog failure record = %#v, want unavailable/%s/%s", record, test.reason, test.proofState)
				}
			}
			if len(seen) != len(wantSiteIDs) {
				t.Fatalf("built catalog failure dropped sites: got=%#v want=%#v", seen, wantSiteIDs)
			}

			replay := runP6BAnalyzeFixtureWithCatalog(t, "p6b-tsstdlib-runtime", test.raw)
			if !reflect.DeepEqual(result.TypeScriptAuthorityResults, replay.TypeScriptAuthorityResults) {
				t.Fatalf("built catalog failure replay drift:\nfirst=%#v\nreplay=%#v", result.TypeScriptAuthorityResults, replay.TypeScriptAuthorityResults)
			}
		})
	}
}

func analyzeP6BCatalogValidationFailureCases(t *testing.T) []struct {
	name       string
	raw        []byte
	reason     tsstdlib.Reason
	proofState tsstdlib.CatalogProofState
} {
	t.Helper()
	tests := []struct {
		name       string
		fixture    string
		reason     tsstdlib.Reason
		proofState tsstdlib.CatalogProofState
	}{
		{name: "empty-missing", reason: tsstdlib.ReasonCatalogMissing, proofState: tsstdlib.CatalogProofMissing},
		{name: "schema", fixture: "schema.json", reason: tsstdlib.ReasonCatalogSchema, proofState: tsstdlib.CatalogProofRejected},
		{name: "version", fixture: "version.json", reason: tsstdlib.ReasonCatalogVersion, proofState: tsstdlib.CatalogProofRejected},
		{name: "input-manifest", fixture: "input-manifest.json", reason: tsstdlib.ReasonCatalogInputManifest, proofState: tsstdlib.CatalogProofRejected},
		{name: "logical-hash", fixture: "logical-hash.json", reason: tsstdlib.ReasonCatalogHash, proofState: tsstdlib.CatalogProofRejected},
		{name: "trailing-artifact-integrity", fixture: "trailing-artifact-integrity.json", reason: tsstdlib.ReasonCatalogInputManifest, proofState: tsstdlib.CatalogProofRejected},
	}
	out := make([]struct {
		name       string
		raw        []byte
		reason     tsstdlib.Reason
		proofState tsstdlib.CatalogProofState
	}, 0, len(tests))
	for _, test := range tests {
		var raw []byte
		if test.fixture != "" {
			var err error
			raw, err = os.ReadFile(filepath.Join("..", "tsstdlib", "testdata", "catalog-failures", test.fixture))
			if err != nil {
				t.Fatalf("read catalog failure fixture %s: %v", test.fixture, err)
			}
		}
		out = append(out, struct {
			name       string
			raw        []byte
			reason     tsstdlib.Reason
			proofState tsstdlib.CatalogProofState
		}{name: test.name, raw: raw, reason: test.reason, proofState: test.proofState})
	}
	return out
}

func TestP6BTypeScriptDeclarationInventoryIsLanguageIsolated(t *testing.T) {
	inventory, ok := typeScriptDeclarationInventory([]scanner.File{
		{Path: "tsconfig.json"},
		{Path: "src/main.ts", Language: scanner.TypeScript},
		{Path: "src/main.go", Language: scanner.Go},
	})
	if !ok || len(inventory) != 3 || inventory[0] != "tsconfig.json" || inventory[1] != "src/main.ts" || inventory[2] != "src/main.go" {
		t.Fatalf("unexpected TypeScript inventory: ok=%v inventory=%#v", ok, inventory)
	}
	if inventory, ok := typeScriptDeclarationInventory([]scanner.File{{Path: "main.go", Language: scanner.Go}}); ok || len(inventory) != 1 {
		t.Fatalf("non-TypeScript inventory activated authority: ok=%v inventory=%#v", ok, inventory)
	}
}

func runP6BAnalyzeFixture(t *testing.T, fixture string) Result {
	t.Helper()
	repoPath, err := filepath.Abs(filepath.Join("testdata", fixture))
	if err != nil {
		t.Fatalf("resolve fixture path: %v", err)
	}
	storagePath := filepath.Join(repoPath, ".anvien")
	if err := os.RemoveAll(storagePath); err != nil {
		t.Fatalf("clear fixture storage before analyze: %v", err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(storagePath); err != nil {
			t.Errorf("clear fixture storage after analyze: %v", err)
		}
	})
	result, err := Run(context.Background(), repoPath, Options{Force: true})
	if err != nil {
		t.Fatalf("run analyze fixture %s: %v", fixture, err)
	}
	return result
}

func runP6BAnalyzeFixtureWithCatalog(t *testing.T, fixture string, raw []byte) Result {
	t.Helper()
	repoPath, err := filepath.Abs(filepath.Join("testdata", fixture))
	if err != nil {
		t.Fatalf("resolve fixture path: %v", err)
	}
	storagePath := filepath.Join(repoPath, ".anvien")
	if err := os.RemoveAll(storagePath); err != nil {
		t.Fatalf("clear fixture storage before analyze: %v", err)
	}
	t.Cleanup(func() {
		if err := os.RemoveAll(storagePath); err != nil {
			t.Errorf("clear fixture storage after analyze: %v", err)
		}
	})
	authority := tsstdlib.NewAuthorityFromCatalog(repoPath, []string{"main.ts", "tsconfig.json"}, raw)
	result, err := Run(context.Background(), repoPath, Options{
		Force: true,
		Resolution: resolution.Options{
			TypeScriptStandardLibrary: authority,
		},
	})
	if err != nil {
		t.Fatalf("run analyze fixture %s with rejected catalog: %v", fixture, err)
	}
	return result
}

func requireAnalyzeP6BSiteCounterEquality(t *testing.T, result Result) {
	t.Helper()
	metrics := result.Metrics.Resolution
	counts := map[tsstdlib.LookupStatus]int{}
	seen := map[string]struct{}{}
	for _, record := range result.TypeScriptAuthorityResults {
		counts[record.Status]++
		if _, duplicate := seen[record.SourceSiteID]; duplicate {
			t.Fatalf("duplicate built source-site authority record: %#v", record)
		}
		seen[record.SourceSiteID] = struct{}{}
		if record.SourceSiteID == "" || record.FilePath == "" || record.RequestedName == "" ||
			record.ProfileHash == "" || record.ConfigHash == "" {
			t.Fatalf("built authority record is incomplete: %#v", record)
		}
		switch record.CatalogProofState {
		case tsstdlib.CatalogProofReady:
			if record.AuthorityHash == "" || record.CatalogHash == "" || record.CatalogArtifactHash == "" {
				t.Fatalf("built ready catalog proof is incomplete: %#v", record)
			}
		case tsstdlib.CatalogProofMissing:
			if record.Reason != tsstdlib.ReasonCatalogMissing || record.AuthorityHash != "" || record.CatalogHash != "" || record.CatalogArtifactHash != "" {
				t.Fatalf("built missing catalog proof fabricated validated identity: %#v", record)
			}
		case tsstdlib.CatalogProofRejected:
			if record.AuthorityHash != "" || record.CatalogHash != "" || record.CatalogArtifactHash == "" {
				t.Fatalf("built rejected catalog proof did not preserve explicit absence plus attempted artifact: %#v", record)
			}
		default:
			t.Fatalf("built authority record has unknown catalog proof state: %#v", record)
		}
	}
	if counts[tsstdlib.LookupResolved] != metrics.ResolvedExternalDeclarations ||
		counts[tsstdlib.LookupCapabilityUnavailable] != metrics.ExternalCapabilityUnavailable ||
		counts[tsstdlib.LookupProfileExcluded] != metrics.ExternalProfileExcluded ||
		counts[tsstdlib.LookupMeaningMismatch] != metrics.ExternalMeaningMismatches {
		t.Fatalf("built authority site/counter drift: counts=%#v metrics=%#v", counts, metrics)
	}
}
