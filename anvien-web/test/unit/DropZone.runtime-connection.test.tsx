import { render, screen, waitFor } from '@testing-library/react';
import { beforeEach, describe, expect, it, vi } from 'vitest';
import { DropZone } from '../../src/components/DropZone';

const fetchReposMock = vi.fn();
const startPollingMock = vi.fn();
const stopPollingMock = vi.fn();

vi.mock('../../src/hooks/useBackend', () => ({
  useBackend: () => ({
    isConnected: false,
    isProbing: false,
    startPolling: startPollingMock,
    stopPolling: stopPollingMock,
    isPolling: false,
    backendUrl: 'http://127.0.0.1:4848',
  }),
}));

vi.mock('../../src/services/backend-client', () => ({
  fetchRepos: (...args: unknown[]) => fetchReposMock(...args),
  startAnalyze: vi.fn(),
  streamAnalyzeProgress: vi.fn(),
  connectToServer: vi.fn(),
  deleteRepo: vi.fn(),
}));

describe('DropZone runtime connection state', () => {
  beforeEach(() => {
    fetchReposMock.mockReset();
    startPollingMock.mockReset();
    stopPollingMock.mockReset();
  });

  it('does not render the retired manual bridge guide while waiting for runtime', async () => {
    render(<DropZone onServerConnect={vi.fn()} />);

    expect(screen.getByText('Connecting to Anvien runtime...')).toBeInTheDocument();
    expect(screen.queryByText('Start Anvien locally')).not.toBeInTheDocument();
    expect(screen.queryByText('anvien serve')).not.toBeInTheDocument();
    expect(screen.queryByText('Copy the command')).not.toBeInTheDocument();
    expect(screen.queryByText('Listening for local bridge')).not.toBeInTheDocument();

    await waitFor(() => expect(startPollingMock).toHaveBeenCalledTimes(1));
    expect(fetchReposMock).not.toHaveBeenCalled();
  });
});
