//go:build ladybugdb

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/tamnguyendinh/anvien/internal/graph"
	"github.com/tamnguyendinh/anvien/internal/lbugnative"
	"github.com/tamnguyendinh/anvien/internal/lbugschema"
)

var definitionLabels = []string{
	"Function", "Class", "Interface", "CodeElement", "Method", "Package", "Struct", "Enum", "Macro", "Typedef", "Union", "Namespace", "Trait", "Impl", "TypeAlias", "Const", "Static", "Variable", "Property", "Record", "Delegate", "Annotation", "Constructor", "Template", "Module",
}

var comparedFields = []string{
	"id", "label", "name", "filePath", "qualifiedName", "startLine", "startCol", "endLine", "endCol", "selectionStartLine", "selectionStartCol", "selectionEndLine", "selectionEndCol",
}

type definitionRecord map[string]*string

type mismatch struct {
	ID      string  `json:"id"`
	Field   string  `json:"field"`
	Graph   *string `json:"graph"`
	Ladybug *string `json:"ladybug"`
}

type graphAudit struct {
	NodeCount                       int                         `json:"nodeCount"`
	RelationshipCount               int                         `json:"relationshipCount"`
	DefinitionCount                 int                         `json:"definitionCount"`
	DefinitionCountsByLabel         map[string]int              `json:"definitionCountsByLabel"`
	MissingConstructCoordinateCount int                         `json:"missingConstructCoordinateCount"`
	SelectionPresentCount           int                         `json:"selectionPresentCount"`
	SelectionAbsentCount            int                         `json:"selectionAbsentCount"`
	SelectionPartialCount           int                         `json:"selectionPartialCount"`
	RealZeroStartColCount           int                         `json:"realZeroStartColCount"`
	RealZeroSelectionStartColCount  int                         `json:"realZeroSelectionStartColCount"`
	DefinesPairCount                int                         `json:"definesPairCount"`
	DefinitionWithoutSingleDefines  int                         `json:"definitionWithoutSingleDefines"`
	MissingEndpointCount            int                         `json:"missingEndpointCount"`
	MissingDefinesEndpointCount     int                         `json:"missingDefinesEndpointCount"`
	Definitions                     map[string]definitionRecord `json:"-"`
	DefinesPairs                    map[string]int              `json:"-"`
	DefinitionIDs                   map[string]struct{}         `json:"-"`
	NodeIDs                         map[string]struct{}         `json:"-"`
	DefinesTargetCounts             map[string]int              `json:"-"`
}

type ladybugAudit struct {
	DefinitionCount         int                         `json:"definitionCount"`
	DefinitionCountsByLabel map[string]int              `json:"definitionCountsByLabel"`
	Definitions             map[string]definitionRecord `json:"-"`
	DefinesPairCount        int                         `json:"definesPairCount"`
	DefinesPairs            map[string]int              `json:"-"`
	EmbeddingRows           int                         `json:"embeddingRows"`
	EmbeddingMissingLabel   int                         `json:"embeddingMissingLabel"`
	EmbeddingSamples        []map[string]string         `json:"embeddingSamples"`
}

func ptr(value string) *string { return &value }

func canonical(value any) *string {
	if value == nil {
		return nil
	}
	switch typed := value.(type) {
	case json.Number:
		return ptr(typed.String())
	case float64:
		if typed == float64(int64(typed)) {
			return ptr(strconv.FormatInt(int64(typed), 10))
		}
		return ptr(strconv.FormatFloat(typed, 'g', -1, 64))
	case string:
		trimmed := strings.TrimSpace(typed)
		if strings.EqualFold(trimmed, "NULL") {
			return nil
		}
		return ptr(typed)
	default:
		return ptr(fmt.Sprint(value))
	}
}

func canonicalOptionalCoordinate(value any) *string {
	canonicalValue := canonical(value)
	if canonicalValue != nil && *canonicalValue == "" {
		return nil
	}
	return canonicalValue
}

func prop(properties graph.NodeProperties, name string) *string {
	value, ok := properties[name]
	if !ok {
		return nil
	}
	return canonical(value)
}

func isDefinitionLabel(label string) bool {
	index := sort.SearchStrings(definitionLabelsSorted, label)
	return index < len(definitionLabelsSorted) && definitionLabelsSorted[index] == label
}

var definitionLabelsSorted = func() []string {
	values := append([]string(nil), definitionLabels...)
	sort.Strings(values)
	return values
}()

func readGraph(path string) (*graphAudit, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	decoder.UseNumber()
	first, err := decoder.Token()
	if err != nil || first != json.Delim('{') {
		return nil, fmt.Errorf("graph root is not an object: %v", err)
	}
	audit := &graphAudit{
		DefinitionCountsByLabel: map[string]int{}, Definitions: map[string]definitionRecord{}, DefinesPairs: map[string]int{}, DefinitionIDs: map[string]struct{}{}, NodeIDs: map[string]struct{}{}, DefinesTargetCounts: map[string]int{},
	}
	for decoder.More() {
		keyToken, err := decoder.Token()
		if err != nil {
			return nil, err
		}
		key := keyToken.(string)
		switch key {
		case "nodes":
			if token, err := decoder.Token(); err != nil || token != json.Delim('[') {
				return nil, fmt.Errorf("nodes is not an array: %v", err)
			}
			for decoder.More() {
				var node graph.Node
				if err := decoder.Decode(&node); err != nil {
					return nil, err
				}
				audit.NodeCount++
				audit.NodeIDs[node.ID] = struct{}{}
				label := string(node.Label)
				if !isDefinitionLabel(label) {
					continue
				}
				record := definitionRecord{"id": ptr(node.ID), "label": ptr(label), "name": prop(node.Properties, "name"), "filePath": prop(node.Properties, "filePath"), "qualifiedName": prop(node.Properties, "qualifiedName"), "startLine": prop(node.Properties, "startLine"), "startCol": prop(node.Properties, "startCol"), "endLine": prop(node.Properties, "endLine"), "endCol": prop(node.Properties, "endCol"), "selectionStartLine": prop(node.Properties, "selectionStartLine"), "selectionStartCol": prop(node.Properties, "selectionStartCol"), "selectionEndLine": prop(node.Properties, "selectionEndLine"), "selectionEndCol": prop(node.Properties, "selectionEndCol")}
				audit.Definitions[node.ID] = record
				audit.DefinitionIDs[node.ID] = struct{}{}
				audit.DefinitionCount++
				audit.DefinitionCountsByLabel[label]++
				construct := []string{"startLine", "startCol", "endLine", "endCol"}
				missingConstruct := false
				for _, field := range construct {
					if record[field] == nil {
						missingConstruct = true
					}
				}
				if missingConstruct {
					audit.MissingConstructCoordinateCount++
				}
				selection := []string{"selectionStartLine", "selectionStartCol", "selectionEndLine", "selectionEndCol"}
				present := 0
				for _, field := range selection {
					if record[field] != nil {
						present++
					}
				}
				switch present {
				case 0:
					audit.SelectionAbsentCount++
				case 4:
					audit.SelectionPresentCount++
				default:
					audit.SelectionPartialCount++
				}
				if value := record["startCol"]; value != nil && *value == "0" {
					audit.RealZeroStartColCount++
				}
				if value := record["selectionStartCol"]; value != nil && *value == "0" {
					audit.RealZeroSelectionStartColCount++
				}
			}
			if _, err := decoder.Token(); err != nil {
				return nil, err
			}
		case "relationships":
			if token, err := decoder.Token(); err != nil || token != json.Delim('[') {
				return nil, fmt.Errorf("relationships is not an array: %v", err)
			}
			for decoder.More() {
				var relationship graph.Relationship
				if err := decoder.Decode(&relationship); err != nil {
					return nil, err
				}
				audit.RelationshipCount++
				_, sourceExists := audit.NodeIDs[relationship.SourceID]
				_, targetExists := audit.NodeIDs[relationship.TargetID]
				if !sourceExists || !targetExists {
					audit.MissingEndpointCount++
				}
				if relationship.Type == graph.RelDefines {
					if _, ok := audit.DefinitionIDs[relationship.TargetID]; ok {
						pair := relationship.SourceID + "\x00" + relationship.TargetID
						audit.DefinesPairs[pair]++
						audit.DefinesTargetCounts[relationship.TargetID]++
						audit.DefinesPairCount++
						if !sourceExists || !targetExists {
							audit.MissingDefinesEndpointCount++
						}
					}
				}
			}
			if _, err := decoder.Token(); err != nil {
				return nil, err
			}
		default:
			var discard any
			if err := decoder.Decode(&discard); err != nil {
				return nil, err
			}
		}
	}
	for id := range audit.DefinitionIDs {
		if audit.DefinesTargetCounts[id] != 1 {
			audit.DefinitionWithoutSingleDefines++
		}
	}
	return audit, nil
}

func readLadybug(path string) (*ladybugAudit, error) {
	runner, err := lbugnative.OpenReadRunner(path)
	if err != nil {
		return nil, err
	}
	defer runner.Close()
	audit := &ladybugAudit{DefinitionCountsByLabel: map[string]int{}, Definitions: map[string]definitionRecord{}, DefinesPairs: map[string]int{}}
	for _, label := range definitionLabels {
		ident := lbugschema.FormatIdent(label)
		query := "MATCH (n:" + ident + ") RETURN n.id AS id, n.name AS name, n.filePath AS filePath, n.qualifiedName AS qualifiedName, n.startLine AS startLine, n.startCol AS startCol, n.endLine AS endLine, n.endCol AS endCol, n.selectionStartLine AS selectionStartLine, n.selectionStartCol AS selectionStartCol, n.selectionEndLine AS selectionEndLine, n.selectionEndCol AS selectionEndCol"
		rows, err := runner.QueryRows(query)
		if err != nil {
			return nil, fmt.Errorf("query %s definitions: %w", label, err)
		}
		for _, row := range rows {
			idValue := canonical(row["id"])
			if idValue == nil {
				return nil, fmt.Errorf("%s row has NULL id", label)
			}
			id := *idValue
			audit.Definitions[id] = definitionRecord{"id": ptr(id), "label": ptr(label), "name": canonical(row["name"]), "filePath": canonical(row["filePath"]), "qualifiedName": canonical(row["qualifiedName"]), "startLine": canonical(row["startLine"]), "startCol": canonical(row["startCol"]), "endLine": canonical(row["endLine"]), "endCol": canonical(row["endCol"]), "selectionStartLine": canonicalOptionalCoordinate(row["selectionStartLine"]), "selectionStartCol": canonicalOptionalCoordinate(row["selectionStartCol"]), "selectionEndLine": canonicalOptionalCoordinate(row["selectionEndLine"]), "selectionEndCol": canonicalOptionalCoordinate(row["selectionEndCol"])}
			audit.DefinitionCount++
			audit.DefinitionCountsByLabel[label]++
		}
		rows, err = runner.QueryRows("MATCH (a:File)-[r:CodeRelation]->(b:" + ident + ") RETURN a.id AS sourceId, b.id AS targetId, r.type AS type")
		if err != nil {
			return nil, fmt.Errorf("query %s DEFINES: %w", label, err)
		}
		for _, row := range rows {
			if value := canonical(row["type"]); value == nil || *value != "DEFINES" {
				continue
			}
			source := canonical(row["sourceId"])
			target := canonical(row["targetId"])
			if source == nil || target == nil {
				continue
			}
			pair := *source + "\x00" + *target
			audit.DefinesPairs[pair]++
			audit.DefinesPairCount++
		}
	}
	embeddings, err := runner.QueryRows("MATCH (e:CodeEmbedding) RETURN e.nodeId AS nodeId, e.label AS label")
	if err != nil {
		return nil, fmt.Errorf("query CodeEmbedding labels: %w", err)
	}
	for _, row := range embeddings {
		audit.EmbeddingRows++
		nodeID := canonical(row["nodeId"])
		label := canonical(row["label"])
		if label == nil || strings.TrimSpace(*label) == "" {
			audit.EmbeddingMissingLabel++
		}
		if len(audit.EmbeddingSamples) < 20 {
			sample := map[string]string{}
			if nodeID != nil {
				sample["nodeId"] = *nodeID
			}
			if label != nil {
				sample["label"] = *label
			}
			audit.EmbeddingSamples = append(audit.EmbeddingSamples, sample)
		}
	}
	return audit, nil
}

func equalPtr(left, right *string) bool {
	if left == nil || right == nil {
		return left == nil && right == nil
	}
	return *left == *right
}

func hashFile(path string) (string, int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", 0, err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return "", 0, err
	}
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", 0, err
	}
	return strings.ToUpper(hex.EncodeToString(hash.Sum(nil))), info.Size(), nil
}

func main() {
	graphPath := flag.String("graph", "", "normal Graph JSON artifact")
	ladybugPath := flag.String("ladybug", "", "normal Ladybug artifact")
	outRoot := flag.String("out", "", "durable evidence directory")
	flag.Parse()
	if *graphPath == "" || *ladybugPath == "" || *outRoot == "" {
		fmt.Fprintln(os.Stderr, "-graph, -ladybug, and -out are required")
		os.Exit(2)
	}
	start := time.Now().UTC()
	graphResult, err := readGraph(*graphPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	ladybugResult, err := readLadybug(*ladybugPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	missing := 0
	extra := 0
	mismatchCount := 0
	mismatchByField := map[string]int{}
	samples := []mismatch{}
	for id, graphRecord := range graphResult.Definitions {
		ladybugRecord, ok := ladybugResult.Definitions[id]
		if !ok {
			missing++
			if len(samples) < 100 {
				samples = append(samples, mismatch{ID: id, Field: "record", Graph: ptr("present"), Ladybug: nil})
			}
			continue
		}
		for _, field := range comparedFields {
			if !equalPtr(graphRecord[field], ladybugRecord[field]) {
				mismatchCount++
				mismatchByField[field]++
				if len(samples) < 100 {
					samples = append(samples, mismatch{ID: id, Field: field, Graph: graphRecord[field], Ladybug: ladybugRecord[field]})
				}
			}
		}
	}
	for id := range ladybugResult.Definitions {
		if _, ok := graphResult.Definitions[id]; !ok {
			extra++
			if len(samples) < 100 {
				samples = append(samples, mismatch{ID: id, Field: "record", Graph: nil, Ladybug: ptr("present")})
			}
		}
	}
	missingPairs := 0
	extraPairs := 0
	pairCountMismatch := 0
	pairSamples := []map[string]any{}
	for pair, count := range graphResult.DefinesPairs {
		other := ladybugResult.DefinesPairs[pair]
		if other == 0 {
			missingPairs += count
		}
		if other != count {
			pairCountMismatch++
			if len(pairSamples) < 100 {
				parts := strings.Split(pair, "\x00")
				pairSamples = append(pairSamples, map[string]any{"sourceId": parts[0], "targetId": parts[1], "graphCount": count, "ladybugCount": other})
			}
		}
	}
	for pair, count := range ladybugResult.DefinesPairs {
		if graphResult.DefinesPairs[pair] == 0 {
			extraPairs += count
		}
	}
	graphHash, graphBytes, err := hashFile(*graphPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	ladybugHash, ladybugBytes, err := hashFile(*ladybugPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	overall := "PASS"
	if missing != 0 || extra != 0 || mismatchCount != 0 || missingPairs != 0 || extraPairs != 0 || pairCountMismatch != 0 || graphResult.MissingConstructCoordinateCount != 0 || graphResult.SelectionPartialCount != 0 || graphResult.MissingEndpointCount != 0 || graphResult.MissingDefinesEndpointCount != 0 || graphResult.DefinitionWithoutSingleDefines != 0 || ladybugResult.EmbeddingMissingLabel != 0 {
		overall = "FAIL"
	}
	evidence := map[string]any{
		"schema": "child02-p2e-field-parity-v1", "generatedUtc": time.Now().UTC().Format(time.RFC3339Nano), "startedUtc": start.Format(time.RFC3339Nano), "durationSeconds": time.Since(start).Seconds(), "overall": overall,
		"artifacts": map[string]any{"graphJson": map[string]any{"path": *graphPath, "sha256": graphHash, "bytes": graphBytes}, "ladybug": map[string]any{"path": *ladybugPath, "sha256": ladybugHash, "bytes": ladybugBytes}},
		"contract":  map[string]any{"definitionLabels": definitionLabels, "comparedFields": comparedFields, "recordKey": "opaque Definition id", "selectionRange": "all four present or all four absent/NULL", "zeroCoordinate": "retained as data", "definesKey": "exact sourceId+targetId pair", "embeddingLabel": "explicit persisted CodeEmbedding.label; current normal artifact may contain zero embedding rows"},
		"graphJson": graphResult, "ladybug": ladybugResult,
		"fieldParity":   map[string]any{"missingLadybugRecords": missing, "extraLadybugRecords": extra, "fieldMismatchCount": mismatchCount, "mismatchByField": mismatchByField, "samples": samples},
		"definesParity": map[string]any{"missingLadybugPairs": missingPairs, "extraLadybugPairs": extraPairs, "pairCountMismatch": pairCountMismatch, "samples": pairSamples},
	}
	if err := os.MkdirAll(*outRoot, 0o755); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	timestamp := time.Now().Format("060102_150405")
	jsonPath := filepath.Join(*outRoot, "qa_child02_p2e_parity_"+timestamp+".json")
	markdownPath := filepath.Join(*outRoot, "qa_child02_p2e_parity_"+timestamp+".md")
	raw, _ := json.MarshalIndent(evidence, "", "  ")
	if err := os.WriteFile(jsonPath, append(raw, '\n'), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	markdown := fmt.Sprintf("# Child 02 P2-E Graph JSON/Ladybug parity\n\n- Overall: **%s**\n- Graph JSON Definitions: `%d`\n- Ladybug Definitions: `%d`\n- Missing / extra / field mismatches: `%d / %d / %d`\n- Selection present / absent / partial: `%d / %d / %d`\n- Real zero startCol / selectionStartCol: `%d / %d`\n- Graph JSON / Ladybug DEFINES pairs: `%d / %d`\n- Missing / extra endpoint pairs: `%d / %d`\n- Missing graph endpoints / missing DEFINES endpoints: `%d / %d`\n- CodeEmbedding rows / missing explicit labels: `%d / %d`\n\nThe paired JSON contains per-label counts, compared-field counts, hashes, mismatch samples, and exact endpoint-pair samples.\n", overall, graphResult.DefinitionCount, ladybugResult.DefinitionCount, missing, extra, mismatchCount, graphResult.SelectionPresentCount, graphResult.SelectionAbsentCount, graphResult.SelectionPartialCount, graphResult.RealZeroStartColCount, graphResult.RealZeroSelectionStartColCount, graphResult.DefinesPairCount, ladybugResult.DefinesPairCount, missingPairs, extraPairs, graphResult.MissingEndpointCount, graphResult.MissingDefinesEndpointCount, ladybugResult.EmbeddingRows, ladybugResult.EmbeddingMissingLabel)
	if err := os.WriteFile(markdownPath, []byte(markdown), 0o644); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	result := map[string]any{"overall": overall, "json": jsonPath, "markdown": markdownPath}
	resultRaw, _ := json.Marshal(result)
	fmt.Println(string(resultRaw))
	if overall != "PASS" {
		os.Exit(1)
	}
}
