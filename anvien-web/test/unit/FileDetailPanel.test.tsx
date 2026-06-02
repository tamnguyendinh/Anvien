import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { beforeEach, describe, expect, it, vi } from "vitest";
import type { FileContextResponse } from "../../src/generated/anvien-contracts";

const fetchFileContext = vi.fn();

vi.mock("../../src/services/backend-client", () => ({
  fetchFileContext: (...args: unknown[]) => fetchFileContext(...args),
}));

import { FileDetailPanel } from "../../src/components/FileDetailPanel";

const fileContext: FileContextResponse = {
  repo: "demo",
  repoPath: "F:/demo",
  graph: {
    path: ".anvien/graph.json",
    stale: false,
  },
  target: {
    type: "file",
    input: "src/app.ts",
    normalizedPath: "src/app.ts",
    dispatchMode: "explicit",
    ambiguityCandidates: [],
  },
  summary: {
    path: "src/app.ts",
    language: "typescript",
    kind: "source",
    appLayer: "api",
    functionalArea: "mcp",
    parseStatus: "parsed",
    symbolCount: 2,
    exportedSymbolCount: 1,
    inboundRefCount: 4,
    outboundRefCount: 7,
    localRelationshipCount: 3,
    unresolvedSourceSiteCount: 1,
    rawUnresolvedSourceSiteCount: 1,
    productionUnresolvedSourceSiteCount: 1,
    testUnresolvedSourceSiteCount: 0,
    nonActionableUnresolvedSourceSiteCount: 0,
    unknownUnresolvedSourceSiteCount: 0,
    defaultVisibleUnresolvedSourceSiteCount: 1,
    linkedFlowCount: 1,
    linkedTestCount: 1,
    risk: "medium",
    rawRisk: "medium",
    defaultVisibleRisk: "medium",
    stale: false,
    changedSinceAnalyze: true,
  },
  symbolTree: [
    {
      id: "Function:src/app.ts:main",
      name: "main",
      kind: "function",
      range: { startLine: 3, endLine: 12 },
      exported: true,
      signature: "function main()",
      relationshipCounts: {
        local: 1,
        inbound: 2,
        outbound: 3,
        unresolved: 1,
      },
      children: [
        {
          id: "Function:src/app.ts:helper",
          name: "helper",
          kind: "function",
          range: { startLine: 6, endLine: 8 },
          exported: false,
          relationshipCounts: {
            local: 1,
            inbound: 0,
            outbound: 1,
            unresolved: 0,
          },
          children: [],
        },
      ],
    },
  ],
  relationships: {
    counts: {
      local: 1,
      outbound: 1,
      inbound: 1,
      samplesReturned: 3,
    },
    local: {
      total: 1,
      counts: { CALLS: 1 },
      samples: [
        {
          sourceFile: "src/app.ts",
          sourceSymbol: "main",
          sourceRange: { startLine: 3 },
          relationshipKind: "CALLS",
          targetFile: "src/app.ts",
          targetSymbol: "helper",
          targetRange: { startLine: 6 },
          sourceSiteId: "site-local",
          proofKind: "source-site",
          sourceSiteStatus: "resolved",
        },
      ],
    },
    outboundByFile: [
      {
        file: "src/router.ts",
        total: 1,
        counts: { USES: 1 },
        samples: [
          {
            sourceFile: "src/app.ts",
            sourceSymbol: "main",
            sourceRange: { startLine: 4 },
            relationshipKind: "USES",
            targetFile: "src/router.ts",
            targetSymbol: "createRouter",
            targetRange: { startLine: 1 },
            sourceSiteId: "site-out",
            proofKind: "source-site",
            sourceSiteStatus: "resolved",
          },
        ],
      },
    ],
    inboundByFile: [
      {
        file: "src/app.test.ts",
        total: 1,
        counts: { CALLS: 1 },
        samples: [
          {
            sourceFile: "src/app.test.ts",
            sourceSymbol: "testMain",
            sourceRange: { startLine: 9 },
            relationshipKind: "CALLS",
            targetFile: "src/app.ts",
            targetSymbol: "main",
            targetRange: { startLine: 3 },
            sourceSiteId: "site-in",
            proofKind: "source-site",
            sourceSiteStatus: "resolved",
          },
        ],
      },
    ],
  },
  unresolved: {
    total: 1,
    byKind: { unresolved_call: 1 },
    byClassification: { in_repo_unresolved: 1 },
    byActionability: { analyzer_gap: 1 },
    groups: [
      {
        sourceSymbol: "main",
        total: 1,
        samples: [
          {
            line: 11,
            column: 4,
            targetText: "missingCall",
            sourceSymbol: "main",
            gapKind: "unresolved_call",
            classification: "in_repo_unresolved",
            actionability: "analyzer_gap",
            proofKind: "none",
            sourceSiteId: "site-gap",
            sourceSiteStatus: "unresolved_local_binding",
          },
        ],
      },
    ],
  },
  linked: {
    counts: {
      flows: 1,
      routes: 1,
      mcpTools: 1,
      tests: 1,
    },
    flows: [{ name: "MCP initialize", kind: "flow", source: "process", confidence: "high" }],
    routes: [{ name: "GET /api/app", kind: "route", source: "route-map" }],
    mcpTools: [{ name: "context", kind: "tool", source: "tool-map" }],
    tests: [{ name: "src/app.test.ts", kind: "test", source: "relationship" }],
  },
  quality: {
    parser: "parsed",
    resolutionConfidence: "degraded",
    unresolvedCalls: 1,
    unresolvedRefs: 0,
    unresolvedImports: 0,
    generated: false,
    stale: false,
    changedSinceAnalyze: true,
  },
  limits: {
    relationshipSamplesPerGroup: 5,
    unresolvedSamplesPerGroup: 5,
    linkedSamplesPerKind: 5,
  },
};

const testFileContext: FileContextResponse = {
  ...fileContext,
  target: {
    ...fileContext.target,
    input: "src/app.test.ts",
    normalizedPath: "src/app.test.ts",
  },
  summary: {
    ...fileContext.summary,
    path: "src/app.test.ts",
    kind: "test",
    appLayer: "api_test",
    unresolvedSourceSiteCount: 3,
    rawUnresolvedSourceSiteCount: 3,
    productionUnresolvedSourceSiteCount: 0,
    testUnresolvedSourceSiteCount: 3,
    nonActionableUnresolvedSourceSiteCount: 0,
    unknownUnresolvedSourceSiteCount: 0,
    defaultVisibleUnresolvedSourceSiteCount: 0,
    linkedTestCount: 0,
    risk: "low",
    rawRisk: "high",
    defaultVisibleRisk: "low",
  },
  relationships: {
    ...fileContext.relationships,
    outboundByFile: [
      {
        file: "src/app.ts",
        total: 1,
        counts: { CALLS: 1 },
        samples: [
          {
            sourceFile: "src/app.test.ts",
            sourceSymbol: "testMain",
            sourceRange: { startLine: 9 },
            relationshipKind: "CALLS",
            targetFile: "src/app.ts",
            targetSymbol: "main",
            targetRange: { startLine: 3 },
            sourceSiteId: "site-tested-target",
            proofKind: "source-site",
            sourceSiteStatus: "resolved",
          },
        ],
      },
    ],
    inboundByFile: [],
  },
  unresolved: {
    total: 3,
    byKind: { unresolved_call: 3 },
    byClassification: { test_framework: 3 },
    byActionability: { non_actionable: 3 },
    groups: [
      {
        sourceSymbol: "testMain",
        total: 3,
        samples: [
          {
            line: 11,
            column: 4,
            targetText: "expectSomething",
            sourceSymbol: "testMain",
            gapKind: "unresolved_call",
            classification: "test_framework",
            actionability: "non_actionable",
            proofKind: "none",
            sourceSiteId: "site-test-gap",
            sourceSiteStatus: "unresolved_local_binding",
          },
        ],
      },
    ],
  },
  linked: {
    ...fileContext.linked,
    counts: {
      ...fileContext.linked.counts,
      tests: 0,
    },
    tests: [],
  },
  quality: {
    ...fileContext.quality,
    unresolvedCalls: 3,
    unresolvedRefs: 0,
    unresolvedImports: 0,
  },
};

describe("FileDetailPanel", () => {
  beforeEach(() => {
    fetchFileContext.mockReset();
    fetchFileContext.mockResolvedValue(fileContext);
  });

  it("loads file context and renders all major detail sections", async () => {
    const onFocusNode = vi.fn();

    render(
      <FileDetailPanel
        repoName="demo"
        filePath="src/app.ts"
        onFocusNode={onFocusNode}
      />,
    );

    expect(await screen.findByText("src/app.ts")).toBeInTheDocument();

    await waitFor(() => {
      expect(fetchFileContext).toHaveBeenCalledWith("src/app.ts", {
        repo: "demo",
        relationships: 5,
        unresolved: 5,
        linked: 5,
      });
    });

    expect(screen.getByTestId("file-detail-section-summary")).toHaveTextContent("Symbols");
    expect(screen.getByTestId("file-detail-section-quality")).toHaveTextContent("Resolution");
    expect(screen.getByTestId("file-detail-section-symbol-tree")).toHaveTextContent("main");
    expect(screen.getByTestId("file-detail-section-symbol-tree")).toHaveTextContent("helper");
    expect(screen.getByTestId("file-detail-section-relationships")).toHaveTextContent("src/router.ts");
    expect(screen.getByTestId("file-detail-section-unresolved")).toHaveTextContent("missingCall");
    expect(screen.getByTestId("file-detail-section-linked")).toHaveTextContent("MCP initialize");
    expect(screen.getByTestId("file-detail-section-linked")).toHaveTextContent("GET /api/app");

    await userEvent.click(screen.getAllByTestId("file-detail-focus-symbol")[0]);
    expect(onFocusNode).toHaveBeenCalledWith("Function:src/app.ts:main");
  });

  it("renders backend errors without keeping stale detail", async () => {
    fetchFileContext.mockRejectedValue(new Error("file context missing"));

    render(<FileDetailPanel repoName="demo" filePath="src/missing.ts" />);

    expect(await screen.findByText("file context missing")).toBeInTheDocument();
  });

  it("renders test file identity without default unresolved samples", async () => {
    fetchFileContext.mockResolvedValueOnce(testFileContext);

    render(<FileDetailPanel repoName="demo" filePath="src/app.test.ts" />);

    expect(await screen.findByText("src/app.test.ts")).toBeInTheDocument();
    expect(screen.getByText("Test File")).toBeInTheDocument();

    const unresolved = screen.getByTestId("file-detail-section-unresolved");
    expect(unresolved).toHaveTextContent("0 sites");
    expect(unresolved).toHaveTextContent("No default unresolved source-site samples.");
    expect(unresolved).not.toHaveTextContent("expectSomething");
    expect(unresolved).not.toHaveTextContent("site-test-gap");
    expect(screen.getByTestId("file-detail-section-relationships")).toHaveTextContent("src/app.ts");
    expect(screen.queryByTestId("file-detail-toggle-raw-unresolved")).not.toBeInTheDocument();
    expect(screen.getByTestId("file-detail-section-quality")).not.toHaveTextContent("Calls");
  });
});
