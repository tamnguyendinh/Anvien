import { test, expect, type Page } from "@playwright/test";

const BACKEND_URL = process.env.BACKEND_URL ?? "http://127.0.0.1:4848";
const FRONTEND_URL = process.env.FRONTEND_URL ?? "http://127.0.0.1:5228";

const makeNode = (
  label: string,
  index: number,
  appLayer: string,
  filePath = `${appLayer}/${label.toLowerCase()}-${index}.ts`,
) => ({
  id: `${label}:${filePath}:${index}`,
  label,
  properties: {
    name: `${label}${index}`,
    filePath,
    startLine: index + 1,
    endLine: index + 2,
    appLayer,
    resolutionConfidence: "direct",
    resolutionHealthBuckets: { clean: 1 },
  },
});

const createOrientationGraph = () => {
  const nodes = [
    ...Array.from({ length: 8 }, (_item, index) =>
      makeNode("Function", index, "backend", `backend/function-${index}.go`),
    ),
    ...Array.from({ length: 5 }, (_item, index) =>
      makeNode("Method", index, "backend", `backend/method-${index}.go`),
    ),
    ...Array.from({ length: 5 }, (_item, index) =>
      makeNode("Route", index, "api", `internal/httpapi/route-${index}.go`),
    ),
    ...Array.from({ length: 4 }, (_item, index) =>
      makeNode("Tool", index, "api", `internal/mcp/tool-${index}.go`),
    ),
    ...Array.from({ length: 6 }, (_item, index) =>
      makeNode("File", index, "frontend", `avmatrix-web/src/file-${index}.tsx`),
    ),
    ...Array.from({ length: 4 }, (_item, index) =>
      makeNode("Class", index, "frontend", `avmatrix-web/src/class-${index}.tsx`),
    ),
    ...Array.from({ length: 4 }, (_item, index) =>
      makeNode("File", index, "docs", `docs/guide-${index}.md`),
    ),
  ];

  const relationships = nodes.slice(1).map((node, index) => ({
    id: `rel-${index}`,
    sourceId: nodes[index].id,
    targetId: node.id,
    type: "CALLS",
    confidence: 1,
    reason: "orientation-label-fixture",
  }));

  return { nodes, relationships };
};

const createDenseSpacingGraph = () => {
  const nodes = [
    ...Array.from({ length: 1677 }, (_item, index) =>
      makeNode(
        "Function",
        index,
        "frontend",
        `avmatrix-web/src/dense/function-${index}.tsx`,
      ),
    ),
    ...Array.from({ length: 300 }, (_item, index) =>
      makeNode(
        "Method",
        index,
        "frontend",
        `avmatrix-web/src/dense/method-${index}.tsx`,
      ),
    ),
    ...Array.from({ length: 60 }, (_item, index) =>
      makeNode("Function", index, "backend", `internal/dense/function-${index}.go`),
    ),
    ...Array.from({ length: 80 }, (_item, index) =>
      makeNode("Route", index, "api", `internal/httpapi/dense-route-${index}.go`),
    ),
    ...Array.from({ length: 40 }, (_item, index) =>
      makeNode("File", index, "docs", `docs/dense-guide-${index}.md`),
    ),
  ];

  const relationships = nodes.slice(1).map((node, index) => ({
    id: `dense-rel-${index}`,
    sourceId: nodes[index].id,
    targetId: node.id,
    type: "CALLS",
    confidence: 1,
    reason: "dense-spacing-fixture",
  }));

  return { nodes, relationships };
};

const countOrientationLabelOverlaps = async (page: Page) =>
  page
    .locator(
      '[data-testid="graph-orientation-label-ring"], [data-testid="graph-orientation-label-island"]',
    )
    .evaluateAll((elements) => {
      const boxes = elements.map((element) => {
        const rect = element.getBoundingClientRect();
        return {
          left: rect.left,
          right: rect.right,
          top: rect.top,
          bottom: rect.bottom,
        };
      });
      let overlaps = 0;
      for (let leftIndex = 0; leftIndex < boxes.length; leftIndex++) {
        for (
          let rightIndex = leftIndex + 1;
          rightIndex < boxes.length;
          rightIndex++
        ) {
          const left = boxes[leftIndex];
          const right = boxes[rightIndex];
          if (
            left.left < right.right &&
            left.right > right.left &&
            left.top < right.bottom &&
            left.bottom > right.top
          ) {
            overlaps++;
          }
        }
      }
      return overlaps;
    });

const getLayoutNodeSpacingDiagnostics = async (page: Page) =>
  page.evaluate(() => {
    const win = window as typeof window & {
      __AVMATRIX_WEB_DIAGNOSTICS__?: {
        layoutNodeSpacing?: {
          nodeCount: number;
          islandCount: number;
          requiredEdgeGap: number;
          requiredCenterDistance: number;
          minObservedCenterDistance: number;
          minObservedEdgeGap: number;
          overlapCount: number;
          targetGapViolationCount: number;
        };
      };
    };
    return win.__AVMATRIX_WEB_DIAGNOSTICS__?.layoutNodeSpacing ?? null;
  });

const getScreenNodeSpacingDiagnostics = async (page: Page) =>
  page.evaluate(() => {
    const win = window as typeof window & {
      __AVMATRIX_WEB_DIAGNOSTICS__?: {
        screenNodeSpacing?: {
          coordinateSpace: "viewport_px";
          nodeCount: number;
          islandCount: number;
          viewportWidth: number;
          viewportHeight: number;
          cameraRatio: number;
          minRenderedRadius: number;
          maxRenderedRadius: number;
          maxRenderedDiameter: number;
          minObservedCenterDistance: number;
          minObservedEdgeGap: number;
          maxRequiredCenterDistance: number;
          overlapCount: number;
          targetGapViolationCount: number;
        };
      };
    };
    return win.__AVMATRIX_WEB_DIAGNOSTICS__?.screenNodeSpacing ?? null;
  });

test.describe("Graph orientation labels", () => {
  test.beforeEach(async ({ page }) => {
    const graph = createOrientationGraph();

    await page.route(`${BACKEND_URL}/api/repos`, (route) =>
      route.fulfill({
        json: [{ name: "orientation-demo", path: "/tmp/orientation-demo" }],
      }),
    );
    await page.route(`${BACKEND_URL}/api/repo**`, (route) =>
      route.fulfill({
        json: {
          name: "orientation-demo",
          path: "/tmp/orientation-demo",
          repoPath: "/tmp/orientation-demo",
        },
      }),
    );
    await page.route(`${BACKEND_URL}/api/graph**`, (route) =>
      route.fulfill({ json: graph }),
    );
    await page.route(`${BACKEND_URL}/api/file**`, (route) =>
      route.fulfill({
        json: {
          content: "export function fixture() {}\n",
          startLine: 0,
          totalLines: 1,
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

  test("shows readable ring and island labels on the desktop graph", async ({
    page,
  }, testInfo) => {
    await page.setViewportSize({ width: 1280, height: 800 });
    await page.goto(
      `${FRONTEND_URL}/?server=${encodeURIComponent(BACKEND_URL)}&project=orientation-demo`,
    );

    await expect(page.locator('[data-testid="status-ready"]')).toBeVisible({
      timeout: 20_000,
    });

    await expect
      .poll(
        async () => page.locator('[data-testid="graph-orientation-label-ring"]').count(),
        { timeout: 10_000 },
      )
      .toBeGreaterThanOrEqual(3);
    await expect
      .poll(
        async () => page.locator('[data-testid="graph-orientation-label-island"]').count(),
        { timeout: 10_000 },
      )
      .toBeGreaterThanOrEqual(4);

    await expect(
      page.locator('[data-testid="graph-orientation-label-ring"][data-label-source="backend"]'),
    ).toContainText("Backend");
    await expect(
      page.locator('[data-testid="graph-orientation-label-ring"][data-label-source="frontend"]'),
    ).toContainText("Frontend");
    await expect(
      page.locator(
        '[data-testid="graph-orientation-label-island"][data-label-source="backend:Method"]',
      ),
    ).toContainText("Method");
    await expect
      .poll(async () => countOrientationLabelOverlaps(page), { timeout: 10_000 })
      .toBe(0);

    await page.screenshot({
      path: testInfo.outputPath("graph-orientation-labels-desktop.png"),
      fullPage: false,
    });
  });

  test("keeps labels visible on a smaller viewport and updates after filters", async ({
    page,
  }, testInfo) => {
    await page.setViewportSize({ width: 480, height: 720 });
    await page.goto(
      `${FRONTEND_URL}/?server=${encodeURIComponent(BACKEND_URL)}&project=orientation-demo`,
    );

    await expect(page.locator('[data-testid="status-ready"]')).toBeVisible({
      timeout: 20_000,
    });

    await expect
      .poll(
        async () => page.locator('[data-testid="graph-orientation-label-ring"]').count(),
        { timeout: 10_000 },
      )
      .toBeGreaterThanOrEqual(2);
    await expect(
      page.locator(
        '[data-testid="graph-orientation-label-island"][data-label-source="backend:Method"]',
      ),
    ).toContainText("Method");

    await page.getByRole("button", { name: "Filters" }).click();
    const methodToggle = page.locator('button[title^="Method ("]').first();
    await expect(methodToggle).toBeVisible({ timeout: 10_000 });
    await methodToggle.click();

    await expect(
      page.locator(
        '[data-testid="graph-orientation-label-island"][data-label-source="backend:Method"]',
      ),
    ).toHaveCount(0);
    await expect
      .poll(async () => countOrientationLabelOverlaps(page), { timeout: 10_000 })
      .toBe(0);

    await page.screenshot({
      path: testInfo.outputPath("graph-orientation-labels-small-filtered.png"),
      fullPage: false,
    });
  });

  test("keeps dense graph nodes separated by the default node gap", async ({
    page,
  }, testInfo) => {
    await page.unroute(`${BACKEND_URL}/api/graph**`);
    await page.route(`${BACKEND_URL}/api/graph**`, (route) =>
      route.fulfill({ json: createDenseSpacingGraph() }),
    );

    await page.setViewportSize({ width: 1280, height: 800 });
    await page.goto(
      `${FRONTEND_URL}/?server=${encodeURIComponent(BACKEND_URL)}&project=orientation-demo`,
    );

    await expect(page.locator('[data-testid="status-ready"]')).toBeVisible({
      timeout: 20_000,
    });

    await expect
      .poll(
        async () => (await getLayoutNodeSpacingDiagnostics(page))?.nodeCount ?? 0,
        { timeout: 10_000 },
      )
      .toBeGreaterThanOrEqual(2157);

    const desktopDiagnostics = await getLayoutNodeSpacingDiagnostics(page);
    expect(desktopDiagnostics?.islandCount).toBeGreaterThanOrEqual(4);
    expect(desktopDiagnostics?.overlapCount).toBe(0);
    expect(desktopDiagnostics?.targetGapViolationCount).toBe(0);
    expect(desktopDiagnostics?.minObservedCenterDistance).toBeGreaterThanOrEqual(
      desktopDiagnostics?.requiredCenterDistance ?? Number.POSITIVE_INFINITY,
    );
    expect(desktopDiagnostics?.minObservedEdgeGap).toBeGreaterThanOrEqual(
      desktopDiagnostics?.requiredEdgeGap ?? Number.POSITIVE_INFINITY,
    );
    const desktopScreenDiagnostics = await getScreenNodeSpacingDiagnostics(page);
    expect(desktopScreenDiagnostics?.coordinateSpace).toBe("viewport_px");
    expect(desktopScreenDiagnostics?.nodeCount).toBeGreaterThanOrEqual(2077);
    expect(desktopScreenDiagnostics?.overlapCount).toBe(0);
    expect(desktopScreenDiagnostics?.targetGapViolationCount).toBe(0);
    expect(desktopScreenDiagnostics?.maxRenderedRadius).toBeGreaterThanOrEqual(0.74);
    expect(desktopScreenDiagnostics?.minObservedEdgeGap).toBeGreaterThanOrEqual(
      desktopScreenDiagnostics?.maxRenderedDiameter ?? Number.POSITIVE_INFINITY,
    );

    await expect
      .poll(
        async () => page.locator('[data-testid="graph-orientation-label-ring"]').count(),
        { timeout: 10_000 },
      )
      .toBeGreaterThanOrEqual(1);
    await expect(
      page.locator('[data-testid="graph-orientation-label-ring"][data-label-source="frontend"]'),
    ).toContainText("Frontend");
    await expect
      .poll(async () => countOrientationLabelOverlaps(page), { timeout: 10_000 })
      .toBe(0);

    await page.screenshot({
      path: testInfo.outputPath("graph-node-spacing-dense-desktop.png"),
      fullPage: false,
    });

    await page.setViewportSize({ width: 520, height: 720 });
    await page.reload();
    await expect(page.locator('[data-testid="status-ready"]')).toBeVisible({
      timeout: 20_000,
    });
    await expect
      .poll(
        async () => (await getScreenNodeSpacingDiagnostics(page))?.nodeCount ?? 0,
        { timeout: 10_000 },
      )
      .toBeGreaterThanOrEqual(2077);
    const smallScreenDiagnostics = await getScreenNodeSpacingDiagnostics(page);
    expect(smallScreenDiagnostics?.viewportWidth).toBeGreaterThan(200);
    expect(smallScreenDiagnostics?.maxRenderedRadius).toBeGreaterThanOrEqual(0.74);
    expect(smallScreenDiagnostics?.overlapCount).toBe(0);
    expect(smallScreenDiagnostics?.targetGapViolationCount).toBe(0);
    expect(smallScreenDiagnostics?.minObservedEdgeGap).toBeGreaterThanOrEqual(
      smallScreenDiagnostics?.maxRenderedDiameter ?? Number.POSITIVE_INFINITY,
    );
    await expect
      .poll(async () => countOrientationLabelOverlaps(page), { timeout: 10_000 })
      .toBe(0);
    await page.screenshot({
      path: testInfo.outputPath("graph-node-spacing-dense-small.png"),
      fullPage: false,
    });
  });
});
