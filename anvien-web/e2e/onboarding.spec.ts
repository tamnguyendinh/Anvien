import { test, expect } from '@playwright/test';

/**
 * E2E tests for the onboarding and analysis user flows.
 *
 * These tests cover:
 *   - Flow 1: Start screen and neutral runtime connection state
 *   - Flow 2: Analyze form when server has zero repos
 *   - Flow 3: Auto-connect when server has repos
 *   - Flow 4: Repo dropdown in exploring view
 *
 * Most tests mock the backend at the network level so they don't
 * require a live anvien server.
 */

const BACKEND_URL = process.env.BACKEND_URL ?? 'http://127.0.0.1:4848';
const FRONTEND_URL = process.env.FRONTEND_URL ?? 'http://127.0.0.1:5228';
const ABSOLUTE_LOCAL_PATH = process.platform === 'win32' ? 'C:\\repos\\demo' : '/tmp/demo';

let firstRepoName = '';

async function launchFromStartScreen(page: import('@playwright/test').Page) {
  await page.goto('/');
  await expect(page.getByRole('button', { name: 'Start Anvien' })).toBeVisible({
    timeout: 10_000,
  });
  await expect(page.getByRole('button', { name: 'RESET RUNTIME' })).toBeVisible();
  await expect(page.getByRole('button', { name: 'User Guide' })).toBeVisible();
  await page.getByRole('button', { name: 'Start Anvien' }).click();
}

async function enterExploringView(page: import('@playwright/test').Page) {
  await page.goto(
    `${FRONTEND_URL}/?server=${encodeURIComponent(BACKEND_URL)}&project=${encodeURIComponent(firstRepoName)}`,
  );

  // Match the 45s budget used by waitForGraphLoaded() in
  // server-connect.spec.ts; under parallel CI workers, downloading the full
  // graph can occasionally exceed 30s.
  await expect(page.locator('[data-testid="status-ready"]')).toBeVisible({ timeout: 45_000 });
}

async function openActiveRepoDropdown(page: import('@playwright/test').Page) {
  const projectButton = page.getByRole('button', { name: firstRepoName, exact: true });
  await expect(projectButton).toBeVisible({ timeout: 5_000 });
  await projectButton.click();
  await expect(page.getByText('Repositories')).toBeVisible({ timeout: 5_000 });
}

// ── Flow 1: Start screen and runtime connection state ──────────────────────

test.describe('Flow 1: Start screen and runtime connection', () => {
  test('shows the exe-served start screen first', async ({ page }) => {
    await page.goto('/');

    await expect(page).toHaveTitle('Anvien');
    await expect(page.getByRole('heading', { name: 'Anvien' })).toBeVisible();
    await expect(page.getByRole('button', { name: 'Start Anvien' })).toBeVisible();
    await expect(page.getByRole('button', { name: 'RESET RUNTIME' })).toBeVisible();
    await expect(page.getByRole('button', { name: 'User Guide' })).toBeVisible();
  });

  test('User Guide displays README.md content', async ({ page }) => {
    await page.goto('/');

    await page.getByRole('button', { name: 'User Guide' }).click();

    await expect(page.getByRole('heading', { name: 'User Guide' })).toBeVisible();
    await expect(page.getByText('# Anvien')).toBeVisible({ timeout: 10_000 });
    await expect(page.getByText('## Why Use Anvien?')).toBeVisible();
  });

  test('shows neutral runtime connection state when backend is unreachable', async ({
    page,
  }, testInfo) => {
    // Block all requests to the backend so the probe fails
    await page.route(`${BACKEND_URL}/**`, (route) => route.abort('connectionrefused'));

    await launchFromStartScreen(page);

    await expect(page.getByText('Connecting to Anvien runtime...')).toBeVisible({
      timeout: 10_000,
    });
    await expect(page.getByText('Start Anvien locally')).toHaveCount(0);
    await expect(page.getByText('anvien serve')).toHaveCount(0);
    await expect(page.getByText('Copy the command')).toHaveCount(0);
    await expect(page.getByText('Listening for local bridge')).toHaveCount(0);
    await page.screenshot({ path: testInfo.outputPath('runtime-connecting.png') });
  });

  test('packaged launcher Start reaches repo chooser or analyze without manual guide', async ({
    page,
  }, testInfo) => {
    test.skip(
      process.env.PACKAGED_LAUNCHER_E2E !== '1',
      'Requires AnvienLauncher.exe serving the packaged Web UI and backend runtime',
    );

    await launchFromStartScreen(page);

    await expect(page.getByText('Start Anvien locally')).toHaveCount(0);
    await expect(page.getByText('anvien serve')).toHaveCount(0);

    await expect
      .poll(
        async () => {
          const repoCardVisible = await page
            .getByTestId('landing-repo-card')
            .first()
            .isVisible()
            .catch(() => false);
          const analyzerVisible = await page
            .getByLabel('Repository Folder')
            .isVisible()
            .catch(() => false);
          return repoCardVisible || analyzerVisible;
        },
        { timeout: 45_000 },
      )
      .toBe(true);
    await page.screenshot({ path: testInfo.outputPath('packaged-start-target.png') });
  });
});

// ── Flow 2: Server detected → auto-connect ────────────────────────────────

test.describe('Flow 2: Server detected — auto-connect', () => {
  test('auto-connects when server becomes reachable', async ({ page }, testInfo) => {
    // Start with server unreachable
    let blockBackend = true;
    await page.route(`${BACKEND_URL}/**`, (route) => {
      if (blockBackend) return route.abort('connectionrefused');
      // Let it through to the real handler below
      return route.fallback();
    });

    // Mock the backend responses for when we "start" the server
    await page.route(`${BACKEND_URL}/api/repos`, async (route) => {
      if (blockBackend) return route.abort('connectionrefused');
      await route.fulfill({ json: [{ name: 'test-repo', path: '/tmp/test' }] });
    });
    await page.route(`${BACKEND_URL}/api/repo`, async (route) => {
      if (blockBackend) return route.abort('connectionrefused');
      await route.fulfill({
        json: { name: 'test-repo', path: '/tmp/test', repoPath: '/tmp/test' },
      });
    });
    await page.route(`${BACKEND_URL}/api/graph**`, async (route) => {
      if (blockBackend) return route.abort('connectionrefused');
      await route.fulfill({ json: { nodes: [], relationships: [] } });
    });
    await page.route(`${BACKEND_URL}/api/heartbeat`, async (route) => {
      if (blockBackend) return route.abort('connectionrefused');
      // SSE response
      await route.fulfill({
        status: 200,
        headers: { 'Content-Type': 'text/event-stream', 'Cache-Control': 'no-cache' },
        body: ':ok\n\n',
      });
    });

    await launchFromStartScreen(page);

    // Verify the neutral runtime connection state is shown first.
    await expect(page.getByText('Connecting to Anvien runtime...')).toBeVisible({
      timeout: 10_000,
    });
    await expect(page.getByText('Start Anvien locally')).toHaveCount(0);
    await expect(page.getByText('anvien serve')).toHaveCount(0);
    await page.screenshot({ path: testInfo.outputPath('before-server-start.png') });

    // "Start" the server by unblocking requests
    blockBackend = false;

    await page.evaluate(() => document.dispatchEvent(new Event('visibilitychange')));

    await expect(page.getByTestId('landing-repo-card').filter({ hasText: 'test-repo' })).toBeVisible({
      timeout: 15_000,
    });
    await page.screenshot({ path: testInfo.outputPath('repo-landing.png') });
  });

  test('transitions to analyze phase when server has zero repos', async ({ page }, testInfo) => {
    // Mock server with zero repos — repos endpoint returns empty array
    await page.route(`${BACKEND_URL}/api/repos`, (route) => route.fulfill({ json: [] }));
    await page.route(`${BACKEND_URL}/api/info`, (route) =>
      route.fulfill({ json: { version: '1.0.0', launchContext: 'npx', nodeVersion: 'v22.0.0' } }),
    );
    await page.route(`${BACKEND_URL}/api/heartbeat`, (route) =>
      route.fulfill({
        status: 200,
        headers: { 'Content-Type': 'text/event-stream' },
        body: ':ok\n\n',
      }),
    );

    await launchFromStartScreen(page);

    // Should transition to analyze when no indexed repos exist.
    await expect(page.getByLabel('Repository Folder')).toBeVisible({ timeout: 20_000 });
    await expect(page.getByRole('button', { name: 'Choose Repository' })).toBeVisible();
    await page.screenshot({ path: testInfo.outputPath('analyze-empty-state.png') });
  });
});

// ── Flow 3: Analyze form ───────────────────────────────────────────────────

test.describe('Flow 3: Analyze form', () => {
  test.beforeEach(async ({ page }) => {
    // Mock server with zero repos to show the analyze form
    await page.route(`${BACKEND_URL}/api/repos`, (route) => route.fulfill({ json: [] }));
    await page.route(`${BACKEND_URL}/api/info`, (route) =>
      route.fulfill({ json: { version: '1.0.0', launchContext: 'npx', nodeVersion: 'v22.0.0' } }),
    );
    await page.route(`${BACKEND_URL}/api/heartbeat`, (route) =>
      route.fulfill({
        status: 200,
        headers: { 'Content-Type': 'text/event-stream' },
        body: ':ok\n\n',
      }),
    );
  });

  test('local path input validates absolute paths', async ({ page }, testInfo) => {
    await launchFromStartScreen(page);

    // Wait for analyze form (transition: runtime connection -> success -> analyze)
    await expect(page.getByLabel('Repository Folder')).toBeVisible({ timeout: 20_000 });

    // Type an invalid relative path
    const input = page.getByLabel('Repository Folder');
    await input.fill('not-a-path');

    // Analyze button should be visible but disabled
    const analyzeBtn = page.getByRole('button', { name: /Analyze Repository/ });
    await expect(analyzeBtn).toBeVisible();
    await expect(analyzeBtn).toBeDisabled();

    // Type a valid absolute local path
    await input.fill(ABSOLUTE_LOCAL_PATH);
    await expect(analyzeBtn).toBeEnabled();
    await page.screenshot({ path: testInfo.outputPath('valid-local-path.png') });
  });

  test('local-only analyze form shows repository chooser', async ({ page }, testInfo) => {
    let pickerCalled = false;
    await page.route(`${BACKEND_URL}/api/local/folder-picker`, (route) => {
      pickerCalled = true;
      return route.fulfill({ json: { path: ABSOLUTE_LOCAL_PATH, cancelled: false } });
    });

    await launchFromStartScreen(page);

    await expect(page.getByLabel('Repository Folder')).toBeVisible({ timeout: 20_000 });
    await page.getByRole('button', { name: 'Choose Repository' }).click();
    await expect(page.getByLabel('Repository Folder')).toHaveValue(ABSOLUTE_LOCAL_PATH);
    expect(pickerCalled).toBe(true);
    await page.screenshot({ path: testInfo.outputPath('local-folder-input.png') });
  });

  test('pending repository chooser can be cancelled before analyzing a pasted path', async ({ page }) => {
    let releasePicker: (() => void) | undefined;
    let analyzeCalled = false;

    await page.route(`${BACKEND_URL}/api/local/folder-picker`, async (route) => {
      await new Promise<void>((resolve) => {
        releasePicker = resolve;
      });
      await route.fulfill({ json: { path: null, cancelled: true } });
    });
    await page.route(`${BACKEND_URL}/api/analyze`, async (route) => {
      expect(route.request().method()).toBe('POST');
      analyzeCalled = true;
      await route.fulfill({
        status: 202,
        json: { jobId: 'job-1', status: 'queued' },
      });
    });
    await page.route(`${BACKEND_URL}/api/analyze/job-1/progress`, async (route) => {
      await route.fulfill({
        status: 200,
        headers: { 'Content-Type': 'text/event-stream' },
        body: `event: complete\ndata: {"repoName":"demo","repoPath":"${ABSOLUTE_LOCAL_PATH.replace(/\\/g, '\\\\')}"}\n\n`,
      });
    });

    await launchFromStartScreen(page);
    await expect(page.getByLabel('Repository Folder')).toBeVisible({ timeout: 20_000 });

    await page.getByRole('button', { name: 'Choose Repository' }).click();
    await expect(page.getByRole('button', { name: 'Cancel Repository Picker' })).toBeVisible();

    await page.getByLabel('Repository Folder').fill(ABSOLUTE_LOCAL_PATH);
    const analyzeButton = page.getByRole('button', { name: /Analyze Repository/i });
    await expect(analyzeButton).toBeEnabled();
    await analyzeButton.click();
    releasePicker?.();

    await expect.poll(() => analyzeCalled).toBe(true);
    await expect(page.getByText('Request aborted')).toHaveCount(0);
  });

  test('repo landing remove button deletes repo and removes card', async ({ page }) => {
    const repos = [
      {
        name: 'demo-repo',
        path: ABSOLUTE_LOCAL_PATH,
        indexedAt: '2026-05-15T00:00:00Z',
        lastCommit: 'abc123',
        stats: { files: 3, nodes: 5 },
      },
    ];
    let deleteCalled = false;
    let deleted = false;

    await page.route(`${BACKEND_URL}/api/repos`, (route) => {
      return route.fulfill({ json: deleted ? [] : repos });
    });
    await page.route(`${BACKEND_URL}/api/info`, (route) =>
      route.fulfill({ json: { version: '1.0.0', launchContext: 'npx', nodeVersion: 'go-test' } }),
    );
    await page.route(`${BACKEND_URL}/api/heartbeat`, (route) =>
      route.fulfill({
        status: 200,
        headers: { 'Content-Type': 'text/event-stream' },
        body: ':ok\n\n',
      }),
    );
    await page.route(`${BACKEND_URL}/api/repo?repo=demo-repo`, (route) => {
      expect(route.request().method()).toBe('DELETE');
      deleteCalled = true;
      deleted = true;
      return route.fulfill({ json: { deleted: 'demo-repo' } });
    });

    await launchFromStartScreen(page);
    await expect(page.getByText('demo-repo')).toBeVisible({ timeout: 20_000 });
    await page.getByRole('button', { name: 'Remove demo-repo from repository list' }).click();

    await expect(page.getByText('demo-repo')).toHaveCount(0);
    await expect(page.getByLabel('Repository Folder')).toBeVisible();
    expect(deleteCalled).toBe(true);
  });

  test('invalid path keeps analyze disabled until corrected', async ({ page }) => {
    await launchFromStartScreen(page);

    await expect(page.getByLabel('Repository Folder')).toBeVisible({ timeout: 20_000 });

    const pathInput = page.getByLabel('Repository Folder');
    await pathInput.fill('relative-folder');
    const analyzeButton = page.getByRole('button', { name: /Analyze Repository/i });
    await expect(analyzeButton).toBeDisabled();

    await pathInput.fill(ABSOLUTE_LOCAL_PATH);
    await expect(analyzeButton).toBeEnabled();
  });
});

// ── Flow 4: Repo dropdown (requires running server) ────────────────────────

test.describe('Flow 4: Repo dropdown in exploring view', () => {
  const SKIP_MSG = 'Requires running anvien server with indexed repos';

  // enterExploringView() can take up to ~45s under parallel CI workers; combined
  // with the dropdown interactions this can exceed the default 60s test budget.
  test.slow();

  test.beforeAll(async () => {
    try {
      const res = await fetch(`${BACKEND_URL}/api/repos`);
      if (!res.ok) {
        test.skip(true, SKIP_MSG);
        return;
      }
      const repos = await res.json();
      if (!repos.length) {
        test.skip(true, 'Server has no indexed repos');
        return;
      }
      firstRepoName = repos[0].name;
    } catch {
      test.skip(true, SKIP_MSG);
    }
  });

  test('project badge opens repo dropdown', async ({ page }, testInfo) => {
    await enterExploringView(page);
    await page.screenshot({ path: testInfo.outputPath('exploring-loaded.png') });

    await openActiveRepoDropdown(page);

    // Repo dropdown should be visible
    await expect(page.getByText('Analyze a new repository')).toBeVisible();
    await page.screenshot({ path: testInfo.outputPath('repo-dropdown-open.png') });
  });

  test('analyze option opens inline form', async ({ page }, testInfo) => {
    await enterExploringView(page);

    await openActiveRepoDropdown(page);

    // Click "Analyze a new repository..."
    await page.getByText('Analyze a new repository').click();

    // Should show the local-only analyze form inline
    await expect(page.getByLabel('Repository Folder')).toBeVisible({ timeout: 5_000 });
    await expect(page.getByRole('button', { name: 'Choose Repository' })).toBeVisible();
    await page.screenshot({ path: testInfo.outputPath('inline-analyze-form.png') });
  });
});
