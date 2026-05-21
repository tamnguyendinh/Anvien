import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { beforeEach, describe, expect, it, vi } from "vitest";
import { LauncherStartScreen } from "../../src/components/LauncherStartScreen";

describe("LauncherStartScreen", () => {
  beforeEach(() => {
    vi.restoreAllMocks();
  });

  it("shows the exe-served start actions", () => {
    render(<LauncherStartScreen onStart={() => {}} />);

    expect(screen.getByRole("heading", { name: "AVmatrix" })).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "Start AVmatrix" })).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "RESET RUNTIME" })).toBeInTheDocument();
    expect(screen.getByRole("button", { name: "User Guide" })).toBeInTheDocument();
  });

  it("starts the in-app flow without using the root html launcher", () => {
    const onStart = vi.fn();

    render(<LauncherStartScreen onStart={onStart} />);

    fireEvent.click(screen.getByRole("button", { name: "Start AVmatrix" }));

    expect(onStart).toHaveBeenCalledTimes(1);
    expect(screen.getByRole("status")).toHaveTextContent("Starting AVmatrix...");
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
      "Resetting AVmatrix runtime...",
    );
  });

  it("shows a graceful user guide fallback when the guide file is missing", async () => {
    vi.stubGlobal(
      "fetch",
      vi.fn().mockResolvedValue({ ok: false, status: 404 }),
    );

    render(<LauncherStartScreen onStart={() => {}} />);

    fireEvent.click(screen.getByRole("button", { name: "User Guide" }));

    await waitFor(() => {
      expect(
        screen.getByText("User guide content is not available in this build."),
      ).toBeInTheDocument();
    });
  });
});
