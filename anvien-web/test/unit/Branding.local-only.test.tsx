import { fireEvent, render, screen } from "@testing-library/react";
import { beforeEach, describe, expect, it, vi } from "vitest";

const setHelpDialogBoxOpen = vi.fn();

vi.mock("../../src/components/RepoAnalyzer", () => ({
  RepoAnalyzer: () => <div data-testid="repo-analyzer">repo-analyzer</div>,
}));
vi.mock("../../src/components/EmbeddingStatus", () => ({
  EmbeddingStatus: () => (
    <div data-testid="embedding-status">embedding-status</div>
  ),
}));
vi.mock("../../src/hooks/useAppState.local-runtime", () => ({
  useAppState: () => ({
    projectName: "Website",
    graph: null,
    openChatPanel: vi.fn(),
    isRightPanelOpen: false,
    rightPanelTab: "chat",
    setSettingsPanelOpen: vi.fn(),
    setHelpDialogBoxOpen,
  }),
}));
vi.mock("../../src/services/backend-client", () => ({
  deleteRepo: vi.fn(),
  fetchRepos: vi.fn(),
  startAnalyze: vi.fn(),
  streamAnalyzeProgress: vi.fn(),
}));

import { AnalyzeOnboarding } from "../../src/components/AnalyzeOnboarding";
import { Header } from "../../src/components/Header";
import { RepoLanding } from "../../src/components/RepoLanding";

describe("Anvien branding on active local surfaces", () => {
  beforeEach(() => {
    setHelpDialogBoxOpen.mockClear();
  });

  it("uses Anvien copy in AnalyzeOnboarding", () => {
    render(<AnalyzeOnboarding onComplete={() => {}} />);

    expect(screen.getByText("Anvien")).toBeInTheDocument();
    expect(
      screen.getByText(/Anvien will parse the code/i),
    ).toBeInTheDocument();
  });

  it("uses Anvien copy in RepoLanding", () => {
    render(
      <RepoLanding
        repos={[]}
        onSelectRepo={() => {}}
        onAnalyzeComplete={() => {}}
      />,
    );

    expect(screen.getByText("Anvien")).toBeInTheDocument();
    expect(screen.getByText("Analyze Repository")).toBeInTheDocument();
    expect(
      screen.getByText(
        /Local only\. No repository data leaves this machine\./i,
      ),
    ).toBeInTheDocument();
  });

  it("uses Anvien branding in Header active path", () => {
    render(<Header availableRepos={[]} />);

    expect(screen.getByText("Anvien")).toBeInTheDocument();
    expect(screen.getByTitle("Help")).toBeInTheDocument();
    expect(screen.queryByText("anvien")).not.toBeInTheDocument();
  });

  it("routes the header help action to local help instead of upstream GitHub", () => {
    render(<Header availableRepos={[]} />);

    fireEvent.click(screen.getByTitle("Help"));

    expect(setHelpDialogBoxOpen).toHaveBeenCalledWith(true);
  });

  it("routes the header back action to the in-app launcher start screen", () => {
    const onNavigateToStart = vi.fn();

    render(
      <Header availableRepos={[]} onNavigateToStart={onNavigateToStart} />,
    );

    fireEvent.click(screen.getByLabelText("Back to Start screen"));

    expect(onNavigateToStart).toHaveBeenCalledTimes(1);
  });
});
