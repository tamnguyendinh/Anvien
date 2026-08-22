//go:build ladybugdb

package lbugnative

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
	"github.com/tamnguyendinh/anvien/internal/lbugload"
	"github.com/tamnguyendinh/anvien/internal/lbugruntime"
	"github.com/tamnguyendinh/anvien/internal/lbugschema"
	"github.com/tamnguyendinh/anvien/internal/resolution"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestP6C3NativeLadybugResolutionOutcomeReadback(t *testing.T) {
	tempRoot := filepath.Join("..", "..", ".tmp", "p6c3-lbugnative-tests")
	if err := os.MkdirAll(tempRoot, 0o755); err != nil {
		t.Fatalf("create repo-local native temp root: %v", err)
	}
	tempDir, err := os.MkdirTemp(tempRoot, "outcome-")
	if err != nil {
		t.Fatalf("create repo-local native temp dir: %v", err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(tempDir) })

	dbPath := filepath.Join(tempDir, "p6c3-outcome.lbug")
	db, err := openNativeDatabase(dbPath, false)
	if err != nil {
		t.Fatalf("open writable outcome db: %v", err)
	}
	conn, err := db.OpenConnection()
	if err != nil {
		db.Close()
		t.Fatalf("open writable outcome connection: %v", err)
	}
	writeRunner := p6c3NativeWriteRunner{conn: conn}
	for _, query := range p6c3SchemaQueries(t) {
		if err := writeRunner.Query(query); err != nil {
			conn.Close()
			db.Close()
			t.Fatalf("outcome schema query failed:\n%s\n%v", query, err)
		}
	}

	resolved := resolution.ResolutionOutcome{
		SchemaVersion:    resolution.ResolutionOutcomeSchemaVersion,
		SourceSiteID:     "SourceSite:src/app.ts#call#helper#3#2#3#10",
		Status:           resolution.ResolutionResolvedInternal,
		Stage:            resolution.ResolutionStageRepository,
		SiteKind:         "call",
		FilePath:         "src/app.ts",
		FileHash:         "hash-app",
		Range:            scopeir.Range{StartLine: 3, StartCol: 2, EndLine: 3, EndCol: 10},
		RequestedName:    "helper",
		RequestedMeaning: "value",
		Language:         "typescript",
		Target: &resolution.ResolutionTarget{
			ID: "def:helper", GraphID: "Function:helper", Name: "helper", Kind: string(scopeir.NodeFunction), FilePath: "src/app.ts",
		},
		Proof: resolution.ResolutionProof{Kind: "scope-binding"},
	}
	unresolved := resolution.ResolutionOutcome{
		SchemaVersion:    resolution.ResolutionOutcomeSchemaVersion,
		SourceSiteID:     "SourceSite:src/app.ts#call#missing#4#2#4#11",
		Status:           resolution.ResolutionUnresolved,
		Stage:            resolution.ResolutionStageRepository,
		SiteKind:         "call",
		FilePath:         "src/app.ts",
		FileHash:         "hash-app",
		Range:            scopeir.Range{StartLine: 4, StartCol: 2, EndLine: 4, EndCol: 11},
		RequestedName:    "missing",
		RequestedMeaning: "value",
		Language:         "typescript",
		Reason:           "call target not resolved",
		Proof:            resolution.ResolutionProof{Kind: "none"},
	}
	resolvedJSON := mustP6C3OutcomeJSON(t, resolved)
	unresolvedJSON := mustP6C3OutcomeJSON(t, unresolved)

	g := graph.New()
	g.AddNode(graph.Node{ID: "Function:caller", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name": "caller", "filePath": "src/app.ts", "startLine": 1, "endLine": 6,
	}})
	g.AddNode(graph.Node{ID: "Function:helper", Label: scopeir.NodeFunction, Properties: graph.NodeProperties{
		"name": "helper", "filePath": "src/app.ts", "startLine": 1, "endLine": 1,
	}})
	g.AddRelationship(graph.Relationship{
		ID: "rel:CALLS:caller->helper", SourceID: "Function:caller", TargetID: "Function:helper", Type: graph.RelCalls,
		Confidence: 1, ResolutionSource: resolution.ResolutionStageRepository,
		SourceSiteID: resolved.SourceSiteID, SourceSiteIDs: []string{resolved.SourceSiteID}, SourceSiteCount: 1,
		SourceSiteStatus: "resolved", ProofKind: resolved.Proof.Kind, TargetRole: "callable", TargetText: resolved.RequestedName,
		FilePath: resolved.FilePath, FileHash: resolved.FileHash,
		StartLine: resolved.Range.StartLine, StartCol: resolved.Range.StartCol, EndLine: resolved.Range.EndLine, EndCol: resolved.Range.EndCol,
		Evidence: []graph.Evidence{{Kind: resolution.ResolutionOutcomeEvidenceKind, Weight: 1, Note: resolvedJSON}},
	})
	gapInput := graphhealth.ResolutionGapInput{
		SourceSiteID: unresolved.SourceSiteID, SourceNodeID: "Function:caller", SourceNodeLabel: string(scopeir.NodeFunction),
		SourceAppLayer: "backend", SourceFunctionalArea: "resolution", FactFamily: unresolved.SiteKind,
		TargetText: unresolved.RequestedName, TargetRole: "callable", SourceSiteStatus: string(unresolved.Status),
		ProofKind: unresolved.Proof.Kind, ResolutionSource: unresolved.Stage, Source: unresolved.Stage,
		FilePath: unresolved.FilePath, FileHash: unresolved.FileHash,
		StartLine: unresolved.Range.StartLine, StartCol: unresolved.Range.StartCol, EndLine: unresolved.Range.EndLine, EndCol: unresolved.Range.EndCol,
		Count: 1, Note: unresolvedJSON,
	}
	g.AddNode(gapInput.GraphNode())
	g.AddRelationship(gapInput.GraphRelationship())

	exported, err := lbugload.ExportGraphCSVs(g, filepath.Join(tempDir, "csv"))
	if err != nil {
		conn.Close()
		db.Close()
		t.Fatalf("export outcome graph: %v", err)
	}
	loadResult, err := lbugload.LoadCSVExport(writeRunner, exported)
	if err != nil {
		conn.Close()
		db.Close()
		t.Fatalf("load outcome graph: %v", err)
	}
	if loadResult.NodeCopyCount != 2 || loadResult.RelationshipCopyCount != 2 ||
		loadResult.FallbackInsertCount != 0 || loadResult.FallbackInsertFailures != 0 ||
		loadResult.SkippedRelationships != 0 {
		conn.Close()
		db.Close()
		t.Fatalf("outcome native load drift: %#v", loadResult)
	}
	conn.Close()
	db.Close()

	readDB, err := openNativeDatabase(dbPath, true)
	if err != nil {
		t.Fatalf("reopen outcome db read-only: %v", err)
	}
	defer readDB.Close()
	readConn, err := readDB.OpenConnection()
	if err != nil {
		t.Fatalf("open outcome read connection: %v", err)
	}
	defer readConn.Close()
	reader := p6c3NativeReadRunner{conn: readConn, silencer: &lbugruntime.StdioSilencer{}}

	relationRows, err := reader.Query("MATCH (a)-[r:CodeRelation]->(b) RETURN r.type AS type, r.sourceSiteId AS sourceSiteId, r.sourceSiteStatus AS sourceSiteStatus, r.evidence AS evidence")
	if err != nil {
		t.Fatalf("read outcome relations: %v", err)
	}
	foundResolved := false
	for _, row := range relationRows {
		if row["sourceSiteId"] != resolved.SourceSiteID {
			continue
		}
		var evidence []graph.Evidence
		if err := json.Unmarshal([]byte(row["evidence"].(string)), &evidence); err != nil {
			t.Fatalf("decode native resolved evidence: %v", err)
		}
		if len(evidence) != 1 || evidence[0].Kind != resolution.ResolutionOutcomeEvidenceKind {
			t.Fatalf("native resolved evidence shape = %#v", evidence)
		}
		var got resolution.ResolutionOutcome
		if err := json.Unmarshal([]byte(evidence[0].Note), &got); err != nil || !reflect.DeepEqual(got, resolved) {
			t.Fatalf("native resolved outcome drift: err=%v want=%#v got=%#v", err, resolved, got)
		}
		foundResolved = true
	}
	if !foundResolved {
		t.Fatalf("native relation readback missing resolved source site %q: %#v", resolved.SourceSiteID, relationRows)
	}

	gapRows, err := reader.Query("MATCH (n:ResolutionGap) RETURN n.sourceSiteId AS sourceSiteId, n.sourceSiteStatus AS sourceSiteStatus, n.note AS note")
	if err != nil {
		t.Fatalf("read native ResolutionGap outcomes: %v", err)
	}
	if len(gapRows) != 1 || gapRows[0]["sourceSiteId"] != unresolved.SourceSiteID ||
		gapRows[0]["sourceSiteStatus"] != string(resolution.ResolutionUnresolved) {
		t.Fatalf("native ResolutionGap identity/status drift: %#v", gapRows)
	}
	var gotUnresolved resolution.ResolutionOutcome
	if err := json.Unmarshal([]byte(gapRows[0]["note"].(string)), &gotUnresolved); err != nil || !reflect.DeepEqual(gotUnresolved, unresolved) {
		t.Fatalf("native unresolved outcome drift: err=%v want=%#v got=%#v", err, unresolved, gotUnresolved)
	}
}

func mustP6C3OutcomeJSON(t *testing.T, outcome resolution.ResolutionOutcome) string {
	t.Helper()
	raw, err := json.Marshal(outcome)
	if err != nil {
		t.Fatalf("marshal outcome: %v", err)
	}
	return string(raw)
}

type p6c3NativeWriteRunner struct {
	conn *nativeConnection
}

func (runner p6c3NativeWriteRunner) Query(query string) error {
	result, err := runner.conn.Query(query)
	if result != nil {
		result.Close()
	}
	return err
}

type p6c3NativeReadRunner struct {
	conn     *nativeConnection
	silencer *lbugruntime.StdioSilencer
}

func (runner p6c3NativeReadRunner) Query(query string) ([]lbugruntime.Row, error) {
	if err := lbugruntime.ValidateReadQuery(query); err != nil {
		return nil, err
	}
	var rows []lbugruntime.Row
	err := runner.silencer.Run(func() error {
		result, err := runner.conn.Query(query)
		if err != nil {
			return err
		}
		defer result.Close()
		rows, err = result.Rows()
		return err
	})
	return rows, err
}

func p6c3SchemaQueries(t *testing.T) []string {
	t.Helper()
	queries, err := lbugschema.SchemaQueries(lbugschema.DefaultEmbeddingDims)
	if err != nil {
		t.Fatalf("build native schema: %v", err)
	}
	return queries
}
