import { act, fireEvent, render, screen, waitFor } from '@testing-library/react';
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { RepoAnalyzer } from '../../src/components/RepoAnalyzer';

const startAnalyzeMock = vi.fn();
const cancelAnalyzeMock = vi.fn();
const streamAnalyzeProgressMock = vi.fn();
const pickLocalFolderMock = vi.fn();

const { TestBackendError } = vi.hoisted(() => {
  class TestBackendError extends Error {
    constructor(
      message: string,
      public readonly status: number,
      public readonly code: 'network' | 'server' | 'client' | 'not_found',
    ) {
      super(message);
      this.name = 'BackendError';
    }
  }
  return { TestBackendError };
});

vi.mock('../../src/services/backend-client', () => ({
  BackendError: TestBackendError,
  pickLocalFolder: (...args: unknown[]) => pickLocalFolderMock(...args),
  startAnalyze: (...args: unknown[]) => startAnalyzeMock(...args),
  cancelAnalyze: (...args: unknown[]) => cancelAnalyzeMock(...args),
  streamAnalyzeProgress: (...args: unknown[]) => streamAnalyzeProgressMock(...args),
}));

vi.mock('../../src/components/AnalyzeProgress', () => ({
  AnalyzeProgress: () => <div>Analyzing</div>,
}));

describe('RepoAnalyzer local-only', () => {
  beforeEach(() => {
    startAnalyzeMock.mockReset();
    cancelAnalyzeMock.mockReset();
    streamAnalyzeProgressMock.mockReset();
    pickLocalFolderMock.mockReset();
    startAnalyzeMock.mockResolvedValue({ jobId: 'job-1', status: 'queued' });
    streamAnalyzeProgressMock.mockReturnValue(new AbortController());
  });

  afterEach(() => {
    vi.useRealTimers();
  });

  const validPath = () => (navigator.userAgent.toLowerCase().includes('win') ? 'C:\\repos\\avmatrix' : '/tmp/avmatrix');

  it('shows only local-folder input', () => {
    render(<RepoAnalyzer variant="sheet" onComplete={vi.fn()} />);

    expect(screen.getByText('Repository Folder')).toBeInTheDocument();
    expect(screen.queryByText('GitHub URL')).not.toBeInTheDocument();
    expect(screen.queryByText('GitHub Repository URL')).not.toBeInTheDocument();
  });

  it('submits an absolute local path to analyze', async () => {
    render(<RepoAnalyzer variant="sheet" onComplete={vi.fn()} />);

    const path = validPath();

    fireEvent.change(screen.getByLabelText('Repository Folder'), {
      target: { value: path },
    });
    fireEvent.click(screen.getByRole('button', { name: /Analyze Repository/i }));

    await waitFor(() => {
      expect(startAnalyzeMock).toHaveBeenCalledWith({ path });
    });
  });

  it('reports the analyzed repo path on completion', async () => {
    const onComplete = vi.fn();
    let completeAnalyze: ((data: { repoName?: string; repoPath?: string }) => void) | undefined;
    streamAnalyzeProgressMock.mockImplementation((_jobId, _onProgress, onCompleteCallback) => {
      completeAnalyze = onCompleteCallback;
      return new AbortController();
    });

    render(<RepoAnalyzer variant="sheet" onComplete={onComplete} />);

    const path = validPath();

    fireEvent.change(screen.getByLabelText('Repository Folder'), {
      target: { value: path },
    });
    fireEvent.click(screen.getByRole('button', { name: /Analyze Repository/i }));

    await waitFor(() => {
      expect(startAnalyzeMock).toHaveBeenCalledWith({ path });
    });

    vi.useFakeTimers();
    act(() => {
      completeAnalyze?.({ repoName: 'DisplayedName' });
      vi.advanceTimersByTime(1200);
    });

    expect(onComplete).toHaveBeenCalledWith({
      repoName: 'DisplayedName',
      repoPath: path,
    });
  });

  it('shows a cancel action while the repository picker is pending', async () => {
    let pickerSignal: AbortSignal | undefined;
    pickLocalFolderMock.mockImplementation((signal: AbortSignal) => {
      pickerSignal = signal;
      return new Promise(() => {});
    });

    render(<RepoAnalyzer variant="sheet" onComplete={vi.fn()} />);

    fireEvent.click(screen.getByRole('button', { name: 'Choose Repository' }));

    await waitFor(() => {
      expect(screen.getByRole('button', { name: 'Cancel Repository Picker' })).toBeEnabled();
    });
    expect(pickerSignal?.aborted).toBe(false);
  });

  it('cancels a pending picker without showing an error', async () => {
    let pickerSignal: AbortSignal | undefined;
    pickLocalFolderMock.mockImplementation((signal: AbortSignal) => {
      pickerSignal = signal;
      return new Promise(() => {});
    });

    render(<RepoAnalyzer variant="sheet" onComplete={vi.fn()} />);

    fireEvent.click(screen.getByRole('button', { name: 'Choose Repository' }));
    await screen.findByRole('button', { name: 'Cancel Repository Picker' });
    fireEvent.click(screen.getByRole('button', { name: 'Cancel Repository Picker' }));

    await waitFor(() => {
      expect(pickerSignal?.aborted).toBe(true);
      expect(screen.getByRole('button', { name: 'Choose Repository' })).toBeEnabled();
    });
    expect(screen.queryByText(/Request aborted/i)).not.toBeInTheDocument();
  });

  it('submits a pasted path after aborting a pending picker', async () => {
    let pickerSignal: AbortSignal | undefined;
    pickLocalFolderMock.mockImplementation((signal: AbortSignal) => {
      pickerSignal = signal;
      return new Promise(() => {});
    });

    render(<RepoAnalyzer variant="sheet" onComplete={vi.fn()} />);

    fireEvent.click(screen.getByRole('button', { name: 'Choose Repository' }));
    await screen.findByRole('button', { name: 'Cancel Repository Picker' });

    const path = validPath();
    fireEvent.change(screen.getByLabelText('Repository Folder'), {
      target: { value: path },
    });

    const analyzeButton = screen.getByRole('button', {
      name: /Analyze Repository/i,
    });
    expect(analyzeButton).toBeEnabled();
    fireEvent.click(analyzeButton);

    await waitFor(() => {
      expect(pickerSignal?.aborted).toBe(true);
      expect(startAnalyzeMock).toHaveBeenCalledWith({ path });
    });
  });

  it('aborts a pending picker request on unmount', async () => {
    let pickerSignal: AbortSignal | undefined;
    pickLocalFolderMock.mockImplementation((signal: AbortSignal) => {
      pickerSignal = signal;
      return new Promise(() => {});
    });

    const { unmount } = render(<RepoAnalyzer variant="sheet" onComplete={vi.fn()} />);

    fireEvent.click(screen.getByRole('button', { name: 'Choose Repository' }));
    await screen.findByRole('button', { name: 'Cancel Repository Picker' });
    unmount();

    expect(pickerSignal?.aborted).toBe(true);
  });

  it('treats wrapped picker abort errors as cancellation', async () => {
    pickLocalFolderMock.mockRejectedValue(new TestBackendError('Request aborted', 0, 'network'));

    render(<RepoAnalyzer variant="sheet" onComplete={vi.fn()} />);

    fireEvent.click(screen.getByRole('button', { name: 'Choose Repository' }));

    await waitFor(() => {
      expect(screen.getByRole('button', { name: 'Choose Repository' })).toBeEnabled();
    });
    expect(screen.queryByText(/Request aborted/i)).not.toBeInTheDocument();
  });
});
