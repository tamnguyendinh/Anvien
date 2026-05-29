import { test, expect } from "@playwright/test";

const BACKEND_URL = process.env.BACKEND_URL ?? "http://127.0.0.1:4848";
const FRONTEND_URL = process.env.FRONTEND_URL ?? "http://127.0.0.1:5228";

test.describe("Graph Health UI", () => {
  test.beforeEach(async ({ page }) => {
    const graph = {
      nodes: [
        {
          id: "file-src-widget",
          label: "File",
          properties: {
            name: "widget.ts",
            filePath: "src/widget.ts",
            graphHealth: {
              topologyStatus: "connected",
              countedIncoming: 1,
              countedOutgoing: 1,
              componentReachableFromRoot: true,
              confidence: "candidate",
            },
          },
        },
        {
          id: "fn-detached-widget",
          label: "Function",
          properties: {
            name: "detachedWidget",
            filePath: "src/widget.ts",
            startLine: 0,
            endLine: 2,
            graphHealth: {
              topologyStatus: "detached_component",
              countedIncoming: 0,
              countedOutgoing: 1,
              excludedEdgeCounts: { structural: 1 },
              componentId: "component-detached",
              componentSize: 2,
              componentReachableFromRoot: false,
              diagnostics: [
                {
                  kind: "unresolved_reference",
                  targetText: "missingCall",
                  count: 1,
                },
              ],
              confidence: "unknown",
            },
          },
        },
        {
          id: "fn-detached-peer",
          label: "Function",
          properties: {
            name: "detachedPeer",
            filePath: "src/widget.ts",
            startLine: 4,
            endLine: 6,
            graphHealth: {
              topologyStatus: "detached_component",
              countedIncoming: 1,
              countedOutgoing: 0,
              componentId: "component-detached",
              componentSize: 2,
              componentReachableFromRoot: false,
              confidence: "candidate",
            },
          },
        },
        {
          id: "fn-connected-diagnostic",
          label: "Function",
          properties: {
            name: "connectedDiagnostic",
            filePath: "src/widget.ts",
            startLine: 8,
            endLine: 10,
            graphHealth: {
              topologyStatus: "connected",
              countedIncoming: 1,
              countedOutgoing: 1,
              componentReachableFromRoot: true,
              diagnostics: [
                {
                  kind: "unresolved_reference",
                  targetText: "testing.T",
                  count: 1,
                  classification: "test_framework",
                  actionability: "non_actionable",
                },
              ],
              confidence: "unknown",
            },
          },
        },
        {
          id: "fn-unknown-topology",
          label: "Function",
          properties: {
            name: "unknownTopology",
            filePath: "src/widget.ts",
            startLine: 12,
            endLine: 14,
            graphHealth: {
              topologyStatus: "unknown_connectivity",
              countedIncoming: 0,
              countedOutgoing: 0,
              componentReachableFromRoot: false,
              diagnostics: [
                {
                  kind: "unresolved_reference",
                  targetText: "mystery",
                  count: 1,
                },
              ],
              confidence: "unknown",
            },
          },
        },
      ],
      relationships: [
        {
          id: "rel-file-defines-widget",
          sourceId: "file-src-widget",
          targetId: "fn-detached-widget",
          type: "DEFINES",
          confidence: 1,
          reason: "fixture",
        },
        {
          id: "rel-widget-calls-peer",
          sourceId: "fn-detached-widget",
          targetId: "fn-detached-peer",
          type: "CALLS",
          confidence: 1,
          reason: "fixture",
        },
        {
          id: "rel-peer-calls-connected-diagnostic",
          sourceId: "fn-detached-peer",
          targetId: "fn-connected-diagnostic",
          type: "CALLS",
          confidence: 1,
          reason: "fixture",
        },
        {
          id: "rel-connected-diagnostic-calls-file",
          sourceId: "fn-connected-diagnostic",
          targetId: "file-src-widget",
          type: "CALLS",
          confidence: 1,
          reason: "fixture",
        },
      ],
    };

    await page.route(`${BACKEND_URL}/api/repos`, (route) =>
      route.fulfill({
        json: [{ name: "graph-health-demo", path: "/tmp/graph-health-demo" }],
      }),
    );
    await page.route(`${BACKEND_URL}/api/repo**`, (route) =>
      route.fulfill({
        json: {
          name: "graph-health-demo",
          path: "/tmp/graph-health-demo",
          repoPath: "/tmp/graph-health-demo",
        },
      }),
    );
    await page.route(`${BACKEND_URL}/api/graph**`, (route) =>
      route.fulfill({ json: graph }),
    );
    await page.route(`${BACKEND_URL}/api/file**`, (route) =>
      route.fulfill({
        json: {
          content: "export function detachedWidget() {\n  missingCall();\n}\n",
          startLine: 0,
          totalLines: 3,
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

  test("shows selected-node explanation and detached component focus action", async ({
    page,
  }) => {
    await page.goto(
      `${FRONTEND_URL}/?server=${encodeURIComponent(BACKEND_URL)}&project=graph-health-demo`,
    );

    await expect(page.locator('[data-testid="status-ready"]')).toBeVisible({
      timeout: 20_000,
    });

    await page.getByPlaceholder("Search nodes...").fill("detachedWidget");
    await page.getByRole("button", { name: /detachedWidget/ }).click();

    const detail = page.getByTestId("graph-health-node-detail");
    await expect(detail).toBeVisible({ timeout: 10_000 });
    await expect(detail).toContainText("Detached component");
    await expect(detail).toContainText("Unresolved reference x1: missingCall");
    await expect(detail).toContainText(
      "Detached: no accepted root reaches this counted-edge component.",
    );

    await page.getByRole("button", { name: "Focus component" }).click();
    await expect(page.getByText("detachedWidget").first()).toBeVisible();
  });

  test("selects all hidden node types from the left dashboard", async ({
    page,
  }) => {
    await page.goto(
      `${FRONTEND_URL}/?server=${encodeURIComponent(BACKEND_URL)}&project=graph-health-demo`,
    );

    await expect(page.locator('[data-testid="status-ready"]')).toBeVisible({
      timeout: 20_000,
    });

    await page.getByRole("button", { name: "Filters" }).click();

    const functionNodeTypeButton = page
      .locator('button[title^="Function ("]')
      .first();
    const selectAllNodeTypesButton = page.getByTitle("Select all node types");

    await expect(functionNodeTypeButton).toBeVisible({ timeout: 10_000 });
    await expect(selectAllNodeTypesButton).toBeDisabled();
    await functionNodeTypeButton.click();
    await expect(functionNodeTypeButton).toHaveAttribute(
      "aria-pressed",
      "false",
    );
    await expect(selectAllNodeTypesButton).toBeEnabled();
    await selectAllNodeTypesButton.click();
    await expect(functionNodeTypeButton).toHaveAttribute(
      "aria-pressed",
      "true",
    );
  });

  test("hiding unknown topology keeps connected diagnostic nodes selectable", async ({
    page,
  }) => {
    await page.goto(
      `${FRONTEND_URL}/?server=${encodeURIComponent(BACKEND_URL)}&project=graph-health-demo`,
    );

    await expect(page.locator('[data-testid="status-ready"]')).toBeVisible({
      timeout: 20_000,
    });

    await page.getByRole("button", { name: "Filters" }).click();
    const unknownTopologyButton = page
      .locator('button[title^="Unknown ("]')
      .first();
    await expect(unknownTopologyButton).toBeVisible({ timeout: 10_000 });
    await unknownTopologyButton.click();
    await expect(unknownTopologyButton).toHaveAttribute(
      "aria-pressed",
      "false",
    );

    await page.getByPlaceholder("Search nodes...").fill("connectedDiagnostic");
    await page.getByRole("button", { name: /connectedDiagnostic/ }).click();

    const detail = page.getByTestId("graph-health-node-detail");
    await expect(detail).toBeVisible({ timeout: 10_000 });
    await expect(detail).toContainText("Connected");
    await expect(detail).toContainText("Unresolved reference x1: testing.T");
    await expect(detail).toContainText("Test framework");
    await expect(detail).toContainText("Non-actionable");
  });

  test("graph shell Back returns to the exe-served start screen", async ({
    page,
  }) => {
    await page.goto(
      `${FRONTEND_URL}/?server=${encodeURIComponent(BACKEND_URL)}&project=graph-health-demo`,
    );

    await expect(page.locator('[data-testid="status-ready"]')).toBeVisible({
      timeout: 20_000,
    });

    await page.getByLabel("Back to Start screen").click();

    await expect(
      page.getByRole("button", { name: "Start Anvien" }),
    ).toBeVisible();
    await expect(page).not.toHaveURL(/\/Start-Anvien\.html$/);
    await expect(
      page.locator('[data-testid="server-reconnect-banner"]'),
    ).toHaveCount(0);
  });
});
