package httpapi

import (
	"errors"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/tamnguyendinh/anvien/internal/filecontext"
	"github.com/tamnguyendinh/anvien/internal/repo"
)

const (
	fileContextDefaultSampleLimit = 10
	fileContextMaxSampleLimit     = 100
	fileHotspotsDefaultLimit      = 20
	fileHotspotsMaxLimit          = 500
)

type fileHotspotsResponse struct {
	Repo     string                    `json:"repo,omitempty"`
	RepoPath string                    `json:"repoPath,omitempty"`
	Graph    filecontext.GraphInfo     `json:"graph"`
	Total    int                       `json:"total"`
	Offset   int                       `json:"offset"`
	Limit    int                       `json:"limit"`
	Sort     string                    `json:"sort"`
	Files    []filecontext.FileSummary `json:"files"`
}

type fileProjectionGraph struct {
	repoName string
	repoPath string
	graph    filecontext.GraphInfo
	builder  *filecontext.Builder
}

type fileProjectionServerCache struct {
	mu      sync.Mutex
	entries map[fileProjectionServerCacheKey]fileProjectionGraph
}

type fileProjectionServerCacheKey struct {
	Repo          string
	RepoPath      string
	GraphPath     string
	GraphModTime  int64
	GraphSize     int64
	IndexedCommit string
	CurrentCommit string
}

func newFileProjectionServerCache() *fileProjectionServerCache {
	return &fileProjectionServerCache{entries: map[fileProjectionServerCacheKey]fileProjectionGraph{}}
}

func (c *fileProjectionServerCache) Get(key fileProjectionServerCacheKey) (fileProjectionGraph, bool) {
	if c == nil {
		return fileProjectionGraph{}, false
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	projection, ok := c.entries[key]
	return projection, ok
}

func (c *fileProjectionServerCache) Set(key fileProjectionServerCacheKey, projection fileProjectionGraph) {
	if c == nil {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	for existing := range c.entries {
		if existing.Repo == key.Repo && existing.RepoPath == key.RepoPath && existing.GraphPath == key.GraphPath {
			delete(c.entries, existing)
		}
	}
	c.entries[key] = projection
}

func (s Server) handleFileContext(w http.ResponseWriter, r *http.Request) {
	if !methodAllowed(w, r, http.MethodGet) {
		return
	}

	path := strings.TrimSpace(r.URL.Query().Get("path"))
	if path == "" {
		writeError(w, http.StatusBadRequest, `Missing "path" query parameter`)
		return
	}

	projection, status, message, err := s.loadFileProjection(r)
	if err != nil {
		writeError(w, status, message)
		return
	}
	context, ok := projection.builder.BuildFileContext(path, filecontext.Options{
		RelationshipSamplesPerGroup: boundedNonNegativeQueryInt(r.URL.Query().Get("relationships"), fileContextDefaultSampleLimit, 1, fileContextMaxSampleLimit),
		UnresolvedSamplesPerGroup:   boundedNonNegativeQueryInt(r.URL.Query().Get("unresolved"), fileContextDefaultSampleLimit, 1, fileContextMaxSampleLimit),
		LinkedSamplesPerKind:        boundedNonNegativeQueryInt(r.URL.Query().Get("linked"), fileContextDefaultSampleLimit, 1, fileContextMaxSampleLimit),
	})
	if !ok {
		writeError(w, http.StatusNotFound, "File not found in graph")
		return
	}
	filecontext.AttachMetadata(&context, projection.repoName, projection.repoPath, projection.graph)
	writeJSON(w, http.StatusOK, context)
}

func (s Server) handleFileHotspots(w http.ResponseWriter, r *http.Request) {
	if !methodAllowed(w, r, http.MethodGet) {
		return
	}

	projection, status, message, err := s.loadFileProjection(r)
	if err != nil {
		writeError(w, status, message)
		return
	}
	sortMode := strings.TrimSpace(r.URL.Query().Get("sort"))
	if !filecontext.IsSupportedFileListSort(sortMode) {
		writeError(w, http.StatusBadRequest, `Unsupported "sort" query parameter`)
		return
	}
	sortMode = filecontext.NormalizeFileListSort(sortMode)
	list := projection.builder.BuildFileList(filecontext.FileListOptions{
		Sort:                sortMode,
		Limit:               boundedNonNegativeQueryInt(r.URL.Query().Get("limit"), fileHotspotsDefaultLimit, 0, fileHotspotsMaxLimit),
		Offset:              boundedNonNegativeQueryInt(r.URL.Query().Get("offset"), 0, 0, 0),
		Kinds:               queryStringList(r, "kind"),
		AppLayer:            strings.TrimSpace(r.URL.Query().Get("appLayer")),
		FunctionalArea:      strings.TrimSpace(r.URL.Query().Get("functionalArea")),
		APIOnly:             queryBool(r, "apiOnly"),
		ChangedOnly:         queryBool(r, "changedOnly"),
		UnresolvedOnly:      queryBool(r, "unresolvedOnly"),
		HighFanInOnly:       queryBool(r, "highFanIn"),
		HighFanOutOnly:      queryBool(r, "highFanOut"),
		HighFanInThreshold:  boundedNonNegativeQueryInt(r.URL.Query().Get("highFanInThreshold"), 10, 1, 0),
		HighFanOutThreshold: boundedNonNegativeQueryInt(r.URL.Query().Get("highFanOutThreshold"), 10, 1, 0),
		ChangedPaths:        changedFilesSinceAnalyze(projection.repoPath, projection.graph.IndexedCommit),
		Stale:               projection.graph.Stale,
	})
	writeJSON(w, http.StatusOK, fileHotspotsResponse{
		Repo:     projection.repoName,
		RepoPath: projection.repoPath,
		Graph:    projection.graph,
		Total:    list.Total,
		Offset:   list.Offset,
		Limit:    list.Limit,
		Sort:     list.Sort,
		Files:    list.Files,
	})
}

func changedFilesSinceAnalyze(repoPath string, indexedCommit string) map[string]struct{} {
	changed := map[string]struct{}{}
	collectGitPathList(repoPath, changed, "diff", "--name-only", "HEAD", "--")
	collectGitPathList(repoPath, changed, "ls-files", "--others", "--exclude-standard")
	if strings.TrimSpace(indexedCommit) != "" {
		collectGitPathList(repoPath, changed, "diff", "--name-only", indexedCommit, "HEAD", "--")
	}
	if len(changed) == 0 {
		return nil
	}
	return changed
}

func collectGitPathList(repoPath string, out map[string]struct{}, args ...string) {
	if strings.TrimSpace(repoPath) == "" || out == nil {
		return
	}
	gitArgs := append([]string{"-C", repoPath}, args...)
	output, err := exec.Command("git", gitArgs...).Output()
	if err != nil {
		return
	}
	for _, line := range strings.Split(string(output), "\n") {
		path := strings.TrimSpace(strings.ReplaceAll(line, "\\", "/"))
		if path != "" {
			out[path] = struct{}{}
		}
	}
}

func (s Server) loadFileProjection(r *http.Request) (fileProjectionGraph, int, string, error) {
	entry, status, message, err := s.resolveRequestedRepo(r)
	if err != nil {
		if status == http.StatusNotFound {
			message = "Repository not found"
		}
		return fileProjectionGraph{}, status, message, err
	}

	graphPath := filepath.Join(storagePathFor(entry), "graph.json")
	stat, err := os.Stat(graphPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fileProjectionGraph{}, http.StatusNotFound, "Graph not found. Run: anvien analyze", err
		}
		return fileProjectionGraph{}, http.StatusInternalServerError, err.Error(), err
	}
	currentCommit := repo.CurrentCommit(entry.Path)
	cacheKey := fileProjectionServerCacheKey{
		Repo:          entry.Name,
		RepoPath:      entry.Path,
		GraphPath:     graphPath,
		GraphModTime:  stat.ModTime().UnixNano(),
		GraphSize:     stat.Size(),
		IndexedCommit: entry.LastCommit,
		CurrentCommit: currentCommit,
	}
	if projection, ok := s.fileProjections.Get(cacheKey); ok {
		return projection, http.StatusOK, "", nil
	}

	g, err := loadGraphSnapshot(graphPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fileProjectionGraph{}, http.StatusNotFound, "Graph not found. Run: anvien analyze", err
		}
		return fileProjectionGraph{}, http.StatusInternalServerError, err.Error(), err
	}

	info := fileProjectionGraphInfo(entry, graphPath, currentCommit)
	builder, _ := s.fileContexts.Get(filecontext.CacheKey{
		Repo:      entry.Name,
		RepoPath:  entry.Path,
		GraphPath: graphPath,
		GraphHash: filecontext.GraphFingerprint(g),
	}, g)
	projection := fileProjectionGraph{
		repoName: entry.Name,
		repoPath: entry.Path,
		graph:    info,
		builder:  builder,
	}
	s.fileProjections.Set(cacheKey, projection)
	return projection, http.StatusOK, "", nil
}

func fileProjectionGraphInfo(entry repo.RegistryEntry, graphPath string, currentCommit string) filecontext.GraphInfo {
	return filecontext.GraphInfo{
		Path:          graphPath,
		IndexedCommit: entry.LastCommit,
		CurrentCommit: currentCommit,
		Stale:         currentCommit != "" && entry.LastCommit != "" && currentCommit != entry.LastCommit,
		AnalyzedAt:    entry.IndexedAt,
	}
}

func boundedNonNegativeQueryInt(value string, fallback int, min int, max int) int {
	value = strings.TrimSpace(value)
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	if parsed < min {
		return min
	}
	if max > 0 && parsed > max {
		return max
	}
	return parsed
}

func queryBool(r *http.Request, name string) bool {
	parsed, err := strconv.ParseBool(strings.TrimSpace(r.URL.Query().Get(name)))
	return err == nil && parsed
}

func queryStringList(r *http.Request, name string) []string {
	values := r.URL.Query()[name]
	out := make([]string, 0, len(values))
	for _, value := range values {
		for _, part := range strings.Split(value, ",") {
			part = strings.TrimSpace(part)
			if part != "" {
				out = append(out, part)
			}
		}
	}
	return out
}
