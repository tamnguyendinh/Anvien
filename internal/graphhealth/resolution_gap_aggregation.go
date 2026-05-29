package graphhealth

import (
	"sort"
	"strings"

	"github.com/tamnguyendinh/anvien/internal/graph"
)

const resolutionGapAggregateKindPrefix = "ResolutionGapAggregate:"

type ResolutionGapAggregationOptions struct {
	MaxSamples int
}

type ResolutionGapAggregate struct {
	ID                   string         `json:"id"`
	BucketKey            string         `json:"bucketKey"`
	SourceNodeID         string         `json:"sourceNodeId"`
	SourceNodeLabel      string         `json:"sourceNodeLabel,omitempty"`
	FactFamily           string         `json:"factFamily"`
	GapKind              string         `json:"gapKind"`
	TargetText           string         `json:"targetText"`
	TargetRole           string         `json:"targetRole"`
	SourceSiteStatus     string         `json:"sourceSiteStatus,omitempty"`
	ProofKind            string         `json:"proofKind,omitempty"`
	Classification       string         `json:"classification,omitempty"`
	Actionability        string         `json:"actionability,omitempty"`
	ResolutionSource     string         `json:"resolutionSource,omitempty"`
	Note                 string         `json:"note,omitempty"`
	InputCount           int            `json:"inputCount"`
	OccurrenceCount      int            `json:"occurrenceCount"`
	SourceSiteCount      int            `json:"sourceSiteCount"`
	SourceSiteIDs        []string       `json:"sourceSiteIds,omitempty"`
	AppLayerCounts       map[string]int `json:"appLayerCounts,omitempty"`
	FunctionalAreaCounts map[string]int `json:"functionalAreaCounts,omitempty"`
	FilePathCounts       map[string]int `json:"filePathCounts,omitempty"`
	Samples              []GapSample    `json:"samples,omitempty"`

	sourceSiteSet map[string]struct{}
}

type GapSample struct {
	SourceSiteID         string `json:"sourceSiteId,omitempty"`
	SourceNodeID         string `json:"sourceNodeId,omitempty"`
	FactFamily           string `json:"factFamily,omitempty"`
	TargetText           string `json:"targetText,omitempty"`
	TargetRole           string `json:"targetRole,omitempty"`
	SourceAppLayer       string `json:"sourceAppLayer,omitempty"`
	SourceFunctionalArea string `json:"sourceFunctionalArea,omitempty"`
	FilePath             string `json:"filePath,omitempty"`
	FileHash             string `json:"fileHash,omitempty"`
	StartLine            int    `json:"startLine,omitempty"`
	StartCol             int    `json:"startCol,omitempty"`
	EndLine              int    `json:"endLine,omitempty"`
	EndCol               int    `json:"endCol,omitempty"`
	Count                int    `json:"count"`
	Note                 string `json:"note,omitempty"`
}

func SourceBackedResolutionGapAggregates(g *graph.Graph, options ResolutionGapAggregationOptions) []ResolutionGapAggregate {
	return ResolutionGapAggregates(SourceBackedResolutionGapInputs(g), options)
}

func ResolutionGapAggregates(inputs []ResolutionGapInput, options ResolutionGapAggregationOptions) []ResolutionGapAggregate {
	if len(inputs) == 0 {
		return nil
	}
	byKey := make(map[string]*ResolutionGapAggregate, len(inputs))
	for _, input := range inputs {
		key := resolutionGapAggregateKey(input)
		aggregate := byKey[key]
		if aggregate == nil {
			aggregate = newResolutionGapAggregate(input, key)
			byKey[key] = aggregate
		}
		aggregate.add(input, options)
	}
	out := make([]ResolutionGapAggregate, 0, len(byKey))
	for _, aggregate := range byKey {
		sort.Strings(aggregate.SourceSiteIDs)
		aggregate.SourceSiteCount = len(aggregate.SourceSiteIDs)
		aggregate.sourceSiteSet = nil
		out = append(out, *aggregate)
	}
	sort.SliceStable(out, func(i int, j int) bool {
		left := out[i]
		right := out[j]
		if left.SourceNodeID != right.SourceNodeID {
			return left.SourceNodeID < right.SourceNodeID
		}
		if left.FactFamily != right.FactFamily {
			return left.FactFamily < right.FactFamily
		}
		if left.TargetText != right.TargetText {
			return left.TargetText < right.TargetText
		}
		if left.SourceSiteStatus != right.SourceSiteStatus {
			return left.SourceSiteStatus < right.SourceSiteStatus
		}
		return left.ID < right.ID
	})
	return out
}

func newResolutionGapAggregate(input ResolutionGapInput, key string) *ResolutionGapAggregate {
	targetRole := input.InferredTargetRole()
	return &ResolutionGapAggregate{
		ID:                   resolutionGapAggregateKindPrefix + key,
		BucketKey:            key,
		SourceNodeID:         strings.TrimSpace(input.SourceNodeID),
		SourceNodeLabel:      strings.TrimSpace(input.SourceNodeLabel),
		FactFamily:           strings.TrimSpace(input.FactFamily),
		GapKind:              input.GapKind(),
		TargetText:           strings.TrimSpace(input.TargetText),
		TargetRole:           targetRole,
		SourceSiteStatus:     strings.TrimSpace(input.SourceSiteStatus),
		ProofKind:            strings.TrimSpace(input.ProofKind),
		Classification:       strings.TrimSpace(input.Classification),
		Actionability:        strings.TrimSpace(input.Actionability),
		ResolutionSource:     strings.TrimSpace(input.ResolutionSource),
		Note:                 strings.TrimSpace(input.Note),
		AppLayerCounts:       map[string]int{},
		FunctionalAreaCounts: map[string]int{},
		FilePathCounts:       map[string]int{},
		sourceSiteSet:        map[string]struct{}{},
	}
}

func (aggregate *ResolutionGapAggregate) add(input ResolutionGapInput, options ResolutionGapAggregationOptions) {
	count := input.Count
	if count <= 0 {
		count = 1
	}
	aggregate.InputCount++
	aggregate.OccurrenceCount += count
	addCount(aggregate.AppLayerCounts, strings.TrimSpace(input.SourceAppLayer), count)
	addCount(aggregate.FunctionalAreaCounts, strings.TrimSpace(input.SourceFunctionalArea), count)
	addCount(aggregate.FilePathCounts, strings.TrimSpace(input.FilePath), count)
	if sourceSiteID := strings.TrimSpace(input.SourceSiteID); sourceSiteID != "" {
		if _, ok := aggregate.sourceSiteSet[sourceSiteID]; !ok {
			aggregate.sourceSiteSet[sourceSiteID] = struct{}{}
			aggregate.SourceSiteIDs = append(aggregate.SourceSiteIDs, sourceSiteID)
		}
	}
	if options.MaxSamples <= 0 || len(aggregate.Samples) < options.MaxSamples {
		aggregate.Samples = append(aggregate.Samples, gapSampleFromInput(input, count))
	}
}

func addCount(counts map[string]int, key string, count int) {
	if key == "" {
		key = "unknown"
	}
	counts[key] += count
}

func gapSampleFromInput(input ResolutionGapInput, count int) GapSample {
	return GapSample{
		SourceSiteID:         strings.TrimSpace(input.SourceSiteID),
		SourceNodeID:         strings.TrimSpace(input.SourceNodeID),
		FactFamily:           strings.TrimSpace(input.FactFamily),
		TargetText:           strings.TrimSpace(input.TargetText),
		TargetRole:           input.InferredTargetRole(),
		SourceAppLayer:       strings.TrimSpace(input.SourceAppLayer),
		SourceFunctionalArea: strings.TrimSpace(input.SourceFunctionalArea),
		FilePath:             strings.TrimSpace(input.FilePath),
		FileHash:             strings.TrimSpace(input.FileHash),
		StartLine:            input.StartLine,
		StartCol:             input.StartCol,
		EndLine:              input.EndLine,
		EndCol:               input.EndCol,
		Count:                count,
		Note:                 strings.TrimSpace(input.Note),
	}
}

func resolutionGapAggregateKey(input ResolutionGapInput) string {
	return strings.Join(nonEmptyParts(
		strings.TrimSpace(input.SourceNodeID),
		strings.TrimSpace(input.FactFamily),
		strings.TrimSpace(input.TargetText),
		input.InferredTargetRole(),
		strings.TrimSpace(input.SourceSiteStatus),
		strings.TrimSpace(input.ProofKind),
		strings.TrimSpace(input.Classification),
		strings.TrimSpace(input.Actionability),
		strings.TrimSpace(input.ResolutionSource),
		strings.TrimSpace(input.FilePath),
		strings.TrimSpace(input.Note),
	), "|")
}

func nonEmptyParts(parts ...string) []string {
	clean := make([]string, 0, len(parts))
	for _, part := range parts {
		if part == "" {
			clean = append(clean, "_")
			continue
		}
		clean = append(clean, part)
	}
	return clean
}
