package tsjs

import (
	"strings"

	sitter "github.com/tree-sitter/go-tree-sitter"

	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

type bindingPatternRequest struct {
	FilePath  string
	FileHash  string
	Source    []byte
	Context   scopeir.BindingContext
	Construct *sitter.Node
	Pattern   *sitter.Node
}

type bindingPatternResult struct {
	Leaves      []scopeir.BindingLeafFact
	Diagnostics []scopeir.ExtractionDiagnosticFact
}

type bindingPatternWalker struct {
	request    bindingPatternRequest
	provenance scopeir.BindingPatternProvenance
	result     bindingPatternResult
}

type bindingModifiers struct {
	rest       bool
	hasDefault bool
	construct  *sitter.Node
}

type bindingContainer uint8

const (
	bindingContainerNone bindingContainer = iota
	bindingContainerArray
	bindingContainerObject
)

// extractBindingPattern enumerates binding leaves only. Declaration-context
// adapters decide when to call it and how to project the returned facts.
func extractBindingPattern(request bindingPatternRequest) bindingPatternResult {
	if request.Pattern == nil {
		if request.Construct == nil {
			return bindingPatternResult{}
		}
		constructRange := nodeRange(request.Construct)
		return bindingPatternResult{Diagnostics: []scopeir.ExtractionDiagnosticFact{{
			Code:     scopeir.DiagnosticMalformedBindingNode,
			FilePath: request.FilePath,
			FileHash: request.FileHash,
			Range:    constructRange,
			NodeKind: request.Construct.Kind(),
			Reason:   "declaration construct is missing its binding pattern",
			Provenance: scopeir.BindingPatternProvenance{
				Context:        request.Context,
				ConstructRange: constructRange,
			},
		}}}
	}
	construct := request.Construct
	if construct == nil {
		construct = request.Pattern
	}
	walker := bindingPatternWalker{
		request: request,
		provenance: scopeir.BindingPatternProvenance{
			Context:        request.Context,
			ConstructRange: nodeRange(construct),
			PatternRange:   nodeRange(request.Pattern),
			PatternKind:    request.Pattern.Kind(),
		},
	}
	walker.walk(request.Pattern, nil, bindingModifiers{}, bindingContainerNone)
	return walker.result
}

func (w *bindingPatternWalker) walk(
	node *sitter.Node,
	path []scopeir.BindingPathSegment,
	modifiers bindingModifiers,
	container bindingContainer,
) {
	if node == nil {
		return
	}
	if node.HasError() {
		w.diagnose(
			node,
			path,
			scopeir.DiagnosticUnsupportedBindingNode,
			"binding-pattern subtree contains grammar errors",
		)
		return
	}
	if node.IsMissing() || node.IsError() || node.Kind() == "ERROR" {
		w.diagnose(
			node,
			path,
			scopeir.DiagnosticMalformedBindingNode,
			"malformed binding-pattern node",
		)
		return
	}

	switch node.Kind() {
	case "identifier", "undefined", "shorthand_property_identifier_pattern":
		w.addLeaf(node, path, modifiers)
	case "array_pattern":
		w.walkArray(node, path, modifiers)
	case "object_pattern":
		w.walkObject(node, path, modifiers)
	case "assignment_pattern", "object_assignment_pattern":
		left := child(node, "left")
		if left == nil {
			w.diagnose(
				node,
				path,
				scopeir.DiagnosticMalformedBindingNode,
				"default binding pattern is missing its left target",
			)
			return
		}
		modifiers.hasDefault = true
		modifiers = withBindingConstruct(modifiers, node)
		w.walk(left, path, modifiers, bindingContainerNone)
	case "rest_pattern":
		w.walkRest(node, path, modifiers, container)
	case "comment":
		return
	default:
		w.diagnose(
			node,
			path,
			scopeir.DiagnosticUnsupportedBindingNode,
			"node kind is not a legal declaration binding target",
		)
	}
}

func (w *bindingPatternWalker) walkArray(
	node *sitter.Node,
	path []scopeir.BindingPathSegment,
	modifiers bindingModifiers,
) {
	arrayIndex := 0
	for childIndex := uint(0); childIndex < node.ChildCount(); childIndex++ {
		candidate := node.Child(childIndex)
		if candidate == nil {
			continue
		}
		switch candidate.Kind() {
		case "[", "]":
			continue
		case ",":
			arrayIndex++
			continue
		case "comment":
			continue
		}
		if !candidate.IsNamed() {
			continue
		}
		index := arrayIndex
		segment := scopeir.BindingPathSegment{
			Kind:        scopeir.BindingPathArrayIndex,
			ArrayIndex:  &index,
			SourceRange: nodeRange(candidate),
		}
		w.walk(candidate, appendBindingPath(path, segment), modifiers, bindingContainerArray)
	}
}

func (w *bindingPatternWalker) walkObject(
	node *sitter.Node,
	path []scopeir.BindingPathSegment,
	modifiers bindingModifiers,
) {
	for childIndex := uint(0); childIndex < node.NamedChildCount(); childIndex++ {
		candidate := node.NamedChild(childIndex)
		if candidate == nil || candidate.Kind() == "comment" {
			continue
		}
		switch candidate.Kind() {
		case "pair_pattern":
			w.walkPair(candidate, path, modifiers)
		case "shorthand_property_identifier_pattern":
			segment := w.staticPropertySegment(candidate)
			w.walk(candidate, appendBindingPath(path, segment), modifiers, bindingContainerObject)
		case "object_assignment_pattern":
			left := child(candidate, "left")
			if left == nil {
				w.diagnose(
					candidate,
					path,
					scopeir.DiagnosticMalformedBindingNode,
					"object default binding is missing its left target",
				)
				continue
			}
			if left.Kind() != "shorthand_property_identifier_pattern" && left.Kind() != "undefined" {
				w.diagnose(
					left,
					path,
					scopeir.DiagnosticUnsupportedBindingNode,
					"object default binding requires a shorthand identifier target",
				)
				continue
			}
			childPath := path
			if left.Kind() == "shorthand_property_identifier_pattern" || left.Kind() == "undefined" {
				childPath = appendBindingPath(path, w.staticPropertySegment(left))
			}
			withDefault := modifiers
			withDefault.hasDefault = true
			withDefault = withBindingConstruct(withDefault, candidate)
			w.walk(left, childPath, withDefault, bindingContainerObject)
		case "rest_pattern":
			w.walkRest(candidate, path, modifiers, bindingContainerObject)
		default:
			w.walk(candidate, path, modifiers, bindingContainerObject)
		}
	}
}

func (w *bindingPatternWalker) walkPair(
	node *sitter.Node,
	path []scopeir.BindingPathSegment,
	modifiers bindingModifiers,
) {
	key := child(node, "key")
	value := child(node, "value")
	if key == nil || value == nil {
		w.diagnose(
			node,
			path,
			scopeir.DiagnosticMalformedBindingNode,
			"object binding pair is missing its key or value",
		)
		return
	}
	if key.HasError() || value.HasError() {
		w.diagnose(
			node,
			path,
			scopeir.DiagnosticUnsupportedBindingNode,
			"object binding pair contains grammar errors",
		)
		return
	}

	segment, ok := w.propertySegment(key)
	if !ok {
		w.diagnose(
			key,
			path,
			scopeir.DiagnosticUnsupportedBindingNode,
			"object binding key cannot be represented deterministically",
		)
		return
	}
	modifiers = withBindingConstruct(modifiers, node)
	w.walk(value, appendBindingPath(path, segment), modifiers, bindingContainerNone)
}

func (w *bindingPatternWalker) walkRest(
	node *sitter.Node,
	path []scopeir.BindingPathSegment,
	modifiers bindingModifiers,
	container bindingContainer,
) {
	if container == bindingContainerNone {
		w.diagnose(
			node,
			path,
			scopeir.DiagnosticUnsupportedBindingNode,
			"rest binding is only legal as a direct array/object pattern child",
		)
		return
	}
	target := firstNamedNonCommentChild(node)
	if target == nil {
		w.diagnose(
			node,
			path,
			scopeir.DiagnosticMalformedBindingNode,
			"rest binding is missing its target",
		)
		return
	}
	switch target.Kind() {
	case "identifier", "undefined":
		modifiers.rest = true
		modifiers = withBindingConstruct(modifiers, node)
		w.walk(target, path, modifiers, bindingContainerNone)
	case "array_pattern", "object_pattern":
		if container != bindingContainerArray {
			w.diagnose(
				target,
				path,
				scopeir.DiagnosticUnsupportedBindingNode,
				"object rest binding requires an identifier target",
			)
			return
		}
		modifiers.rest = true
		modifiers = withBindingConstruct(modifiers, node)
		w.walk(target, path, modifiers, bindingContainerNone)
	default:
		w.diagnose(
			target,
			path,
			scopeir.DiagnosticInvalidRestBinding,
			"rest target is an assignment target, not a declaration binding",
		)
	}
}

func (w *bindingPatternWalker) propertySegment(node *sitter.Node) (scopeir.BindingPathSegment, bool) {
	if node.Kind() == "computed_property_name" {
		expression := firstNamedNonCommentChild(node)
		if expression == nil || expression.IsMissing() || expression.IsError() || node.HasError() || expression.HasError() {
			return scopeir.BindingPathSegment{}, false
		}
		return scopeir.BindingPathSegment{
			Kind:               scopeir.BindingPathComputedProperty,
			ComputedExpression: strings.TrimSpace(expression.Utf8Text(w.request.Source)),
			SourceRange:        nodeRange(node),
		}, true
	}

	switch node.Kind() {
	case "property_identifier", "string", "number":
		return w.staticPropertySegment(node), true
	default:
		return scopeir.BindingPathSegment{}, false
	}
}

func (w *bindingPatternWalker) staticPropertySegment(node *sitter.Node) scopeir.BindingPathSegment {
	name := strings.TrimSpace(node.Utf8Text(w.request.Source))
	if node.Kind() == "string" {
		name = stripQuotes(name)
	}
	return scopeir.BindingPathSegment{
		Kind:         scopeir.BindingPathStaticProperty,
		PropertyName: name,
		SourceRange:  nodeRange(node),
	}
}

func (w *bindingPatternWalker) addLeaf(
	node *sitter.Node,
	path []scopeir.BindingPathSegment,
	modifiers bindingModifiers,
) {
	name := node.Utf8Text(w.request.Source)
	if name == "" {
		w.diagnose(
			node,
			path,
			scopeir.DiagnosticMalformedBindingNode,
			"binding identifier is empty",
		)
		return
	}
	rng := nodeRange(node)
	if modifiers.construct != nil {
		rng = nodeRange(modifiers.construct)
	}
	selectionRange := nodeRange(node)
	w.result.Leaves = append(w.result.Leaves, scopeir.BindingLeafFact{
		FilePath:       w.request.FilePath,
		FileHash:       w.request.FileHash,
		Name:           name,
		Range:          rng,
		SelectionRange: &selectionRange,
		Path:           appendBindingPath(nil, path...),
		Rest:           modifiers.rest,
		Default:        modifiers.hasDefault,
		Provenance:     w.provenance,
	})
}

func withBindingConstruct(modifiers bindingModifiers, node *sitter.Node) bindingModifiers {
	if modifiers.construct == nil {
		modifiers.construct = node
	}
	return modifiers
}

func (w *bindingPatternWalker) diagnose(
	node *sitter.Node,
	path []scopeir.BindingPathSegment,
	code scopeir.ExtractionDiagnosticCode,
	reason string,
) {
	w.result.Diagnostics = append(w.result.Diagnostics, scopeir.ExtractionDiagnosticFact{
		Code:       code,
		FilePath:   w.request.FilePath,
		FileHash:   w.request.FileHash,
		Range:      nodeRange(node),
		NodeKind:   node.Kind(),
		Reason:     reason,
		Path:       appendBindingPath(nil, path...),
		Provenance: w.provenance,
	})
}

func appendBindingPath(
	path []scopeir.BindingPathSegment,
	segments ...scopeir.BindingPathSegment,
) []scopeir.BindingPathSegment {
	result := make([]scopeir.BindingPathSegment, len(path), len(path)+len(segments))
	copy(result, path)
	return append(result, segments...)
}

func firstNamedNonCommentChild(node *sitter.Node) *sitter.Node {
	if node == nil {
		return nil
	}
	for index := uint(0); index < node.NamedChildCount(); index++ {
		candidate := node.NamedChild(index)
		if candidate != nil && candidate.Kind() != "comment" {
			return candidate
		}
	}
	return nil
}
