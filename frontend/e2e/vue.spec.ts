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

// See here how to get started:
// https://playwright.dev/docs/intro
test('check title and number of pages for first book', async ({ page }) => {
  await page.goto('/');
  await expect(page.locator('tbody > tr:first-child > td').first()).toHaveText('some-title');
  await expect(page.locator('tbody > tr:first-child > td:nth-child(2)')).toHaveText('42');
})

test('check error state', async ({ page }) => {
  await page.addInitScript(`window.__E2E_ERROR__ = true`)
  await page.goto('/');
  await expect(page.locator('h3').first()).toHaveText('Error: Error');
})

test('check about page', async ({ page }) => {
  await page.goto('/about');
  await expect(page.locator('h1')).toHaveText('This is an about page');
})
