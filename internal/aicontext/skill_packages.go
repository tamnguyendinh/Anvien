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
	skillManifestFileName = ".anvien-skill-manifest.json"
	skillManifestVersion  = 1
	skillManifestOwner    = "anvien"
)

//go:embed skills
//go:embed skills/*/.env.example
//go:embed skills/*/*/.env.example
var skillSourceFS embed.FS

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
	PackageIDs []string
	Discovered int
	Installed  int
	Updated    int
	Skipped    int
	Adopted    int
	Stale      int
	Preserved  int
	Collisions int
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

func BaseSkillFiles() ([]BaseSkillFile, error) {
	packages, err := SkillPackages()
	if err != nil {
		return nil, err
	}
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
	return files, nil
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

func (result SkillInstallResult) Summary() string {
	parts := []string{
		fmt.Sprintf("discovered=%d", result.Discovered),
		fmt.Sprintf("installed=%d", result.Installed),
		fmt.Sprintf("updated=%d", result.Updated),
		fmt.Sprintf("skipped=%d", result.Skipped),
		fmt.Sprintf("adopted=%d", result.Adopted),
		fmt.Sprintf("stale=%d", result.Stale),
		fmt.Sprintf("preserved=%d", result.Preserved),
		fmt.Sprintf("collisions=%d", result.Collisions),
	}
	return strings.Join(parts, " ")
}

func discoverSkillPackages(source fs.FS) ([]SkillPackage, error) {
	entries, err := fs.ReadDir(source, "skills")
	if err != nil {
		return nil, err
	}
	packages := make([]SkillPackage, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		if shouldSkipSkillPackagePath(name) {
			continue
		}
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

func readSkillPackage(source fs.FS, name string) (SkillPackage, error) {
	sourceRoot := path.Join("skills", name)
	pkg := SkillPackage{
		Name:        name,
		SourceRoot:  sourceRoot,
		InstallRoot: name,
	}

	err := fs.WalkDir(source, sourceRoot, func(filePath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if filePath == sourceRoot {
			return nil
		}
		rel := strings.TrimPrefix(filePath, sourceRoot+"/")
		if entry.IsDir() {
			if shouldSkipSkillPackagePath(entry.Name()) {
				return fs.SkipDir
			}
			return nil
		}
		if !entry.Type().IsRegular() {
			return nil
		}
		if shouldSkipSkillPackagePath(entry.Name()) {
			return nil
		}
		raw, err := fs.ReadFile(source, filePath)
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
	next := newSkillManifest()
	sourceNames := make(map[string]struct{}, len(packages))

	for _, pkg := range packages {
		sourceNames[pkg.Name] = struct{}{}
	}
	result.Preserved = countPreservedSkillTargets(targetDir, manifest, sourceNames)

	for _, pkg := range packages {
		oldEntry, hadManifest := manifest.Skills[pkg.Name]
		targetPkg, err := safeJoin(targetDir, pkg.InstallRoot)
		if err != nil {
			return result, err
		}

		if !hadManifest && pathExists(targetPkg) {
			matches, err := targetMatchesPackage(targetPkg, pkg)
			if err != nil {
				return result, err
			}
			if !matches {
				legacy, err := isLegacyGeneratedAnvienPackage(targetPkg, pkg)
				if err != nil {
					return result, err
				}
				if legacy {
					if err := writeSkillPackage(targetDir, pkg); err != nil {
						return result, err
					}
					next.Skills[pkg.Name] = manifestEntryForPackage(pkg, false)
					result.PackageIDs = append(result.PackageIDs, pkg.InstallRoot)
					result.Adopted++
					continue
				}
				result.Collisions++
				return result, fmt.Errorf("skill package target %s already exists and is not managed by Anvien; preserving existing files", targetPkg)
			}
			next.Skills[pkg.Name] = manifestEntryForPackage(pkg, false)
			result.PackageIDs = append(result.PackageIDs, pkg.InstallRoot)
			result.Adopted++
			continue
		}

		if hadManifest && oldEntry.Hash == pkg.Hash {
			matches, err := packageFilesMatch(targetPkg, pkg)
			if err != nil {
				return result, err
			}
			if matches {
				next.Skills[pkg.Name] = manifestEntryForPackage(pkg, false)
				result.PackageIDs = append(result.PackageIDs, pkg.InstallRoot)
				result.Skipped++
				continue
			}
		}

		if hadManifest {
			if err := removeRetiredManagedFiles(targetDir, oldEntry, pkg); err != nil {
				return result, err
			}
		}
		if err := writeSkillPackage(targetDir, pkg); err != nil {
			return result, err
		}
		next.Skills[pkg.Name] = manifestEntryForPackage(pkg, false)
		result.PackageIDs = append(result.PackageIDs, pkg.InstallRoot)
		if hadManifest {
			result.Updated++
		} else {
			result.Installed++
		}
	}

	for name, entry := range manifest.Skills {
		if _, ok := sourceNames[name]; ok {
			continue
		}
		if entry.Managed {
			entry.Stale = true
			result.Stale++
		}
		next.Skills[name] = entry
	}
	if err := writeSkillManifest(targetDir, next); err != nil {
		return result, err
	}
	return result, nil
}

func writeSkillPackage(targetDir string, pkg SkillPackage) error {
	for _, file := range pkg.Files {
		target, err := safeJoin(targetDir, file.InstallPath)
		if err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(target, file.Content, 0o644); err != nil {
			return err
		}
	}
	return nil
}

func removeRetiredManagedFiles(targetDir string, oldEntry skillManifestEntry, pkg SkillPackage) error {
	if oldEntry.InstallPath == "" {
		oldEntry.InstallPath = pkg.InstallRoot
	}
	currentFiles := make(map[string]struct{}, len(pkg.Files))
	for _, file := range pkg.Files {
		currentFiles[file.PackagePath] = struct{}{}
	}
	for rel := range oldEntry.Files {
		if _, ok := currentFiles[rel]; ok {
			continue
		}
		target, err := safeJoin(targetDir, path.Join(oldEntry.InstallPath, rel))
		if err != nil {
			return err
		}
		info, err := os.Lstat(target)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return err
		}
		if info.IsDir() {
			continue
		}
		if err := os.Remove(target); err != nil {
			return err
		}
		removeEmptyDirsUpTo(targetDir, filepath.Dir(target))
	}
	return nil
}

func countPreservedSkillTargets(targetDir string, manifest skillManifest, sourceNames map[string]struct{}) int {
	entries, err := os.ReadDir(targetDir)
	if err != nil {
		return 0
	}
	count := 0
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		if _, ok := manifest.Skills[name]; ok {
			continue
		}
		if _, ok := sourceNames[name]; ok {
			continue
		}
		count++
	}
	return count
}

func isLegacyGeneratedAnvienPackage(targetPkg string, pkg SkillPackage) (bool, error) {
	if !legacyGeneratedAnvienPackageNames()[pkg.Name] {
		return false, nil
	}
	raw, err := os.ReadFile(filepath.Join(targetPkg, "SKILL.md"))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	frontmatter := parseSkillFrontmatter(string(raw))
	return frontmatter["name"] == pkg.Name, nil
}

func legacyGeneratedAnvienPackageNames() map[string]bool {
	return map[string]bool{
		"anvien-api-surface": true,
		"anvien-debugging":   true,
		"anvien-planner":     true,
		"anvien-refactoring": true,
	}
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

func packageFilesMatch(targetPkg string, pkg SkillPackage) (bool, error) {
	for _, file := range pkg.Files {
		target := filepath.Join(targetPkg, filepath.FromSlash(file.PackagePath))
		raw, err := os.ReadFile(target)
		if err != nil {
			if os.IsNotExist(err) {
				return false, nil
			}
			return false, err
		}
		if hashBytes(raw) != file.Hash {
			return false, nil
		}
	}
	return true, nil
}

func targetMatchesPackage(targetPkg string, pkg SkillPackage) (bool, error) {
	matches, err := packageFilesMatch(targetPkg, pkg)
	if err != nil || !matches {
		return matches, err
	}
	expected := make(map[string]struct{}, len(pkg.Files))
	for _, file := range pkg.Files {
		expected[filepath.ToSlash(file.PackagePath)] = struct{}{}
	}
	err = filepath.WalkDir(targetPkg, func(filePath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return fmt.Errorf("non-regular file exists at %s", filePath)
		}
		rel, err := filepath.Rel(targetPkg, filePath)
		if err != nil {
			return err
		}
		if _, ok := expected[filepath.ToSlash(rel)]; !ok {
			return fmt.Errorf("extra file exists at %s", filePath)
		}
		return nil
	})
	if err != nil {
		return false, nil
	}
	return true, nil
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

func pathExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
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

func shouldSkipSkillPackagePath(name string) bool {
	switch name {
	case ".git", ".hg", ".svn", ".pytest_cache", "__pycache__", "node_modules", "dist", "build", "coverage", ".next":
		return true
	default:
		return false
	}
}

func skillGuideEntries(pkg SkillPackage) string {
	entries := make([]string, 0, len(pkg.Entries))
	for _, entry := range pkg.Entries {
		entries = append(entries, "`.claude/skills/anvien/"+escapeTableCell(entry.InstallPath)+"`")
	}
	return strings.Join(entries, "<br>")
}

func skillGuideUse(pkg SkillPackage) string {
	description := strings.TrimSpace(pkg.Description)
	if description == "" {
		description = fmt.Sprintf("Use the %s skill package.", pkg.Name)
	}
	root := "`.claude/skills/anvien/" + escapeTableCell(pkg.InstallRoot) + "/`"
	return escapeTableCell(description) + "<br>" + root
}

func escapeTableCell(value string) string {
	value = strings.ReplaceAll(value, "\r\n", " ")
	value = strings.ReplaceAll(value, "\n", " ")
	value = strings.ReplaceAll(value, "|", "\\|")
	return strings.TrimSpace(value)
}
