//go:build resolution_parser_integration

package resolution

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"reflect"
	"slices"
	"testing"
	"time"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/parser"
	"github.com/tamnguyendinh/anvien/internal/providers/tsjs"
	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func TestResolveTypeScriptGraphFixture(t *testing.T) {
	irs := parseFixtureWorkspace(t)
	result, err := Resolve(irs, Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}

	serviceSave := requireNode(t, result.Graph, "Method", "src/service.ts", "Service.save")
	format := requireNode(t, result.Graph, "Function", "src/user.ts", "format")
	repoWrite := requireNode(t, result.Graph, "Method", "src/repo.ts", "Repo.write")
	repoClass := requireNode(t, result.Graph, "Class", "src/repo.ts", "Repo")
	modelClass := requireNode(t, result.Graph, "Class", "src/model.ts", "Model")
	userID := requireNode(t, result.Graph, "Property", "src/user.ts", "User.id")
	serviceRepo := requireNode(t, result.Graph, "Property", "src/service.ts", "Service.repo")
	named := requireNode(t, result.Graph, "Interface", "src/contracts.ts", "Named")
	baseService := requireNode(t, result.Graph, "Class", "src/contracts.ts", "BaseService")

	requireRelationship(t, result.Graph, graph.RelCalls, serviceSave.ID, format.ID)
	requireRelationship(t, result.Graph, graph.RelCalls, serviceSave.ID, repoWrite.ID)
	requireRelationship(t, result.Graph, graph.RelCalls, serviceSave.ID, modelClass.ID)
	requireRelationship(t, result.Graph, graph.RelCalls, requireNode(t, result.Graph, "Function", "src/service.ts", "makeRepo").ID, repoClass.ID)
	requireRelationship(t, result.Graph, graph.RelAccesses, serviceSave.ID, userID.ID)
	requireRelationship(t, result.Graph, graph.RelAccesses, serviceSave.ID, serviceRepo.ID)
	requireRelationship(t, result.Graph, graph.RelExtends, requireNode(t, result.Graph, "Class", "src/service.ts", "Service").ID, baseService.ID)
	requireRelationship(t, result.Graph, graph.RelImplements, requireNode(t, result.Graph, "Class", "src/service.ts", "Service").ID, named.ID)
	requireRelationship(t, result.Graph, graph.RelMethodOverrides, requireNode(t, result.Graph, "Class", "src/service.ts", "Service").ID, requireNode(t, result.Graph, "Method", "src/contracts.ts", "BaseService.save").ID)
	requireRelationship(t, result.Graph, graph.RelMethodImplements, serviceSave.ID, requireNode(t, result.Graph, "Method", "src/contracts.ts", "Named.save").ID)

	counts := result.Graph.RelationshipCountsByType()
	for _, relType := range []graph.RelationshipType{
		graph.RelCalls,
		graph.RelImports,
		graph.RelAccesses,
		graph.RelExtends,
		graph.RelImplements,
		graph.RelInherits,
		graph.RelUses,
		graph.RelHasMethod,
		graph.RelHasProperty,
		graph.RelMethodOverrides,
		graph.RelMethodImplements,
	} {
		if counts[relType] == 0 {
			t.Fatalf("expected %s relationships in graph, counts=%v", relType, counts)
		}
	}
	if result.Metrics.ResolvedCalls < 5 || result.Metrics.ResolvedAccesses < 2 || result.Metrics.ResolvedInheritance != 2 {
		t.Fatalf("unexpected metrics: %#v", result.Metrics)
	}
	if len(result.ReferenceIndex.BySourceScope) == 0 || len(result.ReferenceIndex.ByTargetDef) == 0 {
		t.Fatalf("expected reference index to be populated: %#v", result.ReferenceIndex)
	}
}

func TestResolveTypeScriptInterfaceHeritageFromSource(t *testing.T) {
	source := []byte(`export interface Area { id: string; }
export interface AreaWithTableCount extends Area { count: number; }
export interface Shift { id: string; }
export interface ShiftWithCounts extends Shift { assignmentCount: number; }
export interface ExternalBacked extends React.ComponentProps<"button"> { label: string; }
`)
	ir := parseTypeScriptSource(t, "src/types.ts", source)
	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}

	area := requireNode(t, result.Graph, "Interface", "src/types.ts", "Area")
	areaWithCount := requireNode(t, result.Graph, "Interface", "src/types.ts", "AreaWithTableCount")
	shift := requireNode(t, result.Graph, "Interface", "src/types.ts", "Shift")
	shiftWithCounts := requireNode(t, result.Graph, "Interface", "src/types.ts", "ShiftWithCounts")

	requireRelationship(t, result.Graph, graph.RelExtends, areaWithCount.ID, area.ID)
	requireRelationship(t, result.Graph, graph.RelExtends, shiftWithCounts.ID, shift.ID)
	requireRelationship(t, result.Graph, graph.RelInherits, areaWithCount.ID, area.ID)
	requireRelationship(t, result.Graph, graph.RelInherits, shiftWithCounts.ID, shift.ID)
	if result.Metrics.HeritageFactsIndexed != 3 ||
		result.Metrics.ResolvedInheritance != 2 ||
		result.Metrics.UnresolvedInheritance != 1 {
		t.Fatalf("unexpected heritage metrics: %#v", result.Metrics)
	}
}

func TestResolveTypeScriptHeritagePrefersSameFileTargetWhenGlobalNameAmbiguous(t *testing.T) {
	local := parseTypeScriptSource(t, "src/types/area.ts", []byte(`export interface Area { id: string; }
export interface AreaWithTableCount extends Area { tableCount: number; }
`))
	other := parseTypeScriptSource(t, "src/features/tables/types.ts", []byte(`export interface Area { id: string; }
`))

	result, err := Resolve([]scopeir.ScopeIR{local, other}, Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}

	localArea := requireNode(t, result.Graph, "Interface", "src/types/area.ts", "Area")
	otherArea := requireNode(t, result.Graph, "Interface", "src/features/tables/types.ts", "Area")
	areaWithCount := requireNode(t, result.Graph, "Interface", "src/types/area.ts", "AreaWithTableCount")

	requireRelationship(t, result.Graph, graph.RelExtends, areaWithCount.ID, localArea.ID)
	requireRelationship(t, result.Graph, graph.RelInherits, areaWithCount.ID, localArea.ID)
	requireNoRelationship(t, result.Graph, graph.RelExtends, areaWithCount.ID, otherArea.ID)
	if result.Metrics.HeritageFactsIndexed != 1 ||
		result.Metrics.ResolvedInheritance != 1 ||
		result.Metrics.UnresolvedInheritance != 0 {
		t.Fatalf("unexpected heritage metrics: %#v", result.Metrics)
	}
}

func TestResolveAwaitedPromiseReturnMemberAccess(t *testing.T) {
	source := []byte(`type Invoice = {
  invoiceId: string;
};

type InvoiceModel = {
  invoices: Invoice[];
};

type ReadResult = {
  model: InvoiceModel;
};

async function readResult(): Promise<ReadResult> {
  throw new Error("not implemented");
}

async function run() {
  const result = await readResult();
  result.model.invoices;
}
`)
	ir := parseTypeScriptSource(t, "src/awaited-result.ts", source)
	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}

	run := requireNode(t, result.Graph, "Function", "src/awaited-result.ts", "run")
	model := requireNode(t, result.Graph, "Property", "src/awaited-result.ts", "ReadResult.model")
	invoices := requireNode(t, result.Graph, "Property", "src/awaited-result.ts", "InvoiceModel.invoices")

	requireRelationship(t, result.Graph, graph.RelAccesses, run.ID, model.ID)
	requireRelationship(t, result.Graph, graph.RelAccesses, run.ID, invoices.ID)
}

func TestBuildCrossFileBindingFeedsResolveBoundInto(t *testing.T) {
	binding, err := BuildCrossFileBinding(parseFixtureWorkspace(t), Options{})
	if err != nil {
		t.Fatalf("BuildCrossFileBinding() error = %v", err)
	}
	if binding.Metrics.DefinitionsIndexed == 0 || binding.Metrics.ImportsResolved == 0 {
		t.Fatalf("binding metrics missing index/import data: %#v", binding.Metrics)
	}
	if !binding.Metrics.BindingAccumulatorFinalized {
		t.Fatalf("binding accumulator metrics not finalized: %#v", binding.Metrics)
	}

	result, err := ResolveBoundInto(nil, binding, Options{})
	if err != nil {
		t.Fatalf("ResolveBoundInto() error = %v", err)
	}
	if !result.Metrics.BindingAccumulatorDisposed {
		t.Fatalf("binding accumulator was not disposed: %#v", result.Metrics)
	}
	if result.Metrics.ResolvedCalls == 0 || result.Metrics.GraphRelationshipsEmitted == 0 {
		t.Fatalf("resolution metrics missing semantic output: %#v", result.Metrics)
	}
}

func TestResolveTypeScriptGraphSignatureFixture(t *testing.T) {
	result, err := Resolve(parseFixtureWorkspace(t), Options{})
	if err != nil {
		t.Fatalf("resolve failed: %v", err)
	}
	signature := buildGraphSignature(result.Graph)
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(signature); err != nil {
		t.Fatalf("marshal signature: %v", err)
	}
	golden, err := os.ReadFile("testdata/typescript_graph_signature.golden.json")
	if err != nil {
		t.Fatalf("read golden: %v", err)
	}
	if string(buffer.Bytes()) != string(golden) {
		t.Fatalf("graph signature mismatch\nwant:\n%s\ngot:\n%s", golden, buffer.Bytes())
	}
}

func BenchmarkResolveTypeScriptGraphFixture(b *testing.B) {
	irs := parseFixtureWorkspaceForBenchmark(b)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		result, err := Resolve(irs, Options{})
		if err != nil {
			b.Fatalf("resolve failed: %v", err)
		}
		if len(result.Graph.Relationships) == 0 {
			b.Fatalf("empty graph")
		}
	}
}

func TestResolveSkipCompatibilityCrossFileReportsDiagnosticWithoutChangingGraph(t *testing.T) {
	irs := parseFixtureWorkspace(t)
	defaultResult, err := Resolve(irs, Options{})
	if err != nil {
		t.Fatalf("default resolve failed: %v", err)
	}
	skipResult, err := Resolve(irs, Options{SkipCompatibilityCrossFile: true})
	if err != nil {
		t.Fatalf("skip resolve failed: %v", err)
	}

	if !reflect.DeepEqual(buildGraphSignature(defaultResult.Graph), buildGraphSignature(skipResult.Graph)) {
		t.Fatalf("skipCompatibilityCrossFile changed graph signature")
	}
	if !defaultResult.Metrics.CrossFileSkipped || defaultResult.Metrics.CrossFileFilesReprocessed != 0 {
		t.Fatalf("unexpected default cross-file metrics: %#v", defaultResult.Metrics)
	}
	if defaultResult.Metrics.CrossFileSkipReason != "covered-by-scopeir-single-pass-resolution" {
		t.Fatalf("default skip reason = %q", defaultResult.Metrics.CrossFileSkipReason)
	}
	if !skipResult.Metrics.CrossFileSkipped || skipResult.Metrics.CrossFileFilesReprocessed != 0 {
		t.Fatalf("unexpected skip cross-file metrics: %#v", skipResult.Metrics)
	}
	if skipResult.Metrics.CrossFileSkipReason != "disabled-by-pipeline-option" {
		t.Fatalf("skip reason = %q", skipResult.Metrics.CrossFileSkipReason)
	}
	if !skipResult.Metrics.BindingAccumulatorFinalized || !skipResult.Metrics.BindingAccumulatorDisposed {
		t.Fatalf("accumulator lifecycle metrics not closed: %#v", skipResult.Metrics)
	}
}

func TestP1CIdentityConservesSameNameOccurrencesAndEndpoints(t *testing.T) {
	ir := p1cIdentityFixtureIR(t)
	workspace, err := buildWorkspace([]scopeir.ScopeIR{ir})
	if err != nil {
		t.Fatalf("buildWorkspace() error = %v", err)
	}

	filePath := cleanPath(ir.FilePath)
	definitions := workspace.defsByFile[filePath]
	if len(definitions) != len(ir.Definitions) {
		t.Fatalf("defsByFile occurrence denominator = %d, want %d", len(definitions), len(ir.Definitions))
	}
	graphIDs := make(map[string]struct{}, len(definitions))
	for _, definition := range definitions {
		graphIDs[definition.GraphID] = struct{}{}
	}
	if len(graphIDs) != len(definitions) {
		t.Fatalf("unique graph identities = %d, want %d occurrences", len(graphIDs), len(definitions))
	}

	for _, name := range []string{"time", "now"} {
		refs := exactNamedDefinitions(workspace.defsByName[name], name)
		if len(refs) != 2 {
			t.Fatalf("%s occurrences = %d, want 2: %#v", name, len(refs), refs)
		}
		occurrences := map[string]struct{}{}
		scopes := map[string]struct{}{}
		ids := map[string]struct{}{}
		for _, ref := range refs {
			occurrences[ref.Fact.ID] = struct{}{}
			scopes[workspace.scopeByDef[ref.Fact.ID]] = struct{}{}
			ids[ref.GraphID] = struct{}{}
		}
		if len(occurrences) != 2 || len(scopes) != 2 || len(ids) != 2 {
			t.Fatalf("%s identity inputs collapsed: occurrences=%v scopes=%v graphIDs=%v", name, occurrences, scopes, ids)
		}
	}

	shared := exactNamedDefinitions(workspace.defsByName["Shared"], "Shared")
	if len(shared) != 2 {
		t.Fatalf("Shared meaning-lane occurrences = %d, want 2: %#v", len(shared), shared)
	}
	labels := map[scopeir.NodeLabel]struct{}{}
	sharedIDs := map[string]struct{}{}
	for _, ref := range shared {
		labels[ref.Fact.Label] = struct{}{}
		sharedIDs[ref.GraphID] = struct{}{}
	}
	if _, ok := labels[scopeir.NodeInterface]; !ok {
		t.Fatalf("Shared labels = %v, missing Interface", labels)
	}
	if _, ok := labels[scopeir.NodeFunction]; !ok {
		t.Fatalf("Shared labels = %v, missing Function", labels)
	}
	if len(sharedIDs) != 2 {
		t.Fatalf("Shared meaning lanes share graph identity: %v", sharedIDs)
	}

	result, err := Resolve([]scopeir.ScopeIR{ir}, Options{})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	definedTargets := map[string]struct{}{}
	for _, relationship := range result.Graph.Relationships {
		if _, ok := result.Graph.GetNode(relationship.SourceID); !ok {
			t.Fatalf("relationship %s missing source endpoint %s", relationship.ID, relationship.SourceID)
		}
		if _, ok := result.Graph.GetNode(relationship.TargetID); !ok {
			t.Fatalf("relationship %s missing target endpoint %s", relationship.ID, relationship.TargetID)
		}
		if relationship.Type == graph.RelDefines {
			definedTargets[relationship.TargetID] = struct{}{}
		}
	}
	for _, definition := range definitions {
		if _, ok := result.Graph.GetNode(definition.GraphID); !ok {
			t.Fatalf("definition occurrence %s missing graph node %s", definition.Fact.ID, definition.GraphID)
		}
		if _, ok := definedTargets[definition.GraphID]; !ok {
			t.Fatalf("definition occurrence %s missing DEFINES endpoint %s", definition.Fact.ID, definition.GraphID)
		}
	}
}

func TestP1CIdentityIsDeterministicAcrossDefinitionOrder(t *testing.T) {
	first := p1cIdentityFixtureIR(t)
	second := p1cIdentityFixtureIR(t)
	slices.Reverse(second.Definitions)
	slices.Reverse(second.Scopes)
	for index := range second.Scopes {
		slices.Reverse(second.Scopes[index].OwnedDefIDs)
		slices.Reverse(second.Scopes[index].Bindings)
	}

	firstWorkspace, err := buildWorkspace([]scopeir.ScopeIR{first})
	if err != nil {
		t.Fatalf("first buildWorkspace() error = %v", err)
	}
	secondWorkspace, err := buildWorkspace([]scopeir.ScopeIR{second})
	if err != nil {
		t.Fatalf("second buildWorkspace() error = %v", err)
	}
	firstIDs := graphIDsByOccurrence(firstWorkspace)
	secondIDs := graphIDsByOccurrence(secondWorkspace)
	if !reflect.DeepEqual(firstIDs, secondIDs) {
		t.Fatalf("reordered identity set changed\nfirst=%v\nsecond=%v", firstIDs, secondIDs)
	}
}

func p1cIdentityFixtureIR(t *testing.T) scopeir.ScopeIR {
	t.Helper()
	source, err := os.ReadFile("testdata/p1c_identity_repo/src/identity.ts")
	if err != nil {
		t.Fatalf("read P1-C identity fixture: %v", err)
	}
	return parseTypeScriptSource(t, "src/identity.ts", source)
}

func parseFixtureWorkspace(t *testing.T) []scopeir.ScopeIR {
	t.Helper()
	pool := parser.NewPool(nil, parser.PoolOptions{ParseTimeout: time.Second})
	defer pool.Close()
	return parseFixtureWorkspaceWithPool(t, pool)
}

func parseFixtureWorkspaceForBenchmark(b *testing.B) []scopeir.ScopeIR {
	b.Helper()
	pool := parser.NewPool(nil, parser.PoolOptions{ParseTimeout: time.Second})
	defer pool.Close()
	return parseFixtureWorkspaceWithPool(b, pool)
}

func parseTypeScriptSource(t *testing.T, filePath string, source []byte) scopeir.ScopeIR {
	t.Helper()
	pool := parser.NewPool(nil, parser.PoolOptions{ParseTimeout: time.Second})
	defer pool.Close()
	parsed, err := pool.Parse(context.Background(), parser.Request{
		FilePath: filePath,
		Language: scanner.TypeScript,
		Source:   source,
	})
	if err != nil {
		t.Fatalf("parse %s failed: %v", filePath, err)
	}
	defer parsed.Close()
	ir, err := tsjs.Extract(tsjs.Request{
		FilePath: filePath,
		FileHash: "hash-" + filePath,
		Language: scanner.TypeScript,
		Source:   source,
		Root:     parsed.Tree.RootNode(),
	})
	if err != nil {
		t.Fatalf("extract %s failed: %v", filePath, err)
	}
	return ir
}

type testingFataler interface {
	Helper()
	Fatalf(string, ...any)
}

func parseFixtureWorkspaceWithPool(t testingFataler, pool *parser.Pool) []scopeir.ScopeIR {
	t.Helper()
	sources := map[string]string{
		"src/user.ts":      `export default class User { id: string; } export function format(id: string): string { return id; }`,
		"src/repo.ts":      `export class Repo { write(value: string): void {} }`,
		"src/model.ts":     `export class Model {}`,
		"src/contracts.ts": `import User from './user'; export interface Named { id: string; save(user: User): Promise<void>; } export class BaseService { save(user: User): Promise<void> {} }`,
		"src/service.ts": `import User, { format as fmt } from './user';
import { Repo } from './repo';
import { Model } from './model';
import { BaseService, Named } from './contracts';

function makeRepo(): Repo { return new Repo(); }

class Service extends BaseService implements Named {
  public repo: Repo;
  constructor(repo: Repo) { this.repo = repo; }
  save(user: User): Promise<void> {
    const model = new Model();
    const made = makeRepo();
    const formatted = fmt(user.id);
    this.repo.write(formatted);
  }
}
`,
	}
	paths := []string{"src/user.ts", "src/repo.ts", "src/model.ts", "src/contracts.ts", "src/service.ts"}
	irs := make([]scopeir.ScopeIR, 0, len(paths))
	for _, filePath := range paths {
		source := []byte(sources[filePath])
		parsed, err := pool.Parse(context.Background(), parser.Request{
			FilePath: filePath,
			Language: scanner.TypeScript,
			Source:   source,
		})
		if err != nil {
			t.Fatalf("parse %s failed: %v", filePath, err)
		}
		ir, err := tsjs.Extract(tsjs.Request{
			FilePath: filePath,
			FileHash: "hash-" + filePath,
			Language: scanner.TypeScript,
			Source:   source,
			Root:     parsed.Tree.RootNode(),
		})
		parsed.Close()
		if err != nil {
			t.Fatalf("extract %s failed: %v", filePath, err)
		}
		irs = append(irs, ir)
	}
	return irs
}
