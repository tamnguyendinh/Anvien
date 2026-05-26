package repo

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestAcquireStorageLockExcludesConcurrentWritersAndReleases(t *testing.T) {
	lockPath := filepath.Join(t.TempDir(), ".avmatrix", "analyze.lock")

	lock, err := AcquireStorageLock(lockPath)
	if err != nil {
		t.Fatalf("AcquireStorageLock() error = %v", err)
	}
	if _, err := os.Stat(lockPath); err != nil {
		t.Fatalf("lock file missing: %v", err)
	}

	_, err = AcquireStorageLock(lockPath)
	if !errors.Is(err, ErrLockHeld) {
		t.Fatalf("second AcquireStorageLock() error = %v, want ErrLockHeld", err)
	}
	if err := lock.Release(); err != nil {
		t.Fatalf("Release() error = %v", err)
	}
	if _, err := os.Stat(lockPath); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("lock file after release stat error = %v, want not exist", err)
	}
}

func TestAcquireStorageLockRecoversDeadPIDLock(t *testing.T) {
	dir := t.TempDir()
	lockPath := filepath.Join(dir, ".avmatrix", "analyze.lock")
	if err := os.MkdirAll(filepath.Dir(lockPath), 0o755); err != nil {
		t.Fatalf("mkdir lock dir: %v", err)
	}
	oldLock := "pid=424242\nacquiredAt=2026-05-26T07:50:03Z\n"
	if err := os.WriteFile(lockPath, []byte(oldLock), 0o644); err != nil {
		t.Fatalf("write old lock: %v", err)
	}

	options := testStorageLockOptions(func(pid int) bool {
		return false
	})
	lock, err := acquireStorageLock(lockPath, options)
	if err != nil {
		t.Fatalf("acquireStorageLock() error = %v", err)
	}
	defer lock.Release()

	raw, err := os.ReadFile(lockPath)
	if err != nil {
		t.Fatalf("read recovered lock: %v", err)
	}
	if !strings.Contains(string(raw), "version=2") || !strings.Contains(string(raw), "pid=9001") {
		t.Fatalf("recovered lock metadata unexpected:\n%s", raw)
	}
}

func TestAcquireStorageLockKeepsLivePIDLock(t *testing.T) {
	lockPath := filepath.Join(t.TempDir(), ".avmatrix", "analyze.lock")
	if err := os.MkdirAll(filepath.Dir(lockPath), 0o755); err != nil {
		t.Fatalf("mkdir lock dir: %v", err)
	}
	existing := "version=2\npid=424242\nacquiredAt=2026-05-26T07:50:03Z\nhost=test-host\ncommand=avmatrix analyze\ntoken=owner\n"
	if err := os.WriteFile(lockPath, []byte(existing), 0o644); err != nil {
		t.Fatalf("write lock: %v", err)
	}

	_, err := acquireStorageLock(lockPath, testStorageLockOptions(func(pid int) bool {
		return pid == 424242
	}))
	if !errors.Is(err, ErrLockHeld) {
		t.Fatalf("acquireStorageLock() error = %v, want ErrLockHeld", err)
	}
	var held *LockHeldError
	if !errors.As(err, &held) {
		t.Fatalf("acquireStorageLock() error type = %T, want *LockHeldError", err)
	}
	if held.Info.PID != 424242 || held.Info.Command != "avmatrix analyze" {
		t.Fatalf("LockHeldError info = %#v", held.Info)
	}
	raw, err := os.ReadFile(lockPath)
	if err != nil {
		t.Fatalf("read lock: %v", err)
	}
	if string(raw) != existing {
		t.Fatalf("live lock was modified:\n%s", raw)
	}
}

func TestAcquireStorageLockKeepsForeignHostLock(t *testing.T) {
	lockPath := filepath.Join(t.TempDir(), ".avmatrix", "analyze.lock")
	if err := os.MkdirAll(filepath.Dir(lockPath), 0o755); err != nil {
		t.Fatalf("mkdir lock dir: %v", err)
	}
	existing := "version=2\npid=424242\nacquiredAt=2026-05-26T07:50:03Z\nhost=other-host\ncommand=avmatrix analyze\ntoken=owner\n"
	if err := os.WriteFile(lockPath, []byte(existing), 0o644); err != nil {
		t.Fatalf("write lock: %v", err)
	}

	_, err := acquireStorageLock(lockPath, testStorageLockOptions(func(pid int) bool {
		return false
	}))
	if !errors.Is(err, ErrLockHeld) {
		t.Fatalf("acquireStorageLock() error = %v, want ErrLockHeld", err)
	}
	var held *LockHeldError
	if !errors.As(err, &held) {
		t.Fatalf("acquireStorageLock() error type = %T, want *LockHeldError", err)
	}
	if held.Info.Host != "other-host" || !strings.Contains(held.Reason, "another host") {
		t.Fatalf("foreign host error = %#v", held)
	}
	raw, err := os.ReadFile(lockPath)
	if err != nil {
		t.Fatalf("read lock: %v", err)
	}
	if string(raw) != existing {
		t.Fatalf("foreign host lock was modified:\n%s", raw)
	}
}

func TestAcquireStorageLockRecoversMalformedOldLock(t *testing.T) {
	lockPath := filepath.Join(t.TempDir(), ".avmatrix", "analyze.lock")
	if err := os.MkdirAll(filepath.Dir(lockPath), 0o755); err != nil {
		t.Fatalf("mkdir lock dir: %v", err)
	}
	if err := os.WriteFile(lockPath, []byte("not lock metadata\n"), 0o644); err != nil {
		t.Fatalf("write malformed lock: %v", err)
	}
	oldTime := time.Date(2026, 5, 26, 7, 0, 0, 0, time.UTC)
	if err := os.Chtimes(lockPath, oldTime, oldTime); err != nil {
		t.Fatalf("age malformed lock: %v", err)
	}

	lock, err := acquireStorageLock(lockPath, testStorageLockOptions(func(pid int) bool {
		return false
	}))
	if err != nil {
		t.Fatalf("acquireStorageLock() error = %v", err)
	}
	defer lock.Release()
	raw, err := os.ReadFile(lockPath)
	if err != nil {
		t.Fatalf("read recovered lock: %v", err)
	}
	if !strings.Contains(string(raw), "version=2") || !strings.Contains(string(raw), "token=test-token") {
		t.Fatalf("malformed lock was not replaced with owned metadata:\n%s", raw)
	}
}

func TestReleaseDoesNotRemoveReplacedLock(t *testing.T) {
	lockPath := filepath.Join(t.TempDir(), ".avmatrix", "analyze.lock")
	lock, err := acquireStorageLock(lockPath, testStorageLockOptions(func(pid int) bool {
		return false
	}))
	if err != nil {
		t.Fatalf("acquireStorageLock() error = %v", err)
	}
	replacement := "version=2\npid=123\nacquiredAt=2026-05-26T07:50:03Z\nhost=test-host\ncommand=avmatrix analyze\ntoken=replacement\n"
	if err := os.WriteFile(lockPath, []byte(replacement), 0o644); err != nil {
		t.Fatalf("replace lock: %v", err)
	}

	if err := lock.Release(); err != nil {
		t.Fatalf("Release() error = %v", err)
	}
	raw, err := os.ReadFile(lockPath)
	if err != nil {
		t.Fatalf("replacement lock missing after Release: %v", err)
	}
	if string(raw) != replacement {
		t.Fatalf("replacement lock changed:\n%s", raw)
	}
}

func TestDiagnoseStorageLockReportsRecoverableStaleLock(t *testing.T) {
	lockPath := filepath.Join(t.TempDir(), ".avmatrix", "analyze.lock")
	if err := os.MkdirAll(filepath.Dir(lockPath), 0o755); err != nil {
		t.Fatalf("mkdir lock dir: %v", err)
	}
	if err := os.WriteFile(lockPath, []byte("pid=424242\nacquiredAt=2026-05-26T07:50:03Z\n"), 0o644); err != nil {
		t.Fatalf("write lock: %v", err)
	}

	diagnosis, err := diagnoseStorageLockWithOptions(lockPath, testStorageLockOptions(func(pid int) bool {
		return false
	}))
	if err != nil {
		t.Fatalf("diagnoseStorageLockWithOptions() error = %v", err)
	}
	if !diagnosis.Exists || !diagnosis.Stale || !diagnosis.Recoverable || diagnosis.Alive {
		t.Fatalf("diagnosis = %#v", diagnosis)
	}
	if diagnosis.Reason != storageLockRecoveryStalePID {
		t.Fatalf("diagnosis reason = %q", diagnosis.Reason)
	}
}

func testStorageLockOptions(processAlive func(int) bool) storageLockOptions {
	now := time.Date(2026, 5, 26, 8, 0, 0, 0, time.UTC)
	return storageLockOptions{
		now: func() time.Time {
			return now
		},
		hostname: func() string {
			return "test-host"
		},
		pid: func() int {
			return 9001
		},
		commandLine: func() string {
			return "avmatrix analyze --force"
		},
		processAlive: processAlive,
		token: func() string {
			return "test-token"
		},
	}
}
