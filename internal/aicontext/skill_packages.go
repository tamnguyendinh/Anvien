package aicontext

import (
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

const (
	skillManifestFileName   = ".anvien-skill-manifest.json"
	skillManifestVersion    = 1
	skillManifestOwner      = "anvien"
	embeddedSkillSourceRoot = "skills"
	runtimeSkillSourceRoot  = "internal/aicontext/skills"
)

//go:embed all:skills
var skillSourceFS embed.FS

type skillPackageSource struct {
	filesystem fs.FS
	readRoot   string
	sourceRoot string
}

type SkillPackage struct {
	Name        string
	Description string
	SourceRoot  string
	InstallRoot string
	Hash        string
	Entries     []SkillEntry
	Files       []SkillPackageFile
}

type SkillEntry struct {
	Name        string
	Description string
	SourcePath  string
	PackagePath string
	InstallPath string
}

type SkillPackageFile struct {
	SourcePath  string
	PackagePath string
	InstallPath string
	Hash        string
	SizeBytes   int64
	Content     []byte `json:"-"`
}

type BaseSkillFile struct {
	Name        string
	PackageName string
	Path        string
	InstallPath string
	Description string
	Task        string
	Content     string
}

type SkillInstallResult struct {
	PackageIDs   []string
	Discovered   int
	Installed    int
	Updated      int
	Skipped      int
	Adopted      int
	Stale        int
	Preserved    int
	Collisions   int
	Written      int
	Overwritten  int
	Deleted      int
	SkippedFiles int
	Unsafe       int
}

type skillFileSnapshot struct {
	Path        string
	Hash        string
	Content     []byte
	PackageName string
	PackagePath string
}

type skillSyncPlan struct {
	Writes     []skillFileSnapshot
	Overwrites []skillFileSnapshot
	Deletes    []skillFileSnapshot
	Skips      []skillFileSnapshot
}

type skillManifest struct {
	SchemaVersion int                           `json:"schemaVersion"`
	ManagedBy     string                        `json:"managedBy"`
	Skills        map[string]skillManifestEntry `json:"skills"`
}

type skillManifestEntry struct {
	InstallPath string            `json:"installPath"`
	SourceRoot  string            `json:"sourceRoot"`
	Hash        string            `json:"hash"`
	Managed     bool              `json:"managed"`
	Stale       bool              `json:"stale,omitempty"`
	EntryCount  int               `json:"entryCount"`
	FileCount   int               `json:"fileCount"`
	Files       map[string]string `json:"files,omitempty"`
}

func SkillPackages() ([]SkillPackage, error) {
	return discoverSkillPackages(skillSourceFS)
}

func SkillPackagesForRepo(repoPath string) ([]SkillPackage, error) {
	source, ok, err := runtimeSkillPackageSource(repoPath)
	if err != nil {
		return nil, err
	}
	if ok {
		return discoverSkillPackagesFrom(source)
	}
	return SkillPackages()
}

func BaseSkillFiles() ([]BaseSkillFile, error) {
	packages, err := SkillPackages()
	if err != nil {
		return nil, err
	}
	return baseSkillFilesFromPackages(packages), nil
}

func BaseSkillFilesForRepo(repoPath string) ([]BaseSkillFile, error) {
	packages, err := SkillPackagesForRepo(repoPath)
	if err != nil {
		return nil, err
	}
	return baseSkillFilesFromPackages(packages), nil
}

func baseSkillFilesFromPackages(packages []SkillPackage) []BaseSkillFile {
	files := make([]BaseSkillFile, 0)
	for _, pkg := range packages {
		contents := make(map[string]string, len(pkg.Files))
		for _, file := range pkg.Files {
			contents[file.PackagePath] = string(file.Content)
		}
		for _, entry := range pkg.Entries {
			files = append(files, BaseSkillFile{
				Name:        entry.Name,
				PackageName: pkg.Name,
				Path:        entry.PackagePath,
				InstallPath: entry.InstallPath,
				Description: entry.Description,
				Task:        entry.Description,
				Content:     contents[entry.PackagePath],
			})
		}
	}
	return files
}

func InstallBaseSkillsTo(targetDir string) ([]string, error) {
	result, err := InstallSkillPackagesTo(targetDir)
	return result.PackageIDs, err
}

func InstallSkillPackagesTo(targetDir string) (SkillInstallResult, error) {
	packages, err := SkillPackages()
	if err != nil {
		return SkillInstallResult{}, err
	}
	return installSkillPackagesTo(targetDir, packages)
}

func InstallSkillPackagesForRepoTo(targetDir string, repoPath string) (SkillInstallResult, error) {
	packages, err := SkillPackagesForRepo(repoPath)
	if err != nil {
		return SkillInstallResult{}, err
	}
	return installSkillPackagesTo(targetDir, packages)
}

func (result SkillInstallResult) Summary() string {
	parts := []string{
		fmt.Sprintf("discovered=%d", result.Discovered),
		fmt.Sprintf("packages_installed=%d", result.Installed),
		fmt.Sprintf("packages_updated=%d", result.Updated),
		fmt.Sprintf("packages_skipped=%d", result.Skipped),
		fmt.Sprintf("files_written=%d", result.Written),
		fmt.Sprintf("files_overwritten=%d", result.Overwritten),
		fmt.Sprintf("files_deleted=%d", result.Deleted),
		fmt.Sprintf("files_skipped=%d", result.SkippedFiles),
		fmt.Sprintf("collisions=%d", result.Collisions),
		fmt.Sprintf("unsafe=%d", result.Unsafe),
	}
	return strings.Join(parts, " ")
}

func discoverSkillPackages(source fs.FS) ([]SkillPackage, error) {
	return discoverSkillPackagesFrom(skillPackageSource{
		filesystem: source,
		readRoot:   embeddedSkillSourceRoot,
		sourceRoot: embeddedSkillSourceRoot,
	})
}

func discoverSkillPackagesFrom(source skillPackageSource) ([]SkillPackage, error) {
	entries, err := fs.ReadDir(source.filesystem, source.readRoot)
	if err != nil {
		return nil, err
	}
	packages := make([]SkillPackage, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		pkg, err := readSkillPackage(source, name)
		if err != nil {
			return nil, err
		}
		packages = append(packages, pkg)
	}
	sort.Slice(packages, func(i, j int) bool {
		return packages[i].Name < packages[j].Name
	})
	return packages, nil
}

func runtimeSkillPackageSource(repoPath string) (skillPackageSource, bool, error) {
	if strings.TrimSpace(repoPath) == "" {
		return skillPackageSource{}, false, nil
	}
	repoRoot, err := filepath.Abs(repoPath)
	if err != nil {
		return skillPackageSource{}, false, err
	}
	skillsDir := filepath.Join(repoRoot, filepath.FromSlash(runtimeSkillSourceRoot))
	info, err := os.Stat(skillsDir)
	if err != nil {
		if os.IsNotExist(err) {
			return skillPackageSource{}, false, nil
		}
		return skillPackageSource{}, false, err
	}
	if !info.IsDir() {
		return skillPackageSource{}, false, fmt.Errorf("runtime skill source %s is not a directory", skillsDir)
	}
	return skillPackageSource{
		filesystem: os.DirFS(repoRoot),
		readRoot:   runtimeSkillSourceRoot,
		sourceRoot: runtimeSkillSourceRoot,
	}, true, nil
}

func readSkillPackage(source skillPackageSource, name string) (SkillPackage, error) {
	readRoot := path.Join(source.readRoot, name)
	sourceRoot := path.Join(source.sourceRoot, name)
	pkg := SkillPackage{
		Name:        name,
		SourceRoot:  sourceRoot,
		InstallRoot: name,
	}

	err := fs.WalkDir(source.filesystem, readRoot, func(filePath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if filePath == readRoot {
			return nil
		}
		rel := strings.TrimPrefix(filePath, readRoot+"/")
		if entry.IsDir() {
			return nil
		}
		if !entry.Type().IsRegular() {
			return nil
		}
		raw, err := fs.ReadFile(source.filesystem, filePath)
		if err != nil {
			return err
		}
		file := SkillPackageFile{
			SourcePath:  filePath,
			PackagePath: rel,
			InstallPath: path.Join(name, rel),
			Hash:        hashBytes(raw),
			SizeBytes:   int64(len(raw)),
			Content:     raw,
		}
		pkg.Files = append(pkg.Files, file)
		if path.Base(rel) == "SKILL.md" {
			frontmatter := parseSkillFrontmatter(string(raw))
			entryName := frontmatter["name"]
			if entryName == "" {
				entryName = deriveSkillEntryName(name, rel)
			}
			description := frontmatter["description"]
			if description == "" {
				description = deriveSkillDescription(string(raw), entryName)
			}
			pkg.Entries = append(pkg.Entries, SkillEntry{
				Name:        entryName,
				Description: description,
				SourcePath:  filePath,
				PackagePath: rel,
				InstallPath: path.Join(name, rel),
			})
		}
		return nil
	})
	if err != nil {
		return SkillPackage{}, err
	}
	if len(pkg.Entries) == 0 {
		return SkillPackage{}, fmt.Errorf("skill package %q has no SKILL.md entry", name)
	}
	sort.Slice(pkg.Files, func(i, j int) bool {
		return pkg.Files[i].PackagePath < pkg.Files[j].PackagePath
	})
	sort.Slice(pkg.Entries, func(i, j int) bool {
		return pkg.Entries[i].PackagePath < pkg.Entries[j].PackagePath
	})
	primary := primarySkillEntry(pkg)
	pkg.Description = primary.Description
	pkg.Hash = packageHash(pkg.Files)
	return pkg, nil
}

func installSkillPackagesTo(targetDir string, packages []SkillPackage) (SkillInstallResult, error) {
	result := SkillInstallResult{
		PackageIDs: make([]string, 0, len(packages)),
		Discovered: len(packages),
	}
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return result, err
	}
	manifest, err := loadSkillManifest(targetDir)
	if err != nil {
		return result, err
	}
	desired, collisions, err := desiredSkillSnapshot(packages)
	result.Collisions += collisions
	if err != nil {
		return result, err
	}
	actual, unsafe, err := actualSkillSnapshot(targetDir)
	result.Unsafe += unsafe
	if err != nil {
		return result, err
	}
	plan := diffSkillSnapshots(desired, actual)
	result.Written = len(plan.Writes)
	result.Overwritten = len(plan.Overwrites)
	result.Deleted = len(plan.Deletes)
	result.SkippedFiles = len(plan.Skips)

	for _, pkg := range packages {
		result.PackageIDs = append(result.PackageIDs, pkg.InstallRoot)
		oldEntry, hadManifest := manifest.Skills[pkg.Name]
		if !hadManifest {
			result.Installed++
			continue
		}
		if oldEntry.Stale || oldEntry.Hash != pkg.Hash || skillSyncPlanTouchesPackage(plan, pkg.InstallRoot) {
			result.Updated++
		} else {
			result.Skipped++
		}
	}

	if err := applySkillSyncPlan(targetDir, plan); err != nil {
		return result, err
	}
	verified, unsafe, err := actualSkillSnapshot(targetDir)
	result.Unsafe += unsafe
	if err != nil {
		return result, err
	}
	if err := verifySkillSnapshot(desired, verified); err != nil {
		return result, err
	}

	next := newSkillManifest()
	for _, pkg := range packages {
		next.Skills[pkg.Name] = manifestEntryForPackage(pkg, false)
	}
	if err := writeSkillManifest(targetDir, next); err != nil {
		return result, err
	}
	return result, nil
}

func desiredSkillSnapshot(packages []SkillPackage) (map[string]skillFileSnapshot, int, error) {
	snapshot := make(map[string]skillFileSnapshot)
	collisions := 0
	for _, pkg := range packages {
		for _, file := range pkg.Files {
			rel, err := cleanSkillInstallPath(file.InstallPath)
			if err != nil {
				return snapshot, collisions, fmt.Errorf("skill package %s has invalid install path %q: %w", pkg.Name, file.InstallPath, err)
			}
			if previous, ok := snapshot[rel]; ok {
				collisions++
				return snapshot, collisions, fmt.Errorf("duplicate generated skill path %s from packages %s and %s", rel, previous.PackageName, pkg.Name)
			}
			snapshot[rel] = skillFileSnapshot{
				Path:        rel,
				Hash:        file.Hash,
				Content:     file.Content,
				PackageName: pkg.Name,
				PackagePath: file.PackagePath,
			}
		}
	}
	return snapshot, collisions, nil
}

func actualSkillSnapshot(targetDir string) (map[string]skillFileSnapshot, int, error) {
	snapshot := make(map[string]skillFileSnapshot)
	unsafe := 0
	err := filepath.WalkDir(targetDir, func(filePath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if filePath == targetDir || entry.IsDir() {
			return nil
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			unsafe++
			return fmt.Errorf("non-regular skill output exists at %s", filePath)
		}
		rel, err := filepath.Rel(targetDir, filePath)
		if err != nil {
			return err
		}
		rel, err = cleanSkillInstallPath(rel)
		if err != nil {
			return fmt.Errorf("invalid skill output path %s: %w", filePath, err)
		}
		if rel == skillManifestFileName {
			return nil
		}
		raw, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		snapshot[rel] = skillFileSnapshot{
			Path: rel,
			Hash: hashBytes(raw),
		}
		return nil
	})
	if err != nil {
		if os.IsNotExist(err) {
			return snapshot, unsafe, nil
		}
		return snapshot, unsafe, err
	}
	return snapshot, unsafe, nil
}

func diffSkillSnapshots(desired map[string]skillFileSnapshot, actual map[string]skillFileSnapshot) skillSyncPlan {
	plan := skillSyncPlan{}
	desiredPaths := sortedSkillSnapshotPaths(desired)
	for _, rel := range desiredPaths {
		desiredFile := desired[rel]
		actualFile, ok := actual[rel]
		if !ok {
			plan.Writes = append(plan.Writes, desiredFile)
			continue
		}
		if actualFile.Hash != desiredFile.Hash {
			plan.Overwrites = append(plan.Overwrites, desiredFile)
			continue
		}
		plan.Skips = append(plan.Skips, desiredFile)
	}
	for _, rel := range sortedSkillSnapshotPaths(actual) {
		if _, ok := desired[rel]; !ok {
			plan.Deletes = append(plan.Deletes, actual[rel])
		}
	}
	return plan
}

func applySkillSyncPlan(targetDir string, plan skillSyncPlan) error {
	for _, file := range plan.Deletes {
		if err := deleteSkillSnapshotFile(targetDir, file.Path); err != nil {
			return err
		}
	}
	for _, file := range plan.Writes {
		if err := writeSkillSnapshotFile(targetDir, file); err != nil {
			return err
		}
	}
	for _, file := range plan.Overwrites {
		if err := writeSkillSnapshotFile(targetDir, file); err != nil {
			return err
		}
	}
	return removeEmptySkillDirs(targetDir)
}

func writeSkillSnapshotFile(targetDir string, file skillFileSnapshot) error {
	target, err := safeJoin(targetDir, file.Path)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(target, file.Content, 0o644); err != nil {
		return err
	}
	return nil
}

func deleteSkillSnapshotFile(targetDir string, rel string) error {
	target, err := safeJoin(targetDir, rel)
	if err != nil {
		return err
	}
	info, err := os.Lstat(target)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if info.IsDir() || !info.Mode().IsRegular() {
		return fmt.Errorf("refusing to delete non-regular skill output %s", target)
	}
	if err := os.Remove(target); err != nil {
		return err
	}
	removeEmptyDirsUpTo(targetDir, filepath.Dir(target))
	return nil
}

func removeEmptySkillDirs(targetDir string) error {
	dirs := make([]string, 0)
	err := filepath.WalkDir(targetDir, func(filePath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if filePath != targetDir && entry.IsDir() {
			dirs = append(dirs, filePath)
		}
		return nil
	})
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	sort.Slice(dirs, func(i, j int) bool {
		return len(dirs[i]) > len(dirs[j])
	})
	for _, dir := range dirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return err
		}
		if len(entries) == 0 {
			if err := os.Remove(dir); err != nil && !os.IsNotExist(err) {
				return err
			}
		}
	}
	return nil
}

func verifySkillSnapshot(desired map[string]skillFileSnapshot, actual map[string]skillFileSnapshot) error {
	if len(desired) != len(actual) {
		return fmt.Errorf("skill output file count mismatch: desired=%d actual=%d", len(desired), len(actual))
	}
	for _, rel := range sortedSkillSnapshotPaths(desired) {
		desiredFile := desired[rel]
		actualFile, ok := actual[rel]
		if !ok {
			return fmt.Errorf("skill output missing desired file %s", rel)
		}
		if actualFile.Hash != desiredFile.Hash {
			return fmt.Errorf("skill output hash mismatch for %s: desired=%s actual=%s", rel, desiredFile.Hash, actualFile.Hash)
		}
	}
	for _, rel := range sortedSkillSnapshotPaths(actual) {
		if _, ok := desired[rel]; !ok {
			return fmt.Errorf("skill output contains obsolete file %s", rel)
		}
	}
	return nil
}

func sortedSkillSnapshotPaths(snapshot map[string]skillFileSnapshot) []string {
	paths := make([]string, 0, len(snapshot))
	for rel := range snapshot {
		paths = append(paths, rel)
	}
	sort.Strings(paths)
	return paths
}

func skillSyncPlanTouchesPackage(plan skillSyncPlan, installRoot string) bool {
	for _, file := range plan.Writes {
		if skillPathInPackage(file.Path, installRoot) {
			return true
		}
	}
	for _, file := range plan.Overwrites {
		if skillPathInPackage(file.Path, installRoot) {
			return true
		}
	}
	for _, file := range plan.Deletes {
		if skillPathInPackage(file.Path, installRoot) {
			return true
		}
	}
	return false
}

func skillPathInPackage(rel string, installRoot string) bool {
	root, err := cleanSkillInstallPath(installRoot)
	if err != nil {
		return false
	}
	return rel == root || strings.HasPrefix(rel, root+"/")
}

func cleanSkillInstallPath(rel string) (string, error) {
	if strings.TrimSpace(rel) == "" {
		return "", fmt.Errorf("empty path")
	}
	if filepath.IsAbs(rel) {
		return "", fmt.Errorf("absolute path")
	}
	clean := path.Clean(filepath.ToSlash(rel))
	if clean == "." || clean == ".." || strings.HasPrefix(clean, "../") || strings.HasPrefix(clean, "/") {
		return "", fmt.Errorf("path escapes target root")
	}
	return clean, nil
}

func loadSkillManifest(targetDir string) (skillManifest, error) {
	manifest := newSkillManifest()
	raw, err := os.ReadFile(filepath.Join(targetDir, skillManifestFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return manifest, nil
		}
		return manifest, err
	}
	if err := json.Unmarshal(raw, &manifest); err != nil {
		return skillManifest{}, err
	}
	if manifest.ManagedBy != "" && manifest.ManagedBy != skillManifestOwner {
		return skillManifest{}, fmt.Errorf("skill manifest is managed by %q, not %q", manifest.ManagedBy, skillManifestOwner)
	}
	if manifest.SchemaVersion == 0 {
		manifest.SchemaVersion = skillManifestVersion
	}
	if manifest.ManagedBy == "" {
		manifest.ManagedBy = skillManifestOwner
	}
	if manifest.Skills == nil {
		manifest.Skills = map[string]skillManifestEntry{}
	}
	return manifest, nil
}

func writeSkillManifest(targetDir string, manifest skillManifest) error {
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return err
	}
	raw, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(targetDir, skillManifestFileName), append(raw, '\n'), 0o644)
}

func newSkillManifest() skillManifest {
	return skillManifest{
		SchemaVersion: skillManifestVersion,
		ManagedBy:     skillManifestOwner,
		Skills:        map[string]skillManifestEntry{},
	}
}

func manifestEntryForPackage(pkg SkillPackage, stale bool) skillManifestEntry {
	files := make(map[string]string, len(pkg.Files))
	for _, file := range pkg.Files {
		files[file.PackagePath] = file.Hash
	}
	return skillManifestEntry{
		InstallPath: pkg.InstallRoot,
		SourceRoot:  pkg.SourceRoot,
		Hash:        pkg.Hash,
		Managed:     true,
		Stale:       stale,
		EntryCount:  len(pkg.Entries),
		FileCount:   len(pkg.Files),
		Files:       files,
	}
}

func safeJoin(base string, rel string) (string, error) {
	target := filepath.Join(base, filepath.FromSlash(rel))
	absBase, err := filepath.Abs(base)
	if err != nil {
		return "", err
	}
	absTarget, err := filepath.Abs(target)
	if err != nil {
		return "", err
	}
	cleanBase := filepath.Clean(absBase)
	cleanTarget := filepath.Clean(absTarget)
	if cleanTarget != cleanBase && !strings.HasPrefix(cleanTarget, cleanBase+string(os.PathSeparator)) {
		return "", fmt.Errorf("path %s escapes target root %s", cleanTarget, cleanBase)
	}
	return target, nil
}

func removeEmptyDirsUpTo(root string, dir string) {
	root = filepath.Clean(root)
	dir = filepath.Clean(dir)
	for dir != root && strings.HasPrefix(dir, root+string(os.PathSeparator)) {
		entries, err := os.ReadDir(dir)
		if err != nil || len(entries) > 0 {
			return
		}
		if err := os.Remove(dir); err != nil {
			return
		}
		dir = filepath.Dir(dir)
	}
}

func packageHash(files []SkillPackageFile) string {
	h := sha256.New()
	for _, file := range files {
		h.Write([]byte(file.PackagePath))
		h.Write([]byte{0})
		h.Write([]byte(file.Hash))
		h.Write([]byte{0})
	}
	return "sha256:" + hex.EncodeToString(h.Sum(nil))
}

func hashBytes(raw []byte) string {
	sum := sha256.Sum256(raw)
	return "sha256:" + hex.EncodeToString(sum[:])
}

func primarySkillEntry(pkg SkillPackage) SkillEntry {
	for _, entry := range pkg.Entries {
		if entry.PackagePath == "SKILL.md" {
			return entry
		}
	}
	parentPath := path.Join(pkg.Name+"-parent-skill", "SKILL.md")
	for _, entry := range pkg.Entries {
		if entry.PackagePath == parentPath {
			return entry
		}
	}
	return pkg.Entries[0]
}

func deriveSkillEntryName(packageName string, packagePath string) string {
	dir := path.Base(path.Dir(packagePath))
	if dir == "." || dir == "/" {
		return packageName
	}
	return dir
}

func deriveSkillDescription(content string, entryName string) string {
	content = strings.ReplaceAll(content, "\r\n", "\n")
	if strings.HasPrefix(content, "---\n") {
		if end := strings.Index(content[4:], "\n---"); end >= 0 {
			content = content[4+end+len("\n---"):]
		}
	}
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(strings.TrimPrefix(line, "#"))
		if line == "" {
			continue
		}
		if strings.EqualFold(line, entryName) {
			continue
		}
		return line
	}
	return fmt.Sprintf("Use the %s skill.", entryName)
}

func parseSkillFrontmatter(content string) map[string]string {
	out := map[string]string{}
	content = strings.ReplaceAll(content, "\r\n", "\n")
	if !strings.HasPrefix(content, "---\n") {
		return out
	}
	end := strings.Index(content[4:], "\n---")
	if end < 0 {
		return out
	}
	for _, line := range strings.Split(content[4:4+end], "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}
		out[strings.TrimSpace(key)] = cleanFrontmatterValue(value)
	}
	return out
}

func cleanFrontmatterValue(value string) string {
	value = strings.TrimSpace(value)
	value = strings.Trim(value, `"'`)
	return strings.TrimSpace(value)
}

func skillGuideEntries(pkg SkillPackage) string {
	entries := make([]string, 0, len(pkg.Entries))
	for _, entry := range pkg.Entries {
		entries = append(entries, "`.claude/skills/anvien/"+escapeTableCell(entry.InstallPath)+"`")
	}
	return strings.Join(entries, "<br>")
}

func skillGuideNeed(pkg SkillPackage) string {
	description := strings.TrimSpace(pkg.Description)
	if description == "" {
		description = fmt.Sprintf("Use the %s skill package.", pkg.Name)
	}
	return escapeTableCell(description)
}

func skillGuideUse(pkg SkillPackage) string {
	entry := primarySkillEntry(pkg)
	return "`.claude/skills/anvien/" + escapeTableCell(entry.InstallPath) + "`"
}

func escapeTableCell(value string) string {
	value = strings.ReplaceAll(value, "\r\n", " ")
	value = strings.ReplaceAll(value, "\n", " ")
	value = strings.ReplaceAll(value, "|", "\\|")
	return strings.TrimSpace(value)
}
