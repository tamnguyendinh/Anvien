import { describe, expect, it } from 'vitest';
import {
  DOCUMENTATION_NODE_LABEL,
  FILTERABLE_LABELS,
  NODE_COLORS,
} from '../../src/lib/constants';
import {
  NODE_LABELS,
  type NodeLabel,
} from '../../src/generated/avmatrix-contracts';
import * as lucideIcons from '../../src/lib/lucide-icons';

const LEGEND_LABELS: NodeLabel[] = [...NODE_LABELS];

const ICON_MAP: Record<string, string> = {
  Project: 'Zap',
  Package: 'Layers',
  Module: 'Layers',
  Folder: 'Folder',
  File: 'FileCode',
  Documentation: 'Table',
  Class: 'Box',
  Function: 'Braces',
  Method: 'Braces',
  Variable: 'Variable',
  Interface: 'Hash',
  Enum: 'List',
  Decorator: 'AtSign',
  Import: 'FileCode',
  Type: 'Type',
  CodeElement: 'Code',
  Community: 'Globe',
  Process: 'GitBranch',
  Struct: 'Box',
  Macro: 'AtSign',
  Typedef: 'Type',
  Union: 'List',
  Namespace: 'Layers',
  Trait: 'Hash',
  Impl: 'Braces',
  TypeAlias: 'Type',
  Const: 'Variable',
  Static: 'Variable',
  Property: 'Variable',
  Record: 'Box',
  Delegate: 'Braces',
  Annotation: 'AtSign',
  Constructor: 'Braces',
  Template: 'Type',
  Section: 'Table',
  Route: 'Server',
  Tool: 'Zap',
};

describe('filter panel icon mappings', () => {
  it('every filterable label has a mapped icon', () => {
    for (const label of FILTERABLE_LABELS) {
      expect(ICON_MAP).toHaveProperty(label);
    }
  });

  it('every mapped icon is exported from lucide-icons', () => {
    const exportedNames = new Set(Object.keys(lucideIcons));
    const requiredIcons = new Set(Object.values(ICON_MAP));
    for (const iconName of requiredIcons) {
      expect(
        exportedNames.has(iconName),
        `${iconName} should be exported from lucide-icons`,
      ).toBe(true);
    }
  });

  it('covers every generated node label and display-only filter', () => {
    expect(Object.keys(ICON_MAP).sort()).toEqual(
      [...NODE_LABELS, DOCUMENTATION_NODE_LABEL].sort(),
    );
  });
});

describe('color legend', () => {
  it('includes every generated node label', () => {
    expect(LEGEND_LABELS).toEqual([...NODE_LABELS]);
  });

  it('every legend label has a color defined', () => {
    for (const label of LEGEND_LABELS) {
      expect(NODE_COLORS).toHaveProperty(label);
      expect(NODE_COLORS[label]).toMatch(/^#[0-9a-f]{6}$/i);
    }
  });

  it('legend labels match the generated graph contract order', () => {
    expect(LEGEND_LABELS).toEqual([...NODE_LABELS]);
  });

  it('legend labels are a subset of filterable labels', () => {
    const filterableSet = new Set(FILTERABLE_LABELS);
    for (const label of LEGEND_LABELS) {
      expect(
        filterableSet.has(label),
        `${label} should be in FILTERABLE_LABELS`,
      ).toBe(true);
    }
  });

  it('has no duplicate entries', () => {
    const unique = new Set(LEGEND_LABELS);
    expect(unique.size).toBe(LEGEND_LABELS.length);
  });
});
