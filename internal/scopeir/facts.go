package scopeir

type DefinitionFact struct {
	ID       string    `json:"id"`
	FilePath string    `json:"filePath"`
	FileHash string    `json:"fileHash,omitempty"`
	Name     string    `json:"name"`
	Label    NodeLabel `json:"label"`
	Range    Range     `json:"range"`
	// SelectionRange identifies the declaring token when the provider can supply one.
	SelectionRange         *Range   `json:"selectionRange,omitempty"`
	QualifiedName          string   `json:"qualifiedName,omitempty"`
	ParameterCount         *int     `json:"parameterCount,omitempty"`
	RequiredParameterCount *int     `json:"requiredParameterCount,omitempty"`
	ParameterTypes         []string `json:"parameterTypes,omitempty"`
	ReturnType             string   `json:"returnType,omitempty"`
	DeclaredType           string   `json:"declaredType,omitempty"`
	OwnerID                string   `json:"ownerId,omitempty"`
	Visibility             string   `json:"visibility,omitempty"`
	Static                 *bool    `json:"isStatic,omitempty"`
	Readonly               *bool    `json:"isReadonly,omitempty"`
	Abstract               *bool    `json:"isAbstract,omitempty"`
	Final                  *bool    `json:"isFinal,omitempty"`
	Virtual                *bool    `json:"isVirtual,omitempty"`
	Override               *bool    `json:"isOverride,omitempty"`
	Async                  *bool    `json:"isAsync,omitempty"`
	Partial                *bool    `json:"isPartial,omitempty"`
	Annotations            []string `json:"annotations,omitempty"`
	Description            string   `json:"description,omitempty"`
}

type BindingFact struct {
	Name    string        `json:"name"`
	DefID   string        `json:"defId"`
	Origin  BindingOrigin `json:"origin"`
	ViaID   string        `json:"viaId,omitempty"`
	ViaKind ImportKind    `json:"viaKind,omitempty"`
}

type BindingContext string

const (
	BindingContextVariable  BindingContext = "variable"
	BindingContextParameter BindingContext = "parameter"
	BindingContextCatch     BindingContext = "catch"
	BindingContextForIn     BindingContext = "for-in"
	BindingContextForOf     BindingContext = "for-of"
)

type BindingPathSegmentKind string

const (
	BindingPathArrayIndex       BindingPathSegmentKind = "array-index"
	BindingPathStaticProperty   BindingPathSegmentKind = "static-property"
	BindingPathComputedProperty BindingPathSegmentKind = "computed-property"
)

// BindingPathSegment records one typed source step from a binding-pattern root
// to a declaring leaf. Exactly one kind-specific value is populated.
type BindingPathSegment struct {
	Kind               BindingPathSegmentKind `json:"kind"`
	ArrayIndex         *int                   `json:"arrayIndex,omitempty"`
	PropertyName       string                 `json:"propertyName,omitempty"`
	ComputedExpression string                 `json:"computedExpression,omitempty"`
	SourceRange        Range                  `json:"sourceRange"`
}

// BindingPatternProvenance keeps the declaration-context construct distinct
// from the binding-pattern root. Context adapters populate this contract in
// their owning slices; the P3-A walker itself does not emit declarations.
type BindingPatternProvenance struct {
	Context        BindingContext `json:"context,omitempty"`
	ConstructRange Range          `json:"constructRange"`
	PatternRange   Range          `json:"patternRange"`
	PatternKind    string         `json:"patternKind"`
}

// BindingLeafFact is one legal declaring identifier enumerated from a binding
// pattern. Range covers the leaf's declared binding construct, while
// SelectionRange identifies the declaring token.
type BindingLeafFact struct {
	FilePath       string                   `json:"filePath"`
	FileHash       string                   `json:"fileHash,omitempty"`
	Name           string                   `json:"name"`
	Range          Range                    `json:"range"`
	SelectionRange *Range                   `json:"selectionRange,omitempty"`
	Path           []BindingPathSegment     `json:"path,omitempty"`
	Rest           bool                     `json:"rest,omitempty"`
	Default        bool                     `json:"default,omitempty"`
	Provenance     BindingPatternProvenance `json:"provenance"`
}

type ExtractionDiagnosticCode string

const (
	DiagnosticUnsupportedBindingNode ExtractionDiagnosticCode = "tsjs.binding-pattern.unsupported-node"
	DiagnosticMalformedBindingNode   ExtractionDiagnosticCode = "tsjs.binding-pattern.malformed-node"
	DiagnosticInvalidRestBinding     ExtractionDiagnosticCode = "tsjs.binding-pattern.invalid-rest-target"
)

// ExtractionDiagnosticFact makes an unsupported binding-pattern source site
// deterministic and countable instead of allowing the node to disappear.
type ExtractionDiagnosticFact struct {
	Code       ExtractionDiagnosticCode `json:"code"`
	FilePath   string                   `json:"filePath"`
	FileHash   string                   `json:"fileHash,omitempty"`
	Range      Range                    `json:"range"`
	NodeKind   string                   `json:"nodeKind"`
	Reason     string                   `json:"reason"`
	Path       []BindingPathSegment     `json:"path,omitempty"`
	Provenance BindingPatternProvenance `json:"provenance"`
}

type TypeRef struct {
	RawName         string        `json:"rawName"`
	DeclaredAtScope string        `json:"declaredAtScope"`
	Source          TypeRefSource `json:"source"`
	TypeArgs        []TypeRef     `json:"typeArgs,omitempty"`
}

type TypeBindingFact struct {
	Name string  `json:"name"`
	Type TypeRef `json:"type"`
}

type ScopeFact struct {
	ID           string            `json:"id"`
	Parent       *string           `json:"parent"`
	Kind         ScopeKind         `json:"kind"`
	Range        Range             `json:"range"`
	FilePath     string            `json:"filePath"`
	FileHash     string            `json:"fileHash,omitempty"`
	Bindings     []BindingFact     `json:"bindings,omitempty"`
	OwnedDefIDs  []string          `json:"ownedDefIds,omitempty"`
	TypeBindings []TypeBindingFact `json:"typeBindings,omitempty"`
}

type ImportFact struct {
	ID                 string     `json:"id,omitempty"`
	FilePath           string     `json:"filePath"`
	FileHash           string     `json:"fileHash,omitempty"`
	Kind               ImportKind `json:"kind"`
	LocalName          string     `json:"localName,omitempty"`
	ImportedName       string     `json:"importedName,omitempty"`
	Alias              string     `json:"alias,omitempty"`
	TargetRaw          *string    `json:"targetRaw"`
	TargetFile         *string    `json:"targetFile,omitempty"`
	TargetExportedName string     `json:"targetExportedName,omitempty"`
	TargetModuleScope  string     `json:"targetModuleScope,omitempty"`
	TargetDefID        string     `json:"targetDefId,omitempty"`
	TransitiveVia      []string   `json:"transitiveVia,omitempty"`
	LinkStatus         string     `json:"linkStatus,omitempty"`
}

type CallSiteFact struct {
	FilePath         string   `json:"filePath"`
	FileHash         string   `json:"fileHash,omitempty"`
	Name             string   `json:"name"`
	Range            Range    `json:"range"`
	InScope          string   `json:"inScope"`
	CallForm         CallForm `json:"callForm,omitempty"`
	ExplicitReceiver string   `json:"explicitReceiver,omitempty"`
	Arity            *int     `json:"arity,omitempty"`
	ArgTypes         []string `json:"argTypes,omitempty"`
}

type AccessFact struct {
	FilePath         string     `json:"filePath"`
	FileHash         string     `json:"fileHash,omitempty"`
	Name             string     `json:"name"`
	Kind             AccessKind `json:"kind"`
	Range            Range      `json:"range"`
	InScope          string     `json:"inScope"`
	ExplicitReceiver string     `json:"explicitReceiver,omitempty"`
}

type HeritageFact struct {
	FilePath string       `json:"filePath"`
	FileHash string       `json:"fileHash,omitempty"`
	Name     string       `json:"name"`
	Kind     HeritageKind `json:"kind"`
	Range    Range        `json:"range"`
	InScope  string       `json:"inScope"`
}

type TypeAnnotationFact struct {
	FilePath string  `json:"filePath"`
	FileHash string  `json:"fileHash,omitempty"`
	Name     string  `json:"name"`
	Range    Range   `json:"range"`
	InScope  string  `json:"inScope"`
	Type     TypeRef `json:"type"`
}

type ReturnTypeFact struct {
	DefID    string  `json:"defId"`
	FilePath string  `json:"filePath"`
	FileHash string  `json:"fileHash,omitempty"`
	Range    Range   `json:"range"`
	Type     TypeRef `json:"type"`
}

type FrameworkFact struct {
	DefID                string  `json:"defId"`
	FilePath             string  `json:"filePath"`
	FileHash             string  `json:"fileHash,omitempty"`
	Framework            string  `json:"framework,omitempty"`
	Reason               string  `json:"reason"`
	EntryPointMultiplier float64 `json:"entryPointMultiplier"`
	Range                Range   `json:"range"`
}

type DomainFact struct {
	DefID    string `json:"defId"`
	FilePath string `json:"filePath"`
	FileHash string `json:"fileHash,omitempty"`
	Domain   string `json:"domain"`
	Role     string `json:"role,omitempty"`
	Reason   string `json:"reason,omitempty"`
	Range    Range  `json:"range"`
}
