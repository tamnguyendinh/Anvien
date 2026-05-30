import { render, screen, waitFor } from "@testing-library/react";
import userEvent from "@testing-library/user-event";
import { beforeEach, describe, expect, it, vi } from "vitest";
import type { FileHotspotsResponse } from "../../src/generated/anvien-contracts";

const fetchFileHotspots = vi.fn();

vi.mock("../../src/services/backend-client", () => ({
  fetchFileHotspots: (...args: unknown[]) => fetchFileHotspots(...args),
}));

import { FileMapPanel } from "../../src/components/FileMapPanel";

const response: FileHotspotsResponse = {
  repo: "demo",
  repoPath: "F:/demo",
  graph: {
    path: ".anvien/graph.json",
    stale: false,
  },
  total: 2,
  offset: 0,
  limit: 200,
  sort: "unresolved",
  files: [
    {
      path: "internal/mcp/server.go",
      language: "go",
      kind: "source",
      appLayer: "api",
      functionalArea: "mcp",
      parseStatus: "parsed",
      symbolCount: 20,
      exportedSymbolCount: 3,
      inboundRefCount: 14,
      outboundRefCount: 52,
      localRelationshipCount: 31,
      unresolvedSourceSiteCount: 2,
      linkedFlowCount: 4,
      linkedTestCount: 3,
      risk: "medium",
      stale: false,
      changedSinceAnalyze: true,
    },
    {
      path: "internal/mcp/server_test.go",
      language: "go",
      kind: "test",
      appLayer: "api_test",
      functionalArea: "mcp",
      parseStatus: "parsed",
      symbolCount: 11,
      exportedSymbolCount: 0,
      inboundRefCount: 1,
      outboundRefCount: 12,
      localRelationshipCount: 8,
      unresolvedSourceSiteCount: 0,
      linkedFlowCount: 0,
      linkedTestCount: 1,
      risk: "low",
      stale: false,
      changedSinceAnalyze: false,
    },
  ],
};

describe("FileMapPanel", () => {
  beforeEach(() => {
    fetchFileHotspots.mockReset();
    fetchFileHotspots.mockResolvedValue(response);
  });

  it("renders file map rows with summary counts", async () => {
    render(
      <FileMapPanel
        repoName="demo"
        selectedPath={null}
        onOpenFile={vi.fn()}
      />,
    );

    expect(await screen.findByText("internal/mcp/server.go")).toBeInTheDocument();
    expect(screen.getByText("2 files")).toBeInTheDocument();
    expect(screen.getByText("1 changed")).toBeInTheDocument();
    expect(screen.getByText("1 unresolved")).toBeInTheDocument();
    expect(screen.getByText("1 flows")).toBeInTheDocument();
    expect(screen.getByText("2 tests")).toBeInTheDocument();
    expect(screen.getAllByTestId("file-map-row")).toHaveLength(2);
  });

  it("passes sort and filter state to the backend request", async () => {
    render(
      <FileMapPanel
        repoName="demo"
        selectedPath={null}
        onOpenFile={vi.fn()}
      />,
    );

    await screen.findByText("internal/mcp/server.go");
    await userEvent.selectOptions(screen.getByLabelText("File map sort"), "fan-out");
    await userEvent.click(screen.getByRole("button", { name: "Test" }));
    await userEvent.click(screen.getByTestId("file-map-filter-changed"));
    await userEvent.click(screen.getByTestId("file-map-filter-unresolved"));
    await userEvent.click(screen.getByTestId("file-map-filter-api"));

    await waitFor(() => {
      expect(fetchFileHotspots).toHaveBeenLastCalledWith(
        expect.objectContaining({
          repo: "demo",
          sort: "fan-out",
          limit: 200,
          kinds: ["test"],
          changedOnly: true,
          unresolvedOnly: true,
          apiOnly: true,
        }),
      );
    });
  });

  it("opens the selected file row", async () => {
    const onOpenFile = vi.fn();
    render(
      <FileMapPanel
        repoName="demo"
        selectedPath={null}
        onOpenFile={onOpenFile}
      />,
    );

    await userEvent.click(await screen.findByText("internal/mcp/server.go"));

    expect(onOpenFile).toHaveBeenCalledWith("internal/mcp/server.go");
  });
});
