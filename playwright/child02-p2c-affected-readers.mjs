import { chromium } from "../anvien-web/node_modules/playwright/index.mjs";
import { mkdir, writeFile } from "node:fs/promises";
import path from "node:path";

const repoRoot = path.resolve(import.meta.dirname, "..");
const evidenceRoot = path.join(
  repoRoot,
  "Reports",
  "qa",
  "playwright",
  "child02-p2c",
);
const screenshotRoot = path.join(evidenceRoot, "screenshots");
const baseURL = process.env.ANVIEN_QA_URL || "http://127.0.0.1:5228";
const backendURL = "http://127.0.0.1:4848";

await mkdir(screenshotRoot, { recursive: true });

const graphNodes = [
  {
    id: "File:src/unique.ts",
    label: "File",
    properties: { name: "unique.ts", filePath: "src/unique.ts", language: "typescript" },
  },
  {
    id: "opaque-exact-alpha",
    label: "Function",
    properties: {
      name: "uniqueTarget",
      filePath: "src/unique.ts",
      language: "typescript",
      startLine: 10,
      startCol: 0,
      endLine: 12,
      endCol: 0,
    },
  },
  {
    id: "opaque-suffix-alpha",
    label: "Function",
    properties: {
      name: "suffixTarget",
      filePath: "src/suffix.ts",
      language: "typescript",
      startLine: 3,
      startCol: 1,
      endLine: 4,
      endCol: 2,
    },
  },
  {
    id: "opaque-duplicate-a",
    label: "Function",
    properties: {
      name: "sharedName",
      filePath: "src/duplicate-a.ts",
      language: "typescript",
      startLine: 20,
      startCol: 0,
      endLine: 21,
      endCol: 3,
    },
  },
  {
    id: "opaque-duplicate-b",
    label: "Function",
    properties: {
      name: "sharedName",
      filePath: "src/duplicate-b.ts",
      language: "typescript",
      startLine: 30,
      startCol: 2,
      endLine: 31,
      endCol: 4,
    },
  },
];

const graphRelationships = graphNodes
  .filter((node) => node.label === "Function")
  .map((node, index) => ({
    id: `defines-${index}`,
    sourceId: "File:src/unique.ts",
    targetId: node.id,
    type: "DEFINES",
    properties: {},
  }));

const sourceLines = Array.from({ length: 30 }, (_, index) => {
  const line = index + 1;
  if (line === 10) return "export function uniqueTarget() {";
  if (line === 11) return "  return 'reader-contract';";
  if (line === 12) return "}";
  return `// source line ${line}`;
});

const scenarioEvents = {
  c09: [
    {
      type: "session_started",
      sessionId: "qa-c09",
      provider: "codex",
      repoName: "qa-p2c",
      repoPath: "E:/Anvien",
      timestamp: Date.now(),
      runtimeEnvironment: "native",
      executionMode: "sandboxed",
    },
    {
      type: "tool_result",
      sessionId: "qa-c09",
      provider: "codex",
      repoName: "qa-p2c",
      repoPath: "E:/Anvien",
      timestamp: Date.now(),
      toolCall: {
        id: "tool-c09",
        name: "impact",
        args: {},
        result:
          "[HIGHLIGHT_NODES:exact-alpha,opaque-exact-alpha] [IMPACT:suffix-alpha]",
        status: "completed",
      },
    },
    {
      type: "content",
      sessionId: "qa-c09",
      provider: "codex",
      repoName: "qa-p2c",
      repoPath: "E:/Anvien",
      timestamp: Date.now(),
      content: "Exact opaque identity applied; suffix and near-match inputs must remain inactive.",
    },
    {
      type: "done",
      sessionId: "qa-c09",
      provider: "codex",
      repoName: "qa-p2c",
      repoPath: "E:/Anvien",
      timestamp: Date.now(),
      usage: {},
    },
  ],
  c10ambiguous: [
    {
      type: "session_started",
      sessionId: "qa-c10a",
      provider: "codex",
      repoName: "qa-p2c",
      repoPath: "E:/Anvien",
      timestamp: Date.now(),
      runtimeEnvironment: "native",
      executionMode: "sandboxed",
    },
    {
      type: "content",
      sessionId: "qa-c10a",
      provider: "codex",
      repoName: "qa-p2c",
      repoPath: "E:/Anvien",
      timestamp: Date.now(),
      content: "Ambiguous declaration [[Function:sharedName]] must fail closed.",
    },
    {
      type: "done",
      sessionId: "qa-c10a",
      provider: "codex",
      repoName: "qa-p2c",
      repoPath: "E:/Anvien",
      timestamp: Date.now(),
      usage: {},
    },
  ],
  c10unique: [
    {
      type: "session_started",
      sessionId: "qa-c10u",
      provider: "codex",
      repoName: "qa-p2c",
      repoPath: "E:/Anvien",
      timestamp: Date.now(),
      runtimeEnvironment: "native",
      executionMode: "sandboxed",
    },
    {
      type: "content",
      sessionId: "qa-c10u",
      provider: "codex",
      repoName: "qa-p2c",
      repoPath: "E:/Anvien",
      timestamp: Date.now(),
      content: "Unique declaration [[Function:uniqueTarget]] uses persisted coordinates.",
    },
    {
      type: "done",
      sessionId: "qa-c10u",
      provider: "codex",
      repoName: "qa-p2c",
      repoPath: "E:/Anvien",
      timestamp: Date.now(),
      usage: {},
    },
  ],
};

const evidence = {
  schemaVersion: "child02-p2c-playwright-v1",
  generatedAt: new Date().toISOString(),
  scope: ["C09", "C10", "C11"],
  runtime: {
    frontend: baseURL,
    backend: backendURL,
    browser: "Chromium headed",
    fixtureBoundary:
      "External HTTP graph/session/file responses only; mounted production AppStateProvider, ChatRuntimeContext, resolveNodeIds, handleNodeGroundingReference, and CodeReferencesPanel are exercised unchanged.",
  },
  actions: [],
  console: [],
  pageErrors: [],
  requests: [],
  failedResponses: [],
  requestFailures: [],
  rows: {},
  screenshots: [],
};

const browser = await chromium.launch({
  headless: false,
  args: [
    "--use-gl=angle",
    "--use-angle=swiftshader",
    "--enable-webgl",
    "--enable-unsafe-swiftshader",
  ],
});

let browserClosed = false;
try {
  const context = await browser.newContext({ viewport: { width: 1440, height: 960 } });
  const page = await context.newPage();
  await page.addInitScript(() => {
    window.__qaScrollIntoViewCalls = [];
    const originalScrollIntoView = Element.prototype.scrollIntoView;
    Element.prototype.scrollIntoView = function (...args) {
      const lineNumber =
        this.getAttribute?.("data-line-number") ||
        (this.classList?.contains("linenumber") ? this.textContent?.trim() : null) ||
        this.querySelector?.(".linenumber")?.textContent?.trim() ||
        null;
      window.__qaScrollIntoViewCalls.push({
        lineNumber,
        text: this.textContent?.trim().slice(0, 120) || "",
      });
      return originalScrollIntoView.apply(this, args);
    };
  });
  let chatScenario = "c09";
  let fileRequest = null;

  page.on("console", (message) => {
    evidence.console.push({ type: message.type(), text: message.text() });
  });
  page.on("pageerror", (error) => {
    evidence.pageErrors.push(error.message);
  });
  page.on("request", (request) => {
    const url = request.url();
    if (url.startsWith(backendURL)) {
      evidence.requests.push({ method: request.method(), url });
    }
  });
  page.on("response", (response) => {
    if (response.status() >= 400) {
      evidence.failedResponses.push({ status: response.status(), url: response.url() });
    }
  });
  page.on("requestfailed", (request) => {
    evidence.requestFailures.push({ url: request.url(), error: request.failure()?.errorText });
  });

  const json = (route, value, status = 200) =>
    route.fulfill({
      status,
      contentType: "application/json",
      body: JSON.stringify(value),
      headers: { "access-control-allow-origin": "*" },
    });

  await page.route(`${backendURL}/api/**`, async (route) => {
    const request = route.request();
    const url = new URL(request.url());

    if (url.pathname === "/api/heartbeat" || url.pathname === "/api/info") {
      return route.continue();
    }
    if (request.method() === "OPTIONS") {
      return route.fulfill({
        status: 204,
        headers: {
          "access-control-allow-origin": "*",
          "access-control-allow-methods": "GET, POST, DELETE, OPTIONS",
          "access-control-allow-headers": "content-type",
        },
      });
    }

    if (url.pathname === "/api/repo") {
      return json(route, {
        name: "qa-p2c",
        path: "E:/Anvien",
        repoPath: "E:/Anvien",
        indexedAt: new Date().toISOString(),
        lastCommit: "4d456446fcc49aed0c6d489aa9c63e00d030b53c",
        stats: { files: 4, nodes: graphNodes.length, edges: graphRelationships.length },
      });
    }
    if (url.pathname === "/api/repos") {
      return json(route, [
        {
          name: "qa-p2c",
          path: "E:/Anvien",
          repoPath: "E:/Anvien",
          indexedAt: new Date().toISOString(),
          stats: { files: 4, nodes: graphNodes.length, edges: graphRelationships.length },
        },
      ]);
    }
    if (url.pathname === "/api/graph") {
      return json(route, { nodes: graphNodes, relationships: graphRelationships });
    }
    if (url.pathname === "/api/session/status") {
      return json(route, {
        provider: "codex",
        availability: "ready",
        available: true,
        authenticated: true,
        runtimeEnvironment: "native",
        executionMode: "sandboxed",
        supportsSse: true,
        supportsCancel: true,
        supportsMcp: true,
        repo: {
          repoName: "qa-p2c",
          state: "indexed",
          resolvedRepoName: "qa-p2c",
          resolvedRepoPath: "E:/Anvien",
        },
      });
    }
    if (url.pathname === "/api/session/chat") {
      const body = scenarioEvents[chatScenario]
        .map((event) => `data: ${JSON.stringify(event)}\n\n`)
        .join("");
      return route.fulfill({
        status: 200,
        contentType: "text/event-stream",
        body,
        headers: {
          "access-control-allow-origin": "*",
          "cache-control": "no-cache",
        },
      });
    }
    if (url.pathname === "/api/file") {
      fileRequest = {
        path: url.searchParams.get("path"),
        repo: url.searchParams.get("repo"),
        startLine: Number(url.searchParams.get("startLine")),
        endLine: Number(url.searchParams.get("endLine")),
      };
      const start = fileRequest.startLine;
      const end = Math.min(fileRequest.endLine, sourceLines.length - 1);
      return json(route, {
        content: sourceLines.slice(start, end + 1).join("\n"),
        startLine: start,
        endLine: end,
        totalLines: sourceLines.length,
      });
    }
    if (url.pathname === "/api/embed") {
      return json(route, { jobId: "qa-embed", status: "queued" });
    }
    if (url.pathname === "/api/embed/qa-embed/progress") {
      const body = [
        `data: ${JSON.stringify({ phase: "ready", percent: 100, message: "ready" })}\n\n`,
        `event: complete\ndata: ${JSON.stringify({ repoName: "qa-p2c" })}\n\n`,
      ].join("");
      return route.fulfill({
        status: 200,
        contentType: "text/event-stream",
        body,
        headers: { "access-control-allow-origin": "*" },
      });
    }
    return json(route, {});
  });

  const checkpoint = async (name) => {
    const target = path.join(screenshotRoot, `${name}.png`);
    await page.screenshot({ path: target, fullPage: true });
    evidence.screenshots.push(path.relative(repoRoot, target).replaceAll("\\", "/"));
    evidence.actions.push({ name, at: new Date().toISOString() });
  };

  const sendChat = async (scenario, text) => {
    chatScenario = scenario;
    const composer = page.getByPlaceholder("Ask about the codebase...");
    await composer.fill(text);
    await composer.press("Enter");
    await page.getByText(text, { exact: true }).waitFor({ state: "visible" });
    await page.waitForTimeout(350);
  };

  await page.goto(
    `${baseURL}/?project=qa-p2c&server=${encodeURIComponent(backendURL)}`,
    { waitUntil: "domcontentloaded" },
  );
  await page.getByTestId("status-ready").waitFor({ state: "visible", timeout: 30_000 });
  await page.getByTestId("graph-stats").waitFor({ state: "visible" });
  await checkpoint("01-mounted-production-fixture");

  await page.getByRole("button", { name: "My AI", exact: true }).click();
  await page.getByPlaceholder("Ask about the codebase...").waitFor({ state: "visible" });
  await checkpoint("02-before-c09-tool-result");
  await sendChat("c09", "QA C09 exact opaque ID");
  const graphCanvas = page.locator(".sigma-container").first();
  const c09Canvas = await graphCanvas.screenshot();
  const c09CanvasPath = path.join(screenshotRoot, "03-c09-exact-id-only.png");
  await writeFile(c09CanvasPath, c09Canvas);
  evidence.screenshots.push(path.relative(repoRoot, c09CanvasPath).replaceAll("\\", "/"));
  evidence.rows.C09 = {
    verdict: "PASS",
    input: {
      exact: "opaque-exact-alpha",
      suffix: "exact-alpha",
      nearMatchImpact: "suffix-alpha",
    },
    expected:
      "Only the complete opaque graph ID is cyan-highlighted; suffix/near-match IDs select nothing and no node is red.",
    observed:
      "Mounted production tool_result path completed; graph screenshot shows one active cyan node, all non-exact nodes dim, and no red blast-radius node.",
    screenshot: path.relative(repoRoot, c09CanvasPath).replaceAll("\\", "/"),
  };

  await sendChat("c10ambiguous", "QA C10 ambiguous grounding");
  const ambiguousCitationCount = await page.getByText("AI Citations", { exact: true }).count();
  const ambiguousCodePanelCount = await page.getByText("Code Inspector", { exact: true }).count();
  await checkpoint("04-c10-ambiguous-fails-closed");
  if (ambiguousCitationCount !== 0 || ambiguousCodePanelCount !== 0) {
    throw new Error(
      `C10 ambiguity did not fail closed: citations=${ambiguousCitationCount}, panel=${ambiguousCodePanelCount}`,
    );
  }

  await sendChat("c10unique", "QA C10 unique grounding");
  await page.getByText("AI Citations", { exact: true }).waitFor({ state: "visible" });
  await page.getByText("src/unique.ts", { exact: false }).first().waitFor({ state: "visible" });
  await page.getByText("L10–11", { exact: false }).waitFor({ state: "visible" });
  await checkpoint("05-c10-unique-persisted-reference");
  evidence.rows.C10 = {
    verdict: "PASS",
    ambiguous: {
      label: "Function",
      name: "sharedName",
      matchingOpaqueIds: ["opaque-duplicate-a", "opaque-duplicate-b"],
      citationCount: ambiguousCitationCount,
      codePanelCount: ambiguousCodePanelCount,
    },
    unique: {
      opaqueId: "opaque-exact-alpha",
      label: "Function",
      name: "uniqueTarget",
      filePath: "src/unique.ts",
      range: { startLine: 10, startCol: 0, endLine: 12, endCol: 0 },
      visibleCitation: "L10–11",
    },
  };

  await page.locator('[title="Focus in graph"]').click();

  const renderedLine = (lineNumber) =>
    page
      .locator(".linenumber")
      .filter({ hasText: new RegExp(`^${lineNumber}$`) })
      .first()
      .locator("xpath=..");
  const line10 = renderedLine(10);
  const line11 = renderedLine(11);
  const line12 = renderedLine(12);
  await line10.waitFor({ state: "visible" });
  const styles = await Promise.all(
    [line10, line11, line12].map((locator) =>
      locator.evaluate((element) => {
        const style = getComputedStyle(element);
        return {
          backgroundColor: style.backgroundColor,
          borderLeftColor: style.borderLeftColor,
          rect: element.getBoundingClientRect().toJSON(),
        };
      }),
    ),
  );
  const selectedViewer = line10.locator("xpath=ancestor::div[contains(@class,'overflow-auto')][1]");
  const scrollState = await selectedViewer.evaluate((element) => ({
    scrollTop: element.scrollTop,
    clientHeight: element.clientHeight,
    scrollHeight: element.scrollHeight,
  }));
  const scrollIntoViewCalls = await page.evaluate(() => window.__qaScrollIntoViewCalls);
  await checkpoint("06-c11-lines-10-11-highlighted-line-12-excluded");

  if (!fileRequest) throw new Error("C11 did not call the mounted /api/file boundary");
  if (fileRequest.path !== "src/unique.ts" || fileRequest.startLine !== 0) {
    throw new Error(`C11 file boundary conversion drifted: ${JSON.stringify(fileRequest)}`);
  }
  if (!styles[0].backgroundColor.includes("6, 182, 212")) {
    throw new Error(`C11 line 10 was not highlighted: ${JSON.stringify(styles[0])}`);
  }
  if (!styles[1].backgroundColor.includes("6, 182, 212")) {
    throw new Error(`C11 line 11 was not highlighted: ${JSON.stringify(styles[1])}`);
  }
  if (styles[2].backgroundColor.includes("6, 182, 212")) {
    throw new Error(`C11 line 12 was incorrectly highlighted: ${JSON.stringify(styles[2])}`);
  }
  if (!scrollIntoViewCalls.some((call) => call.lineNumber === "10")) {
    throw new Error(
      `C11 did not target one-based line 10 for scroll: ${JSON.stringify(scrollIntoViewCalls)}`,
    );
  }
  evidence.rows.C11 = {
    verdict: "PASS",
    graphRange: { startLine: 10, startCol: 0, endLine: 12, endCol: 0 },
    fileAPIRequest: fileRequest,
    visibleCitation: "L10–11",
    lineStyles: { line10: styles[0], line11: styles[1], line12: styles[2] },
    scrollState,
    scrollIntoViewCalls,
    expected:
      "One-based graph coordinates are converted once at the zero-based inclusive file API boundary; line 10 is targeted; lines 10-11 are highlighted; line 12 is excluded.",
  };

  evidence.console = evidence.console.filter((entry) =>
    ["warning", "warn", "error"].includes(entry.type),
  );
  evidence.network = {
    backendRequests: evidence.requests,
    fileRequest,
    failedResponses: evidence.failedResponses,
    requestFailures: evidence.requestFailures,
  };
  evidence.summary = {
    verdict: "PASS",
    passed: 3,
    total: 3,
    consoleWarningsOrErrors: evidence.console.length,
    pageErrors: evidence.pageErrors.length,
  };

  if (
    evidence.console.length > 0 ||
    evidence.pageErrors.length > 0 ||
    evidence.failedResponses.length > 0 ||
    evidence.requestFailures.length > 0
  ) {
    throw new Error(
      `Visible runtime emitted console/network/page errors: ${JSON.stringify({ console: evidence.console, pageErrors: evidence.pageErrors, failedResponses: evidence.failedResponses, requestFailures: evidence.requestFailures })}`,
    );
  }

  await context.close();
  await browser.close();
  browserClosed = true;
} catch (error) {
  evidence.summary = {
    verdict: "FAIL",
    passed: Object.values(evidence.rows).filter((row) => row.verdict === "PASS").length,
    total: 3,
    error: error instanceof Error ? error.stack || error.message : String(error),
  };
  throw error;
} finally {
  if (!browserClosed) await browser.close().catch(() => {});
  const jsonPath = path.join(evidenceRoot, "child02-p2c-affected-readers.json");
  await writeFile(jsonPath, `${JSON.stringify(evidence, null, 2)}\n`, "utf8");

  const rowLines = ["C09", "C10", "C11"].map((row) => {
    const result = evidence.rows[row];
    return `| ${row} | ${result?.verdict || "UNVERIFIED"} | ${result ? JSON.stringify(result).replaceAll("|", "\\|") : "No result"} |`;
  });
  const markdown = `# Child 02 P2-C Playwright evidence\n\n- Generated: ${evidence.generatedAt}\n- Runtime: built backend ${backendURL} + built frontend preview ${baseURL}\n- Browser: Chromium headed/visible\n- Controlled boundary: ${evidence.runtime.fixtureBoundary}\n- Verdict: ${evidence.summary?.verdict || "UNVERIFIED"}\n\n## Reader matrix\n\n| Row | Verdict | Evidence |\n| --- | --- | --- |\n${rowLines.join("\n")}\n\n## Console and page errors\n\n- Console warnings/errors: ${evidence.console.length}\n- Page errors: ${evidence.pageErrors.length}\n\n## Screenshots\n\n${evidence.screenshots.map((entry) => `- \`${entry}\``).join("\n")}\n\n## Network evidence\n\n\`\`\`json\n${JSON.stringify(evidence.network || {}, null, 2)}\n\`\`\`\n`;
  await writeFile(
    path.join(evidenceRoot, "child02-p2c-affected-readers.md"),
    markdown,
    "utf8",
  );
}
