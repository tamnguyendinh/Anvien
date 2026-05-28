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

type RuntimeGraphInteractionSample = {
  recordedAt: number;
  mode: "overview" | "zoom-in" | "zoom-out" | "wheel-zoom" | "detail-focus";
  targetNodeId: string;
  coordinateSpace: "viewport_px";
  nodeCount: number;
  islandCount: number;
  viewportWidth: number;
  viewportHeight: number;
  visibleViewportNodeCount: number;
  visibleViewportIslandCount: number;
  cameraRatio: number;
  cameraX: number;
  cameraY: number;
  minRenderedRadius: number;
  maxRenderedRadius: number;
  maxRenderedDiameter: number;
  minObservedCenterDistance: number;
  minObservedEdgeGap: number;
  maxRequiredCenterDistance: number;
  overlapCount: number;
  targetGapViolationCount: number;
};

type RuntimeGraphInteractionDiagnostics = {
  currentMode: RuntimeGraphInteractionSample["mode"];
  currentTargetNodeId: string;
  lastModeChangedAt: number;
  overviewSamples: RuntimeGraphInteractionSample[];
  zoomSamples: RuntimeGraphInteractionSample[];
  detailFocusSamples: RuntimeGraphInteractionSample[];
  dynamicGapSamples: RuntimeGraphInteractionSample[];
};

type BrowserPerformanceProbeSnapshot = {
  longTaskCount: number;
  totalLongTaskMs: number;
  longestLongTaskMs: number;
  maxFrameDeltaMs: number;
  frameDropCount: number;
  screenWrites: number;
  overviewWrites: number;
};

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

const getLayoutRingDiagnostics = async (page: Page) =>
  page.evaluate(() => {
    const win = window as typeof window & {
      __AVMATRIX_WEB_DIAGNOSTICS__?: {
        layoutRings?: {
          nodeCount: number;
          ringCount: number;
          ringNodeCounts: Record<string, number>;
          ringCenters: Record<string, { x: number; y: number }>;
          ringIslandCounts: Record<string, number>;
          apiBetweenBackendAndFrontend: boolean;
          docsCentered: boolean;
          sameColorIslandViolations: number;
        };
      };
    };
    return win.__AVMATRIX_WEB_DIAGNOSTICS__?.layoutRings ?? null;
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

const getGraphInteractionDiagnostics = async (page: Page) =>
  page.evaluate(() => {
    const win = window as typeof window & {
      __AVMATRIX_WEB_DIAGNOSTICS__?: {
        graphInteraction?: RuntimeGraphInteractionDiagnostics;
      };
    };
    return win.__AVMATRIX_WEB_DIAGNOSTICS__?.graphInteraction ?? null;
  });

const installBrowserPerformanceProbe = async (page: Page) =>
  page.addInitScript(() => {
    const probe = {
      longTasks: [] as number[],
      maxFrameDelta: 0,
      frameDrops: 0,
      screenWrites: 0,
      overviewWrites: 0,
      lastScreenRecordedAt: 0,
      lastOverviewRecordedAt: 0,
      reset() {
        const diagnostics = (
          window as typeof window & {
            __AVMATRIX_WEB_DIAGNOSTICS__?: {
              screenNodeSpacing?: { recordedAt: number };
              graphOverview?: { recordedAt: number };
            };
          }
        ).__AVMATRIX_WEB_DIAGNOSTICS__;
        this.longTasks = [];
        this.maxFrameDelta = 0;
        this.frameDrops = 0;
        this.screenWrites = 0;
        this.overviewWrites = 0;
        this.lastScreenRecordedAt =
          diagnostics?.screenNodeSpacing?.recordedAt ?? 0;
        this.lastOverviewRecordedAt =
          diagnostics?.graphOverview?.recordedAt ?? 0;
      },
      snapshot(): BrowserPerformanceProbeSnapshot {
        const totalLongTaskMs = this.longTasks.reduce(
          (total, duration) => total + duration,
          0,
        );
        return {
          longTaskCount: this.longTasks.length,
          totalLongTaskMs,
          longestLongTaskMs: this.longTasks.reduce(
            (maximum, duration) => Math.max(maximum, duration),
            0,
          ),
          maxFrameDeltaMs: this.maxFrameDelta,
          frameDropCount: this.frameDrops,
          screenWrites: this.screenWrites,
          overviewWrites: this.overviewWrites,
        };
      },
    };
    (
      window as typeof window & {
        __AVMATRIX_PHASE8_PROBE__?: typeof probe;
      }
    ).__AVMATRIX_PHASE8_PROBE__ = probe;
    try {
      const observer = new PerformanceObserver((list) => {
        for (const entry of list.getEntries()) {
          probe.longTasks.push(entry.duration);
        }
      });
      observer.observe({ entryTypes: ["longtask"] });
    } catch (_error) {
      // The probe still records frame and diagnostics-write counts.
    }

    let previousFrame = 0;
    const tick = (now: number) => {
      if (previousFrame > 0) {
        const delta = now - previousFrame;
        probe.maxFrameDelta = Math.max(probe.maxFrameDelta, delta);
        if (delta > 32) probe.frameDrops++;
      }
      previousFrame = now;

      const diagnostics = (
        window as typeof window & {
          __AVMATRIX_WEB_DIAGNOSTICS__?: {
            screenNodeSpacing?: { recordedAt: number };
            graphOverview?: { recordedAt: number };
          };
        }
      ).__AVMATRIX_WEB_DIAGNOSTICS__;
      const screenRecordedAt = diagnostics?.screenNodeSpacing?.recordedAt ?? 0;
      const overviewRecordedAt = diagnostics?.graphOverview?.recordedAt ?? 0;
      if (
        screenRecordedAt > 0 &&
        screenRecordedAt !== probe.lastScreenRecordedAt
      ) {
        probe.screenWrites++;
        probe.lastScreenRecordedAt = screenRecordedAt;
      }
      if (
        overviewRecordedAt > 0 &&
        overviewRecordedAt !== probe.lastOverviewRecordedAt
      ) {
        probe.overviewWrites++;
        probe.lastOverviewRecordedAt = overviewRecordedAt;
      }
      requestAnimationFrame(tick);
    };
    requestAnimationFrame(tick);
  });

const resetBrowserPerformanceProbe = async (page: Page) =>
  page.evaluate(() => {
    (
      window as typeof window & {
        __AVMATRIX_PHASE8_PROBE__?: { reset: () => void };
      }
    ).__AVMATRIX_PHASE8_PROBE__?.reset();
  });

const getBrowserPerformanceProbe = async (page: Page) =>
  page.evaluate(() => {
    return (
      window as typeof window & {
        __AVMATRIX_PHASE8_PROBE__?: {
          snapshot: () => BrowserPerformanceProbeSnapshot;
        };
      }
    ).__AVMATRIX_PHASE8_PROBE__?.snapshot() ?? null;
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
    const expectedDenseLayoutRingInventory = uniqueSorted(
      denseGraph.nodes.map(
        (node) => node.properties.appLayer ?? "missing_app_layer",
      ),
    );
    const expectedDenseIslandInventory = getFixtureIslandInventory(denseGraph);
    const expectedDenseVisibleNodeCount =
      getDefaultVisibleFixtureNodes(denseGraph).length;
    const expectedSearchTargetId =
      denseGraph.nodes.find((node) => node.properties.name === "Function0")?.id ??
      "";

    await page.unroute(`${BACKEND_URL}/api/graph**`);
    await page.route(`${BACKEND_URL}/api/graph**`, (route) =>
      route.fulfill({ json: denseGraph }),
    );

    await installBrowserPerformanceProbe(page);
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
    const desktopRingDiagnostics = await getLayoutRingDiagnostics(page);
    const desktopOverviewDiagnostics = await getGraphOverviewDiagnostics(page);
    expect(desktopOverviewDiagnostics).not.toBeNull();
    expect(desktopRingDiagnostics).not.toBeNull();
    expect(desktopRingDiagnostics?.nodeCount).toBe(denseGraph.nodes.length);
    expect(desktopRingDiagnostics?.ringCount).toBe(
      expectedDenseLayoutRingInventory.length,
    );
    expect(
      Object.keys(desktopRingDiagnostics?.ringNodeCounts ?? {}).sort((left, right) =>
        left.localeCompare(right),
      ),
    ).toEqual(expectedDenseLayoutRingInventory);
    expect(desktopRingDiagnostics?.apiBetweenBackendAndFrontend).toBe(true);
    expect(desktopRingDiagnostics?.docsCentered).toBe(true);
    expect(desktopRingDiagnostics?.sameColorIslandViolations).toBe(0);
    expect(desktopRingDiagnostics?.ringIslandCounts.frontend).toBeGreaterThan(1);
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
        async () =>
          (await getGraphInteractionDiagnostics(page))?.overviewSamples.length ?? 0,
        { timeout: 10_000 },
      )
      .toBeGreaterThan(0);
    const overviewInteraction = await getGraphInteractionDiagnostics(page);
    const latestOverviewSample = overviewInteraction?.overviewSamples.at(-1);
    expect(latestOverviewSample?.mode).toBe("overview");
    expect(latestOverviewSample?.nodeCount).toBe(expectedDenseVisibleNodeCount);
    expect(latestOverviewSample?.visibleViewportNodeCount).toBe(
      desktopScreenDiagnostics?.visibleViewportNodeCount,
    );
    expect(latestOverviewSample?.visibleViewportIslandCount).toBe(
      expectedDenseIslandInventory.length,
    );
    expect(overviewInteraction?.dynamicGapSamples.at(-1)?.mode).toBe("overview");

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
    const loadProbe = await getBrowserPerformanceProbe(page);
    expect(loadProbe?.screenWrites).toBeGreaterThan(0);
    expect(loadProbe?.overviewWrites).toBeGreaterThan(0);

    const initialZoomDiagnostics = await getScreenNodeSpacingDiagnostics(page);
    expect(initialZoomDiagnostics).not.toBeNull();

    await resetBrowserPerformanceProbe(page);
    const graphBox = await page.locator(".sigma-container").boundingBox();
    expect(graphBox).not.toBeNull();
    await page.mouse.move(
      (graphBox?.x ?? 0) + (graphBox?.width ?? 0) / 2,
      (graphBox?.y ?? 0) + (graphBox?.height ?? 0) / 2,
    );
    await page.mouse.wheel(0, -600);
    await expect
      .poll(
        async () => (await getScreenNodeSpacingDiagnostics(page))?.cameraRatio ?? 1,
        { timeout: 10_000 },
      )
      .toBeLessThan((initialZoomDiagnostics?.cameraRatio ?? 0) * 0.75);
    const wheelZoomDiagnostics = await getScreenNodeSpacingDiagnostics(page);
    expect(wheelZoomDiagnostics?.maxRenderedRadius).toBeGreaterThan(
      (initialZoomDiagnostics?.maxRenderedRadius ?? Number.POSITIVE_INFINITY) *
        1.1,
    );
    await expect
      .poll(
        async () =>
          (await getGraphInteractionDiagnostics(page))?.zoomSamples.some(
            (sample) => sample.mode === "wheel-zoom",
          ) ?? false,
        { timeout: 10_000 },
      )
      .toBe(true);
    const wheelProbe = await getBrowserPerformanceProbe(page);
    expect(wheelProbe?.screenWrites).toBeLessThanOrEqual(3);
    expect(wheelProbe?.overviewWrites).toBeLessThanOrEqual(3);

    await resetBrowserPerformanceProbe(page);
    await page.getByTitle("Zoom In").click();
    await expect
      .poll(
        async () => (await getScreenNodeSpacingDiagnostics(page))?.cameraRatio ?? 1,
        { timeout: 10_000 },
      )
      .toBeLessThan((wheelZoomDiagnostics?.cameraRatio ?? 0) * 0.75);
    const zoomInOneDiagnostics = await getScreenNodeSpacingDiagnostics(page);
    expect(zoomInOneDiagnostics?.maxRenderedRadius).toBeGreaterThan(
      (wheelZoomDiagnostics?.maxRenderedRadius ?? Number.POSITIVE_INFINITY) *
        1.1,
    );
    const buttonProbe = await getBrowserPerformanceProbe(page);
    expect(buttonProbe?.screenWrites).toBeLessThanOrEqual(3);
    expect(buttonProbe?.overviewWrites).toBeLessThanOrEqual(3);

    await page.getByTitle("Zoom In").click();
    await expect
      .poll(
        async () => (await getScreenNodeSpacingDiagnostics(page))?.cameraRatio ?? 1,
        { timeout: 10_000 },
      )
      .toBeLessThan((zoomInOneDiagnostics?.cameraRatio ?? 0) * 0.75);
    const zoomInTwoDiagnostics = await getScreenNodeSpacingDiagnostics(page);
    expect(zoomInTwoDiagnostics?.maxRenderedRadius).toBeGreaterThan(
      (zoomInOneDiagnostics?.maxRenderedRadius ?? Number.POSITIVE_INFINITY) *
        1.1,
    );

    await page.getByTitle("Zoom Out").click();
    await expect
      .poll(
        async () => (await getScreenNodeSpacingDiagnostics(page))?.cameraRatio ?? 0,
        { timeout: 10_000 },
      )
      .toBeGreaterThan(
        (zoomInTwoDiagnostics?.cameraRatio ?? Number.POSITIVE_INFINITY) * 1.25,
      );
    const zoomOutDiagnostics = await getScreenNodeSpacingDiagnostics(page);
    expect(zoomOutDiagnostics?.maxRenderedRadius).toBeLessThan(
      (zoomInTwoDiagnostics?.maxRenderedRadius ?? 0) * 0.9,
    );
    const zoomInteraction = await getGraphInteractionDiagnostics(page);
    const zoomInSamples =
      zoomInteraction?.zoomSamples.filter((sample) => sample.mode === "zoom-in") ??
      [];
    const wheelZoomSamples =
      zoomInteraction?.zoomSamples.filter(
        (sample) => sample.mode === "wheel-zoom",
      ) ?? [];
    const zoomOutSamples =
      zoomInteraction?.zoomSamples.filter((sample) => sample.mode === "zoom-out") ??
      [];
    expect(zoomInSamples.length).toBeGreaterThan(0);
    expect(wheelZoomSamples.length).toBeGreaterThan(0);
    expect(zoomOutSamples.length).toBeGreaterThan(0);
    expect(wheelZoomSamples.at(-1)?.maxRenderedRadius).toBeGreaterThan(
      (initialZoomDiagnostics?.maxRenderedRadius ?? Number.POSITIVE_INFINITY) *
        1.1,
    );
    expect(zoomInSamples.at(-1)?.maxRenderedRadius).toBeGreaterThan(
      (initialZoomDiagnostics?.maxRenderedRadius ?? Number.POSITIVE_INFINITY) *
        1.1,
    );
    expect(zoomOutSamples.at(-1)?.maxRenderedRadius).toBeLessThan(
      (zoomInTwoDiagnostics?.maxRenderedRadius ?? 0) * 0.9,
    );

    await page.getByPlaceholder("Search nodes...").fill("Function0");
    await page.locator("button").filter({ hasText: "Function0" }).first().click();
    await expect
      .poll(
        async () => (await getScreenNodeSpacingDiagnostics(page))?.cameraRatio ?? 1,
        { timeout: 10_000 },
      )
      .toBeLessThan((zoomOutDiagnostics?.cameraRatio ?? 0) * 0.5);
    await expect
      .poll(
        async () =>
          (await getScreenNodeSpacingDiagnostics(page))
            ?.targetGapViolationCount ?? -1,
        { timeout: 10_000 },
      )
      .toBe(0);
    const searchFocusDiagnostics = await getScreenNodeSpacingDiagnostics(page);
    expect(searchFocusDiagnostics?.visibleViewportNodeCount).toBeGreaterThan(0);
    expect(searchFocusDiagnostics?.overlapCount).toBe(0);
    expect(searchFocusDiagnostics?.targetGapViolationCount).toBe(0);
    expect(searchFocusDiagnostics?.minObservedEdgeGap).toBeGreaterThanOrEqual(
      searchFocusDiagnostics?.maxRenderedDiameter ?? Number.POSITIVE_INFINITY,
    );
    await expect
      .poll(
        async () =>
          (await getGraphInteractionDiagnostics(page))?.detailFocusSamples
            .length ?? 0,
        { timeout: 10_000 },
      )
      .toBeGreaterThan(0);
    const detailInteraction = await getGraphInteractionDiagnostics(page);
    const searchFocusSample = detailInteraction?.detailFocusSamples.at(-1);
    expect(expectedSearchTargetId).not.toBe("");
    expect(searchFocusSample?.mode).toBe("detail-focus");
    expect(searchFocusSample?.targetNodeId).toBe(expectedSearchTargetId);
    expect(searchFocusSample?.visibleViewportNodeCount).toBeGreaterThan(0);
    expect(searchFocusSample?.overlapCount).toBe(0);
    expect(searchFocusSample?.targetGapViolationCount).toBe(0);
    expect(searchFocusSample?.maxRequiredCenterDistance).toBeGreaterThan(
      searchFocusSample?.maxRenderedDiameter ?? Number.POSITIVE_INFINITY,
    );

    await page.getByTitle("Zoom Out").click();
    await expect
      .poll(
        async () => (await getScreenNodeSpacingDiagnostics(page))?.cameraRatio ?? 0,
        { timeout: 10_000 },
      )
      .toBeGreaterThan(
        (searchFocusDiagnostics?.cameraRatio ?? Number.POSITIVE_INFINITY) * 1.25,
      );
    const sameSelectionShiftedDiagnostics =
      await getScreenNodeSpacingDiagnostics(page);

    await page.getByTitle("Focus on Selected Node").click();
    await expect
      .poll(
        async () => (await getScreenNodeSpacingDiagnostics(page))?.cameraRatio ?? 1,
        { timeout: 10_000 },
      )
      .toBeLessThan((sameSelectionShiftedDiagnostics?.cameraRatio ?? 0) * 0.75);
    await expect
      .poll(
        async () =>
          (await getScreenNodeSpacingDiagnostics(page))
            ?.targetGapViolationCount ?? -1,
        { timeout: 10_000 },
      )
      .toBe(0);
    const sameSelectionFocusDiagnostics =
      await getScreenNodeSpacingDiagnostics(page);
    expect(
      sameSelectionFocusDiagnostics?.visibleViewportNodeCount,
    ).toBeGreaterThan(0);
    expect(sameSelectionFocusDiagnostics?.maxRenderedRadius).toBeGreaterThan(
      (sameSelectionShiftedDiagnostics?.maxRenderedRadius ??
        Number.POSITIVE_INFINITY) * 1.1,
    );
    expect(sameSelectionFocusDiagnostics?.overlapCount).toBe(0);
    expect(sameSelectionFocusDiagnostics?.targetGapViolationCount).toBe(0);
    const sameSelectionInteraction = await getGraphInteractionDiagnostics(page);
    const sameSelectionFocusSample =
      sameSelectionInteraction?.detailFocusSamples.at(-1);
    expect(sameSelectionFocusSample?.mode).toBe("detail-focus");
    expect(sameSelectionFocusSample?.maxRenderedRadius).toBeGreaterThan(
      (sameSelectionShiftedDiagnostics?.maxRenderedRadius ??
        Number.POSITIVE_INFINITY) * 1.1,
    );
    expect(sameSelectionInteraction?.dynamicGapSamples.at(-1)?.mode).toBe(
      "detail-focus",
    );

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
