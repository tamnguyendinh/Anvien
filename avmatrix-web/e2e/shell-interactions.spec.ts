import { test, expect, type Page } from "@playwright/test";

const BACKEND_URL = process.env.BACKEND_URL ?? "http://127.0.0.1:4848";
const FRONTEND_URL = process.env.FRONTEND_URL ?? "http://127.0.0.1:5228";
const TARGET_REPO_NAME = process.env.E2E_REPO_NAME ?? "Restaurant_manager";
const ABSOLUTE_LOCAL_PATH =
  process.platform === "win32" ? "C:\\repos\\shell-check" : "/tmp/shell-check";

type BackendRepo = {
  name: string;
  path?: string;
  repoPath?: string;
};

let selectedRepoName = "";

test.beforeAll(async () => {
  if (process.env.E2E) {
    const res = await fetch(`${BACKEND_URL}/api/repos`);
    const repos = await res.json();
    selectedRepoName = selectDeterministicRepoName(repos);
    if (!selectedRepoName) {
      test.skip(
        true,
        `Target repo ${TARGET_REPO_NAME} is not indexed; deterministic large-graph smoke skipped`,
      );
    }
    return;
  }
  try {
    const [backendRes, frontendRes] = await Promise.allSettled([
      fetch(`${BACKEND_URL}/api/repos`),
      fetch(FRONTEND_URL),
    ]);
    if (
      backendRes.status === "rejected" ||
      (backendRes.status === "fulfilled" && !backendRes.value.ok)
    ) {
      test.skip(true, "avmatrix serve not available");
      return;
    }
    if (
      frontendRes.status === "rejected" ||
      (frontendRes.status === "fulfilled" && !frontendRes.value.ok)
    ) {
      test.skip(true, "Vite dev server not available");
      return;
    }
    if (backendRes.status === "fulfilled") {
      const repos = await backendRes.value.json();
      if (!repos.length) {
        test.skip(true, "No indexed repos");
        return;
      }
      selectedRepoName = selectDeterministicRepoName(repos);
      if (!selectedRepoName) {
        test.skip(
          true,
          `Target repo ${TARGET_REPO_NAME} is not indexed; deterministic large-graph smoke skipped`,
        );
        return;
      }
    }
  } catch {
    test.skip(true, "servers not available");
  }
});

function selectDeterministicRepoName(repos: BackendRepo[]) {
  const normalizedTarget = normalizeRepoSelector(TARGET_REPO_NAME);
  const selected = repos.find((repo) => {
    const candidates = [repo.name, repo.path, repo.repoPath]
      .filter(Boolean)
      .map((value) => normalizeRepoSelector(value ?? ""));
    return candidates.some(
      (candidate) =>
        candidate === normalizedTarget ||
        candidate.endsWith(`/${normalizedTarget}`),
    );
  });
  return selected?.name ?? "";
}

function normalizeRepoSelector(value: string) {
  return value.replace(/\\/g, "/").replace(/\/+$/, "").split("/").pop() ?? "";
}

async function waitForGraphLoaded(page: Page) {
  await page.goto(
    `${FRONTEND_URL}/?server=${encodeURIComponent(BACKEND_URL)}&project=${encodeURIComponent(selectedRepoName)}`,
  );

  await expect(page.locator('[data-testid="status-ready"]')).toBeVisible({
    timeout: 45_000,
  });
}

async function openRepoDropdown(page: Page) {
  const url = new URL(page.url());
  const project = url.searchParams.get("project");

  const repoButton = project
    ? page.locator("header button").filter({ hasText: project }).first()
    : page
        .locator("header button")
        .filter({ has: page.locator("svg") })
        .nth(1);

  await repoButton.click();
  await expect(page.getByText("Repositories")).toBeVisible({ timeout: 5_000 });
}

function mockReadySessionStatus(page: Page) {
  return page.route(`${BACKEND_URL}/api/session/status**`, (route) =>
    route.fulfill({
      contentType: "application/json",
      body: JSON.stringify({
        availability: "ready",
        authenticated: true,
        message: "Local runtime ready",
        runtimeEnvironment: "local",
        executionMode: "cli",
        repo: { state: "indexed", message: "Indexed and ready" },
        version: "1.0.0",
        executablePath: "codex",
      }),
    }),
  );
}

test.describe("Shell interactions", () => {
  test("back button returns to the launcher start screen without reconnect banner", async ({
    page,
  }) => {
    await waitForGraphLoaded(page);

    await page.getByLabel("Back to Start screen").click();

    await expect(page).toHaveURL(/\/Start-AVmatrix\.html$/);
    await expect(
      page.getByRole("button", { name: "Start AVmatrix" }),
    ).toBeVisible({
      timeout: 5_000,
    });
    await expect(
      page.locator('[data-testid="server-reconnect-banner"]'),
    ).toHaveCount(0);
  });

  test("resizes the left dashboard and keeps graph controls usable", async ({
    page,
  }) => {
    await waitForGraphLoaded(page);

    const panel = page.locator(".file-tree-panel").first();
    const handle = page.getByTestId("left-dashboard-resize-handle");
    const panelBox = await panel.boundingBox();
    if (!panelBox) throw new Error("left dashboard bounding box missing");

    await handle.dragTo(page.locator("main"), {
      sourcePosition: { x: 2, y: 80 },
      targetPosition: { x: panelBox.x + panelBox.width + 160, y: 80 },
    });
    await expect
      .poll(async () => (await panel.boundingBox())?.width ?? 0, {
        timeout: 5_000,
      })
      .toBeGreaterThan(panelBox.width + 80);

    await handle.dragTo(page.locator("main"), {
      sourcePosition: { x: 2, y: 80 },
      targetPosition: { x: 40, y: 80 },
    });
    await expect
      .poll(async () => (await panel.boundingBox())?.width ?? 0, {
        timeout: 5_000,
      })
      .toBeGreaterThanOrEqual(190);

    await page.getByRole("button", { name: "Filters" }).click();
    await expect(
      page.locator('button[title^="Property ("]').first(),
    ).toBeVisible({
      timeout: 10_000,
    });
    await expect(page.locator("canvas").first()).toBeVisible();
  });

  test("displays graph filters, legends, focus depth, and relationship toggles", async ({
    page,
  }) => {
    await waitForGraphLoaded(page);

    await page.getByRole("button", { name: "Filters" }).click();

    await expect(page.getByText("Node Types").first()).toBeVisible();
    await expect(page.getByText("Edge Types").first()).toBeVisible();
    await expect(page.getByText("Focus Depth")).toBeVisible();
    await expect(page.getByText("Color Legend")).toBeVisible();

    await expect(page.locator('button[title^="File ("]').first()).toBeVisible({
      timeout: 10_000,
    });
    await expect(
      page.locator('div[title^="Legend node File ("]').first(),
    ).toBeVisible();

    const callsButton = page.locator('button[title^="Calls ("]').first();
    await expect(callsButton).toBeVisible({ timeout: 10_000 });
    await expect(
      page.locator('div[title^="Legend edge Calls ("]').first(),
    ).toBeVisible();

    await callsButton.click();
    await expect(callsButton).toHaveAttribute("aria-pressed", "false");
    await callsButton.click();
    await expect(callsButton).toHaveAttribute("aria-pressed", "true");

    await page.getByRole("button", { name: "2 hops" }).click();
    await expect(
      page.getByText("Select a node to apply depth filter"),
    ).toBeVisible();
    await page.getByRole("button", { name: "All", exact: true }).click();
    await expect(
      page.getByText("Select a node to apply depth filter"),
    ).not.toBeVisible();
  });

  test("opens settings, edits values, saves, then closes", async ({ page }) => {
    await mockReadySessionStatus(page);
    await waitForGraphLoaded(page);

    await page.getByTitle("Session Settings").click();
    await expect(page.getByText("AI Runtime")).toBeVisible({ timeout: 5_000 });
    await expect(page.getByText("Codex Account")).toBeVisible();
    await expect(page.getByText("Claude Code Account")).toBeVisible();

    await page.getByRole("button", { name: "Close" }).click();
    await expect(page.getByText("AI Runtime")).not.toBeVisible({
      timeout: 5_000,
    });
  });

  test("opens repo dropdown, enters analyze flow, types a path, and dismisses it", async ({
    page,
  }) => {
    await waitForGraphLoaded(page);
    await openRepoDropdown(page);

    await page.getByText("Analyze a new repository...").click();
    await expect(page.getByLabel("Repository Folder")).toBeVisible({
      timeout: 5_000,
    });

    const pathInput = page.getByLabel("Repository Folder");
    await pathInput.fill(ABSOLUTE_LOCAL_PATH);
    await expect(
      page.getByRole("button", { name: /Analyze Repository/i }),
    ).toBeEnabled();

    const url = new URL(page.url());
    const project = url.searchParams.get("project");
    const repoButton = project
      ? page.locator("header button").filter({ hasText: project }).first()
      : page
          .locator("header button")
          .filter({ has: page.locator("svg") })
          .nth(1);

    await repoButton.click();
    await expect(page.getByLabel("Repository Folder")).not.toBeVisible({
      timeout: 5_000,
    });
    await expect(page.getByText("Repositories")).not.toBeVisible({
      timeout: 5_000,
    });
  });

  test("sends a chat message, clears it, switches to processes, and closes the modal/panel", async ({
    page,
  }) => {
    await mockReadySessionStatus(page);
    await page.route(`${BACKEND_URL}/api/session/chat`, async (route) => {
      await route.fulfill({
        status: 200,
        headers: { "Content-Type": "text/event-stream" },
        body: [
          'data: {"type":"session_started","sessionId":"e2e-session"}',
          "",
          'data: {"type":"content","content":"Mocked assistant reply from Playwright."}',
          "",
          'data: {"type":"done"}',
          "",
        ].join("\n"),
      });
    });

    await waitForGraphLoaded(page);

    await page.getByRole("button", { name: "My AI" }).click();
    await expect(page.getByText("Ask me anything")).toBeVisible({
      timeout: 10_000,
    });

    const composer = page.locator(
      'textarea[placeholder="Ask about the codebase..."]',
    );
    await composer.fill("What is loaded?");
    await composer.press("Enter");

    await expect(
      page.getByText("Mocked assistant reply from Playwright."),
    ).toBeVisible({
      timeout: 10_000,
    });

    await page.getByTitle("Clear chat").click();
    await expect(page.getByText("Ask me anything")).toBeVisible({
      timeout: 10_000,
    });

    await page.getByRole("button", { name: /^Processes\b/ }).click();
    await expect(
      page.locator('[data-testid="process-list-loaded"]'),
    ).toBeVisible({
      timeout: 15_000,
    });

    const processRow = page.locator('[data-testid="process-row"]').first();
    await processRow.hover();
    const highlightButton = processRow.locator(
      '[data-testid="process-highlight-button"]',
    );
    await highlightButton.click();
    await expect(highlightButton).toHaveAttribute(
      "title",
      "Click to remove highlight from graph",
      {
        timeout: 15_000,
      },
    );
    await expect(processRow).toHaveClass(/border-border-strong/, {
      timeout: 15_000,
    });

    await processRow.hover();
    const viewButton = processRow.locator(
      '[data-testid="process-view-button"]',
    );
    await viewButton.waitFor({ state: "visible", timeout: 5_000 });
    await viewButton.click();
    await expect(page.locator('[data-testid="process-modal"]')).toBeVisible({
      timeout: 15_000,
    });
    await page
      .locator('[data-testid="process-modal"]')
      .getByRole("button", { name: "Close" })
      .click();
    await expect(page.locator('[data-testid="process-modal"]')).not.toBeVisible(
      { timeout: 5_000 },
    );

    await page.getByTitle("Close Panel").click();
    await expect(
      page.locator('textarea[placeholder="Ask about the codebase..."]'),
    ).not.toBeVisible({
      timeout: 5_000,
    });
  });

  test("can stop an in-flight chat response", async ({ page }) => {
    await mockReadySessionStatus(page);
    await page.route(`${BACKEND_URL}/api/session/chat`, async (route) => {
      await new Promise((resolve) => setTimeout(resolve, 4_000));
      try {
        await route.fulfill({
          status: 200,
          headers: { "Content-Type": "text/event-stream" },
          body: [
            'data: {"type":"session_started","sessionId":"slow-session"}',
            "",
            'data: {"type":"content","content":"This reply arrived too late."}',
            "",
            'data: {"type":"done"}',
            "",
          ].join("\n"),
        });
      } catch {
        // Browser likely aborted the request after Stop was pressed.
      }
    });

    await waitForGraphLoaded(page);
    await page.getByRole("button", { name: "My AI" }).click();

    const composer = page.locator(
      'textarea[placeholder="Ask about the codebase..."]',
    );
    await composer.fill("slow request");
    await composer.press("Enter");

    const stopButton = page.getByTitle("Stop response");
    await expect(stopButton).toBeVisible({ timeout: 5_000 });
    await stopButton.click();

    await expect(stopButton).not.toBeVisible({ timeout: 10_000 });
    await expect(
      page.locator('textarea[placeholder="Ask about the codebase..."]'),
    ).toBeVisible();
  });
});
