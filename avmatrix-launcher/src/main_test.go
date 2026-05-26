package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestParseAction(t *testing.T) {
	tests := []struct {
		args []string
		want string
	}{
		{nil, "start"},
		{[]string{"register"}, "register"},
		{[]string{"avmatrix://reset"}, "reset"},
		{[]string{"stop"}, "stop"},
		{[]string{"unknown"}, "start"},
	}

	for _, tt := range tests {
		if got := parseAction(tt.args); got != tt.want {
			t.Fatalf("parseAction(%v) = %q, want %q", tt.args, got, tt.want)
		}
	}
}

func TestStaticHandlerServesFilesAndFallsBackToIndex(t *testing.T) {
	webDist := t.TempDir()
	if err := os.WriteFile(filepath.Join(webDist, "index.html"), []byte("index"), 0o644); err != nil {
		t.Fatalf("write index: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(webDist, "assets"), 0o755); err != nil {
		t.Fatalf("mkdir assets: %v", err)
	}
	if err := os.WriteFile(filepath.Join(webDist, "assets", "app.js"), []byte("asset"), 0o644); err != nil {
		t.Fatalf("write asset: %v", err)
	}

	server := httptest.NewServer(staticHandler(webDist))
	defer server.Close()

	assertBody(t, server.URL+"/assets/app.js", "asset")
	assertBody(t, server.URL+"/repo/detail", "index")
}

func TestStaticHandlerDoesNotServeRootStartScreen(t *testing.T) {
	webDist := t.TempDir()
	if err := os.WriteFile(filepath.Join(webDist, "index.html"), []byte("index"), 0o644); err != nil {
		t.Fatalf("write index: %v", err)
	}

	server := httptest.NewServer(staticHandler(webDist))
	defer server.Close()

	assertBody(t, server.URL+"/Start-AVmatrix.html", "index")
	assertBody(t, server.URL+"/repo/detail", "index")
}

func TestStaticHandlerInjectsLauncherLifecycleAndRecordsHeartbeat(t *testing.T) {
	webDist := t.TempDir()
	if err := os.WriteFile(filepath.Join(webDist, "index.html"), []byte("<html><body>app</body></html>"), 0o644); err != nil {
		t.Fatalf("write index: %v", err)
	}
	lifecycle := newWebLifecycleMonitor(100 * time.Millisecond)
	server := httptest.NewServer(staticHandlerWithLifecycle(webDist, lifecycle))
	defer server.Close()

	response, err := http.Get(server.URL + "/")
	if err != nil {
		t.Fatalf("GET index: %v", err)
	}
	raw, err := io.ReadAll(response.Body)
	_ = response.Body.Close()
	if err != nil {
		t.Fatalf("read index: %v", err)
	}
	body := string(raw)
	if !strings.Contains(body, "data-avmatrix-launcher-lifecycle") || !strings.Contains(body, launcherHeartbeatPath) {
		t.Fatalf("index missing launcher lifecycle script:\n%s", body)
	}

	req, err := http.NewRequest(http.MethodPost, server.URL+launcherHeartbeatPath, nil)
	if err != nil {
		t.Fatalf("heartbeat request: %v", err)
	}
	heartbeat, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("POST heartbeat: %v", err)
	}
	_ = heartbeat.Body.Close()
	if heartbeat.StatusCode != http.StatusNoContent {
		t.Fatalf("heartbeat status = %d, want 204", heartbeat.StatusCode)
	}
	if lifecycle.expired(time.Now().Add(500 * time.Millisecond)) {
		t.Fatalf("fresh heartbeat should keep lifecycle alive")
	}
}

func TestLauncherWebReadyRequiresLifecycleHeartbeat(t *testing.T) {
	launcherServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == launcherHeartbeatPath && r.Method == http.MethodPost {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer launcherServer.Close()

	if !launcherWebReadyAt(launcherServer.URL + launcherHeartbeatPath) {
		t.Fatalf("launcher heartbeat endpoint should mark web runtime ready")
	}

	devServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("vite index"))
	}))
	defer devServer.Close()

	if launcherWebReadyAt(devServer.URL + launcherHeartbeatPath) {
		t.Fatalf("plain 200 response must not be accepted as launcher web runtime")
	}
}

func TestWebLifecycleMonitorDoesNotExpireAfterHeartbeatStops(t *testing.T) {
	lifecycle := newWebLifecycleMonitor(40 * time.Millisecond)
	lifecycle.start()
	defer lifecycle.stop()

	lifecycle.recordHeartbeat(time.Now().Add(-500 * time.Millisecond))
	select {
	case <-lifecycle.Done():
		t.Fatalf("stale heartbeat alone must not expire lifecycle")
	case <-time.After(180 * time.Millisecond):
	}
}

func TestWebLifecycleMonitorAllowsUnboundedHeavyGraphLoadHeartbeatGap(t *testing.T) {
	lifecycle := newWebLifecycleMonitor(launcherUICloseGrace)
	now := time.Date(2026, 5, 19, 15, 0, 0, 0, time.UTC)
	lifecycle.recordHeartbeat(now)

	gap := 24 * time.Hour
	if snapshot := lifecycle.snapshot(now.Add(gap)); snapshot.Expired {
		t.Fatalf("heartbeat gap %s should not expire lifecycle: %#v", gap, snapshot)
	} else if snapshot.HeartbeatAge != gap {
		t.Fatalf("snapshot heartbeat age = %s, want %s", snapshot.HeartbeatAge, gap)
	}
}

func TestWebLifecycleSnapshotReportsStaleHeartbeatWithoutTimeout(t *testing.T) {
	lifecycle := newWebLifecycleMonitor(launcherUICloseGrace)
	now := time.Date(2026, 5, 19, 15, 0, 0, 0, time.UTC)
	lifecycle.recordHeartbeat(now)

	gap := 3 * time.Hour
	snapshot := lifecycle.snapshot(now.Add(gap))
	if snapshot.Expired {
		t.Fatalf("stale heartbeat should remain non-expiring: %#v", snapshot)
	}
	if snapshot.HeartbeatAge != gap {
		t.Fatalf("snapshot heartbeat age = %s, want %s", snapshot.HeartbeatAge, gap)
	}
}

func TestWebLifecycleClosedSignalUsesGraceBeforeShutdown(t *testing.T) {
	lifecycle := newWebLifecycleMonitor(100 * time.Millisecond)
	now := time.Now()
	lifecycle.recordHeartbeat(now)
	lifecycle.recordClosed(now)

	if lifecycle.expired(now.Add(50 * time.Millisecond)) {
		t.Fatalf("close signal should keep a short reload grace window")
	}
	if !lifecycle.expired(now.Add(150 * time.Millisecond)) {
		t.Fatalf("close signal should expire after grace window")
	}
}

func TestWebLifecycleClosedSignalWithoutHeartbeatUsesGrace(t *testing.T) {
	lifecycle := newWebLifecycleMonitor(100 * time.Millisecond)
	now := time.Date(2026, 5, 19, 15, 0, 0, 0, time.UTC)
	lifecycle.recordClosed(now)

	if lifecycle.expired(now.Add(50 * time.Millisecond)) {
		t.Fatalf("close signal should keep grace even without prior heartbeat")
	}
	if !lifecycle.expired(now.Add(150 * time.Millisecond)) {
		t.Fatalf("close signal should expire after grace even without prior heartbeat")
	}
}

func TestWebLifecycleSnapshotReportsClosedGraceExpiry(t *testing.T) {
	lifecycle := newWebLifecycleMonitor(100 * time.Millisecond)
	now := time.Date(2026, 5, 19, 15, 0, 0, 0, time.UTC)
	lifecycle.recordHeartbeat(now)
	lifecycle.recordClosed(now)

	snapshot := lifecycle.snapshot(now.Add(150 * time.Millisecond))
	if !snapshot.Expired {
		t.Fatalf("snapshot should be expired after close grace: %#v", snapshot)
	}
	if snapshot.Reason != "closed_grace_expired" {
		t.Fatalf("snapshot reason = %q, want closed_grace_expired", snapshot.Reason)
	}
	if !snapshot.ClosedSeen || !snapshot.LastClosed.Equal(now) {
		t.Fatalf("snapshot did not record close signal: %#v", snapshot)
	}
}

func TestVerifyWebDist(t *testing.T) {
	webDist := t.TempDir()
	if err := verifyWebDist(webDist); err == nil {
		t.Fatalf("verifyWebDist without index returned nil")
	}
	if err := os.WriteFile(filepath.Join(webDist, "index.html"), []byte("index"), 0o644); err != nil {
		t.Fatalf("write index: %v", err)
	}
	if err := verifyWebDist(webDist); err != nil {
		t.Fatalf("verifyWebDist with index: %v", err)
	}
}

func TestOpenBrowserCanBeSuppressedForSmokeTests(t *testing.T) {
	t.Setenv("AVMATRIX_LAUNCHER_NO_BROWSER", "1")
	if err := openBrowser("http://127.0.0.1:5228"); err != nil {
		t.Fatalf("openBrowser with suppression: %v", err)
	}
}

func TestHiddenCommandAppliesHiddenProcAttr(t *testing.T) {
	cmd := hiddenCommand("taskkill", "/PID", "123", "/T")
	if cmd.SysProcAttr == nil {
		t.Fatalf("hiddenCommand SysProcAttr = nil")
	}
	if runtime.GOOS == "windows" {
		if !cmd.SysProcAttr.HideWindow {
			t.Fatalf("hiddenCommand HideWindow = false, want true")
		}
		if cmd.SysProcAttr.CreationFlags&0x08000000 == 0 {
			t.Fatalf("hiddenCommand CreationFlags = %#x, want CREATE_NO_WINDOW", cmd.SysProcAttr.CreationFlags)
		}
	}
}

func TestRuntimeProcessSweepIncludesRepoViteServer(t *testing.T) {
	paths := launcherPaths{
		exePath:    filepath.Join("C:", "repo", "avmatrix-launcher", "AVmatrixLauncher.exe"),
		rootDir:    filepath.Join("C:", "repo"),
		homeDir:    filepath.Join("C:", "repo", "avmatrix-launcher"),
		serverExe:  filepath.Join("C:", "repo", "avmatrix-launcher", "server-bundle", "avmatrix-server.exe"),
		backendExe: filepath.Join("C:", "repo", "avmatrix", "bin", "avmatrix.exe"),
	}

	script := buildStopRuntimeProcessesScript(paths, 1234)
	for _, want := range []string{"node.exe", "avmatrix-web", "vite", "--port 5228", "avmatrix\\bin\\avmatrix.exe", " serve", "--port 4848"} {
		if !strings.Contains(script, want) {
			t.Fatalf("runtime sweep script missing %q:\n%s", want, script)
		}
	}
	if strings.Contains(script, "$bundleDir") || strings.Contains(script, "server-bundle', 'avmatrix.exe") {
		t.Fatalf("runtime sweep should not target a server-bundle avmatrix.exe authority:\n%s", script)
	}
}

func TestConflictingWebRuntimeSweepTargetsOnlyRepoViteServer(t *testing.T) {
	paths := launcherPaths{
		rootDir: filepath.Join("C:", "repo"),
	}

	script := buildStopWebDevServerScript(paths, 1234)
	for _, want := range []string{"node.exe", "avmatrix-web", "vite", "--port 5228"} {
		if !strings.Contains(script, want) {
			t.Fatalf("web runtime sweep script missing %q:\n%s", want, script)
		}
	}
	if strings.Contains(script, "avmatrix-server.exe") || strings.Contains(script, "avmatrix.exe") {
		t.Fatalf("web runtime sweep should not target packaged backend processes:\n%s", script)
	}
}

func TestTasklistCommandRunsHiddenOnWindows(t *testing.T) {
	cmd := tasklistCommand(1234)
	if len(cmd.Args) < 2 || cmd.Args[0] != "tasklist" {
		t.Fatalf("tasklist command args = %#v", cmd.Args)
	}
	if cmd.SysProcAttr == nil {
		t.Fatalf("tasklist command SysProcAttr = nil")
	}
	if runtime.GOOS == "windows" {
		if !cmd.SysProcAttr.HideWindow {
			t.Fatalf("tasklist HideWindow = false, want true")
		}
		if cmd.SysProcAttr.CreationFlags&0x08000000 == 0 {
			t.Fatalf("tasklist CreationFlags = %#x, want CREATE_NO_WINDOW", cmd.SysProcAttr.CreationFlags)
		}
	}
}

func TestWriteAndReadState(t *testing.T) {
	paths := launcherPaths{
		rootDir:   t.TempDir(),
		stateFile: filepath.Join(t.TempDir(), "launcher-state.json"),
	}

	writeState(paths, "ready", 1234)
	state, err := readState(paths)
	if err != nil {
		t.Fatalf("readState: %v", err)
	}
	if state.RootDir != paths.rootDir || state.BackendPID != 1234 || state.Status != "ready" {
		t.Fatalf("state = %#v", state)
	}
}

func assertBody(t *testing.T, url string, want string) {
	t.Helper()
	response, err := http.Get(url)
	if err != nil {
		t.Fatalf("GET %s: %v", url, err)
	}
	defer response.Body.Close()
	raw, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("read response %s: %v", url, err)
	}
	if string(raw) != want {
		t.Fatalf("GET %s body = %q, want %q", url, string(raw), want)
	}
}
