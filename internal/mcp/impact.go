package mcp

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/tamnguyendinh/avmatrix-go/internal/graph"
	"github.com/tamnguyendinh/avmatrix-go/internal/graphhealth"
	"github.com/tamnguyendinh/avmatrix-go/internal/scopeir"
	"github.com/tamnguyendinh/avmatrix-go/internal/semantic"
)

var (
	impactAllowedDirections = map[string]bool{
		"upstream":   true,
		"downstream": true,
	}
	impactAllowedRelationTypes = map[string]bool{
		string(graph.RelCalls):            true,
		string(graph.RelImports):          true,
		string(graph.RelUses):             true,
		string(graph.RelInherits):         true,
		string(graph.RelExtends):          true,
		string(graph.RelImplements):       true,
		string(graph.RelHasMethod):        true,
		string(graph.RelHasProperty):      true,
		string(graph.RelMethodOverrides):  true,
		"OVERRIDES":                       true,
		string(graph.RelMethodImplements): true,
		string(graph.RelAccesses):         true,
		string(graph.RelHandlesRoute):     true,
		string(graph.RelFetches):          true,
		string(graph.RelHandlesTool):      true,
		string(graph.RelEntryPointOf):     true,
		string(graph.RelWraps):            true,
	}
	impactDefaultRelationTypes = []string{
		string(graph.RelCalls),
		string(graph.RelImports),
		string(graph.RelUses),
		string(graph.RelInherits),
		string(graph.RelExtends),
		string(graph.RelImplements),
		string(graph.RelHasProperty),
		string(graph.RelMethodOverrides),
		"OVERRIDES",
		string(graph.RelMethodImplements),
		string(graph.RelAccesses),
	}
)

type impactOptions struct {
	Target        string
	TargetUID     string
	Direction     string
	FilePath      string
	Kind          string
	MaxDepth      int
	RelationTypes []string
	IncludeTests  bool
	MinConfidence float64
}

func (s Server) impactTool(args map[string]any) (map[string]any, error) {
	payload, _, err := s.impactToolInternal(args, false)
	return payload, err
}

type impactToolProfile struct {
	RepoResolve       time.Duration
	TargetLookup      time.Duration
	IndexBuild        time.Duration
	Traversal         time.Duration
	AffectedSummaries time.Duration
	Formatting        time.Duration
}

func (s Server) impactToolProfiled(args map[string]any) (map[string]any, impactToolProfile, error) {
	return s.impactToolInternal(args, true)
}

func (s Server) impactToolInternal(args map[string]any, collectProfile bool) (map[string]any, impactToolProfile, error) {
	var profile impactToolProfile
	last := time.Now()
	mark := func(target *time.Duration) {
		if !collectProfile {
			return
		}
		now := time.Now()
		*target += impactProfileElapsed(last, now)
		last = now
	}

	options, validation := parseImpactArgs(args)
	if validation != nil {
		return impactValidationResult(args, validation), profile, nil
	}

	g, err := s.graphForResource(stringArg(args, "repo", ""))
	if err != nil {
		return nil, profile, err
	}
	mark(&profile.RepoResolve)

	candidates := contextCandidates(g, options.Target, options.TargetUID, options.FilePath, options.Kind)
	targetLabel := firstNonEmptyString(options.Target, options.TargetUID)
	mark(&profile.TargetLookup)
	if len(candidates) == 0 {
		return map[string]any{
			"error":         fmt.Sprintf("Target '%s' not found", targetLabel),
			"target":        map[string]any{"name": targetLabel},
			"direction":     options.Direction,
			"impactedCount": 0,
			"risk":          "UNKNOWN",
		}, profile, nil
	}
	if len(candidates) > 1 && candidates[0].Score >= 0.9 && candidates[0].Score-candidates[1].Score > 0.09 {
		candidates = candidates[:1]
	}
	if len(candidates) > 1 {
		return map[string]any{
			"status":        "ambiguous",
			"message":       fmt.Sprintf("Found %d symbols matching '%s'. Use target_uid, file_path, or kind to disambiguate.", len(candidates), targetLabel),
			"target":        map[string]any{"name": targetLabel},
			"direction":     options.Direction,
			"impactedCount": 0,
			"risk":          "UNKNOWN",
			"candidates":    contextCandidatePayloads(candidates),
		}, profile, nil
	}

	payload, runProfile := runImpactBFSProfiled(g, candidates[0].Node, options, collectProfile)
	profile.IndexBuild += runProfile.IndexBuild
	profile.Traversal += runProfile.Traversal
	profile.AffectedSummaries += runProfile.AffectedSummaries
	profile.Formatting += runProfile.Formatting
	return payload, profile, nil
}

func parseImpactArgs(args map[string]any) (impactOptions, map[string]any) {
	options := impactOptions{
		Target:        strings.TrimSpace(stringArg(args, "target", "")),
		TargetUID:     strings.TrimSpace(stringArg(args, "target_uid", "")),
		Direction:     strings.TrimSpace(stringArg(args, "direction", "")),
		FilePath:      stringArg(args, "file_path", ""),
		Kind:          stringArg(args, "kind", ""),
		MaxDepth:      intArg(args, "maxDepth", 3, 1, 20),
		RelationTypes: impactDefaultRelationTypes,
		IncludeTests:  boolArg(args, "includeTests", false),
		MinConfidence: floatArg(args, "minConfidence", 0, 0, 1),
	}
	if options.Target == "" && options.TargetUID == "" {
		return options, map[string]any{"error": "Impact validation failed: provide either target or target_uid.", "field": "target"}
	}
	if !impactAllowedDirections[options.Direction] {
		return options, map[string]any{
			"error":         "Impact validation failed: direction must be one of upstream, downstream.",
			"field":         "direction",
			"allowedValues": []string{"upstream", "downstream"},
		}
	}
	if rawTypes, ok := args["relationTypes"].([]any); ok {
		types := make([]string, 0, len(rawTypes))
		for _, rawType := range rawTypes {
			relType := strings.TrimSpace(fmt.Sprint(rawType))
			if !impactAllowedRelationTypes[relType] {
				return options, map[string]any{
					"error":         "Impact validation failed: invalid relationTypes: " + relType + ".",
					"field":         "relationTypes",
					"allowedValues": sortedImpactAllowedRelationTypes(),
				}
			}
			types = append(types, relType)
			if relType == "OVERRIDES" {
				types = append(types, string(graph.RelMethodOverrides))
			}
		}
		if len(types) == 0 {
			return options, map[string]any{
				"error":         "Impact validation failed: relationTypes must include at least one valid relation when provided.",
				"field":         "relationTypes",
				"allowedValues": sortedImpactAllowedRelationTypes(),
			}
		}
		options.RelationTypes = dedupeStrings(types)
	}
	return options, nil
}

func impactValidationResult(args map[string]any, issue map[string]any) map[string]any {
	target := firstNonEmptyString(stringArg(args, "target", ""), stringArg(args, "target_uid", ""))
	result := map[string]any{
		"error":         issue["error"],
		"field":         issue["field"],
		"target":        map[string]any{"name": target},
		"direction":     firstNonEmptyString(stringArg(args, "direction", ""), "upstream"),
		"impactedCount": 0,
		"risk":          "UNKNOWN",
		"suggestion":    "Fix the invalid impact parameters and retry the request.",
	}
	if allowed, ok := issue["allowedValues"]; ok {
		result["allowedValues"] = allowed
	}
	return result
}

func runImpactBFS(g *graph.Graph, target graph.Node, options impactOptions) map[string]any {
	result, _ := runImpactBFSProfiled(g, target, options, false)
	return result
}

func runImpactBFSProfiled(g *graph.Graph, target graph.Node, options impactOptions, collectProfile bool) (map[string]any, impactToolProfile) {
	var profile impactToolProfile
	last := time.Now()
	nodeByID := resourceGraphNodesByID(g)
	relationSet := stringSet(options.RelationTypes)
	if collectProfile {
		now := time.Now()
		profile.IndexBuild = impactProfileElapsed(last, now)
		last = now
	}
	visited := map[string]bool{target.ID: true}
	frontier := []string{target.ID}
	impacted := make([]map[string]any, 0)
	traversalComplete := true

	if target.Label == scopeir.NodeClass || target.Label == scopeir.NodeInterface {
		for _, seedID := range impactClassLikeSeedIDs(g, target.ID) {
			if !visited[seedID] {
				visited[seedID] = true
				frontier = append(frontier, seedID)
			}
		}
	}

	for depth := 1; depth <= options.MaxDepth && len(frontier) > 0; depth++ {
		frontierSet := stringSet(frontier)
		nextFrontier := make([]string, 0)
		for _, relationship := range g.Relationships {
			relType := string(relationship.Type)
			if !relationSet[relType] {
				continue
			}
			if options.MinConfidence > 0 && relationship.Confidence < options.MinConfidence {
				continue
			}

			var relatedID string
			if options.Direction == "upstream" {
				if !frontierSet[relationship.TargetID] {
					continue
				}
				relatedID = relationship.SourceID
			} else {
				if !frontierSet[relationship.SourceID] {
					continue
				}
				relatedID = relationship.TargetID
			}
			if visited[relatedID] {
				continue
			}
			related, ok := nodeByID[relatedID]
			if !ok {
				continue
			}
			filePath := resourceNodeString(related, "filePath")
			if !options.IncludeTests && isImpactTestPath(filePath) {
				continue
			}
			visited[relatedID] = true
			nextFrontier = append(nextFrontier, relatedID)
			impacted = append(impacted, impactItemPayload(depth, related, relationship))
		}
		sort.Strings(nextFrontier)
		frontier = nextFrontier
	}
	if collectProfile {
		now := time.Now()
		profile.Traversal = impactProfileElapsed(last, now)
		last = now
	}

	grouped := groupImpactByDepth(impacted)
	affectedProcesses := impactAffectedProcesses(g, impacted)
	affectedModules := impactAffectedModules(g, impacted)
	semanticStatus := semantic.GraphSemanticStatus(g)
	semanticSummary := impactSemanticSummary(impacted)
	risk := impactRisk(len(grouped["1"]), len(affectedProcesses), len(affectedModules), len(impacted))
	if collectProfile {
		now := time.Now()
		profile.AffectedSummaries = impactProfileElapsed(last, now)
		last = now
	}

	targetPayload := map[string]any{
		"id":       target.ID,
		"name":     firstResourceNodeString(target, "name", "label", "heuristicLabel"),
		"type":     string(target.Label),
		"filePath": resourceNodeString(target, "filePath"),
	}
	addContextNodeSemanticFields(targetPayload, target)
	summary := map[string]any{
		"direct":             len(grouped["1"]),
		"processes_affected": len(affectedProcesses),
		"modules_affected":   len(affectedModules),
	}
	if appLayers := impactCountMapResult(semanticSummary, "affectedAppLayers"); len(appLayers) > 0 {
		summary["affected_app_layers"] = appLayers
	}
	if areas := impactCountMapResult(semanticSummary, "affectedFunctionalAreas"); len(areas) > 0 {
		summary["affected_functional_areas"] = areas
	}
	if risks, ok := semanticSummary["resolutionHealthRisks"].(map[string]any); ok && impactResolutionHealthRiskHasEvidence(risks) {
		summary["resolution_health_risks"] = risks
	}
	result := map[string]any{
		"target":                  targetPayload,
		"direction":               options.Direction,
		"impactedCount":           len(impacted),
		"risk":                    risk,
		"summary":                 summary,
		"semanticStatus":          semanticStatus,
		"affectedAppLayers":       semanticSummary["affectedAppLayers"],
		"affectedFunctionalAreas": semanticSummary["affectedFunctionalAreas"],
		"resolutionHealthRisks":   semanticSummary["resolutionHealthRisks"],
		"affected_processes":      affectedProcesses,
		"affected_modules":        affectedModules,
		"byDepth":                 grouped,
	}
	if warning := querySemanticWarning(semanticStatus); warning != "" {
		result["semanticWarning"] = warning
	}
	if warning := impactWorkflowWarning(risk); warning != "" {
		result["workflowWarning"] = warning
		result["workflowWarningBlocksOutput"] = false
	}
	if !traversalComplete {
		result["partial"] = true
	}
	if collectProfile {
		profile.Formatting = impactProfileElapsed(last, time.Now())
	}
	return result, profile
}

func impactProfileElapsed(start time.Time, end time.Time) time.Duration {
	elapsed := end.Sub(start)
	if elapsed <= 0 {
		return time.Nanosecond
	}
	return elapsed
}

func impactClassLikeSeedIDs(g *graph.Graph, targetID string) []string {
	seeds := make([]string, 0)
	for _, relationship := range g.Relationships {
		if relationship.Type == graph.RelHasMethod && relationship.SourceID == targetID {
			seeds = append(seeds, relationship.TargetID)
		}
		if relationship.Type == graph.RelDefines && relationship.TargetID == targetID {
			seeds = append(seeds, relationship.SourceID)
		}
	}
	return seeds
}

func impactItemPayload(depth int, node graph.Node, relationship graph.Relationship) map[string]any {
	payload := map[string]any{
		"depth":        depth,
		"id":           node.ID,
		"name":         firstResourceNodeString(node, "name", "label", "heuristicLabel"),
		"type":         string(node.Label),
		"filePath":     resourceNodeString(node, "filePath"),
		"relationType": string(relationship.Type),
		"confidence":   impactConfidence(relationship),
	}
	addContextNodeSemanticFields(payload, node)
	addImpactRelationshipProofFields(payload, relationship)
	if relationship.Reason != "" {
		payload["reason"] = relationship.Reason
	}
	if relationship.ResolutionSource != "" {
		payload["resolutionSource"] = relationship.ResolutionSource
	}
	if relationship.FileHash != "" {
		payload["fileHash"] = relationship.FileHash
	}
	if len(relationship.Evidence) > 0 {
		payload["evidence"] = relationship.Evidence
	}
	return payload
}

func addImpactRelationshipProofFields(payload map[string]any, relationship graph.Relationship) {
	if relationship.SourceSiteID != "" {
		payload["sourceSiteId"] = relationship.SourceSiteID
	}
	if len(relationship.SourceSiteIDs) > 0 {
		payload["sourceSiteIds"] = relationship.SourceSiteIDs
	}
	if relationship.SourceSiteCount > 0 {
		payload["sourceSiteCount"] = relationship.SourceSiteCount
	}
	if relationship.SourceSiteStatus != "" {
		payload["sourceSiteStatus"] = relationship.SourceSiteStatus
	}
	if relationship.ProofKind != "" {
		payload["proofKind"] = relationship.ProofKind
	}
	if relationship.TargetRole != "" {
		payload["targetRole"] = relationship.TargetRole
	}
	if relationship.TargetText != "" {
		payload["targetText"] = relationship.TargetText
	}
	if relationship.FilePath != "" {
		payload["relationshipFilePath"] = relationship.FilePath
	}
	if relationship.StartLine != 0 {
		payload["relationshipStartLine"] = relationship.StartLine
	}
	if relationship.StartCol != 0 {
		payload["relationshipStartCol"] = relationship.StartCol
	}
	if relationship.EndLine != 0 {
		payload["relationshipEndLine"] = relationship.EndLine
	}
	if relationship.EndCol != 0 {
		payload["relationshipEndCol"] = relationship.EndCol
	}
}

func impactSemanticSummary(impacted []map[string]any) map[string]any {
	appLayers := map[string]int{}
	functionalAreas := map[string]int{}
	resolutionConfidence := map[string]int{}
	resolutionBuckets := map[string]int{}
	riskNodes := make([]map[string]any, 0)
	totalGapCount := 0
	nodesWithGaps := 0
	degradedNodes := 0

	for _, item := range impacted {
		if appLayer := impactStringValue(item["appLayer"]); appLayer != "" {
			incrementQueryCount(appLayers, appLayer, 1)
		}
		if functionalArea := impactStringValue(item["functionalArea"]); functionalArea != "" {
			incrementQueryCount(functionalAreas, functionalArea, 1)
		}
		confidence := impactStringValue(item["resolutionConfidence"])
		if confidence != "" {
			incrementQueryCount(resolutionConfidence, confidence, 1)
		}
		if confidence == graphhealth.ResolutionConfidenceDegraded {
			degradedNodes++
		}
		gapCount := impactIntValue(item["resolutionGapCount"])
		if gapCount > 0 {
			totalGapCount += gapCount
			nodesWithGaps++
		}
		itemBuckets := impactCountMapValue(item["resolutionHealthBuckets"])
		for bucket, count := range itemBuckets {
			incrementQueryCount(resolutionBuckets, bucket, count)
		}
		if gapCount > 0 || confidence == graphhealth.ResolutionConfidenceDegraded || len(itemBuckets) > 0 {
			riskNode := map[string]any{
				"id":   item["id"],
				"name": item["name"],
				"type": item["type"],
			}
			if filePath := impactStringValue(item["filePath"]); filePath != "" {
				riskNode["filePath"] = filePath
			}
			if appLayer := impactStringValue(item["appLayer"]); appLayer != "" {
				riskNode["appLayer"] = appLayer
			}
			if functionalArea := impactStringValue(item["functionalArea"]); functionalArea != "" {
				riskNode["functionalArea"] = functionalArea
			}
			if confidence != "" {
				riskNode["resolutionConfidence"] = confidence
			}
			if gapCount > 0 {
				riskNode["resolutionGapCount"] = gapCount
			}
			if len(itemBuckets) > 0 {
				riskNode["resolutionHealthBuckets"] = itemBuckets
			}
			riskNodes = append(riskNodes, riskNode)
		}
	}

	return map[string]any{
		"affectedAppLayers":       cloneQueryCountMap(appLayers),
		"affectedFunctionalAreas": cloneQueryCountMap(functionalAreas),
		"resolutionHealthRisks": map[string]any{
			"nodesWithGaps":              nodesWithGaps,
			"degradedNodes":              degradedNodes,
			"totalResolutionGapCount":    totalGapCount,
			"resolutionConfidenceCounts": cloneQueryCountMap(resolutionConfidence),
			"resolutionHealthBuckets":    cloneQueryCountMap(resolutionBuckets),
			"nodes":                      riskNodes,
		},
	}
}

func impactWorkflowWarning(risk string) string {
	switch strings.ToUpper(strings.TrimSpace(risk)) {
	case "HIGH", "CRITICAL":
		return "Impact risk is " + strings.ToUpper(risk) + "; this is workflow safety information, not a blocker. Review affected symbols and flows before editing."
	default:
		return ""
	}
}

func impactResolutionHealthRiskHasEvidence(risks map[string]any) bool {
	return impactIntValue(risks["nodesWithGaps"]) > 0 ||
		impactIntValue(risks["degradedNodes"]) > 0 ||
		impactIntValue(risks["totalResolutionGapCount"]) > 0 ||
		len(impactCountMapValue(risks["resolutionHealthBuckets"])) > 0
}

func impactCountMapResult(summary map[string]any, key string) map[string]int {
	return impactCountMapValue(summary[key])
}

func impactStringValue(value any) string {
	text, ok := value.(string)
	if !ok {
		return ""
	}
	return strings.TrimSpace(text)
}

func impactIntValue(value any) int {
	switch typed := value.(type) {
	case int:
		return typed
	case int32:
		return int(typed)
	case int64:
		return int(typed)
	case float64:
		return int(typed)
	case float32:
		return int(typed)
	default:
		return 0
	}
}

func impactCountMapValue(value any) map[string]int {
	out := map[string]int{}
	switch typed := value.(type) {
	case map[string]int:
		for key, count := range typed {
			incrementQueryCount(out, key, count)
		}
	case map[string]any:
		for key, rawCount := range typed {
			incrementQueryCount(out, key, impactIntValue(rawCount))
		}
	}
	return cloneQueryCountMap(out)
}

func impactConfidence(relationship graph.Relationship) float64 {
	if relationship.Confidence > 0 {
		return relationship.Confidence
	}
	switch relationship.Type {
	case graph.RelCalls, graph.RelImports, graph.RelDefines:
		return 0.95
	case graph.RelUses, graph.RelInherits, graph.RelExtends, graph.RelImplements:
		return 0.85
	default:
		return 0.7
	}
}

func groupImpactByDepth(impacted []map[string]any) map[string][]map[string]any {
	grouped := make(map[string][]map[string]any)
	for _, item := range impacted {
		depth := fmt.Sprint(item["depth"])
		grouped[depth] = append(grouped[depth], item)
	}
	for depth := range grouped {
		sort.Slice(grouped[depth], func(i, j int) bool {
			left, right := grouped[depth][i], grouped[depth][j]
			if left["filePath"] != right["filePath"] {
				return fmt.Sprint(left["filePath"]) < fmt.Sprint(right["filePath"])
			}
			return fmt.Sprint(left["id"]) < fmt.Sprint(right["id"])
		})
	}
	return grouped
}

func impactAffectedProcesses(g *graph.Graph, impacted []map[string]any) []map[string]any {
	impactedSet := impactIDSet(impacted)
	nodeByID := resourceGraphNodesByID(g)
	type processStats struct {
		Name        string
		Type        string
		FilePath    string
		Count       int
		TotalHits   int
		EarliestSet bool
		Earliest    int
	}
	processes := make(map[string]*processStats)
	for _, relationship := range g.Relationships {
		if relationship.Type != graph.RelStepInProcess || !impactedSet[relationship.SourceID] {
			continue
		}
		process, ok := nodeByID[relationship.TargetID]
		if !ok {
			continue
		}
		stats := processes[process.ID]
		if stats == nil {
			stats = &processStats{
				Name:     firstResourceNodeString(process, "heuristicLabel", "label", "name"),
				Type:     resourceNodeString(process, "processType"),
				FilePath: resourceNodeString(process, "filePath"),
			}
			processes[process.ID] = stats
		}
		stats.Count++
		stats.TotalHits++
		if relationship.Step != nil && (!stats.EarliestSet || *relationship.Step < stats.Earliest) {
			stats.Earliest = *relationship.Step
			stats.EarliestSet = true
		}
	}
	out := make([]map[string]any, 0, len(processes))
	for processID, stats := range processes {
		earliest := any(nil)
		if stats.EarliestSet {
			earliest = stats.Earliest
		}
		out = append(out, map[string]any{
			"name":                   stats.Name,
			"processType":            firstNonEmptyString(stats.Type, "process"),
			"filePath":               stats.FilePath,
			"affected_process_count": stats.Count,
			"total_hits":             stats.TotalHits,
			"earliest_broken_step":   earliest,
		})
		if processNode, ok := nodeByID[processID]; ok {
			addContextNodeSemanticFields(out[len(out)-1], processNode)
		}
	}
	sort.Slice(out, func(i, j int) bool {
		left, _ := out[i]["total_hits"].(int)
		right, _ := out[j]["total_hits"].(int)
		if left != right {
			return left > right
		}
		return fmt.Sprint(out[i]["name"]) < fmt.Sprint(out[j]["name"])
	})
	return out
}

func impactAffectedModules(g *graph.Graph, impacted []map[string]any) []map[string]any {
	impactedSet := impactIDSet(impacted)
	directSet := make(map[string]bool)
	for _, item := range impacted {
		if item["depth"] == 1 {
			directSet[fmt.Sprint(item["id"])] = true
		}
	}
	nodeByID := resourceGraphNodesByID(g)
	type moduleStats struct {
		Node   graph.Node
		Name   string
		Hits   int
		Direct bool
	}
	modules := make(map[string]*moduleStats)
	for _, relationship := range g.Relationships {
		if relationship.Type != graph.RelMemberOf || !impactedSet[relationship.SourceID] {
			continue
		}
		module, ok := nodeByID[relationship.TargetID]
		if !ok {
			continue
		}
		name := firstResourceNodeString(module, "heuristicLabel", "label", "name")
		if name == "" {
			continue
		}
		stats := modules[name]
		if stats == nil {
			stats = &moduleStats{Name: name, Node: module}
			modules[name] = stats
		}
		stats.Hits++
		if directSet[relationship.SourceID] {
			stats.Direct = true
		}
	}
	out := make([]map[string]any, 0, len(modules))
	for _, stats := range modules {
		impact := "indirect"
		if stats.Direct {
			impact = "direct"
		}
		row := map[string]any{"name": stats.Name, "hits": stats.Hits, "impact": impact}
		addContextNodeSemanticFields(row, stats.Node)
		out = append(out, row)
	}
	sort.Slice(out, func(i, j int) bool {
		left, _ := out[i]["hits"].(int)
		right, _ := out[j]["hits"].(int)
		if left != right {
			return left > right
		}
		return fmt.Sprint(out[i]["name"]) < fmt.Sprint(out[j]["name"])
	})
	return out
}

func impactIDSet(items []map[string]any) map[string]bool {
	out := make(map[string]bool, len(items))
	for _, item := range items {
		id := fmt.Sprint(item["id"])
		if id != "" {
			out[id] = true
		}
	}
	return out
}

func impactRisk(directCount int, processCount int, moduleCount int, impactedCount int) string {
	switch {
	case directCount >= 30 || processCount >= 5 || moduleCount >= 5 || impactedCount >= 200:
		return "CRITICAL"
	case directCount >= 15 || processCount >= 3 || moduleCount >= 3 || impactedCount >= 100:
		return "HIGH"
	case directCount >= 5 || impactedCount >= 30:
		return "MEDIUM"
	default:
		return "LOW"
	}
}

func isImpactTestPath(filePath string) bool {
	normalized := normalizeContextPath(filePath)
	return strings.Contains(normalized, "/test/") ||
		strings.Contains(normalized, "/tests/") ||
		strings.Contains(normalized, "/__tests__/") ||
		strings.HasPrefix(normalized, "test/") ||
		strings.HasPrefix(normalized, "tests/") ||
		strings.HasPrefix(normalized, "__tests__/") ||
		strings.HasSuffix(normalized, "_test.go") ||
		strings.HasSuffix(normalized, "_test.py") ||
		strings.HasSuffix(normalized, ".test.ts") ||
		strings.HasSuffix(normalized, ".test.tsx") ||
		strings.HasSuffix(normalized, ".spec.ts") ||
		strings.HasSuffix(normalized, ".spec.tsx")
}

func sortedImpactAllowedRelationTypes() []string {
	out := make([]string, 0, len(impactAllowedRelationTypes))
	for relType := range impactAllowedRelationTypes {
		out = append(out, relType)
	}
	sort.Strings(out)
	return out
}

func dedupeStrings(values []string) []string {
	seen := make(map[string]bool, len(values))
	out := make([]string, 0, len(values))
	for _, value := range values {
		if seen[value] {
			continue
		}
		seen[value] = true
		out = append(out, value)
	}
	return out
}

func floatArg(args map[string]any, name string, fallback float64, minValue float64, maxValue float64) float64 {
	if args == nil {
		return fallback
	}
	var value float64
	switch raw := args[name].(type) {
	case int:
		value = float64(raw)
	case float64:
		value = raw
	default:
		return fallback
	}
	if value < minValue {
		return minValue
	}
	if value > maxValue {
		return maxValue
	}
	return value
}
