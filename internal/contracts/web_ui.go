package contracts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
	"github.com/tamnguyendinh/anvien/internal/lbugschema"
	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/semantic"
)

const (
	WebUIContractSchemaPath     = "contracts/web-ui/anvien-web-contract.schema.json"
	WebUIContractTypeScriptPath = "anvien-web/src/generated/anvien-contracts.ts"
)

type LanguageContract struct {
	Name       string   `json:"name"`
	Value      string   `json:"value"`
	Extensions []string `json:"extensions"`
	Syntax     string   `json:"syntax"`
}

type RelationshipDisplayPolicy struct {
	Type          string `json:"type"`
	DisplayLabel  string `json:"displayLabel"`
	SemanticGroup string `json:"semanticGroup"`
	DisplayPolicy string `json:"displayPolicy"`
}

type APIRouteContract struct {
	Method       string   `json:"method"`
	Path         string   `json:"path"`
	QueryParams  []string `json:"queryParams,omitempty"`
	ResponseType string   `json:"responseType"`
	Description  string   `json:"description,omitempty"`
}

type LanguageFactCoverage struct {
	Family            string   `json:"family"`
	Status            string   `json:"status"`
	NodeLabels        []string `json:"nodeLabels"`
	RelationshipTypes []string `json:"relationshipTypes"`
	Evidence          string   `json:"evidence"`
}

type LanguageGraphCoverage struct {
	Language                   string                 `json:"language"`
	ExtractorStatus            string                 `json:"extractorStatus"`
	SourceFactFamilies         []string               `json:"sourceFactFamilies"`
	SupportedNodeLabels        []string               `json:"supportedNodeLabels"`
	SupportedRelationshipTypes []string               `json:"supportedRelationshipTypes"`
	FactFamilies               []LanguageFactCoverage `json:"factFamilies"`
	ResolutionStatus           string                 `json:"resolutionStatus"`
	UnresolvedPolicy           string                 `json:"unresolvedPolicy"`
	WebDisplayPolicy           string                 `json:"webDisplayPolicy"`
	FixtureCoverage            []string               `json:"fixtureCoverage"`
	ProviderParityProofLevel   string                 `json:"providerParityProofLevel"`
	DefaultRegressionGate      string                 `json:"defaultRegressionGate"`
	OptionalExternalAudit      string                 `json:"optionalExternalAudit"`
}

type WebUIContractManifest struct {
	Status        string `json:"status"`
	GeneratedFrom string `json:"generatedFrom"`
	Artifacts     struct {
		SchemaManifest    string `json:"schemaManifest"`
		TypeScriptAdapter string `json:"typescriptAdapter"`
	} `json:"artifacts"`
	Graph struct {
		NodeLabels                           []string                    `json:"nodeLabels"`
		GraphRelationshipTypes               []string                    `json:"graphRelationshipTypes"`
		LadybugDBNodeTables                  []string                    `json:"ladybugdbNodeTables"`
		LadybugDBRelationshipTypes           []string                    `json:"ladybugdbRelationshipTypes"`
		RelationshipDisplayPolicy            []RelationshipDisplayPolicy `json:"relationshipDisplayPolicy"`
		GraphHealthDiagnosticClassifications []string                    `json:"graphHealthDiagnosticClassifications"`
		GraphHealthDiagnosticActionabilities []string                    `json:"graphHealthDiagnosticActionabilities"`
		GraphHealthResolutionHealthBuckets   []string                    `json:"graphHealthResolutionHealthBuckets"`
		GraphHealthResolutionConfidence      []string                    `json:"graphHealthResolutionConfidence"`
		GraphHealthReportTriageDimensions    []string                    `json:"graphHealthReportTriageDimensions"`
		AppLayers                            []string                    `json:"appLayers"`
		AppLayerLabels                       []semantic.TermDefinition   `json:"appLayerLabels"`
		SemanticTerms                        []semantic.TermDefinition   `json:"semanticTerms"`
		SemanticStatusValues                 []string                    `json:"semanticStatusValues"`
		SemanticSchemaVersion                string                      `json:"semanticSchemaVersion"`
		FunctionalAreas                      []string                    `json:"functionalAreas"`
		FunctionalAreaLabels                 []semantic.TermDefinition   `json:"functionalAreaLabels"`
		FileGroups                           []string                    `json:"fileGroups"`
		FileGroupLabels                      []semantic.TermDefinition   `json:"fileGroupLabels"`
		FileRoles                            []string                    `json:"fileRoles"`
		FileRoleLabels                       []semantic.TermDefinition   `json:"fileRoleLabels"`
		RelationshipTableName                string                      `json:"relationshipTableName"`
		EmbeddingTableName                   string                      `json:"embeddingTableName"`
	} `json:"graph"`
	Languages struct {
		CodeLanguages          []LanguageContract      `json:"codeLanguages"`
		GraphCoverage          []LanguageGraphCoverage `json:"graphCoverage"`
		RubyExtensionlessFiles []string                `json:"rubyExtensionlessFiles"`
		AuxiliarySyntax        map[string]string       `json:"auxiliarySyntax"`
		AuxiliaryBasenames     map[string]string       `json:"auxiliaryBasenames"`
	} `json:"languages"`
	Pipeline struct {
		Phases []string `json:"phases"`
	} `json:"pipeline"`
	API struct {
		FileProjectionRoutes []APIRouteContract `json:"fileProjectionRoutes"`
	} `json:"api"`
	Session struct {
		Providers           []string `json:"providers"`
		Availability        []string `json:"availability"`
		ExecutionModes      []string `json:"executionModes"`
		RuntimeEnvironments []string `json:"runtimeEnvironments"`
		ErrorCodes          []string `json:"errorCodes"`
		StreamEvents        []string `json:"streamEvents"`
	} `json:"session"`
}

var nodeLabels = []scopeir.NodeLabel{
	scopeir.NodeProject,
	scopeir.NodePackage,
	scopeir.NodeModule,
	scopeir.NodeFolder,
	scopeir.NodeFile,
	scopeir.NodeClass,
	scopeir.NodeFunction,
	scopeir.NodeMethod,
	scopeir.NodeVariable,
	scopeir.NodeInterface,
	scopeir.NodeEnum,
	scopeir.NodeDecorator,
	scopeir.NodeImport,
	scopeir.NodeType,
	scopeir.NodeCodeElement,
	scopeir.NodeCommunity,
	scopeir.NodeProcess,
	scopeir.NodeStruct,
	scopeir.NodeMacro,
	scopeir.NodeTypedef,
	scopeir.NodeUnion,
	scopeir.NodeNamespace,
	scopeir.NodeTrait,
	scopeir.NodeImpl,
	scopeir.NodeTypeAlias,
	scopeir.NodeConst,
	scopeir.NodeStatic,
	scopeir.NodeProperty,
	scopeir.NodeRecord,
	scopeir.NodeDelegate,
	scopeir.NodeAnnotation,
	scopeir.NodeConstructor,
	scopeir.NodeTemplate,
	scopeir.NodeSection,
	scopeir.NodeRoute,
	scopeir.NodeTool,
	scopeir.NodeResolutionGap,
}

var graphRelationshipTypes = []graph.RelationshipType{
	graph.RelContains,
	graph.RelCalls,
	graph.RelInherits,
	graph.RelMethodOverrides,
	graph.RelMethodImplements,
	graph.RelImports,
	graph.RelUses,
	graph.RelDefines,
	graph.RelDecorates,
	graph.RelImplements,
	graph.RelExtends,
	graph.RelHasMethod,
	graph.RelHasProperty,
	graph.RelAccesses,
	graph.RelMemberOf,
	graph.RelStepInProcess,
	graph.RelHandlesRoute,
	graph.RelFetches,
	graph.RelHandlesTool,
	graph.RelEntryPointOf,
	graph.RelWraps,
	graph.RelQueries,
	graph.RelHasResolutionGap,
}

var codeLanguages = []LanguageContract{
	{Name: "JavaScript", Value: string(scanner.JavaScript), Extensions: []string{".js", ".jsx", ".mjs", ".cjs"}, Syntax: "javascript"},
	{Name: "TypeScript", Value: string(scanner.TypeScript), Extensions: []string{".ts", ".tsx", ".mts", ".cts"}, Syntax: "typescript"},
	{Name: "Python", Value: string(scanner.Python), Extensions: []string{".py"}, Syntax: "python"},
	{Name: "Java", Value: string(scanner.Java), Extensions: []string{".java"}, Syntax: "java"},
	{Name: "C", Value: string(scanner.C), Extensions: []string{".c"}, Syntax: "c"},
	{Name: "CPlusPlus", Value: string(scanner.CPlusPlus), Extensions: []string{".cpp", ".cc", ".cxx", ".h", ".hpp", ".hxx", ".hh"}, Syntax: "cpp"},
	{Name: "CSharp", Value: string(scanner.CSharp), Extensions: []string{".cs"}, Syntax: "csharp"},
	{Name: "Go", Value: string(scanner.Go), Extensions: []string{".go"}, Syntax: "go"},
	{Name: "Ruby", Value: string(scanner.Ruby), Extensions: []string{".rb", ".rake", ".gemspec"}, Syntax: "ruby"},
	{Name: "Rust", Value: string(scanner.Rust), Extensions: []string{".rs"}, Syntax: "rust"},
	{Name: "PHP", Value: string(scanner.PHP), Extensions: []string{".php", ".phtml", ".php3", ".php4", ".php5", ".php8"}, Syntax: "php"},
	{Name: "Kotlin", Value: string(scanner.Kotlin), Extensions: []string{".kt", ".kts"}, Syntax: "kotlin"},
	{Name: "Swift", Value: string(scanner.Swift), Extensions: []string{".swift"}, Syntax: "swift"},
	{Name: "Dart", Value: string(scanner.Dart), Extensions: []string{".dart"}, Syntax: "dart"},
	{Name: "Vue", Value: string(scanner.Vue), Extensions: []string{".vue"}, Syntax: "typescript"},
	{Name: "Svelte", Value: string(scanner.Svelte), Extensions: []string{".svelte"}, Syntax: "typescript"},
	{Name: "Astro", Value: string(scanner.Astro), Extensions: []string{".astro"}, Syntax: "typescript"},
	{Name: "Cobol", Value: string(scanner.Cobol), Extensions: []string{".cbl", ".cob", ".cpy", ".cobol", ".copybook", ".jcl", ".job", ".proc"}, Syntax: "cobol"},
}

var languageGraphCoverage = []LanguageGraphCoverage{
	languageCoverage(
		scanner.JavaScript,
		"scopeir-provider-backed",
		labels(scopeir.NodeClass, scopeir.NodeFunction, scopeir.NodeMethod, scopeir.NodeConstructor, scopeir.NodeVariable, scopeir.NodeProperty),
		providerFactCoverage("resolved-graph-output", "class extends facts resolve to EXTENDS plus compatibility INHERITS when in-repo targets are found"),
		[]string{"tsjs provider unit fixture", "provider call/import parity", "provider heritage parity"},
		"representative endpoint/count-level provider parity",
		"default provider and contract tests",
		"",
	),
	languageCoverage(
		scanner.TypeScript,
		"scopeir-provider-backed",
		labels(scopeir.NodeClass, scopeir.NodeFunction, scopeir.NodeMethod, scopeir.NodeConstructor, scopeir.NodeVariable, scopeir.NodeProperty, scopeir.NodeInterface, scopeir.NodeTypeAlias, scopeir.NodeEnum),
		providerFactCoverage("resolved-graph-output", "class/interface extends and implements facts resolve to EXTENDS/IMPLEMENTS plus compatibility INHERITS when in-repo targets are found"),
		[]string{"tsjs provider unit fixture", "provider call/import/owner/heritage parity", "Restaurant_manager committed heritage fixture"},
		"representative endpoint/count-level provider parity plus Restaurant_manager heritage target fixture",
		"default committed Restaurant_manager heritage fixture",
		"external Restaurant_manager source trace when ANVIEN_RESTAURANT_MANAGER_ROOT is set",
	),
	languageCoverage(
		scanner.Python,
		"scopeir-provider-backed",
		labels(scopeir.NodeClass, scopeir.NodeFunction, scopeir.NodeMethod, scopeir.NodeVariable, scopeir.NodeProperty),
		providerFactCoverage("resolved-graph-output", "class base facts resolve to EXTENDS plus compatibility INHERITS when in-repo targets are found"),
		[]string{"python provider golden fixture", "provider call/import/owner/heritage parity"},
		"representative endpoint/count-level provider parity",
		"default provider and contract tests",
		"",
	),
	languageCoverage(
		scanner.Java,
		"scopeir-provider-backed",
		labels(scopeir.NodePackage, scopeir.NodeClass, scopeir.NodeInterface, scopeir.NodeEnum, scopeir.NodeRecord, scopeir.NodeMethod, scopeir.NodeConstructor, scopeir.NodeProperty, scopeir.NodeVariable),
		providerFactCoverage("resolved-graph-output", "extends/implements facts resolve to EXTENDS/IMPLEMENTS plus compatibility INHERITS when in-repo targets are found"),
		[]string{"java provider golden fixture", "provider call/import/owner/heritage parity", "endpoint graph parity for methods/properties/calls/accesses/uses"},
		"representative endpoint/count-level provider parity",
		"default provider and contract tests",
		"",
	),
	languageCoverage(
		scanner.C,
		"scopeir-provider-backed",
		labels(scopeir.NodeStruct, scopeir.NodeFunction, scopeir.NodeVariable, scopeir.NodeProperty),
		providerFactCoverage("not-applicable-by-language-semantics", "C has no inheritance syntax in the current ScopeIR provider; struct fields and type references are tracked through members/USES"),
		[]string{"c provider golden fixture", "provider call/import parity", "endpoint graph parity for structs/properties/calls/accesses/uses"},
		"representative endpoint/count-level provider parity",
		"default provider and contract tests",
		"",
	),
	languageCoverage(
		scanner.CPlusPlus,
		"scopeir-provider-backed",
		labels(scopeir.NodePackage, scopeir.NodeClass, scopeir.NodeStruct, scopeir.NodeFunction, scopeir.NodeMethod, scopeir.NodeConstructor, scopeir.NodeVariable, scopeir.NodeProperty),
		providerFactCoverage("resolved-graph-output", "base-class facts resolve to EXTENDS plus compatibility INHERITS when in-repo targets are found"),
		[]string{"cpp provider golden fixture", "provider call/import/owner/heritage parity"},
		"representative endpoint/count-level provider parity",
		"default provider and contract tests",
		"",
	),
	languageCoverage(
		scanner.CSharp,
		"scopeir-provider-backed",
		labels(scopeir.NodePackage, scopeir.NodeClass, scopeir.NodeInterface, scopeir.NodeMethod, scopeir.NodeConstructor, scopeir.NodeVariable, scopeir.NodeProperty),
		providerFactCoverage("resolved-graph-output", "base/interface facts resolve to EXTENDS/IMPLEMENTS plus compatibility INHERITS when in-repo targets are found"),
		[]string{"csharp provider golden fixture", "provider call/import/owner/heritage parity"},
		"representative endpoint/count-level provider parity",
		"default provider and contract tests",
		"",
	),
	languageCoverage(
		scanner.Go,
		"scopeir-provider-backed",
		labels(scopeir.NodePackage, scopeir.NodeStruct, scopeir.NodeInterface, scopeir.NodeFunction, scopeir.NodeMethod, scopeir.NodeVariable, scopeir.NodeConst, scopeir.NodeProperty),
		providerFactCoverage("resolved-graph-output", "embedded struct/interface facts resolve to EXTENDS plus compatibility INHERITS when in-repo targets are found"),
		[]string{"go provider golden fixture", "provider call/import/owner/heritage parity"},
		"representative endpoint/count-level provider parity",
		"default provider and contract tests",
		"",
	),
	languageCoverage(
		scanner.Ruby,
		"scopeir-provider-backed",
		labels(scopeir.NodeTrait, scopeir.NodeClass, scopeir.NodeFunction, scopeir.NodeMethod, scopeir.NodeVariable, scopeir.NodeProperty),
		providerFactCoverage("resolved-graph-output", "class inheritance and include/extend/prepend facts resolve to EXTENDS/IMPLEMENTS plus compatibility INHERITS when in-repo targets are found"),
		[]string{"ruby provider golden fixture", "provider call/import/heritage parity"},
		"representative endpoint/count-level provider parity",
		"default provider and contract tests",
		"",
	),
	languageCoverage(
		scanner.Rust,
		"scopeir-provider-backed",
		labels(scopeir.NodeStruct, scopeir.NodeTrait, scopeir.NodeImpl, scopeir.NodeFunction, scopeir.NodeMethod, scopeir.NodeVariable, scopeir.NodeConst, scopeir.NodeProperty),
		providerFactCoverage("resolved-graph-output", "trait impl facts resolve to IMPLEMENTS plus compatibility INHERITS when in-repo targets are found"),
		[]string{"rust provider golden fixture", "provider call/import/owner/heritage parity"},
		"representative endpoint/count-level provider parity",
		"default provider and contract tests",
		"",
	),
	languageCoverage(
		scanner.PHP,
		"scopeir-provider-backed",
		labels(scopeir.NodePackage, scopeir.NodeClass, scopeir.NodeInterface, scopeir.NodeTrait, scopeir.NodeFunction, scopeir.NodeMethod, scopeir.NodeConstructor, scopeir.NodeVariable, scopeir.NodeProperty),
		providerFactCoverage("resolved-graph-output", "extends/implements/use facts resolve to EXTENDS/IMPLEMENTS plus compatibility INHERITS when in-repo targets are found"),
		[]string{"php provider golden fixture", "provider call/heritage parity"},
		"representative endpoint/count-level provider parity",
		"default provider and contract tests",
		"",
	),
	languageCoverage(
		scanner.Kotlin,
		"scopeir-provider-backed",
		labels(scopeir.NodePackage, scopeir.NodeClass, scopeir.NodeInterface, scopeir.NodeFunction, scopeir.NodeMethod, scopeir.NodeConstructor, scopeir.NodeVariable, scopeir.NodeProperty),
		providerFactCoverage("resolved-graph-output", "base/interface facts resolve to EXTENDS/IMPLEMENTS plus compatibility INHERITS when in-repo targets are found"),
		[]string{"kotlin provider golden fixture", "provider call/import/heritage parity"},
		"representative endpoint/count-level provider parity",
		"default provider and contract tests",
		"",
	),
	languageCoverage(
		scanner.Swift,
		"scopeir-provider-backed",
		labels(scopeir.NodeClass, scopeir.NodeStruct, scopeir.NodeInterface, scopeir.NodeEnum, scopeir.NodeFunction, scopeir.NodeMethod, scopeir.NodeConstructor, scopeir.NodeVariable, scopeir.NodeProperty),
		providerFactCoverage("resolved-graph-output", "protocol conformance facts resolve to IMPLEMENTS plus compatibility INHERITS when in-repo targets are found"),
		[]string{"swift provider golden fixture", "provider import/heritage parity"},
		"representative endpoint/count-level provider parity",
		"default provider and contract tests",
		"",
	),
	languageCoverage(
		scanner.Dart,
		"scopeir-provider-backed",
		labels(scopeir.NodeClass, scopeir.NodeInterface, scopeir.NodeFunction, scopeir.NodeMethod, scopeir.NodeConstructor, scopeir.NodeVariable, scopeir.NodeProperty),
		providerFactCoverage("resolved-graph-output", "extends/implements facts resolve to EXTENDS/IMPLEMENTS plus compatibility INHERITS when in-repo targets are found"),
		[]string{"dart provider golden fixture", "provider import/heritage parity"},
		"representative endpoint/count-level provider parity",
		"default provider and contract tests",
		"",
	),
	scriptContainerCoverage(scanner.Vue),
	scriptContainerCoverage(scanner.Svelte),
	scriptContainerCoverage(scanner.Astro),
	languageCoverage(
		scanner.Cobol,
		"dedicated-analyzer-phase",
		labels(scopeir.NodeFile, scopeir.NodeModule, scopeir.NodeSection),
		[]LanguageFactCoverage{
			factCoverage("cobol-programs", "dedicated-analyzer-phase", labels(scopeir.NodeModule), rels(graph.RelDefines, graph.RelContains), "COBOL PROGRAM-ID nodes and nesting relationships are emitted by internal/cobol"),
			factCoverage("cobol-sections-paragraphs", "dedicated-analyzer-phase", labels(scopeir.NodeSection), rels(graph.RelDefines, graph.RelCalls), "COBOL sections, paragraphs, PERFORM, and CALL metrics are emitted by internal/cobol"),
			factCoverage("copybooks", "dedicated-analyzer-phase", labels(scopeir.NodeFile), rels(graph.RelUses), "COPY statements are tracked by dedicated COBOL metrics and copybook expansion tests"),
			factCoverage("jcl", "dedicated-analyzer-phase", labels(scopeir.NodeFile, scopeir.NodeModule), rels(graph.RelCalls), "JCL job/step program links are tracked by dedicated COBOL metrics and JCL tests"),
			factCoverage("scopeir-provider-facts", "scanned-not-extracted", []string{}, []string{}, "COBOL is scanned and processed by the dedicated analyzer phase; it is not routed through ScopeIR provider extraction"),
		},
		[]string{"internal/cobol tests", "analyze pipeline COBOL metrics"},
		"dedicated analyzer metrics, not ScopeIR provider parity",
		"default analyze/cobol tests",
		"",
	),
}

var rubyExtensionlessFiles = []string{
	"Rakefile",
	"Gemfile",
	"Guardfile",
	"Vagrantfile",
	"Brewfile",
}

var auxiliarySyntax = map[string]string{
	"json":       "json",
	"yaml":       "yaml",
	"yml":        "yaml",
	"md":         "markdown",
	"mdx":        "markdown",
	"html":       "markup",
	"htm":        "markup",
	"erb":        "markup",
	"xml":        "markup",
	"css":        "css",
	"scss":       "css",
	"sass":       "css",
	"sh":         "bash",
	"bash":       "bash",
	"zsh":        "bash",
	"sql":        "sql",
	"toml":       "toml",
	"ini":        "ini",
	"dockerfile": "docker",
}

var auxiliaryBasenames = map[string]string{
	"Makefile":   "makefile",
	"Dockerfile": "docker",
}

var pipelinePhases = []string{
	"idle",
	"extracting",
	"structure",
	"parsing",
	"imports",
	"calls",
	"heritage",
	"communities",
	"processes",
	"enriching",
	"complete",
	"error",
}

var sessionProviders = []string{"codex", "claude-code"}
var sessionAvailability = []string{"ready", "not_installed", "not_signed_in", "error"}
var sessionExecutionModes = []string{"sandboxed", "bypass"}
var sessionRuntimeEnvironments = []string{"native", "wsl2"}
var sessionErrorCodes = []string{
	"BAD_REQUEST",
	"INVALID_REPO_BINDING",
	"INVALID_REPO_PATH",
	"REPO_NOT_FOUND",
	"INDEX_REQUIRED",
	"SESSION_NOT_FOUND",
	"SESSION_RUNTIME_UNAVAILABLE",
	"SESSION_NOT_SIGNED_IN",
	"SESSION_START_FAILED",
	"SESSION_CANCELLED",
}
var sessionStreamEvents = []string{
	"session_started",
	"reasoning",
	"content",
	"tool_call",
	"tool_result",
	"error",
	"cancelled",
	"done",
}

func WebUIContract() WebUIContractManifest {
	var manifest WebUIContractManifest
	manifest.Status = "go_owned_web_contract_generated"
	manifest.GeneratedFrom = "internal/contracts"
	manifest.Artifacts.SchemaManifest = WebUIContractSchemaPath
	manifest.Artifacts.TypeScriptAdapter = WebUIContractTypeScriptPath
	manifest.Graph.NodeLabels = labelStrings(nodeLabels)
	manifest.Graph.GraphRelationshipTypes = relationshipStrings(graphRelationshipTypes)
	manifest.Graph.LadybugDBNodeTables = append([]string(nil), lbugschema.NodeTables...)
	manifest.Graph.LadybugDBRelationshipTypes = append([]string(nil), lbugschema.RelationshipTypes...)
	manifest.Graph.RelationshipDisplayPolicy = relationshipDisplayPolicies(graphRelationshipTypes)
	manifest.Graph.GraphHealthDiagnosticClassifications = append([]string(nil), graphhealth.DiagnosticClassifications...)
	manifest.Graph.GraphHealthDiagnosticActionabilities = append([]string(nil), graphhealth.DiagnosticActionabilities...)
	manifest.Graph.GraphHealthResolutionHealthBuckets = graphHealthResolutionHealthBucketStrings()
	manifest.Graph.GraphHealthResolutionConfidence = append([]string(nil), graphhealth.ResolutionConfidenceLevels...)
	manifest.Graph.GraphHealthReportTriageDimensions = []string{"topology", "diagnostic"}
	manifest.Graph.AppLayers = semantic.AppLayerStrings()
	manifest.Graph.AppLayerLabels = semantic.AppLayerDefinitions()
	manifest.Graph.SemanticTerms = semantic.SemanticTermDefinitions()
	manifest.Graph.SemanticStatusValues = semantic.StatusValues()
	manifest.Graph.SemanticSchemaVersion = semantic.SchemaVersion
	manifest.Graph.FunctionalAreas = semantic.FunctionalAreaStrings()
	manifest.Graph.FunctionalAreaLabels = semantic.FunctionalAreaDefinitions()
	manifest.Graph.FileGroups = semantic.FileGroupStrings()
	manifest.Graph.FileGroupLabels = semantic.FileGroupDefinitions()
	manifest.Graph.FileRoles = semantic.FileRoleStrings()
	manifest.Graph.FileRoleLabels = semantic.FileRoleDefinitions()
	manifest.Graph.RelationshipTableName = lbugschema.RelTableName
	manifest.Graph.EmbeddingTableName = lbugschema.EmbeddingTableName
	manifest.Languages.CodeLanguages = append([]LanguageContract(nil), codeLanguages...)
	manifest.Languages.GraphCoverage = append([]LanguageGraphCoverage(nil), languageGraphCoverage...)
	manifest.Languages.RubyExtensionlessFiles = append([]string(nil), rubyExtensionlessFiles...)
	manifest.Languages.AuxiliarySyntax = copyStringMap(auxiliarySyntax)
	manifest.Languages.AuxiliaryBasenames = copyStringMap(auxiliaryBasenames)
	manifest.Pipeline.Phases = append([]string(nil), pipelinePhases...)
	manifest.API.FileProjectionRoutes = []APIRouteContract{
		{
			Method:       "GET",
			Path:         "/api/file-detail",
			QueryParams:  []string{"repo", "path", "relationships", "unresolved", "linked"},
			ResponseType: "FileContextResponse",
			Description:  "File-first graph projection detail for one indexed repository file.",
		},
		{
			Method:       "GET",
			Path:         "/api/file-hotspots",
			QueryParams:  []string{"repo", "sort", "limit", "offset", "kind", "appLayer", "functionalArea", "apiOnly", "changedOnly", "unresolvedOnly", "highFanIn", "highFanOut", "highFanInThreshold", "highFanOutThreshold"},
			ResponseType: "FileHotspotsResponse",
			Description:  "Paginated file projection summary list for file map and hotspot views.",
		},
	}
	manifest.Session.Providers = append([]string(nil), sessionProviders...)
	manifest.Session.Availability = append([]string(nil), sessionAvailability...)
	manifest.Session.ExecutionModes = append([]string(nil), sessionExecutionModes...)
	manifest.Session.RuntimeEnvironments = append([]string(nil), sessionRuntimeEnvironments...)
	manifest.Session.ErrorCodes = append([]string(nil), sessionErrorCodes...)
	manifest.Session.StreamEvents = append([]string(nil), sessionStreamEvents...)
	return manifest
}

func WebUIContractJSON() ([]byte, error) {
	raw, err := json.MarshalIndent(WebUIContract(), "", "  ")
	if err != nil {
		return nil, err
	}
	return append(raw, '\n'), nil
}

func WebUIContractTypeScript() (string, error) {
	manifest := WebUIContract()
	var b strings.Builder
	b.WriteString("/* eslint-disable */\n")
	b.WriteString("// Code generated by `go run ./cmd/generate-web-contracts`; DO NOT EDIT.\n")
	b.WriteString("// Source of truth: internal/contracts and Go runtime contract packages.\n\n")
	writeConstArray(&b, "NODE_LABELS", manifest.Graph.NodeLabels)
	b.WriteString("export type NodeLabel = (typeof NODE_LABELS)[number];\n\n")
	writeConstArray(&b, "GRAPH_RELATIONSHIP_TYPES", manifest.Graph.GraphRelationshipTypes)
	b.WriteString("export type RelationshipType = (typeof GRAPH_RELATIONSHIP_TYPES)[number];\n\n")
	writeConstArray(&b, "NODE_TABLES", manifest.Graph.LadybugDBNodeTables)
	b.WriteString("export type NodeTableName = (typeof NODE_TABLES)[number];\n\n")
	writeConstArray(&b, "REL_TYPES", manifest.Graph.LadybugDBRelationshipTypes)
	b.WriteString("export type RelType = (typeof REL_TYPES)[number];\n\n")
	writeConstObjectArray(&b, "RELATIONSHIP_DISPLAY_POLICY", manifest.Graph.RelationshipDisplayPolicy)
	b.WriteString("export type RelationshipDisplayPolicy = (typeof RELATIONSHIP_DISPLAY_POLICY)[number];\n\n")
	writeConstObjectArray(&b, "FILE_PROJECTION_API_ROUTES", manifest.API.FileProjectionRoutes)
	b.WriteString("export type FileProjectionAPIRoute = (typeof FILE_PROJECTION_API_ROUTES)[number];\n\n")
	writeConstString(&b, "REL_TABLE_NAME", manifest.Graph.RelationshipTableName)
	writeConstString(&b, "EMBEDDING_TABLE_NAME", manifest.Graph.EmbeddingTableName)
	b.WriteString("\n")
	if err := writeLanguageSection(&b, manifest.Languages.CodeLanguages); err != nil {
		return "", err
	}
	writeConstObjectArray(&b, "LANGUAGE_GRAPH_COVERAGE", manifest.Languages.GraphCoverage)
	b.WriteString("export type LanguageGraphCoverage = (typeof LANGUAGE_GRAPH_COVERAGE)[number];\n\n")
	writeConstArray(&b, "RUBY_EXTENSIONLESS_FILES", manifest.Languages.RubyExtensionlessFiles)
	writeConstObject(&b, "AUXILIARY_SYNTAX_MAP", manifest.Languages.AuxiliarySyntax)
	writeConstObject(&b, "AUXILIARY_BASENAME_MAP", manifest.Languages.AuxiliaryBasenames)
	b.WriteString(languageFunctions)
	writeConstArray(&b, "GRAPH_HEALTH_TOPOLOGY_STATUSES", graphHealthTopologyStatusStrings())
	b.WriteString("export type GraphHealthTopologyStatus = (typeof GRAPH_HEALTH_TOPOLOGY_STATUSES)[number];\n\n")
	writeConstArray(&b, "GRAPH_HEALTH_CONFIDENCE_LEVELS", graphhealth.ConfidenceLevels)
	b.WriteString("export type GraphHealthConfidence = (typeof GRAPH_HEALTH_CONFIDENCE_LEVELS)[number];\n\n")
	writeConstArray(&b, "GRAPH_HEALTH_EXPECTED_ISOLATION_REASONS", graphhealth.ExpectedIsolationReasons)
	b.WriteString("export type GraphHealthExpectedIsolationReason = (typeof GRAPH_HEALTH_EXPECTED_ISOLATION_REASONS)[number];\n\n")
	writeConstArray(&b, "GRAPH_HEALTH_DIAGNOSTIC_CLASSIFICATIONS", graphhealth.DiagnosticClassifications)
	b.WriteString("export type GraphHealthDiagnosticClassification = (typeof GRAPH_HEALTH_DIAGNOSTIC_CLASSIFICATIONS)[number];\n\n")
	writeConstArray(&b, "GRAPH_HEALTH_DIAGNOSTIC_ACTIONABILITIES", graphhealth.DiagnosticActionabilities)
	b.WriteString("export type GraphHealthDiagnosticActionability = (typeof GRAPH_HEALTH_DIAGNOSTIC_ACTIONABILITIES)[number];\n\n")
	writeConstArray(&b, "GRAPH_HEALTH_RESOLUTION_HEALTH_BUCKETS", manifest.Graph.GraphHealthResolutionHealthBuckets)
	b.WriteString("export type GraphHealthResolutionHealthBucket = (typeof GRAPH_HEALTH_RESOLUTION_HEALTH_BUCKETS)[number];\n\n")
	writeConstArray(&b, "GRAPH_HEALTH_RESOLUTION_CONFIDENCE_LEVELS", manifest.Graph.GraphHealthResolutionConfidence)
	b.WriteString("export type GraphHealthResolutionConfidence = (typeof GRAPH_HEALTH_RESOLUTION_CONFIDENCE_LEVELS)[number];\n\n")
	writeConstArray(&b, "APP_LAYERS", manifest.Graph.AppLayers)
	b.WriteString("export type AppLayer = (typeof APP_LAYERS)[number];\n\n")
	writeConstObjectArray(&b, "APP_LAYER_LABELS", manifest.Graph.AppLayerLabels)
	b.WriteString("export type AppLayerLabel = (typeof APP_LAYER_LABELS)[number];\n\n")
	writeConstObjectArray(&b, "SEMANTIC_TERMS", manifest.Graph.SemanticTerms)
	b.WriteString("export type SemanticTerm = (typeof SEMANTIC_TERMS)[number];\n\n")
	writeConstArray(&b, "SEMANTIC_STATUS_VALUES", manifest.Graph.SemanticStatusValues)
	b.WriteString("export type SemanticStatusValue = (typeof SEMANTIC_STATUS_VALUES)[number];\n\n")
	writeConstString(&b, "SEMANTIC_SCHEMA_VERSION", manifest.Graph.SemanticSchemaVersion)
	b.WriteString("\n")
	writeConstArray(&b, "FUNCTIONAL_AREAS", manifest.Graph.FunctionalAreas)
	b.WriteString("export type FunctionalArea = (typeof FUNCTIONAL_AREAS)[number];\n\n")
	writeConstObjectArray(&b, "FUNCTIONAL_AREA_LABELS", manifest.Graph.FunctionalAreaLabels)
	b.WriteString("export type FunctionalAreaLabel = (typeof FUNCTIONAL_AREA_LABELS)[number];\n\n")
	writeConstArray(&b, "FILE_GROUPS", manifest.Graph.FileGroups)
	b.WriteString("export type FileGroup = (typeof FILE_GROUPS)[number];\n\n")
	writeConstObjectArray(&b, "FILE_GROUP_LABELS", manifest.Graph.FileGroupLabels)
	b.WriteString("export type FileGroupLabel = (typeof FILE_GROUP_LABELS)[number];\n\n")
	writeConstArray(&b, "FILE_ROLES", manifest.Graph.FileRoles)
	b.WriteString("export type FileRole = (typeof FILE_ROLES)[number];\n\n")
	writeConstObjectArray(&b, "FILE_ROLE_LABELS", manifest.Graph.FileRoleLabels)
	b.WriteString("export type FileRoleLabel = (typeof FILE_ROLE_LABELS)[number];\n\n")
	b.WriteString(graphTypes)
	b.WriteString(fileContextTypes)
	writeConstArray(&b, "PIPELINE_PHASES", manifest.Pipeline.Phases)
	b.WriteString("export type PipelinePhase = (typeof PIPELINE_PHASES)[number];\n\n")
	b.WriteString(pipelineTypes)
	writeConstArray(&b, "SESSION_PROVIDERS", manifest.Session.Providers)
	b.WriteString("export type LocalSessionProvider = (typeof SESSION_PROVIDERS)[number];\n\n")
	writeConstArray(&b, "SESSION_AVAILABILITY", manifest.Session.Availability)
	b.WriteString("export type SessionAvailability = (typeof SESSION_AVAILABILITY)[number];\n\n")
	writeConstArray(&b, "SESSION_EXECUTION_MODES", manifest.Session.ExecutionModes)
	b.WriteString("export type SessionExecutionMode = (typeof SESSION_EXECUTION_MODES)[number];\n\n")
	writeConstArray(&b, "SESSION_RUNTIME_ENVIRONMENTS", manifest.Session.RuntimeEnvironments)
	b.WriteString("export type SessionRuntimeEnvironment = (typeof SESSION_RUNTIME_ENVIRONMENTS)[number];\n\n")
	writeConstArray(&b, "SESSION_ERROR_CODES", manifest.Session.ErrorCodes)
	b.WriteString("export type SessionErrorCode = (typeof SESSION_ERROR_CODES)[number];\n\n")
	writeConstArray(&b, "SESSION_STREAM_EVENTS", manifest.Session.StreamEvents)
	b.WriteString("export type SessionStreamEventType = (typeof SESSION_STREAM_EVENTS)[number];\n\n")
	b.WriteString(sessionTypes)
	return b.String(), nil
}

func labelStrings(labels []scopeir.NodeLabel) []string {
	out := make([]string, 0, len(labels))
	for _, label := range labels {
		out = append(out, string(label))
	}
	return out
}

func relationshipStrings(types []graph.RelationshipType) []string {
	out := make([]string, 0, len(types))
	for _, relType := range types {
		out = append(out, string(relType))
	}
	return out
}

func graphHealthTopologyStatusStrings() []string {
	out := make([]string, 0, len(graphhealth.TopologyStatuses))
	for _, status := range graphhealth.TopologyStatuses {
		out = append(out, string(status))
	}
	return out
}

func graphHealthResolutionHealthBucketStrings() []string {
	out := make([]string, 0, len(graphhealth.ResolutionHealthBuckets))
	for _, bucket := range graphhealth.ResolutionHealthBuckets {
		out = append(out, string(bucket))
	}
	return out
}

func relationshipDisplayPolicies(types []graph.RelationshipType) []RelationshipDisplayPolicy {
	out := make([]RelationshipDisplayPolicy, 0, len(types))
	for _, relType := range types {
		policy := RelationshipDisplayPolicy{
			Type:          string(relType),
			DisplayLabel:  displayLabelForRelationship(relType),
			SemanticGroup: "first-class",
			DisplayPolicy: "count and draw as an independent graph relationship",
		}
		if relType == graph.RelInherits {
			policy.SemanticGroup = "normalized-heritage"
			policy.DisplayPolicy = "group with matching EXTENDS or IMPLEMENTS source-target pairs; count and draw only standalone INHERITS edges as independent relationships"
		}
		if relType == graph.RelHasResolutionGap {
			policy.SemanticGroup = "resolution-health"
			policy.DisplayPolicy = "diagnostic relation from a real source node to a persisted ResolutionGap; do not treat as a resolved code edge"
		}
		out = append(out, policy)
	}
	return out
}

func displayLabelForRelationship(relType graph.RelationshipType) string {
	switch relType {
	case graph.RelInherits:
		return "Normalized Heritage"
	case graph.RelMethodOverrides:
		return "Method Overrides"
	case graph.RelMethodImplements:
		return "Method Implements"
	case graph.RelHasMethod:
		return "Has Method"
	case graph.RelHasProperty:
		return "Has Property"
	case graph.RelMemberOf:
		return "Member Of"
	case graph.RelStepInProcess:
		return "Step In Process"
	case graph.RelHandlesRoute:
		return "Handles Route"
	case graph.RelHandlesTool:
		return "Handles Tool"
	case graph.RelEntryPointOf:
		return "Entry Point Of"
	case graph.RelHasResolutionGap:
		return "Has Resolution Gap"
	default:
		return titleWords(string(relType))
	}
}

func titleWords(value string) string {
	parts := strings.Fields(strings.ToLower(strings.ReplaceAll(value, "_", " ")))
	for index, part := range parts {
		if part == "" {
			continue
		}
		parts[index] = strings.ToUpper(part[:1]) + part[1:]
	}
	return strings.Join(parts, " ")
}

func scriptContainerCoverage(language scanner.Language) LanguageGraphCoverage {
	return languageCoverage(
		language,
		"script-container-backed",
		labels(scopeir.NodeClass, scopeir.NodeFunction, scopeir.NodeMethod, scopeir.NodeConstructor, scopeir.NodeVariable, scopeir.NodeProperty, scopeir.NodeInterface, scopeir.NodeTypeAlias),
		[]LanguageFactCoverage{
			factCoverage("embedded-script-definitions", "resolved-graph-output", labels(scopeir.NodeClass, scopeir.NodeFunction, scopeir.NodeMethod, scopeir.NodeConstructor, scopeir.NodeVariable, scopeir.NodeProperty, scopeir.NodeInterface, scopeir.NodeTypeAlias), rels(graph.RelDefines), "embedded JS/TS definitions are extracted from component script blocks"),
			factCoverage("embedded-script-imports", "resolved-graph-output", labels(scopeir.NodeFile), rels(graph.RelImports, graph.RelUses), "embedded JS/TS imports resolve to in-repo files/definitions when available"),
			factCoverage("embedded-script-calls", "resolved-graph-output", labels(scopeir.NodeFunction, scopeir.NodeMethod, scopeir.NodeConstructor), rels(graph.RelCalls), "embedded JS/TS calls resolve to in-repo callable definitions when available"),
			factCoverage("embedded-script-accesses", "resolved-graph-output", labels(scopeir.NodeProperty, scopeir.NodeVariable), rels(graph.RelAccesses), "embedded JS/TS property/variable accesses resolve when receiver or scope bindings are available"),
			factCoverage("embedded-script-type-references", "resolved-graph-output", labels(scopeir.NodeClass, scopeir.NodeInterface, scopeir.NodeTypeAlias), rels(graph.RelUses), "embedded JS/TS type references resolve to in-repo type definitions when available"),
			factCoverage("embedded-script-members", "resolved-graph-output", labels(scopeir.NodeMethod, scopeir.NodeConstructor, scopeir.NodeProperty), rels(graph.RelHasMethod, graph.RelHasProperty, graph.RelMemberOf), "embedded JS/TS owner/member facts emit owner relationships for methods, constructors, and properties"),
			factCoverage("embedded-script-heritage", "resolved-graph-output", labels(scopeir.NodeClass, scopeir.NodeInterface), rels(graph.RelExtends, graph.RelImplements, graph.RelInherits), "embedded JS/TS heritage facts resolve to EXTENDS/IMPLEMENTS plus compatibility INHERITS when in-repo targets are found"),
			factCoverage("framework-routes-tools-processes", "not-applicable-by-language-semantics", []string{}, []string{}, "route/tool/process edges are framework/analyzer-derived rather than script-container provider facts"),
		},
		[]string{string(language) + " script-container tests", "script-container graph parity count tests"},
		"representative script-container count-level graph parity",
		"default provider and contract tests",
		"",
	)
}

func languageCoverage(
	language scanner.Language,
	extractorStatus string,
	supportedNodeLabels []string,
	facts []LanguageFactCoverage,
	fixtureCoverage []string,
	providerParityProofLevel string,
	defaultRegressionGate string,
	optionalExternalAudit string,
) LanguageGraphCoverage {
	facts = coverageFactsForLanguage(supportedNodeLabels, facts)
	return LanguageGraphCoverage{
		Language:                   string(language),
		ExtractorStatus:            extractorStatus,
		SourceFactFamilies:         factFamilyNames(facts),
		SupportedNodeLabels:        supportedNodeLabels,
		SupportedRelationshipTypes: factRelationshipTypes(facts),
		FactFamilies:               facts,
		ResolutionStatus:           resolutionStatusForExtractor(extractorStatus),
		UnresolvedPolicy:           unresolvedPolicyForExtractor(extractorStatus),
		WebDisplayPolicy:           webDisplayPolicyForExtractor(extractorStatus),
		FixtureCoverage:            append([]string(nil), fixtureCoverage...),
		ProviderParityProofLevel:   providerParityProofLevel,
		DefaultRegressionGate:      defaultRegressionGate,
		OptionalExternalAudit:      optionalExternalAudit,
	}
}

func providerFactCoverage(heritageStatus string, heritageEvidence string) []LanguageFactCoverage {
	heritageLabels := []string{}
	if heritageStatus == "resolved-graph-output" {
		heritageLabels = labels(scopeir.NodeClass, scopeir.NodeInterface, scopeir.NodeStruct, scopeir.NodeTrait)
	}
	return []LanguageFactCoverage{
		factCoverage("definitions", "resolved-graph-output", labels(scopeir.NodeFile), rels(graph.RelDefines), "definitions emit graph nodes and file DEFINES edges for provider-backed files"),
		factCoverage("imports", "resolved-graph-output", labels(scopeir.NodeFile), rels(graph.RelImports, graph.RelUses), "in-repo import targets emit IMPORTS and import-use USES edges; external imports remain unresolved evidence"),
		factCoverage("calls", "resolved-graph-output", labels(scopeir.NodeFunction, scopeir.NodeMethod, scopeir.NodeConstructor), rels(graph.RelCalls), "resolved callable targets emit CALLS edges; unresolved/external calls remain metrics/evidence"),
		factCoverage("accesses", "resolved-graph-output", labels(scopeir.NodeProperty, scopeir.NodeVariable, scopeir.NodeConst, scopeir.NodeStatic), rels(graph.RelAccesses), "resolved member/property accesses emit ACCESSES edges; unresolved accesses remain metrics/evidence"),
		factCoverage("type-references", "resolved-graph-output", labels(scopeir.NodeClass, scopeir.NodeInterface, scopeir.NodeStruct, scopeir.NodeTrait, scopeir.NodeRecord, scopeir.NodeTypeAlias, scopeir.NodeEnum), rels(graph.RelUses), "resolved type annotations and bindings emit USES edges; built-in/external types remain unresolved or intentionally ignored"),
		factCoverage("members", "resolved-graph-output", labels(scopeir.NodeMethod, scopeir.NodeConstructor, scopeir.NodeProperty), rels(graph.RelHasMethod, graph.RelHasProperty, graph.RelMemberOf), "owner/member definitions emit HAS_METHOD, HAS_PROPERTY, and MEMBER_OF-style graph relationships where supported"),
		factCoverage("heritage", heritageStatus, heritageLabels, heritageRelationshipTypes(heritageStatus), heritageEvidence),
		factCoverage("framework-routes-tools-processes", "not-applicable-by-language-semantics", []string{}, []string{}, "route/tool/process edges are framework/analyzer-derived and covered outside provider parity fixtures"),
	}
}

func factCoverage(family string, status string, nodeLabels []string, relationshipTypes []string, evidence string) LanguageFactCoverage {
	return LanguageFactCoverage{
		Family:            family,
		Status:            status,
		NodeLabels:        cloneStrings(nodeLabels),
		RelationshipTypes: cloneStrings(relationshipTypes),
		Evidence:          evidence,
	}
}

func cloneStrings(values []string) []string {
	out := make([]string, len(values))
	copy(out, values)
	return out
}

func labels(values ...scopeir.NodeLabel) []string {
	out := make([]string, 0, len(values))
	for _, value := range values {
		out = append(out, string(value))
	}
	return out
}

func rels(values ...graph.RelationshipType) []string {
	out := make([]string, 0, len(values))
	for _, value := range values {
		out = append(out, string(value))
	}
	return out
}

func heritageRelationshipTypes(status string) []string {
	if status != "resolved-graph-output" {
		return []string{}
	}
	return rels(graph.RelExtends, graph.RelImplements, graph.RelInherits)
}

func factFamilyNames(facts []LanguageFactCoverage) []string {
	out := make([]string, 0, len(facts))
	for _, fact := range facts {
		out = append(out, fact.Family)
	}
	return out
}

func factRelationshipTypes(facts []LanguageFactCoverage) []string {
	seen := map[string]bool{}
	out := make([]string, 0)
	for _, fact := range facts {
		for _, relType := range fact.RelationshipTypes {
			if seen[relType] {
				continue
			}
			seen[relType] = true
			out = append(out, relType)
		}
	}
	return out
}

func coverageFactsForLanguage(supportedNodeLabels []string, facts []LanguageFactCoverage) []LanguageFactCoverage {
	out := make([]LanguageFactCoverage, len(facts))
	copy(out, facts)
	for index := range out {
		if out[index].Family == "definitions" {
			out[index].NodeLabels = uniqueStrings(append([]string{string(scopeir.NodeFile)}, supportedNodeLabels...))
		}
	}
	return out
}

func uniqueStrings(values []string) []string {
	seen := map[string]bool{}
	out := make([]string, 0, len(values))
	for _, value := range values {
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		out = append(out, value)
	}
	return out
}

func resolutionStatusForExtractor(extractorStatus string) string {
	if extractorStatus == "dedicated-analyzer-phase" {
		return "not-scopeir-resolved"
	}
	return "scopeir-resolved-in-repo-targets"
}

func unresolvedPolicyForExtractor(extractorStatus string) string {
	if extractorStatus == "dedicated-analyzer-phase" {
		return "recorded in dedicated analyzer metrics/evidence"
	}
	return "unresolved or external targets are retained in resolution metrics/evidence rather than emitted as resolved graph edges"
}

func webDisplayPolicyForExtractor(extractorStatus string) string {
	if extractorStatus == "dedicated-analyzer-phase" {
		return "graph-present labels and relationships are filterable; ScopeIR provider parity is not implied"
	}
	return "graph-present labels and relationships are filterable; unknown future labels/types use fallback display"
}

func copyStringMap(in map[string]string) map[string]string {
	out := make(map[string]string, len(in))
	for key, value := range in {
		out[key] = value
	}
	return out
}

func writeConstString(b *strings.Builder, name string, value string) {
	raw, _ := json.Marshal(value)
	fmt.Fprintf(b, "export const %s = %s as const;\n", name, raw)
}

func writeConstArray(b *strings.Builder, name string, values []string) {
	raw, _ := json.MarshalIndent(values, "", "  ")
	fmt.Fprintf(b, "export const %s = %s as const;\n\n", name, string(raw))
}

func writeConstObjectArray[T any](b *strings.Builder, name string, values []T) {
	raw, _ := marshalStable(values)
	fmt.Fprintf(b, "export const %s = %s as const;\n\n", name, raw)
}

func writeConstObject(b *strings.Builder, name string, values map[string]string) {
	raw, _ := json.MarshalIndent(values, "", "  ")
	fmt.Fprintf(b, "const %s: Record<string, string> = %s;\n\n", name, string(raw))
}

func writeLanguageSection(b *strings.Builder, languages []LanguageContract) error {
	b.WriteString("export const SupportedLanguages = {\n")
	for _, lang := range languages {
		raw, _ := json.Marshal(lang.Value)
		fmt.Fprintf(b, "  %s: %s,\n", lang.Name, raw)
	}
	b.WriteString("} as const;\n")
	b.WriteString("export type SupportedLanguages = (typeof SupportedLanguages)[keyof typeof SupportedLanguages];\n\n")

	extensionMap := make(map[string][]string, len(languages))
	syntaxMap := make(map[string]string, len(languages))
	for _, lang := range languages {
		extensionMap[lang.Value] = lang.Extensions
		syntaxMap[lang.Value] = lang.Syntax
	}
	extensionRaw, err := marshalStable(extensionMap)
	if err != nil {
		return err
	}
	syntaxRaw, err := marshalStable(syntaxMap)
	if err != nil {
		return err
	}
	fmt.Fprintf(b, "const EXTENSION_MAP = %s as const satisfies Record<SupportedLanguages, readonly string[]>;\n\n", extensionRaw)
	fmt.Fprintf(b, "const SYNTAX_MAP = %s as const satisfies Record<SupportedLanguages, string>;\n\n", syntaxRaw)
	return nil
}

func marshalStable(value any) (string, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetIndent("", "  ")
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(value); err != nil {
		return "", err
	}
	return strings.TrimSpace(buf.String()), nil
}

const languageFunctions = `const extToLang = new Map<string, SupportedLanguages>();
for (const [lang, exts] of Object.entries(EXTENSION_MAP) as [
  SupportedLanguages,
  readonly string[],
][]) {
  for (const ext of exts) {
    extToLang.set(ext, lang);
  }
}

export const getLanguageFromFilename = (filename: string): SupportedLanguages | null => {
  const normalized = filename.replace(/\\/g, '/');
  const lastDot = normalized.lastIndexOf('.');
  if (lastDot >= 0) {
    const ext = normalized.slice(lastDot).toLowerCase();
    const lang = extToLang.get(ext);
    if (lang !== undefined) return lang;
  }

  const basename = normalized.split('/').pop() || normalized;
  if ((RUBY_EXTENSIONLESS_FILES as readonly string[]).includes(basename)) {
    return SupportedLanguages.Ruby;
  }

  return null;
};

export const getSyntaxLanguageFromFilename = (filePath: string): string => {
  const normalized = filePath.replace(/\\/g, '/');
  const lang = getLanguageFromFilename(normalized);
  if (lang) return SYNTAX_MAP[lang];
  const ext = normalized.split('.').pop()?.toLowerCase();
  if (ext && ext in AUXILIARY_SYNTAX_MAP) return AUXILIARY_SYNTAX_MAP[ext];
  const basename = normalized.split('/').pop() || '';
  if (basename in AUXILIARY_BASENAME_MAP) return AUXILIARY_BASENAME_MAP[basename];
  return 'text';
};

`

const graphTypes = `export interface GraphHealthDiagnostic {
  kind: string;
  factFamily?: string;
  sourceNodeId?: string;
  targetText?: string;
  resolutionSource?: string;
  classification?: GraphHealthDiagnosticClassification;
  actionability?: GraphHealthDiagnosticActionability;
  filePath?: string;
  fileHash?: string;
  startLine?: number;
  startCol?: number;
  endLine?: number;
  endCol?: number;
  sourceSiteId?: string;
  sourceSiteStatus?: string;
  proofKind?: string;
  targetRole?: string;
  count?: number;
  note?: string;
  source?: string;
}

export interface GraphHealthNodeMetadata {
  topologyStatus: GraphHealthTopologyStatus;
  countedIncoming: number;
  countedOutgoing: number;
  excludedEdgeCounts?: Record<string, number>;
  componentId?: string;
  componentSize?: number;
  componentRootNodeIds?: string[];
  componentReachableFromRoot: boolean;
  expectedIsolationReasons?: GraphHealthExpectedIsolationReason[];
  diagnostics?: GraphHealthDiagnostic[];
  confidence: GraphHealthConfidence;
  resolutionHealthBuckets?: Partial<Record<GraphHealthResolutionHealthBucket, number>>;
  resolutionGapCount?: number;
  resolutionConfidence: GraphHealthResolutionConfidence;
}

export interface GraphHealthComponentSummary {
  id: string;
  nodeCount: number;
  countedEdgeCount: number;
  detached: boolean;
  reachableFromRoot: boolean;
  rootNodeIds?: string[];
  sampleNodeIds?: string[];
}

export interface GraphHealthTopologyResolutionOverlay {
  nodesWithNoGaps: number;
  nodesWithGaps: number;
  nodesWithDegradedResolution: number;
}

export interface GraphHealthSummary {
  policyVersion: string;
  nodeCount: number;
  countedRelationshipCount: number;
  componentCount: number;
  detachedComponentCount: number;
  rootNodeCount: number;
  unresolvedReferenceCount: number;
  sourceBackedUnresolvedReferenceCount: number;
  unattributedUnresolvedReferenceCount: number;
  topologyStatusCounts: Partial<Record<GraphHealthTopologyStatus, number>>;
  expectedIsolationReasonCounts: Partial<Record<GraphHealthExpectedIsolationReason, number>>;
  confidenceCounts: Partial<Record<GraphHealthConfidence, number>>;
  diagnosticCounts: Record<string, number>;
  diagnosticClassificationCounts: Partial<Record<GraphHealthDiagnosticClassification, number>>;
  diagnosticActionabilityCounts: Partial<Record<GraphHealthDiagnosticActionability, number>>;
  excludedEdgeCounts: Record<string, number>;
  resolutionGapNodeCount: number;
  hasResolutionGapRelationshipCount: number;
  resolutionGapCount: number;
  resolvedReferenceCount: number;
  resolutionHealthBucketCounts: Partial<Record<GraphHealthResolutionHealthBucket, number>>;
  resolutionConfidenceCounts: Partial<Record<GraphHealthResolutionConfidence, number>>;
  resolutionGapFactFamilyCounts: Record<string, number>;
  resolutionGapTargetRoleCounts: Record<string, number>;
  resolutionGapClassificationCounts: Partial<Record<GraphHealthDiagnosticClassification, number>>;
  resolutionGapActionabilityCounts: Partial<Record<GraphHealthDiagnosticActionability, number>>;
  resolutionGapAppLayerCounts: Partial<Record<AppLayer, number>>;
  resolutionGapFunctionalAreaCounts: Partial<Record<FunctionalArea, number>>;
  resolutionGapTopologyStatusCounts: Partial<Record<GraphHealthTopologyStatus, number>>;
  topologyResolutionOverlayCounts: Partial<Record<GraphHealthTopologyStatus, GraphHealthTopologyResolutionOverlay>>;
  largestDetachedComponents?: GraphHealthComponentSummary[];
}

export interface GraphSemanticFieldStatus {
  field: string;
  status: SemanticStatusValue;
  required: boolean;
  totalNodes: number;
  nodesWithField: number;
  missingNodes: number;
  unknownNodes: number;
  nodesWithSource: number;
  missingSourceNodes: number;
  staleIncompleteSchemaEvidence: boolean;
  message?: string;
}

export interface GraphSemanticStatus {
  schemaVersion: string;
  appLayer: GraphSemanticFieldStatus;
  functionalArea: GraphSemanticFieldStatus;
}

export type NodeProperties = {
  name: string;
  filePath: string;
  appLayer?: AppLayer;
  appLayerSource?: string;
  functionalArea?: FunctionalArea;
  functionalAreaSource?: string;
  startLine?: number;
  endLine?: number;
  language?: SupportedLanguages | string;
  isExported?: boolean;
  astFrameworkMultiplier?: number;
  astFrameworkReason?: string;
  heuristicLabel?: string;
  cohesion?: number;
  symbolCount?: number;
  keywords?: string[];
  description?: string;
  enrichedBy?: 'heuristic' | 'llm';
  processType?: 'intra_community' | 'cross_community';
  stepCount?: number;
  communities?: string[];
  entryPointId?: string;
  terminalId?: string;
  entryPointScore?: number;
  entryPointReason?: string;
  parameterCount?: number;
  level?: number;
  returnType?: string;
  declaredType?: string;
  visibility?: string;
  isStatic?: boolean;
  isReadonly?: boolean;
  isAbstract?: boolean;
  isFinal?: boolean;
  isVirtual?: boolean;
  isOverride?: boolean;
  isAsync?: boolean;
  isPartial?: boolean;
  annotations?: string[];
  responseKeys?: string[];
  errorKeys?: string[];
  middleware?: string[];
  topologyStatus?: GraphHealthTopologyStatus;
  countedIncoming?: number;
  countedOutgoing?: number;
  excludedEdgeCounts?: Record<string, number>;
  componentId?: string;
  componentSize?: number;
  componentRootNodeIds?: string[];
  componentReachableFromRoot?: boolean;
  expectedIsolationReasons?: GraphHealthExpectedIsolationReason[];
  diagnostics?: GraphHealthDiagnostic[];
  confidence?: GraphHealthConfidence;
  resolutionHealthBuckets?: Partial<Record<GraphHealthResolutionHealthBucket, number>>;
  resolutionGapCount?: number;
  resolutionConfidence?: GraphHealthResolutionConfidence;
  graphHealth?: GraphHealthNodeMetadata;
  [key: string]: unknown;
};

export interface GraphNode {
  id: string;
  label: NodeLabel;
  properties: NodeProperties;
}

export interface GraphRelationship {
  id: string;
  sourceId: string;
  targetId: string;
  type: RelationshipType;
  confidence: number;
  reason: string;
  step?: number;
  resolutionSource?: string;
  fileHash?: string;
  sourceSiteId?: string;
  sourceSiteIds?: readonly string[];
  sourceSiteCount?: number;
  sourceSiteStatus?: string;
  proofKind?: string;
  targetRole?: string;
  targetText?: string;
  filePath?: string;
  startLine?: number;
  startCol?: number;
  endLine?: number;
  endCol?: number;
  evidence?: readonly {
    readonly kind: string;
    readonly weight: number;
    readonly note?: string;
  }[];
}

export interface GraphResponse {
  nodes: GraphNode[];
  relationships: GraphRelationship[];
  graphHealth?: GraphHealthSummary;
  semanticStatus?: GraphSemanticStatus;
}

export interface GraphHealthComponentExplanation {
  id: string;
  nodeCount: number;
  countedEdgeCount: number;
  detached: boolean;
  reachableFromRoot: boolean;
  rootNodeIds?: string[];
  sampleNodeIds?: string[];
  topologyStatusCounts: Partial<Record<GraphHealthTopologyStatus, number>>;
  expectedIsolationReasonCounts: Partial<Record<GraphHealthExpectedIsolationReason, number>>;
  confidenceCounts: Partial<Record<GraphHealthConfidence, number>>;
  diagnosticCounts: Record<string, number>;
  diagnosticClassificationCounts: Partial<Record<GraphHealthDiagnosticClassification, number>>;
  diagnosticActionabilityCounts: Partial<Record<GraphHealthDiagnosticActionability, number>>;
  resolutionGapCount: number;
  resolutionHealthBucketCounts: Partial<Record<GraphHealthResolutionHealthBucket, number>>;
  resolutionConfidenceCounts: Partial<Record<GraphHealthResolutionConfidence, number>>;
}

export interface GraphHealthExplainResponse {
  kind: 'node' | 'component';
  nodeId?: string;
  componentId?: string;
  node?: GraphNode;
  health?: GraphHealthNodeMetadata;
  component?: GraphHealthComponentExplanation;
  countedIncomingRelationships?: GraphRelationship[];
  countedOutgoingRelationships?: GraphRelationship[];
  excludedRelationships?: GraphRelationship[];
  sampleNodes?: GraphNode[];
  countedRelationshipSamples?: GraphRelationship[];
  excludedRelationshipSamples?: GraphRelationship[];
  sampleLimit?: number;
}

export type GraphHealthReportPriority =
  | 'no_incoming'
  | 'detached_component'
  | 'unresolved_reference'
  | 'true_isolated'
  | 'no_outgoing'
  | 'unknown_connectivity';

export const GRAPH_HEALTH_REPORT_TRIAGE_DIMENSIONS = ['topology', 'diagnostic'] as const;
export type GraphHealthReportTriageDimension = (typeof GRAPH_HEALTH_REPORT_TRIAGE_DIMENSIONS)[number];

export interface GraphHealthReportCandidate {
  nodeId: string;
  label: NodeLabel | string;
  name?: string;
  filePath?: string;
  triagePriority: GraphHealthReportPriority;
  triageDimension: GraphHealthReportTriageDimension;
  topologyStatus: GraphHealthTopologyStatus;
  confidence: GraphHealthConfidence;
  countedIncoming: number;
  countedOutgoing: number;
  excludedEdgeCounts?: Record<string, number>;
  expectedIsolationReasons?: GraphHealthExpectedIsolationReason[];
  diagnostics?: GraphHealthDiagnostic[];
  componentId?: string;
  componentSize?: number;
  componentReachableFromRoot: boolean;
  resolutionHealthBuckets?: Partial<Record<GraphHealthResolutionHealthBucket, number>>;
  resolutionGapCount?: number;
  resolutionConfidence: GraphHealthResolutionConfidence;
}

export interface GraphHealthReportResponse {
  reportType: 'graph_health_candidate_review';
  verdictPolicy: 'candidate_not_confirmed';
  limit: number;
  includeExpected: boolean;
  summary: GraphHealthSummary;
  totalCandidates: number;
  returnedCandidates: number;
  candidates: GraphHealthReportCandidate[];
}

`

const fileContextTypes = `export interface FileProjectionGraphInfo {
  path?: string;
  indexedCommit?: string;
  currentCommit?: string;
  stale: boolean;
  analyzedAt?: string;
}

export interface FileContextAmbiguityCandidate {
  type: string;
  id?: string;
  name?: string;
  file?: string;
  line?: number;
  confidence?: number;
  command?: string;
}

export interface FileContextTarget {
  type: string;
  input: string;
  normalizedPath?: string;
  dispatchMode?: string;
  ambiguityCandidates?: FileContextAmbiguityCandidate[];
}

export interface FileSummary {
  path: string;
  language?: SupportedLanguages | string;
  kind?: string;
  fileGroup?: FileGroup | string;
  fileRole?: FileRole | string;
  appLayer?: AppLayer | string;
  functionalArea?: FunctionalArea | string;
  parseStatus?: string;
  symbolCount: number;
  exportedSymbolCount: number;
  inboundRefCount: number;
  outboundRefCount: number;
  localRelationshipCount: number;
  unresolved: number;
  linkedFlowCount: number;
  linkedTestCount: number;
  risk?: string;
  stale: boolean;
  changedSinceAnalyze: boolean;
}

export interface FileGroupSummary {
  key: FileGroup | string;
  label: string;
  files: number;
  unresolved: number;
  roles?: Record<string, number>;
  appLayers?: Record<string, number>;
  functionalAreas?: Record<string, number>;
  sampleFiles?: string[];
}

export interface FileSourceRange {
  startLine?: number;
  startColumn?: number;
  endLine?: number;
  endColumn?: number;
}

export interface FileSymbolRelationshipCounts {
  local: number;
  inbound: number;
  outbound: number;
  unresolved: number;
}

export interface FileSymbolTreeNode {
  id: string;
  name: string;
  kind: string;
  range: FileSourceRange;
  exported: boolean;
  signature?: string;
  relationshipCounts: FileSymbolRelationshipCounts;
  children?: FileSymbolTreeNode[];
}

export interface FileRelationshipSample {
  sourceFile?: string;
  sourceSymbol?: string;
  sourceRange: FileSourceRange;
  relationshipKind: RelationshipType | string;
  targetFile?: string;
  targetSymbol?: string;
  targetRange: FileSourceRange;
  sourceSiteId?: string;
  proofKind?: string;
  sourceSiteStatus?: string;
}

export interface FileRelationshipGroup {
  total: number;
  counts?: Record<string, number>;
  samples: FileRelationshipSample[];
}

export interface FileRelationshipByFileGroup extends FileRelationshipGroup {
  file: string;
}

export interface FileRelationshipCounts {
  local: number;
  outbound: number;
  inbound: number;
  samplesReturned: number;
}

export interface FileRelationshipSections {
  counts: FileRelationshipCounts;
  local: FileRelationshipGroup;
  outboundByFile: FileRelationshipByFileGroup[];
  inboundByFile: FileRelationshipByFileGroup[];
}

export interface FileUnresolvedSample {
  line?: number;
  column?: number;
  targetText?: string;
  sourceSymbol?: string;
  gapKind?: string;
  classification?: GraphHealthDiagnosticClassification | string;
  actionability?: GraphHealthDiagnosticActionability | string;
  proofKind?: string;
  sourceSiteId?: string;
  sourceSiteStatus?: string;
}

export interface FileUnresolvedGroup {
  sourceSymbol?: string;
  total: number;
  samples: FileUnresolvedSample[];
}

export interface FileUnresolvedSummary {
  total: number;
  byKind?: Record<string, number>;
  byClassification?: Record<string, number>;
  byActionability?: Record<string, number>;
  groups: FileUnresolvedGroup[];
}

export interface FileLinkedItem {
  name: string;
  kind?: string;
  source?: string;
  confidence?: string;
  trace?: string;
}

export interface FileLinkedCounts {
  flows: number;
  routes: number;
  mcpTools: number;
  tests: number;
}

export interface FileLinkedSummary {
  counts: FileLinkedCounts;
  flows: FileLinkedItem[];
  routes: FileLinkedItem[];
  mcpTools: FileLinkedItem[];
  tests: FileLinkedItem[];
}

export interface FileQualitySignals {
  parser?: string;
  resolutionConfidence?: GraphHealthResolutionConfidence | string;
  unresolvedCalls: number;
  unresolvedRefs: number;
  unresolvedImports: number;
  generated: boolean;
  stale: boolean;
  changedSinceAnalyze: boolean;
}

export interface FileContextLimits {
  relationshipSamplesPerGroup: number;
  unresolvedSamplesPerGroup: number;
  linkedSamplesPerKind: number;
}

export interface FileContextResponse {
  repo?: string;
  repoPath?: string;
  graph: FileProjectionGraphInfo;
  target: FileContextTarget;
  summary: FileSummary;
  symbolTree: FileSymbolTreeNode[];
  relationships: FileRelationshipSections;
  unresolved: FileUnresolvedSummary;
  linked: FileLinkedSummary;
  quality: FileQualitySignals;
  limits: FileContextLimits;
}

export interface FileHotspotsResponse {
  repo?: string;
  repoPath?: string;
  graph: FileProjectionGraphInfo;
  total: number;
  offset: number;
  limit: number;
  sort: string;
  fileGroups?: FileGroupSummary[];
  files: FileSummary[];
}

`

const pipelineTypes = `export interface PipelineProgress {
  phase: PipelinePhase;
  percent: number;
  showPercent?: boolean;
  message: string;
  detail?: string;
  targetRepoName?: string;
  stats?: {
    filesProcessed: number;
    totalFiles: number;
    nodesCreated: number;
  };
}

`

const sessionTypes = `export interface SessionRepoBinding {
  repoName?: string;
  repoPath?: string;
}

export interface ResolvedSessionRepo {
  repoName: string;
  repoPath: string;
  indexed: boolean;
  storagePath?: string;
}

export interface SessionRepoResolution extends SessionRepoBinding {
  state: 'indexed' | 'index_required' | 'not_found' | 'invalid';
  resolvedRepoName?: string;
  resolvedRepoPath?: string;
  message?: string;
}

export interface SessionStatus {
  provider: LocalSessionProvider;
  availability: SessionAvailability;
  available: boolean;
  authenticated: boolean;
  executablePath?: string;
  version?: string;
  message?: string;
  recommendedEnvironment?: 'native' | 'wsl2';
  runtimeEnvironment: SessionRuntimeEnvironment;
  executionMode: SessionExecutionMode;
  supportsSse: boolean;
  supportsCancel: boolean;
  supportsMcp: boolean;
}

export interface SessionStatusResponse extends SessionStatus {
  repo?: SessionRepoResolution;
}

export interface SessionChatRequest extends SessionRepoBinding {
  message: string;
}

export interface SessionToolCall {
  id: string;
  name: string;
  args?: Record<string, unknown>;
  result?: string;
  status: 'pending' | 'running' | 'completed' | 'error';
}

interface SessionEventBase {
  sessionId: string;
  provider: LocalSessionProvider;
  repoName: string;
  repoPath: string;
  timestamp: number;
}

export interface SessionStartedEvent extends SessionEventBase {
  type: 'session_started';
  runtimeEnvironment: SessionRuntimeEnvironment;
  executionMode: SessionExecutionMode;
}

export interface SessionReasoningEvent extends SessionEventBase {
  type: 'reasoning';
  reasoning: string;
}

export interface SessionContentEvent extends SessionEventBase {
  type: 'content';
  content: string;
}

export interface SessionToolCallEvent extends SessionEventBase {
  type: 'tool_call';
  toolCall: SessionToolCall;
}

export interface SessionToolResultEvent extends SessionEventBase {
  type: 'tool_result';
  toolCall: SessionToolCall;
}

export interface SessionErrorEvent extends SessionEventBase {
  type: 'error';
  code: SessionErrorCode;
  error: string;
}

export interface SessionCancelledEvent extends SessionEventBase {
  type: 'cancelled';
  reason: string;
}

export interface SessionDoneEvent extends SessionEventBase {
  type: 'done';
  usage?: Record<string, number>;
}

export type SessionStreamEvent =
  | SessionStartedEvent
  | SessionReasoningEvent
  | SessionContentEvent
  | SessionToolCallEvent
  | SessionToolResultEvent
  | SessionErrorEvent
  | SessionCancelledEvent
  | SessionDoneEvent;
`
