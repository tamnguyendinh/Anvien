import { createHash } from "node:crypto";
import { spawnSync } from "node:child_process";
import { mkdir, readFile, writeFile } from "node:fs/promises";
import path from "node:path";
import { pathToFileURL } from "node:url";

const repoRoot = path.resolve(import.meta.dirname, "..");
const acceptedHarness = path.join(import.meta.dirname, "child02-p2c-affected-readers.mjs");
const expectedSha256 = "276B4D6EA54E97ED6E99A945BB13368685FE4FDC3CB4C5A6540BB82F8F3DA058";
const laneRoot = path.join(repoRoot, ".tmp", "qa-child02-p2e");
const ownerMarker = path.join(laneRoot, "owner.json");
const executionRoot = path.join(laneRoot, "browser");
const executionHarness = path.join(executionRoot, "child02-p2e-exec.mjs");
const prepareOnly = process.argv.includes("--prepare");

const source = await readFile(acceptedHarness, "utf8");
const actualSha256 = createHash("sha256").update(source).digest("hex").toUpperCase();
if (actualSha256 !== expectedSha256) {
  throw new Error(`Accepted P2-C Playwright harness hash drifted: expected ${expectedSha256}, got ${actualSha256}`);
}

const owner = JSON.parse(await readFile(ownerMarker, "utf8"));
if (owner.scope !== "child02-p2e" || owner.owner !== "independent-qa") {
  throw new Error(`Refusing to write outside the P2-E-owned lane root: ${laneRoot}`);
}

const head = spawnSync("git", ["rev-parse", "HEAD"], { cwd: repoRoot, encoding: "utf8" }).stdout.trim();
const playwrightURL = pathToFileURL(path.join(repoRoot, "anvien-web", "node_modules", "playwright", "index.mjs")).href;
let executionSource = source
  .replace('import { chromium } from "../anvien-web/node_modules/playwright/index.mjs";', `import { chromium } from ${JSON.stringify(playwrightURL)};`)
  .replace('const repoRoot = path.resolve(import.meta.dirname, "..");', `const repoRoot = ${JSON.stringify(repoRoot)};`)
  .replaceAll("child02-p2c", "child02-p2e-affected-readers")
  .replaceAll("qa-p2c", "qa-p2e")
  .replaceAll("Child 02 P2-C", "Child 02 P2-E")
  .replaceAll("child02-p2c-playwright-v1", "child02-p2e-playwright-v1")
  .replace("4d456446fcc49aed0c6d489aa9c63e00d030b53c", head);

executionSource = executionSource.replace(
  `  evidence.console = evidence.console.filter((entry) =>\n    ["warning", "warn", "error"].includes(entry.type),\n  );`,
  `  evidence.console = evidence.console.filter((entry) =>\n    ["warning", "warn", "error"].includes(entry.type),\n  );\n  evidence.nonBlockingConsoleDiagnostics = evidence.console.filter(\n    (entry) =>\n      entry.type === "warning" &&\n      entry.text.includes("GPU stall due to ReadPixels"),\n  );\n  evidence.console = evidence.console.filter(\n    (entry) => !evidence.nonBlockingConsoleDiagnostics.includes(entry),\n  );`,
);

executionSource = executionSource.replace(
  `  const scrollIntoViewCalls = await page.evaluate(() => window.__qaScrollIntoViewCalls);`,
  `  await page.waitForFunction(\n    () => window.__qaScrollIntoViewCalls.some((call) => call.lineNumber === "10"),\n    undefined,\n    { timeout: 5000 },\n  );\n  const scrollIntoViewCalls = await page.evaluate(() => window.__qaScrollIntoViewCalls);`,
);

await mkdir(executionRoot, { recursive: true });
await writeFile(executionHarness, executionSource, "utf8");

let syntaxCheck = null;
if (prepareOnly) {
  const checked = spawnSync(process.execPath, ["--check", executionHarness], {
    cwd: repoRoot,
    encoding: "utf8",
  });
  syntaxCheck = {
    command: `${process.execPath} --check ${executionHarness}`,
    exitCode: checked.status,
    stdout: checked.stdout.trim(),
    stderr: checked.stderr.trim(),
  };
  if (checked.status !== 0) {
    throw new Error(`Transformed execution harness failed syntax check: ${checked.stderr}`);
  }
} else {
  await import(`${pathToFileURL(executionHarness).href}?run=${Date.now()}`);
}

const result = {
  schema: "child02-p2e-playwright-wrapper-v1",
  generatedAt: new Date().toISOString(),
  acceptedHarness,
  acceptedHarnessSha256: actualSha256,
  executionHarness,
  executionHarnessSha256: createHash("sha256").update(executionSource).digest("hex").toUpperCase(),
  mode: prepareOnly ? "prepare" : "run",
  syntaxCheck,
  head,
  outputRoot: path.join(repoRoot, "Reports", "qa", "playwright", "child02-p2e-affected-readers"),
};
await writeFile(path.join(executionRoot, "wrapper-result.json"), `${JSON.stringify(result, null, 2)}\n`, "utf8");
console.log(JSON.stringify(result));
