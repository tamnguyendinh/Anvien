package tsjs

import (
	sitter "github.com/tree-sitter/go-tree-sitter"

	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

func (c *collector) emitReference(node *sitter.Node) {
	c.emitReferenceKind(node, node.Kind())
}

func (c *collector) emitReferenceKind(node *sitter.Node, kind string) {
	switch kind {
	case "for_in_statement":
		c.emitLoopAssignmentWrites(node)
	case "call_expression":
		fn := child(node, "function")
		if fn == nil {
			return
		}
		arity := countArguments(child(node, "arguments"))
		member := unwrapAwaitExpression(fn)
		if member != nil && member.Kind() == "member_expression" {
			property := child(member, "property")
			receiver := child(member, "object")
			if property != nil {
				c.addCall(node, property, receiver, scopeir.CallMember, arity)
			}
			return
		}
		if member != nil && member.Kind() == "identifier" {
			c.addCall(node, member, nil, scopeir.CallFree, arity)
		}
	case "member_expression":
		if isCallFunctionMember(node) || c.isLoopAssignmentMemberTarget(node) {
			return
		}
		property := child(node, "property")
		receiver := child(node, "object")
		if property != nil {
			c.addAccess(node, property, receiver, memberAccessKind(node))
		}
	case "new_expression":
		ctor := child(node, "constructor")
		if ctor != nil {
			c.addCall(node, ctor, nil, scopeir.CallConstructor, countArguments(child(node, "arguments")))
		}
	case "extends_clause", "extends_type_clause":
		for _, target := range heritageTargetNodes(node) {
			c.addHeritage(node, target, scopeir.HeritageExtends)
		}
	case "implements_clause":
		for _, target := range heritageTargetNodes(node) {
			c.addHeritage(target, target, scopeir.HeritageImplements)
		}
	}
}

func (c *collector) emitLoopAssignmentWrites(node *sitter.Node) {
	if child(node, "kind") != nil {
		return
	}
	c.emitAssignmentTargetWrites(child(node, "left"))
}

func (c *collector) emitAssignmentTargetWrites(target *sitter.Node) {
	target = unwrapLoopAssignmentTarget(target)
	if target == nil {
		return
	}
	switch target.Kind() {
	case "identifier", "undefined", "shorthand_property_identifier_pattern":
		c.addAccess(target, target, nil, scopeir.AccessWrite)
	case "member_expression":
		property := child(target, "property")
		if property != nil {
			c.addAccess(target, property, child(target, "object"), scopeir.AccessWrite)
		}
	case "array_pattern", "object_pattern":
		for index := uint(0); index < target.NamedChildCount(); index++ {
			candidate := target.NamedChild(index)
			if candidate != nil && candidate.Kind() != "comment" {
				c.emitAssignmentTargetWrites(candidate)
			}
		}
	case "pair_pattern":
		c.emitAssignmentTargetWrites(child(target, "value"))
	case "assignment_pattern", "object_assignment_pattern":
		c.emitAssignmentTargetWrites(child(target, "left"))
	case "rest_pattern":
		c.emitAssignmentTargetWrites(firstNamedNonCommentChild(target))
	}
}

func unwrapLoopAssignmentTarget(node *sitter.Node) *sitter.Node {
	for node != nil {
		switch node.Kind() {
		case "parenthesized_expression", "non_null_expression":
			next := node.NamedChild(0)
			if next == nil {
				return node
			}
			node = next
		default:
			return node
		}
	}
	return nil
}

func (c *collector) isLoopAssignmentMemberTarget(node *sitter.Node) bool {
	for current := node.Parent(); current != nil; current = current.Parent() {
		if current.Kind() != "for_in_statement" {
			continue
		}
		if child(current, "kind") != nil {
			return false
		}
		return assignmentTargetContainsMember(child(current, "left"), node)
	}
	return false
}

func assignmentTargetContainsMember(target *sitter.Node, member *sitter.Node) bool {
	target = unwrapLoopAssignmentTarget(target)
	if target == nil || member == nil {
		return false
	}
	switch target.Kind() {
	case "member_expression":
		return target.Id() == member.Id()
	case "array_pattern", "object_pattern":
		for index := uint(0); index < target.NamedChildCount(); index++ {
			candidate := target.NamedChild(index)
			if candidate != nil && candidate.Kind() != "comment" && assignmentTargetContainsMember(candidate, member) {
				return true
			}
		}
	case "pair_pattern":
		return assignmentTargetContainsMember(child(target, "value"), member)
	case "assignment_pattern", "object_assignment_pattern":
		return assignmentTargetContainsMember(child(target, "left"), member)
	case "rest_pattern":
		return assignmentTargetContainsMember(firstNamedNonCommentChild(target), member)
	}
	return false
}

func (c *collector) addCall(anchor *sitter.Node, nameNode *sitter.Node, receiver *sitter.Node, form scopeir.CallForm, arity *int) {
	rng := nodeRange(anchor)
	fact := scopeir.CallSiteFact{
		FilePath: c.filePath,
		FileHash: c.fileHash,
		Name:     c.text(nameNode),
		Range:    rng,
		InScope:  c.innermostScopeID(rng),
		CallForm: form,
		Arity:    arity,
	}
	if receiver != nil {
		fact.ExplicitReceiver = c.text(receiver)
	}
	c.calls = append(c.calls, fact)
}

func (c *collector) addAccess(anchor *sitter.Node, nameNode *sitter.Node, receiver *sitter.Node, kind scopeir.AccessKind) {
	rng := nodeRange(anchor)
	fact := scopeir.AccessFact{
		FilePath: c.filePath,
		FileHash: c.fileHash,
		Name:     c.text(nameNode),
		Kind:     kind,
		Range:    rng,
		InScope:  c.innermostScopeID(rng),
	}
	if receiver != nil {
		fact.ExplicitReceiver = c.text(receiver)
	}
	c.accesses = append(c.accesses, fact)
}

func (c *collector) addHeritage(anchor *sitter.Node, nameNode *sitter.Node, kind scopeir.HeritageKind) {
	rng := nodeRange(anchor)
	c.heritage = append(c.heritage, scopeir.HeritageFact{
		FilePath: c.filePath,
		FileHash: c.fileHash,
		Name:     c.text(nameNode),
		Kind:     kind,
		Range:    rng,
		InScope:  c.innermostScopeID(rng),
	})
}

func isCallFunctionMember(node *sitter.Node) bool {
	parent := node.Parent()
	if parent == nil || parent.Kind() != "call_expression" {
		return false
	}
	fn := child(parent, "function")
	return fn != nil && fn.Id() == node.Id()
}

func memberAccessKind(node *sitter.Node) scopeir.AccessKind {
	parent := node.Parent()
	if parent == nil {
		return scopeir.AccessRead
	}
	switch parent.Kind() {
	case "assignment_expression", "augmented_assignment_expression":
		left := child(parent, "left")
		if left == nil {
			left = parent.NamedChild(0)
		}
		if left != nil && left.Id() == node.Id() {
			return scopeir.AccessWrite
		}
	case "update_expression":
		return scopeir.AccessWrite
	}
	return scopeir.AccessRead
}

func heritageTargetNodes(node *sitter.Node) []*sitter.Node {
	if node == nil {
		return nil
	}
	if value := child(node, "value"); value != nil {
		return []*sitter.Node{value}
	}

	out := make([]*sitter.Node, 0, node.NamedChildCount())
	for index := uint(0); index < node.NamedChildCount(); index++ {
		candidate := node.NamedChild(index)
		if candidate == nil {
			continue
		}
		switch candidate.Kind() {
		case "comment", "type_arguments":
			continue
		}
		out = append(out, candidate)
	}
	if len(out) == 0 {
		if value := firstIdentifierLikeChild(node); value != nil {
			out = append(out, value)
		}
	}
	return out
}
