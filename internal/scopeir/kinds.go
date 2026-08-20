package scopeir

type NodeLabel string

const (
	NodeProject       NodeLabel = "Project"
	NodePackage       NodeLabel = "Package"
	NodeModule        NodeLabel = "Module"
	NodeFolder        NodeLabel = "Folder"
	NodeFile          NodeLabel = "File"
	NodeClass         NodeLabel = "Class"
	NodeFunction      NodeLabel = "Function"
	NodeMethod        NodeLabel = "Method"
	NodeVariable      NodeLabel = "Variable"
	NodeInterface     NodeLabel = "Interface"
	NodeEnum          NodeLabel = "Enum"
	NodeDecorator     NodeLabel = "Decorator"
	NodeImport        NodeLabel = "Import"
	NodeType          NodeLabel = "Type"
	NodeCodeElement   NodeLabel = "CodeElement"
	NodeCommunity     NodeLabel = "Community"
	NodeProcess       NodeLabel = "Process"
	NodeStruct        NodeLabel = "Struct"
	NodeMacro         NodeLabel = "Macro"
	NodeTypedef       NodeLabel = "Typedef"
	NodeUnion         NodeLabel = "Union"
	NodeNamespace     NodeLabel = "Namespace"
	NodeTrait         NodeLabel = "Trait"
	NodeImpl          NodeLabel = "Impl"
	NodeTypeAlias     NodeLabel = "TypeAlias"
	NodeConst         NodeLabel = "Const"
	NodeStatic        NodeLabel = "Static"
	NodeProperty      NodeLabel = "Property"
	NodeRecord        NodeLabel = "Record"
	NodeDelegate      NodeLabel = "Delegate"
	NodeAnnotation    NodeLabel = "Annotation"
	NodeConstructor   NodeLabel = "Constructor"
	NodeTemplate      NodeLabel = "Template"
	NodeSection       NodeLabel = "Section"
	NodeRoute         NodeLabel = "Route"
	NodeTool          NodeLabel = "Tool"
	NodeResolutionGap NodeLabel = "ResolutionGap"
)

type ScopeKind string

const (
	ScopeModule     ScopeKind = "Module"
	ScopeNamespace  ScopeKind = "Namespace"
	ScopeClass      ScopeKind = "Class"
	ScopeFunction   ScopeKind = "Function"
	ScopeBlock      ScopeKind = "Block"
	ScopeExpression ScopeKind = "Expression"
)

type BindingOrigin string

const (
	BindingLocal     BindingOrigin = "local"
	BindingImport    BindingOrigin = "import"
	BindingNamespace BindingOrigin = "namespace"
	BindingWildcard  BindingOrigin = "wildcard"
	BindingReexport  BindingOrigin = "reexport"
)

type ImportKind string

const (
	ImportNamed             ImportKind = "named"
	ImportAlias             ImportKind = "alias"
	ImportNamespace         ImportKind = "namespace"
	ImportWildcard          ImportKind = "wildcard"
	ImportWildcardExpanded  ImportKind = "wildcard-expanded"
	ImportReexport          ImportKind = "reexport"
	ImportDynamicUnresolved ImportKind = "dynamic-unresolved"
)

// ExportKind records the source-level form of one exported binding or
// specifier. It does not describe terminal resolution or public-API
// reachability.
type ExportKind string

const (
	// ExportDirect is a declaration carrying its own export modifier.
	ExportDirect ExportKind = "direct"
	// ExportNamed is a local named export specifier without an alias.
	ExportNamed ExportKind = "named"
	// ExportDefault is a default declaration or anonymous default expression.
	ExportDefault ExportKind = "default"
	// ExportAlias is a local export specifier whose exported name differs from
	// its local name.
	ExportAlias ExportKind = "alias"
	// ExportReexport is a named/default specifier with a source module clause.
	ExportReexport ExportKind = "reexport"
	// ExportStar is an `export * from` source clause.
	ExportStar ExportKind = "star"
	// ExportNamespace is an `export * as ns from` source clause.
	ExportNamespace ExportKind = "namespace"
)

// ExportMeaning identifies one semantic lane supplied by an export site.
// ExportFact stores these lanes as a canonical set because declarations such
// as classes and enums can contribute both value and type meanings.
type ExportMeaning string

const (
	ExportMeaningValue     ExportMeaning = "value"
	ExportMeaningType      ExportMeaning = "type"
	ExportMeaningNamespace ExportMeaning = "namespace"
)

type TypeRefSource string

const (
	TypeSourceAnnotation        TypeRefSource = "annotation"
	TypeSourceParameter         TypeRefSource = "parameter-annotation"
	TypeSourceReturn            TypeRefSource = "return-annotation"
	TypeSourceSelf              TypeRefSource = "self"
	TypeSourceAssignment        TypeRefSource = "assignment-inferred"
	TypeSourceConstructor       TypeRefSource = "constructor-inferred"
	TypeSourceCallReturn        TypeRefSource = "call-return"
	TypeSourceCallReturnElement TypeRefSource = "call-return-element"
	TypeSourceFieldAccess       TypeRefSource = "field-access"
	TypeSourceMethodReturn      TypeRefSource = "method-return"
	TypeSourceReceiver          TypeRefSource = "receiver-propagated"
)

type CallForm string

const (
	CallFree        CallForm = "free"
	CallMember      CallForm = "member"
	CallConstructor CallForm = "constructor"
	CallIndex       CallForm = "index"
)

type AccessKind string

const (
	AccessRead  AccessKind = "read"
	AccessWrite AccessKind = "write"
)

type HeritageKind string

const (
	HeritageExtends    HeritageKind = "extends"
	HeritageImplements HeritageKind = "implements"
	HeritageTraitImpl  HeritageKind = "trait-impl"
	HeritageInclude    HeritageKind = "include"
	HeritageExtend     HeritageKind = "extend"
	HeritagePrepend    HeritageKind = "prepend"
)
