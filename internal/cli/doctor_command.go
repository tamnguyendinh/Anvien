package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tamnguyendinh/avmatrix-go/internal/repo"
)

type doctorLockTarget struct {
	RepoPath    string `json:"repoPath"`
	StoragePath string `json:"storagePath"`
	LockPath    string `json:"lockPath"`
}

type doctorLockReport struct {
	RepoPath    string             `json:"repoPath"`
	StoragePath string             `json:"storagePath"`
	LockPath    string             `json:"lockPath"`
	Status      string             `json:"status"`
	Diagnosis   repo.LockDiagnosis `json:"diagnosis"`
}

type doctorProcess struct {
	PID               int      `json:"pid"`
	ParentPID         int      `json:"parentPid,omitempty"`
	Name              string   `json:"name"`
	ExecutablePath    string   `json:"executablePath,omitempty"`
	CommandLine       string   `json:"commandLine,omitempty"`
	ParentName        string   `json:"parentName,omitempty"`
	ParentCommandLine string   `json:"parentCommandLine,omitempty"`
	Role              string   `json:"role"`
	Ownership         string   `json:"ownership"`
	Notes             []string `json:"notes,omitempty"`
}

func newDoctorCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "doctor",
		Short: "Inspect AVmatrix local runtime locks and processes",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}
	cmd.AddCommand(newDoctorLocksCommand(), newDoctorProcessesCommand())
	return cmd
}

func newDoctorLocksCommand() *cobra.Command {
	var repoQuery string
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "locks",
		Short: "Inspect repository analyze lock state",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			target, err := resolveDoctorLockTarget(repoQuery)
			if err != nil {
				return err
			}
			diagnosis, err := repo.DiagnoseStorageLock(target.LockPath)
			if err != nil {
				return err
			}
			report := doctorLockReport{
				RepoPath:    target.RepoPath,
				StoragePath: target.StoragePath,
				LockPath:    target.LockPath,
				Status:      doctorLockStatus(diagnosis),
				Diagnosis:   diagnosis,
			}
			if jsonOutput {
				raw, err := json.MarshalIndent(report, "", "  ")
				if err != nil {
					return err
				}
				_, err = fmt.Fprintln(cmd.OutOrStdout(), string(raw))
				return err
			}
			return writeDoctorLockReport(cmd, report)
		},
	}
	cmd.Flags().StringVar(&repoQuery, "repo", "", "repository name or path; defaults to current repository")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "write JSON lock report to stdout")
	return cmd
}

func newDoctorProcessesCommand() *cobra.Command {
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "processes",
		Short: "Inspect AVmatrix runtime processes without stopping them",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			processes, err := collectAVmatrixProcesses(contextFromCommand(cmd))
			if err != nil {
				return err
			}
			if jsonOutput {
				raw, err := json.MarshalIndent(processes, "", "  ")
				if err != nil {
					return err
				}
				_, err = fmt.Fprintln(cmd.OutOrStdout(), string(raw))
				return err
			}
			return writeDoctorProcesses(cmd, processes)
		},
	}
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "write JSON process report to stdout")
	return cmd
}

func resolveDoctorLockTarget(repoQuery string) (doctorLockTarget, error) {
	query := strings.TrimSpace(repoQuery)
	if query == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return doctorLockTarget{}, err
		}
		if gitRoot := repo.GitRoot(cwd); gitRoot != "" {
			cwd = gitRoot
		}
		return doctorLockTargetFromPath(cwd)
	}

	if looksLikeRepoPath(query) || pathExists(query) {
		return doctorLockTargetFromPath(query)
	}

	entries, err := repo.NewEnvStore().ReadRegistry()
	if err != nil {
		return doctorLockTarget{}, err
	}
	entry, err := repo.ResolveEntry(entries, query)
	if err != nil {
		return doctorLockTarget{}, err
	}
	storagePath := entry.StoragePath
	if storagePath == "" {
		storagePath = repo.StoragePath(entry.Path)
	}
	return doctorLockTarget{
		RepoPath:    entry.Path,
		StoragePath: storagePath,
		LockPath:    filepath.Join(storagePath, "analyze.lock"),
	}, nil
}

func doctorLockTargetFromPath(path string) (doctorLockTarget, error) {
	resolved, err := filepath.Abs(path)
	if err != nil {
		return doctorLockTarget{}, err
	}
	resolved = filepath.Clean(resolved)
	paths := repo.Paths(resolved)
	return doctorLockTarget{
		RepoPath:    resolved,
		StoragePath: paths.StoragePath,
		LockPath:    paths.AnalyzeLockPath,
	}, nil
}

func looksLikeRepoPath(query string) bool {
	return filepath.IsAbs(query) ||
		strings.HasPrefix(query, ".") ||
		strings.Contains(query, "/") ||
		strings.Contains(query, `\`)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func doctorLockStatus(diagnosis repo.LockDiagnosis) string {
	switch {
	case !diagnosis.Exists:
		return "free"
	case diagnosis.Stale && diagnosis.Recoverable:
		return "stale-recoverable"
	case diagnosis.ForeignHost:
		return "foreign-host"
	case diagnosis.Alive:
		return "held-live"
	default:
		return "held-unknown"
	}
}

func writeDoctorLockReport(cmd *cobra.Command, report doctorLockReport) error {
	out := cmd.OutOrStdout()
	if _, err := fmt.Fprintln(out, "AVmatrix analyze lock"); err != nil {
		return err
	}
	for _, line := range []string{
		"Repo: " + report.RepoPath,
		"Storage: " + report.StoragePath,
		"Lock: " + report.LockPath,
		"Status: " + report.Status,
	} {
		if _, err := fmt.Fprintln(out, line); err != nil {
			return err
		}
	}
	if !report.Diagnosis.Exists {
		return nil
	}
	info := report.Diagnosis.Info
	if info.PID > 0 {
		if _, err := fmt.Fprintf(out, "PID: %d\n", info.PID); err != nil {
			return err
		}
	}
	if info.Host != "" {
		if _, err := fmt.Fprintf(out, "Host: %s\n", info.Host); err != nil {
			return err
		}
	}
	if !info.AcquiredAt.IsZero() {
		if _, err := fmt.Fprintf(out, "Acquired: %s\n", info.AcquiredAt.Format(timeLayoutRFC3339Nano())); err != nil {
			return err
		}
	}
	if info.Command != "" {
		if _, err := fmt.Fprintf(out, "Command: %s\n", info.Command); err != nil {
			return err
		}
	}
	if report.Diagnosis.Reason != "" {
		if _, err := fmt.Fprintf(out, "Reason: %s\n", report.Diagnosis.Reason); err != nil {
			return err
		}
	}
	return nil
}

func timeLayoutRFC3339Nano() string {
	return "2006-01-02T15:04:05.999999999Z07:00"
}

func collectAVmatrixProcesses(ctx context.Context) ([]doctorProcess, error) {
	var processes []doctorProcess
	var err error
	if runtime.GOOS == "windows" {
		processes, err = collectWindowsAVmatrixProcesses(ctx)
	} else {
		processes, err = collectUnixAVmatrixProcesses(ctx)
	}
	if err != nil {
		return nil, err
	}
	currentPID := os.Getpid()
	filtered := processes[:0]
	for index := range processes {
		process := classifyDoctorProcess(processes[index])
		if process.PID == currentPID || process.ParentPID == currentPID || process.Role == "doctor" {
			continue
		}
		filtered = append(filtered, process)
	}
	return filtered, nil
}

func collectWindowsAVmatrixProcesses(ctx context.Context) ([]doctorProcess, error) {
	script := fmt.Sprintf(`
$currentPid = %d
$all = Get-CimInstance Win32_Process
$byPid = @{}
foreach ($p in $all) { $byPid[[int]$p.ProcessId] = $p }
$rows = $all | Where-Object {
  if ($_.ProcessId -eq $currentPid -or $_.ParentProcessId -eq $currentPid) { return $false }
  $name = if ($_.Name) { $_.Name } else { '' }
  $cmd = if ($_.CommandLine) { $_.CommandLine } else { '' }
  $name -match 'avmatrix' -or $cmd -match 'avmatrix'
} | ForEach-Object {
  $parent = $byPid[[int]$_.ParentProcessId]
  [PSCustomObject]@{
    pid = [int]$_.ProcessId
    parentPid = [int]$_.ParentProcessId
    name = if ($_.Name) { $_.Name } else { '' }
    executablePath = if ($_.ExecutablePath) { $_.ExecutablePath } else { '' }
    commandLine = if ($_.CommandLine) { $_.CommandLine } else { '' }
    parentName = if ($parent -and $parent.Name) { $parent.Name } else { '' }
    parentCommandLine = if ($parent -and $parent.CommandLine) { $parent.CommandLine } else { '' }
  }
}
@($rows) | ConvertTo-Json -Depth 4
`, os.Getpid())
	cmd := exec.CommandContext(ctx, "powershell", "-NoProfile", "-ExecutionPolicy", "Bypass", "-Command", script)
	raw, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	raw = []byte(strings.TrimSpace(string(raw)))
	if len(raw) == 0 {
		return []doctorProcess{}, nil
	}
	var processes []doctorProcess
	if err := json.Unmarshal(raw, &processes); err != nil {
		var single doctorProcess
		if singleErr := json.Unmarshal(raw, &single); singleErr != nil {
			return nil, err
		}
		processes = []doctorProcess{single}
	}
	return processes, nil
}

func collectUnixAVmatrixProcesses(ctx context.Context) ([]doctorProcess, error) {
	cmd := exec.CommandContext(ctx, "ps", "-eo", "pid=,ppid=,comm=,args=")
	raw, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	processes := []doctorProcess{}
	for _, line := range strings.Split(string(raw), "\n") {
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}
		pid, _ := strconv.Atoi(fields[0])
		ppid, _ := strconv.Atoi(fields[1])
		if pid == os.Getpid() || ppid == os.Getpid() {
			continue
		}
		name := fields[2]
		command := strings.Join(fields[3:], " ")
		if !strings.Contains(strings.ToLower(name+" "+command), "avmatrix") {
			continue
		}
		processes = append(processes, doctorProcess{
			PID:         pid,
			ParentPID:   ppid,
			Name:        name,
			CommandLine: command,
		})
	}
	return processes, nil
}

func classifyDoctorProcess(process doctorProcess) doctorProcess {
	haystack := strings.ToLower(process.Name + " " + process.CommandLine)
	parent := strings.ToLower(process.ParentName + " " + process.ParentCommandLine)

	switch {
	case strings.Contains(haystack, " doctor processes") || strings.Contains(haystack, " doctor locks"):
		process.Role = "doctor"
		process.Ownership = "diagnostic-command"
	case strings.Contains(haystack, " mcp") || strings.HasSuffix(strings.TrimSpace(haystack), "mcp"):
		process.Role = "mcp"
		process.Ownership = "editor-owned"
		process.Notes = append(process.Notes, "expected to stay running while the editor or agent session owns it")
	case strings.Contains(haystack, " analyze") || strings.Contains(haystack, ".exe analyze"):
		process.Role = "analyze"
		process.Ownership = "user-command-or-job"
		process.Notes = append(process.Notes, "should exit when analysis completes or is cancelled")
	case strings.Contains(haystack, " embed") || strings.Contains(haystack, " embedding"):
		process.Role = "embed"
		process.Ownership = "user-command-or-job"
		process.Notes = append(process.Notes, "should exit when embedding completes or is cancelled")
	case strings.Contains(haystack, " serve") || strings.Contains(haystack, ".exe serve"):
		process.Role = "serve"
		if strings.Contains(haystack, "--port 4848") || strings.Contains(haystack, "--port=4848") {
			process.Ownership = "launcher-owned"
		} else {
			process.Ownership = "user-command-or-dev-server"
		}
	case strings.Contains(haystack, "avmatrixlauncher.exe") || strings.Contains(haystack, "avmatrix-server.exe"):
		process.Role = "packaged-launcher"
		process.Ownership = "launcher-owned"
	default:
		process.Role = "avmatrix-runtime"
		process.Ownership = "unknown"
	}

	if process.Role == "mcp" && (strings.Contains(parent, "codex") || strings.Contains(parent, "claude") || strings.Contains(parent, "cursor")) {
		process.Notes = append(process.Notes, "parent process looks like an editor or AI agent integration")
	}
	return process
}

func writeDoctorProcesses(cmd *cobra.Command, processes []doctorProcess) error {
	out := cmd.OutOrStdout()
	if len(processes) == 0 {
		_, err := fmt.Fprintln(out, "No AVmatrix processes found.")
		return err
	}
	if _, err := fmt.Fprintln(out, "AVmatrix processes:"); err != nil {
		return err
	}
	for _, process := range processes {
		if _, err := fmt.Fprintf(out, "- pid=%d role=%s ownership=%s name=%s\n", process.PID, process.Role, process.Ownership, process.Name); err != nil {
			return err
		}
		if process.ParentPID > 0 || process.ParentName != "" {
			if _, err := fmt.Fprintf(out, "  parent: pid=%d name=%s\n", process.ParentPID, process.ParentName); err != nil {
				return err
			}
		}
		if process.CommandLine != "" {
			if _, err := fmt.Fprintf(out, "  command: %s\n", process.CommandLine); err != nil {
				return err
			}
		}
		for _, note := range process.Notes {
			if _, err := fmt.Fprintf(out, "  note: %s\n", note); err != nil {
				return err
			}
		}
	}
	return nil
}
