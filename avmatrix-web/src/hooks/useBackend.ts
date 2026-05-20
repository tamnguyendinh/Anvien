import { useState, useEffect, useCallback, useRef } from 'react';
import {
  normalizeServerUrl,
  probeBackend,
  setBackendUrl as setServiceUrl,
} from '../services/backend-client';
import { DEFAULT_BACKEND_URL } from '../config/ui-constants';

// ── localStorage keys ────────────────────────────────────────────────────────

const LS_URL_KEY = 'avmatrix-backend-url';

// ── Public interface ─────────────────────────────────────────────────────────

export interface UseBackendResult {
  /** Backend probe succeeded */
  isConnected: boolean;
  /** Currently checking connection */
  isProbing: boolean;
  /** Current backend URL */
  backendUrl: string;
  /** Start checking for server availability */
  startPolling: () => void;
  /** Stop checking */
  stopPolling: () => void;
  /** Whether background checking is active */
  isPolling: boolean;
}

// ── Hook implementation ──────────────────────────────────────────────────────

export function useBackend(): UseBackendResult {
  const [backendUrl] = useState<string>(() => {
    try {
      const saved = localStorage.getItem(LS_URL_KEY);
      if (!saved) {
        return DEFAULT_BACKEND_URL;
      }
      const normalized = normalizeServerUrl(saved);
      localStorage.setItem(LS_URL_KEY, normalized);
      return normalized;
    } catch {
      try {
        localStorage.removeItem(LS_URL_KEY);
      } catch {
        // ignore cleanup failure
      }
      return DEFAULT_BACKEND_URL;
    }
  });

  const [isConnected, setIsConnected] = useState(false);
  const [isProbing, setIsProbing] = useState(false);

  // Race-condition guard: monotonically increasing probe ID
  const probeIdRef = useRef(0);

  // ── Core probe logic ───────────────────────────────────────────────────────

  const probe = useCallback(async (): Promise<boolean> => {
    const id = ++probeIdRef.current;
    setIsProbing(true);

    try {
      const ok = await probeBackend();
      if (id !== probeIdRef.current) return false;
      setIsConnected(ok);
      return ok;
    } catch {
      if (id === probeIdRef.current) {
        setIsConnected(false);
      }
      return false;
    } finally {
      if (id === probeIdRef.current) {
        setIsProbing(false);
      }
    }
  }, []);

  // ── Server detection ─────────────────────────────────────────────────────

  const [isPolling, setIsPolling] = useState(false);

  const probeRef = useRef(probe);
  probeRef.current = probe;

  const stopPolling = useCallback(() => {
    setIsPolling(false);
  }, []);

  const startPolling = useCallback(() => {
    setIsPolling(true);
    void probeRef.current().then((ok) => {
      if (ok) setIsPolling(false);
    });
  }, []);

  // On tab return during background checking, probe immediately.
  useEffect(() => {
    if (!isPolling) return;
    const handleVisibility = () => {
      if (!document.hidden) {
        // Probe immediately when the user returns to the tab.
        void probeRef.current().then((ok) => {
          if (ok) setIsPolling(false);
        });
      }
    };
    document.addEventListener('visibilitychange', handleVisibility);
    return () => document.removeEventListener('visibilitychange', handleVisibility);
  }, [isPolling]);

  // ── Mount: sync service URL + auto-probe ─────────────────────────────────

  useEffect(() => {
    setServiceUrl(backendUrl);
    void probe();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return {
    isConnected,
    isProbing,
    backendUrl,
    startPolling,
    stopPolling,
    isPolling,
  };
}
