import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { beforeEach, describe, expect, it, vi } from "vitest";
import { LauncherStartScreen } from "../../src/components/LauncherStartScreen";

describe("LauncherStartScreen", () => {
  beforeEach(() => {
    vi.restoreAllMocks();
  });

  it("shows the exe-served start actions", () => {
    render(<LauncherStartScreen onStart={() => {}} />);

    expect(screen.getByRole("heading", { name: "Anvien" })).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Start Anvien" })).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "RESET RUNTIME" })).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "User Guide" })).toBeInTheDocument();
  });

  it("starts the in-app flow without using the root html launcher", () => {
    const onStart = vi.fn();

    render(<LauncherStartScreen onStart={onStart} />);

    fireEvent.click(screen.getByRole("button", { name: "Start Anvien" }));

    expect(onStart).toHaveBeenCalledTimes(1);
    expect(screen.getByRole("status")).toHaveTextContent("Starting Anvien...");
  });

  it("runs reset through the provided launcher reset action", () => {
    const onResetRuntime = vi.fn();

    render(
      <LauncherStartScreen
        onStart={() => {}}
        onResetRuntime={onResetRuntime}
      />,
    );

    fireEvent.click(screen.getByRole("button", { name: "RESET RUNTIME" }));

    expect(onResetRuntime).toHaveBeenCalledTimes(1);
    expect(screen.getByRole("status")).toHaveTextContent(
      "Resetting Anvien runtime...",
    );
  });

  it("shows a graceful user guide fallback when the guide file is missing", async () => {
    vi.stubGlobal(
      "fetch",
      vi.fn().mockResolvedValue({ ok: false, status: 404 }),
    );

    render(<LauncherStartScreen onStart={() => {}} />);

    fireEvent.click(screen.getByRole("button", { name: "User Guide" }));

    expect(fetch).toHaveBeenCalledWith("/README.md", { cache: "no-store" });

    await waitFor(() => {
      expect(
        screen.getByText("README.md content is not available in this build."),
      ).toBeInTheDocument();
    });
  });

  it("loads README.md into the user guide panel", async () => {
    vi.stubGlobal(
      "fetch",
      vi.fn().mockResolvedValue({
        ok: true,
        text: () => Promise.resolve("# Anvien\n\nREADME guide content"),
      }),
    );

    render(<LauncherStartScreen onStart={() => {}} />);

    fireEvent.click(screen.getByRole("button", { name: "User Guide" }));

    expect(fetch).toHaveBeenCalledWith("/README.md", { cache: "no-store" });
    await waitFor(() => {
      expect(screen.getByText(/README guide content/)).toBeInTheDocument();
    });
  });
});
