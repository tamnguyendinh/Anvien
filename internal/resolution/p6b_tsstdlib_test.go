package resolution

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/tsstdlib"
)

func TestP6BResolverUsesCatalogForEligibleTypeScriptMisses(t *testing.T) {
	ir := p6bTypeScriptIR()
	ir.Calls = []scopeir.CallSiteFact{
		p6bCall("max", scopeir.CallMember, "Math", 2),
		p6bCall("Promise", scopeir.CallConstructor, "", 3),
		p6bCall("parseInt", scopeir.CallFree, "", 4),
	}
	ir.Accesses = []scopeir.AccessFact{{
		FilePath:         ir.FilePath,
		FileHash:         ir.FileHash,
		Name:             "PI",
		Kind:             scopeir.AccessRead,
		Range:            scopeir.Range{StartLine: 5, StartCol: 2, EndLine: 5, EndCol: 9},
		InScope:          p6bFunctionScope,
		ExplicitReceiver: "Math",
	}}
	ir.TypeAnnotations = []scopeir.TypeAnnotationFact{
		{
			FilePath: ir.FilePath,
			FileHash: ir.FileHash,
			Name:     "value",
			Range:    scopeir.Range{StartLine: 6, StartCol: 2, EndLine: 6, EndCol: 18},
			InScope:  p6bFunctionScope,
			Type: scopeir.TypeRef{
				RawName:         "Promise<number>",
				DeclaredAtScope: p6bFunctionScope,
				Source:          scopeir.TypeSourceAnnotation,
			},
		},
		{
			FilePath: ir.FilePath,
			FileHash: ir.FileHash,
			Name:     "maximum",
			Range:    scopeir.Range{StartLine: 7, StartCol: 2, EndLine: 7, EndCol: 18},
			InScope:  p6bFunctionScope,
			Type: scopeir.TypeRef{
				RawName:         "Math.max",
				DeclaredAtScope: p6bFunctionScope,
				Source:          scopeir.TypeSourceMethodReturn,
			},
		},
	}

	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{
		TypeScriptStandardLibrary: p6bAuthority("es2022"),
	})
	if err != nil {
		t.Fatalf("resolve TypeScript standard-library sites: %v", err)
	}
	if result.Metrics.ResolvedExternalDeclarations != 6 ||
		result.Metrics.ExternalCapabilityUnavailable != 0 ||
		result.Metrics.ExternalProfileExcluded != 0 ||
		result.Metrics.ExternalMeaningMismatches != 0 ||
		result.Metrics.UnresolvedReferences != 0 ||
		result.Metrics.UnresolvedReferenceDiagnostics != 0 {
		t.Fatalf("unexpected external resolution metrics: %#v", result.Metrics)
	}
	requireP6BSiteCounterEquality(t, result)
	wantSites := map[string]int{
		"call|Math.max":           1,
		"call|Promise":            1,
		"call|parseInt":           1,
		"access|Math.PI":          1,
		"type-reference|Promise":  1,
		"type-reference|Math.max": 1,
	}
	for _, record := range result.TypeScriptAuthorityResults {
		key := record.SiteKind + "|" + record.RequestedName
		wantSites[key]--
		if record.Status != tsstdlib.LookupResolved || record.ResolvedSymbolID == "" || len(record.DeclarationRanges) == 0 {
			t.Fatalf("resolved site record is incomplete: %#v", record)
		}
		if record.RequestedName == "Math.max" || record.RequestedName == "Math.PI" {
			if record.ResolvedOwnerID == "" {
				t.Fatalf("member site has no resolved external owner: %#v", record)
			}
		}
	}
	for key, remaining := range wantSites {
		if remaining != 0 {
			t.Fatalf("source-site inventory %s remaining=%d; records=%#v", key, remaining, result.TypeScriptAuthorityResults)
		}
	}
	for _, node := range result.Graph.Nodes {
		if node.Label == "ExternalSymbol" {
			t.Fatalf("P6-B must not materialize P6-C2 external nodes: %#v", node)
		}
	}
}

func TestP6BResolverKeepsRepositoryResolutionAuthoritative(t *testing.T) {
	ir := p6bTypeScriptIR()
	target := scopeir.DefinitionFact{
		ID:            "def:src/main.ts#8:0:Function:parseInt",
		FilePath:      ir.FilePath,
		FileHash:      ir.FileHash,
		Name:          "parseInt",
		QualifiedName: "parseInt",
		Label:         scopeir.NodeFunction,
		Range:         scopeir.Range{StartLine: 8, EndLine: 8},
	}
	ir.Definitions = append(ir.Definitions, target)
	ir.Scopes[1].Bindings = append(ir.Scopes[1].Bindings, scopeir.BindingFact{
		Name:   target.Name,
		DefID:  target.ID,
		Origin: scopeir.BindingLocal,
	})
	ir.Calls = []scopeir.CallSiteFact{p6bCall("parseInt", scopeir.CallFree, "", 2)}

	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{
		TypeScriptStandardLibrary: p6bAuthority("es2022"),
	})
	if err != nil {
		t.Fatalf("resolve repository precedence: %v", err)
	}
	if result.Metrics.ResolvedCalls != 1 || result.Metrics.ResolvedExternalDeclarations != 0 || result.Metrics.UnresolvedReferences != 0 {
		t.Fatalf("repository result did not remain authoritative: %#v", result.Metrics)
	}
	requireP6BSiteCounterEquality(t, result)
}

func TestP6BResolverDoesNotRescueExplicitImportFailure(t *testing.T) {
	ir := p6bTypeScriptIR()
	targetRaw := "./missing"
	ir.Imports = []scopeir.ImportFact{{
		ID:                "import:promise",
		FilePath:          ir.FilePath,
		FileHash:          ir.FileHash,
		Kind:              scopeir.ImportNamed,
		LocalName:         "Promise",
		ImportedName:      "Promise",
		RequestedMeanings: []scopeir.ExportMeaning{scopeir.ExportMeaningValue},
		TargetRaw:         &targetRaw,
	}}
	ir.Calls = []scopeir.CallSiteFact{p6bCall("Promise", scopeir.CallConstructor, "", 2)}

	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{
		TypeScriptStandardLibrary: p6bAuthority("es2022"),
	})
	if err != nil {
		t.Fatalf("resolve explicit import failure: %v", err)
	}
	if result.Metrics.ResolvedExternalDeclarations != 0 ||
		result.Metrics.UnresolvedReferences != 1 ||
		result.Metrics.UnresolvedReferenceDiagnostics != 1 {
		t.Fatalf("explicit import failure was not terminal: %#v", result.Metrics)
	}
	requireP6BSiteCounterEquality(t, result)
}

func TestP6BResolverMakesUnavailableAndExcludedProfilesExplicit(t *testing.T) {
	t.Run("no-lib", func(t *testing.T) {
		ir := p6bTypeScriptIR()
		ir.TypeAnnotations = []scopeir.TypeAnnotationFact{{
			FilePath: ir.FilePath,
			FileHash: ir.FileHash,
			Name:     "value",
			Range:    scopeir.Range{StartLine: 2, StartCol: 2, EndLine: 2, EndCol: 18},
			InScope:  p6bFunctionScope,
			Type:     scopeir.TypeRef{RawName: "Promise<number>", Source: scopeir.TypeSourceAnnotation},
		}}
		result, err := Resolve([]scopeir.ScopeIR{ir}, Options{TypeScriptStandardLibrary: p6bAuthority("no-lib")})
		if err != nil {
			t.Fatalf("resolve noLib profile: %v", err)
		}
		if result.Metrics.ExternalCapabilityUnavailable != 1 || result.Metrics.UnresolvedReferences != 0 {
			t.Fatalf("noLib miss became a repository gap: %#v", result.Metrics)
		}
		requireP6BSiteRecord(t, result, tsstdlib.LookupCapabilityUnavailable, tsstdlib.ReasonDisabledByNoLib)
	})

	t.Run("profile-excluded", func(t *testing.T) {
		ir := p6bTypeScriptIR()
		ir.Calls = []scopeir.CallSiteFact{p6bCall("Promise", scopeir.CallConstructor, "", 2)}
		result, err := Resolve([]scopeir.ScopeIR{ir}, Options{
			TypeScriptStandardLibrary: tsstdlib.NewAuthority(".", []string{"src/main.ts"}),
		})
		if err != nil {
			t.Fatalf("resolve default profile exclusion: %v", err)
		}
		if result.Metrics.ExternalProfileExcluded != 1 || result.Metrics.UnresolvedReferences != 0 {
			t.Fatalf("profile exclusion became a repository gap: %#v", result.Metrics)
		}
		requireP6BSiteRecord(t, result, tsstdlib.LookupProfileExcluded, tsstdlib.ReasonProfileExcludes)
	})

	t.Run("meaning-mismatch", func(t *testing.T) {
		ir := p6bTypeScriptIR()
		ir.Calls = []scopeir.CallSiteFact{p6bCall("ArrayLike", scopeir.CallFree, "", 2)}
		result, err := Resolve([]scopeir.ScopeIR{ir}, Options{TypeScriptStandardLibrary: p6bAuthority("es2022")})
		if err != nil {
			t.Fatalf("resolve meaning mismatch: %v", err)
		}
		if result.Metrics.ExternalMeaningMismatches != 1 || result.Metrics.UnresolvedReferences != 0 {
			t.Fatalf("meaning mismatch became a repository gap: %#v", result.Metrics)
		}
		requireP6BSiteRecord(t, result, tsstdlib.LookupMeaningMismatch, tsstdlib.ReasonMeaningMismatch)
	})
}

func TestP6BResolverCarriesCatalogValidationFailuresPerSourceSite(t *testing.T) {
	for _, test := range p6bCatalogValidationFailureCases(t) {
		t.Run(test.name, func(t *testing.T) {
			authority := tsstdlib.NewAuthorityFromCatalog(".", []string{"src/main.ts"}, test.raw)
			ir := p6bTypeScriptIR()
			valueCall := p6bCall("parseInt", scopeir.CallFree, "", 2)
			memberCall := p6bCall("max", scopeir.CallMember, "Math", 3)
			annotation := scopeir.TypeAnnotationFact{
				FilePath: ir.FilePath,
				FileHash: ir.FileHash,
				Name:     "value",
				Range:    scopeir.Range{StartLine: 4, StartCol: 2, EndLine: 4, EndCol: 18},
				InScope:  p6bFunctionScope,
				Type: scopeir.TypeRef{
					RawName:         "Promise<number>",
					DeclaredAtScope: p6bFunctionScope,
					Source:          scopeir.TypeSourceAnnotation,
				},
			}
			ir.Calls = []scopeir.CallSiteFact{valueCall, valueCall, memberCall, memberCall}
			ir.TypeAnnotations = []scopeir.TypeAnnotationFact{annotation, annotation}

			result, err := Resolve([]scopeir.ScopeIR{ir}, Options{TypeScriptStandardLibrary: authority})
			if err != nil {
				t.Fatalf("resolve %s catalog failure: %v", test.name, err)
			}
			requireP6BSiteCounterEquality(t, result)
			if len(result.TypeScriptAuthorityResults) != 3 || result.Metrics.ExternalCapabilityUnavailable != 3 ||
				result.Metrics.UnresolvedReferences != 0 || result.Metrics.UnresolvedReferenceDiagnostics != 0 {
				t.Fatalf("catalog failure did not canonicalize to one handled record per unique site: records=%#v metrics=%#v", result.TypeScriptAuthorityResults, result.Metrics)
			}
			want := map[string]int{"parseInt": 1, "Math.max": 1, "Promise": 1}
			for _, record := range result.TypeScriptAuthorityResults {
				want[record.RequestedName]--
				if record.Status != tsstdlib.LookupCapabilityUnavailable || record.Reason != test.reason || record.CatalogProofState != test.proofState {
					t.Fatalf("catalog failure record = %#v, want unavailable/%s/%s", record, test.reason, test.proofState)
				}
			}
			for name, remaining := range want {
				if remaining != 0 {
					t.Fatalf("catalog failure site %s remaining=%d: %#v", name, remaining, result.TypeScriptAuthorityResults)
				}
			}

			replay, err := Resolve([]scopeir.ScopeIR{ir}, Options{TypeScriptStandardLibrary: authority})
			if err != nil {
				t.Fatalf("replay %s catalog failure: %v", test.name, err)
			}
			if !reflect.DeepEqual(result.TypeScriptAuthorityResults, replay.TypeScriptAuthorityResults) {
				t.Fatalf("catalog failure replay drift:\nfirst=%#v\nreplay=%#v", result.TypeScriptAuthorityResults, replay.TypeScriptAuthorityResults)
			}

			conflict := cloneTypeScriptAuthorityResult(result.TypeScriptAuthorityResults[0])
			conflict.RequestedName += ".conflict"
			metrics := result.Metrics
			if _, err := finalizeTypeScriptAuthorityResults(
				[]TypeScriptAuthorityResult{result.TypeScriptAuthorityResults[0], conflict},
				&metrics,
			); err == nil {
				t.Fatal("conflicting catalog-failure payloads for one source site must fail closed")
			}
		})
	}
}

func TestP6BTypeScriptAuthorityValidationProofStatusReasonMatrix(t *testing.T) {
	readyCases := []struct {
		name   string
		status tsstdlib.LookupStatus
		reason tsstdlib.Reason
	}{
		{name: "resolved", status: tsstdlib.LookupResolved},
		{name: "profile-excluded", status: tsstdlib.LookupProfileExcluded, reason: tsstdlib.ReasonProfileExcludes},
		{name: "meaning-mismatch", status: tsstdlib.LookupMeaningMismatch, reason: tsstdlib.ReasonMeaningMismatch},
		{name: "no-lib", status: tsstdlib.LookupCapabilityUnavailable, reason: tsstdlib.ReasonDisabledByNoLib},
		{name: "config-invalid", status: tsstdlib.LookupCapabilityUnavailable, reason: tsstdlib.ReasonConfigInvalid},
		{name: "config-topology", status: tsstdlib.LookupCapabilityUnavailable, reason: tsstdlib.ReasonConfigTopology},
		{name: "config-unreadable", status: tsstdlib.LookupCapabilityUnavailable, reason: tsstdlib.ReasonConfigUnreadable},
	}
	for _, test := range readyCases {
		t.Run("accept-ready-"+test.name, func(t *testing.T) {
			record := p6bAuthorityValidationRecord(test.status, test.reason)
			if err := validateTypeScriptAuthorityResult(record); err != nil {
				t.Fatalf("valid ready proof was rejected: %v\nrecord=%#v", err, record)
			}
		})
	}

	invalidReadyCases := []struct {
		name   string
		status tsstdlib.LookupStatus
		reason tsstdlib.Reason
	}{
		{name: "resolved-with-profile-exclusion", status: tsstdlib.LookupResolved, reason: tsstdlib.ReasonProfileExcludes},
		{name: "profile-excluded-with-meaning-mismatch", status: tsstdlib.LookupProfileExcluded, reason: tsstdlib.ReasonMeaningMismatch},
		{name: "meaning-mismatch-with-profile-exclusion", status: tsstdlib.LookupMeaningMismatch, reason: tsstdlib.ReasonProfileExcludes},
		{name: "unavailable-with-empty-reason", status: tsstdlib.LookupCapabilityUnavailable},
		{name: "unavailable-with-unknown-reason", status: tsstdlib.LookupCapabilityUnavailable, reason: tsstdlib.Reason("unknown_ready_reason")},
		{name: "not-found", status: tsstdlib.LookupNotFound},
	}
	for _, test := range invalidReadyCases {
		t.Run("reject-ready-"+test.name, func(t *testing.T) {
			record := p6bAuthorityValidationRecord(test.status, test.reason)
			if err := validateTypeScriptAuthorityResult(record); err == nil {
				t.Fatalf("invalid ready status/reason combination was accepted: %#v", record)
			}
		})
	}

	for _, test := range p6bCatalogValidationFailureCases(t) {
		t.Run("accept-"+test.name+"-proof", func(t *testing.T) {
			record := p6bAuthorityValidationRecord(tsstdlib.LookupCapabilityUnavailable, test.reason)
			record.CatalogProofState = test.proofState
			record.AuthorityHash = ""
			record.CatalogHash = ""
			if test.proofState == tsstdlib.CatalogProofMissing {
				record.CatalogArtifactHash = ""
			} else {
				record.CatalogArtifactHash = "attempted-artifact-hash"
			}
			if err := validateTypeScriptAuthorityResult(record); err != nil {
				t.Fatalf("valid %s catalog proof was rejected: %v\nrecord=%#v", test.name, err, record)
			}
		})

		t.Run("reject-ready-"+test.name+"-reason", func(t *testing.T) {
			record := p6bAuthorityValidationRecord(tsstdlib.LookupCapabilityUnavailable, test.reason)
			if err := validateTypeScriptAuthorityResult(record); err == nil {
				t.Fatalf("ready proof accepted catalog failure reason %q with complete hashes: %#v", test.reason, record)
			}
		})
	}
}

func p6bAuthorityValidationRecord(status tsstdlib.LookupStatus, reason tsstdlib.Reason) TypeScriptAuthorityResult {
	record := TypeScriptAuthorityResult{
		SourceSiteID:        "SourceSite:src/main.ts#call#parseInt#2#2#2#12",
		Stage:               TypeScriptStandardLibraryStage,
		FilePath:            "src/main.ts",
		FileHash:            "hash-main",
		Range:               scopeir.Range{StartLine: 2, StartCol: 2, EndLine: 2, EndCol: 12},
		SiteKind:            "call",
		RequestedName:       "parseInt",
		RequestedMeaning:    tsstdlib.MeaningValue,
		Status:              status,
		Reason:              reason,
		AuthorityKind:       tsstdlib.AuthorityKind,
		CatalogProofState:   tsstdlib.CatalogProofReady,
		AuthorityHash:       "authority-hash",
		TypeScriptVersion:   tsstdlib.TypeScriptVersion,
		CatalogHash:         "logical-catalog-hash",
		CatalogArtifactHash: "artifact-hash",
		ProfileHash:         "profile-hash",
		ConfigHash:          "config-hash",
	}
	if status == tsstdlib.LookupResolved {
		record.ResolvedSymbolID = "tsstdlib:resolved-symbol"
		record.DeclarationRanges = []tsstdlib.Declaration{{
			Library:   "lib.es5.d.ts",
			StartLine: 42,
			StartCol:  1,
			EndLine:   42,
			EndCol:    20,
		}}
	}
	return record
}

func p6bCatalogValidationFailureCases(t *testing.T) []struct {
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

func TestP6BResolverDoesNotConsultCatalogForJavaScript(t *testing.T) {
	ir := p6bTypeScriptIR()
	ir.Language = scanner.JavaScript
	ir.Calls = []scopeir.CallSiteFact{p6bCall("max", scopeir.CallMember, "Math", 2)}
	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{
		TypeScriptStandardLibrary: p6bAuthority("es2022"),
	})
	if err != nil {
		t.Fatalf("resolve JavaScript isolation case: %v", err)
	}
	if result.Metrics.ResolvedExternalDeclarations != 0 ||
		result.Metrics.ExternalCapabilityUnavailable != 0 ||
		result.Metrics.ExternalProfileExcluded != 0 ||
		result.Metrics.ExternalMeaningMismatches != 0 ||
		result.Metrics.UnresolvedReferences != 1 {
		t.Fatalf("non-TypeScript site consulted the catalog: %#v", result.Metrics)
	}
	requireP6BSiteCounterEquality(t, result)
}

func TestP6BExternalMemberLookupRequiresExternalReceiverForCallAndAccess(t *testing.T) {
	tests := []struct {
		name           string
		prepare        func(scopeir.ScopeIR) scopeir.ScopeIR
		wantExternal   int
		wantUnresolved int
	}{
		{
			name: "local-value-claim",
			prepare: func(ir scopeir.ScopeIR) scopeir.ScopeIR {
				local := scopeir.DefinitionFact{
					ID: "def:src/main.ts#1:1:Variable:Math", FilePath: ir.FilePath, FileHash: ir.FileHash,
					Name: "Math", QualifiedName: "Math", Label: scopeir.NodeVariable,
					Range: scopeir.Range{StartLine: 1, StartCol: 1, EndLine: 1, EndCol: 5},
				}
				ir.Definitions = append(ir.Definitions, local)
				ir.Scopes[1].OwnedDefIDs = append(ir.Scopes[1].OwnedDefIDs, local.ID)
				ir.Scopes[1].Bindings = append(ir.Scopes[1].Bindings, scopeir.BindingFact{Name: "Math", DefID: local.ID, Origin: scopeir.BindingLocal})
				return ir
			},
			wantUnresolved: 2,
		},
		{
			name: "repository-type-claim-missing-member",
			prepare: func(ir scopeir.ScopeIR) scopeir.ScopeIR {
				owner := scopeir.DefinitionFact{
					ID: "def:src/main.ts#1:1:Class:Math", FilePath: ir.FilePath, FileHash: ir.FileHash,
					Name: "Math", QualifiedName: "Math", Label: scopeir.NodeClass,
					Range: scopeir.Range{StartLine: 1, StartCol: 1, EndLine: 1, EndCol: 5},
				}
				ir.Definitions = append(ir.Definitions, owner)
				ir.Scopes[0].OwnedDefIDs = append(ir.Scopes[0].OwnedDefIDs, owner.ID)
				return ir
			},
			wantUnresolved: 2,
		},
		{
			name: "explicit-import-receiver-failure",
			prepare: func(ir scopeir.ScopeIR) scopeir.ScopeIR {
				targetRaw := "./missing"
				ir.Imports = []scopeir.ImportFact{{
					ID: "import:math", FilePath: ir.FilePath, FileHash: ir.FileHash,
					Kind: scopeir.ImportNamespace, LocalName: "Math", ImportedName: "*", TargetRaw: &targetRaw,
				}}
				return ir
			},
			wantUnresolved: 2,
		},
		{
			name:         "genuine-external-receiver",
			prepare:      func(ir scopeir.ScopeIR) scopeir.ScopeIR { return ir },
			wantExternal: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ir := test.prepare(p6bTypeScriptIR())
			ir.Calls = []scopeir.CallSiteFact{p6bCall("max", scopeir.CallMember, "Math", 2)}
			ir.Accesses = []scopeir.AccessFact{p6bAccess("PI", "Math", 3)}
			result, err := Resolve([]scopeir.ScopeIR{ir}, Options{TypeScriptStandardLibrary: p6bAuthority("es2022")})
			if err != nil {
				t.Fatalf("resolve member receiver case: %v", err)
			}
			if result.Metrics.ResolvedExternalDeclarations != test.wantExternal ||
				result.Metrics.UnresolvedReferences != test.wantUnresolved ||
				result.Metrics.UnresolvedReferenceDiagnostics != test.wantUnresolved {
				t.Fatalf("receiver precedence metrics = %#v, want external=%d unresolved=%d", result.Metrics, test.wantExternal, test.wantUnresolved)
			}
			requireP6BSiteCounterEquality(t, result)
			if len(result.TypeScriptAuthorityResults) != test.wantExternal {
				t.Fatalf("authority results = %d, want %d: %#v", len(result.TypeScriptAuthorityResults), test.wantExternal, result.TypeScriptAuthorityResults)
			}
		})
	}
}

func TestP6BAuthorityResultsCanonicalizeDuplicateSourceFacts(t *testing.T) {
	ir := p6bTypeScriptIR()
	annotation := scopeir.TypeAnnotationFact{
		FilePath: ir.FilePath,
		FileHash: ir.FileHash,
		Name:     "value",
		Range:    scopeir.Range{StartLine: 6, StartCol: 2, EndLine: 6, EndCol: 18},
		InScope:  p6bFunctionScope,
		Type: scopeir.TypeRef{
			RawName:         "Promise<number>",
			DeclaredAtScope: p6bFunctionScope,
			Source:          scopeir.TypeSourceAnnotation,
		},
	}
	ir.TypeAnnotations = []scopeir.TypeAnnotationFact{annotation, annotation}

	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{TypeScriptStandardLibrary: p6bAuthority("es2022")})
	if err != nil {
		t.Fatalf("resolve duplicate TypeScript facts: %v", err)
	}
	requireP6BSiteCounterEquality(t, result)
	if len(result.TypeScriptAuthorityResults) != 1 || result.Metrics.ResolvedExternalDeclarations != 1 {
		t.Fatalf("duplicate facts were not canonicalized to one site/counter: records=%#v metrics=%#v", result.TypeScriptAuthorityResults, result.Metrics)
	}

	conflict := cloneTypeScriptAuthorityResult(result.TypeScriptAuthorityResults[0])
	conflict.RequestedName = "Map"
	metrics := result.Metrics
	metrics.ResolvedExternalDeclarations = 2
	if _, err := finalizeTypeScriptAuthorityResults(
		[]TypeScriptAuthorityResult{result.TypeScriptAuthorityResults[0], conflict},
		&metrics,
	); err == nil {
		t.Fatal("conflicting payloads for one stable source site must fail closed")
	}
}

func requireP6BSiteRecord(t *testing.T, result Result, status tsstdlib.LookupStatus, reason tsstdlib.Reason) {
	t.Helper()
	requireP6BSiteCounterEquality(t, result)
	if len(result.TypeScriptAuthorityResults) != 1 {
		t.Fatalf("authority result count = %d, want 1: %#v", len(result.TypeScriptAuthorityResults), result.TypeScriptAuthorityResults)
	}
	record := result.TypeScriptAuthorityResults[0]
	if record.Status != status || record.Reason != reason {
		t.Fatalf("authority result = %#v, want %s/%s", record, status, reason)
	}
}

func requireP6BSiteCounterEquality(t *testing.T, result Result) {
	t.Helper()
	counts := map[tsstdlib.LookupStatus]int{}
	seen := map[string]struct{}{}
	for _, record := range result.TypeScriptAuthorityResults {
		counts[record.Status]++
		if _, duplicate := seen[record.SourceSiteID]; duplicate {
			t.Fatalf("duplicate source-site authority result: %#v", record)
		}
		seen[record.SourceSiteID] = struct{}{}
		if record.SourceSiteID == "" || record.Stage != TypeScriptStandardLibraryStage || record.FilePath == "" ||
			record.RequestedName == "" || record.AuthorityKind != tsstdlib.AuthorityKind ||
			record.TypeScriptVersion != tsstdlib.TypeScriptVersion || record.ProfileHash == "" || record.ConfigHash == "" {
			t.Fatalf("authority record is not lossless: %#v", record)
		}
		switch record.CatalogProofState {
		case tsstdlib.CatalogProofReady:
			if record.AuthorityHash == "" || record.CatalogHash == "" || record.CatalogArtifactHash == "" {
				t.Fatalf("ready catalog proof is incomplete: %#v", record)
			}
		case tsstdlib.CatalogProofMissing:
			if record.Reason != tsstdlib.ReasonCatalogMissing || record.AuthorityHash != "" || record.CatalogHash != "" || record.CatalogArtifactHash != "" {
				t.Fatalf("missing catalog proof fabricated validated identity: %#v", record)
			}
		case tsstdlib.CatalogProofRejected:
			if record.AuthorityHash != "" || record.CatalogHash != "" || record.CatalogArtifactHash == "" {
				t.Fatalf("rejected catalog proof did not preserve explicit absence plus attempted artifact: %#v", record)
			}
		default:
			t.Fatalf("unknown catalog proof state: %#v", record)
		}
	}
	if counts[tsstdlib.LookupResolved] != result.Metrics.ResolvedExternalDeclarations ||
		counts[tsstdlib.LookupCapabilityUnavailable] != result.Metrics.ExternalCapabilityUnavailable ||
		counts[tsstdlib.LookupProfileExcluded] != result.Metrics.ExternalProfileExcluded ||
		counts[tsstdlib.LookupMeaningMismatch] != result.Metrics.ExternalMeaningMismatches {
		t.Fatalf("authority site/counter drift: counts=%#v metrics=%#v", counts, result.Metrics)
	}
}

const (
	p6bModuleScope   = "scope:src/main.ts#1:0-10:1:Module"
	p6bFunctionScope = "scope:src/main.ts#1:0-10:1:Function"
)

func p6bTypeScriptIR() scopeir.ScopeIR {
	parent := p6bModuleScope
	function := scopeir.DefinitionFact{
		ID:            "def:src/main.ts#1:0:Function:run",
		FilePath:      "src/main.ts",
		FileHash:      "hash-main",
		Name:          "run",
		QualifiedName: "run",
		Label:         scopeir.NodeFunction,
		Range:         scopeir.Range{StartLine: 1, EndLine: 10},
	}
	return scopeir.ScopeIR{
		FilePath:    "src/main.ts",
		FileHash:    "hash-main",
		Language:    scanner.TypeScript,
		ModuleScope: p6bModuleScope,
		Scopes: []scopeir.ScopeFact{
			{ID: p6bModuleScope, Kind: scopeir.ScopeModule, FilePath: "src/main.ts", FileHash: "hash-main", Range: scopeir.Range{StartLine: 1, EndLine: 10}},
			{
				ID:          p6bFunctionScope,
				Parent:      &parent,
				Kind:        scopeir.ScopeFunction,
				FilePath:    "src/main.ts",
				FileHash:    "hash-main",
				Range:       scopeir.Range{StartLine: 1, EndLine: 10},
				OwnedDefIDs: []string{function.ID},
				Bindings: []scopeir.BindingFact{{
					Name:   function.Name,
					DefID:  function.ID,
					Origin: scopeir.BindingLocal,
				}},
			},
		},
		Definitions: []scopeir.DefinitionFact{function},
	}
}

func p6bCall(name string, form scopeir.CallForm, receiver string, line int) scopeir.CallSiteFact {
	return scopeir.CallSiteFact{
		FilePath:         "src/main.ts",
		FileHash:         "hash-main",
		Name:             name,
		Range:            scopeir.Range{StartLine: line, StartCol: 2, EndLine: line, EndCol: 12},
		InScope:          p6bFunctionScope,
		CallForm:         form,
		ExplicitReceiver: receiver,
	}
}

func p6bAccess(name string, receiver string, line int) scopeir.AccessFact {
	return scopeir.AccessFact{
		FilePath:         "src/main.ts",
		FileHash:         "hash-main",
		Name:             name,
		Kind:             scopeir.AccessRead,
		Range:            scopeir.Range{StartLine: line, StartCol: 2, EndLine: line, EndCol: 12},
		InScope:          p6bFunctionScope,
		ExplicitReceiver: receiver,
	}
}

func p6bAuthority(profile string) *tsstdlib.Authority {
	return tsstdlib.NewAuthority(
		filepath.Join("..", "tsstdlib", "testdata", "profiles", profile),
		[]string{"src/main.ts", "tsconfig.json"},
	)
}
