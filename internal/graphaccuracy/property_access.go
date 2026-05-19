package graphaccuracy

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type PropertyAccessAuditOptions struct {
	Repo        string
	GraphPath   string
	OutPath     string
	MaxExamples int
}

type PropertyAccessAuditResult struct {
	GeneratedAt string `json:"generatedAt"`
	Inputs      struct {
		Repo  string `json:"repo"`
		Graph string `json:"graph"`
	} `json:"inputs"`
	Totals        PropertyAccessTotals             `json:"totals"`
	Languages     map[string]PropertyLanguageStats `json:"languages"`
	Categories    map[string]PropertyAccessBucket  `json:"categories"`
	OrphanStatus  map[string]PropertyAccessBucket  `json:"orphanStatus"`
	GraphTruth    map[string]PropertyAccessBucket  `json:"graphTruth"`
	InvalidLinks  PropertyAccessBucket             `json:"invalidLinks"`
	Relationships PropertyRelationshipSummary      `json:"relationships"`
	Notes         []string                         `json:"notes"`
}

type PropertyAccessTotals struct {
	Nodes                   int `json:"nodes"`
	Relationships           int `json:"relationships"`
	PropertyNodes           int `json:"propertyNodes"`
	OwnerLinkedProperties   int `json:"ownerLinkedProperties"`
	StandaloneProperties    int `json:"standaloneProperties"`
	HasPropertyEdges        int `json:"hasPropertyEdges"`
	AccessesEdges           int `json:"accessesEdges"`
	InvalidHasPropertyEdges int `json:"invalidHasPropertyEdges"`
}

type PropertyLanguageStats struct {
	PropertyNodes         int            `json:"propertyNodes"`
	OwnerLinkedProperties int            `json:"ownerLinkedProperties"`
	StandaloneProperties  int            `json:"standaloneProperties"`
	Categories            map[string]int `json:"categories,omitempty"`
	OrphanStatus          map[string]int `json:"orphanStatus,omitempty"`
	GraphTruth            map[string]int `json:"graphTruth,omitempty"`
}

type PropertyRelationshipSummary struct {
	HasPropertyByOwnerLabel map[string]int `json:"hasPropertyByOwnerLabel"`
	AccessesBySourceLabel   map[string]int `json:"accessesBySourceLabel"`
	AccessesByTargetLabel   map[string]int `json:"accessesByTargetLabel"`
}

type PropertyAccessBucket struct {
	Count    int                     `json:"count"`
	Examples []PropertyAccessExample `json:"examples,omitempty"`
}

type PropertyAccessExample struct {
	ID            string `json:"id"`
	Name          string `json:"name,omitempty"`
	FilePath      string `json:"filePath,omitempty"`
	Language      string `json:"language,omitempty"`
	QualifiedName string `json:"qualifiedName,omitempty"`
	DeclaredType  string `json:"declaredType,omitempty"`
	StartLine     int    `json:"startLine,omitempty"`
	SourceLine    string `json:"sourceLine,omitempty"`
	OwnerID       string `json:"ownerId,omitempty"`
	OwnerLabel    string `json:"ownerLabel,omitempty"`
	Category      string `json:"category,omitempty"`
	OrphanStatus  string `json:"orphanStatus,omitempty"`
	GraphTruth    string `json:"graphTruth,omitempty"`
	Reason        string `json:"reason,omitempty"`
}

type sourceLineCache struct {
	repoAbs string
	files   map[string][]string
}

func RunPropertyAccessAudit(options PropertyAccessAuditOptions) (PropertyAccessAuditResult, error) {
	if options.GraphPath == "" {
		return PropertyAccessAuditResult{}, fmt.Errorf("graph path is required")
	}
	if options.MaxExamples <= 0 {
		options.MaxExamples = 50
	}
	repo := options.Repo
	if strings.TrimSpace(repo) == "" {
		repo = "."
	}
	repoAbs, err := filepath.Abs(repo)
	if err != nil {
		return PropertyAccessAuditResult{}, fmt.Errorf("resolve repo: %w", err)
	}
	graphFile, err := ReadGraph(options.GraphPath)
	if err != nil {
		return PropertyAccessAuditResult{}, err
	}
	result := buildPropertyAccessAudit(repoAbs, options.GraphPath, graphFile, options.MaxExamples)
	if options.OutPath != "" {
		if err := WritePropertyAccessAuditResult(options.OutPath, result); err != nil {
			return PropertyAccessAuditResult{}, err
		}
	}
	return result, nil
}

func WritePropertyAccessAuditResult(path string, result PropertyAccessAuditResult) error {
	raw, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal property/access audit: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, append(raw, '\n'), 0o644)
}

func PropertyAccessAuditSummaryLines(result PropertyAccessAuditResult) []string {
	lines := []string{
		fmt.Sprintf("properties.total=%d ownerLinked=%d standalone=%d hasPropertyEdges=%d accessesEdges=%d invalidHasPropertyEdges=%d",
			result.Totals.PropertyNodes,
			result.Totals.OwnerLinkedProperties,
			result.Totals.StandaloneProperties,
			result.Totals.HasPropertyEdges,
			result.Totals.AccessesEdges,
			result.Totals.InvalidHasPropertyEdges,
		),
	}
	for _, language := range sortedLanguageKeys(result.Languages) {
		stats := result.Languages[language]
		lines = append(lines, fmt.Sprintf("language.%s.properties=%d ownerLinked=%d standalone=%d",
			language,
			stats.PropertyNodes,
			stats.OwnerLinkedProperties,
			stats.StandaloneProperties,
		))
	}
	for _, key := range sortedBucketKeys(result.OrphanStatus) {
		lines = append(lines, fmt.Sprintf("orphan.%s=%d", key, result.OrphanStatus[key].Count))
	}
	for _, key := range sortedBucketKeys(result.GraphTruth) {
		lines = append(lines, fmt.Sprintf("graphTruth.%s=%d", key, result.GraphTruth[key].Count))
	}
	return lines
}

func buildPropertyAccessAudit(repoAbs string, graphPath string, graphFile GraphFile, maxExamples int) PropertyAccessAuditResult {
	nodesByID := make(map[string]GraphNode, len(graphFile.Nodes))
	for _, node := range graphFile.Nodes {
		nodesByID[node.ID] = node
	}
	incomingHasProperty := map[string][]GraphRelationship{}
	cache := sourceLineCache{repoAbs: repoAbs, files: map[string][]string{}}

	var result PropertyAccessAuditResult
	result.GeneratedAt = time.Now().Format(time.RFC3339)
	result.Inputs.Repo = repoAbs
	result.Inputs.Graph = graphPath
	result.Totals.Nodes = len(graphFile.Nodes)
	result.Totals.Relationships = len(graphFile.Relationships)
	result.Languages = map[string]PropertyLanguageStats{}
	result.Categories = map[string]PropertyAccessBucket{}
	result.OrphanStatus = map[string]PropertyAccessBucket{}
	result.GraphTruth = map[string]PropertyAccessBucket{}
	result.Relationships.HasPropertyByOwnerLabel = map[string]int{}
	result.Relationships.AccessesBySourceLabel = map[string]int{}
	result.Relationships.AccessesByTargetLabel = map[string]int{}

	for _, relationship := range graphFile.Relationships {
		switch relationship.Type {
		case "HAS_PROPERTY":
			result.Totals.HasPropertyEdges++
			source, sourceOK := nodesByID[relationship.SourceID]
			target, targetOK := nodesByID[relationship.TargetID]
			if !sourceOK || !targetOK || target.Label != "Property" {
				result.Totals.InvalidHasPropertyEdges++
				example := relationshipExample(relationship, source, target, "invalid_has_property_edge", "HAS_PROPERTY source/target is missing or target is not Property")
				addAuditExample(&result.InvalidLinks, example, maxExamples)
				continue
			}
			incomingHasProperty[relationship.TargetID] = append(incomingHasProperty[relationship.TargetID], relationship)
			result.Relationships.HasPropertyByOwnerLabel[source.Label]++
		case "ACCESSES":
			result.Totals.AccessesEdges++
			source, sourceOK := nodesByID[relationship.SourceID]
			target, targetOK := nodesByID[relationship.TargetID]
			if sourceOK {
				result.Relationships.AccessesBySourceLabel[source.Label]++
			}
			if targetOK {
				result.Relationships.AccessesByTargetLabel[target.Label]++
			}
		}
	}

	for _, node := range graphFile.Nodes {
		if node.Label != "Property" {
			continue
		}
		language := propertyLanguage(node)
		result.Totals.PropertyNodes++
		ownerRel, owner, ownerLinked := firstOwnerLink(node.ID, incomingHasProperty, nodesByID)
		if ownerLinked {
			result.Totals.OwnerLinkedProperties++
		} else {
			result.Totals.StandaloneProperties++
		}
		line, context := cache.contextFor(node)
		category := classifyProperty(node, owner, ownerLinked, language, line, context)
		orphanStatus := classifyPropertyOrphanStatus(ownerLinked, category)
		graphTruth := classifyPropertyGraphTruth(orphanStatus)
		example := propertyExample(node, ownerRel, owner, language, category, orphanStatus, graphTruth, line)
		addAuditBucket(result.Categories, category, example, maxExamples)
		addAuditBucket(result.OrphanStatus, orphanStatus, example, maxExamples)
		addAuditBucket(result.GraphTruth, graphTruth, example, maxExamples)
		addLanguageStats(result.Languages, language, category, orphanStatus, graphTruth, ownerLinked)
	}
	ensureAuditBuckets(result.OrphanStatus,
		"owner_linked",
		"false_orphan",
		"true_orphan",
		"unknown",
		"external_library_owned",
		"intentionally_unmodeled",
	)
	ensureAuditBuckets(result.GraphTruth,
		"edge_present",
		"real_edge_missing",
		"true_no_edge",
		"unknown_no_edge",
		"invalid_synthetic_edge_risk",
	)

	result.Notes = []string{
		"ACCESSES and HAS_PROPERTY are counted from final graph relationships.",
		"The gate includes every Property node in the graph, across every language AVmatrix-Go currently emits.",
		"True orphan and false orphan classifications are conservative taxonomy outputs; they are not instructions to force-link all Property nodes.",
		"Unknown no-edge cases must remain unlinked until source evidence proves a real relationship.",
		"Invalid HAS_PROPERTY edges are structural problems: missing endpoint nodes or non-Property targets.",
	}
	return result
}

func firstOwnerLink(propertyID string, incoming map[string][]GraphRelationship, nodes map[string]GraphNode) (GraphRelationship, GraphNode, bool) {
	rels := incoming[propertyID]
	if len(rels) == 0 {
		return GraphRelationship{}, GraphNode{}, false
	}
	rel := rels[0]
	owner, ok := nodes[rel.SourceID]
	if !ok {
		return rel, GraphNode{}, false
	}
	return rel, owner, true
}

func propertyLanguage(node GraphNode) string {
	language := strings.ToLower(propString(node.Properties, "language"))
	if language != "" {
		return language
	}
	filePath := filepath.ToSlash(propString(node.Properties, "filePath"))
	switch strings.ToLower(filepath.Ext(filePath)) {
	case ".js", ".jsx", ".mjs", ".cjs":
		return "javascript"
	case ".ts", ".tsx", ".mts", ".cts":
		return "typescript"
	case ".py":
		return "python"
	case ".java":
		return "java"
	case ".c":
		return "c"
	case ".cpp", ".cc", ".cxx", ".h", ".hpp", ".hxx", ".hh":
		return "cpp"
	case ".cs":
		return "csharp"
	case ".go":
		return "go"
	case ".rb", ".rake", ".gemspec":
		return "ruby"
	case ".rs":
		return "rust"
	case ".php", ".phtml", ".php3", ".php4", ".php5", ".php8":
		return "php"
	case ".kt", ".kts":
		return "kotlin"
	case ".swift":
		return "swift"
	case ".dart":
		return "dart"
	case ".vue":
		return "vue"
	case ".svelte":
		return "svelte"
	case ".astro":
		return "astro"
	case ".cbl", ".cob", ".cpy", ".cobol", ".copybook", ".jcl", ".job", ".proc":
		return "cobol"
	}
	return "unknown"
}

func classifyProperty(node GraphNode, owner GraphNode, ownerLinked bool, language string, line string, context []string) string {
	if isTSJSLanguage(language) {
		return classifyTSJSProperty(node, owner, ownerLinked, line, context)
	}
	if language == "go" {
		return classifyGoProperty(node, owner, ownerLinked, line, context)
	}
	if ownerLinked {
		return language + "_owner_linked_" + strings.ToLower(owner.Label)
	}
	if propertyInExternalPath(node) {
		return language + "_external_library_property"
	}
	if qualifiedNameHasOwner(propString(node.Properties, "qualifiedName"), propString(node.Properties, "name")) {
		return language + "_qualified_member_without_owner"
	}
	if propString(node.Properties, "declaredType") != "" {
		return language + "_typed_property_without_owner"
	}
	return language + "_unclassified_property"
}

func classifyGoProperty(node GraphNode, owner GraphNode, ownerLinked bool, line string, context []string) string {
	if ownerLinked {
		return "go_owner_linked_" + strings.ToLower(owner.Label)
	}
	if propertyInExternalPath(node) {
		return "go_external_library_property"
	}
	if qualifiedNameHasOwner(propString(node.Properties, "qualifiedName"), propString(node.Properties, "name")) {
		return "go_qualified_member_without_owner"
	}
	if looksGoAnonymousStructField(line, context) {
		return "go_anonymous_struct_field"
	}
	if propString(node.Properties, "declaredType") != "" {
		return "go_typed_property_without_owner"
	}
	return "go_unclassified_property"
}

func classifyTSJSProperty(node GraphNode, owner GraphNode, ownerLinked bool, line string, context []string) string {
	if ownerLinked {
		switch owner.Label {
		case "Class":
			return "tsjs_class_field"
		case "Interface":
			return "tsjs_interface_property_signature"
		case "TypeAlias":
			return "tsjs_type_alias_object_literal_member"
		default:
			return "tsjs_owner_linked_" + strings.ToLower(owner.Label)
		}
	}
	if propertyInExternalPath(node) {
		return "tsjs_external_library_property"
	}
	if looksInsideInterface(context) {
		return "tsjs_interface_property_signature"
	}
	if looksInsideTypeAliasObject(context) {
		return "tsjs_type_alias_object_literal_member"
	}
	if looksRuntimeObjectLiteral(context) {
		return "tsjs_runtime_object_literal_key"
	}
	if looksDestructuringOrBinding(line, context) {
		return "tsjs_destructuring_or_binding_pattern"
	}
	if propString(node.Properties, "declaredType") != "" {
		return "tsjs_typed_shape_or_binding_property"
	}
	if qualifiedNameHasOwner(propString(node.Properties, "qualifiedName"), propString(node.Properties, "name")) {
		return "tsjs_qualified_member_without_owner"
	}
	return "tsjs_unclassified_property"
}

func classifyPropertyOrphanStatus(ownerLinked bool, category string) string {
	if ownerLinked {
		return "owner_linked"
	}
	switch category {
	case "tsjs_interface_property_signature", "tsjs_type_alias_object_literal_member", "tsjs_class_field":
		return "false_orphan"
	case "tsjs_runtime_object_literal_key", "tsjs_destructuring_or_binding_pattern", "go_anonymous_struct_field":
		return "true_orphan"
	case "tsjs_typed_shape_or_binding_property":
		return "unknown"
	default:
		if strings.Contains(category, "_external_library_property") {
			return "external_library_owned"
		}
		if strings.Contains(category, "_intentionally_unmodeled_property") {
			return "intentionally_unmodeled"
		}
		if strings.Contains(category, "_qualified_member_without_owner") {
			return "false_orphan"
		}
		return "unknown"
	}
}

func propertyInExternalPath(node GraphNode) bool {
	filePath := "/" + filepath.ToSlash(strings.ToLower(propString(node.Properties, "filePath"))) + "/"
	return strings.Contains(filePath, "/node_modules/") ||
		strings.Contains(filePath, "/vendor/") ||
		strings.Contains(filePath, "/.cargo/") ||
		strings.Contains(filePath, "/pkg/mod/")
}

func qualifiedNameHasOwner(qualifiedName string, name string) bool {
	qualifiedName = strings.TrimSpace(qualifiedName)
	name = strings.TrimSpace(name)
	if qualifiedName == "" || qualifiedName == name {
		return false
	}
	return strings.Contains(qualifiedName, ".") || strings.Contains(qualifiedName, "::")
}

func isTSJSLanguage(language string) bool {
	switch language {
	case "javascript", "typescript", "tsx", "jsx", "vue", "svelte", "astro":
		return true
	default:
		return false
	}
}

func classifyPropertyGraphTruth(orphanStatus string) string {
	switch orphanStatus {
	case "owner_linked":
		return "edge_present"
	case "false_orphan":
		return "real_edge_missing"
	case "true_orphan":
		return "true_no_edge"
	default:
		return "unknown_no_edge"
	}
}

func looksInsideInterface(context []string) bool {
	return contextContainsBlockHeader(context, "interface ")
}

func looksInsideTypeAliasObject(context []string) bool {
	for i := len(context) - 1; i >= 0; i-- {
		line := strings.TrimSpace(context[i])
		if strings.Contains(line, "type ") && strings.Contains(line, "= {") {
			return true
		}
		if strings.Contains(line, "= {") {
			return false
		}
		if strings.Contains(line, "interface ") || strings.Contains(line, "class ") || strings.Contains(line, "function ") {
			return false
		}
	}
	return false
}

func looksRuntimeObjectLiteral(context []string) bool {
	for i := len(context) - 1; i >= 0; i-- {
		line := strings.TrimSpace(context[i])
		if strings.Contains(line, "type ") || strings.Contains(line, "interface ") || strings.Contains(line, "class ") {
			return false
		}
		if strings.Contains(line, "return {") || strings.Contains(line, "= {") {
			return true
		}
	}
	return false
}

func looksDestructuringOrBinding(line string, context []string) bool {
	trimmed := strings.TrimSpace(line)
	if strings.Contains(trimmed, "{") && strings.Contains(trimmed, "}") {
		return true
	}
	for i := len(context) - 1; i >= 0; i-- {
		previous := strings.TrimSpace(context[i])
		if strings.Contains(previous, "function ") && strings.Contains(previous, "({") {
			return true
		}
		if strings.Contains(previous, "=>") && strings.Contains(previous, "({") {
			return true
		}
	}
	return false
}

func looksGoAnonymousStructField(line string, context []string) bool {
	if !goFieldLineHasType(line) {
		return false
	}
	for i := len(context) - 1; i >= 0; i-- {
		trimmed := strings.TrimSpace(context[i])
		if trimmed == "" {
			continue
		}
		if strings.Contains(trimmed, "struct {") {
			return !strings.HasPrefix(trimmed, "type ")
		}
		if strings.HasPrefix(trimmed, "type ") || strings.HasPrefix(trimmed, "func ") {
			return false
		}
	}
	return false
}

func goFieldLineHasType(line string) bool {
	line = strings.TrimSpace(line)
	if line == "" || strings.HasPrefix(line, "}") || strings.HasPrefix(line, "//") {
		return false
	}
	line = strings.TrimSpace(strings.Split(line, "`")[0])
	fields := strings.Fields(line)
	return len(fields) >= 2
}

func contextContainsBlockHeader(context []string, needle string) bool {
	for i := len(context) - 1; i >= 0; i-- {
		line := strings.TrimSpace(context[i])
		if strings.Contains(line, needle) {
			return true
		}
		if strings.Contains(line, "type ") || strings.Contains(line, "class ") || strings.Contains(line, "function ") || strings.Contains(line, "return {") || strings.Contains(line, "= {") {
			return false
		}
	}
	return false
}

func (c *sourceLineCache) contextFor(node GraphNode) (string, []string) {
	filePath := filepath.ToSlash(propString(node.Properties, "filePath"))
	startLine := propInt(node.Properties, "startLine")
	if filePath == "" || startLine <= 0 {
		return "", nil
	}
	lines, ok := c.files[filePath]
	if !ok {
		raw, err := os.ReadFile(filepath.Join(c.repoAbs, filepath.FromSlash(filePath)))
		if err != nil {
			c.files[filePath] = nil
			return "", nil
		}
		lines = strings.Split(string(raw), "\n")
		c.files[filePath] = lines
	}
	index := startLine - 1
	if index < 0 || index >= len(lines) {
		return "", nil
	}
	contextStart := index - 40
	if contextStart < 0 {
		contextStart = 0
	}
	return strings.TrimSpace(lines[index]), lines[contextStart : index+1]
}

func propertyExample(node GraphNode, ownerRel GraphRelationship, owner GraphNode, language string, category string, orphanStatus string, graphTruth string, line string) PropertyAccessExample {
	example := PropertyAccessExample{
		ID:            node.ID,
		Name:          propString(node.Properties, "name"),
		FilePath:      filepath.ToSlash(propString(node.Properties, "filePath")),
		Language:      language,
		QualifiedName: propString(node.Properties, "qualifiedName"),
		DeclaredType:  propString(node.Properties, "declaredType"),
		StartLine:     propInt(node.Properties, "startLine"),
		SourceLine:    line,
		Category:      category,
		OrphanStatus:  orphanStatus,
		GraphTruth:    graphTruth,
	}
	if ownerRel.ID != "" {
		example.OwnerID = ownerRel.SourceID
		example.OwnerLabel = owner.Label
	}
	return example
}

func relationshipExample(relationship GraphRelationship, source GraphNode, target GraphNode, category string, reason string) PropertyAccessExample {
	return PropertyAccessExample{
		ID:         relationship.ID,
		OwnerID:    relationship.SourceID,
		OwnerLabel: source.Label,
		Category:   category,
		GraphTruth: "invalid_synthetic_edge_risk",
		Reason:     reason,
		Name:       target.Label,
	}
}

func addAuditBucket(buckets map[string]PropertyAccessBucket, key string, example PropertyAccessExample, maxExamples int) {
	bucket := buckets[key]
	addAuditExample(&bucket, example, maxExamples)
	buckets[key] = bucket
}

func addAuditExample(bucket *PropertyAccessBucket, example PropertyAccessExample, maxExamples int) {
	bucket.Count++
	if len(bucket.Examples) < maxExamples {
		bucket.Examples = append(bucket.Examples, example)
	}
}

func ensureAuditBuckets(buckets map[string]PropertyAccessBucket, keys ...string) {
	for _, key := range keys {
		if _, ok := buckets[key]; !ok {
			buckets[key] = PropertyAccessBucket{}
		}
	}
}

func addLanguageStats(languages map[string]PropertyLanguageStats, language string, category string, orphanStatus string, graphTruth string, ownerLinked bool) {
	stats := languages[language]
	if stats.Categories == nil {
		stats.Categories = map[string]int{}
		stats.OrphanStatus = map[string]int{}
		stats.GraphTruth = map[string]int{}
	}
	stats.PropertyNodes++
	if ownerLinked {
		stats.OwnerLinkedProperties++
	} else {
		stats.StandaloneProperties++
	}
	stats.Categories[category]++
	stats.OrphanStatus[orphanStatus]++
	stats.GraphTruth[graphTruth]++
	languages[language] = stats
}

func sortedBucketKeys(buckets map[string]PropertyAccessBucket) []string {
	keys := make([]string, 0, len(buckets))
	for key := range buckets {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func sortedLanguageKeys(languages map[string]PropertyLanguageStats) []string {
	keys := make([]string, 0, len(languages))
	for key := range languages {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func propInt(props map[string]any, key string) int {
	if props == nil {
		return 0
	}
	value, ok := props[key]
	if !ok || value == nil {
		return 0
	}
	switch x := value.(type) {
	case int:
		return x
	case int64:
		return int(x)
	case float64:
		return int(x)
	case json.Number:
		n, _ := x.Int64()
		return int(n)
	default:
		return 0
	}
}
