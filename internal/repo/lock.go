package repo

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var ErrLockHeld = errors.New("repository index lock is already held")

const (
	storageLockVersion          = "2"
	storageLockRetryLimit       = 3
	storageLockMalformedGrace   = 5 * time.Minute
	storageLockRecoveryStalePID = "owning process is not running"
)

type StorageLock struct {
	path  string
	file  *os.File
	token string
}

type LockInfo struct {
	Path       string            `json:"path"`
	Version    string            `json:"version,omitempty"`
	PID        int               `json:"pid,omitempty"`
	AcquiredAt time.Time         `json:"acquiredAt,omitempty"`
	Host       string            `json:"host,omitempty"`
	Command    string            `json:"command,omitempty"`
	Malformed  bool              `json:"malformed,omitempty"`
	ModTime    time.Time         `json:"modTime,omitempty"`
	Raw        map[string]string `json:"raw,omitempty"`
	Token      string            `json:"-"`
}

type LockDiagnosis struct {
	Info        LockInfo      `json:"info"`
	Exists      bool          `json:"exists"`
	Alive       bool          `json:"alive"`
	Stale       bool          `json:"stale"`
	Recoverable bool          `json:"recoverable"`
	ForeignHost bool          `json:"foreignHost"`
	Age         time.Duration `json:"age"`
	Reason      string        `json:"reason,omitempty"`
}

type LockHeldError struct {
	Info   LockInfo
	Reason string
}

func (e *LockHeldError) Error() string {
	if e == nil {
		return ErrLockHeld.Error()
	}
	details := []string{}
	if e.Info.PID > 0 {
		details = append(details, fmt.Sprintf("pid=%d", e.Info.PID))
	}
	if e.Info.Host != "" {
		details = append(details, "host="+e.Info.Host)
	}
	if !e.Info.AcquiredAt.IsZero() {
		details = append(details, "acquiredAt="+e.Info.AcquiredAt.Format(time.RFC3339Nano))
	}
	if e.Info.Command != "" {
		details = append(details, "command="+e.Info.Command)
	}
	if e.Info.Malformed {
		details = append(details, "malformed=true")
	}
	if e.Reason != "" {
		details = append(details, "reason="+e.Reason)
	}
	if len(details) == 0 {
		return ErrLockHeld.Error()
	}
	return fmt.Sprintf("%s (%s)", ErrLockHeld, strings.Join(details, ", "))
}

func (e *LockHeldError) Is(target error) bool {
	return target == ErrLockHeld
}

type storageLockOptions struct {
	now            func() time.Time
	hostname       func() string
	pid            func() int
	commandLine    func() string
	processAlive   func(int) bool
	token          func() string
	malformedGrace time.Duration
}

func AcquireStorageLock(lockPath string) (*StorageLock, error) {
	return acquireStorageLock(lockPath, storageLockOptions{})
}

func acquireStorageLock(lockPath string, options storageLockOptions) (*StorageLock, error) {
	options = withStorageLockDefaults(options)
	if err := os.MkdirAll(filepath.Dir(lockPath), 0o755); err != nil {
		return nil, err
	}

	for attempt := 0; attempt < storageLockRetryLimit; attempt++ {
		token := options.token()
		file, err := os.OpenFile(lockPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
		if err == nil {
			lock := &StorageLock{path: lockPath, file: file, token: token}
			info := LockInfo{
				Path:       lockPath,
				Version:    storageLockVersion,
				PID:        options.pid(),
				AcquiredAt: options.now().UTC(),
				Host:       options.hostname(),
				Command:    options.commandLine(),
				Token:      token,
			}
			if err := writeStorageLockInfo(file, info); err != nil {
				_ = file.Close()
				_ = os.Remove(lockPath)
				return nil, err
			}
			return lock, nil
		}
		if !errors.Is(err, os.ErrExist) {
			return nil, err
		}

		diagnosis, diagnosisErr := diagnoseStorageLockWithOptions(lockPath, options)
		if diagnosisErr != nil {
			return nil, diagnosisErr
		}
		if diagnosis.Stale && diagnosis.Recoverable {
			if err := os.Remove(lockPath); err != nil && !errors.Is(err, os.ErrNotExist) {
				return nil, &LockHeldError{Info: diagnosis.Info, Reason: fmt.Sprintf("stale lock could not be removed: %v", err)}
			}
			continue
		}
		return nil, &LockHeldError{Info: diagnosis.Info, Reason: diagnosis.Reason}
	}

	diagnosis, err := diagnoseStorageLockWithOptions(lockPath, options)
	if err != nil {
		return nil, err
	}
	return nil, &LockHeldError{Info: diagnosis.Info, Reason: diagnosis.Reason}
}

func ReadStorageLockInfo(lockPath string) (LockInfo, error) {
	raw, err := os.ReadFile(lockPath)
	if err != nil {
		return LockInfo{Path: lockPath}, err
	}
	info := parseStorageLockInfo(raw)
	info.Path = lockPath
	if stat, err := os.Stat(lockPath); err == nil {
		info.ModTime = stat.ModTime()
	}
	return info, nil
}

func DiagnoseStorageLock(lockPath string) (LockDiagnosis, error) {
	return diagnoseStorageLockWithOptions(lockPath, storageLockOptions{})
}

func diagnoseStorageLockWithOptions(lockPath string, options storageLockOptions) (LockDiagnosis, error) {
	options = withStorageLockDefaults(options)
	info, err := ReadStorageLockInfo(lockPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return LockDiagnosis{Info: LockInfo{Path: lockPath}, Exists: false, Reason: "lock file does not exist"}, nil
		}
		return LockDiagnosis{Info: LockInfo{Path: lockPath}}, err
	}

	currentHost := options.hostname()
	foreignHost := info.Host != "" && currentHost != "" && !strings.EqualFold(info.Host, currentHost)
	age := lockInfoAge(info, options.now())
	alive := false
	if info.PID > 0 && !foreignHost {
		alive = options.processAlive(info.PID)
	}

	diagnosis := LockDiagnosis{
		Info:        info,
		Exists:      true,
		Alive:       alive,
		ForeignHost: foreignHost,
		Age:         age,
	}

	switch {
	case foreignHost:
		diagnosis.Reason = "lock belongs to another host"
	case info.PID > 0 && alive:
		diagnosis.Reason = "owning process is still running"
	case info.PID > 0:
		diagnosis.Stale = true
		diagnosis.Recoverable = true
		diagnosis.Reason = storageLockRecoveryStalePID
	case info.Malformed && age >= options.malformedGrace:
		diagnosis.Stale = true
		diagnosis.Recoverable = true
		diagnosis.Reason = "lock metadata is malformed and older than stale grace"
	case info.Malformed:
		diagnosis.Reason = "lock metadata is malformed but recent"
	default:
		diagnosis.Reason = "lock owner cannot be determined"
	}
	return diagnosis, nil
}

func (lock *StorageLock) Release() error {
	if lock == nil {
		return nil
	}
	var closeErr error
	if lock.file != nil {
		closeErr = lock.file.Close()
		lock.file = nil
	}
	if lock.token != "" {
		info, err := ReadStorageLockInfo(lock.path)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				if closeErr != nil {
					return closeErr
				}
				return nil
			}
			if closeErr != nil {
				return closeErr
			}
			return err
		}
		if info.Token != lock.token {
			if closeErr != nil {
				return closeErr
			}
			return nil
		}
	}
	removeErr := os.Remove(lock.path)
	if errors.Is(removeErr, os.ErrNotExist) {
		removeErr = nil
	}
	if closeErr != nil {
		return closeErr
	}
	return removeErr
}

func withStorageLockDefaults(options storageLockOptions) storageLockOptions {
	if options.now == nil {
		options.now = time.Now
	}
	if options.hostname == nil {
		options.hostname = currentHostname
	}
	if options.pid == nil {
		options.pid = os.Getpid
	}
	if options.commandLine == nil {
		options.commandLine = currentCommandLine
	}
	if options.processAlive == nil {
		options.processAlive = storageLockProcessAlive
	}
	if options.token == nil {
		options.token = newStorageLockToken
	}
	if options.malformedGrace <= 0 {
		options.malformedGrace = storageLockMalformedGrace
	}
	return options
}

func writeStorageLockInfo(file *os.File, info LockInfo) error {
	lines := []string{
		"version=" + storageLockVersion,
		fmt.Sprintf("pid=%d", info.PID),
		"acquiredAt=" + info.AcquiredAt.UTC().Format(time.RFC3339Nano),
		"host=" + sanitizeLockValue(info.Host),
		"command=" + sanitizeLockValue(info.Command),
		"token=" + sanitizeLockValue(info.Token),
	}
	if _, err := fmt.Fprintln(file, strings.Join(lines, "\n")); err != nil {
		return err
	}
	return file.Sync()
}

func parseStorageLockInfo(raw []byte) LockInfo {
	info := LockInfo{Raw: map[string]string{}}
	scanner := bufio.NewScanner(strings.NewReader(string(raw)))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok || strings.TrimSpace(key) == "" {
			info.Malformed = true
			continue
		}
		info.Raw[strings.TrimSpace(key)] = strings.TrimSpace(value)
	}
	if err := scanner.Err(); err != nil {
		info.Malformed = true
	}
	if len(info.Raw) == 0 {
		info.Malformed = true
		return info
	}

	info.Version = info.Raw["version"]
	info.Host = info.Raw["host"]
	info.Command = info.Raw["command"]
	info.Token = info.Raw["token"]

	pidText := info.Raw["pid"]
	if pidText == "" {
		info.Malformed = true
	} else if pid, err := strconv.Atoi(pidText); err == nil {
		info.PID = pid
	} else {
		info.Malformed = true
	}

	acquiredAtText := info.Raw["acquiredAt"]
	if acquiredAtText == "" {
		info.Malformed = true
	} else if acquiredAt, err := time.Parse(time.RFC3339Nano, acquiredAtText); err == nil {
		info.AcquiredAt = acquiredAt
	} else {
		info.Malformed = true
	}

	if info.Version != "" && info.Version != storageLockVersion {
		info.Malformed = true
	}
	if info.Version == storageLockVersion && info.Token == "" {
		info.Malformed = true
	}
	return info
}

func lockInfoAge(info LockInfo, now time.Time) time.Duration {
	basis := info.AcquiredAt
	if basis.IsZero() {
		basis = info.ModTime
	}
	if basis.IsZero() {
		return 0
	}
	age := now.Sub(basis)
	if age < 0 {
		return 0
	}
	return age
}

func currentHostname() string {
	host, err := os.Hostname()
	if err != nil {
		return ""
	}
	return host
}

func currentCommandLine() string {
	if len(os.Args) == 0 {
		return ""
	}
	return strings.Join(os.Args, " ")
}

func newStorageLockToken() string {
	var bytes [16]byte
	if _, err := rand.Read(bytes[:]); err == nil {
		return hex.EncodeToString(bytes[:])
	}
	return fmt.Sprintf("%d-%d", os.Getpid(), time.Now().UnixNano())
}

func sanitizeLockValue(value string) string {
	value = strings.ReplaceAll(value, "\r", " ")
	value = strings.ReplaceAll(value, "\n", " ")
	return strings.TrimSpace(value)
}

func storageLockProcessAlive(pid int) bool {
	if pid <= 0 {
		return false
	}
	if runtime.GOOS == "windows" {
		out, err := exec.Command("tasklist", "/FI", fmt.Sprintf("PID eq %d", pid), "/FO", "CSV", "/NH").Output()
		return err == nil && strings.Contains(string(out), fmt.Sprintf("\"%d\"", pid))
	}
	proc, err := os.FindProcess(pid)
	return err == nil && proc.Signal(syscall.Signal(0)) == nil
}
