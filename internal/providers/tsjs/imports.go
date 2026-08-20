package tsjs

import (
	"strings"

	sitter "github.com/tree-sitter/go-tree-sitter"

	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func (c *collector) emitImport(node *sitter.Node) {
	c.emitImportKind(node, node.Kind())
}

func (c *collector) emitImportKind(node *sitter.Node, kind string) {
	switch kind {
	case "import_statement":
		c.emitImportStatement(node)
	case "export_statement":
		c.emitExportStatement(node)
	case "ERROR", "expression_statement":
		c.emitRecoveredExportDiagnostic(node)
	}
}

func (c *collector) emitImportStatement(node *sitter.Node) {
	sourceNode := child(node, "source")
	if sourceNode == nil {
		return
	}
	targetRaw := stripQuotes(c.text(sourceNode))
	importClause := firstNamedChildOfType(node, "import_clause")
	if importClause == nil {
		return
	}

	namespaceImport := firstDescendantOfType(importClause, "namespace_import")
	if namespaceImport != nil {
		localName := c.text(firstDescendantOfType(namespaceImport, "identifier"))
		if localName != "" {
			c.addImport(scopeir.ImportNamespace, localName, moduleNameFromTarget(targetRaw), "", targetRaw)
			c.importedLocalNames[localName] = struct{}{}
		}
		return
	}

	defaultName := c.text(firstNamedChildOfType(importClause, "identifier"))
	if defaultName != "" {
		c.addImport(scopeir.ImportNamed, defaultName, "default", "", targetRaw)
		c.importedLocalNames[defaultName] = struct{}{}
	}

	for _, specifier := range descendantsOfType(importClause, "import_specifier") {
		names := namedIdentifierChildren(specifier)
		imported := c.text(child(specifier, "name"))
		if imported == "" && len(names) > 0 {
			imported = c.text(names[0])
		}
		if imported == "" {
			continue
		}
		alias := ""
		if len(names) > 1 {
			alias = c.text(names[len(names)-1])
		}
		localName := imported
		kind := scopeir.ImportNamed
		if alias != "" && alias != imported {
			localName = alias
			kind = scopeir.ImportAlias
		}
		c.addImport(kind, localName, imported, alias, targetRaw)
		c.importedLocalNames[localName] = struct{}{}
	}
}

func (c *collector) emitExportStatement(node *sitter.Node) {
	sourceNode := child(node, "source")
	if sourceNode == nil {
		if hasSourceClauseSyntax(node) {
			return
		}
		c.exportStatements = append(c.exportStatements, node)
		return
	}
	targetRaw := stripQuotes(c.text(sourceNode))
	// Preserve the existing compatibility wildcard for both plain star
	// re-exports and namespace-star syntax (`export * as ns from ...`).
	// The latter is represented by a namespace_export descendant rather than
	// an anonymous `*` child by the TypeScript grammar.
	if hasAnonymousChild(node, "*") || containsNodeKind(node, "namespace_export") {
		c.addImport(scopeir.ImportWildcard, "", "", "", targetRaw)
		return
	}
	for _, specifier := range descendantsOfType(node, "export_specifier") {
		names := namedIdentifierChildren(specifier)
		imported := c.text(child(specifier, "name"))
		if imported == "" && len(names) > 0 {
			imported = c.text(names[0])
		}
		if imported == "" {
			continue
		}
		alias := ""
		if len(names) > 1 {
			alias = c.text(names[len(names)-1])
		}
		localName := imported
		if alias != "" {
			localName = alias
		}
		c.addImport(scopeir.ImportReexport, localName, imported, alias, targetRaw)
	}
}

func (c *collector) emitPendingExportFacts() {
	for _, statement := range c.exportStatements {
		c.emitExportFacts(statement)
	}
	c.exportStatements = nil
}

func (c *collector) emitExportFacts(statement *sitter.Node) {
	if statement == nil || hasSourceClauseSyntax(statement) {
		return
	}
	if parent := statement.Parent(); parent == nil || parent.Kind() != "program" {
		c.addExportDiagnostic(
			scopeir.ExportDiagnosticUnsupportedSyntax,
			statement,
			statement,
			"module export extraction requires a top-level export statement",
		)
		return
	}
	if declaration := child(statement, "declaration"); declaration != nil {
		c.emitExportDeclaration(statement, declaration, hasAnonymousChild(statement, "default"))
		return
	}
	if value := child(statement, "value"); value != nil {
		if nodeHasMalformedSyntax(value) {
			c.addExportDiagnostic(
				scopeir.ExportDiagnosticMalformedSyntax,
				value,
				statement,
				"malformed default export expression",
			)
			return
		}
		if !hasAnonymousChild(statement, "default") {
			c.addExportDiagnostic(
				scopeir.ExportDiagnosticUnsupportedSyntax,
				value,
				statement,
				"source-less export expression is not a default export",
			)
			return
		}
		c.emitDefaultExportExpression(statement, value)
		return
	}
	if clause := firstNamedChildOfType(statement, "export_clause"); clause != nil {
		c.emitLocalExportClause(statement, clause)
		return
	}

	code := scopeir.ExportDiagnosticUnsupportedSyntax
	reason := "unsupported source-less export syntax"
	if nodeHasMalformedSyntax(statement) {
		code = scopeir.ExportDiagnosticMalformedSyntax
		reason = "malformed direct or local export syntax"
	}
	c.addExportDiagnostic(
		code,
		statement,
		statement,
		reason,
	)
}

func (c *collector) emitExportDeclaration(
	statement *sitter.Node,
	declaration *sitter.Node,
	defaultExport bool,
) {
	if declaration.Kind() == "ambient_declaration" {
		declaration = firstExportDeclarationChild(declaration)
		if declaration == nil {
			c.addExportDiagnostic(
				scopeir.ExportDiagnosticMalformedSyntax,
				statement,
				statement,
				"ambient export declaration has no declaration payload",
			)
			return
		}
	}

	switch declaration.Kind() {
	case "lexical_declaration", "variable_declaration":
		if defaultExport {
			c.addExportDiagnostic(
				scopeir.ExportDiagnosticUnsupportedSyntax,
				declaration,
				statement,
				"default variable export declaration is unsupported syntax",
			)
			return
		}
		c.emitVariableExportDeclarations(statement, declaration)
		return
	}
	if nodeHasMalformedSyntax(declaration) {
		c.addExportDiagnostic(
			scopeir.ExportDiagnosticMalformedSyntax,
			declaration,
			statement,
			"malformed direct export declaration",
		)
		return
	}

	meanings, supported := exportMeaningsForDeclarationKind(declaration.Kind())
	if !supported {
		c.addExportDiagnostic(
			scopeir.ExportDiagnosticUnsupportedSyntax,
			declaration,
			statement,
			"unsupported direct export declaration kind "+declaration.Kind(),
		)
		return
	}
	nameNode := child(declaration, "name")
	if nameNode == nil || c.text(nameNode) == "" {
		c.addExportDiagnostic(
			scopeir.ExportDiagnosticMalformedSyntax,
			declaration,
			statement,
			"export declaration is missing its binding name",
		)
		return
	}

	localName := c.text(nameNode)
	exportedName := localName
	kind := scopeir.ExportDirect
	if defaultExport {
		exportedName = "default"
		kind = scopeir.ExportDefault
	}
	c.addExportFact(scopeir.ExportFact{
		FilePath:       c.filePath,
		FileHash:       c.fileHash,
		Kind:           kind,
		ExportedName:   exportedName,
		LocalName:      localName,
		LocalDefID:     c.definitionIDForSelection(localName, nodeRange(nameNode)),
		Meanings:       meanings,
		TypeOnly:       meaningsAreTypeOnly(meanings),
		Range:          nodeRange(declaration),
		SelectionRange: rangePointer(nodeRange(nameNode)),
		Provenance: scopeir.ExportProvenance{
			StatementRange: nodeRange(statement),
			SiteKind:       "export_declaration",
		},
	})
}

func (c *collector) emitVariableExportDeclarations(statement, declaration *sitter.Node) {
	declarators := 0
	diagnosticCount := 0
	for index := uint(0); index < declaration.NamedChildCount(); index++ {
		declarator := declaration.NamedChild(index)
		if declarator == nil {
			continue
		}
		if declarator.Kind() != "variable_declarator" {
			if nodeHasMalformedSyntax(declarator) {
				c.addExportDiagnostic(
					scopeir.ExportDiagnosticMalformedSyntax,
					declarator,
					statement,
					"malformed exported variable declaration site",
				)
				diagnosticCount++
			}
			continue
		}
		declarators++
		if nodeHasMalformedSyntax(declarator) {
			c.addExportDiagnostic(
				scopeir.ExportDiagnosticMalformedSyntax,
				declarator,
				statement,
				"malformed exported variable declarator",
			)
			diagnosticCount++
			continue
		}
		nameNode := child(declarator, "name")
		if nameNode == nil {
			c.addExportDiagnostic(
				scopeir.ExportDiagnosticMalformedSyntax,
				declarator,
				statement,
				"exported variable declarator is missing its binding",
			)
			diagnosticCount++
			continue
		}
		if isIdentifierLike(nameNode) {
			name := c.text(nameNode)
			if name == "" {
				c.addExportDiagnostic(
					scopeir.ExportDiagnosticMalformedSyntax,
					declarator,
					statement,
					"exported variable declarator has an empty binding name",
				)
				diagnosticCount++
				continue
			}
			c.addExportFact(scopeir.ExportFact{
				FilePath:       c.filePath,
				FileHash:       c.fileHash,
				Kind:           scopeir.ExportDirect,
				ExportedName:   name,
				LocalName:      name,
				LocalDefID:     c.definitionIDForSelection(name, nodeRange(nameNode)),
				Meanings:       []scopeir.ExportMeaning{scopeir.ExportMeaningValue},
				Range:          nodeRange(declarator),
				SelectionRange: rangePointer(nodeRange(nameNode)),
				Provenance: scopeir.ExportProvenance{
					StatementRange: nodeRange(statement),
					SiteKind:       "export_declaration",
				},
			})
			continue
		}

		declaratorRange := nodeRange(declarator)
		leafCount := 0
		bindingDiagnosticCount := 0
		for _, leaf := range c.bindingLeaves {
			if leaf.Provenance.Context != scopeir.BindingContextVariable ||
				leaf.Provenance.ConstructRange != declaratorRange {
				continue
			}
			selection := copyRangePointer(leaf.SelectionRange)
			localDefID := ""
			if selection != nil {
				localDefID = c.definitionIDForSelection(leaf.Name, *selection)
			}
			c.addExportFact(scopeir.ExportFact{
				FilePath:       c.filePath,
				FileHash:       c.fileHash,
				Kind:           scopeir.ExportDirect,
				ExportedName:   leaf.Name,
				LocalName:      leaf.Name,
				LocalDefID:     localDefID,
				Meanings:       []scopeir.ExportMeaning{scopeir.ExportMeaningValue},
				Range:          leaf.Range,
				SelectionRange: selection,
				Provenance: scopeir.ExportProvenance{
					StatementRange: nodeRange(statement),
					SiteKind:       "export_declaration",
				},
			})
			leafCount++
		}
		for _, diagnostic := range c.diagnostics {
			if diagnostic.Provenance.Context != scopeir.BindingContextVariable ||
				diagnostic.Provenance.ConstructRange != declaratorRange {
				continue
			}
			c.addExportDiagnosticFromExtraction(diagnostic, statement)
			bindingDiagnosticCount++
			diagnosticCount++
		}
		if leafCount == 0 && bindingDiagnosticCount == 0 {
			c.addExportDiagnostic(
				scopeir.ExportDiagnosticUnsupportedSyntax,
				nameNode,
				statement,
				"exported binding pattern has no eligible binding leaves",
			)
			diagnosticCount++
		}
	}
	if declarators == 0 && diagnosticCount == 0 {
		code := scopeir.ExportDiagnosticUnsupportedSyntax
		reason := "export declaration has no eligible variable bindings"
		if nodeHasMalformedSyntax(declaration) {
			code = scopeir.ExportDiagnosticMalformedSyntax
			reason = "malformed export declaration has no eligible variable bindings"
		}
		c.addExportDiagnostic(
			code,
			declaration,
			statement,
			reason,
		)
	}
}

func (c *collector) emitDefaultExportExpression(statement, value *sitter.Node) {
	meanings := []scopeir.ExportMeaning{scopeir.ExportMeaningValue}
	if value.Kind() == "class" || value.Kind() == "class_expression" {
		meanings = []scopeir.ExportMeaning{scopeir.ExportMeaningValue, scopeir.ExportMeaningType}
	}
	c.addExportFact(scopeir.ExportFact{
		FilePath:     c.filePath,
		FileHash:     c.fileHash,
		Kind:         scopeir.ExportDefault,
		ExportedName: "default",
		Meanings:     meanings,
		Range:        nodeRange(value),
		Provenance: scopeir.ExportProvenance{
			StatementRange: nodeRange(statement),
			SiteKind:       "export_default_expression",
		},
	})
}

func (c *collector) emitLocalExportClause(statement, clause *sitter.Node) {
	statementTypeOnly := hasAnonymousChild(statement, "type")
	diagnosticCount := 0
	for index := uint(0); index < clause.NamedChildCount(); index++ {
		specifier := clause.NamedChild(index)
		if specifier == nil || specifier.Kind() == "comment" {
			continue
		}
		if specifier.Kind() != "export_specifier" {
			code := scopeir.ExportDiagnosticUnsupportedSyntax
			reason := "unsupported local export clause node kind " + specifier.Kind()
			if nodeHasMalformedSyntax(specifier) {
				code = scopeir.ExportDiagnosticMalformedSyntax
				reason = "malformed local export specifier site"
			}
			c.addExportDiagnostic(code, specifier, statement, reason)
			diagnosticCount++
			continue
		}
		if nodeHasMalformedSyntax(specifier) {
			c.addExportDiagnostic(
				scopeir.ExportDiagnosticMalformedSyntax,
				specifier,
				statement,
				"malformed local export specifier",
			)
			diagnosticCount++
			continue
		}
		if c.exportSpecifierHasTrailingAliasError(clause, specifier, index) {
			continue
		}
		nameNode := child(specifier, "name")
		if nameNode == nil {
			c.addExportDiagnostic(
				scopeir.ExportDiagnosticMalformedSyntax,
				specifier,
				statement,
				"local export specifier is missing its local name",
			)
			diagnosticCount++
			continue
		}
		if !isIdentifierLike(nameNode) {
			c.addExportDiagnostic(
				scopeir.ExportDiagnosticUnsupportedSyntax,
				nameNode,
				statement,
				"unsupported local export name kind "+nameNode.Kind(),
			)
			diagnosticCount++
			continue
		}

		localName := c.text(nameNode)
		aliasNode := child(specifier, "alias")
		exportedName := localName
		selectionNode := nameNode
		if aliasNode != nil && c.text(aliasNode) != "" {
			exportedName = c.text(aliasNode)
			selectionNode = aliasNode
		}
		if localName == "" || exportedName == "" {
			c.addExportDiagnostic(
				scopeir.ExportDiagnosticMalformedSyntax,
				specifier,
				statement,
				"local export specifier has an empty name",
			)
			diagnosticCount++
			continue
		}

		kind := scopeir.ExportNamed
		if exportedName != localName {
			kind = scopeir.ExportAlias
		}
		typeOnly := statementTypeOnly || hasAnonymousChild(specifier, "type")
		localDefID, meanings := c.localDefinitionEvidence(statement, localName)
		if typeOnly {
			meanings = []scopeir.ExportMeaning{scopeir.ExportMeaningType}
		}
		c.addExportFact(scopeir.ExportFact{
			FilePath:       c.filePath,
			FileHash:       c.fileHash,
			Kind:           kind,
			ExportedName:   exportedName,
			LocalName:      localName,
			LocalDefID:     localDefID,
			Meanings:       meanings,
			TypeOnly:       typeOnly || meaningsAreTypeOnly(meanings),
			Range:          nodeRange(specifier),
			SelectionRange: rangePointer(nodeRange(selectionNode)),
			Provenance: scopeir.ExportProvenance{
				StatementRange: nodeRange(statement),
				SiteKind:       "export_specifier",
			},
		})
	}
	if nodeHasMalformedSyntax(clause) && diagnosticCount == 0 {
		c.addExportDiagnostic(
			scopeir.ExportDiagnosticMalformedSyntax,
			clause,
			statement,
			"malformed local export clause",
		)
	}
}

func (c *collector) exportSpecifierHasTrailingAliasError(
	clause *sitter.Node,
	specifier *sitter.Node,
	specifierIndex uint,
) bool {
	for index := specifierIndex + 1; index < clause.NamedChildCount(); index++ {
		next := clause.NamedChild(index)
		if next == nil || next.Kind() == "comment" {
			continue
		}
		if !nodeHasMalformedSyntax(next) {
			return false
		}
		malformedText := strings.TrimSpace(c.text(next))
		if malformedText != "as" && !strings.HasPrefix(malformedText, "as ") {
			return false
		}
		start := int(specifier.EndByte())
		end := int(next.StartByte())
		if start < 0 || end < start || end > len(c.source) {
			return false
		}
		return !strings.Contains(string(c.source[start:end]), ",")
	}
	return false
}

func (c *collector) emitRecoveredExportDiagnostic(node *sitter.Node) {
	if node == nil || hasSourceClauseSyntax(node) ||
		!hasExportKeywordPrefix(c.text(node)) || !nodeHasMalformedSyntax(node) {
		return
	}
	parent := node.Parent()
	if parent == nil || parent.Kind() != "program" {
		return
	}
	c.addExportDiagnostic(
		scopeir.ExportDiagnosticMalformedSyntax,
		node,
		node,
		"malformed direct or local export syntax",
	)
}

func (c *collector) addExportFact(fact scopeir.ExportFact) {
	c.exports = append(c.exports, fact)
}

func (c *collector) addExportDiagnostic(
	code scopeir.ExportDiagnosticCode,
	site *sitter.Node,
	statement *sitter.Node,
	reason string,
) {
	if site == nil || statement == nil {
		return
	}
	c.exportDiagnostics = append(c.exportDiagnostics, scopeir.ExportDiagnosticFact{
		Code:     code,
		FilePath: c.filePath,
		FileHash: c.fileHash,
		Range:    nodeRange(site),
		NodeKind: site.Kind(),
		Reason:   reason,
		Provenance: scopeir.ExportProvenance{
			StatementRange: nodeRange(statement),
			SiteKind:       statement.Kind(),
		},
	})
}

func (c *collector) addExportDiagnosticFromExtraction(
	diagnostic scopeir.ExtractionDiagnosticFact,
	statement *sitter.Node,
) {
	if statement == nil {
		return
	}
	code := scopeir.ExportDiagnosticUnsupportedSyntax
	if diagnostic.Code == scopeir.DiagnosticMalformedBindingNode {
		code = scopeir.ExportDiagnosticMalformedSyntax
	}
	c.exportDiagnostics = append(c.exportDiagnostics, scopeir.ExportDiagnosticFact{
		Code:     code,
		FilePath: c.filePath,
		FileHash: c.fileHash,
		Range:    diagnostic.Range,
		NodeKind: diagnostic.NodeKind,
		Reason:   "exported binding pattern: " + diagnostic.Reason,
		Provenance: scopeir.ExportProvenance{
			StatementRange: nodeRange(statement),
			SiteKind:       "export_binding_pattern",
		},
	})
}

func (c *collector) definitionIDForSelection(name string, selection scopeir.Range) string {
	matches := make([]string, 0, 1)
	for _, definition := range c.definitions {
		if definition.Name != name || definition.SelectionRange == nil ||
			*definition.SelectionRange != selection {
			continue
		}
		matches = append(matches, definition.ID)
	}
	if len(matches) == 1 {
		return matches[0]
	}
	return ""
}

func (c *collector) localDefinitionEvidence(
	statement *sitter.Node,
	name string,
) (string, []scopeir.ExportMeaning) {
	ids := make(map[string]struct{})
	meanings := make([]scopeir.ExportMeaning, 0, 2)
	for _, definition := range c.definitions {
		if definition.Name != name || !c.definitionIsModuleLocal(statement, definition) {
			continue
		}
		ids[definition.ID] = struct{}{}
		meanings = append(meanings, exportMeaningsForDefinitionLabel(definition.Label)...)
	}
	localDefID := ""
	if len(ids) == 1 {
		for id := range ids {
			localDefID = id
		}
	}
	return localDefID, meanings
}

func (c *collector) definitionIsModuleLocal(
	statement *sitter.Node,
	definition scopeir.DefinitionFact,
) bool {
	if statement == nil || definition.SelectionRange == nil {
		return false
	}
	program := statement.Parent()
	if program == nil || program.Kind() != "program" {
		return false
	}
	for index := uint(0); index < program.NamedChildCount(); index++ {
		if c.topLevelDeclarationDefines(program.NamedChild(index), definition) {
			return true
		}
	}
	return false
}

func (c *collector) topLevelDeclarationDefines(
	declaration *sitter.Node,
	definition scopeir.DefinitionFact,
) bool {
	if declaration == nil || definition.SelectionRange == nil {
		return false
	}
	switch declaration.Kind() {
	case "export_statement":
		if child(declaration, "source") != nil {
			return false
		}
		return c.topLevelDeclarationDefines(child(declaration, "declaration"), definition)
	case "ambient_declaration":
		return c.topLevelDeclarationDefines(firstExportDeclarationChild(declaration), definition)
	case "lexical_declaration", "variable_declaration":
		for index := uint(0); index < declaration.NamedChildCount(); index++ {
			declarator := declaration.NamedChild(index)
			if declarator == nil || declarator.Kind() != "variable_declarator" {
				continue
			}
			nameNode := child(declarator, "name")
			if nameNode == nil {
				continue
			}
			if isIdentifierLike(nameNode) {
				if definition.Name == c.text(nameNode) &&
					*definition.SelectionRange == nodeRange(nameNode) &&
					definition.Range == nodeRange(declarator) {
					return true
				}
				continue
			}
			declaratorRange := nodeRange(declarator)
			for _, leaf := range c.bindingLeaves {
				if leaf.Provenance.Context == scopeir.BindingContextVariable &&
					leaf.Provenance.ConstructRange == declaratorRange &&
					leaf.Name == definition.Name && leaf.SelectionRange != nil &&
					*leaf.SelectionRange == *definition.SelectionRange &&
					leaf.Range == definition.Range {
					return true
				}
			}
		}
		return false
	case "function_declaration", "function_signature", "generator_function_declaration",
		"class_declaration", "abstract_class_declaration", "enum_declaration",
		"interface_declaration", "type_alias_declaration", "internal_module":
		nameNode := child(declaration, "name")
		return nameNode != nil && definition.Name == c.text(nameNode) &&
			*definition.SelectionRange == nodeRange(nameNode) &&
			definition.Range == nodeRange(declaration)
	default:
		return false
	}
}

func exportMeaningsForDeclarationKind(kind string) ([]scopeir.ExportMeaning, bool) {
	switch kind {
	case "function_declaration", "function_signature", "generator_function_declaration":
		return []scopeir.ExportMeaning{scopeir.ExportMeaningValue}, true
	case "class_declaration", "abstract_class_declaration", "enum_declaration":
		return []scopeir.ExportMeaning{scopeir.ExportMeaningValue, scopeir.ExportMeaningType}, true
	case "interface_declaration", "type_alias_declaration":
		return []scopeir.ExportMeaning{scopeir.ExportMeaningType}, true
	case "internal_module":
		return []scopeir.ExportMeaning{scopeir.ExportMeaningValue, scopeir.ExportMeaningNamespace}, true
	default:
		return nil, false
	}
}

func exportMeaningsForDefinitionLabel(label scopeir.NodeLabel) []scopeir.ExportMeaning {
	switch label {
	case scopeir.NodeClass, scopeir.NodeEnum:
		return []scopeir.ExportMeaning{scopeir.ExportMeaningValue, scopeir.ExportMeaningType}
	case scopeir.NodeInterface, scopeir.NodeTypeAlias, scopeir.NodeType:
		return []scopeir.ExportMeaning{scopeir.ExportMeaningType}
	case scopeir.NodeNamespace:
		return []scopeir.ExportMeaning{scopeir.ExportMeaningValue, scopeir.ExportMeaningNamespace}
	default:
		return []scopeir.ExportMeaning{scopeir.ExportMeaningValue}
	}
}

func meaningsAreTypeOnly(meanings []scopeir.ExportMeaning) bool {
	if len(meanings) == 0 {
		return false
	}
	for _, meaning := range meanings {
		if meaning != scopeir.ExportMeaningType {
			return false
		}
	}
	return true
}

func firstExportDeclarationChild(node *sitter.Node) *sitter.Node {
	if node == nil {
		return nil
	}
	for index := uint(0); index < node.NamedChildCount(); index++ {
		candidate := node.NamedChild(index)
		if candidate != nil && candidate.Kind() != "comment" && candidate.Kind() != "decorator" {
			return candidate
		}
	}
	return nil
}

func rangePointer(value scopeir.Range) *scopeir.Range {
	return &value
}

func copyRangePointer(value *scopeir.Range) *scopeir.Range {
	if value == nil {
		return nil
	}
	cloned := *value
	return &cloned
}

func nodeHasMalformedSyntax(node *sitter.Node) bool {
	if node == nil {
		return true
	}
	if node.IsError() || node.IsMissing() {
		return true
	}
	for index := uint(0); index < node.ChildCount(); index++ {
		if nodeHasMalformedSyntax(node.Child(index)) {
			return true
		}
	}
	return false
}

func hasExportKeywordPrefix(value string) bool {
	trimmed := strings.TrimSpace(value)
	if !strings.HasPrefix(trimmed, "export") {
		return false
	}
	if len(trimmed) == len("export") {
		return true
	}
	switch trimmed[len("export")] {
	case ' ', '\t', '\r', '\n', '{', '*', '=':
		return true
	default:
		return false
	}
}

func hasSourceClauseSyntax(node *sitter.Node) bool {
	if node == nil {
		return false
	}
	if child(node, "source") != nil {
		return true
	}
	if node.Kind() == "from" {
		return true
	}
	for index := uint(0); index < node.ChildCount(); index++ {
		candidate := node.Child(index)
		if candidate == nil || (node.Kind() == "export_statement" && candidate.Kind() == "export_clause") {
			continue
		}
		if containsNodeKind(candidate, "from") {
			return true
		}
	}
	return false
}

func containsNodeKind(node *sitter.Node, kind string) bool {
	if node == nil {
		return false
	}
	if node.Kind() == kind {
		return true
	}
	for index := uint(0); index < node.ChildCount(); index++ {
		if containsNodeKind(node.Child(index), kind) {
			return true
		}
	}
	return false
}

func (c *collector) addImport(kind scopeir.ImportKind, localName, importedName, alias, targetRaw string) {
	target := targetRaw
	c.imports = append(c.imports, scopeir.ImportFact{
		FilePath:     c.filePath,
		FileHash:     c.fileHash,
		Kind:         kind,
		LocalName:    localName,
		ImportedName: importedName,
		Alias:        alias,
		TargetRaw:    &target,
	})
}

func hasAnonymousChild(node *sitter.Node, text string) bool {
	if node == nil {
		return false
	}
	for index := uint(0); index < node.ChildCount(); index++ {
		child := node.Child(index)
		if child != nil && child.Kind() == text {
			return true
		}
	}
	return false
}
