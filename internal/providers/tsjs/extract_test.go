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
	if len(ir.BindingLeaves) != 7 {
		t.Fatalf("variable binding leaves = %d, want 7: %#v", len(ir.BindingLeaves), ir.BindingLeaves)
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
	for _, leaf := range ir.BindingLeaves {
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
	if len(ir.BindingLeaves) != 1 || ir.BindingLeaves[0].Name != "local" || bindingPathText(ir.BindingLeaves[0].Path) != "property:untyped" {
		t.Fatalf("inference-miss binding leaves = %#v", ir.BindingLeaves)
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
