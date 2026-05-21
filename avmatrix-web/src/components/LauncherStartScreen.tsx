import { useState } from "react";
import { HelpCircle, Play, RefreshCw } from "@/lib/lucide-icons";

interface LauncherStartScreenProps {
  onStart: () => void;
  onResetRuntime?: () => void;
}

const GUIDE_UNAVAILABLE =
  "README.md content is not available in this build.";
const GUIDE_PATH = "/README.md";

export const LauncherStartScreen = ({
  onStart,
  onResetRuntime,
}: LauncherStartScreenProps) => {
  const [status, setStatus] = useState("");
  const [guideOpen, setGuideOpen] = useState(false);
  const [guideLoaded, setGuideLoaded] = useState(false);
  const [guideText, setGuideText] = useState("Loading README.md...");

  const loadGuide = async () => {
    if (guideLoaded) return;
    setGuideLoaded(true);

    try {
      const response = await fetch(GUIDE_PATH, { cache: "no-store" });
      if (!response.ok) {
        throw new Error(`Unable to load README.md (${response.status})`);
      }
      const markdown = await response.text();
      setGuideText(markdown.trim() || GUIDE_UNAVAILABLE);
    } catch {
      setGuideText(GUIDE_UNAVAILABLE);
    }
  };

  const handleStart = () => {
    setStatus("Starting AVmatrix...");
    onStart();
  };

  const handleReset = () => {
    setStatus("Resetting AVmatrix runtime...");
    if (onResetRuntime) {
      onResetRuntime();
      return;
    }
    window.location.href = "avmatrix://reset";
  };

  const handleGuideToggle = () => {
    const nextOpen = !guideOpen;
    setGuideOpen(nextOpen);
    if (nextOpen) {
      void loadGuide();
    }
  };

  return (
    <main className="press-shell press-ruled flex min-h-screen items-center justify-center p-8">
      <section className="grid w-full max-w-xl gap-4" aria-labelledby="launcher-start-title">
        <h1
          id="launcher-start-title"
          className="press-title text-center text-4xl leading-tight text-text-primary"
        >
          AVmatrix
        </h1>

        <button
          type="button"
          onClick={handleStart}
          className="press-filled-button flex min-h-14 w-full cursor-pointer items-center justify-center gap-2 px-5 py-3.5 text-base font-semibold text-text-inverse"
        >
          <Play className="h-4.5 w-4.5" />
          <span>Start AVmatrix</span>
        </button>

        <button
          type="button"
          onClick={handleReset}
          className="press-ghost-button flex min-h-12 w-full cursor-pointer items-center justify-center gap-2 px-5 py-3 text-sm font-semibold text-text-primary"
        >
          <RefreshCw className="h-4 w-4" />
          <span>RESET RUNTIME</span>
        </button>

        <button
          type="button"
          onClick={handleGuideToggle}
          aria-expanded={guideOpen}
          className="press-ghost-button flex min-h-12 w-full cursor-pointer items-center justify-center gap-2 px-5 py-3 text-sm font-semibold text-text-primary"
        >
          <HelpCircle className="h-4 w-4" />
          <span>User Guide</span>
        </button>

        {guideOpen && (
          <section className="press-panel max-h-[min(50vh,420px)] overflow-auto p-5">
            <h2 className="mb-3 text-left text-base font-semibold text-text-primary">
              User Guide
            </h2>
            <pre className="font-mono text-[13px] leading-relaxed whitespace-pre-wrap text-text-secondary">
              {guideText}
            </pre>
          </section>
        )}

        <p
          className="min-h-6 text-center text-sm text-text-secondary"
          aria-live="polite"
          role="status"
        >
          {status}
        </p>
      </section>
    </main>
  );
};
