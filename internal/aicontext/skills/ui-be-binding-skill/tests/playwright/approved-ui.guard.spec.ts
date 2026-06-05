import { expect, test } from '@playwright/test';

const forbiddenTexts = [
  'MVP',
  'Demo',
  'Sample',
  'Placeholder',
  'Mock data',
  'Coming soon',
  'TODO',
  'No data available',
  'Failed to load',
  'Loading...',
  'Backend',
  'API',
  'For testing only',
  'This is a demo',
  'Please configure',
  'Example',
  'Test data',
  'Development only',
  'Work in progress',
  'Under construction',
  'Beta',
  'Alpha',
];

test.describe('approved UI guard', () => {
  test('home page has no forbidden visible text', async ({ page }) => {
    await page.goto('/');

    for (const text of forbiddenTexts) {
      await expect(page.getByText(text, { exact: false })).toHaveCount(0);
    }
  });

  test('home page visible text snapshot is stable', async ({ page }) => {
    await page.goto('/');

    const visibleText = await page.locator('body').innerText();
    expect(normalizeVisibleText(visibleText)).toMatchSnapshot('home-visible-text.txt');
  });

  test('home page visual output matches approved baseline', async ({ page }) => {
    await page.goto('/');

    await expect(page).toHaveScreenshot('home-approved.png', {
      maxDiffPixels: 100,
      fullPage: true,
    });
  });

  test('approved pricing slots are visible', async ({ page }) => {
    await page.goto('/');

    await expect(page.getByText('$3/month', { exact: false })).toBeVisible();
    await expect(page.getByText('30-day free trial', { exact: false })).toBeVisible();
  });
});

function normalizeVisibleText(input: string): string {
  return input
    .split('\n')
    .map((line) => line.trim())
    .filter(Boolean)
    .join('\n');
}
