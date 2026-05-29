import { defineConfig, type Plugin } from "vite";
import react from "@vitejs/plugin-react";
import tailwindcss from "@tailwindcss/vite";
import fs from "fs";
import path from "path";
import { createRequire } from "module";

const _require = createRequire(import.meta.url);

function findSiblingPackageDir(packageName: string): string {
  const monorepoRoot = path.resolve(__dirname, "..");

  for (const entry of fs.readdirSync(monorepoRoot, { withFileTypes: true })) {
    if (!entry.isDirectory()) continue;

    const packageJsonPath = path.join(monorepoRoot, entry.name, "package.json");
    if (!fs.existsSync(packageJsonPath)) continue;

    try {
      const pkg = _require(packageJsonPath);
      if (pkg.name === packageName) {
        return path.join(monorepoRoot, entry.name);
      }
    } catch {
      // Ignore malformed or unrelated package manifests.
    }
  }

  throw new Error(
    `Could not find sibling package '${packageName}' from ${__dirname}`,
  );
}

const CLI_ROOT = findSiblingPackageDir("anvien");
const anvienCliPkg = _require(path.join(CLI_ROOT, "package.json"));
const REPO_ROOT = path.resolve(__dirname, "..");
const README_PATH = path.join(REPO_ROOT, "README.md");

function rootReadmePlugin(): Plugin {
  return {
    name: "anvien-root-readme",
    configureServer(server) {
      server.middlewares.use((req, res, next) => {
        const requestPath = (req.url ?? "").split("?")[0];
        if (requestPath !== "/README.md") {
          next();
          return;
        }
        if (!fs.existsSync(README_PATH)) {
          res.statusCode = 404;
          res.end("README.md not found");
          return;
        }
        res.setHeader("Content-Type", "text/markdown; charset=utf-8");
        fs.createReadStream(README_PATH).pipe(res);
      });
    },
    closeBundle() {
      const outDir = path.resolve(__dirname, "dist");
      if (!fs.existsSync(README_PATH) || !fs.existsSync(outDir)) {
        return;
      }
      fs.copyFileSync(README_PATH, path.join(outDir, "README.md"));
    },
  };
}

export default defineConfig({
  plugins: [rootReadmePlugin(), react(), tailwindcss()],
  define: {
    __REQUIRED_NODE_VERSION__: JSON.stringify(
      anvienCliPkg.engines.node.replace(/[>=^~\s]/g, ""),
    ),
  },
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
      // Fix for Rollup failing to resolve this deep import from @langchain/anthropic
      "@anthropic-ai/sdk/lib/transform-json-schema": path.resolve(
        __dirname,
        "node_modules/@anthropic-ai/sdk/lib/transform-json-schema.mjs",
      ),
      // Fix for mermaid d3-color prototype crash on Vercel (known issue with mermaid 10.9.0+ and Vite)
      mermaid: path.resolve(
        __dirname,
        "node_modules/mermaid/dist/mermaid.esm.min.mjs",
      ),
    },
  },
  server: {
    // Allow serving files from node_modules
    fs: {
      allow: [".."],
    },
  },
});
