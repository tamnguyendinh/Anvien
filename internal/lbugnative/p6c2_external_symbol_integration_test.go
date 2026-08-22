//go:build ladybugdb

package lbugnative

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/lbugload"
	"github.com/tamnguyendinh/anvien/internal/lbugruntime"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestP6C2NativeLadybugExternalSymbolReadback(t *testing.T) {
	dbPath := filepath.Join(t.TempDir(), "p6c2-external.lbug")
	db, err := openNativeDatabase(dbPath, false)
	if err != nil {
		t.Fatalf("open writable external-symbol db: %v", err)
	}
	conn, err := db.OpenConnection()
	if err != nil {
		db.Close()
		t.Fatalf("open writable external-symbol connection: %v", err)
	}
	writeRunner := nativeWriteRunner{conn: conn}
	for _, query := range schemaQueries(t) {
		if err := writeRunner.Query(query); err != nil {
			conn.Close()
			db.Close()
			t.Fatalf("external-symbol schema query failed:\n%s\n%v", query, err)
		}
	}

	externalID := "ExternalSymbol:tsstdlib:math-max"
	sourceSiteID := "SourceSite:src/app.ts#call#Math.max#7#2#7#10"
	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:run", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name": "run", "filePath": "src/app.ts", "startLine": 1, "endLine": 10,
	}})
	g.AddNode(graph.Node{ID: externalID, Label: scopeir.NodeExternalSymbol, Properties: graph.NodeProperties{
		"name":                 "Math.max",
		"qualifiedName":        "Math.max",
		"requestedNames":       []string{"Math.max"},
		"requestedTargetTexts": []string{"Math.max"},
		"meaning":              "value",
		"meanings":             []string{"value"},
		"semanticSymbolId":     "tsstdlib:math-max",
		"semanticOwnerId":      "tsstdlib:math",
		"authorityKind":        "typescript_standard_library",
		"typeScriptVersion":    "5.9.3",
		"catalogProofState":    "ready",
		"authorityHash":        "authority-hash",
		"catalogHash":          "catalog-hash",
		"catalogArtifactHash":  "artifact-hash",
		"profileHashes":        []string{"profile-hash"},
		"configHashes":         []string{"config-hash"},
		"declarationLibraries": []string{"lib.es2015.core.d.ts"},
		"declarationRanges": []map[string]any{{
			"library": "lib.es2015.core.d.ts", "startLine": 120, "startCol": 5, "endLine": 120, "endCol": 48,
		}},
		"sourceSiteIds":   []string{sourceSiteID},
		"sourceSiteCount": 1,
		"origin":          "typescript_standard_library",
		"external":        true,
		"editable":        false,
		"repositoryOwned": false,
	}})
	g.AddRelationship(graph.Relationship{
		ID: "rel:CALLS:run->Math.max", SourceID: "Function:run", TargetID: externalID,
		Type: graph.RelCalls, Confidence: 1, ResolutionSource: "typescript_standard_library",
		SourceSiteID: sourceSiteID, SourceSiteIDs: []string{sourceSiteID}, SourceSiteCount: 1,
		SourceSiteStatus: "resolved", ProofKind: "typescript-standard-library-authority",
		TargetRole: "callable", TargetText: "Math.max", FilePath: "src/app.ts",
		StartLine: 7, StartCol: 2, EndLine: 7, EndCol: 10,
		Evidence: []graph.Evidence{{Kind: "typescript-authority-result-v1", Weight: 1, Note: `{"sourceSiteId":"` + sourceSiteID + `"}`}},
	})
	export, err := lbugload.ExportGraphCSVs(g, filepath.Join(t.TempDir(), "csv"))
	if err != nil {
		conn.Close()
		db.Close()
		t.Fatalf("export external-symbol graph: %v", err)
	}
	loadResult, err := lbugload.LoadCSVExport(writeRunner, export)
	if err != nil {
		conn.Close()
		db.Close()
		t.Fatalf("load external-symbol graph: %v", err)
	}
	if loadResult.NodeCopyCount != 2 || loadResult.RelationshipCopyCount != 1 ||
		loadResult.FallbackInsertCount != 0 || loadResult.FallbackInsertFailures != 0 ||
		loadResult.SkippedRelationships != 0 {
		conn.Close()
		db.Close()
		t.Fatalf("external-symbol native load drift: %#v", loadResult)
	}
	conn.Close()
	db.Close()

	readDB, err := openNativeDatabase(dbPath, true)
	if err != nil {
		t.Fatalf("reopen external-symbol db read-only: %v", err)
	}
	defer readDB.Close()
	readConn, err := readDB.OpenConnection()
	if err != nil {
		t.Fatalf("open external-symbol read connection: %v", err)
	}
	defer readConn.Close()
	reader := nativeReadRunner{conn: readConn, silencer: &lbugruntime.StdioSilencer{}}
	rows, err := reader.Query("MATCH (n:ExternalSymbol) RETURN n.id AS id, n.semanticSymbolId AS semanticSymbolId, n.semanticOwnerId AS semanticOwnerId, n.authorityKind AS authorityKind, n.typeScriptVersion AS typeScriptVersion, n.catalogProofState AS catalogProofState, n.declarationRanges AS declarationRanges, n.sourceSiteCount AS sourceSiteCount, n.origin AS origin, n.external AS external, n.editable AS editable, n.repositoryOwned AS repositoryOwned")
	if err != nil {
		t.Fatalf("read external-symbol row: %v", err)
	}
	if len(rows) != 1 || rows[0]["id"] != externalID ||
		rows[0]["semanticSymbolId"] != "tsstdlib:math-max" || rows[0]["semanticOwnerId"] != "tsstdlib:math" ||
		rows[0]["authorityKind"] != "typescript_standard_library" || rows[0]["typeScriptVersion"] != "5.9.3" ||
		rows[0]["catalogProofState"] != "ready" || rows[0]["sourceSiteCount"] != "1" ||
		rows[0]["origin"] != "typescript_standard_library" || !nativeBoolTrue(rows[0]["external"]) ||
		nativeBoolTrue(rows[0]["editable"]) || nativeBoolTrue(rows[0]["repositoryOwned"]) ||
		!strings.Contains(rows[0]["declarationRanges"].(string), `"library":"lib.es2015.core.d.ts"`) {
		t.Fatalf("native external-symbol readback lost provenance/non-editable fields: %#v", rows)
	}
	relationRows, err := reader.Query("MATCH (a:Function)-[r:CodeRelation]->(b:ExternalSymbol) RETURN a.id AS fromId, b.id AS toId, r.type AS type, r.resolutionSource AS resolutionSource, r.sourceSiteId AS sourceSiteId, r.proofKind AS proofKind, r.evidence AS evidence")
	if err != nil {
		t.Fatalf("read external-symbol relation: %v", err)
	}
	if len(relationRows) != 1 || relationRows[0]["fromId"] != "Function:run" || relationRows[0]["toId"] != externalID ||
		relationRows[0]["type"] != "CALLS" || relationRows[0]["resolutionSource"] != "typescript_standard_library" ||
		relationRows[0]["sourceSiteId"] != sourceSiteID || relationRows[0]["proofKind"] != "typescript-standard-library-authority" ||
		!strings.Contains(relationRows[0]["evidence"].(string), "typescript-authority-result-v1") {
		t.Fatalf("native external relation lost source-site proof: %#v", relationRows)
	}
}
