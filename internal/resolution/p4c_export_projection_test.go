package resolution

import (
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestP4CProjectsExportFactsAndRuntimeCompatibility(t *testing.T) {
	const (
		filePath    = "src/exports.ts"
		moduleScope = "scope:src/exports.ts#1:0-20:0:Module"
	)
	runtimeDef := scopeir.DefinitionFact{
		ID:            "def:src/exports.ts#1:0:Variable:runtimeValue",
		FilePath:      filePath,
		Name:          "runtimeValue",
		QualifiedName: "runtimeValue",
		Label:         scopeir.NodeVariable,
		Range:         scopeir.Range{StartLine: 1, StartCol: 0, EndLine: 1, EndCol: 21},
		Visibility:    "private",
	}
	typeDef := scopeir.DefinitionFact{
		ID:            "def:src/exports.ts#2:0:Interface:TypeOnly",
		FilePath:      filePath,
		Name:          "TypeOnly",
		QualifiedName: "TypeOnly",
		Label:         scopeir.NodeInterface,
		Range:         scopeir.Range{StartLine: 2, StartCol: 0, EndLine: 2, EndCol: 20},
		Visibility:    "public",
	}
	targetRaw := "./remote"
	facts := []scopeir.ExportFact{
		{
			FilePath: filePath, FileHash: "hash-exports", Kind: scopeir.ExportDirect,
			ExportedName: "runtimeValue", LocalName: "runtimeValue", LocalDefID: runtimeDef.ID,
			Meanings:   []scopeir.ExportMeaning{scopeir.ExportMeaningValue},
			Range:      scopeir.Range{StartLine: 1, StartCol: 0, EndLine: 1, EndCol: 21},
			Provenance: scopeir.ExportProvenance{StatementRange: scopeir.Range{StartLine: 1, EndLine: 1}, SiteKind: "export_declaration"},
		},
		{
			FilePath: filePath, FileHash: "hash-exports", Kind: scopeir.ExportAlias,
			ExportedName: "renamed", LocalName: "runtimeValue", LocalDefID: runtimeDef.ID,
			Meanings:   []scopeir.ExportMeaning{scopeir.ExportMeaningValue},
			Range:      scopeir.Range{StartLine: 4, StartCol: 0, EndLine: 4, EndCol: 28},
			Provenance: scopeir.ExportProvenance{StatementRange: scopeir.Range{StartLine: 4, EndLine: 4}, SiteKind: "export_specifier"},
		},
		{
			FilePath: filePath, FileHash: "hash-exports", Kind: scopeir.ExportNamed,
			ExportedName: "TypeOnly", LocalName: "TypeOnly", LocalDefID: typeDef.ID,
			Meanings: []scopeir.ExportMeaning{scopeir.ExportMeaningType}, TypeOnly: true,
			Range:      scopeir.Range{StartLine: 5, StartCol: 0, EndLine: 5, EndCol: 24},
			Provenance: scopeir.ExportProvenance{StatementRange: scopeir.Range{StartLine: 5, EndLine: 5}, SiteKind: "export_specifier"},
		},
		{
			FilePath: filePath, FileHash: "hash-exports", Kind: scopeir.ExportDefault,
			ExportedName: "default", Meanings: []scopeir.ExportMeaning{scopeir.ExportMeaningValue},
			Range:      scopeir.Range{StartLine: 7, StartCol: 15, EndLine: 7, EndCol: 28},
			Provenance: scopeir.ExportProvenance{StatementRange: scopeir.Range{StartLine: 7, EndLine: 7}, SiteKind: "export_default_expression"},
		},
		{
			FilePath: filePath, FileHash: "hash-exports", Kind: scopeir.ExportReexport,
			ExportedName: "remote", TargetRaw: &targetRaw, TargetExportedName: "remote",
			Meanings:   []scopeir.ExportMeaning{scopeir.ExportMeaningValue},
			Range:      scopeir.Range{StartLine: 9, StartCol: 9, EndLine: 9, EndCol: 15},
			Provenance: scopeir.ExportProvenance{StatementRange: scopeir.Range{StartLine: 9, EndLine: 9}, SiteKind: "export_specifier"},
		},
		{
			FilePath: filePath, FileHash: "hash-exports", Kind: scopeir.ExportStar,
			TargetRaw: &targetRaw, Meanings: []scopeir.ExportMeaning{scopeir.ExportMeaningValue},
			Range:      scopeir.Range{StartLine: 10, StartCol: 7, EndLine: 10, EndCol: 8},
			Provenance: scopeir.ExportProvenance{StatementRange: scopeir.Range{StartLine: 10, EndLine: 10}, SiteKind: "export_star"},
		},
		{
			FilePath: filePath, FileHash: "hash-exports", Kind: scopeir.ExportNamespace,
			ExportedName: "ns", TargetRaw: &targetRaw,
			Meanings:   []scopeir.ExportMeaning{scopeir.ExportMeaningNamespace},
			Range:      scopeir.Range{StartLine: 11, StartCol: 13, EndLine: 11, EndCol: 15},
			Provenance: scopeir.ExportProvenance{StatementRange: scopeir.Range{StartLine: 11, EndLine: 11}, SiteKind: "export_namespace"},
		},
	}
	ir := scopeir.ScopeIR{
		FilePath: filePath, FileHash: "hash-exports", Language: scanner.TypeScript, ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{{
			ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: filePath,
			OwnedDefIDs: []string{runtimeDef.ID, typeDef.ID},
			Bindings: []scopeir.BindingFact{
				{Name: runtimeDef.Name, DefID: runtimeDef.ID, Origin: scopeir.BindingLocal},
				{Name: typeDef.Name, DefID: typeDef.ID, Origin: scopeir.BindingLocal},
			},
		}},
		Definitions: []scopeir.DefinitionFact{runtimeDef, typeDef}, Exports: facts,
	}

	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}

	exportNodes := make([]graph.Node, 0, len(facts))
	for _, node := range result.Graph.Nodes {
		if node.Label == exportNodeLabel {
			exportNodes = append(exportNodes, node)
		}
	}
	if len(exportNodes) != len(facts) {
		t.Fatalf("export graph records = %d, want %d", len(exportNodes), len(facts))
	}
	seenIDs := map[string]struct{}{}
	for _, node := range exportNodes {
		if _, duplicate := seenIDs[node.ID]; duplicate {
			t.Fatalf("duplicate export graph record %q", node.ID)
		}
		seenIDs[node.ID] = struct{}{}
		if node.Properties["filePath"] != filePath || node.Properties["fileHash"] != "hash-exports" {
			t.Fatalf("export provenance drift: %#v", node.Properties)
		}
		for key := range node.Properties {
			lower := strings.ToLower(key)
			if strings.Contains(lower, "terminal") || strings.Contains(lower, "publicapi") || strings.Contains(lower, "resolvedtarget") {
				t.Fatalf("export record invented later-slice state %q: %#v", key, node.Properties)
			}
		}
		if localDefID, ok := node.Properties["localDefId"].(string); ok && localDefID != "" {
			localNodeID, ok := node.Properties["localDefinitionNodeId"].(string)
			if !ok || localNodeID == "" {
				t.Fatalf("local export %q lost graph definition reference: %#v", localDefID, node.Properties)
			}
			if _, ok := result.Graph.GetNode(localNodeID); !ok {
				t.Fatalf("local export %q has orphan graph definition reference %q", localDefID, localNodeID)
			}
		}
	}

	runtimeNode := requireNode(t, result.Graph, "Variable", filePath, runtimeDef.Name)
	if got, ok := runtimeNode.Properties["isExported"].(bool); !ok || !got {
		t.Fatalf("runtime definition isExported = %#v, want true", runtimeNode.Properties["isExported"])
	}
	if runtimeNode.Properties["visibility"] != "private" {
		t.Fatalf("runtime access visibility changed: %#v", runtimeNode.Properties["visibility"])
	}
	typeNode := requireNode(t, result.Graph, "Interface", filePath, typeDef.Name)
	if got, ok := typeNode.Properties["isExported"].(bool); !ok || got {
		t.Fatalf("type-only definition isExported = %#v, want false", typeNode.Properties["isExported"])
	}
	if typeNode.Properties["visibility"] != "public" {
		t.Fatalf("type-only access visibility changed: %#v", typeNode.Properties["visibility"])
	}

	contains := 0
	for _, relationship := range result.Graph.Relationships {
		if relationship.Type == graph.RelContains && relationship.SourceID == "File:"+filePath {
			if node, ok := result.Graph.GetNode(relationship.TargetID); ok && node.Label == exportNodeLabel {
				contains++
			}
		}
	}
	if contains != len(facts) {
		t.Fatalf("file-to-export containment edges = %d, want %d", contains, len(facts))
	}
}

func TestP4CRejectsOrphanLocalExportFact(t *testing.T) {
	const filePath = "src/orphan.ts"
	const moduleScope = "scope:src/orphan.ts#1:0-2:0:Module"
	ir := scopeir.ScopeIR{
		FilePath: filePath, Language: scanner.TypeScript, ModuleScope: moduleScope,
		Scopes: []scopeir.ScopeFact{{ID: moduleScope, Kind: scopeir.ScopeModule, FilePath: filePath}},
		Exports: []scopeir.ExportFact{{
			FilePath: filePath, Kind: scopeir.ExportNamed, ExportedName: "missing", LocalName: "missing",
			LocalDefID: "def:src/orphan.ts#1:0:Variable:missing", Meanings: []scopeir.ExportMeaning{scopeir.ExportMeaningValue},
			Range:      scopeir.Range{StartLine: 1, EndLine: 1},
			Provenance: scopeir.ExportProvenance{StatementRange: scopeir.Range{StartLine: 1, EndLine: 1}, SiteKind: "export_specifier"},
		}},
	}
	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{})
	if err == nil || !strings.Contains(err.Error(), "references missing definition") {
		t.Fatalf("Resolve() error = %v, want fail-closed orphan reference", err)
	}
	if result.Graph != nil {
		t.Fatalf("Resolve() returned partial graph for orphan export: %#v", result.Graph)
	}
}
