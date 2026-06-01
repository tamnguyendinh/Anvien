package analyze

import (
	"path"
	"strings"

	"github.com/tamnguyendinh/anvien/internal/documents"
	"github.com/tamnguyendinh/anvien/internal/scanner"
)

const fileClassificationSampleLimit = 5

type FileClassificationBucket string

const (
	FileBucketParsedCode          FileClassificationBucket = "parsedCode"
	FileBucketDocuments           FileClassificationBucket = "documents"
	FileBucketMetadataOnly        FileClassificationBucket = "metadataOnly"
	FileBucketScriptNoExtractor   FileClassificationBucket = "scriptNoExtractor"
	FileBucketStaticAssets        FileClassificationBucket = "staticAssets"
	FileBucketUnsupportedLanguage FileClassificationBucket = "unsupportedLanguage"
	FileBucketUnknown             FileClassificationBucket = "unknown"
	FileBucketFailed              FileClassificationBucket = "failed"
)

type FileClassificationSample struct {
	Bucket   FileClassificationBucket `json:"bucket"`
	Path     string                   `json:"path"`
	Language string                   `json:"language,omitempty"`
	Reason   string                   `json:"reason,omitempty"`
}

type fileClassificationOutcome struct {
	Parsed              map[string]struct{}
	Failed              map[string]struct{}
	UnsupportedLanguage map[string]struct{}
}

func classifyFileMetrics(files []scanner.File, outcome fileClassificationOutcome) FileMetrics {
	metrics := FileMetrics{Scanned: len(files)}
	sampler := newFileClassificationSampler()
	for _, file := range files {
		bucket, reason := classifyFile(file, outcome)
		metrics.add(bucket)
		if bucket != FileBucketParsedCode {
			sampler.add(file, bucket, reason)
		}
	}
	metrics.Parsed = metrics.ParsedCode
	metrics.Unsupported = metrics.UnsupportedLanguage
	metrics.ClassificationSamples = sampler.samples
	return metrics
}

func classifyFile(file scanner.File, outcome fileClassificationOutcome) (FileClassificationBucket, string) {
	normalizedPath := normalizeClassificationPath(file.Path)
	switch {
	case containsClassificationPath(outcome.Failed, normalizedPath):
		return FileBucketFailed, "processing failed"
	case containsClassificationPath(outcome.Parsed, normalizedPath):
		return FileBucketParsedCode, "ScopeIR produced"
	case documents.Kind(normalizedPath) != "":
		return FileBucketDocuments, "document indexing phase"
	case metadataOnlyFile(normalizedPath):
		return FileBucketMetadataOnly, "metadata or structured non-code input"
	case scriptNoExtractorFile(normalizedPath):
		return FileBucketScriptNoExtractor, "script-like file without ScopeIR extractor"
	case staticAssetFile(normalizedPath):
		return FileBucketStaticAssets, "static asset or Web artifact"
	case containsClassificationPath(outcome.UnsupportedLanguage, normalizedPath):
		return FileBucketUnsupportedLanguage, "parser reported unsupported language"
	case file.Language != "" && !hasExtractor(file.Language):
		return FileBucketUnsupportedLanguage, "recognized language has no ScopeIR extractor"
	default:
		return FileBucketUnknown, "no classifier matched"
	}
}

func (metrics *FileMetrics) add(bucket FileClassificationBucket) {
	switch bucket {
	case FileBucketParsedCode:
		metrics.ParsedCode++
	case FileBucketDocuments:
		metrics.Documents++
	case FileBucketMetadataOnly:
		metrics.MetadataOnly++
	case FileBucketScriptNoExtractor:
		metrics.ScriptNoExtractor++
	case FileBucketStaticAssets:
		metrics.StaticAssets++
	case FileBucketUnsupportedLanguage:
		metrics.UnsupportedLanguage++
	case FileBucketUnknown:
		metrics.Unknown++
	case FileBucketFailed:
		metrics.Failed++
	}
}

func (metrics FileMetrics) ClassifiedTotal() int {
	return metrics.ParsedCode +
		metrics.Documents +
		metrics.MetadataOnly +
		metrics.ScriptNoExtractor +
		metrics.StaticAssets +
		metrics.UnsupportedLanguage +
		metrics.Unknown +
		metrics.Failed
}

type fileClassificationSampler struct {
	samples []FileClassificationSample
	seen    map[FileClassificationBucket]int
}

func newFileClassificationSampler() fileClassificationSampler {
	return fileClassificationSampler{seen: map[FileClassificationBucket]int{}}
}

func (sampler *fileClassificationSampler) add(file scanner.File, bucket FileClassificationBucket, reason string) {
	if sampler.seen[bucket] >= fileClassificationSampleLimit {
		return
	}
	sampler.seen[bucket]++
	sampler.samples = append(sampler.samples, FileClassificationSample{
		Bucket:   bucket,
		Path:     normalizeClassificationPath(file.Path),
		Language: string(file.Language),
		Reason:   reason,
	})
}

func containsClassificationPath(paths map[string]struct{}, filePath string) bool {
	if len(paths) == 0 {
		return false
	}
	_, ok := paths[filePath]
	return ok
}

func addClassificationPath(paths map[string]struct{}, filePath string) {
	paths[normalizeClassificationPath(filePath)] = struct{}{}
}

func normalizeClassificationPath(filePath string) string {
	return strings.ReplaceAll(filePath, "\\", "/")
}

func metadataOnlyFile(filePath string) bool {
	base := strings.ToLower(path.Base(filePath))
	ext := strings.ToLower(path.Ext(filePath))
	if strings.HasPrefix(base, "dockerfile") ||
		base == "go.mod" ||
		base == "go.sum" ||
		strings.HasPrefix(base, ".env") {
		return true
	}
	switch base {
	case "package.json", "package-lock.json", "pnpm-lock.yaml", "yarn.lock",
		"tsconfig.json", "jsconfig.json", "docker-compose.yaml", "docker-compose.yml",
		".mcp.json", ".editorconfig", ".gitattributes", ".gitignore", ".dockerignore",
		".npmrc", ".nvmrc", ".prettierrc", ".eslintrc", ".babelrc", ".browserslistrc":
		return true
	}
	switch ext {
	case ".json", ".jsonl", ".yaml", ".yml", ".toml", ".xml", ".mod", ".sum",
		".lock", ".conf", ".config", ".ini", ".properties":
		return true
	default:
		return false
	}
}

func scriptNoExtractorFile(filePath string) bool {
	base := strings.ToLower(path.Base(filePath))
	ext := strings.ToLower(path.Ext(filePath))
	switch base {
	case "makefile", "rakefile":
		return true
	}
	switch ext {
	case ".sh", ".bash", ".zsh", ".fish", ".ps1", ".psm1", ".bat", ".cmd",
		".mk", ".sql":
		return true
	default:
		return false
	}
}

func staticAssetFile(filePath string) bool {
	ext := strings.ToLower(path.Ext(filePath))
	switch ext {
	case ".html", ".htm", ".css", ".scss", ".sass", ".less", ".svg",
		".png", ".jpg", ".jpeg", ".gif", ".webp", ".ico", ".map",
		".woff", ".woff2", ".ttf", ".eot":
		return true
	default:
		return false
	}
}
