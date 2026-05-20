import { describe, expect, it } from 'vitest';
import type { GraphNode } from '../../src/generated/avmatrix-contracts';
import {
  DEFAULT_GRAPH_HEALTH_FILTERS,
  getGraphHealthDiagnosticCounts,
  getGraphHealthExpectedReasonCounts,
  getGraphHealthTopologyCounts,
  graphNodeMatchesHealthFilters,
} from '../../src/lib/graph-health-filters';

const makeHealthNode = (
  id: string,
  topologyStatus: GraphNode['properties']['topologyStatus'],
  extras: Partial<GraphNode['properties']['graphHealth']> = {},
): GraphNode =>
  ({
    id,
    label: 'Function',
    properties: {
      name: id,
      filePath: 'src/app.ts',
      graphHealth: {
        topologyStatus,
        countedIncoming: 0,
        countedOutgoing: 0,
        componentReachableFromRoot: false,
        confidence: 'candidate',
        ...extras,
      },
    },
  }) as GraphNode;

describe('graph health filters', () => {
  it('counts topology, expected reasons, and diagnostic occurrences', () => {
    const nodes = [
      makeHealthNode('a', 'no_incoming', {
        expectedIsolationReasons: ['test'],
        diagnostics: [{ kind: 'unresolved_reference', count: 3 }],
      }),
      makeHealthNode('b', 'detached_component', {
        expectedIsolationReasons: ['fixture'],
      }),
    ];

    expect(getGraphHealthTopologyCounts(nodes).get('no_incoming')).toBe(1);
    expect(getGraphHealthTopologyCounts(nodes).get('detached_component')).toBe(1);
    expect(getGraphHealthExpectedReasonCounts(nodes).get('test')).toBe(1);
    expect(getGraphHealthExpectedReasonCounts(nodes).get('fixture')).toBe(1);
    expect(getGraphHealthDiagnosticCounts(nodes).get('unresolved_reference')).toBe(3);
  });

  it('matches topology, expected-reason, and diagnostic visibility filters', () => {
    const node = makeHealthNode('a', 'unknown_connectivity', {
      expectedIsolationReasons: ['test'],
      diagnostics: [{ kind: 'unresolved_reference' }],
    });

    expect(graphNodeMatchesHealthFilters(node, DEFAULT_GRAPH_HEALTH_FILTERS)).toBe(true);
    expect(
      graphNodeMatchesHealthFilters(node, {
        ...DEFAULT_GRAPH_HEALTH_FILTERS,
        visibleTopologyStatuses: ['no_incoming'],
      }),
    ).toBe(false);
    expect(
      graphNodeMatchesHealthFilters(node, {
        ...DEFAULT_GRAPH_HEALTH_FILTERS,
        hiddenExpectedIsolationReasons: ['test'],
      }),
    ).toBe(false);
    expect(
      graphNodeMatchesHealthFilters(node, {
        ...DEFAULT_GRAPH_HEALTH_FILTERS,
        visibleDiagnosticKinds: [],
      }),
    ).toBe(false);
  });

  it('does not hide future diagnostic kinds without an explicit UI toggle', () => {
    const node = makeHealthNode('future-diagnostic', 'unknown_connectivity', {
      diagnostics: [{ kind: 'future_diagnostic' }],
    });

    expect(
      graphNodeMatchesHealthFilters(node, {
        ...DEFAULT_GRAPH_HEALTH_FILTERS,
        visibleDiagnosticKinds: [],
      }),
    ).toBe(true);
  });
});
