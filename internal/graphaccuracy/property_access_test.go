package graphaccuracy

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunPropertyAccessAuditClassifiesCrossLanguagePropertyTruth(t *testing.T) {
	repo := t.TempDir()
	writeAccuracyFixtureFile(t, repo, "src/model.ts", `interface User {
  id: string
}

type Shape = {
  title: string
}

const runtime = {
  ok: true
}

class Box {
  result: string
}
`)
	writeAccuracyFixtureFile(t, repo, "service/service.go", `package service

type Service struct {
	Repo string
}
`)
	graphPath := writeGraphFixture(t, repo, "graph.json", GraphFile{
		Nodes: []GraphNode{
			propertyNode("src/model.ts", "typescript", "id", "id", "string", 2),
			propertyNode("src/model.ts", "typescript", "title", "title", "string", 6),
			propertyNode("src/model.ts", "typescript", "ok", "ok", "boolean", 10),
			propertyNode("src/model.ts", "typescript", "Box.result", "result", "string", 14),
			propertyNode("service/service.go", "go", "Service.Repo", "Repo", "string", 4),
			{ID: "Class:src/model.ts:Box", Label: "Class", Properties: map[string]any{"name": "Box", "filePath": "src/model.ts", "language": "typescript"}},
			{ID: "Struct:service/service.go:Service", Label: "Struct", Properties: map[string]any{"name": "Service", "filePath": "service/service.go", "language": "go"}},
			{ID: "Function:src/model.ts:read", Label: "Function", Properties: map[string]any{"name": "read", "filePath": "src/model.ts", "language": "typescript"}},
		},
		Relationships: []GraphRelationship{
			{ID: "has:box-result", Type: "HAS_PROPERTY", SourceID: "Class:src/model.ts:Box", TargetID: "Property:src/model.ts:Box.result"},
			{ID: "has:service-repo", Type: "HAS_PROPERTY", SourceID: "Struct:service/service.go:Service", TargetID: "Property:service/service.go:Service.Repo"},
			{ID: "access:read-result", Type: "ACCESSES", SourceID: "Function:src/model.ts:read", TargetID: "Property:src/model.ts:Box.result"},
		},
	})
	outPath := filepath.Join(repo, "audit.json")

	result, err := RunPropertyAccessAudit(PropertyAccessAuditOptions{
		Repo:        repo,
		GraphPath:   graphPath,
		OutPath:     outPath,
		MaxExamples: 2,
	})
	if err != nil {
		t.Fatalf("RunPropertyAccessAudit() error = %v", err)
	}
	if _, err := os.Stat(outPath); err != nil {
		t.Fatalf("audit output missing: %v", err)
	}
	if result.Totals.PropertyNodes != 5 || result.Totals.OwnerLinkedProperties != 2 || result.Totals.HasPropertyEdges != 2 || result.Totals.AccessesEdges != 1 {
		t.Fatalf("totals = %#v", result.Totals)
	}
	if result.Languages["typescript"].PropertyNodes != 4 || result.Languages["go"].PropertyNodes != 1 {
		t.Fatalf("languages = %#v", result.Languages)
	}
	assertBucketCount(t, result.Categories, "tsjs_interface_property_signature", 1)
	assertBucketCount(t, result.Categories, "tsjs_type_alias_object_literal_member", 1)
	assertBucketCount(t, result.Categories, "tsjs_runtime_object_literal_key", 1)
	assertBucketCount(t, result.Categories, "tsjs_class_field", 1)
	assertBucketCount(t, result.Categories, "go_owner_linked_struct", 1)
	assertBucketCount(t, result.OrphanStatus, "false_orphan", 2)
	assertBucketCount(t, result.OrphanStatus, "true_orphan", 1)
	assertBucketCount(t, result.OrphanStatus, "owner_linked", 2)
	assertBucketCount(t, result.GraphTruth, "real_edge_missing", 2)
	assertBucketCount(t, result.GraphTruth, "true_no_edge", 1)
	assertBucketCount(t, result.GraphTruth, "edge_present", 2)
}

func TestRunPropertyAccessAuditClassifiesGoAnonymousStructFieldsAsTrueNoEdge(t *testing.T) {
	repo := t.TempDir()
	writeAccuracyFixtureFile(t, repo, "service/anonymous_test.go", `package service

func TestParseAction() {
	tests := []struct {
		args []string
		want string
	}{}
	_ = tests
}
`)
	graphPath := writeGraphFixture(t, repo, "graph.json", GraphFile{
		Nodes: []GraphNode{
			propertyNode("service/anonymous_test.go", "go", "args", "args", "[]string", 5),
			propertyNode("service/anonymous_test.go", "go", "want", "want", "string", 6),
		},
	})

	result, err := RunPropertyAccessAudit(PropertyAccessAuditOptions{
		Repo:        repo,
		GraphPath:   graphPath,
		MaxExamples: 2,
	})
	if err != nil {
		t.Fatalf("RunPropertyAccessAudit() error = %v", err)
	}
	assertBucketCount(t, result.Categories, "go_anonymous_struct_field", 2)
	assertBucketCount(t, result.OrphanStatus, "true_orphan", 2)
	assertBucketCount(t, result.GraphTruth, "true_no_edge", 2)
}

func TestRunPropertyAccessAuditClassifiesTSJSInlineTypeLiteralAsTrueNoEdge(t *testing.T) {
	repo := t.TempDir()
	writeAccuracyFixtureFile(t, repo, "src/panel.tsx", `import { useRef, useState } from "react";

interface PanelProps {
  title: string;
}

export function Panel() {
  const resizeRef = useRef<{ startX: number; startWidth: number } | null>(null);
  const [error] = useState<{
    message: string;
  } | null>(null);
  return resizeRef.current?.startX ?? error?.message;
}
`)
	graphPath := writeGraphFixture(t, repo, "graph.json", GraphFile{
		Nodes: []GraphNode{
			propertyNode("src/panel.tsx", "typescript", "startX", "startX", "number", 8),
			propertyNode("src/panel.tsx", "typescript", "startWidth", "startWidth", "number", 8),
			propertyNode("src/panel.tsx", "typescript", "message", "message", "string", 10),
		},
	})

	result, err := RunPropertyAccessAudit(PropertyAccessAuditOptions{
		Repo:        repo,
		GraphPath:   graphPath,
		MaxExamples: 3,
	})
	if err != nil {
		t.Fatalf("RunPropertyAccessAudit() error = %v", err)
	}
	assertBucketCount(t, result.Categories, "tsjs_inline_type_literal_property", 3)
	assertBucketCount(t, result.OrphanStatus, "true_orphan", 3)
	assertBucketCount(t, result.GraphTruth, "true_no_edge", 3)
}

func propertyNode(rel string, language string, qualifiedName string, name string, declaredType string, startLine int) GraphNode {
	return GraphNode{
		ID:    "Property:" + rel + ":" + qualifiedName,
		Label: "Property",
		Properties: map[string]any{
			"name":          name,
			"qualifiedName": qualifiedName,
			"declaredType":  declaredType,
			"filePath":      rel,
			"language":      language,
			"startLine":     startLine,
		},
	}
}

func assertBucketCount(t *testing.T, buckets map[string]PropertyAccessBucket, key string, want int) {
	t.Helper()
	if got := buckets[key].Count; got != want {
		t.Fatalf("bucket %s = %d, want %d (%#v)", key, got, want, buckets)
	}
}
