package mcp

import (
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/tamnguyendinh/anvien/internal/filecontext"
	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/graphhealth"
	"github.com/tamnguyendinh/anvien/internal/repo"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
	"github.com/tamnguyendinh/anvien/internal/semantic"
)

var diffHunkPattern = regexp.MustCompile(`@@ -(\d+)(?:,(\d+))? \+(\d+)(?:,(\d+))? @@`)

type detectFileDiff struct {
	FilePath string
	Hunks    []detectHunk
	Deleted  bool
}

type detectHunk struct {
	StartLine int
	EndLine   int
}

func (s Server) detectChangesTool(args map[string]any) (map[string]any, error) {
	entry, err := s.resolveResourceRepo(stringArg(args, "repo", ""))
	if err != nil {
		return nil, err
	}
	diffOutput, err := gitDiffForDetectChanges(entry, args)
	if err != nil {
		return map[string]any{"error": "Git diff failed: " + err.Error()}, nil
	}
	fileDiffs := parseDetectDiffHunks(diffOutput)
	targetType := normalizedTargetType(args, "", detectChangesTargetTypeAllowed())
	dispatchMode := normalizedDispatchMode(args, targetType)
	if len(fileDiffs) == 0 {
		result := map[string]any{
			"summary": map[string]any{
				"changed_count":  0,
				"affected_count": 0,
				"risk_level":     "none",
				"message":        "No changes detected.",
			},
			"changed_symbols":    []map[string]any{},
			"changed_files":      []map[string]any{},
			"affected_files":     []map[string]any{},
			"affected_processes": []map[string]any{},
			"fileLayer": map[string]any{
				"changedFiles":     0,
				"affectedFiles":    0,
				"derivedEdgesNote": filecontext.DerivedFileEdgesNote,
			},
		}
		if targetType != "" {
			return detectChangesTargetPayload(result, nil, targetType, dispatchMode), nil
		}
		return result, nil
	}

	g, err := loadResourceGraphSnapshot(filepath.Join(storagePathForEntry(entry), "graph.json"))
	if err != nil {
		return nil, err
	}
	changedSymbols := detectChangedSymbols(g, fileDiffs)
	affectedProcesses := detectAffectedProcesses(g, changedSymbols)
	changedFiles := detectChangedFileRows(g, fileDiffs, changedSymbols)
	affectedFiles := detectAffectedFileRows(g, changedSymbols, affectedProcesses)
	semanticStatus := semantic.GraphSemanticStatus(g)
	semanticSummary := detectChangesSemanticSummary(changedSymbols, affectedProcesses)
	summary := map[string]any{
		"changed_count":  len(changedSymbols),
		"affected_count": len(affectedProcesses),
		"changed_files":  len(fileDiffs),
		"affected_files": len(affectedFiles),
		"risk_level":     detectRisk(len(affectedProcesses)),
	}
	if changedAppLayers := impactCountMapResult(semanticSummary, "changedAppLayers"); len(changedAppLayers) > 0 {
		summary["changed_app_layers"] = changedAppLayers
	}
	if changedAreas := impactCountMapResult(semanticSummary, "changedFunctionalAreas"); len(changedAreas) > 0 {
		summary["changed_functional_areas"] = changedAreas
	}
	if affectedAppLayers := impactCountMapResult(semanticSummary, "affectedAppLayers"); len(affectedAppLayers) > 0 {
		summary["affected_app_layers"] = affectedAppLayers
	}
	if affectedAreas := impactCountMapResult(semanticSummary, "affectedFunctionalAreas"); len(affectedAreas) > 0 {
		summary["affected_functional_areas"] = affectedAreas
	}
	if gapChanges, ok := semanticSummary["resolutionGapChanges"].(map[string]any); ok && detectChangesResolutionGapChangesHasEvidence(gapChanges) {
		summary["resolution_gap_changes"] = gapChanges
	}
	if healthImpact, ok := semanticSummary["resolutionHealthImpact"].(map[string]any); ok && detectChangesResolutionHealthImpactHasEvidence(healthImpact) {
		summary["resolution_health_impact"] = healthImpact
	}
	result := map[string]any{
		"summary":                 summary,
		"semanticStatus":          semanticStatus,
		"changedAppLayers":        semanticSummary["changedAppLayers"],
		"changedFunctionalAreas":  semanticSummary["changedFunctionalAreas"],
		"affectedAppLayers":       semanticSummary["affectedAppLayers"],
		"affectedFunctionalAreas": semanticSummary["affectedFunctionalAreas"],
		"resolutionGapChanges":    semanticSummary["resolutionGapChanges"],
		"resolutionHealthImpact":  semanticSummary["resolutionHealthImpact"],
		"changed_symbols":         changedSymbols,
		"changed_files":           changedFiles,
		"affected_files":          affectedFiles,
		"affected_processes":      affectedProcesses,
		"fileLayer": map[string]any{
			"changedFiles":     len(changedFiles),
			"affectedFiles":    len(affectedFiles),
			"changedFileRisk":  detectFileRisk(changedFiles),
			"derivedEdgesNote": filecontext.DerivedFileEdgesNote,
		},
	}
	if warning := querySemanticWarning(semanticStatus); warning != "" {
		result["semanticWarning"] = warning
	}
	if targetType != "" {
		return detectChangesTargetPayload(result, changedFiles, targetType, dispatchMode), nil
	}
	return result, nil
}

func detectChangesTargetPayload(result map[string]any, changedFiles []map[string]any, targetType string, dispatchMode string) map[string]any {
	payload := map[string]any{
		"summary":      result["summary"],
		"targetType":   targetType,
		"dispatchMode": dispatchMode,
	}
	for _, key := range []string{
		"semanticStatus",
		"semanticWarning",
		"changedAppLayers",
		"changedFunctionalAreas",
		"affectedAppLayers",
		"affectedFunctionalAreas",
		"resolutionGapChanges",
		"resolutionHealthImpact",
	} {
		if value, ok := result[key]; ok {
			payload[key] = value
		}
	}
	switch targetType {
	case targetTypeFiles:
		if changedFiles == nil {
			changedFiles = []map[string]any{}
		}
		payload["changed_files"] = changedFiles
		payload["total"] = len(changedFiles)
	case targetTypeSymbols:
		symbols, _ := result["changed_symbols"].([]map[string]any)
		if symbols == nil {
			symbols = []map[string]any{}
		}
		payload["changed_symbols"] = symbols
		payload["total"] = len(symbols)
	case targetTypeFlows:
		flows, _ := result["affected_processes"].([]map[string]any)
		if flows == nil {
			flows = []map[string]any{}
		}
		payload["affected_processes"] = flows
		payload["flows"] = flows
		payload["total"] = len(flows)
	}
	return payload
}

func detectChangedFileRows(g *graph.Graph, fileDiffs []detectFileDiff, changedSymbols []map[string]any) []map[string]any {
	symbolsByFile := detectSymbolsByFile(changedSymbols)
	rows := make([]map[string]any, 0, len(fileDiffs))
	for _, fileDiff := range fileDiffs {
		row := map[string]any{
			"path":    fileDiff.FilePath,
			"deleted": fileDiff.Deleted,
			"hunks":   fileDiff.Hunks,
		}
		if summary, ok := mcpFileSummaryForPath(g, fileDiff.FilePath); ok {
			row["summary"] = summary
			row["relationshipHints"] = mcpFileRelationshipHints(summary)
			row["linkedFlows"] = summary.LinkedFlowCount
			row["linkedTests"] = summary.LinkedTestCount
			row["fileRisk"] = summary.Risk
		}
		symbols := symbolsByFile[normalizeContextPath(fileDiff.FilePath)]
		if symbols == nil {
			symbols = []map[string]any{}
		}
		row["changedSymbols"] = symbols
		row["changedSymbolCount"] = len(symbols)
		row["unresolvedDelta"] = detectChangedSymbolsResolutionGapCount(symbols)
		rows = append(rows, row)
	}
	sort.Slice(rows, func(i, j int) bool {
		return fmt.Sprint(rows[i]["path"]) < fmt.Sprint(rows[j]["path"])
	})
	return rows
}

func detectAffectedFileRows(g *graph.Graph, changedSymbols []map[string]any, affectedProcesses []map[string]any) []map[string]any {
	counts := map[string]int{}
	for _, symbol := range changedSymbols {
		incrementQueryCount(counts, normalizeContextPath(impactStringValue(symbol["filePath"])), 1)
	}
	for _, process := range affectedProcesses {
		steps, _ := process["changed_steps"].([]map[string]any)
		for _, step := range steps {
			incrementQueryCount(counts, normalizeContextPath(impactStringValue(step["filePath"])), 1)
		}
	}
	paths := make([]string, 0, len(counts))
	for path := range counts {
		if strings.TrimSpace(path) != "" {
			paths = append(paths, path)
		}
	}
	sort.Strings(paths)
	rows := make([]map[string]any, 0, len(paths))
	for _, path := range paths {
		row := map[string]any{"path": path, "affectedSymbols": counts[path]}
		if summary, ok := mcpFileSummaryForPath(g, path); ok {
			row["summary"] = summary
			row["relationshipHints"] = mcpFileRelationshipHints(summary)
			row["fileRisk"] = summary.Risk
		}
		rows = append(rows, row)
	}
	return rows
}

func detectSymbolsByFile(symbols []map[string]any) map[string][]map[string]any {
	out := map[string][]map[string]any{}
	for _, symbol := range symbols {
		path := normalizeContextPath(impactStringValue(symbol["filePath"]))
		if path == "" {
			continue
		}
		out[path] = append(out[path], symbol)
	}
	return out
}

func detectChangedSymbolsResolutionGapCount(symbols []map[string]any) int {
	total := 0
	for _, symbol := range symbols {
		total += impactIntValue(symbol["resolutionGapCount"])
		if entity, _ := symbol["resolutionGapEntity"].(bool); entity {
			total += maxDetectCount(impactIntValue(symbol["count"]), 1)
		}
	}
	return total
}

func detectFileRisk(files []map[string]any) string {
	riskOrder := map[string]int{"low": 1, "medium": 2, "high": 3, "critical": 4}
	selected := "low"
	for _, file := range files {
		risk := strings.ToLower(strings.TrimSpace(impactStringValue(file["fileRisk"])))
		if riskOrder[risk] > riskOrder[selected] {
			selected = risk
		}
	}
	return selected
}

func gitDiffForDetectChanges(entry repo.RegistryEntry, args map[string]any) (string, error) {
	scope := firstNonEmptyString(stringArg(args, "scope", ""), "unstaged")
	diffArgs := []string{"-C", entry.Path, "diff", "-U0"}
	switch scope {
	case "staged":
		diffArgs = []string{"-C", entry.Path, "diff", "--staged", "-U0"}
	case "all":
		diffArgs = []string{"-C", entry.Path, "diff", "HEAD", "-U0"}
	case "compare":
		baseRef := strings.TrimSpace(stringArg(args, "base_ref", ""))
		if baseRef == "" {
			return "", errors.New(`base_ref is required for "compare" scope`)
		}
		diffArgs = []string{"-C", entry.Path, "diff", baseRef, "-U0"}
	case "unstaged":
	default:
		scope = "unstaged"
	}
	output, err := exec.Command("git", diffArgs...).Output()
	return string(output), err
}

func parseDetectDiffHunks(diffOutput string) []detectFileDiff {
	var files []detectFileDiff
	currentIndex := -1
	oldPath := ""
	for _, line := range strings.Split(diffOutput, "\n") {
		switch {
		case strings.HasPrefix(line, "--- "):
			oldPath = strings.TrimSpace(strings.TrimPrefix(line, "--- "))
			if oldPath != "/dev/null" {
				oldPath = strings.TrimPrefix(oldPath, "a/")
			}
		case strings.HasPrefix(line, "+++ "):
			newPath := strings.TrimSpace(strings.TrimPrefix(line, "+++ "))
			deleted := newPath == "/dev/null"
			filePath := newPath
			if deleted {
				filePath = oldPath
			}
			if filePath == "" || filePath == "/dev/null" {
				currentIndex = -1
				continue
			}
			filePath = strings.TrimPrefix(filePath, "b/")
			filePath = strings.TrimPrefix(filePath, "a/")
			files = append(files, detectFileDiff{FilePath: normalizeContextPath(filePath), Deleted: deleted})
			currentIndex = len(files) - 1
		case strings.HasPrefix(line, "@@ ") && currentIndex >= 0:
			hunk, ok := parseDetectHunk(line, files[currentIndex].Deleted)
			if ok {
				files[currentIndex].Hunks = append(files[currentIndex].Hunks, hunk)
			}
		}
	}
	out := files[:0]
	for _, file := range files {
		if file.FilePath != "" && len(file.Hunks) > 0 {
			out = append(out, file)
		}
	}
	return out
}

func parseDetectHunk(line string, deleted bool) (detectHunk, bool) {
	matches := diffHunkPattern.FindStringSubmatch(line)
	if len(matches) == 0 {
		return detectHunk{}, false
	}
	startIndex := 3
	countIndex := 4
	if deleted {
		startIndex = 1
		countIndex = 2
	}
	start, err := strconv.Atoi(matches[startIndex])
	if err != nil {
		return detectHunk{}, false
	}
	count := 1
	if len(matches) > countIndex && matches[countIndex] != "" {
		parsed, parseErr := strconv.Atoi(matches[countIndex])
		if parseErr == nil {
			count = parsed
		}
	}
	if count == 0 {
		return detectHunk{}, false
	}
	end := start
	if count > 0 {
		end = start + count - 1
	}
	return detectHunk{StartLine: start, EndLine: end}, true
}

func detectChangedSymbols(g *graph.Graph, fileDiffs []detectFileDiff) []map[string]any {
	seen := make(map[string]bool)
	out := make([]map[string]any, 0)
	for _, fileDiff := range fileDiffs {
		for _, node := range g.Nodes {
			filePath := normalizeContextPath(resourceNodeString(node, "filePath"))
			if filePath == "" || !(filePath == fileDiff.FilePath || strings.HasSuffix(filePath, "/"+fileDiff.FilePath)) {
				continue
			}
			startLine := resourceNodeInt(node, "startLine")
			endLine := resourceNodeInt(node, "endLine")
			if startLine == 0 || endLine == 0 || !detectOverlapsAnyHunk(startLine, endLine, fileDiff.Hunks) {
				continue
			}
			if seen[node.ID] {
				continue
			}
			seen[node.ID] = true
			changeType := "touched"
			if fileDiff.Deleted {
				changeType = "deleted"
			}
			out = append(out, map[string]any{
				"id":          node.ID,
				"name":        firstResourceNodeString(node, "name", "label", "heuristicLabel"),
				"type":        string(node.Label),
				"filePath":    resourceNodeString(node, "filePath"),
				"change_type": changeType,
			})
			addContextNodeSemanticFields(out[len(out)-1], node)
			addContextResolutionGapEntityFields(out[len(out)-1], node)
		}
	}
	return out
}

func detectOverlapsAnyHunk(startLine int, endLine int, hunks []detectHunk) bool {
	for _, hunk := range hunks {
		if startLine <= hunk.EndLine && endLine >= hunk.StartLine {
			return true
		}
	}
	return false
}

func detectAffectedProcesses(g *graph.Graph, changedSymbols []map[string]any) []map[string]any {
	if len(changedSymbols) == 0 {
		return []map[string]any{}
	}
	changedSet := make(map[string]string, len(changedSymbols))
	changedByID := make(map[string]map[string]any, len(changedSymbols))
	for _, symbol := range changedSymbols {
		id := fmt.Sprint(symbol["id"])
		changedSet[id] = fmt.Sprint(symbol["name"])
		changedByID[id] = symbol
	}
	nodeByID := resourceGraphNodesByID(g)
	processes := make(map[string]map[string]any)
	for _, relationship := range g.Relationships {
		symbolName, ok := changedSet[relationship.SourceID]
		if relationship.Type != graph.RelStepInProcess || !ok {
			continue
		}
		process, exists := nodeByID[relationship.TargetID]
		if !exists {
			continue
		}
		current := processes[process.ID]
		if current == nil {
			current = map[string]any{
				"id":            process.ID,
				"name":          firstResourceNodeString(process, "heuristicLabel", "label", "name"),
				"process_type":  resourceNodeString(process, "processType"),
				"step_count":    resourceNodeInt(process, "stepCount"),
				"changed_steps": []map[string]any{},
			}
			addContextNodeSemanticFields(current, process)
			processes[process.ID] = current
		}
		step := 0
		if relationship.Step != nil {
			step = *relationship.Step
		}
		changedSymbol := changedByID[relationship.SourceID]
		stepRow := map[string]any{"symbol": symbolName, "step": step}
		if changedSymbol != nil {
			stepRow["id"] = changedSymbol["id"]
			detectCopySemanticRowFields(stepRow, changedSymbol)
			if gapCount := impactIntValue(changedSymbol["resolutionGapCount"]); gapCount > 0 {
				stepRow["resolutionGapCount"] = gapCount
			}
			if buckets := impactCountMapValue(changedSymbol["resolutionHealthBuckets"]); len(buckets) > 0 {
				stepRow["resolutionHealthBuckets"] = buckets
			}
			detectIncrementCountField(current, "changedStepAppLayers", detectRowAppLayer(changedSymbol), 1)
			detectIncrementCountField(current, "changedStepFunctionalAreas", detectRowFunctionalArea(changedSymbol), 1)
		}
		current["changed_steps"] = append(current["changed_steps"].([]map[string]any), stepRow)
	}
	out := make([]map[string]any, 0, len(processes))
	for _, process := range processes {
		if steps, ok := process["changed_steps"].([]map[string]any); ok {
			if healthImpact := detectChangesResolutionHealthImpact(steps); detectChangesResolutionHealthImpactHasEvidence(healthImpact) {
				process["resolutionHealthImpact"] = healthImpact
			}
		}
		out = append(out, process)
	}
	sort.Slice(out, func(i, j int) bool {
		left, right := out[i], out[j]
		if fmt.Sprint(left["name"]) != fmt.Sprint(right["name"]) {
			return fmt.Sprint(left["name"]) < fmt.Sprint(right["name"])
		}
		return fmt.Sprint(left["id"]) < fmt.Sprint(right["id"])
	})
	return out
}

func detectChangesSemanticSummary(changedSymbols []map[string]any, affectedProcesses []map[string]any) map[string]any {
	return map[string]any{
		"changedAppLayers":        detectCountRowsByAppLayer(changedSymbols),
		"changedFunctionalAreas":  detectCountRowsByFunctionalArea(changedSymbols),
		"affectedAppLayers":       detectCountRowsByAppLayer(affectedProcesses),
		"affectedFunctionalAreas": detectCountRowsByFunctionalArea(affectedProcesses),
		"resolutionGapChanges":    detectChangesResolutionGapChanges(changedSymbols),
		"resolutionHealthImpact":  detectChangesResolutionHealthImpact(changedSymbols),
	}
}

func detectCountRowsByAppLayer(rows []map[string]any) map[string]int {
	counts := map[string]int{}
	for _, row := range rows {
		incrementQueryCount(counts, detectRowAppLayer(row), 1)
	}
	return cloneQueryCountMap(counts)
}

func detectCountRowsByFunctionalArea(rows []map[string]any) map[string]int {
	counts := map[string]int{}
	for _, row := range rows {
		incrementQueryCount(counts, detectRowFunctionalArea(row), 1)
	}
	return cloneQueryCountMap(counts)
}

func detectRowAppLayer(row map[string]any) string {
	return firstNonEmptyString(impactStringValue(row["appLayer"]), impactStringValue(row["sourceAppLayer"]))
}

func detectRowFunctionalArea(row map[string]any) string {
	return firstNonEmptyString(impactStringValue(row["functionalArea"]), impactStringValue(row["sourceFunctionalArea"]))
}

func detectCopySemanticRowFields(dst map[string]any, src map[string]any) {
	for _, key := range []string{
		"type",
		"filePath",
		"appLayer",
		"appLayerSource",
		"functionalArea",
		"functionalAreaSource",
		"topologyStatus",
		"resolutionConfidence",
		"sourceAppLayer",
		"sourceFunctionalArea",
		"factFamily",
		"targetRole",
		"targetText",
		"gapKind",
		"classification",
		"actionability",
		"sourceSiteStatus",
		"proofKind",
	} {
		if value, ok := src[key]; ok && value != nil && fmt.Sprint(value) != "" {
			dst[key] = value
		}
	}
}

func detectIncrementCountField(payload map[string]any, key string, value string, count int) {
	if value == "" || count <= 0 {
		return
	}
	counts, _ := payload[key].(map[string]int)
	if counts == nil {
		counts = map[string]int{}
		payload[key] = counts
	}
	incrementQueryCount(counts, value, count)
}

func detectChangesResolutionGapChanges(changedSymbols []map[string]any) map[string]any {
	byGapKind := map[string]int{}
	byFactFamily := map[string]int{}
	byTargetRole := map[string]int{}
	byClassification := map[string]int{}
	byActionability := map[string]int{}
	byAppLayer := map[string]int{}
	byFunctionalArea := map[string]int{}
	topTargets := map[string]int{}
	gapEntities := 0
	gapOccurrenceCount := 0
	sourceNodesWithGaps := 0
	totalGapCount := 0

	for _, row := range changedSymbols {
		if impactIntValue(row["resolutionGapCount"]) > 0 {
			sourceNodesWithGaps++
			totalGapCount += impactIntValue(row["resolutionGapCount"])
		}
		isGapEntity := fmt.Sprint(row["type"]) == string(scopeir.NodeResolutionGap)
		if entity, _ := row["resolutionGapEntity"].(bool); entity {
			isGapEntity = true
		}
		if !isGapEntity {
			continue
		}
		gapEntities++
		occurrenceCount := maxDetectCount(impactIntValue(row["count"]), 1)
		gapOccurrenceCount += occurrenceCount
		incrementQueryCount(byGapKind, impactStringValue(row["gapKind"]), 1)
		incrementQueryCount(byFactFamily, impactStringValue(row["factFamily"]), 1)
		incrementQueryCount(byTargetRole, impactStringValue(row["targetRole"]), 1)
		incrementQueryCount(byClassification, impactStringValue(row["classification"]), 1)
		incrementQueryCount(byActionability, impactStringValue(row["actionability"]), 1)
		incrementQueryCount(byAppLayer, detectRowAppLayer(row), 1)
		incrementQueryCount(byFunctionalArea, detectRowFunctionalArea(row), 1)
		incrementQueryCount(topTargets, impactStringValue(row["targetText"]), occurrenceCount)
	}

	return map[string]any{
		"changedGapEntities":         gapEntities,
		"changedGapOccurrenceCount":  gapOccurrenceCount,
		"changedSourceNodesWithGaps": sourceNodesWithGaps,
		"totalResolutionGapCount":    totalGapCount,
		"gapKinds":                   cloneQueryCountMap(byGapKind),
		"factFamilies":               cloneQueryCountMap(byFactFamily),
		"targetRoles":                cloneQueryCountMap(byTargetRole),
		"classifications":            cloneQueryCountMap(byClassification),
		"actionability":              cloneQueryCountMap(byActionability),
		"appLayers":                  cloneQueryCountMap(byAppLayer),
		"functionalAreas":            cloneQueryCountMap(byFunctionalArea),
		"topTargets":                 topQueryCountMap(topTargets, 10),
	}
}

func detectChangesResolutionHealthImpact(rows []map[string]any) map[string]any {
	resolutionConfidence := map[string]int{}
	resolutionBuckets := map[string]int{}
	totalGapCount := 0
	nodesWithGaps := 0
	degradedNodes := 0
	riskNodes := make([]map[string]any, 0)

	for _, row := range rows {
		confidence := impactStringValue(row["resolutionConfidence"])
		incrementQueryCount(resolutionConfidence, confidence, 1)
		if confidence == graphhealth.ResolutionConfidenceDegraded {
			degradedNodes++
		}
		gapCount := impactIntValue(row["resolutionGapCount"])
		if gapCount > 0 {
			nodesWithGaps++
			totalGapCount += gapCount
		}
		buckets := impactCountMapValue(row["resolutionHealthBuckets"])
		for bucket, count := range buckets {
			incrementQueryCount(resolutionBuckets, bucket, count)
		}
		if gapCount == 0 && confidence != graphhealth.ResolutionConfidenceDegraded && len(buckets) == 0 {
			continue
		}
		riskNode := map[string]any{
			"id":   row["id"],
			"name": row["name"],
			"type": row["type"],
		}
		detectCopySemanticRowFields(riskNode, row)
		if gapCount > 0 {
			riskNode["resolutionGapCount"] = gapCount
		}
		if len(buckets) > 0 {
			riskNode["resolutionHealthBuckets"] = buckets
		}
		riskNodes = append(riskNodes, riskNode)
	}

	return map[string]any{
		"nodesWithGaps":              nodesWithGaps,
		"degradedNodes":              degradedNodes,
		"totalResolutionGapCount":    totalGapCount,
		"resolutionConfidenceCounts": cloneQueryCountMap(resolutionConfidence),
		"resolutionHealthBuckets":    cloneQueryCountMap(resolutionBuckets),
		"nodes":                      riskNodes,
	}
}

func detectChangesResolutionGapChangesHasEvidence(changes map[string]any) bool {
	return impactIntValue(changes["changedGapEntities"]) > 0 ||
		impactIntValue(changes["changedGapOccurrenceCount"]) > 0 ||
		impactIntValue(changes["changedSourceNodesWithGaps"]) > 0 ||
		impactIntValue(changes["totalResolutionGapCount"]) > 0 ||
		len(impactCountMapValue(changes["gapKinds"])) > 0 ||
		len(impactCountMapValue(changes["topTargets"])) > 0
}

func detectChangesResolutionHealthImpactHasEvidence(impact map[string]any) bool {
	return impactIntValue(impact["nodesWithGaps"]) > 0 ||
		impactIntValue(impact["degradedNodes"]) > 0 ||
		impactIntValue(impact["totalResolutionGapCount"]) > 0 ||
		len(impactCountMapValue(impact["resolutionHealthBuckets"])) > 0
}

func maxDetectCount(left int, right int) int {
	if left > right {
		return left
	}
	return right
}

func detectRisk(processCount int) string {
	switch {
	case processCount == 0:
		return "low"
	case processCount <= 5:
		return "medium"
	case processCount <= 15:
		return "high"
	default:
		return "critical"
	}
}
