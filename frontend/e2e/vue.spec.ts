import { test, expect } from '@playwright/test';


// take a screenshot for each test and add it to the report
test.afterEach(async ({ page }, testInfo) => {
  // Get a unique place for the screenshot.
  const screenshotPath = testInfo.outputPath(`result.png`);
  // Add it to the report.
  testInfo.attachments.push({ name: 'screenshot', path: screenshotPath, contentType: 'image/png' });
  // Take the screenshot itself.
  await page.screenshot({ path: screenshotPath, timeout: 5000 });
});

test.beforeEach(async ({ page }, testInfo) => {
  // log console logs to test output
  page.on('console', msg => console.log(msg.text()))
})

test('check home', async ({ page }) => {
  await page.goto('/');
  await expect(page.locator('tbody > tr:first-child > td').first()).toHaveText('some-title');
  await expect(page.locator('tbody > tr:first-child > td:nth-child(2)')).toHaveText('42');
  await expect(page).toHaveScreenshot()
})

test('check error state', async ({ page }) => {
  await page.addInitScript(`window.__E2E_ERROR__ = true`)
  await page.goto('/');
  await expect(page.locator('h3').first()).toContainText('Error:')
})

test('check about page', async ({ page }) => {
  await page.goto('/');
  await page.getByRole('link', { name: 'About' }).click();
  await expect(page.locator('h1')).toHaveText('This is an about page');
  await expect(page).toHaveScreenshot()
})
