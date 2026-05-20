import { test, expect } from '@playwright/test';

/**
 * E2E tests for the AVmatrix web UI — exploring view features.
 *
 * Requires:
 *   - avmatrix serve running on 127.0.0.1:4848 with at least one indexed repo
 *   - avmatrix-web dev server running on 127.0.0.1:5228
 *
 * Skipped when servers aren't available (CI without services, etc.).
 * Set E2E=1 to force-run even without the availability check.
 */

const BACKEND_URL = process.env.BACKEND_URL ?? 'http://127.0.0.1:4848';
const FRONTEND_URL = process.env.FRONTEND_URL ?? 'http://127.0.0.1:5228';
const POST_LOAD_STABILITY_WINDOW_MS = 30_000;

let firstRepoName = '';

type RuntimeDiagnostics = {
  graphConversion: {
    count: number;
    lastNodeCount: number;
    lastRelationshipCount: number;
  };
  visualScale: {
    nodeCount: number;
    minNodeSize: number;
    maxNodeSize: number;
    maxRenderedNodeSizeCap: number;
    structuralToLeafRatio: number;
    maxSizeByLabel: Record<string, number>;
  };
  layout: {
    starts: number;
    stops: number;
    isRunning: boolean;
    lastDurationBudgetMs: number;
    lastRunMs: number;
    lastNoverlapMs: number;
    lastReason: string;
  };
  heartbeat: {
    connects: number;
    reconnects: number;
  };
  reconnectBanner: {
    shows: number;
    visible: boolean;
  };
};

test.beforeAll(async () => {
  if (process.env.E2E) {
    const res = await fetch(`${BACKEND_URL}/api/repos`);
    const repos = await res.json();
    firstRepoName = repos[0]?.name ?? '';
    if (!firstRepoName) test.skip(true, 'No indexed repos - run avmatrix analyze first');
    return;
  }
  try {
    const [backendRes, frontendRes] = await Promise.allSettled([
      fetch(`${BACKEND_URL}/api/repos`),
      fetch(FRONTEND_URL),
    ]);
    if (
      backendRes.status === 'rejected' ||
      (backendRes.status === 'fulfilled' && !backendRes.value.ok)
    ) {
      test.skip(true, 'avmatrix serve not available on :4848');
      return;
    }
    if (
      frontendRes.status === 'rejected' ||
      (frontendRes.status === 'fulfilled' && !frontendRes.value.ok)
    ) {
      test.skip(true, 'Vite dev server not available on :5228');
      return;
    }
    // Check there's at least one indexed repo
    if (backendRes.status === 'fulfilled') {
      const repos = await backendRes.value.json();
      if (!repos.length) {
        test.skip(true, 'No indexed repos — run avmatrix analyze first');
        return;
      }
      firstRepoName = repos[0].name;
    }
  } catch {
    test.skip(true, 'servers not available');
  }
});

/**
 * Wait for the repo-scoped graph flow to complete.
 * The current architecture does not rely on a process-global active repo, so
 * E2E tests must pass the target repo context explicitly.
 */
async function waitForGraphLoaded(page: import('@playwright/test').Page) {
  await page.goto(
    `${FRONTEND_URL}/?server=${encodeURIComponent(BACKEND_URL)}&project=${encodeURIComponent(firstRepoName)}`,
  );

  const statusBar = page.getByRole('contentinfo');
  await expect(statusBar.getByText('Ready', { exact: true })).toBeVisible({ timeout: 45_000 });
  await expect(statusBar).toContainText(/nodes/, {
    timeout: 20_000,
  });
}

async function getRuntimeDiagnostics(
  page: import('@playwright/test').Page,
): Promise<RuntimeDiagnostics | null> {
  return page.evaluate(() => {
    const win = window as typeof window & {
      __AVMATRIX_WEB_DIAGNOSTICS__?: RuntimeDiagnostics;
    };
    return win.__AVMATRIX_WEB_DIAGNOSTICS__ ?? null;
  });
}

async function openFirstProcessModal(page: import('@playwright/test').Page) {
  const viewBtn = page
    .locator('[data-testid="process-view-button"]:not(:disabled)')
    .filter({ hasText: /^View$/ })
    .first();
  await expect(viewBtn).toBeAttached({ timeout: 30_000 });
  await viewBtn.scrollIntoViewIfNeeded();

  const processRow = viewBtn.locator('xpath=ancestor::*[@data-testid="process-row"][1]');
  await processRow.hover();
  await expect(viewBtn).toBeEnabled({ timeout: 5_000 });
  await viewBtn.click();
  await expect(page.locator('[data-testid="process-modal"]')).toBeVisible({ timeout: 60_000 });
}

test.describe('Server Connection & Graph Loading', () => {
  test('selects a repo from landing and loads graph', async ({ page }) => {
    await waitForGraphLoaded(page);
  });

  test('keeps connection stable after large graph load and layout window', async ({ page }) => {
    await waitForGraphLoaded(page);

    await expect(page.locator('[data-testid="server-reconnect-banner"]')).toHaveCount(0);
    await expect
      .poll(
        async () =>
          (await getRuntimeDiagnostics(page))?.graphConversion.count ?? 0,
        { timeout: 10_000 },
      )
      .toBeGreaterThan(0);
    await expect
      .poll(
        async () => {
          const diagnostics = await getRuntimeDiagnostics(page);
          if (!diagnostics) return 0;
          return diagnostics.layout.starts - diagnostics.layout.stops;
        },
        { timeout: 10_000, intervals: [500] },
      )
      .toBeGreaterThan(0);
    await page.waitForTimeout(POST_LOAD_STABILITY_WINDOW_MS);

    const diagnostics = await getRuntimeDiagnostics(page);
    expect(diagnostics?.heartbeat.connects).toBeGreaterThan(0);
    expect(diagnostics?.heartbeat.reconnects).toBe(0);
    expect(diagnostics?.reconnectBanner.shows).toBe(0);
    expect(diagnostics?.reconnectBanner.visible).toBe(false);
    expect(diagnostics?.graphConversion.lastNodeCount).toBeGreaterThan(0);
    expect(diagnostics?.graphConversion.lastRelationshipCount).toBeGreaterThan(0);
    expect(diagnostics?.layout.lastDurationBudgetMs).toBeGreaterThan(0);
    expect(diagnostics?.layout.lastNoverlapMs).toBeGreaterThanOrEqual(0);
  });
});

test.describe('Graph Dashboard Controls', () => {
  test('toggles uncommon node and edge types from the left dashboard', async ({ page }) => {
    test.slow();
    await waitForGraphLoaded(page);
    await page.getByRole('button', { name: 'Filters' }).click();

    const propertyControl = page.locator('button[title^="Property ("]').first();
    const accessesControl = page.locator('button[title^="Accesses ("]').first();
    const graphHealthSection = page.locator('[data-testid="graph-health-filter-section"]');
    await expect(propertyControl).toBeVisible({ timeout: 10_000 });
    await expect(accessesControl).toBeVisible({ timeout: 10_000 });
    await expect(graphHealthSection).toBeVisible({ timeout: 10_000 });
    await expect(graphHealthSection.getByText('Graph Health')).toBeVisible();
    await expect(graphHealthSection.getByText('Expected Isolation')).toBeVisible();
    await expect(graphHealthSection.getByText('Diagnostics')).toBeVisible();
    await expect(graphHealthSection.getByText('Confidence')).toBeVisible();

    const initialPropertyState = await propertyControl.getAttribute('aria-pressed');
    expect(initialPropertyState).toMatch(/^(true|false)$/);
    await propertyControl.click();
    await expect(propertyControl).toHaveAttribute(
      'aria-pressed',
      initialPropertyState === 'true' ? 'false' : 'true',
    );
    await propertyControl.click();
    await expect(propertyControl).toHaveAttribute('aria-pressed', initialPropertyState!);

    const initialAccessesState = await accessesControl.getAttribute('aria-pressed');
    expect(initialAccessesState).toMatch(/^(true|false)$/);
    await accessesControl.click();
    await expect(accessesControl).toHaveAttribute(
      'aria-pressed',
      initialAccessesState === 'true' ? 'false' : 'true',
    );
    await accessesControl.click();
    await expect(accessesControl).toHaveAttribute('aria-pressed', initialAccessesState!);

    const noIncomingControl = page.locator('button[title^="No incoming ("]').first();
    const testReasonControl = page.locator('button[title^="Test ("]').first();
    const unresolvedControl = page.locator('button[title^="Unresolved reference ("]').first();
    await expect(noIncomingControl).toBeVisible({ timeout: 10_000 });
    await expect(testReasonControl).toBeVisible({ timeout: 10_000 });
    await expect(unresolvedControl).toBeVisible({ timeout: 10_000 });

    await noIncomingControl.click();
    await expect(noIncomingControl).toHaveAttribute('aria-pressed', 'false');
    await noIncomingControl.click();
    await expect(noIncomingControl).toHaveAttribute('aria-pressed', 'true');

    await testReasonControl.click();
    await expect(testReasonControl).toHaveAttribute('aria-pressed', 'false');
    await testReasonControl.click();
    await expect(testReasonControl).toHaveAttribute('aria-pressed', 'true');

    await unresolvedControl.click();
    await expect(unresolvedControl).toHaveAttribute('aria-pressed', 'false');
    await unresolvedControl.click();
    await expect(unresolvedControl).toHaveAttribute('aria-pressed', 'true');

    await expect(page.locator('[title^="Legend node Property ("]').first()).toBeVisible();
    await expect(page.locator('[title^="Legend edge Accesses ("]').first()).toBeVisible();

    await page.getByRole('button', { name: 'Explorer' }).click();
    const packageJson = page.getByText('package.json').first();
    await expect(packageJson).toBeVisible({ timeout: 10_000 });
    await packageJson.click();
    const graphHealthDetail = page.getByTestId('graph-health-node-detail');
    await expect(graphHealthDetail).toBeVisible({ timeout: 10_000 });
    await expect(graphHealthDetail).toContainText('Next:');
  });

  test('keeps loaded graph node visual scale bounded', async ({ page }) => {
    await waitForGraphLoaded(page);

    await expect
      .poll(
        async () => (await getRuntimeDiagnostics(page))?.visualScale.nodeCount ?? 0,
        { timeout: 10_000 },
      )
      .toBeGreaterThan(0);

    const diagnostics = await getRuntimeDiagnostics(page);
    const visualScale = diagnostics!.visualScale;
    const expectedMaxSize = visualScale.nodeCount > 20_000 ? 3 : 4.5;
    const expectedRenderedMaxSize = visualScale.nodeCount > 20_000 ? 3 : 9;

    expect(visualScale.maxNodeSize).toBeLessThanOrEqual(expectedMaxSize);
    expect(visualScale.maxRenderedNodeSizeCap).toBe(expectedRenderedMaxSize);
    expect(visualScale.structuralToLeafRatio).toBeLessThanOrEqual(3);
    expect(visualScale.maxSizeByLabel.Package ?? 0).toBeLessThanOrEqual(1.5);
    expect(visualScale.maxSizeByLabel.Section ?? 0).toBeLessThanOrEqual(1);
  });
});

test.describe('My AI', () => {
  test('panel opens and agent initializes without error', async ({ page }) => {
    await waitForGraphLoaded(page);

    await page.getByRole('button', { name: 'My AI' }).click();
    await expect(page.getByText('Ask me anything')).toBeVisible({ timeout: 15_000 });

    const errorBanner = page.getByText('Database not ready');
    expect(await errorBanner.isVisible().catch(() => false)).toBe(false);
  });
});

test.describe('Processes Panel', () => {
  test('shows process list and View button works', async ({ page }) => {
    await waitForGraphLoaded(page);

    await page.getByRole('button', { name: 'My AI' }).click();
    await page.getByRole('button', { name: /^Processes\b/ }).click();

    await expect(page.locator('[data-testid="process-list-loaded"]')).toBeVisible({
      timeout: 15_000,
    });

    await openFirstProcessModal(page);
  });

  test('lightbulb highlights nodes in graph', async ({ page }) => {
    await waitForGraphLoaded(page);

    await page.getByRole('button', { name: 'My AI' }).click();
    await page.getByRole('button', { name: /^Processes\b/ }).click();
    await expect(page.locator('[data-testid="process-list-loaded"]')).toBeVisible({
      timeout: 15_000,
    });

    const processRow = page.locator('[data-testid="process-row"]').first();
    await expect(processRow).toBeVisible({ timeout: 10_000 });
    await processRow.hover();

    const lightbulb = processRow.locator('[data-testid="process-highlight-button"]');
    await lightbulb.waitFor({ state: 'visible', timeout: 5_000 });
    await lightbulb.click();
    await expect(lightbulb).toHaveAttribute('title', 'Click to remove highlight from graph', {
      timeout: 15_000,
    });
    await expect(processRow).toHaveClass(/border-border-strong/, { timeout: 15_000 });
  });
});

test.describe('Turn Off All Highlights', () => {
  test('selecting a node dims others, button clears it', async ({ page }) => {
    await waitForGraphLoaded(page);

    await expect(page.locator('canvas').first()).toBeVisible({ timeout: 10_000 });

    const fileItem = page.getByText('package.json').first();
    await expect(fileItem).toBeVisible({ timeout: 10_000 });
    await fileItem.click();

    const highlightToggle = page.locator('[data-testid="ai-highlights-toggle"]');
    await expect(highlightToggle).toHaveAttribute('title', 'Turn off AI-driven highlights', {
      timeout: 5_000,
    });

    await highlightToggle.click();
    await expect(highlightToggle).toHaveAttribute('title', 'Turn on AI-driven highlights', {
      timeout: 5_000,
    });
  });
});
