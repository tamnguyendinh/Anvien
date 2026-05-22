package lbugload

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/tamnguyendinh/avmatrix-go/internal/lbugschema"
)

const copyCSVOptions = `(HEADER=true, ESCAPE='"', DELIM=',', QUOTE='"', PARALLEL=false, auto_detect=false)`

type RelationshipCSVRow struct {
	FromID           string
	ToID             string
	Type             string
	Confidence       string
	Reason           string
	Step             string
	ResolutionSource string
	Evidence         string
	FileHash         string
	SourceSiteID     string
	SourceSiteIDs    string
	SourceSiteCount  string
	SourceSiteStatus string
	ProofKind        string
	TargetRole       string
	TargetText       string
	FilePath         string
	StartLine        string
	StartCol         string
	EndLine          string
	EndCol           string
}

func NodeCopyQuery(table string, csvPath string) (string, error) {
	columns, ok := nodeColumns(table)
	if !ok {
		return "", fmt.Errorf("unsupported node table %q", table)
	}
	return fmt.Sprintf(
		`COPY %s(%s) FROM "%s" %s`,
		lbugschema.FormatIdent(table),
		strings.Join(columns, ", "),
		NormalizeCopyPath(csvPath),
		copyCSVOptions,
	), nil
}

func RelationshipCopyQuery(fromLabel string, toLabel string, csvPath string) string {
	return fmt.Sprintf(
		`COPY %s FROM "%s" (from="%s", to="%s", HEADER=true, ESCAPE='"', DELIM=',', QUOTE='"', PARALLEL=false, auto_detect=false)`,
		lbugschema.RelTableName,
		NormalizeCopyPath(csvPath),
		fromLabel,
		toLabel,
	)
}

func RetryCopyQuery(query string) string {
	return strings.Replace(query, "auto_detect=false)", "auto_detect=false, IGNORE_ERRORS=true)", 1)
}

func FallbackRelationshipInsertQuery(row RelationshipCSVRow, fromLabel string, toLabel string) string {
	confidence := parseFloatString(row.Confidence, 1)
	step := parseIntString(row.Step, 0)
	sourceSiteCount := parseIntString(row.SourceSiteCount, 0)
	startLine := parseIntString(row.StartLine, 0)
	startCol := parseIntString(row.StartCol, 0)
	endLine := parseIntString(row.EndLine, 0)
	endCol := parseIntString(row.EndCol, 0)
	return fmt.Sprintf(
		`MATCH (a:%s {id: %s}), (b:%s {id: %s}) CREATE (a)-[:%s {type: %s, confidence: %s, reason: %s, step: %d, resolutionSource: %s, evidence: %s, fileHash: %s, sourceSiteId: %s, sourceSiteIds: %s, sourceSiteCount: %d, sourceSiteStatus: %s, proofKind: %s, targetRole: %s, targetText: %s, filePath: %s, startLine: %d, startCol: %d, endLine: %d, endCol: %d}]->(b)`,
		lbugschema.FormatIdent(fromLabel),
		cypherString(row.FromID),
		lbugschema.FormatIdent(toLabel),
		cypherString(row.ToID),
		lbugschema.RelTableName,
		cypherString(row.Type),
		strconv.FormatFloat(confidence, 'f', -1, 64),
		cypherString(row.Reason),
		step,
		cypherString(row.ResolutionSource),
		cypherString(row.Evidence),
		cypherString(row.FileHash),
		cypherString(row.SourceSiteID),
		cypherString(row.SourceSiteIDs),
		sourceSiteCount,
		cypherString(row.SourceSiteStatus),
		cypherString(row.ProofKind),
		cypherString(row.TargetRole),
		cypherString(row.TargetText),
		cypherString(row.FilePath),
		startLine,
		startCol,
		endLine,
		endCol,
	)
}

func NormalizeCopyPath(path string) string {
	return strings.ReplaceAll(path, `\`, `/`)
}

func ReadRelationshipCSVRows(csvPath string) ([]RelationshipCSVRow, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, fmt.Errorf("open relationship csv %s: %w", csvPath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	if _, err := reader.Read(); err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, fmt.Errorf("read relationship csv header %s: %w", csvPath, err)
	}

	var rows []RelationshipCSVRow
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("read relationship csv row %s: %w", csvPath, err)
		}
		row, ok := relationshipRowFromRecord(record)
		if !ok {
			continue
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func relationshipRowFromRecord(record []string) (RelationshipCSVRow, bool) {
	if len(record) < len(relationshipColumns) {
		return RelationshipCSVRow{}, false
	}
	return RelationshipCSVRow{
		FromID:           record[0],
		ToID:             record[1],
		Type:             record[2],
		Confidence:       record[3],
		Reason:           record[4],
		Step:             record[5],
		ResolutionSource: record[6],
		Evidence:         record[7],
		FileHash:         record[8],
		SourceSiteID:     record[9],
		SourceSiteIDs:    record[10],
		SourceSiteCount:  record[11],
		SourceSiteStatus: record[12],
		ProofKind:        record[13],
		TargetRole:       record[14],
		TargetText:       record[15],
		FilePath:         record[16],
		StartLine:        record[17],
		StartCol:         record[18],
		EndLine:          record[19],
		EndCol:           record[20],
	}, true
}

func cypherString(value string) string {
	value = strings.ReplaceAll(value, `\`, `\\`)
	value = strings.ReplaceAll(value, `'`, `''`)
	value = strings.ReplaceAll(value, "\n", `\n`)
	value = strings.ReplaceAll(value, "\r", `\r`)
	return "'" + value + "'"
}

func parseFloatString(value string, fallback float64) float64 {
	parsed, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fallback
	}
	return parsed
}

func parseIntString(value string, fallback int) int {
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}
