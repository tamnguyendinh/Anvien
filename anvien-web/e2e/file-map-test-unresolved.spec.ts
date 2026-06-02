import { expect, test } from "@playwright/test";

const BACKEND_URL = process.env.BACKEND_URL ?? "http://127.0.0.1:4848";
const FRONTEND_URL = process.env.FRONTEND_URL ?? "http://127.0.0.1:5228";

const graph = {
  nodes: [
    {
      id: "File:src/app.ts",
      label: "File",
      properties: {
        name: "app.ts",
        filePath: "src/app.ts",
        appLayer: "api",
        functionalArea: "mcp",
      },
    },
    {
      id: "File:src/app.test.ts",
      label: "File",
      properties: {
        name: "app.test.ts",
        filePath: "src/app.test.ts",
        appLayer: "api_test",
        functionalArea: "mcp",
      },
    },
    {
      id: "Function:src/app.ts:main",
      label: "Function",
      properties: {
        name: "main",
        filePath: "src/app.ts",
        startLine: 3,
        endLine: 12,
      },
    },
    {
      id: "Function:src/app.test.ts:testMain",
      label: "Function",
      properties: {
        name: "testMain",
        filePath: "src/app.test.ts",
        startLine: 9,
        endLine: 14,
      },
    },
  ],
  relationships: [
    {
      id: "rel-app-defines-main",
      sourceId: "File:src/app.ts",
      targetId: "Function:src/app.ts:main",
      type: "DEFINES",
      confidence: 1,
      reason: "fixture",
    },
    {
      id: "rel-test-defines-test-main",
      sourceId: "File:src/app.test.ts",
      targetId: "Function:src/app.test.ts:testMain",
      type: "DEFINES",
      confidence: 1,
      reason: "fixture",
    },
    {
      id: "rel-test-calls-main",
      sourceId: "Function:src/app.test.ts:testMain",
      targetId: "Function:src/app.ts:main",
      type: "CALLS",
      confidence: 1,
      reason: "fixture",
    },
  ],
};

const sourceFileSummary = {
  path: "src/app.ts",
  language: "typescript",
  kind: "source",
  appLayer: "api",
  functionalArea: "mcp",
  parseStatus: "parsed",
  symbolCount: 2,
  exportedSymbolCount: 1,
  inboundRefCount: 1,
  outboundRefCount: 0,
  localRelationshipCount: 1,
  unresolvedSourceSiteCount: 1,
  rawUnresolvedSourceSiteCount: 1,
  productionUnresolvedSourceSiteCount: 1,
  testUnresolvedSourceSiteCount: 0,
  nonActionableUnresolvedSourceSiteCount: 0,
  unknownUnresolvedSourceSiteCount: 0,
  defaultVisibleUnresolvedSourceSiteCount: 1,
  linkedFlowCount: 0,
  linkedTestCount: 1,
  risk: "medium",
  rawRisk: "medium",
  defaultVisibleRisk: "medium",
  stale: false,
  changedSinceAnalyze: false,
};

const testFileSummary = {
  path: "src/app.test.ts",
  language: "typescript",
  kind: "test",
  appLayer: "api_test",
  functionalArea: "mcp",
  parseStatus: "parsed",
  symbolCount: 1,
  exportedSymbolCount: 0,
  inboundRefCount: 0,
  outboundRefCount: 1,
  localRelationshipCount: 1,
  unresolvedSourceSiteCount: 3,
  rawUnresolvedSourceSiteCount: 3,
  productionUnresolvedSourceSiteCount: 0,
  testUnresolvedSourceSiteCount: 3,
  nonActionableUnresolvedSourceSiteCount: 0,
  unknownUnresolvedSourceSiteCount: 0,
  defaultVisibleUnresolvedSourceSiteCount: 0,
  linkedFlowCount: 0,
  linkedTestCount: 0,
  risk: "low",
  rawRisk: "high",
  defaultVisibleRisk: "low",
  stale: false,
  changedSinceAnalyze: false,
};

test.describe("file map test unresolved defaults", () => {
  test.beforeEach(async ({ page }) => {
    await page.route(`${BACKEND_URL}/api/repos`, (route) =>
      route.fulfill({
        json: [{ name: "test-files-demo", path: "/tmp/test-files-demo" }],
      }),
    );
    await page.route(`${BACKEND_URL}/api/repo**`, (route) =>
      route.fulfill({
        json: {
          name: "test-files-demo",
          path: "/tmp/test-files-demo",
          repoPath: "/tmp/test-files-demo",
        },
      }),
    );
    await page.route(`${BACKEND_URL}/api/graph**`, (route) =>
      route.fulfill({ json: graph }),
    );
    await page.route(`${BACKEND_URL}/api/file-hotspots**`, (route) =>
      route.fulfill({
        json: {
          repo: "test-files-demo",
          repoPath: "/tmp/test-files-demo",
          graph: { path: ".anvien/graph.json", stale: false },
          total: 2,
          offset: 0,
          limit: 200,
          sort: "unresolved",
          files: [sourceFileSummary, testFileSummary],
        },
      }),
    );
    await page.route(`${BACKEND_URL}/api/file-context**`, (route) =>
      route.fulfill({
        json: {
          repo: "test-files-demo",
          repoPath: "/tmp/test-files-demo",
          graph: { path: ".anvien/graph.json", stale: false },
          target: {
            type: "file",
            input: "src/app.test.ts",
            normalizedPath: "src/app.test.ts",
            dispatchMode: "explicit",
            ambiguityCandidates: [],
          },
          summary: testFileSummary,
          symbolTree: [
            {
              id: "Function:src/app.test.ts:testMain",
              name: "testMain",
              kind: "function",
              range: { startLine: 9, endLine: 14 },
              exported: false,
              relationshipCounts: {
                local: 1,
                inbound: 0,
                outbound: 1,
                unresolved: 3,
              },
              children: [],
            },
          ],
          relationships: {
            counts: {
              local: 1,
              outbound: 1,
              inbound: 0,
              samplesReturned: 2,
            },
            local: {
              total: 1,
              counts: { DEFINES: 1 },
              samples: [
                {
                  sourceFile: "src/app.test.ts",
                  sourceSymbol: "File:src/app.test.ts",
                  sourceRange: {},
                  relationshipKind: "DEFINES",
                  targetFile: "src/app.test.ts",
                  targetSymbol: "testMain",
                  targetRange: { startLine: 9 },
                },
              ],
            },
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
            counts: {
              flows: 0,
              routes: 0,
              mcpTools: 0,
              tests: 0,
            },
            flows: [],
            routes: [],
            mcpTools: [],
            tests: [],
          },
          quality: {
            parser: "parsed",
            resolutionConfidence: "degraded",
            unresolvedCalls: 3,
            unresolvedRefs: 0,
            unresolvedImports: 0,
            generated: false,
            stale: false,
            changedSinceAnalyze: false,
          },
          limits: {
            relationshipSamplesPerGroup: 5,
            unresolvedSamplesPerGroup: 5,
            linkedSamplesPerKind: 5,
          },
        },
      }),
    );
    await page.route(`${BACKEND_URL}/api/heartbeat`, (route) =>
      route.fulfill({
        status: 200,
        headers: {
          "Content-Type": "text/event-stream",
          "Cache-Control": "no-cache",
        },
        body: ":ok\n\n",
      }),
    );
  });

  test("shows test file identity while hiding raw test unresolved samples by default", async ({
    page,
  }) => {
    await page.goto(
      `${FRONTEND_URL}/?server=${encodeURIComponent(BACKEND_URL)}&project=test-files-demo`,
    );

    await expect(page.locator('[data-testid="status-ready"]')).toBeVisible({
      timeout: 20_000,
    });

    await page.getByTestId("file-map-tab").click();
    await expect(page.getByTestId("file-map-panel")).toBeVisible();
    await expect(page.getByText("1 unresolved")).toBeVisible();
    await expect(page.getByText("Test File")).toBeVisible();

    await page
      .getByTestId("file-map-row")
      .filter({ hasText: "src/app.test.ts" })
      .locator("button")
      .first()
      .click();
    await expect(page.getByTestId("file-detail-section-summary")).toContainText("Test File");
    await expect(page.getByTestId("file-detail-section-unresolved")).toContainText("0 sites");
    await expect(page.getByTestId("file-detail-section-unresolved")).not.toContainText(
      "expectSomething",
    );
    await expect(page.getByTestId("file-detail-section-relationships")).toContainText("src/app.ts");
  });
});
