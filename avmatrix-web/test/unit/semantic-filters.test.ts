import { describe, expect, it } from "vitest";
import type { GraphNode } from "../../src/generated/avmatrix-contracts";
import {
  DEFAULT_SEMANTIC_FILTERS,
  getAppLayerCounts,
  getResolutionGapActionabilityCounts,
  getResolutionGapTargetTextCounts,
  getResolutionHealthBucketCounts,
  getResolutionLensRows,
  getSemanticFilterableFromNode,
  semanticMatchesFilters,
} from "../../src/lib/semantic-filters";

const makeNode = (
  id: string,
  properties: Partial<GraphNode["properties"]>,
  label: GraphNode["label"] = "Function",
): GraphNode =>
  ({
    id,
    label,
    properties: {
      name: id,
      filePath: "src/app.ts",
      ...properties,
    },
  }) as GraphNode;

describe("semantic filters", () => {
  it("counts App Layer and ResolutionGap metadata without inventing missing fields", () => {
    const nodes = [
      makeNode("backend", { appLayer: "backend" }),
      makeNode("missing", {}),
      makeNode(
        "gap",
        {
          appLayer: "api",
          sourceAppLayer: "api",
          functionalArea: "api",
          factFamily: "call",
          targetRole: "callable",
          classification: "in_repo_unresolved",
          actionability: "analyzer_gap",
          targetText: "missingHandler",
          resolutionHealthBuckets: {
            in_repo_analyzer_gap: 2,
            unresolved_call_target: 1,
          },
          resolutionConfidence: "degraded",
        },
        "ResolutionGap",
      ),
    ];

    expect(getAppLayerCounts(nodes).get("backend")).toBe(1);
    expect(getAppLayerCounts(nodes).get("missing")).toBe(1);
    expect(getResolutionHealthBucketCounts(nodes).get("in_repo_analyzer_gap")).toBe(2);
    expect(getResolutionGapActionabilityCounts(nodes).get("analyzer_gap")).toBe(1);
    expect(getResolutionGapTargetTextCounts(nodes).get("missingHandler")).toBe(1);

    const missingFilterable = getSemanticFilterableFromNode(nodes[1]);
    expect(missingFilterable.appLayer).toBeUndefined();
  });

  it("matches App Layer, missing metadata, resolution bucket, and target text filters", () => {
    const backend = getSemanticFilterableFromNode(
      makeNode("backend", { appLayer: "backend", resolutionConfidence: "clear" }),
    );
    const missing = getSemanticFilterableFromNode(makeNode("missing", {}));
    const gap = getSemanticFilterableFromNode(
      makeNode(
        "gap",
        {
          appLayer: "api",
          sourceAppLayer: "api",
          factFamily: "call",
          targetRole: "callable",
          actionability: "analyzer_gap",
          targetText: "missingHandler",
          resolutionHealthBuckets: { in_repo_analyzer_gap: 1 },
        },
        "ResolutionGap",
      ),
    );

    expect(
      semanticMatchesFilters(backend, {
        ...DEFAULT_SEMANTIC_FILTERS,
        visibleAppLayers: ["backend"],
      }),
    ).toBe(true);
    expect(
      semanticMatchesFilters(gap, {
        ...DEFAULT_SEMANTIC_FILTERS,
        visibleAppLayers: ["backend"],
      }),
    ).toBe(false);
    expect(
      semanticMatchesFilters(missing, {
        ...DEFAULT_SEMANTIC_FILTERS,
        showNodesMissingAppLayer: false,
      }),
    ).toBe(false);
    expect(
      semanticMatchesFilters(gap, {
        ...DEFAULT_SEMANTIC_FILTERS,
        visibleResolutionHealthBuckets: ["external_unresolved"],
      }),
    ).toBe(false);
    expect(
      semanticMatchesFilters(gap, {
        ...DEFAULT_SEMANTIC_FILTERS,
        visibleResolutionGapTargetTexts: ["missingHandler"],
      }),
    ).toBe(true);
  });

  it("builds the required resolution lens rows from persisted gap fields", () => {
    const rows = getResolutionLensRows([
      makeNode(
        "api-gap",
        {
          appLayer: "api",
          sourceAppLayer: "api",
          functionalArea: "api",
          factFamily: "call",
          targetRole: "callable",
          classification: "in_repo_unresolved",
          actionability: "analyzer_gap",
          targetText: "missingHandler",
        },
        "ResolutionGap",
      ),
      makeNode(
        "frontend-type-gap",
        {
          appLayer: "frontend",
          sourceAppLayer: "frontend",
          functionalArea: "web_graph_ui",
          factFamily: "type-reference",
          targetRole: "type",
          classification: "in_repo_unresolved",
          actionability: "analyzer_gap",
          targetText: "GraphNode",
        },
        "ResolutionGap",
      ),
    ]);

    expect(rows.map((row) => row.id)).toEqual([
      "backend-unresolved-calls",
      "api-unresolved-handlers-contracts",
      "frontend-unresolved-type-refs",
      "shared-contract-analyzer-gaps",
      "external-unresolved-symbols",
      "builtin-test-stdlib-non-actionable",
      "in-repo-analyzer-gaps",
      "resolution-gaps-by-functional-area",
      "top-app-layers-by-analyzer-gap",
      "top-functional-areas-by-unresolved",
      "top-unresolved-target-text",
    ]);
    expect(rows.find((row) => row.id === "api-unresolved-handlers-contracts")?.count).toBe(1);
    expect(rows.find((row) => row.id === "frontend-unresolved-type-refs")?.count).toBe(1);
    expect(rows.find((row) => row.id === "in-repo-analyzer-gaps")?.count).toBe(2);
  });
});
