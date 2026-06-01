package analyze

import (
	"encoding/json"
	"testing"

	"github.com/tamnguyendinh/anvien/internal/scanner"
)

func TestClassifyFileMetricsUsesCausalBuckets(t *testing.T) {
	parsed := map[string]struct{}{}
	failed := map[string]struct{}{}
	addClassificationPath(parsed, "src/main.go")
	addClassificationPath(failed, "broken/input.ts")

	files := []scanner.File{
		{Path: "src/main.go", Language: scanner.Go},
		{Path: "README.md", Language: scanner.Markdown},
		{Path: "data/sheet.csv", Language: scanner.Spreadsheet},
		{Path: "package.json"},
		{Path: "testdata/golden/response.json"},
		{Path: "go.mod"},
		{Path: "config/app.yaml"},
		{Path: "scripts/build.ps1"},
		{Path: "scripts/run.sh"},
		{Path: "web/index.html"},
		{Path: "web/index.css"},
		{Path: "mainframe/program.cob", Language: scanner.Cobol},
		{Path: "notes/blob.weird"},
		{Path: "broken/input.ts", Language: scanner.TypeScript},
	}

	metrics := classifyFileMetrics(files, fileClassificationOutcome{
		Parsed: parsed,
		Failed: failed,
	})

	if metrics.Scanned != len(files) || metrics.ClassifiedTotal() != metrics.Scanned {
		t.Fatalf("metrics do not reconcile: %#v", metrics)
	}
	assertFileMetric(t, "parsedCode", metrics.ParsedCode, 1)
	assertFileMetric(t, "parsed legacy alias", metrics.Parsed, metrics.ParsedCode)
	assertFileMetric(t, "documents", metrics.Documents, 2)
	assertFileMetric(t, "metadataOnly", metrics.MetadataOnly, 4)
	assertFileMetric(t, "scriptNoExtractor", metrics.ScriptNoExtractor, 2)
	assertFileMetric(t, "staticAssets", metrics.StaticAssets, 2)
	assertFileMetric(t, "unsupportedLanguage", metrics.UnsupportedLanguage, 1)
	assertFileMetric(t, "unsupported legacy alias", metrics.Unsupported, metrics.UnsupportedLanguage)
	assertFileMetric(t, "unknown", metrics.Unknown, 1)
	assertFileMetric(t, "failed", metrics.Failed, 1)
	requireSample(t, metrics, FileBucketDocuments, "README.md")
	requireSample(t, metrics, FileBucketMetadataOnly, "package.json")
	requireSample(t, metrics, FileBucketScriptNoExtractor, "scripts/build.ps1")
	requireSample(t, metrics, FileBucketStaticAssets, "web/index.html")
	requireSample(t, metrics, FileBucketUnsupportedLanguage, "mainframe/program.cob")
	requireSample(t, metrics, FileBucketUnknown, "notes/blob.weird")
	requireSample(t, metrics, FileBucketFailed, "broken/input.ts")
}

func TestFileMetricsJSONContract(t *testing.T) {
	metrics := FileMetrics{
		Scanned:             2,
		Parsed:              1,
		ParsedCode:          1,
		Unsupported:         1,
		UnsupportedLanguage: 1,
		ClassificationSamples: []FileClassificationSample{
			{Bucket: FileBucketUnsupportedLanguage, Path: "mainframe/program.cob", Language: string(scanner.Cobol), Reason: "recognized language has no ScopeIR extractor"},
		},
	}
	raw, err := json.Marshal(metrics)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	for _, key := range []string{"scanned", "parsed", "parsedCode", "unsupported", "unsupportedLanguage", "classificationSamples"} {
		if _, ok := payload[key]; !ok {
			t.Fatalf("JSON missing %q: %s", key, raw)
		}
	}
}

func assertFileMetric(t *testing.T, name string, got int, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("%s = %d, want %d", name, got, want)
	}
}

func requireSample(t *testing.T, metrics FileMetrics, bucket FileClassificationBucket, filePath string) {
	t.Helper()
	for _, sample := range metrics.ClassificationSamples {
		if sample.Bucket == bucket && sample.Path == filePath {
			return
		}
	}
	t.Fatalf("missing sample bucket=%s path=%s in %#v", bucket, filePath, metrics.ClassificationSamples)
}
