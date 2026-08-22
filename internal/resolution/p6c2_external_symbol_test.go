package resolution

import (
	"reflect"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/tsstdlib"
)

func TestP6C2UnavailableAuthorityDoesNotMaterializeExternalSymbol(t *testing.T) {
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
		t.Fatalf("resolve noLib external capability: %v", err)
	}
	requireP6BSiteRecord(t, result, tsstdlib.LookupCapabilityUnavailable, tsstdlib.ReasonDisabledByNoLib)
	for _, node := range result.Graph.Nodes {
		if node.Label == scopeir.NodeExternalSymbol {
			t.Fatalf("capability-unavailable site materialized a fake external target: %#v", node)
		}
	}
}

func TestP6C2ExternalMaterializationDeduplicatesSitesAndReplaysDeterministically(t *testing.T) {
	ir := p6bTypeScriptIR()
	call := p6bCall("max", scopeir.CallMember, "Math", 2)
	ir.Calls = []scopeir.CallSiteFact{call, call}

	first, err := Resolve([]scopeir.ScopeIR{ir}, Options{TypeScriptStandardLibrary: p6bAuthority("es2022")})
	if err != nil {
		t.Fatalf("resolve duplicated external site: %v", err)
	}
	second, err := Resolve([]scopeir.ScopeIR{ir}, Options{TypeScriptStandardLibrary: p6bAuthority("es2022")})
	if err != nil {
		t.Fatalf("replay duplicated external site: %v", err)
	}
	if !reflect.DeepEqual(first.Graph, second.Graph) || !reflect.DeepEqual(first.ReferenceIndex, second.ReferenceIndex) {
		t.Fatalf("external graph replay drift:\nfirst=%#v\nsecond=%#v", first.Graph, second.Graph)
	}
	if len(first.TypeScriptAuthorityResults) != 1 || first.Metrics.ResolvedExternalDeclarations != 1 {
		t.Fatalf("duplicate external site did not canonicalize: records=%#v metrics=%#v", first.TypeScriptAuthorityResults, first.Metrics)
	}
	externalNodes, externalEdges := 0, 0
	for _, node := range first.Graph.Nodes {
		if node.Label == scopeir.NodeExternalSymbol {
			externalNodes++
		}
	}
	for _, relationship := range first.Graph.Relationships {
		if relationship.TargetID == externalSymbolGraphID(first.TypeScriptAuthorityResults[0].ResolvedSymbolID) {
			externalEdges++
		}
	}
	if externalNodes != 1 || externalEdges != 1 {
		t.Fatalf("referenced-only dedupe = nodes %d edges %d, want 1/1", externalNodes, externalEdges)
	}
}

func TestP6C2ExternalMaterializationPreservesCoalescedSiteProofs(t *testing.T) {
	ir := p6bTypeScriptIR()
	ir.Calls = []scopeir.CallSiteFact{
		p6bCall("max", scopeir.CallMember, "Math", 2),
		p6bCall("max", scopeir.CallMember, "Math", 3),
	}
	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{TypeScriptStandardLibrary: p6bAuthority("es2022")})
	if err != nil {
		t.Fatalf("resolve coalesced external sites: %v", err)
	}
	if len(result.TypeScriptAuthorityResults) != 2 || result.Metrics.ResolvedExternalDeclarations != 2 {
		t.Fatalf("coalesced external sites lost authority records: records=%#v metrics=%#v", result.TypeScriptAuthorityResults, result.Metrics)
	}
	targetID := externalSymbolGraphID(result.TypeScriptAuthorityResults[0].ResolvedSymbolID)
	nodeSiteCount := 0
	for _, node := range result.Graph.Nodes {
		if node.ID == targetID {
			nodeSiteCount, _ = node.Properties["sourceSiteCount"].(int)
		}
	}
	externalEdges := 0
	for _, relationship := range result.Graph.Relationships {
		if relationship.TargetID != targetID {
			continue
		}
		externalEdges++
		proofCount := 0
		for _, evidence := range relationship.Evidence {
			if evidence.Kind == typeScriptAuthorityEvidenceKind {
				proofCount++
			}
		}
		if relationship.SourceSiteCount != 2 || len(relationship.SourceSiteIDs) != 2 || proofCount != 2 {
			t.Fatalf("coalesced external edge lost site-level proof: %#v", relationship)
		}
	}
	if nodeSiteCount != 2 || externalEdges != 1 {
		t.Fatalf("coalesced external materialization = node sites %d edges %d, want 2/1", nodeSiteCount, externalEdges)
	}
}
