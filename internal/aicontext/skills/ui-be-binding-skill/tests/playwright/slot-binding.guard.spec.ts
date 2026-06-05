import { expect, test } from '@playwright/test';

// Adjust this file to match the actual backend mocking strategy of the project.
// Purpose: verify backend values enter approved slots only and do not create extra UI text.

test.describe('approved slot binding guard', () => {
  test('pricing data binds only to approved pricing slots', async ({ page }) => {
    await page.route('**/api/pricing/current', async (route) => {
      await route.fulfill({
        status: 200,
        contentType: 'application/json',
        body: JSON.stringify({
          monthlyPriceLabel: '$3/month',
          trialLabel: '30-day free trial',
          paymentUrl: '/pricing',
          internalPlanId: 'DO_NOT_RENDER',
          debugMode: 'DO_NOT_RENDER',
          adminNote: 'DO_NOT_RENDER',
        }),
      });
    });

    await page.goto('/');

    await expect(page.getByText('$3/month', { exact: false })).toBeVisible();
    await expect(page.getByText('30-day free trial', { exact: false })).toBeVisible();

    await expect(page.getByText('DO_NOT_RENDER', { exact: false })).toHaveCount(0);
    await expect(page.getByText('internalPlanId', { exact: false })).toHaveCount(0);
    await expect(page.getByText('debugMode', { exact: false })).toHaveCount(0);
    await expect(page.getByText('adminNote', { exact: false })).toHaveCount(0);
  });
});
