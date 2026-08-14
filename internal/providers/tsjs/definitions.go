package tsjs

import (
	sitter "github.com/tree-sitter/go-tree-sitter"

	"github.com/tamnguyendinh/anvien/internal/scanner"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func (c *collector) emitDefinition(node *sitter.Node) {
	c.emitDefinitionKind(node, node.Kind())
}

func (c *collector) emitDefinitionKind(node *sitter.Node, kind string) {
	switch kind {
	case "class_declaration", "abstract_class_declaration":
		c.addDefinition(node, scopeir.NodeClass, child(node, "name"), "", "", "", "")
	case "interface_declaration":
		c.addDefinition(node, scopeir.NodeInterface, child(node, "name"), "", "", "", "")
	case "type_alias_declaration":
		c.addDefinition(node, scopeir.NodeTypeAlias, child(node, "name"), "", "", "", "")
	case "enum_declaration":
		c.addDefinition(node, scopeir.NodeEnum, child(node, "name"), "", "", "", "")
	case "function_declaration", "function_signature":
		c.addDefinition(node, scopeir.NodeFunction, child(node, "name"), "", returnTypeNameForCallable(c, node), "", "")
	case "required_parameter", "optional_parameter":
		c.emitParameterBindingPattern(node, parameterPattern(node), formalParameterCallable(node))
	case "arrow_function":
		if c.language == scanner.TypeScript {
			if parameter := child(node, "parameter"); parameter != nil {
				c.emitParameterBindingPattern(parameter, parameter, node)
			}
		}
	case "method_definition", "abstract_method_signature", "method_signature":
		nameNode := child(node, "name")
		label := scopeir.NodeMethod
		if c.text(nameNode) == "constructor" {
			label = scopeir.NodeConstructor
		}
		ownerName := c.ownerDeclarationNameFor(node)
		qualified := ""
		if ownerName != "" && nameNode != nil {
			qualified = ownerName + "." + c.text(nameNode)
		}
		c.addDefinition(
			node,
			label,
			nameNode,
			c.ownerDefIDFor(node),
			returnTypeNameForCallable(c, node),
			"",
			qualified,
		)
		if ownerName != "" {
			c.addTypeBinding(node, "this", ownerName, scopeir.TypeSourceAnnotation)
		}
	case "public_field_definition", "property_signature":
		nameNode := child(node, "name")
		ownerName := c.ownerDeclarationNameFor(node)
		qualified := ""
		if ownerName != "" && nameNode != nil {
			qualified = ownerName + "." + c.text(nameNode)
		}
		c.addDefinition(
			node,
			scopeir.NodeProperty,
			nameNode,
			c.ownerDefIDFor(node),
			"",
			declaredTypeNameForNode(c, node),
			qualified,
		)
	case "variable_declarator":
		nameNode := child(node, "name")
		if nameNode == nil {
			return
		}
		if nameNode.Kind() != "identifier" {
			c.emitVariableBindingPattern(node, nameNode)
			return
		}
		label := scopeir.NodeVariable
		returnType := ""
		if isFunctionExpression(child(node, "value")) {
			label = scopeir.NodeFunction
			returnType = returnTypeNameForCallable(c, child(node, "value"))
		}
		c.addDefinition(node, label, nameNode, "", returnType, declaredTypeNameForNode(c, node), "")
	}
}

func parameterPattern(node *sitter.Node) *sitter.Node {
	pattern := child(node, "pattern")
	if pattern == nil {
		pattern = child(node, "name")
	}
	return pattern
}

func formalParameterCallable(node *sitter.Node) *sitter.Node {
	parameters := node.Parent()
	if parameters == nil || parameters.Kind() != "formal_parameters" {
		return nil
	}
	callable := parameters.Parent()
	if !isFunctionScopeNode(callable) {
		return nil
	}
	ownedParameters := child(callable, "parameters")
	if ownedParameters == nil || ownedParameters.Id() != parameters.Id() {
		return nil
	}
	return callable
}

func (c *collector) emitParameterBindingPattern(node *sitter.Node, pattern *sitter.Node, callable *sitter.Node) {
	if pattern != nil && pattern.Kind() == "this" {
		return
	}
	scopeID := c.parameterCallableScopeID(callable)
	if scopeID == "" {
		return
	}

	result := c.extractParameterBindingPattern(node, pattern)
	c.bindingLeaves = append(c.bindingLeaves, result.Leaves...)
	c.diagnostics = append(c.diagnostics, result.Diagnostics...)
	for _, leaf := range result.Leaves {
		c.addParameterBindingLeafDefinition(leaf, scopeID)
	}
}

func (c *collector) extractParameterBindingPattern(
	construct *sitter.Node,
	pattern *sitter.Node,
) bindingPatternResult {
	request := bindingPatternRequest{
		FilePath:  c.filePath,
		FileHash:  c.fileHash,
		Source:    c.source,
		Context:   scopeir.BindingContextParameter,
		Construct: construct,
		Pattern:   pattern,
	}

	rootPattern := pattern
	isRest := rootPattern != nil && rootPattern.Kind() == "rest_pattern"
	if isRest {
		request.Pattern = firstNamedNonCommentChild(rootPattern)
	}
	result := extractBindingPattern(request)
	if rootPattern == nil {
		return result
	}

	provenance := scopeir.BindingPatternProvenance{
		Context:        scopeir.BindingContextParameter,
		ConstructRange: nodeRange(construct),
		PatternRange:   nodeRange(rootPattern),
		PatternKind:    rootPattern.Kind(),
	}
	hasDefault := child(construct, "value") != nil
	for index := range result.Leaves {
		result.Leaves[index].Provenance = provenance
		if isRest {
			result.Leaves[index].Rest = true
			result.Leaves[index].Range = nodeRange(rootPattern)
		}
		if hasDefault {
			result.Leaves[index].Default = true
			result.Leaves[index].Range = nodeRange(construct)
		} else if rootPattern.Kind() == "identifier" || rootPattern.Kind() == "undefined" {
			result.Leaves[index].Range = nodeRange(construct)
		}
	}
	for index := range result.Diagnostics {
		result.Diagnostics[index].Provenance = provenance
	}
	return result
}

func (c *collector) parameterCallableScopeID(callable *sitter.Node) string {
	if !isFunctionScopeNode(callable) {
		return ""
	}
	id := scopeID(c.filePath, nodeRange(callable), scopeir.ScopeFunction)
	if c.scopeByID(id) != nil {
		return id
	}
	return ""
}

func (c *collector) addParameterBindingLeafDefinition(leaf scopeir.BindingLeafFact, scopeID string) {
	selectionRange := leaf.SelectionRange
	if selectionRange != nil {
		cloned := *selectionRange
		selectionRange = &cloned
	}
	id := defID(c.filePath, leaf.Range, scopeir.NodeVariable, leaf.Name)
	c.definitions = append(c.definitions, scopeir.DefinitionFact{
		ID:             id,
		FilePath:       c.filePath,
		FileHash:       c.fileHash,
		Name:           leaf.Name,
		Label:          scopeir.NodeVariable,
		Range:          leaf.Range,
		SelectionRange: selectionRange,
		QualifiedName:  leaf.Name,
	})

	if scope := c.scopeByID(scopeID); scope != nil {
		scope.OwnedDefIDs = append(scope.OwnedDefIDs, id)
		scope.Bindings = append(scope.Bindings, scopeir.BindingFact{
			Name:   leaf.Name,
			DefID:  id,
			Origin: scopeir.BindingLocal,
		})
	}
}

func (c *collector) emitVariableBindingPattern(node *sitter.Node, pattern *sitter.Node) {
	result := extractBindingPattern(bindingPatternRequest{
		FilePath:  c.filePath,
		FileHash:  c.fileHash,
		Source:    c.source,
		Context:   scopeir.BindingContextVariable,
		Construct: node,
		Pattern:   pattern,
	})
	c.bindingLeaves = append(c.bindingLeaves, result.Leaves...)
	c.diagnostics = append(c.diagnostics, result.Diagnostics...)
	for _, leaf := range result.Leaves {
		c.addVariableBindingLeafDefinition(leaf)
	}
}

func (c *collector) addVariableBindingLeafDefinition(leaf scopeir.BindingLeafFact) {
	selectionRange := leaf.SelectionRange
	if selectionRange != nil {
		cloned := *selectionRange
		selectionRange = &cloned
	}
	id := defID(c.filePath, leaf.Range, scopeir.NodeVariable, leaf.Name)
	c.definitions = append(c.definitions, scopeir.DefinitionFact{
		ID:             id,
		FilePath:       c.filePath,
		FileHash:       c.fileHash,
		Name:           leaf.Name,
		Label:          scopeir.NodeVariable,
		Range:          leaf.Range,
		SelectionRange: selectionRange,
		QualifiedName:  leaf.Name,
	})

	scopeID := c.variableBindingScopeID(leaf.Range)
	if scope := c.scopeByID(scopeID); scope != nil {
		scope.OwnedDefIDs = append(scope.OwnedDefIDs, id)
		scope.Bindings = append(scope.Bindings, scopeir.BindingFact{
			Name:   leaf.Name,
			DefID:  id,
			Origin: scopeir.BindingLocal,
		})
	}
}

func (c *collector) variableBindingScopeID(rng scopeir.Range) string {
	bestID := ""
	bestSpan := int(^uint(0) >> 1)
	bestRank := -1
	for _, scope := range c.scopes {
		if scope.FilePath != c.filePath || !rangeContains(scope.Range, rng) {
			continue
		}
		span := rangeSpan(scope.Range)
		rank := variableBindingScopeRank(scope.Kind)
		if span < bestSpan || (span == bestSpan && rank > bestRank) {
			bestID = scope.ID
			bestSpan = span
			bestRank = rank
		}
	}
	return bestID
}

func variableBindingScopeRank(kind scopeir.ScopeKind) int {
	switch kind {
	case scopeir.ScopeFunction:
		return 2
	case scopeir.ScopeClass:
		return 1
	default:
		return 0
	}
}

func (c *collector) addDefinition(
	node *sitter.Node,
	label scopeir.NodeLabel,
	nameNode *sitter.Node,
	ownerID string,
	returnType string,
	declaredType string,
	qualifiedName string,
) {
	if nameNode == nil {
		return
	}
	name := c.text(nameNode)
	if name == "" {
		return
	}
	rng := nodeRange(node)
	selectionRange := nodeRange(nameNode)
	id := defID(c.filePath, rng, label, name)
	if qualifiedName == "" {
		qualifiedName = name
	}
	fact := scopeir.DefinitionFact{
		ID:             id,
		FilePath:       c.filePath,
		FileHash:       c.fileHash,
		Name:           name,
		Label:          label,
		Range:          rng,
		SelectionRange: &selectionRange,
		QualifiedName:  qualifiedName,
		ReturnType:     returnType,
		DeclaredType:   declaredType,
		OwnerID:        ownerID,
	}
	c.definitions = append(c.definitions, fact)

	scopeID := c.innermostScopeID(rng)
	if scope := c.scopeByID(scopeID); scope != nil {
		scope.OwnedDefIDs = append(scope.OwnedDefIDs, id)
		scope.Bindings = append(scope.Bindings, scopeir.BindingFact{
			Name:   name,
			DefID:  id,
			Origin: scopeir.BindingLocal,
		})
	}

	if returnType != "" {
		if returnNode := child(node, "return_type"); returnNode != nil {
			c.returnTypes = append(c.returnTypes, scopeir.ReturnTypeFact{
				DefID:    id,
				FilePath: c.filePath,
				FileHash: c.fileHash,
				Range:    nodeRange(returnNode),
				Type: scopeir.TypeRef{
					RawName:         returnType,
					DeclaredAtScope: scopeID,
					Source:          scopeir.TypeSourceReturn,
				},
			})
		}
	}
}

func defID(filePath string, rng scopeir.Range, label scopeir.NodeLabel, name string) string {
	return "def:" + filePath + "#" + intString(rng.StartLine) + ":" + intString(rng.StartCol) + ":" + string(label) + ":" + name
}

func returnTypeNameForCallable(c *collector, node *sitter.Node) string {
	returnNode := child(node, "return_type")
	if returnNode == nil {
		return ""
	}
	return stripTypeAnnotation(c.text(returnNode))
}

func declaredTypeNameForNode(c *collector, node *sitter.Node) string {
	typeNode := child(node, "type")
	if typeNode == nil {
		return ""
	}
	return stripTypeAnnotation(c.text(typeNode))
}

func (c *collector) ownerDefIDFor(node *sitter.Node) string {
	owner, label, ok := c.ownerDeclarationFor(node)
	if !ok {
		return ""
	}
	nameNode := child(owner, "name")
	if nameNode == nil {
		return ""
	}
	return defID(c.filePath, nodeRange(owner), label, c.text(nameNode))
}

func (c *collector) ownerDeclarationNameFor(node *sitter.Node) string {
	owner, label, ok := c.ownerDeclarationFor(node)
	if !ok {
		return ""
	}
	if label == scopeir.NodeProperty {
		return c.propertyQualifiedNameFor(owner)
	}
	return c.text(child(owner, "name"))
}

func (c *collector) ownerDeclarationFor(node *sitter.Node) (*sitter.Node, scopeir.NodeLabel, bool) {
	if node.Kind() == "property_signature" {
		if owner := parentPropertySignatureOwner(node); owner != nil {
			return owner, scopeir.NodeProperty, true
		}
		if owner := directTypeAliasObjectOwner(node); owner != nil {
			return owner, scopeir.NodeTypeAlias, true
		}
	}
	current := node.Parent()
	for current != nil {
		if label, ok := ownerDeclarationLabel(current); ok {
			return current, label, true
		}
		current = current.Parent()
	}
	return nil, "", false
}

func (c *collector) propertyQualifiedNameFor(node *sitter.Node) string {
	nameNode := child(node, "name")
	if nameNode == nil {
		return ""
	}
	name := c.text(nameNode)
	ownerName := c.ownerDeclarationNameFor(node)
	if ownerName == "" {
		return name
	}
	return ownerName + "." + name
}

func ownerDeclarationLabel(node *sitter.Node) (scopeir.NodeLabel, bool) {
	switch node.Kind() {
	case "class_declaration", "abstract_class_declaration":
		return scopeir.NodeClass, true
	case "interface_declaration":
		return scopeir.NodeInterface, true
	default:
		return "", false
	}
}

func directTypeAliasObjectOwner(node *sitter.Node) *sitter.Node {
	current := node.Parent()
	for current != nil {
		switch current.Kind() {
		case "type_alias_declaration":
			return current
		case "property_signature", "class_declaration", "abstract_class_declaration", "interface_declaration":
			return nil
		}
		current = current.Parent()
	}
	return nil
}

func parentPropertySignatureOwner(node *sitter.Node) *sitter.Node {
	current := node.Parent()
	for current != nil {
		switch current.Kind() {
		case "property_signature":
			return current
		case "type_alias_declaration", "class_declaration", "abstract_class_declaration", "interface_declaration":
			return nil
		}
		current = current.Parent()
	}
	return nil
}
