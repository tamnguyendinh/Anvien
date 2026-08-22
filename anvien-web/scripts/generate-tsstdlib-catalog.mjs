import crypto from "node:crypto";
import fs from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";

import ts from "../node_modules/typescript/lib/typescript.js";

const schemaVersion = "tsstdlib.catalog.v1";
const authorityKind = "typescript_standard_library";
const identityVersion = "tsstdlib.semantic.v1";
const expectedTypeScriptVersion = "5.9.3";
const expectedIntegrity =
  "sha512-jl1vZzPDinLr9eUt3J/t7V6FgNEw9QjvBPdysz9KfQDD41fQrC2Y4vKQdiaUpFT4bXlb1RHhLpp8wtm6M5TgSw==";
const generationCommand = "node anvien-web/scripts/generate-tsstdlib-catalog.mjs";

const scriptPath = fileURLToPath(import.meta.url);
const webRoot = path.resolve(path.dirname(scriptPath), "..");
const repoRoot = path.resolve(webRoot, "..");
const typeScriptRoot = path.join(webRoot, "node_modules", "typescript");
const libRoot = path.join(typeScriptRoot, "lib");
const outputPath = path.join(repoRoot, "internal", "tsstdlib", "catalog.v1.json");

verifyLockedCompiler();

const inputPaths = fs
  .readdirSync(libRoot)
  .filter((name) => /^lib(?:\..+)?\.d\.ts$/.test(name))
  .sort(compareText);
const inputIndex = new Map(inputPaths.map((name, index) => [name, index]));
const inputs = inputPaths.map((name) => {
  const raw = fs.readFileSync(path.join(libRoot, name));
  const text = raw.toString("utf8");
  const references = ts
    .preProcessFile(text, true, true)
    .libReferenceDirectives.map((reference) => resolveLibReference(reference.fileName))
    .filter((reference, index, values) => values.indexOf(reference) === index)
    .sort(compareText)
    .map((reference) => requiredInputIndex(reference));
  const input = {
    p: name,
    b: raw.byteLength,
    h: sha256(raw),
  };
  if (references.length > 0) {
    input.r = references;
  }
  return input;
});

const aliases = sortedObject(
  [...ts.libMap.entries()].map(([name, fileName]) => [
    name.toLowerCase(),
    requiredInputIndex(fileName),
  ]),
);
const targets = buildTargetProfiles();
const symbols = buildSymbols();

const catalog = {
  schema: schemaVersion,
  authority: authorityKind,
  identity: identityVersion,
  ts: expectedTypeScriptVersion,
  integrity: expectedIntegrity,
  generator: generationCommand,
  hash: "",
  inputs,
  aliases,
  targets,
  symbols,
};
catalog.hash = sha256(Buffer.from(JSON.stringify(catalog), "utf8"));
assignSemanticIdentities(catalog.symbols, catalog.hash);

fs.mkdirSync(path.dirname(outputPath), { recursive: true });
fs.writeFileSync(outputPath, `${JSON.stringify(catalog)}\n`, "utf8");

const output = fs.readFileSync(outputPath);
process.stdout.write(
  `${JSON.stringify({
    output: path.relative(repoRoot, outputPath).replaceAll(path.sep, "/"),
    typescript: ts.version,
    inputs: inputs.length,
    inputBytes: inputs.reduce((sum, input) => sum + input.b, 0),
    symbols: symbols.length,
    members: symbols.reduce(
      (sum, symbol) =>
        sum +
        (symbol.mv?.length ?? 0) +
        (symbol.mt?.length ?? 0) +
        (symbol.ms?.length ?? 0),
      0,
    ),
    semanticIDs: countSemanticIdentities(symbols),
    catalogHash: catalog.hash,
    artifactBytes: output.byteLength,
    artifactSha256: sha256(output),
  })}\n`,
);

function verifyLockedCompiler() {
  if (ts.version !== expectedTypeScriptVersion) {
    throw new Error(
      `TypeScript version mismatch: got ${ts.version}, want ${expectedTypeScriptVersion}`,
    );
  }
  const lock = JSON.parse(
    fs.readFileSync(path.join(webRoot, "package-lock.json"), "utf8"),
  );
  const locked = lock.packages?.["node_modules/typescript"];
  if (
    locked?.version !== expectedTypeScriptVersion ||
    locked?.integrity !== expectedIntegrity
  ) {
    throw new Error("package-lock TypeScript version or integrity mismatch");
  }
  const installed = JSON.parse(
    fs.readFileSync(path.join(typeScriptRoot, "package.json"), "utf8"),
  );
  if (installed.version !== expectedTypeScriptVersion) {
    throw new Error("installed TypeScript package does not match the locked version");
  }
}

function buildTargetProfiles() {
  const targetOption = ts.optionDeclarations.find(
    (option) => option.name === "target",
  );
  if (!(targetOption?.type instanceof Map)) {
    throw new Error("TypeScript target option metadata is unavailable");
  }
  const entries = [["default", requiredInputIndex(ts.getDefaultLibFileName({}))]];
  for (const [name, target] of targetOption.type.entries()) {
    entries.push([
      name.toLowerCase(),
      requiredInputIndex(ts.getDefaultLibFileName({ target })),
    ]);
  }
  return sortedObject(entries);
}

function buildSymbols() {
  const rootNames = inputPaths.map((name) => path.join(libRoot, name));
  const program = ts.createProgram({
    rootNames,
    options: {
      noLib: true,
      skipLibCheck: true,
      target: ts.ScriptTarget.ESNext,
    },
  });
  const checker = program.getTypeChecker();
  const sourceFiles = program
    .getSourceFiles()
    .filter((sourceFile) => inputIndex.has(path.basename(sourceFile.fileName)));
  if (sourceFiles.length !== inputPaths.length) {
    throw new Error(
      `compiler program loaded ${sourceFiles.length} declaration inputs, want ${inputPaths.length}`,
    );
  }
  const anchor =
    sourceFiles.find(
      (sourceFile) => path.basename(sourceFile.fileName) === "lib.es5.d.ts",
    ) ?? sourceFiles[0];
  const globals = checker.getSymbolsInScope(
    anchor,
    ts.SymbolFlags.Value | ts.SymbolFlags.Type | ts.SymbolFlags.Namespace,
  );
  const catalogSymbols = [];
  for (const symbol of globals) {
    const name = symbol.getName();
    if (!name || name.startsWith("__@")) {
      continue;
    }
    const lanes = declarationLanes(symbol.getDeclarations() ?? []);
    if (laneCount(lanes) === 0) {
      continue;
    }
    const item = { n: name };
    attachLanes(item, lanes);
    const baseTypes = collectBaseTypes(checker, symbol, lanes);
    const valueOwners = collectValueOwners(checker, symbol, lanes);
    if (baseTypes.length > 0) item.bt = baseTypes;
    if (valueOwners.length > 0) item.ov = valueOwners;
    const valueMembers = collectMembers(
      checker,
      symbol,
      "v",
      lanes,
      valueOwners,
    );
    const typeMembers = collectMembers(
      checker,
      symbol,
      "t",
      lanes,
      valueOwners,
    );
    const namespaceMembers = collectMembers(
      checker,
      symbol,
      "s",
      lanes,
      valueOwners,
    );
    if (valueMembers.length > 0) item.mv = valueMembers;
    if (typeMembers.length > 0) item.mt = typeMembers;
    if (namespaceMembers.length > 0) item.ms = namespaceMembers;
    catalogSymbols.push(item);
  }
  catalogSymbols.sort((left, right) => compareText(left.n, right.n));
  return catalogSymbols;

  function collectMembers(
    typeChecker,
    owner,
    ownerLane,
    ownerLanes,
    valueOwners,
  ) {
    const candidates = new Map();
    const add = (member) => {
      const memberName = member?.getName?.();
      if (!memberName || memberName.startsWith("__@")) return;
      const current = candidates.get(memberName) ?? [];
      if (!current.some((candidate) => candidate === member)) {
        current.push(member);
      }
      candidates.set(memberName, current);
    };

    try {
      if (ownerLane === "v" && ownerLanes.v.length > 0) {
        for (const member of owner.exports?.values?.() ?? []) {
          add(member);
        }
        const valueDeclaration =
          owner.valueDeclaration ?? owner.getDeclarations()?.[0];
        if (valueOwners.length === 0 && valueDeclaration) {
          for (const member of typeChecker.getPropertiesOfType(
            typeChecker.getTypeOfSymbolAtLocation(owner, valueDeclaration),
          )) {
            add(member);
          }
        }
      }
      if (ownerLane === "t" && ownerLanes.t.length > 0) {
        for (const member of owner.members?.values?.() ?? []) {
          add(member);
        }
      }
      if (ownerLane === "s" && ownerLanes.s.length > 0) {
        for (const member of owner.exports?.values?.() ?? []) {
          add(member);
        }
      }
    } catch (error) {
      throw new Error(
        `failed to collect ${ownerLane} members for ${owner.getName()}: ${error.message}`,
      );
    }

    const members = [];
    for (const [memberName, symbolsForName] of candidates.entries()) {
      const mergedLanes = emptyLanes();
      for (const member of symbolsForName) {
        mergeLanes(
          mergedLanes,
          declarationLanes(member.getDeclarations() ?? []),
        );
      }
      if (laneCount(mergedLanes) === 0) continue;
      const member = { n: memberName };
      attachLanes(member, mergedLanes);
      members.push(member);
    }
    members.sort((left, right) => compareText(left.n, right.n));
    return members;
  }

  function collectBaseTypes(typeChecker, owner, ownerLanes) {
    if (ownerLanes.t.length === 0) return [];
    const names = new Set();
    try {
      const declaredType = typeChecker.getDeclaredTypeOfSymbol(owner);
      if (
        declaredType &&
        typeof typeChecker.getBaseTypes === "function" &&
        (declaredType.objectFlags &
          (ts.ObjectFlags.Class | ts.ObjectFlags.Interface)) !==
          0
      ) {
        for (const baseType of typeChecker.getBaseTypes(declaredType) ?? []) {
          addCatalogTypeName(names, baseType.aliasSymbol ?? baseType.symbol);
        }
      }
    } catch (error) {
      throw new Error(
        `failed to collect base types for ${owner.getName()}: ${error.message}`,
      );
    }
    return [...names].sort(compareText);
  }

  function collectValueOwners(typeChecker, owner, ownerLanes) {
    if (ownerLanes.v.length === 0 || hasClassDeclaration(owner)) return [];
    const valueDeclaration = owner.valueDeclaration ?? owner.getDeclarations()?.[0];
    if (!valueDeclaration) return [];
    const names = new Set();
    const visit = (type) => {
      for (const constituent of type?.types ?? []) visit(constituent);
      addCatalogTypeName(names, type?.aliasSymbol ?? type?.symbol);
    };
    try {
      visit(typeChecker.getTypeOfSymbolAtLocation(owner, valueDeclaration));
    } catch (error) {
      throw new Error(
        `failed to collect value owner for ${owner.getName()}: ${error.message}`,
      );
    }
    return [...names].sort(compareText);
  }

  function addCatalogTypeName(names, symbol) {
    const name = symbol?.getName?.();
    if (!name || name.startsWith("__")) return;
    if (
      !(symbol.getDeclarations?.() ?? []).some((declaration) =>
        inputIndex.has(path.basename(declaration.getSourceFile().fileName)),
      )
    ) {
      return;
    }
    names.add(name);
  }

  function hasClassDeclaration(symbol) {
    return (symbol.getDeclarations?.() ?? []).some((declaration) =>
      ts.isClassDeclaration(declaration),
    );
  }
}

function declarationLanes(declarations) {
  const lanes = emptyLanes();
  for (const declaration of declarations) {
    const reference = declarationReference(declaration);
    if (!reference) continue;
    for (const lane of meaningsForDeclaration(declaration)) {
      lanes[lane].push(reference);
    }
  }
  for (const lane of ["v", "t", "s"]) {
    lanes[lane] = uniqueReferences(lanes[lane]);
  }
  return lanes;
}

function meaningsForDeclaration(declaration) {
  if (
    ts.isClassDeclaration(declaration) ||
    ts.isClassExpression(declaration) ||
    ts.isEnumDeclaration(declaration)
  ) {
    return ["v", "t"];
  }
  if (
    ts.isInterfaceDeclaration(declaration) ||
    ts.isTypeAliasDeclaration(declaration) ||
    ts.isTypeParameterDeclaration(declaration)
  ) {
    return ["t"];
  }
  if (ts.isModuleDeclaration(declaration)) {
    return ["v", "s"];
  }
  if (
    ts.isFunctionDeclaration(declaration) ||
    ts.isFunctionExpression(declaration) ||
    ts.isMethodDeclaration(declaration) ||
    ts.isMethodSignature(declaration) ||
    ts.isPropertyDeclaration(declaration) ||
    ts.isPropertySignature(declaration) ||
    ts.isGetAccessorDeclaration(declaration) ||
    ts.isSetAccessorDeclaration(declaration) ||
    ts.isVariableDeclaration(declaration) ||
    ts.isBindingElement(declaration) ||
    ts.isParameter(declaration) ||
    ts.isEnumMember(declaration)
  ) {
    return ["v"];
  }
  return [];
}

function declarationReference(declaration) {
  const sourceFile = declaration.getSourceFile();
  const libraryName = path.basename(sourceFile.fileName);
  const library = inputIndex.get(libraryName);
  if (library === undefined) return null;
  const startPosition = declaration.getStart(sourceFile, false);
  const endPosition = declaration.getEnd();
  const start = sourceFile.getLineAndCharacterOfPosition(startPosition);
  const end = sourceFile.getLineAndCharacterOfPosition(endPosition);
  return [library, start.line + 1, start.character + 1, end.line + 1, end.character + 1];
}

function emptyLanes() {
  return { v: [], t: [], s: [] };
}

function mergeLanes(target, source) {
  for (const lane of ["v", "t", "s"]) {
    target[lane].push(...source[lane]);
    target[lane] = uniqueReferences(target[lane]);
  }
}

function attachLanes(target, lanes) {
  if (lanes.v.length > 0) target.v = lanes.v;
  if (lanes.t.length > 0) target.t = lanes.t;
  if (lanes.s.length > 0) target.s = lanes.s;
}

function assignSemanticIdentities(catalogSymbols, catalogHash) {
  for (const symbol of catalogSymbols) {
    attachSemanticIDs(
      symbol,
      catalogHash,
      [["global"]],
      symbol.n,
    );
    for (const [membersKey, ownerMeaning] of [
      ["mv", "value"],
      ["mt", "type"],
      ["ms", "namespace"],
    ]) {
      for (const member of symbol[membersKey] ?? []) {
        attachSemanticIDs(
          member,
          catalogHash,
          [["global", symbol.n, ownerMeaning]],
          member.n,
        );
      }
    }
  }
}

function attachSemanticIDs(target, catalogHash, ownerPath, name) {
  for (const [laneKey, idKey, meaning] of [
    ["v", "iv", "value"],
    ["t", "it", "type"],
    ["s", "is", "namespace"],
  ]) {
    const declarations = target[laneKey];
    if (!declarations?.length) continue;
    target[idKey] = semanticID(
      catalogHash,
      ownerPath,
      name,
      meaning,
      declarations,
    );
  }
}

function semanticID(catalogHash, ownerPath, name, meaning, declarations) {
  const payload = {
    a: authorityKind,
    v: expectedTypeScriptVersion,
    c: catalogHash,
    o: ownerPath,
    n: name,
    m: meaning,
    d: declarations.map(([library, ...range]) => [
      inputs[library].p,
      ...range,
    ]),
  };
  return `tsstdlib:${sha256(Buffer.from(JSON.stringify(payload), "utf8"))}`;
}

function countSemanticIdentities(catalogSymbols) {
  let count = 0;
  const countItem = (item) => {
    for (const key of ["iv", "it", "is"]) {
      if (item[key]) count += 1;
    }
  };
  for (const symbol of catalogSymbols) {
    countItem(symbol);
    for (const membersKey of ["mv", "mt", "ms"]) {
      for (const member of symbol[membersKey] ?? []) countItem(member);
    }
  }
  return count;
}

function laneCount(lanes) {
  return lanes.v.length + lanes.t.length + lanes.s.length;
}

function uniqueReferences(references) {
  references.sort(compareReference);
  return references.filter(
    (reference, index) =>
      index === 0 || compareReference(references[index - 1], reference) !== 0,
  );
}

function compareReference(left, right) {
  for (let index = 0; index < left.length; index += 1) {
    if (left[index] !== right[index]) return left[index] - right[index];
  }
  return 0;
}

function resolveLibReference(referenceName) {
  const normalized = referenceName.toLowerCase();
  return ts.libMap.get(normalized) ?? `lib.${normalized}.d.ts`;
}

function requiredInputIndex(name) {
  const index = inputIndex.get(name);
  if (index === undefined) {
    throw new Error(`TypeScript declaration input ${name} is not in the catalog corpus`);
  }
  return index;
}

function sortedObject(entries) {
  return Object.fromEntries(
    [...entries].sort(([left], [right]) => compareText(left, right)),
  );
}

function compareText(left, right) {
  if (left < right) return -1;
  if (left > right) return 1;
  return 0;
}

function sha256(raw) {
  return crypto.createHash("sha256").update(raw).digest("hex");
}
