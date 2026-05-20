import { describe, it, expect, vi, beforeEach, afterEach } from 'vitest';
import { connectHeartbeat } from '../../src/services/backend-client';
import {
  getWebRuntimeDiagnostics,
  resetWebRuntimeDiagnostics,
} from '../../src/lib/runtime-diagnostics';

// Mock EventSource to simulate SSE behavior
class MockEventSource {
  onopen: (() => void) | null = null;
  onerror: (() => void) | null = null;
  closed = false;

  close() {
    this.closed = true;
  }
}

let lastEventSource: MockEventSource | null = null;

beforeEach(() => {
  lastEventSource = null;
  vi.stubGlobal(
    'EventSource',
    vi.fn().mockImplementation(() => {
      lastEventSource = new MockEventSource();
      return lastEventSource;
    }),
  );
  resetWebRuntimeDiagnostics();
});

afterEach(() => {
  vi.unstubAllGlobals();
});

describe('connectHeartbeat', () => {
  it('calls onConnect when EventSource opens', () => {
    const onConnect = vi.fn();
    const onReconnecting = vi.fn();
    connectHeartbeat(onConnect, onReconnecting);

    lastEventSource!.onopen!();
    expect(onConnect).toHaveBeenCalledOnce();
    expect(onReconnecting).not.toHaveBeenCalled();
  });

  it('calls onReconnecting on first error and leaves reconnecting to EventSource', () => {
    const onConnect = vi.fn();
    const onReconnecting = vi.fn();
    connectHeartbeat(onConnect, onReconnecting);

    // Simulate connection drop
    lastEventSource!.onerror!();

    expect(onReconnecting).toHaveBeenCalledOnce();
    expect(lastEventSource!.closed).toBe(false);
    expect(EventSource).toHaveBeenCalledTimes(1);
  });

  it('fires onReconnecting only once per disconnect', () => {
    const onConnect = vi.fn();
    const onReconnecting = vi.fn();
    connectHeartbeat(onConnect, onReconnecting);

    // First error
    lastEventSource!.onerror!();
    expect(onReconnecting).toHaveBeenCalledOnce();

    lastEventSource!.onerror!();
    expect(onReconnecting).toHaveBeenCalledOnce(); // still 1

    lastEventSource!.onerror!();
    expect(onReconnecting).toHaveBeenCalledOnce(); // still 1
  });

  it('does not create timer-driven EventSource instances on repeated errors', () => {
    const onConnect = vi.fn();
    const onReconnecting = vi.fn();
    connectHeartbeat(onConnect, onReconnecting);

    for (let i = 0; i < 10; i++) {
      lastEventSource!.onerror!();
    }

    expect(EventSource).toHaveBeenCalledTimes(1);
    expect(onReconnecting).toHaveBeenCalledOnce();
  });

  it('resets reconnecting state when connection recovers', () => {
    const onConnect = vi.fn();
    const onReconnecting = vi.fn();
    connectHeartbeat(onConnect, onReconnecting);

    // Drop
    lastEventSource!.onerror!();
    expect(onReconnecting).toHaveBeenCalledOnce();

    lastEventSource!.onopen!();
    expect(onConnect).toHaveBeenCalledOnce();

    // Drop again — should fire onReconnecting again (reset after recovery)
    lastEventSource!.onerror!();
    expect(onReconnecting).toHaveBeenCalledTimes(2);
  });

  it('closes the EventSource when cleanup is called', () => {
    const onConnect = vi.fn();
    const onReconnecting = vi.fn();
    const cleanup = connectHeartbeat(onConnect, onReconnecting);

    lastEventSource!.onerror!();
    cleanup();

    expect(lastEventSource!.closed).toBe(true);
    expect(EventSource).toHaveBeenCalledTimes(1);
  });

  it('records heartbeat connect and reconnect diagnostics', () => {
    const onConnect = vi.fn();
    const onReconnecting = vi.fn();
    connectHeartbeat(onConnect, onReconnecting);

    lastEventSource!.onopen!();
    lastEventSource!.onerror!();

    const diagnostics = getWebRuntimeDiagnostics();
    expect(diagnostics?.heartbeat.connects).toBe(1);
    expect(diagnostics?.heartbeat.reconnects).toBe(1);
    expect(diagnostics?.heartbeat.lastRetryAttempt).toBe(1);
  });
});
