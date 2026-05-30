package mcp

import (
	"fmt"
	"sort"
	"strings"

	"github.com/tamnguyendinh/anvien/internal/filecontext"
	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

const (
	targetTypeAuto    = "auto"
	targetTypeSymbol  = "symbol"
	targetTypeFile    = "file"
	targetTypeRoute   = "route"
	targetTypeTool    = "tool"
	targetTypeFiles   = "files"
	targetTypeSymbols = "symbols"
	targetTypeFlows   = "flows"
	targetTypeAPI     = "api"

	dispatchModeExplicit = "explicit"
	dispatchModeSmart    = "smart"
	dispatchModeLegacy   = "legacy"
)

func normalizedTargetType(args map[string]any, defaultType string, allowed map[string]bool) string {
	targetType := strings.TrimSpace(stringArg(args, "target_type", ""))
	if targetType == "" {
		targetType = defaultType
	}
	targetType = strings.ToLower(strings.ReplaceAll(targetType, "-", "_"))
	if allowed[targetType] {
		return targetType
	}
	return defaultType
}

func normalizedDispatchMode(args map[string]any, targetType string) string {
	dispatchMode := strings.TrimSpace(stringArg(args, "dispatch_mode", ""))
	if dispatchMode != "" {
		return dispatchMode
	}
	if targetType == targetTypeAuto {
		return dispatchModeSmart
	}
	if strings.TrimSpace(stringArg(args, "target_type", "")) == "" {
		return dispatchModeLegacy
	}
	return dispatchModeExplicit
}

func mcpBuildFileContext(g *graph.Graph, path string, dispatchMode string) (filecontext.FileContext, bool) {
	context, ok := filecontext.NewBuilder(g).BuildFileContext(path, filecontext.Options{
		RelationshipSamplesPerGroup: 5,
		UnresolvedSamplesPerGroup:   5,
		LinkedSamplesPerKind:        5,
	})
	if !ok {
		return filecontext.FileContext{}, false
	}
	context.Target.DispatchMode = dispatchMode
	return context, true
}

func mcpFileSummaryForPath(g *graph.Graph, path string) (filecontext.FileSummary, bool) {
	context, ok := mcpBuildFileContext(g, path, dispatchModeExplicit)
	if !ok {
		return filecontext.FileSummary{}, false
	}
	return context.Summary, true
}

func addMCPDispatchFields(payload map[string]any, targetType string, dispatchMode string) {
	payload["targetType"] = targetType
	payload["dispatchMode"] = dispatchMode
}

func addMCPSymbolTargetFields(payload map[string]any, symbol map[string]any, summary filecontext.FileSummary, hasSummary bool, dispatchMode string) {
	addMCPDispatchFields(payload, targetTypeSymbol, dispatchMode)
	payload["selectedSymbol"] = symbol
	if hasSummary {
		payload["selectedFile"] = map[string]any{
			"path":           summary.Path,
			"kind":           summary.Kind,
			"appLayer":       summary.AppLayer,
			"functionalArea": summary.FunctionalArea,
			"risk":           summary.Risk,
			"symbolCount":    summary.SymbolCount,
			"unresolved":     summary.UnresolvedSourceSiteCount,
			"fanIn":          summary.InboundRefCount,
			"fanOut":         summary.OutboundRefCount,
		}
		payload["fileSummary"] = summary
	}
}

func contextFileCandidate(path string, input string) map[string]any {
	return map[string]any{
		"targetType":       targetTypeFile,
		"type":             targetTypeFile,
		"name":             path,
		"filePath":         path,
		"confidence":       1.0,
		"suggestedCommand": fmt.Sprintf("anvien context file %q", path),
		"input":            input,
	}
}

func addContextCandidateSuggestions(candidates []map[string]any, commandName string) []map[string]any {
	out := make([]map[string]any, 0, len(candidates))
	for _, candidate := range candidates {
		item := cloneMap(candidate)
		item["targetType"] = targetTypeSymbol
		if _, ok := item["suggestedCommand"]; !ok {
			name := strings.TrimSpace(fmt.Sprint(item["name"]))
			uid := strings.TrimSpace(fmt.Sprint(item["uid"]))
			if uid != "" {
				item["suggestedCommand"] = fmt.Sprintf("anvien %s symbol %q --uid %q", commandName, name, uid)
			} else {
				item["suggestedCommand"] = fmt.Sprintf("anvien %s symbol %q", commandName, name)
			}
		}
		out = append(out, item)
	}
	return out
}

func cloneMap(input map[string]any) map[string]any {
	out := make(map[string]any, len(input))
	for key, value := range input {
		out[key] = value
	}
	return out
}

func flattenFileSymbols(nodes []filecontext.SymbolTreeNode) []filecontext.SymbolTreeNode {
	out := make([]filecontext.SymbolTreeNode, 0)
	var walk func([]filecontext.SymbolTreeNode)
	walk = func(items []filecontext.SymbolTreeNode) {
		for _, item := range items {
			out = append(out, item)
			if len(item.Children) > 0 {
				walk(item.Children)
			}
		}
	}
	walk(nodes)
	return out
}

func nodeByID(g *graph.Graph) map[string]graph.Node {
	out := make(map[string]graph.Node, len(g.Nodes))
	for _, node := range g.Nodes {
		out[node.ID] = node
	}
	return out
}

func fileImpactAffectedFiles(rows []map[string]any) []map[string]any {
	counts := map[string]int{}
	for _, row := range rows {
		path := strings.TrimSpace(impactStringValue(row["filePath"]))
		if path == "" {
			continue
		}
		counts[path]++
	}
	paths := make([]string, 0, len(counts))
	for path := range counts {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	out := make([]map[string]any, 0, len(paths))
	for _, path := range paths {
		out = append(out, map[string]any{"path": path, "impactedSymbols": counts[path]})
	}
	return out
}

func combineImpactProcesses(payloads []map[string]any) []map[string]any {
	seen := map[string]map[string]any{}
	for _, payload := range payloads {
		processes, _ := payload["affected_processes"].([]map[string]any)
		for _, process := range processes {
			id := strings.TrimSpace(fmt.Sprint(process["id"]))
			if id == "" {
				id = strings.TrimSpace(fmt.Sprint(process["name"]))
			}
			if id == "" || seen[id] != nil {
				continue
			}
			seen[id] = process
		}
	}
	keys := make([]string, 0, len(seen))
	for key := range seen {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	out := make([]map[string]any, 0, len(keys))
	for _, key := range keys {
		out = append(out, seen[key])
	}
	return out
}

func symbolRowsFromFileContext(context filecontext.FileContext, nodes map[string]graph.Node) []map[string]any {
	symbols := flattenFileSymbols(context.SymbolTree)
	rows := make([]map[string]any, 0, len(symbols))
	for _, symbol := range symbols {
		row := map[string]any{
			"id":        symbol.ID,
			"name":      symbol.Name,
			"type":      symbol.Kind,
			"filePath":  context.Summary.Path,
			"startLine": symbol.Range.StartLine,
			"endLine":   symbol.Range.EndLine,
		}
		if node, ok := nodes[symbol.ID]; ok {
			addContextNodeSemanticFields(row, node)
			addContextResolutionGapEntityFields(row, node)
		}
		rows = append(rows, row)
	}
	return rows
}

func queryTargetTypeAllowed() map[string]bool {
	return map[string]bool{
		"":                true,
		targetTypeFiles:   true,
		targetTypeSymbols: true,
		targetTypeFlows:   true,
		targetTypeAPI:     true,
	}
}

func contextTargetTypeAllowed() map[string]bool {
	return map[string]bool{
		targetTypeAuto:   true,
		targetTypeSymbol: true,
		targetTypeFile:   true,
	}
}

func impactTargetTypeAllowed() map[string]bool {
	return map[string]bool{
		targetTypeAuto:   true,
		targetTypeSymbol: true,
		targetTypeFile:   true,
		targetTypeRoute:  true,
		targetTypeTool:   true,
	}
}

func detectChangesTargetTypeAllowed() map[string]bool {
	return map[string]bool{
		"":                true,
		targetTypeFiles:   true,
		targetTypeSymbols: true,
		targetTypeFlows:   true,
	}
}

func graphDefinitionLabelsForSymbols(label scopeir.NodeLabel) bool {
	return label != scopeir.NodeFile && label != scopeir.NodeRoute && label != scopeir.NodeTool && queryDefinitionLabel(label)
}
