package tsjs

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	sitter "github.com/tree-sitter/go-tree-sitter"

	"github.com/tamnguyendinh/anvien/internal/parser"
	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

const typescriptParityFixture = `import User, { format as fmt } from './user';
export { Service as UserService } from './service';

interface Named { id: string; }
type UserLike = User | Named;

function makeRepo(): Repo { return new Repo(); }

class Service extends BaseService implements Named {
  public repo: Repo;
  constructor(repo: Repo) { this.repo = repo; }
  async save(user: User): Promise<void> {
    const model = new Model();
    const made = makeRepo();
    const formatted = fmt(user.id);
    this.repo.write(formatted);
  }
}
`

func TestExtractTypeScriptScopeIR(t *testing.T) {
	source := []byte(typescriptParityFixture)
	ir := parseAndExtract(t, "src/service.ts", "hash-ts", scanner.TypeScript, source)

	for _, def := range ir.Definitions {
		if def.FileHash != "hash-ts" {
			t.Fatalf("definition %s missing file hash: %#v", def.Name, def)
		}
	}
	for _, call := range ir.Calls {
		if call.FileHash != "hash-ts" {
			t.Fatalf("call %s missing file hash: %#v", call.Name, call)
		}
	}

	service := requireDefinition(t, ir, "Service", scopeir.NodeClass)
	requireDefinition(t, ir, "Named", scopeir.NodeInterface)
	requireDefinition(t, ir, "UserLike", scopeir.NodeTypeAlias)
	requireDefinition(t, ir, "makeRepo", scopeir.NodeFunction)
	requireDefinition(t, ir, "repo", scopeir.NodeProperty)
	requireDefinition(t, ir, "id", scopeir.NodeProperty)
	requireDefinition(t, ir, "model", scopeir.NodeVariable)
	requireDefinition(t, ir, "made", scopeir.NodeVariable)
	requireDefinition(t, ir, "formatted", scopeir.NodeVariable)
	save := requireDefinition(t, ir, "save", scopeir.NodeMethod)
	if save.OwnerID != service.ID || save.QualifiedName != "Service.save" {
		t.Fatalf("save owner/qualified mismatch: %#v service=%s", save, service.ID)
	}

	requireImport(t, ir, scopeir.ImportNamed, "User", "default", "./user")
	requireImport(t, ir, scopeir.ImportAlias, "fmt", "format", "./user")
	requireImport(t, ir, scopeir.ImportReexport, "UserService", "Service", "./service")

	requireCall(t, ir, "Model", scopeir.CallConstructor)
	requireCall(t, ir, "makeRepo", scopeir.CallFree)
	requireCall(t, ir, "fmt", scopeir.CallFree)
	requireCall(t, ir, "write", scopeir.CallMember)
	requireAccess(t, ir, "repo", scopeir.AccessWrite)
	requireAccess(t, ir, "id", scopeir.AccessRead)
	requireHeritage(t, ir, "BaseService", scopeir.HeritageExtends)
	requireHeritage(t, ir, "Named", scopeir.HeritageImplements)
	requireTypeBinding(t, ir, "this", "Service")
	requireTypeBinding(t, ir, "repo", "Repo")
	requireTypeBinding(t, ir, "user", "User")
	requireTypeBinding(t, ir, "made", "Repo")
	requireTypeAnnotation(t, ir, "Named")
	requireReturnType(t, ir, save.ID, "Promise<void>")
}

func TestExtractTypeScriptInterfaceHeritage(t *testing.T) {
	source := []byte(`interface Area { id: string; }
interface CountedArea extends Area { count: number; }
interface RankedArea extends Area, CountedArea { rank: number; }
interface GenericArea extends Pick<Area, "id"> { label: string; }
`)
	ir := parseAndExtract(t, "src/area.ts", "hash-area", scanner.TypeScript, source)

	requireDefinition(t, ir, "Area", scopeir.NodeInterface)
	requireDefinition(t, ir, "CountedArea", scopeir.NodeInterface)
	requireDefinition(t, ir, "RankedArea", scopeir.NodeInterface)
	requireHeritage(t, ir, "Area", scopeir.HeritageExtends)
	requireHeritage(t, ir, "CountedArea", scopeir.HeritageExtends)
	requireHeritage(t, ir, `Pick<Area, "id">`, scopeir.HeritageExtends)
}

type restaurantManagerHeritageCase struct {
	path   string
	source string
	want   []string
}

var restaurantManagerHeritageCases = []restaurantManagerHeritageCase{
	{
		path: "electron/renderer/src/utils/performance.ts",
		source: `interface RenderPerformanceEntry extends PerformanceEntry {}
interface BrowserPerformance extends Performance {}
`,
		want: []string{"PerformanceEntry", "Performance"},
	},
	{
		path: "electron/renderer/src/utils/dateUtils.ts",
		source: `interface DateOptions {}
interface TimeOptions {}
interface DateTimeOptions extends DateOptions, TimeOptions {}
`,
		want: []string{"DateOptions", "TimeOptions"},
	},
	{
		path: "electron/renderer/src/types/table.ts",
		source: `interface Table {}
interface TableWithUser extends Table {}
`,
		want: []string{"Table"},
	},
	{
		path: "electron/renderer/src/types/area.ts",
		source: `interface Area {}
interface AreaWithTableCount extends Area {}
`,
		want: []string{"Area"},
	},
	{
		path: "electron/renderer/src/features/tables/types.ts",
		source: `interface Table {}
interface TableWithUser extends Table {}
`,
		want: []string{"Table"},
	},
	{
		path:   "electron/renderer/src/components/shared/Form/FormTextarea.tsx",
		source: `interface FormTextareaProps extends React.TextareaHTMLAttributes<HTMLTextAreaElement> {}`,
		want:   []string{"React.TextareaHTMLAttributes<HTMLTextAreaElement>"},
	},
	{
		path:   "electron/renderer/src/components/shared/Form/FormSelect.tsx",
		source: `interface FormSelectProps extends React.SelectHTMLAttributes<HTMLSelectElement> {}`,
		want:   []string{"React.SelectHTMLAttributes<HTMLSelectElement>"},
	},
	{
		path:   "electron/renderer/src/components/shared/Form/FormInput.tsx",
		source: `interface FormInputProps extends React.InputHTMLAttributes<HTMLInputElement> {}`,
		want:   []string{"React.InputHTMLAttributes<HTMLInputElement>"},
	},
	{
		path:   "electron/renderer/src/components/shared/Form/FormCheckbox.tsx",
		source: `interface FormCheckboxProps extends React.InputHTMLAttributes<HTMLInputElement> {}`,
		want:   []string{"React.InputHTMLAttributes<HTMLInputElement>"},
	},
	{
		path:   "electron/renderer/src/components/shared/ErrorState/ErrorBoundary.tsx",
		source: `class ErrorBoundary extends Component<Props, State> {}`,
		want:   []string{"Component"},
	},
	{
		path:   "electron/renderer/src/components/shared/Button/Button.tsx",
		source: `interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {}`,
		want:   []string{"React.ButtonHTMLAttributes<HTMLButtonElement>"},
	},
	{
		path:   "electron/renderer/src/api/client.ts",
		source: `export class ApiError extends Error {}`,
		want:   []string{"Error"},
	},
	{
		path: "electron/renderer/src/features/shifts/types.ts",
		source: `interface Shift {}
interface ShiftAssignment {}
interface ShiftDTO {}
interface ShiftWithCounts extends Shift {}
interface AssignmentWithUser extends ShiftAssignment {}
interface ShiftWithCountsDTO extends ShiftDTO {}
`,
		want: []string{"Shift", "ShiftAssignment", "ShiftDTO"},
	},
}

func TestExtractRestaurantManagerTypeScriptHeritageFixture(t *testing.T) {
	total := assertRestaurantManagerHeritageCases(t, func(t *testing.T, test restaurantManagerHeritageCase) []byte {
		t.Helper()
		return []byte(test.source)
	})
	if total != 17 {
		t.Fatalf("committed Restaurant_manager TS heritage target facts = %d, want 17", total)
	}
}

func TestExtractRestaurantManagerTypeScriptHeritageSites(t *testing.T) {
	root := os.Getenv("ANVIEN_RESTAURANT_MANAGER_ROOT")
	if root == "" {
		t.Skip("set ANVIEN_RESTAURANT_MANAGER_ROOT to trace Restaurant_manager TS heritage sites")
	}

	total := assertRestaurantManagerHeritageCases(t, func(t *testing.T, test restaurantManagerHeritageCase) []byte {
		t.Helper()
		fullPath := filepath.Join(root, filepath.FromSlash(test.path))
		source, err := os.ReadFile(fullPath)
		if err != nil {
			t.Fatalf("read %s: %v", fullPath, err)
		}
		return source
	})
	if total != 17 {
		t.Fatalf("Restaurant_manager TS heritage target facts = %d, want 17", total)
	}
}

func assertRestaurantManagerHeritageCases(t *testing.T, sourceFor func(*testing.T, restaurantManagerHeritageCase) []byte) int {
	t.Helper()
	total := 0
	for _, tt := range restaurantManagerHeritageCases {
		t.Run(tt.path, func(t *testing.T) {
			source := sourceFor(t, tt)
			ir := parseAndExtract(t, tt.path, "hash-restaurant-manager", scanner.TypeScript, source)
			for _, target := range tt.want {
				requireHeritage(t, ir, target, scopeir.HeritageExtends)
			}
			if len(ir.Heritage) != len(tt.want) {
				t.Fatalf("heritage fact count for %s = %d, want %d: %#v", tt.path, len(ir.Heritage), len(tt.want), ir.Heritage)
			}
			total += len(ir.Heritage)
		})
	}
	return total
}

func TestExtractTypeScriptScopeIRParityFixture(t *testing.T) {
	ir := parseAndExtract(t, "src/service.ts", "hash-ts", scanner.TypeScript, []byte(typescriptParityFixture))
	signature := buildParitySignature(ir)
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(signature); err != nil {
		t.Fatalf("marshal signature: %v", err)
	}
	raw := buffer.Bytes()
	golden, err := os.ReadFile("testdata/typescript_scopeir_signature.golden.json")
	if err != nil {
		t.Fatalf("read golden: %v", err)
	}
	if string(raw) != string(golden) {
		t.Fatalf("parity signature mismatch\nwant:\n%s\ngot:\n%s", golden, raw)
	}
}

func TestExtractJavaScriptScopeIR(t *testing.T) {
	source := []byte(`import { createService } from './factory';

export function start() {
  const service = createService();
  service.run();
}
`)
	ir := parseAndExtract(t, "src/start.js", "hash-js", scanner.JavaScript, source)

	requireDefinition(t, ir, "start", scopeir.NodeFunction)
	requireDefinition(t, ir, "service", scopeir.NodeVariable)
	requireImport(t, ir, scopeir.ImportNamed, "createService", "createService", "./factory")
	requireCall(t, ir, "createService", scopeir.CallFree)
	requireCall(t, ir, "run", scopeir.CallMember)
}

func TestExtractVariableBindingPatternsEmitScopeIRFacts(t *testing.T) {
	source := []byte(`function bind(input: any, fallback: any) { const [first,,{source: alias = fallback, nested: [deep]}, ...tail] = input; const {direct, outer: {inner}, defaulted = fallback} = input; }`)
	ir := parseAndExtract(t, "src/variables.ts", "hash-variables", scanner.TypeScript, source)

	if len(ir.ExtractionDiagnostics) != 0 {
		t.Fatalf("variable binding diagnostics = %#v, want none", ir.ExtractionDiagnostics)
	}
	variableLeaves := make([]scopeir.BindingLeafFact, 0, 7)
	for _, leaf := range ir.BindingLeaves {
		if leaf.Provenance.Context == scopeir.BindingContextVariable {
			variableLeaves = append(variableLeaves, leaf)
		}
	}
	if len(variableLeaves) != 7 {
		t.Fatalf("variable binding leaves = %d, want 7: %#v", len(variableLeaves), variableLeaves)
	}

	functionScopeID := ""
	for _, scope := range ir.Scopes {
		if scope.Kind == scopeir.ScopeFunction {
			if functionScopeID != "" {
				t.Fatalf("multiple function scopes in focused fixture: %#v", ir.Scopes)
			}
			functionScopeID = scope.ID
		}
	}
	if functionScopeID == "" {
		t.Fatalf("missing function scope in %#v", ir.Scopes)
	}

	want := map[string]struct {
		path      string
		rangeText string
		rest      bool
		defaults  bool
	}{
		"first":     {path: "array:0", rangeText: "first"},
		"alias":     {path: "array:2/property:source", rangeText: "source: alias = fallback", defaults: true},
		"deep":      {path: "array:2/property:nested/array:0", rangeText: "nested: [deep]"},
		"tail":      {path: "array:3", rangeText: "...tail", rest: true},
		"direct":    {path: "property:direct", rangeText: "direct"},
		"inner":     {path: "property:outer/property:inner", rangeText: "outer: {inner}"},
		"defaulted": {path: "property:defaulted", rangeText: "defaulted = fallback", defaults: true},
	}
	seen := map[string]int{}
	for _, leaf := range variableLeaves {
		expected, ok := want[leaf.Name]
		if !ok {
			t.Fatalf("unexpected production binding leaf %#v", leaf)
		}
		seen[leaf.Name]++
		if got := bindingPathText(leaf.Path); got != expected.path {
			t.Fatalf("leaf %s path = %q, want %q", leaf.Name, got, expected.path)
		}
		if got := sourceTextForRange(source, leaf.Range); got != expected.rangeText {
			t.Fatalf("leaf %s range text = %q, want %q", leaf.Name, got, expected.rangeText)
		}
		if leaf.SelectionRange == nil || sourceTextForRange(source, *leaf.SelectionRange) != leaf.Name {
			t.Fatalf("leaf %s selection range = %#v", leaf.Name, leaf.SelectionRange)
		}
		if leaf.Rest != expected.rest || leaf.Default != expected.defaults {
			t.Fatalf("leaf %s modifiers = rest:%t default:%t, want rest:%t default:%t", leaf.Name, leaf.Rest, leaf.Default, expected.rest, expected.defaults)
		}
		if leaf.Provenance.Context != scopeir.BindingContextVariable {
			t.Fatalf("leaf %s context = %q, want variable", leaf.Name, leaf.Provenance.Context)
		}

		definitions := p3bDefinitionsNamed(ir, leaf.Name, scopeir.NodeVariable)
		if len(definitions) != 1 {
			t.Fatalf("leaf %s definitions = %d, want 1: %#v", leaf.Name, len(definitions), definitions)
		}
		definition := definitions[0]
		if definition.Range != leaf.Range || definition.SelectionRange == nil || *definition.SelectionRange != *leaf.SelectionRange {
			t.Fatalf("leaf %s definition ranges mismatch: leaf=%#v definition=%#v", leaf.Name, leaf, definition)
		}
		if definition.DeclaredType != "" || definition.ReturnType != "" {
			t.Fatalf("leaf %s invented type data: %#v", leaf.Name, definition)
		}

		bindingCount := 0
		ownedCount := 0
		bindingScopeID := ""
		for _, scope := range ir.Scopes {
			for _, binding := range scope.Bindings {
				if binding.Name == leaf.Name && binding.DefID == definition.ID && binding.Origin == scopeir.BindingLocal {
					bindingCount++
					bindingScopeID = scope.ID
				}
			}
			for _, defID := range scope.OwnedDefIDs {
				if defID == definition.ID {
					ownedCount++
				}
			}
		}
		if bindingCount != 1 || ownedCount != 1 || bindingScopeID != functionScopeID {
			t.Fatalf("leaf %s scope facts = bindings:%d owned:%d scope:%q, want 1/1/%q", leaf.Name, bindingCount, ownedCount, bindingScopeID, functionScopeID)
		}
	}
	for name := range want {
		if seen[name] != 1 {
			t.Fatalf("leaf %s emitted %d times, want 1", name, seen[name])
		}
	}
}

func TestExtractVariableBindingPatternSurvivesTypeInferenceMiss(t *testing.T) {
	source := []byte(`function bind(source: unknown) { const {untyped: local} = source; }`)
	ir := parseAndExtract(t, "src/inference-miss.ts", "hash-inference-miss", scanner.TypeScript, source)

	definitions := p3bDefinitionsNamed(ir, "local", scopeir.NodeVariable)
	if len(definitions) != 1 {
		t.Fatalf("local definitions = %d, want 1: %#v", len(definitions), definitions)
	}
	variableLeaves := make([]scopeir.BindingLeafFact, 0, 1)
	for _, leaf := range ir.BindingLeaves {
		if leaf.Provenance.Context == scopeir.BindingContextVariable {
			variableLeaves = append(variableLeaves, leaf)
		}
	}
	if len(variableLeaves) != 1 || variableLeaves[0].Name != "local" || bindingPathText(variableLeaves[0].Path) != "property:untyped" {
		t.Fatalf("inference-miss variable binding leaves = %#v", variableLeaves)
	}
	if definitions[0].DeclaredType != "" || definitions[0].ReturnType != "" {
		t.Fatalf("inference-miss leaf has invented type: %#v", definitions[0])
	}
	for _, scope := range ir.Scopes {
		for _, binding := range scope.TypeBindings {
			if binding.Name == "local" {
				t.Fatalf("inference-miss leaf unexpectedly has type binding: %#v", binding)
			}
		}
	}
	for _, annotation := range ir.TypeAnnotations {
		if annotation.Name == "local" {
			t.Fatalf("inference-miss leaf unexpectedly has type annotation: %#v", annotation)
		}
	}
}

func TestExtractVariableBindingPatternsPreserveSiblingBoundaries(t *testing.T) {
	baselineSource := []byte("import base, { named as alias } from './dep';\nconst existing = 1;\n")
	patternSource := append(append([]byte(nil), baselineSource...), []byte("const {bound} = source;\n")...)
	assignmentSource := append(append([]byte(nil), baselineSource...), []byte("({assigned} = source);\n[written] = source;\n")...)

	baseline := parseAndExtract(t, "src/siblings.ts", "hash-siblings", scanner.TypeScript, baselineSource)
	withPattern := parseAndExtract(t, "src/siblings.ts", "hash-siblings", scanner.TypeScript, patternSource)
	withAssignments := parseAndExtract(t, "src/siblings.ts", "hash-siblings", scanner.TypeScript, assignmentSource)

	requireSameImports(t, baseline, withPattern)
	requireSameImports(t, baseline, withAssignments)
	if len(baseline.BindingLeaves) != 0 || len(withAssignments.BindingLeaves) != 0 || len(withPattern.BindingLeaves) != 1 {
		t.Fatalf("binding leaf deltas baseline/pattern/assignment = %d/%d/%d", len(baseline.BindingLeaves), len(withPattern.BindingLeaves), len(withAssignments.BindingLeaves))
	}
	if len(withAssignments.Definitions) != len(baseline.Definitions) {
		t.Fatalf("assignment destructuring declaration delta = %d, want 0", len(withAssignments.Definitions)-len(baseline.Definitions))
	}
	if len(p3bDefinitionsNamed(withAssignments, "assigned", scopeir.NodeVariable)) != 0 || len(p3bDefinitionsNamed(withAssignments, "written", scopeir.NodeVariable)) != 0 {
		t.Fatalf("assignment destructuring emitted declarations: %#v", withAssignments.Definitions)
	}
	if len(p3bDefinitionsNamed(withPattern, "bound", scopeir.NodeVariable)) != 1 {
		t.Fatalf("pattern declaration missing or duplicated: %#v", withPattern.Definitions)
	}

	baselineIdentifier := requireDefinition(t, baseline, "existing", scopeir.NodeVariable)
	patternIdentifier := requireDefinition(t, withPattern, "existing", scopeir.NodeVariable)
	assignmentIdentifier := requireDefinition(t, withAssignments, "existing", scopeir.NodeVariable)
	if baselineIdentifier.ID != patternIdentifier.ID || baselineIdentifier.ID != assignmentIdentifier.ID ||
		baselineIdentifier.Range != patternIdentifier.Range || baselineIdentifier.Range != assignmentIdentifier.Range ||
		baselineIdentifier.SelectionRange == nil || patternIdentifier.SelectionRange == nil || assignmentIdentifier.SelectionRange == nil ||
		*baselineIdentifier.SelectionRange != *patternIdentifier.SelectionRange || *baselineIdentifier.SelectionRange != *assignmentIdentifier.SelectionRange {
		t.Fatalf("identifier regression baseline=%#v pattern=%#v assignment=%#v", baselineIdentifier, patternIdentifier, assignmentIdentifier)
	}
}

func TestExtractVariableBindingPatternSurfacesStructuredDiagnostic(t *testing.T) {
	ir := parseAndExtract(t, "src/invalid-variable.ts", "hash-invalid-variable", scanner.TypeScript, []byte(`const [...target.member] = input;`))
	if len(ir.BindingLeaves) != 0 || len(ir.ExtractionDiagnostics) != 1 {
		t.Fatalf("invalid variable pattern result = leaves:%#v diagnostics:%#v", ir.BindingLeaves, ir.ExtractionDiagnostics)
	}
	if ir.ExtractionDiagnostics[0].Code != scopeir.DiagnosticInvalidRestBinding || ir.ExtractionDiagnostics[0].Provenance.Context != scopeir.BindingContextVariable {
		t.Fatalf("invalid variable diagnostic = %#v", ir.ExtractionDiagnostics[0])
	}
	if len(p3bDefinitionsNamed(ir, "target", scopeir.NodeVariable)) != 0 || len(p3bDefinitionsNamed(ir, "member", scopeir.NodeVariable)) != 0 {
		t.Fatalf("invalid variable pattern emitted declarations: %#v", ir.Definitions)
	}
}

func TestExtractParameterBindingPatternsEmitScopeIRFacts(t *testing.T) {
	source := []byte(`function bind(fallback: any, [first,,{source: alias = fallback}, ...tail]: any[], {outer: {deep}, optional = fallback} = fallback, ...rest: any[]) {} class Service { constructor({repo: ctorRepo}: any) {} method([methodValue]: any) {} } const parenthesized = ({item: arrowItem}: any) => arrowItem; const single = singleValue => singleValue;`)
	ir := parseAndExtract(t, "src/parameters.ts", "hash-parameters", scanner.TypeScript, source)

	if len(ir.ExtractionDiagnostics) != 0 {
		t.Fatalf("parameter binding diagnostics = %#v, want none", ir.ExtractionDiagnostics)
	}

	type expectedParameterLeaf struct {
		path          string
		rangeText     string
		patternText   string
		constructText string
		scopeMarker   string
		rest          bool
		defaults      bool
	}
	want := map[string]expectedParameterLeaf{
		"fallback": {
			rangeText:     "fallback: any",
			patternText:   "fallback",
			constructText: "fallback: any",
			scopeMarker:   "function bind(",
		},
		"first": {
			path:          "array:0",
			rangeText:     "first",
			patternText:   "[first,,{source: alias = fallback}, ...tail]",
			constructText: "[first,,{source: alias = fallback}, ...tail]: any[]",
			scopeMarker:   "function bind(",
		},
		"alias": {
			path:          "array:2/property:source",
			rangeText:     "source: alias = fallback",
			patternText:   "[first,,{source: alias = fallback}, ...tail]",
			constructText: "[first,,{source: alias = fallback}, ...tail]: any[]",
			scopeMarker:   "function bind(",
			defaults:      true,
		},
		"tail": {
			path:          "array:3",
			rangeText:     "...tail",
			patternText:   "[first,,{source: alias = fallback}, ...tail]",
			constructText: "[first,,{source: alias = fallback}, ...tail]: any[]",
			scopeMarker:   "function bind(",
			rest:          true,
		},
		"deep": {
			path:          "property:outer/property:deep",
			rangeText:     "{outer: {deep}, optional = fallback} = fallback",
			patternText:   "{outer: {deep}, optional = fallback}",
			constructText: "{outer: {deep}, optional = fallback} = fallback",
			scopeMarker:   "function bind(",
			defaults:      true,
		},
		"optional": {
			path:          "property:optional",
			rangeText:     "{outer: {deep}, optional = fallback} = fallback",
			patternText:   "{outer: {deep}, optional = fallback}",
			constructText: "{outer: {deep}, optional = fallback} = fallback",
			scopeMarker:   "function bind(",
			defaults:      true,
		},
		"rest": {
			rangeText:     "...rest",
			patternText:   "...rest",
			constructText: "...rest: any[]",
			scopeMarker:   "function bind(",
			rest:          true,
		},
		"ctorRepo": {
			path:          "property:repo",
			rangeText:     "repo: ctorRepo",
			patternText:   "{repo: ctorRepo}",
			constructText: "{repo: ctorRepo}: any",
			scopeMarker:   "constructor(",
		},
		"methodValue": {
			path:          "array:0",
			rangeText:     "methodValue",
			patternText:   "[methodValue]",
			constructText: "[methodValue]: any",
			scopeMarker:   "method(",
		},
		"arrowItem": {
			path:          "property:item",
			rangeText:     "item: arrowItem",
			patternText:   "{item: arrowItem}",
			constructText: "{item: arrowItem}: any",
			scopeMarker:   "({item: arrowItem}: any) => arrowItem",
		},
		"singleValue": {
			rangeText:     "singleValue",
			patternText:   "singleValue",
			constructText: "singleValue",
			scopeMarker:   "singleValue => singleValue",
		},
	}

	parameterLeaves := make([]scopeir.BindingLeafFact, 0, len(want))
	for _, leaf := range ir.BindingLeaves {
		if leaf.Provenance.Context == scopeir.BindingContextParameter {
			parameterLeaves = append(parameterLeaves, leaf)
		}
	}
	if len(parameterLeaves) != len(want) {
		t.Fatalf("parameter binding leaves = %d, want %d: %#v", len(parameterLeaves), len(want), parameterLeaves)
	}

	scopeIDs := map[string]string{}
	for _, expected := range want {
		if _, ok := scopeIDs[expected.scopeMarker]; !ok {
			scopeIDs[expected.scopeMarker] = p3b1FunctionScopeID(t, ir, source, expected.scopeMarker)
		}
	}

	seen := map[string]int{}
	for _, leaf := range parameterLeaves {
		expected, ok := want[leaf.Name]
		if !ok {
			t.Fatalf("unexpected parameter binding leaf %#v", leaf)
		}
		seen[leaf.Name]++
		if got := bindingPathText(leaf.Path); got != expected.path {
			t.Fatalf("parameter leaf %s path = %q, want %q", leaf.Name, got, expected.path)
		}
		if got := sourceTextForRange(source, leaf.Range); got != expected.rangeText {
			t.Fatalf("parameter leaf %s range text = %q, want %q", leaf.Name, got, expected.rangeText)
		}
		if leaf.SelectionRange == nil || sourceTextForRange(source, *leaf.SelectionRange) != leaf.Name {
			t.Fatalf("parameter leaf %s selection range = %#v", leaf.Name, leaf.SelectionRange)
		}
		if leaf.Rest != expected.rest || leaf.Default != expected.defaults {
			t.Fatalf("parameter leaf %s modifiers = rest:%t default:%t, want rest:%t default:%t", leaf.Name, leaf.Rest, leaf.Default, expected.rest, expected.defaults)
		}
		if got := sourceTextForRange(source, leaf.Provenance.PatternRange); got != expected.patternText {
			t.Fatalf("parameter leaf %s pattern range = %q, want %q", leaf.Name, got, expected.patternText)
		}
		if got := sourceTextForRange(source, leaf.Provenance.ConstructRange); got != expected.constructText {
			t.Fatalf("parameter leaf %s construct range = %q, want %q", leaf.Name, got, expected.constructText)
		}

		definitions := p3bDefinitionsNamed(ir, leaf.Name, scopeir.NodeVariable)
		if len(definitions) != 1 {
			t.Fatalf("parameter leaf %s definitions = %d, want 1: %#v", leaf.Name, len(definitions), definitions)
		}
		definition := definitions[0]
		if definition.Range != leaf.Range || definition.SelectionRange == nil || *definition.SelectionRange != *leaf.SelectionRange {
			t.Fatalf("parameter leaf %s definition ranges mismatch: leaf=%#v definition=%#v", leaf.Name, leaf, definition)
		}
		if definition.DeclaredType != "" || definition.ReturnType != "" {
			t.Fatalf("parameter leaf %s invented type data: %#v", leaf.Name, definition)
		}

		bindingCount := 0
		ownedCount := 0
		bindingScopeID := ""
		for _, scope := range ir.Scopes {
			for _, binding := range scope.Bindings {
				if binding.Name == leaf.Name && binding.DefID == definition.ID && binding.Origin == scopeir.BindingLocal {
					bindingCount++
					bindingScopeID = scope.ID
				}
			}
			for _, defID := range scope.OwnedDefIDs {
				if defID == definition.ID {
					ownedCount++
				}
			}
		}
		wantScopeID := scopeIDs[expected.scopeMarker]
		if bindingCount != 1 || ownedCount != 1 || bindingScopeID != wantScopeID {
			t.Fatalf("parameter leaf %s scope facts = bindings:%d owned:%d scope:%q, want 1/1/%q", leaf.Name, bindingCount, ownedCount, bindingScopeID, wantScopeID)
		}
	}
	for name := range want {
		if seen[name] != 1 {
			t.Fatalf("parameter leaf %s emitted %d times, want 1", name, seen[name])
		}
	}
	requireTypeBinding(t, ir, "fallback", "any")
}

func TestExtractParameterBindingPatternsExcludeTypeOnlyParameterSyntax(t *testing.T) {
	source := []byte(`function runtime(arg: {
		annotation: [annotationRequired: string, annotationOptional?: number, ...annotationRest: boolean[]];
		nestedFunction: (nestedRequired: string, nestedOptional?: number, ...nestedRest: boolean[]) => void;
		nestedConstructor: new (constructorRequired: string, constructorOptional?: number, ...constructorRest: boolean[]) => object;
	}): [returnRequired: string, returnOptional?: number, ...returnRest: boolean[]] {
		const bodyLocal: [bodyRequired: string, bodyOptional?: number, ...bodyRest: boolean[]] = ["", 1, true];
		return ["", 1, true];
	}`)
	ir := parseAndExtract(t, "src/parameter-type-syntax.ts", "hash-parameter-type-syntax", scanner.TypeScript, source)

	if len(ir.ExtractionDiagnostics) != 0 {
		t.Fatalf("type-only parameter syntax diagnostics = %#v, want none", ir.ExtractionDiagnostics)
	}

	var argLeaves []scopeir.BindingLeafFact
	for _, leaf := range ir.BindingLeaves {
		if leaf.Name == "arg" && leaf.Provenance.Context == scopeir.BindingContextParameter {
			argLeaves = append(argLeaves, leaf)
		}
	}
	if len(argLeaves) != 1 {
		t.Fatalf("runtime arg parameter leaves = %d, want 1: %#v", len(argLeaves), argLeaves)
	}

	argDefinitions := p3bDefinitionsNamed(ir, "arg", scopeir.NodeVariable)
	if len(argDefinitions) != 1 {
		t.Fatalf("runtime arg variable definitions = %d, want 1: %#v", len(argDefinitions), argDefinitions)
	}
	argDefinition := argDefinitions[0]
	if argDefinition.Range != argLeaves[0].Range || argDefinition.SelectionRange == nil || argLeaves[0].SelectionRange == nil || *argDefinition.SelectionRange != *argLeaves[0].SelectionRange {
		t.Fatalf("runtime arg definition ranges mismatch: leaf=%#v definition=%#v", argLeaves[0], argDefinition)
	}
	runtimeScopeID := ""
	for _, scope := range ir.Scopes {
		if scope.Kind != scopeir.ScopeFunction {
			continue
		}
		if runtimeScopeID != "" {
			t.Fatalf("type-only callable syntax emitted an extra runtime function scope: %#v", ir.Scopes)
		}
		runtimeScopeID = scope.ID
	}
	if runtimeScopeID == "" {
		t.Fatalf("runtime function scope missing: %#v", ir.Scopes)
	}

	argOwned := 0
	argBindings := 0
	for _, scope := range ir.Scopes {
		for _, defID := range scope.OwnedDefIDs {
			if defID == argDefinition.ID {
				argOwned++
				if scope.ID != runtimeScopeID {
					t.Fatalf("runtime arg definition owned by %q, want %q", scope.ID, runtimeScopeID)
				}
			}
		}
		for _, binding := range scope.Bindings {
			if binding.Name == "arg" && binding.DefID == argDefinition.ID && binding.Origin == scopeir.BindingLocal {
				argBindings++
				if scope.ID != runtimeScopeID {
					t.Fatalf("runtime arg local binding emitted in %q, want %q", scope.ID, runtimeScopeID)
				}
			}
		}
	}
	if argOwned != 1 || argBindings != 1 {
		t.Fatalf("runtime arg scope facts = owned:%d bindings:%d, want 1/1", argOwned, argBindings)
	}

	typeOnlyNames := []string{
		"annotationRequired", "annotationOptional", "annotationRest",
		"nestedRequired", "nestedOptional", "nestedRest",
		"constructorRequired", "constructorOptional", "constructorRest",
		"returnRequired", "returnOptional", "returnRest",
		"bodyRequired", "bodyOptional", "bodyRest",
	}
	definitionByID := make(map[string]scopeir.DefinitionFact, len(ir.Definitions))
	for _, definition := range ir.Definitions {
		definitionByID[definition.ID] = definition
	}
	for _, name := range typeOnlyNames {
		leaves := 0
		for _, leaf := range ir.BindingLeaves {
			if leaf.Name == name {
				leaves++
			}
		}
		definitions := p3bDefinitionsNamed(ir, name, scopeir.NodeVariable)
		owned := 0
		bindings := 0
		for _, scope := range ir.Scopes {
			for _, defID := range scope.OwnedDefIDs {
				if definition, ok := definitionByID[defID]; ok && definition.Name == name {
					owned++
				}
			}
			for _, binding := range scope.Bindings {
				if binding.Name == name && binding.Origin == scopeir.BindingLocal {
					bindings++
				}
			}
		}
		if leaves != 0 || len(definitions) != 0 || owned != 0 || bindings != 0 {
			t.Fatalf("type-only label %q emitted facts = leaves:%d variable-definitions:%d owned:%d local-bindings:%d", name, leaves, len(definitions), owned, bindings)
		}
	}

	bodyDefinitions := p3bDefinitionsNamed(ir, "bodyLocal", scopeir.NodeVariable)
	if len(bodyDefinitions) != 1 {
		t.Fatalf("body-local variable traversal not preserved: definitions=%#v", bodyDefinitions)
	}
}

func TestExtractParameterBindingPatternsOptionalThisAndJavaScriptControls(t *testing.T) {
	typescriptSource := []byte(`function controlled(this: { marker: string }, name?: string) { return name; }`)
	typescriptIR := parseAndExtract(t, "src/parameter-controls.ts", "hash-parameter-controls", scanner.TypeScript, typescriptSource)
	if len(typescriptIR.ExtractionDiagnostics) != 0 {
		t.Fatalf("parameter control diagnostics = %#v, want none", typescriptIR.ExtractionDiagnostics)
	}

	controlledScopeID := p3b1FunctionScopeID(t, typescriptIR, typescriptSource, "function controlled(")
	var optionalLeaves []scopeir.BindingLeafFact
	for _, leaf := range typescriptIR.BindingLeaves {
		if leaf.Provenance.Context != scopeir.BindingContextParameter {
			continue
		}
		switch leaf.Name {
		case "name":
			optionalLeaves = append(optionalLeaves, leaf)
		case "this":
			t.Fatalf("explicit this pseudo-parameter emitted a binding leaf: %#v", leaf)
		default:
			t.Fatalf("unexpected parameter control leaf: %#v", leaf)
		}
	}
	if len(optionalLeaves) != 1 {
		t.Fatalf("optional name parameter leaves = %d, want 1: %#v", len(optionalLeaves), optionalLeaves)
	}
	optionalLeaf := optionalLeaves[0]
	if got := sourceTextForRange(typescriptSource, optionalLeaf.Range); got != "name?: string" {
		t.Fatalf("optional name range = %q, want %q", got, "name?: string")
	}
	if got := sourceTextForRange(typescriptSource, optionalLeaf.Provenance.PatternRange); got != "name" {
		t.Fatalf("optional name pattern = %q, want %q", got, "name")
	}
	if got := sourceTextForRange(typescriptSource, optionalLeaf.Provenance.ConstructRange); got != "name?: string" {
		t.Fatalf("optional name construct = %q, want %q", got, "name?: string")
	}

	optionalDefinitions := p3bDefinitionsNamed(typescriptIR, "name", scopeir.NodeVariable)
	if len(optionalDefinitions) != 1 {
		t.Fatalf("optional name definitions = %d, want 1: %#v", len(optionalDefinitions), optionalDefinitions)
	}
	optionalDefinition := optionalDefinitions[0]
	optionalBindings := 0
	optionalOwned := 0
	for _, scope := range typescriptIR.Scopes {
		for _, binding := range scope.Bindings {
			if binding.Name == "this" && binding.Origin == scopeir.BindingLocal {
				t.Fatalf("explicit this pseudo-parameter emitted a local binding in %q: %#v", scope.ID, binding)
			}
			if scope.ID == controlledScopeID && binding.Name == "name" && binding.DefID == optionalDefinition.ID && binding.Origin == scopeir.BindingLocal {
				optionalBindings++
			}
		}
		for _, defID := range scope.OwnedDefIDs {
			if scope.ID == controlledScopeID && defID == optionalDefinition.ID {
				optionalOwned++
			}
		}
	}
	if optionalBindings != 1 || optionalOwned != 1 {
		t.Fatalf("optional name scope facts = bindings:%d owned:%d, want 1/1", optionalBindings, optionalOwned)
	}
	if definitions := p3bDefinitionsNamed(typescriptIR, "this", scopeir.NodeVariable); len(definitions) != 0 {
		t.Fatalf("explicit this pseudo-parameter definitions = %d, want 0: %#v", len(definitions), definitions)
	}

	javascriptSource := []byte(`const unchanged = plain => plain;`)
	javascriptIR := parseAndExtract(t, "src/parameter-controls.js", "hash-parameter-controls-js", scanner.JavaScript, javascriptSource)
	for _, leaf := range javascriptIR.BindingLeaves {
		if leaf.Provenance.Context == scopeir.BindingContextParameter {
			t.Fatalf("plain JavaScript arrow emitted a parameter binding leaf: %#v", leaf)
		}
	}
	if definitions := p3bDefinitionsNamed(javascriptIR, "plain", scopeir.NodeVariable); len(definitions) != 0 {
		t.Fatalf("plain JavaScript arrow parameter definitions = %d, want 0: %#v", len(definitions), definitions)
	}
	for _, scope := range javascriptIR.Scopes {
		for _, binding := range scope.Bindings {
			if binding.Name == "plain" && binding.Origin == scopeir.BindingLocal {
				t.Fatalf("plain JavaScript arrow parameter emitted a local binding in %q: %#v", scope.ID, binding)
			}
		}
	}
}

func TestExtractParameterBindingPatternsPreserveShadowingAndSiblingContexts(t *testing.T) {
	source := []byte(`import { consume } from './dep'; function outer({value}: any) { const {source: localValue} = input; function inner({value}: any) { const existing = value; consume(value, localValue, existing); } } try {} catch (caught) { consume(caught); } for (const loopValue of values) { consume(loopValue); } ({written} = input);`)
	ir := parseAndExtract(t, "src/parameter-shadowing.ts", "hash-parameter-shadowing", scanner.TypeScript, source)

	outerScopeID := p3b1FunctionScopeID(t, ir, source, "function outer(")
	innerScopeID := p3b1FunctionScopeID(t, ir, source, "function inner(")
	if outerScopeID == innerScopeID {
		t.Fatalf("outer/inner parameter scopes both = %q", outerScopeID)
	}

	valueDefinitions := p3bDefinitionsNamed(ir, "value", scopeir.NodeVariable)
	if len(valueDefinitions) != 2 {
		t.Fatalf("shadowed parameter definitions = %d, want 2: %#v", len(valueDefinitions), valueDefinitions)
	}
	valueBindingsByScope := map[string]int{}
	valueOwnedByScope := map[string]int{}
	for _, definition := range valueDefinitions {
		for _, scope := range ir.Scopes {
			for _, binding := range scope.Bindings {
				if binding.Name == "value" && binding.DefID == definition.ID && binding.Origin == scopeir.BindingLocal {
					valueBindingsByScope[scope.ID]++
				}
			}
			for _, defID := range scope.OwnedDefIDs {
				if defID == definition.ID {
					valueOwnedByScope[scope.ID]++
				}
			}
		}
	}
	for _, scopeID := range []string{outerScopeID, innerScopeID} {
		if valueBindingsByScope[scopeID] != 1 || valueOwnedByScope[scopeID] != 1 {
			t.Fatalf("shadowed value facts in %q = bindings:%d owned:%d, want 1/1", scopeID, valueBindingsByScope[scopeID], valueOwnedByScope[scopeID])
		}
	}
	if len(valueBindingsByScope) != 2 || len(valueOwnedByScope) != 2 {
		t.Fatalf("shadowed value leaked across scopes: bindings=%#v owned=%#v", valueBindingsByScope, valueOwnedByScope)
	}

	localValue := p3bDefinitionsNamed(ir, "localValue", scopeir.NodeVariable)
	if len(localValue) != 1 || !p3b1ScopeOwnsAndBinds(ir, outerScopeID, localValue[0].ID, "localValue") {
		t.Fatalf("variable sibling localValue not preserved in outer scope: definitions=%#v scopes=%#v", localValue, ir.Scopes)
	}
	existing := p3bDefinitionsNamed(ir, "existing", scopeir.NodeVariable)
	if len(existing) != 1 || !p3b1ScopeOwnsAndBinds(ir, innerScopeID, existing[0].ID, "existing") {
		t.Fatalf("identifier variable sibling existing not preserved in inner scope: definitions=%#v scopes=%#v", existing, ir.Scopes)
	}

	parameterLeaves := 0
	variableLeaves := 0
	catchLeaves := 0
	for _, leaf := range ir.BindingLeaves {
		switch leaf.Provenance.Context {
		case scopeir.BindingContextParameter:
			parameterLeaves++
		case scopeir.BindingContextVariable:
			variableLeaves++
		case scopeir.BindingContextCatch:
			catchLeaves++
		case scopeir.BindingContextForIn, scopeir.BindingContextForOf:
			t.Fatalf("locked sibling context unexpectedly emitted leaf %#v", leaf)
		}
	}
	if parameterLeaves != 2 || variableLeaves != 1 || catchLeaves != 1 {
		t.Fatalf("parameter/variable/catch leaf counts = %d/%d/%d, want 2/1/1: %#v", parameterLeaves, variableLeaves, catchLeaves, ir.BindingLeaves)
	}
	caughtDefinitions := p3bDefinitionsNamed(ir, "caught", scopeir.NodeVariable)
	if len(caughtDefinitions) != 1 {
		t.Fatalf("catch sibling definitions = %d, want 1: %#v", len(caughtDefinitions), caughtDefinitions)
	}
	catchScope := p3b2CatchScope(t, ir, source, "catch (caught) { consume(caught); }")
	caughtOwned, caughtBindings := p3b2ScopeFactCounts(ir, catchScope.ID, caughtDefinitions[0].ID, "caught")
	if caughtOwned != 1 || caughtBindings != 1 {
		t.Fatalf("catch sibling scope facts = owned:%d bindings:%d, want 1/1", caughtOwned, caughtBindings)
	}
	if len(p3bDefinitionsNamed(ir, "written", scopeir.NodeVariable)) != 0 {
		t.Fatalf("assignment sibling emitted a false declaration: %#v", ir.Definitions)
	}
	requireImport(t, ir, scopeir.ImportNamed, "consume", "consume", "./dep")
	if got := countCalls(ir, "consume", scopeir.CallFree); got != 3 {
		t.Fatalf("consume calls = %d, want 3", got)
	}

	patternTypeBindings := 0
	for _, scope := range ir.Scopes {
		for _, binding := range scope.TypeBindings {
			if binding.Name == "{value}" && binding.Type.RawName == "any" && binding.Type.Source == scopeir.TypeSourceParameter {
				patternTypeBindings++
			}
		}
	}
	if patternTypeBindings != 2 {
		t.Fatalf("separate parameter type-binding path count = %d, want 2", patternTypeBindings)
	}
}

func TestExtractCatchBindingPatternsEmitScopeIRFacts(t *testing.T) {
	source := []byte(`function run(input: any, fallback: any) { try { throw input; } catch (caught) { consume(caught); } try { throw input; } catch ([first,,{source: alias = fallback}, ...tail]) { consume(first, alias, tail); } try { throw input; } catch ({outer: {deep}, optional = fallback, ...rest}) { consume(deep, optional, rest); } }`)
	ir := parseAndExtract(t, "src/catch-patterns.ts", "hash-catch-patterns", scanner.TypeScript, source)

	if len(ir.ExtractionDiagnostics) != 0 {
		t.Fatalf("catch binding diagnostics = %#v, want none", ir.ExtractionDiagnostics)
	}

	type expectedCatchLeaf struct {
		path          string
		rangeText     string
		patternText   string
		patternKind   string
		constructText string
		rest          bool
		defaults      bool
	}
	want := map[string]expectedCatchLeaf{
		"caught": {
			rangeText:     "caught",
			patternText:   "caught",
			patternKind:   "identifier",
			constructText: "catch (caught) { consume(caught); }",
		},
		"first": {
			path:          "array:0",
			rangeText:     "first",
			patternText:   "[first,,{source: alias = fallback}, ...tail]",
			patternKind:   "array_pattern",
			constructText: "catch ([first,,{source: alias = fallback}, ...tail]) { consume(first, alias, tail); }",
		},
		"alias": {
			path:          "array:2/property:source",
			rangeText:     "source: alias = fallback",
			patternText:   "[first,,{source: alias = fallback}, ...tail]",
			patternKind:   "array_pattern",
			constructText: "catch ([first,,{source: alias = fallback}, ...tail]) { consume(first, alias, tail); }",
			defaults:      true,
		},
		"tail": {
			path:          "array:3",
			rangeText:     "...tail",
			patternText:   "[first,,{source: alias = fallback}, ...tail]",
			patternKind:   "array_pattern",
			constructText: "catch ([first,,{source: alias = fallback}, ...tail]) { consume(first, alias, tail); }",
			rest:          true,
		},
		"deep": {
			path:          "property:outer/property:deep",
			rangeText:     "outer: {deep}",
			patternText:   "{outer: {deep}, optional = fallback, ...rest}",
			patternKind:   "object_pattern",
			constructText: "catch ({outer: {deep}, optional = fallback, ...rest}) { consume(deep, optional, rest); }",
		},
		"optional": {
			path:          "property:optional",
			rangeText:     "optional = fallback",
			patternText:   "{outer: {deep}, optional = fallback, ...rest}",
			patternKind:   "object_pattern",
			constructText: "catch ({outer: {deep}, optional = fallback, ...rest}) { consume(deep, optional, rest); }",
			defaults:      true,
		},
		"rest": {
			rangeText:     "...rest",
			patternText:   "{outer: {deep}, optional = fallback, ...rest}",
			patternKind:   "object_pattern",
			constructText: "catch ({outer: {deep}, optional = fallback, ...rest}) { consume(deep, optional, rest); }",
			rest:          true,
		},
	}

	catchLeaves := make([]scopeir.BindingLeafFact, 0, len(want))
	for _, leaf := range ir.BindingLeaves {
		if leaf.Provenance.Context == scopeir.BindingContextCatch {
			catchLeaves = append(catchLeaves, leaf)
		}
	}
	if len(catchLeaves) != len(want) {
		t.Fatalf("catch binding leaves = %d, want %d: %#v", len(catchLeaves), len(want), catchLeaves)
	}

	functionScopeID := p3b1FunctionScopeID(t, ir, source, "function run(")
	catchScopes := make(map[string]scopeir.ScopeFact, 3)
	for _, expected := range want {
		if _, ok := catchScopes[expected.constructText]; ok {
			continue
		}
		scope := p3b2CatchScope(t, ir, source, expected.constructText)
		if scope.Parent == nil || *scope.Parent != functionScopeID {
			t.Fatalf("catch scope %q parent = %#v, want %q", expected.constructText, scope.Parent, functionScopeID)
		}
		catchScopes[expected.constructText] = scope
	}

	seen := map[string]int{}
	for _, leaf := range catchLeaves {
		expected, ok := want[leaf.Name]
		if !ok {
			t.Fatalf("unexpected catch binding leaf %#v", leaf)
		}
		seen[leaf.Name]++
		if got := bindingPathText(leaf.Path); got != expected.path {
			t.Fatalf("catch leaf %s path = %q, want %q", leaf.Name, got, expected.path)
		}
		if got := sourceTextForRange(source, leaf.Range); got != expected.rangeText {
			t.Fatalf("catch leaf %s range text = %q, want %q", leaf.Name, got, expected.rangeText)
		}
		if leaf.SelectionRange == nil || sourceTextForRange(source, *leaf.SelectionRange) != leaf.Name {
			t.Fatalf("catch leaf %s selection range = %#v", leaf.Name, leaf.SelectionRange)
		}
		if leaf.Rest != expected.rest || leaf.Default != expected.defaults {
			t.Fatalf("catch leaf %s modifiers = rest:%t default:%t, want rest:%t default:%t", leaf.Name, leaf.Rest, leaf.Default, expected.rest, expected.defaults)
		}
		if leaf.Provenance.PatternKind != expected.patternKind ||
			sourceTextForRange(source, leaf.Provenance.PatternRange) != expected.patternText ||
			sourceTextForRange(source, leaf.Provenance.ConstructRange) != expected.constructText {
			t.Fatalf("catch leaf %s provenance mismatch: %#v", leaf.Name, leaf.Provenance)
		}

		definitions := p3bDefinitionsNamed(ir, leaf.Name, scopeir.NodeVariable)
		if len(definitions) != 1 {
			t.Fatalf("catch leaf %s definitions = %d, want 1: %#v", leaf.Name, len(definitions), definitions)
		}
		definition := definitions[0]
		if definition.Range != leaf.Range || definition.SelectionRange == nil || *definition.SelectionRange != *leaf.SelectionRange {
			t.Fatalf("catch leaf %s definition ranges mismatch: leaf=%#v definition=%#v", leaf.Name, leaf, definition)
		}
		if definition.DeclaredType != "" || definition.ReturnType != "" {
			t.Fatalf("catch leaf %s invented type data: %#v", leaf.Name, definition)
		}

		scope := catchScopes[expected.constructText]
		owned, bindings := p3b2ScopeFactCounts(ir, scope.ID, definition.ID, leaf.Name)
		globalOwned, globalBindings := p3b2GlobalScopeFactCounts(ir, definition.ID, leaf.Name)
		if owned != 1 || bindings != 1 || globalOwned != 1 || globalBindings != 1 {
			t.Fatalf("catch leaf %s scope facts = catch:%d/%d global:%d/%d, want 1/1 and 1/1 in %q", leaf.Name, owned, bindings, globalOwned, globalBindings, scope.ID)
		}
	}
	for name := range want {
		if seen[name] != 1 {
			t.Fatalf("catch leaf %s emitted %d times, want 1", name, seen[name])
		}
	}

	consumeCallsByScope := map[string]int{}
	for _, call := range ir.Calls {
		if call.Name == "consume" && call.CallForm == scopeir.CallFree {
			consumeCallsByScope[call.InScope]++
		}
	}
	for _, scope := range catchScopes {
		if consumeCallsByScope[scope.ID] != 1 {
			t.Fatalf("consume calls in catch scope %q = %d, want 1: %#v", scope.ID, consumeCallsByScope[scope.ID], ir.Calls)
		}
	}

	for name := range want {
		for _, scope := range ir.Scopes {
			for _, binding := range scope.TypeBindings {
				if binding.Name == name {
					t.Fatalf("catch leaf %s unexpectedly has a type binding: %#v", name, binding)
				}
			}
		}
		for _, annotation := range ir.TypeAnnotations {
			if annotation.Name == name {
				t.Fatalf("catch leaf %s unexpectedly has a type annotation: %#v", name, annotation)
			}
		}
	}
}

func TestExtractCatchBindingPatternsOptionalAndJavaScriptControls(t *testing.T) {
	optionalSource := []byte(`try { work(); } catch { consume(); }`)
	optionalIR := parseAndExtract(t, "src/optional-catch.ts", "hash-optional-catch", scanner.TypeScript, optionalSource)
	if len(optionalIR.BindingLeaves) != 0 || len(optionalIR.ExtractionDiagnostics) != 0 || len(optionalIR.Definitions) != 0 {
		t.Fatalf("optional catch facts = leaves:%#v diagnostics:%#v definitions:%#v, want all zero", optionalIR.BindingLeaves, optionalIR.ExtractionDiagnostics, optionalIR.Definitions)
	}
	optionalScope := p3b2CatchScope(t, optionalIR, optionalSource, "catch { consume(); }")
	if optionalScope.Parent == nil || *optionalScope.Parent != optionalIR.ModuleScope {
		t.Fatalf("optional catch scope parent = %#v, want module %q", optionalScope.Parent, optionalIR.ModuleScope)
	}
	if len(optionalScope.OwnedDefIDs) != 0 || len(optionalScope.Bindings) != 0 {
		t.Fatalf("optional catch scope owns binding facts: %#v", optionalScope)
	}
	optionalConsumeCalls := 0
	for _, call := range optionalIR.Calls {
		if call.Name == "consume" && call.InScope == optionalScope.ID {
			optionalConsumeCalls++
		}
	}
	if optionalConsumeCalls != 1 {
		t.Fatalf("optional catch consume calls in %q = %d, want 1: %#v", optionalScope.ID, optionalConsumeCalls, optionalIR.Calls)
	}

	javascriptSource := []byte(`try {} catch (error) { consume(error); } try {} catch ({message: local = fallback, ...rest}) { consume(local, rest); }`)
	javascriptIR := parseAndExtract(t, "src/catch-controls.js", "hash-catch-controls-js", scanner.JavaScript, javascriptSource)
	if len(javascriptIR.ExtractionDiagnostics) != 0 {
		t.Fatalf("JavaScript catch diagnostics = %#v, want none", javascriptIR.ExtractionDiagnostics)
	}
	want := map[string]struct {
		path          string
		rangeText     string
		patternText   string
		patternKind   string
		constructText string
		rest          bool
		defaults      bool
	}{
		"error": {
			rangeText:     "error",
			patternText:   "error",
			patternKind:   "identifier",
			constructText: "catch (error) { consume(error); }",
		},
		"local": {
			path:          "property:message",
			rangeText:     "message: local = fallback",
			patternText:   "{message: local = fallback, ...rest}",
			patternKind:   "object_pattern",
			constructText: "catch ({message: local = fallback, ...rest}) { consume(local, rest); }",
			defaults:      true,
		},
		"rest": {
			rangeText:     "...rest",
			patternText:   "{message: local = fallback, ...rest}",
			patternKind:   "object_pattern",
			constructText: "catch ({message: local = fallback, ...rest}) { consume(local, rest); }",
			rest:          true,
		},
	}
	seen := map[string]int{}
	for _, leaf := range javascriptIR.BindingLeaves {
		if leaf.Provenance.Context != scopeir.BindingContextCatch {
			t.Fatalf("JavaScript control emitted a non-catch leaf: %#v", leaf)
		}
		expected, ok := want[leaf.Name]
		if !ok {
			t.Fatalf("unexpected JavaScript catch leaf: %#v", leaf)
		}
		seen[leaf.Name]++
		if bindingPathText(leaf.Path) != expected.path || sourceTextForRange(javascriptSource, leaf.Range) != expected.rangeText || leaf.Rest != expected.rest || leaf.Default != expected.defaults {
			t.Fatalf("JavaScript catch leaf %s mismatch: %#v", leaf.Name, leaf)
		}
		if leaf.SelectionRange == nil || sourceTextForRange(javascriptSource, *leaf.SelectionRange) != leaf.Name {
			t.Fatalf("JavaScript catch leaf %s selection range = %#v", leaf.Name, leaf.SelectionRange)
		}
		if leaf.Provenance.PatternKind != expected.patternKind ||
			sourceTextForRange(javascriptSource, leaf.Provenance.PatternRange) != expected.patternText ||
			sourceTextForRange(javascriptSource, leaf.Provenance.ConstructRange) != expected.constructText {
			t.Fatalf("JavaScript catch leaf %s provenance mismatch: %#v", leaf.Name, leaf.Provenance)
		}
		definitions := p3bDefinitionsNamed(javascriptIR, leaf.Name, scopeir.NodeVariable)
		if len(definitions) != 1 {
			t.Fatalf("JavaScript catch leaf %s definitions = %d, want 1: %#v", leaf.Name, len(definitions), definitions)
		}
		definition := definitions[0]
		if definition.Range != leaf.Range || definition.SelectionRange == nil || *definition.SelectionRange != *leaf.SelectionRange {
			t.Fatalf("JavaScript catch leaf %s definition ranges mismatch: leaf=%#v definition=%#v", leaf.Name, leaf, definition)
		}
		if definition.DeclaredType != "" || definition.ReturnType != "" {
			t.Fatalf("JavaScript catch leaf %s invented type data: %#v", leaf.Name, definition)
		}
		scope := p3b2CatchScope(t, javascriptIR, javascriptSource, expected.constructText)
		if scope.Parent == nil || *scope.Parent != javascriptIR.ModuleScope {
			t.Fatalf("JavaScript catch leaf %s scope parent = %#v, want module %q", leaf.Name, scope.Parent, javascriptIR.ModuleScope)
		}
		owned, bindings := p3b2ScopeFactCounts(javascriptIR, scope.ID, definition.ID, leaf.Name)
		globalOwned, globalBindings := p3b2GlobalScopeFactCounts(javascriptIR, definition.ID, leaf.Name)
		if owned != 1 || bindings != 1 || globalOwned != 1 || globalBindings != 1 {
			t.Fatalf("JavaScript catch leaf %s scope facts = catch:%d/%d global:%d/%d, want 1/1 and 1/1", leaf.Name, owned, bindings, globalOwned, globalBindings)
		}
	}
	if len(seen) != len(want) {
		t.Fatalf("JavaScript catch leaves = %#v, want names %#v", seen, want)
	}
	for name := range want {
		if seen[name] != 1 {
			t.Fatalf("JavaScript catch leaf %s emitted %d times, want 1", name, seen[name])
		}
	}
}

func TestExtractCatchBindingPatternsPreserveShadowingAndSiblingContexts(t *testing.T) {
	source := []byte(`import { consume } from './dep'; function shadow(input: any) { const caught = input; try {} catch (caught) { consume(caught); } consume(caught); ({written} = input); }`)
	ir := parseAndExtract(t, "src/catch-shadowing.ts", "hash-catch-shadowing", scanner.TypeScript, source)

	requireImport(t, ir, scopeir.ImportNamed, "consume", "consume", "./dep")
	if len(ir.ExtractionDiagnostics) != 0 {
		t.Fatalf("catch shadowing diagnostics = %#v, want none", ir.ExtractionDiagnostics)
	}

	functionScopeID := p3b1FunctionScopeID(t, ir, source, "function shadow(")
	catchScope := p3b2CatchScope(t, ir, source, "catch (caught) { consume(caught); }")
	if catchScope.Parent == nil || *catchScope.Parent != functionScopeID {
		t.Fatalf("shadowing catch scope parent = %#v, want %q", catchScope.Parent, functionScopeID)
	}

	caughtDefinitions := p3bDefinitionsNamed(ir, "caught", scopeir.NodeVariable)
	if len(caughtDefinitions) != 2 {
		t.Fatalf("shadowed caught definitions = %d, want 2: %#v", len(caughtDefinitions), caughtDefinitions)
	}
	var outerDefinition scopeir.DefinitionFact
	var catchDefinition scopeir.DefinitionFact
	for _, definition := range caughtDefinitions {
		globalOwned, globalBindings := p3b2GlobalScopeFactCounts(ir, definition.ID, "caught")
		if globalOwned != 1 || globalBindings != 1 {
			t.Fatalf("shadowed definition %q global scope facts = %d/%d, want 1/1", definition.ID, globalOwned, globalBindings)
		}
		catchOwned, catchBindings := p3b2ScopeFactCounts(ir, catchScope.ID, definition.ID, "caught")
		outerOwned, outerBindings := p3b2ScopeFactCounts(ir, functionScopeID, definition.ID, "caught")
		switch {
		case catchOwned == 1 && catchBindings == 1 && outerOwned == 0 && outerBindings == 0:
			catchDefinition = definition
		case outerOwned == 1 && outerBindings == 1 && catchOwned == 0 && catchBindings == 0:
			outerDefinition = definition
		default:
			t.Fatalf("shadowed definition %q leaked scope facts: catch=%d/%d outer=%d/%d", definition.ID, catchOwned, catchBindings, outerOwned, outerBindings)
		}
	}
	if outerDefinition.ID == "" || catchDefinition.ID == "" || outerDefinition.ID == catchDefinition.ID {
		t.Fatalf("shadowed definitions are not distinct: outer=%#v catch=%#v", outerDefinition, catchDefinition)
	}
	if got := sourceTextForRange(source, outerDefinition.Range); got != "caught = input" {
		t.Fatalf("outer caught range = %q, want %q", got, "caught = input")
	}
	if got := sourceTextForRange(source, catchDefinition.Range); got != "caught" {
		t.Fatalf("catch caught range = %q, want %q", got, "caught")
	}

	catchLeaves := 0
	for _, leaf := range ir.BindingLeaves {
		if leaf.Name == "caught" && leaf.Provenance.Context == scopeir.BindingContextCatch {
			catchLeaves++
		}
		if leaf.Name == "written" {
			t.Fatalf("assignment sibling emitted a binding leaf: %#v", leaf)
		}
	}
	if catchLeaves != 1 {
		t.Fatalf("shadowed caught catch leaves = %d, want 1: %#v", catchLeaves, ir.BindingLeaves)
	}
	if len(p3bDefinitionsNamed(ir, "written", scopeir.NodeVariable)) != 0 {
		t.Fatalf("assignment sibling emitted a false declaration: %#v", ir.Definitions)
	}

	consumeCallsByScope := map[string]int{}
	for _, call := range ir.Calls {
		if call.Name == "consume" && call.CallForm == scopeir.CallFree {
			consumeCallsByScope[call.InScope]++
		}
	}
	if consumeCallsByScope[catchScope.ID] != 1 || consumeCallsByScope[functionScopeID] != 1 || len(consumeCallsByScope) != 2 {
		t.Fatalf("shadowed consume call scopes = %#v, want catch/function 1/1", consumeCallsByScope)
	}
}

func TestExtractBindingPatternEnumeratesTypedLeaves(t *testing.T) {
	source := []byte(`const [first,,{source: alias = fallback, short = fallback2, [dynamicKey]: computed}, ...tail] = input;`)
	result := extractPatternFromVariableDeclarator(t, source)

	if len(result.Diagnostics) != 0 {
		t.Fatalf("binding diagnostics = %#v, want none", result.Diagnostics)
	}
	if len(result.Leaves) != 5 {
		t.Fatalf("binding leaves = %d, want 5: %#v", len(result.Leaves), result.Leaves)
	}

	want := map[string]struct {
		path      string
		rangeText string
		rest      bool
		defaults  bool
	}{
		"first":    {path: "array:0", rangeText: "first"},
		"alias":    {path: "array:2/property:source", rangeText: "source: alias = fallback", defaults: true},
		"short":    {path: "array:2/property:short", rangeText: "short = fallback2", defaults: true},
		"computed": {path: "array:2/computed:dynamicKey", rangeText: "[dynamicKey]: computed"},
		"tail":     {path: "array:3", rangeText: "...tail", rest: true},
	}
	for _, leaf := range result.Leaves {
		expected, ok := want[leaf.Name]
		if !ok {
			t.Fatalf("unexpected binding leaf %#v", leaf)
		}
		if got := bindingPathText(leaf.Path); got != expected.path {
			t.Fatalf("leaf %s path = %q, want %q", leaf.Name, got, expected.path)
		}
		if got := sourceTextForRange(source, leaf.Range); got != expected.rangeText {
			t.Fatalf("leaf %s range text = %q, want %q", leaf.Name, got, expected.rangeText)
		}
		if leaf.SelectionRange == nil {
			t.Fatalf("leaf %s has nil selection range", leaf.Name)
		}
		if got := sourceTextForRange(source, *leaf.SelectionRange); got != leaf.Name {
			t.Fatalf("leaf %s selection text = %q", leaf.Name, got)
		}
		if leaf.Rest != expected.rest || leaf.Default != expected.defaults {
			t.Fatalf("leaf %s modifiers = rest:%t default:%t, want rest:%t default:%t", leaf.Name, leaf.Rest, leaf.Default, expected.rest, expected.defaults)
		}
		if leaf.Provenance.Context != scopeir.BindingContextVariable ||
			leaf.Provenance.PatternKind != "array_pattern" ||
			sourceTextForRange(source, leaf.Provenance.PatternRange) != `[first,,{source: alias = fallback, short = fallback2, [dynamicKey]: computed}, ...tail]` ||
			sourceTextForRange(source, leaf.Provenance.ConstructRange) != string(source[len("const "):len(source)-1]) {
			t.Fatalf("leaf %s provenance mismatch: %#v", leaf.Name, leaf.Provenance)
		}
	}
}

func TestExtractBindingPatternNestedAndLegalEmptyPatterns(t *testing.T) {
	nestedSource := []byte(`const {outer: [nested,,{deep: leaf}]} = input;`)
	nested := extractPatternFromVariableDeclarator(t, nestedSource)
	if len(nested.Diagnostics) != 0 || len(nested.Leaves) != 2 {
		t.Fatalf("nested result = %#v", nested)
	}
	paths := map[string]string{}
	for _, leaf := range nested.Leaves {
		paths[leaf.Name] = bindingPathText(leaf.Path)
	}
	if paths["nested"] != "property:outer/array:0" || paths["leaf"] != "property:outer/array:2/property:deep" {
		t.Fatalf("nested paths = %#v", paths)
	}

	for _, source := range [][]byte{[]byte(`const [] = input;`), []byte(`const {} = input;`), []byte(`const [,,] = input;`)} {
		result := extractPatternFromVariableDeclarator(t, source)
		if len(result.Leaves) != 0 || len(result.Diagnostics) != 0 {
			t.Fatalf("legal empty pattern %q produced %#v", source, result)
		}
	}
}

func TestExtractBindingPatternReportsStructuredDiagnostics(t *testing.T) {
	source := []byte(`const [...target.member] = input;`)
	result := extractPatternFromVariableDeclarator(t, source)
	if len(result.Leaves) != 0 || len(result.Diagnostics) != 1 {
		t.Fatalf("invalid rest result = %#v", result)
	}
	diagnostic := result.Diagnostics[0]
	if diagnostic.Code != scopeir.DiagnosticInvalidRestBinding ||
		diagnostic.FilePath != "src/pattern.ts" ||
		diagnostic.NodeKind != "member_expression" ||
		diagnostic.Reason == "" ||
		bindingPathText(diagnostic.Path) != "array:0" ||
		diagnostic.Provenance.Context != scopeir.BindingContextVariable ||
		diagnostic.Provenance.PatternKind != "array_pattern" ||
		sourceTextForRange(source, diagnostic.Range) != "target.member" {
		t.Fatalf("invalid rest diagnostic = %#v", diagnostic)
	}

	constructSource := []byte(`const value = input;`)
	missing := extractMissingPatternFromVariableDeclarator(t, constructSource)
	if len(missing.Diagnostics) != 1 || missing.Diagnostics[0].Code != scopeir.DiagnosticMalformedBindingNode {
		t.Fatalf("missing-pattern diagnostic = %#v", missing.Diagnostics)
	}
}

func TestExtractBindingPatternHandlesUndefinedAndContextAwareRest(t *testing.T) {
	valid := []struct {
		source string
		name   string
		path   string
		rest   bool
	}{
		{source: `const [undefined] = [1];`, name: "undefined", path: "array:0"},
		{source: `const {x: undefined} = {x: 1};`, name: "undefined", path: "property:x"},
		{source: `const [...undefined] = [1];`, name: "undefined", path: "array:0", rest: true},
		{source: `const {...undefined} = {x: 1};`, name: "undefined", path: "", rest: true},
	}
	for _, test := range valid {
		result := extractPatternFromVariableDeclarator(t, []byte(test.source))
		if len(result.Diagnostics) != 0 || len(result.Leaves) != 1 {
			t.Fatalf("valid undefined pattern %q produced %#v", test.source, result)
		}
		leaf := result.Leaves[0]
		if leaf.Name != test.name || bindingPathText(leaf.Path) != test.path || leaf.Rest != test.rest {
			t.Fatalf("valid undefined leaf for %q = %#v", test.source, leaf)
		}
	}

	nested := extractPatternFromVariableDeclarator(t, []byte(`const [...[a,b]] = [1,2];`))
	if len(nested.Diagnostics) != 0 || len(nested.Leaves) != 2 {
		t.Fatalf("nested array rest produced %#v", nested)
	}
	for _, leaf := range nested.Leaves {
		if !leaf.Rest || (leaf.Name != "a" && leaf.Name != "b") {
			t.Fatalf("nested array rest leaf = %#v", leaf)
		}
	}
}

func TestExtractBindingPatternRejectsMalformedContextInvalidForms(t *testing.T) {
	invalid := []struct {
		source string
		path   string
	}{
		{source: `const {#x: y} = obj;`, path: ""},
		{source: `const {...{a}} = obj;`, path: ""},
		{source: `const {{a}=foo} = obj;`, path: ""},
		{source: `const {a: ...b} = obj;`, path: "property:a"},
	}
	for _, source := range invalid {
		result := extractPatternFromVariableDeclarator(t, []byte(source.source))
		if len(result.Leaves) != 0 || len(result.Diagnostics) != 1 {
			t.Fatalf("invalid pattern %q produced %#v", source.source, result)
		}
		diagnostic := result.Diagnostics[0]
		if diagnostic.Code != scopeir.DiagnosticUnsupportedBindingNode ||
			diagnostic.FilePath == "" || diagnostic.NodeKind == "" || diagnostic.Reason == "" ||
			diagnostic.Provenance.Context != scopeir.BindingContextVariable ||
			bindingPathText(diagnostic.Path) != source.path ||
			diagnostic.Provenance.PatternRange.EndCol <= diagnostic.Provenance.PatternRange.StartCol ||
			diagnostic.Provenance.ConstructRange.EndCol <= diagnostic.Provenance.ConstructRange.StartCol {
			t.Fatalf("invalid pattern diagnostic for %q = %#v", source.source, diagnostic)
		}
	}
}

func extractPatternFromVariableDeclarator(t *testing.T, source []byte) bindingPatternResult {
	t.Helper()
	pool := parser.NewPool(nil, parser.PoolOptions{ParseTimeout: time.Second})
	defer pool.Close()
	parsed, err := pool.Parse(context.Background(), parser.Request{
		FilePath: "src/pattern.ts",
		Language: scanner.TypeScript,
		Source:   source,
	})
	if err != nil {
		t.Fatalf("parse binding pattern: %v", err)
	}
	defer parsed.Close()
	declarator := firstNodeOfKind(parsed.Tree.RootNode(), "variable_declarator")
	if declarator == nil {
		t.Fatalf("missing variable_declarator for %q", source)
	}
	pattern := child(declarator, "name")
	if pattern == nil {
		t.Fatalf("missing binding pattern for %q", source)
	}
	return extractBindingPattern(bindingPatternRequest{
		FilePath:  "src/pattern.ts",
		FileHash:  "hash-pattern",
		Source:    source,
		Context:   scopeir.BindingContextVariable,
		Construct: declarator,
		Pattern:   pattern,
	})
}

func extractMissingPatternFromVariableDeclarator(t *testing.T, source []byte) bindingPatternResult {
	t.Helper()
	pool := parser.NewPool(nil, parser.PoolOptions{ParseTimeout: time.Second})
	defer pool.Close()
	parsed, err := pool.Parse(context.Background(), parser.Request{
		FilePath: "src/pattern.ts",
		Language: scanner.TypeScript,
		Source:   source,
	})
	if err != nil {
		t.Fatalf("parse node: %v", err)
	}
	defer parsed.Close()
	construct := firstNodeOfKind(parsed.Tree.RootNode(), "variable_declarator")
	if construct == nil {
		t.Fatalf("missing variable_declarator for %q", source)
	}
	return extractBindingPattern(bindingPatternRequest{
		FilePath:  "src/pattern.ts",
		FileHash:  "hash-pattern",
		Source:    source,
		Context:   scopeir.BindingContextVariable,
		Construct: construct,
	})
}

func firstNodeOfKind(root *sitter.Node, kind string) *sitter.Node {
	var found *sitter.Node
	walk(root, func(node *sitter.Node) {
		if found == nil && node.Kind() == kind {
			found = node
		}
	})
	return found
}

func bindingPathText(path []scopeir.BindingPathSegment) string {
	parts := make([]string, 0, len(path))
	for _, segment := range path {
		switch segment.Kind {
		case scopeir.BindingPathArrayIndex:
			if segment.ArrayIndex == nil {
				parts = append(parts, "array:nil")
			} else {
				parts = append(parts, "array:"+strconv.Itoa(*segment.ArrayIndex))
			}
		case scopeir.BindingPathStaticProperty:
			parts = append(parts, "property:"+segment.PropertyName)
		case scopeir.BindingPathComputedProperty:
			parts = append(parts, "computed:"+segment.ComputedExpression)
		default:
			parts = append(parts, "unknown:"+string(segment.Kind))
		}
	}
	return strings.Join(parts, "/")
}

func sourceTextForRange(source []byte, rng scopeir.Range) string {
	if rng.StartLine != 1 || rng.EndLine != 1 {
		return ""
	}
	return string(source[rng.StartCol:rng.EndCol])
}

func TestExtractTypeAliasObjectPropertiesHaveNestedOwners(t *testing.T) {
	source := []byte(`type Shape = {
  title: string;
  nested: {
    count: number;
  };
}
`)
	ir := parseAndExtract(t, "src/shape.ts", "hash-shape", scanner.TypeScript, source)

	shape := requireDefinition(t, ir, "Shape", scopeir.NodeTypeAlias)
	title := requireExtractQualifiedDefinition(t, ir, "Shape.title", scopeir.NodeProperty)
	if title.OwnerID != shape.ID {
		t.Fatalf("title owner = %q, want %q; title=%#v shape=%#v", title.OwnerID, shape.ID, title, shape)
	}
	nested := requireExtractQualifiedDefinition(t, ir, "Shape.nested", scopeir.NodeProperty)
	count := requireExtractQualifiedDefinition(t, ir, "Shape.nested.count", scopeir.NodeProperty)
	if count.OwnerID != nested.ID {
		t.Fatalf("nested count owner = %q, want %q; count=%#v nested=%#v", count.OwnerID, nested.ID, count, nested)
	}
}

func TestExtractInlineTypeLiteralPropertiesStayUnowned(t *testing.T) {
	source := []byte(`import { useRef } from "react";

export function Panel() {
  const resizeRef = useRef<{ startX: number; startWidth: number } | null>(null);
  return resizeRef.current?.startX;
}
`)
	ir := parseAndExtract(t, "src/panel.tsx", "hash-panel", scanner.TypeScript, source)

	startX := requireDefinition(t, ir, "startX", scopeir.NodeProperty)
	if startX.OwnerID != "" || startX.QualifiedName != "startX" {
		t.Fatalf("inline type literal property should stay unowned: %#v", startX)
	}
}

func BenchmarkExtractTypeScriptScopeIR(b *testing.B) {
	source := []byte(typescriptParityFixture)
	pool := parser.NewPool(nil, parser.PoolOptions{ParseTimeout: time.Second})
	defer pool.Close()
	parsed, err := pool.Parse(context.Background(), parser.Request{
		FilePath: "src/service.ts",
		Language: scanner.TypeScript,
		Source:   source,
	})
	if err != nil {
		b.Fatalf("parse failed: %v", err)
	}
	defer parsed.Close()
	root := parsed.Tree.RootNode()

	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ir, err := Extract(Request{
			FilePath: "src/service.ts",
			FileHash: "hash-ts",
			Language: scanner.TypeScript,
			Source:   source,
			Root:     root,
		})
		if err != nil {
			b.Fatalf("extract failed: %v", err)
		}
		if len(ir.Definitions) == 0 || len(ir.Calls) == 0 {
			b.Fatalf("incomplete extraction: %#v", ir)
		}
	}
}

type paritySignature struct {
	Scopes         []string `json:"scopes"`
	Definitions    []string `json:"definitions"`
	Imports        []string `json:"imports"`
	Calls          []string `json:"calls"`
	Accesses       []string `json:"accesses"`
	Heritage       []string `json:"heritage"`
	TypeReferences []string `json:"typeReferences"`
	TypeBindings   []string `json:"typeBindings"`
}

func buildParitySignature(ir scopeir.ScopeIR) paritySignature {
	signature := paritySignature{}
	for _, scope := range ir.Scopes {
		signature.Scopes = append(signature.Scopes, string(scope.Kind)+":"+scope.ID)
		for _, binding := range scope.TypeBindings {
			signature.TypeBindings = append(
				signature.TypeBindings,
				binding.Name+":"+binding.Type.RawName+":"+string(binding.Type.Source),
			)
		}
	}
	for _, def := range ir.Definitions {
		signature.Definitions = append(signature.Definitions,
			string(def.Label)+":"+def.QualifiedName+":"+def.ReturnType+":"+def.DeclaredType+":"+def.OwnerID,
		)
	}
	for _, item := range ir.Imports {
		target := ""
		if item.TargetRaw != nil {
			target = *item.TargetRaw
		}
		signature.Imports = append(signature.Imports,
			string(item.Kind)+":"+item.LocalName+":"+item.ImportedName+":"+item.Alias+":"+target,
		)
	}
	for _, call := range ir.Calls {
		signature.Calls = append(signature.Calls,
			call.Name+":"+string(call.CallForm)+":"+call.ExplicitReceiver+":"+formatOptionalInt(call.Arity),
		)
	}
	for _, access := range ir.Accesses {
		signature.Accesses = append(signature.Accesses,
			string(access.Kind)+":"+access.Name+":"+access.ExplicitReceiver,
		)
	}
	for _, item := range ir.Heritage {
		signature.Heritage = append(signature.Heritage, string(item.Kind)+":"+item.Name)
	}
	for _, item := range ir.TypeAnnotations {
		if item.Name == item.Type.RawName {
			signature.TypeReferences = append(signature.TypeReferences, item.Name)
		}
	}
	sort.Strings(signature.Scopes)
	sort.Strings(signature.Definitions)
	sort.Strings(signature.Imports)
	sort.Strings(signature.Calls)
	sort.Strings(signature.Accesses)
	sort.Strings(signature.Heritage)
	sort.Strings(signature.TypeReferences)
	sort.Strings(signature.TypeBindings)
	return signature
}

func formatOptionalInt(value *int) string {
	if value == nil {
		return ""
	}
	return strconv.Itoa(*value)
}

func parseAndExtract(t *testing.T, filePath string, fileHash string, language scanner.Language, source []byte) scopeir.ScopeIR {
	t.Helper()
	pool := parser.NewPool(nil, parser.PoolOptions{ParseTimeout: time.Second})
	defer pool.Close()

	parsed, err := pool.Parse(context.Background(), parser.Request{
		FilePath: filePath,
		Language: language,
		Source:   source,
	})
	if err != nil {
		t.Fatalf("parse failed: %v", err)
	}
	defer parsed.Close()

	ir, err := Extract(Request{
		FilePath: filePath,
		FileHash: fileHash,
		Language: language,
		Source:   source,
		Root:     parsed.Tree.RootNode(),
	})
	if err != nil {
		t.Fatalf("extract failed: %v", err)
	}
	return ir
}

func requireDefinition(t *testing.T, ir scopeir.ScopeIR, name string, label scopeir.NodeLabel) scopeir.DefinitionFact {
	t.Helper()
	for _, def := range ir.Definitions {
		if def.Name == name && def.Label == label {
			return def
		}
	}
	t.Fatalf("missing definition %s/%s in %#v", name, label, ir.Definitions)
	return scopeir.DefinitionFact{}
}

func p3bDefinitionsNamed(ir scopeir.ScopeIR, name string, label scopeir.NodeLabel) []scopeir.DefinitionFact {
	var matches []scopeir.DefinitionFact
	for _, definition := range ir.Definitions {
		if definition.Name == name && definition.Label == label {
			matches = append(matches, definition)
		}
	}
	return matches
}

func p3b1FunctionScopeID(t *testing.T, ir scopeir.ScopeIR, source []byte, marker string) string {
	t.Helper()
	found := ""
	for _, scope := range ir.Scopes {
		if scope.Kind != scopeir.ScopeFunction || !strings.HasPrefix(sourceTextForRange(source, scope.Range), marker) {
			continue
		}
		if found != "" {
			t.Fatalf("multiple function scopes contain %q: %#v", marker, ir.Scopes)
		}
		found = scope.ID
	}
	if found == "" {
		t.Fatalf("missing function scope containing %q: %#v", marker, ir.Scopes)
	}
	return found
}

func p3b1ScopeOwnsAndBinds(ir scopeir.ScopeIR, scopeID string, defID string, name string) bool {
	for _, scope := range ir.Scopes {
		if scope.ID != scopeID {
			continue
		}
		owned := false
		for _, candidate := range scope.OwnedDefIDs {
			if candidate == defID {
				owned = true
				break
			}
		}
		if !owned {
			return false
		}
		for _, binding := range scope.Bindings {
			if binding.Name == name && binding.DefID == defID && binding.Origin == scopeir.BindingLocal {
				return true
			}
		}
	}
	return false
}

func p3b2CatchScope(t *testing.T, ir scopeir.ScopeIR, source []byte, constructText string) scopeir.ScopeFact {
	t.Helper()
	var found *scopeir.ScopeFact
	for index := range ir.Scopes {
		scope := &ir.Scopes[index]
		if scope.Kind != scopeir.ScopeBlock || sourceTextForRange(source, scope.Range) != constructText {
			continue
		}
		if found != nil {
			t.Fatalf("multiple catch scopes match %q: %#v", constructText, ir.Scopes)
		}
		found = scope
	}
	if found == nil {
		t.Fatalf("missing catch scope %q: %#v", constructText, ir.Scopes)
	}
	return *found
}

func p3b2ScopeFactCounts(ir scopeir.ScopeIR, scopeID string, defID string, name string) (int, int) {
	owned := 0
	bindings := 0
	for _, scope := range ir.Scopes {
		if scope.ID != scopeID {
			continue
		}
		for _, candidate := range scope.OwnedDefIDs {
			if candidate == defID {
				owned++
			}
		}
		for _, binding := range scope.Bindings {
			if binding.Name == name && binding.DefID == defID && binding.Origin == scopeir.BindingLocal {
				bindings++
			}
		}
	}
	return owned, bindings
}

func p3b2GlobalScopeFactCounts(ir scopeir.ScopeIR, defID string, name string) (int, int) {
	owned := 0
	bindings := 0
	for _, scope := range ir.Scopes {
		for _, candidate := range scope.OwnedDefIDs {
			if candidate == defID {
				owned++
			}
		}
		for _, binding := range scope.Bindings {
			if binding.Name == name && binding.DefID == defID && binding.Origin == scopeir.BindingLocal {
				bindings++
			}
		}
	}
	return owned, bindings
}

func requireSameImports(t *testing.T, left scopeir.ScopeIR, right scopeir.ScopeIR) {
	t.Helper()
	leftRaw, err := json.Marshal(left.Imports)
	if err != nil {
		t.Fatalf("marshal left imports: %v", err)
	}
	rightRaw, err := json.Marshal(right.Imports)
	if err != nil {
		t.Fatalf("marshal right imports: %v", err)
	}
	if !bytes.Equal(leftRaw, rightRaw) {
		t.Fatalf("import-binding delta: left=%s right=%s", leftRaw, rightRaw)
	}
}

func requireExtractQualifiedDefinition(t *testing.T, ir scopeir.ScopeIR, qualifiedName string, label scopeir.NodeLabel) scopeir.DefinitionFact {
	t.Helper()
	for _, def := range ir.Definitions {
		if def.QualifiedName == qualifiedName && def.Label == label {
			return def
		}
	}
	t.Fatalf("missing qualified definition %s/%s in %#v", qualifiedName, label, ir.Definitions)
	return scopeir.DefinitionFact{}
}

func requireImport(t *testing.T, ir scopeir.ScopeIR, kind scopeir.ImportKind, local string, imported string, target string) {
	t.Helper()
	for _, item := range ir.Imports {
		if item.Kind == kind && item.LocalName == local && item.ImportedName == imported && item.TargetRaw != nil && *item.TargetRaw == target {
			return
		}
	}
	t.Fatalf("missing import kind=%s local=%s imported=%s target=%s in %#v", kind, local, imported, target, ir.Imports)
}

func requireCall(t *testing.T, ir scopeir.ScopeIR, name string, form scopeir.CallForm) {
	t.Helper()
	for _, call := range ir.Calls {
		if call.Name == name && call.CallForm == form {
			return
		}
	}
	t.Fatalf("missing call %s/%s in %#v", name, form, ir.Calls)
}

func requireAccess(t *testing.T, ir scopeir.ScopeIR, name string, kind scopeir.AccessKind) {
	t.Helper()
	for _, access := range ir.Accesses {
		if access.Name == name && access.Kind == kind {
			return
		}
	}
	t.Fatalf("missing access %s/%s in %#v", name, kind, ir.Accesses)
}

func requireHeritage(t *testing.T, ir scopeir.ScopeIR, name string, kind scopeir.HeritageKind) {
	t.Helper()
	for _, item := range ir.Heritage {
		if item.Name == name && item.Kind == kind {
			return
		}
	}
	t.Fatalf("missing heritage %s/%s in %#v", name, kind, ir.Heritage)
}

func requireTypeBinding(t *testing.T, ir scopeir.ScopeIR, name string, rawName string) {
	t.Helper()
	for _, scope := range ir.Scopes {
		for _, binding := range scope.TypeBindings {
			if binding.Name == name && binding.Type.RawName == rawName {
				return
			}
		}
	}
	t.Fatalf("missing type binding %s -> %s in %#v", name, rawName, ir.Scopes)
}

func requireTypeAnnotation(t *testing.T, ir scopeir.ScopeIR, name string) {
	t.Helper()
	for _, item := range ir.TypeAnnotations {
		if item.Name == name {
			return
		}
	}
	t.Fatalf("missing type annotation %s in %#v", name, ir.TypeAnnotations)
}

func requireReturnType(t *testing.T, ir scopeir.ScopeIR, defID string, rawName string) {
	t.Helper()
	for _, item := range ir.ReturnTypes {
		if item.DefID == defID && item.Type.RawName == rawName {
			return
		}
	}
	t.Fatalf("missing return type %s -> %s in %#v", defID, rawName, ir.ReturnTypes)
}
