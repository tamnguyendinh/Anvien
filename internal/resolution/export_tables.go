package resolution

import (
	"encoding/json"
	"path/filepath"
	"sort"
	"strings"

	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

// exportTableEntry is a syntax-derived export entry.  Fact is copied from the
// accepted ScopeIR export fact and is deliberately kept separate from any
// terminal definition or resolution outcome.  TargetFiles contains only the
// already-resolved module/file candidates for a source-bearing export.
type exportTableEntry struct {
	Fact        scopeir.ExportFact
	TargetFiles []string
}

// exportStarAdjacency records an export * edge without expanding it.  P5-C
// owns traversal, precedence, ambiguity, and cycle outcomes; P5-B only keeps
// the source fact and the existing module/file result together.
type exportStarAdjacency struct {
	Fact        scopeir.ExportFact
	TargetFiles []string
}

// exportTable is the syntax-derived export surface for one repository file.
// Explicit is keyed by the source-written exported name. StarAdjacency is
// intentionally a separate list because export * has no exported name and
// must never synthesize a default entry.
type exportTable struct {
	FilePath      string
	Explicit      map[string][]exportTableEntry
	StarAdjacency []exportStarAdjacency
}

type exportTables map[string]exportTable

// buildExportTables builds deterministic syntax-derived export tables from
// accepted ScopeIR facts and the module/file results already produced by the
// workspace import pass. It does not inspect physical definitions and does
// not resolve a terminal symbol.
func buildExportTables(files []scopeir.ScopeIR, imports []resolvedImport) exportTables {
	resolvedByFact := indexResolvedExportImports(imports)
	tables := make(exportTables, len(files))

	for _, input := range files {
		filePath := cleanExportTablePath(input.FilePath)
		if filePath == "" {
			continue
		}
		table, ok := tables[filePath]
		if !ok {
			table = exportTable{
				FilePath: filePath,
				Explicit: make(map[string][]exportTableEntry),
			}
		}

		for _, sourceFact := range input.Exports {
			fact := cloneExportTableFact(sourceFact)
			if fact.FilePath == "" {
				fact.FilePath = filePath
			} else {
				fact.FilePath = cleanExportTablePath(fact.FilePath)
			}
			targetFiles := resolvedByFact[exportImportKeyForFact(fact)]
			targetFiles = cloneSortedExportTablePaths(targetFiles)

			if fact.Kind == scopeir.ExportStar {
				table.StarAdjacency = append(table.StarAdjacency, exportStarAdjacency{
					Fact:        fact,
					TargetFiles: targetFiles,
				})
				continue
			}

			name := fact.ExportedName
			table.Explicit[name] = append(table.Explicit[name], exportTableEntry{
				Fact:        fact,
				TargetFiles: targetFiles,
			})
		}

		sortExportTable(&table)
		tables[filePath] = table
	}

	return tables
}

// buildExportTablesForWorkspace is the integration seam used by
// buildWorkspace. Keeping the call here makes the table construction consume
// the normalized workspace files and the existing resolvedImport results
// without changing resolveImports or path lookup.
func (w *workspace) buildExportTables() {
	w.exportTables = buildExportTables(w.files, w.imports)
}

type exportImportKey struct {
	filePath     string
	kind         scopeir.ImportKind
	localName    string
	importedName string
	targetRaw    string
}

func indexResolvedExportImports(imports []resolvedImport) map[exportImportKey][]string {
	indexed := make(map[exportImportKey][]string)
	for _, item := range imports {
		if item.Fact.TargetRaw == nil {
			continue
		}
		if item.Fact.Kind != scopeir.ImportReexport && item.Fact.Kind != scopeir.ImportWildcard {
			continue
		}
		key := exportImportKey{
			filePath:     cleanExportTablePath(item.Fact.FilePath),
			kind:         item.Fact.Kind,
			localName:    item.Fact.LocalName,
			importedName: item.Fact.ImportedName,
			targetRaw:    *item.Fact.TargetRaw,
		}
		for _, targetFile := range item.TargetFiles {
			indexed[key] = appendUniqueExportTablePath(indexed[key], targetFile)
		}
	}
	for key, targetFiles := range indexed {
		indexed[key] = cloneSortedExportTablePaths(targetFiles)
	}
	return indexed
}

func exportImportKeyForFact(fact scopeir.ExportFact) exportImportKey {
	kind := scopeir.ImportReexport
	localName := fact.ExportedName
	importedName := fact.TargetExportedName
	if fact.Kind == scopeir.ExportStar || fact.Kind == scopeir.ExportNamespace {
		kind = scopeir.ImportWildcard
		localName = ""
		importedName = ""
	}
	targetRaw := ""
	if fact.TargetRaw != nil {
		targetRaw = *fact.TargetRaw
	}
	return exportImportKey{
		filePath:     cleanExportTablePath(fact.FilePath),
		kind:         kind,
		localName:    localName,
		importedName: importedName,
		targetRaw:    targetRaw,
	}
}

func sortExportTable(table *exportTable) {
	for name, entries := range table.Explicit {
		sort.Slice(entries, func(i, j int) bool {
			return exportTableFactKey(entries[i].Fact) < exportTableFactKey(entries[j].Fact)
		})
		for index := range entries {
			entries[index].TargetFiles = cloneSortedExportTablePaths(entries[index].TargetFiles)
		}
		table.Explicit[name] = entries
	}
	sort.Slice(table.StarAdjacency, func(i, j int) bool {
		return exportTableFactKey(table.StarAdjacency[i].Fact) < exportTableFactKey(table.StarAdjacency[j].Fact)
	})
	for index := range table.StarAdjacency {
		table.StarAdjacency[index].TargetFiles = cloneSortedExportTablePaths(table.StarAdjacency[index].TargetFiles)
	}
}

func cloneExportTableFact(fact scopeir.ExportFact) scopeir.ExportFact {
	cloned := fact
	cloned.Meanings = append([]scopeir.ExportMeaning(nil), fact.Meanings...)
	if fact.SelectionRange != nil {
		rangeCopy := *fact.SelectionRange
		cloned.SelectionRange = &rangeCopy
	}
	if fact.TargetRaw != nil {
		targetCopy := *fact.TargetRaw
		cloned.TargetRaw = &targetCopy
	}
	return cloned
}

func cloneSortedExportTablePaths(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	paths := make([]string, 0, len(values))
	for _, value := range values {
		value = cleanExportTablePath(value)
		if value == "" || (len(paths) > 0 && paths[len(paths)-1] == value) {
			continue
		}
		paths = append(paths, value)
	}
	sort.Strings(paths)
	writeIndex := 0
	for _, value := range paths {
		if writeIndex > 0 && paths[writeIndex-1] == value {
			continue
		}
		paths[writeIndex] = value
		writeIndex++
	}
	return paths[:writeIndex]
}

func appendUniqueExportTablePath(values []string, value string) []string {
	value = cleanExportTablePath(value)
	if value == "" {
		return values
	}
	for _, existing := range values {
		if existing == value {
			return values
		}
	}
	return append(values, value)
}

func cleanExportTablePath(value string) string {
	if value == "" {
		return ""
	}
	return filepath.ToSlash(filepath.Clean(strings.ReplaceAll(value, "\\", "/")))
}

func exportTableFactKey(fact scopeir.ExportFact) string {
	encoded, err := json.Marshal(fact)
	if err != nil {
		// ExportFact contains only JSON-safe scalar, pointer, and slice fields.
		// Keep a deterministic fallback if that contract ever changes.
		return strings.Join([]string{
			fact.FilePath,
			string(fact.Kind),
			fact.ExportedName,
			fact.LocalName,
			fact.TargetExportedName,
		}, "\x00")
	}
	return string(encoded)
}
