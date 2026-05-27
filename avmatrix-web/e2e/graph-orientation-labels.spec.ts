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

type FixtureGraph = ReturnType<typeof createDenseSpacingGraph>;
type FixtureNode = FixtureGraph["nodes"][number];

const DEFAULT_VISIBLE_FIXTURE_LABELS = new Set([
  "Project",
  "Package",
  "Module",
  "Folder",
  "File",
  "Documentation",
  "Class",
  "Function",
  "Method",
  "Interface",
  "Enum",
  "Type",
]);
const DOCUMENTATION_FILE_EXTENSIONS = new Set([
  ".md",
  ".mdx",
  ".rst",
  ".txt",
]);
const DOCUMENTATION_PATH_SEGMENTS = new Set([
  "doc",
  "docs",
  "documentation",
  "wiki",
]);

const sortKeys = (values: string[]) =>
  [...values].sort((left, right) => left.localeCompare(right));

const uniqueSorted = (values: string[]) => sortKeys([...new Set(values)]);

const getFixtureDisplayLabel = (node: FixtureNode): string => {
  if (node.label === "Documentation") return "Documentation";
  const path = node.properties.filePath.replace(/\\/g, "/").toLowerCase();
  const pathSegments = path.split("/").filter(Boolean);
  if (pathSegments.some((segment) => DOCUMENTATION_PATH_SEGMENTS.has(segment))) {
    return "Documentation";
  }
  const baseName = pathSegments.at(-1) ?? "";
  const extensionStart = baseName.lastIndexOf(".");
  const extension = extensionStart > 0 ? baseName.slice(extensionStart) : "";
  if (DOCUMENTATION_FILE_EXTENSIONS.has(extension)) return "Documentation";
  return node.label;
};

const getDefaultVisibleFixtureNodes = (graph: FixtureGraph) =>
  graph.nodes.filter((node) =>
    DEFAULT_VISIBLE_FIXTURE_LABELS.has(getFixtureDisplayLabel(node)),
  );

const getFixtureRingInventory = (graph: FixtureGraph) =>
  uniqueSorted(
    getDefaultVisibleFixtureNodes(graph).map(
      (node) => node.properties.appLayer ?? "missing_app_layer",
    ),
  );

const getFixtureIslandInventory = (graph: FixtureGraph) =>
  uniqueSorted(
    getDefaultVisibleFixtureNodes(graph).map(
      (node) =>
        `${node.properties.appLayer ?? "missing_app_layer"}:${getFixtureDisplayLabel(node)}`,
    ),
  );

const sumCounts = (counts: Record<string, number>) =>
  Object.values(counts).reduce((total, count) => total + count, 0);

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

const getOrientationLabelSources = async (
  page: Page,
  kind: "ring" | "island",
) =>
  page
    .locator(`[data-testid="graph-orientation-label-${kind}"]`)
    .evaluateAll((elements) =>
      elements
        .map((element) => element.getAttribute("data-label-source") ?? "")
        .filter(Boolean)
        .sort((left, right) => left.localeCompare(right)),
    );

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
          visibleViewportNodeCount: number;
          visibleViewportIslandCounts: Record<string, number>;
          cameraRatio: number;
          cameraX: number;
          cameraY: number;
          viewportGraphCenterX: number;
          viewportGraphCenterY: number;
          viewportGraphMinX: number;
          viewportGraphMaxX: number;
          viewportGraphMinY: number;
          viewportGraphMaxY: number;
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

const getGraphOverviewDiagnostics = async (page: Page) =>
  page.evaluate(() => {
    const win = window as typeof window & {
      __AVMATRIX_WEB_DIAGNOSTICS__?: {
        graphOverview?: {
          nodeCount: number;
          visibleViewportNodeCount: number;
          visibleColorCount: number;
          visibleRingCount: number;
          visibleIslandCount: number;
          dominantIslandKey: string;
          dominantIslandShare: number;
          visibleColorCounts: Record<string, number>;
          visibleRingCounts: Record<string, number>;
          visibleIslandCounts: Record<string, number>;
          visibleNodeTypeCounts: Record<string, number>;
          graphRingCounts: Record<string, number>;
          graphIslandCounts: Record<string, number>;
          graphNodeTypeCounts: Record<string, number>;
          visibleRingInventory: string[];
          visibleNodeTypeInventory: string[];
          graphRingInventory: string[];
          graphIslandInventory: string[];
          filterNodeTypeInventory: string[];
          cameraRatio: number;
        };
      };
    };
    return win.__AVMATRIX_WEB_DIAGNOSTICS__?.graphOverview ?? null;
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
    const orientationGraph = createOrientationGraph();
    const expectedRingInventory = getFixtureRingInventory(orientationGraph);
    const expectedIslandInventory = getFixtureIslandInventory(orientationGraph);

    await page.setViewportSize({ width: 1280, height: 800 });
    await page.goto(
      `${FRONTEND_URL}/?server=${encodeURIComponent(BACKEND_URL)}&project=orientation-demo`,
    );

    await expect(page.locator('[data-testid="status-ready"]')).toBeVisible({
      timeout: 20_000,
    });

    await expect
      .poll(
        async () => (await getGraphOverviewDiagnostics(page))?.graphRingInventory.length ?? 0,
        { timeout: 10_000 },
      )
      .toBe(expectedRingInventory.length);
    const overviewDiagnostics = await getGraphOverviewDiagnostics(page);
    expect(overviewDiagnostics).not.toBeNull();
    expect(overviewDiagnostics?.graphRingInventory).toEqual(expectedRingInventory);
    expect(overviewDiagnostics?.graphIslandInventory).toEqual(expectedIslandInventory);
    await expect
      .poll(
        async () => getOrientationLabelSources(page, "ring"),
        { timeout: 10_000 },
      )
      .toEqual(expectedRingInventory);
    await expect
      .poll(
        async () => getOrientationLabelSources(page, "island"),
        { timeout: 10_000 },
      )
      .toEqual(expectedIslandInventory);

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
        async () => (await getGraphOverviewDiagnostics(page))?.visibleRingInventory.length ?? 0,
        { timeout: 10_000 },
      )
      .toBeGreaterThan(0);
    const smallOverviewDiagnostics = await getGraphOverviewDiagnostics(page);
    expect(smallOverviewDiagnostics).not.toBeNull();
    await expect
      .poll(
        async () => getOrientationLabelSources(page, "ring"),
        { timeout: 10_000 },
      )
      .toEqual(smallOverviewDiagnostics?.visibleRingInventory);
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

  test("keeps dense graph overview inventory visible on default load", async ({
    page,
  }, testInfo) => {
    const denseGraph = createDenseSpacingGraph();
    const expectedDenseRingInventory = getFixtureRingInventory(denseGraph);
    const expectedDenseIslandInventory = getFixtureIslandInventory(denseGraph);
    const expectedDenseVisibleNodeCount =
      getDefaultVisibleFixtureNodes(denseGraph).length;

    await page.unroute(`${BACKEND_URL}/api/graph**`);
    await page.route(`${BACKEND_URL}/api/graph**`, (route) =>
      route.fulfill({ json: denseGraph }),
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
      .toBe(denseGraph.nodes.length);

    const desktopDiagnostics = await getLayoutNodeSpacingDiagnostics(page);
    const desktopOverviewDiagnostics = await getGraphOverviewDiagnostics(page);
    expect(desktopOverviewDiagnostics).not.toBeNull();
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
    expect(desktopScreenDiagnostics?.nodeCount).toBe(expectedDenseVisibleNodeCount);
    expect(sumCounts(desktopOverviewDiagnostics?.graphNodeTypeCounts ?? {})).toBe(
      expectedDenseVisibleNodeCount,
    );
    expect(desktopOverviewDiagnostics?.graphRingInventory).toEqual(
      expectedDenseRingInventory,
    );
    expect(desktopOverviewDiagnostics?.graphIslandInventory).toEqual(
      expectedDenseIslandInventory,
    );
    expect(desktopOverviewDiagnostics?.visibleRingInventory).toEqual(
      Object.keys(desktopOverviewDiagnostics?.visibleRingCounts ?? {}).sort(
        (left, right) => left.localeCompare(right),
      ),
    );
    expect(desktopOverviewDiagnostics?.filterNodeTypeInventory).toEqual(
      Object.keys(desktopOverviewDiagnostics?.graphNodeTypeCounts ?? {}).sort(
        (left, right) => left.localeCompare(right),
      ),
    );
    expect(desktopOverviewDiagnostics?.graphRingInventory).toEqual(
      Object.keys(desktopOverviewDiagnostics?.graphRingCounts ?? {}).sort(
        (left, right) => left.localeCompare(right),
      ),
    );
    expect(
      Object.keys(desktopOverviewDiagnostics?.visibleRingCounts ?? {}).sort(
        (left, right) => left.localeCompare(right),
      ),
    ).toEqual(expectedDenseRingInventory);
    expect(
      Object.keys(desktopOverviewDiagnostics?.visibleIslandCounts ?? {}).sort(
        (left, right) => left.localeCompare(right),
      ),
    ).toEqual(expectedDenseIslandInventory);
    expect(desktopOverviewDiagnostics?.visibleNodeTypeInventory).toEqual(
      desktopOverviewDiagnostics?.filterNodeTypeInventory,
    );
    expect(desktopOverviewDiagnostics?.visibleColorCount).toBe(
      Object.keys(desktopOverviewDiagnostics?.visibleColorCounts ?? {}).length,
    );
    expect(desktopOverviewDiagnostics?.visibleIslandCount).toBe(
      expectedDenseIslandInventory.length,
    );
    expect(desktopOverviewDiagnostics?.dominantIslandShare).toBeLessThan(0.85);
    const desktopVisibleIslandCounts =
      desktopScreenDiagnostics?.visibleViewportIslandCounts ?? {};
    expect(
      Object.keys(desktopVisibleIslandCounts).sort((left, right) =>
        left.localeCompare(right),
      ),
    ).toEqual(expectedDenseIslandInventory);
    expect(desktopScreenDiagnostics?.visibleViewportNodeCount).toBe(
      sumCounts(desktopVisibleIslandCounts),
    );
    expect(desktopScreenDiagnostics?.cameraRatio).toBeGreaterThan(0);

    await expect
      .poll(
        async () => getOrientationLabelSources(page, "ring"),
        { timeout: 10_000 },
      )
      .toEqual(expectedDenseRingInventory);
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
      .toBe(expectedDenseVisibleNodeCount);
    const smallScreenDiagnostics = await getScreenNodeSpacingDiagnostics(page);
    expect(smallScreenDiagnostics?.viewportWidth).toBeGreaterThan(200);
    const smallVisibleIslandCounts =
      smallScreenDiagnostics?.visibleViewportIslandCounts ?? {};
    expect(smallScreenDiagnostics?.visibleViewportNodeCount).toBe(
      sumCounts(smallVisibleIslandCounts),
    );
    expect(
      Object.keys(smallVisibleIslandCounts).sort((left, right) =>
        left.localeCompare(right),
      ).length,
    ).toBeGreaterThan(0);
    expect(smallScreenDiagnostics?.cameraRatio).toBeGreaterThan(0);
    await expect
      .poll(async () => countOrientationLabelOverlaps(page), { timeout: 10_000 })
      .toBe(0);
    await page.screenshot({
      path: testInfo.outputPath("graph-node-spacing-dense-small.png"),
      fullPage: false,
    });
  });
});
