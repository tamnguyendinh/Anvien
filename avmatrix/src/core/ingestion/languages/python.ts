/**
 * Python Language Provider
 *
 * Assembles all Python-specific ingestion capabilities into a single
 * LanguageProvider, following the Strategy pattern used by the pipeline.
 *
 * Key Python traits:
 *   - importSemantics: 'namespace' (Python uses namespace imports, not wildcard)
 *   - mroStrategy: 'c3' (Python C3 linearization for multiple inheritance)
 *   - namedBindingExtractor: present (from X import Y)
 */

import { SupportedLanguages } from 'avmatrix-shared';
import type { CaptureMatch, Scope, ScopeId } from 'avmatrix-shared';
import { createClassExtractor } from '../class-extractors/generic.js';
import { pythonClassConfig } from '../class-extractors/configs/python.js';
import { defineLanguage } from '../language-provider.js';
import { typeConfig as pythonConfig } from '../type-extractors/python.js';
import { pythonExportChecker } from '../export-detection.js';
import { createImportResolver } from '../import-resolvers/resolver-factory.js';
import { pythonImportConfig } from '../import-resolvers/configs/python.js';
import { extractPythonNamedBindings } from '../named-bindings/python.js';
import { PYTHON_QUERIES } from '../tree-sitter-queries.js';
import { createFieldExtractor } from '../field-extractors/generic.js';
import { pythonConfig as pythonFieldConfig } from '../field-extractors/configs/python.js';
import { createMethodExtractor } from '../method-extractors/generic.js';
import { pythonMethodConfig } from '../method-extractors/configs/python.js';
import { createVariableExtractor } from '../variable-extractors/generic.js';
import { pythonVariableConfig } from '../variable-extractors/configs/python.js';
import { createCallExtractor } from '../call-extractors/generic.js';
import { pythonCallConfig } from '../call-extractors/configs/python.js';
import { createHeritageExtractor } from '../heritage-extractors/generic.js';
import {
  emitPythonScopeCapturesFromTree,
  interpretPythonImport,
  interpretPythonTypeBinding,
} from '../scope-captures/python.js';

const BUILT_INS: ReadonlySet<string> = new Set([
  'print',
  'len',
  'range',
  'str',
  'int',
  'float',
  'list',
  'dict',
  'set',
  'tuple',
  'append',
  'extend',
  'update',
  'type',
  'isinstance',
  'issubclass',
  'getattr',
  'setattr',
  'hasattr',
  'enumerate',
  'zip',
  'sorted',
  'reversed',
  'min',
  'max',
  'sum',
  'abs',
]);

const PYTHON_DECLARATION_SUBTAGS: ReadonlySet<string> = new Set([
  '@declaration.name',
  '@declaration.owner',
  '@declaration.qualified_name',
  '@declaration.return_type',
  '@declaration.declared_type',
]);

const pythonBindingScopeFor = (match: CaptureMatch, innermostScope: Scope): ScopeId | null => {
  const declarationAnchor = Object.keys(match).find(
    (name) => name.startsWith('@declaration.') && !PYTHON_DECLARATION_SUBTAGS.has(name),
  );
  if (declarationAnchor === undefined) return null;

  const declarationKind = declarationAnchor.slice('@declaration.'.length);
  if (declarationKind === 'class' || declarationKind === 'function') {
    return innermostScope.parent ?? null;
  }

  if (declarationKind === 'method' || declarationKind === 'property') {
    return innermostScope.kind === 'Function' ? innermostScope.parent : null;
  }

  return null;
};

export const pythonProvider = defineLanguage({
  id: SupportedLanguages.Python,
  extensions: ['.py'],
  treeSitterQueries: PYTHON_QUERIES,
  typeConfig: pythonConfig,
  exportChecker: pythonExportChecker,
  importResolver: createImportResolver(pythonImportConfig),
  namedBindingExtractor: extractPythonNamedBindings,
  importSemantics: 'namespace',
  mroStrategy: 'c3',
  callExtractor: createCallExtractor(pythonCallConfig),
  fieldExtractor: createFieldExtractor(pythonFieldConfig),
  methodExtractor: createMethodExtractor(pythonMethodConfig),
  variableExtractor: createVariableExtractor(pythonVariableConfig),
  classExtractor: createClassExtractor(pythonClassConfig),
  heritageExtractor: createHeritageExtractor(SupportedLanguages.Python),
  emitScopeCapturesFromTree: emitPythonScopeCapturesFromTree,
  interpretImport: interpretPythonImport,
  interpretTypeBinding: interpretPythonTypeBinding,
  bindingScopeFor: pythonBindingScopeFor,
  builtInNames: BUILT_INS,
});
