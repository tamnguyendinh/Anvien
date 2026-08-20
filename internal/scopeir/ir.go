package scopeir

import (
	"bytes"
	"encoding/json"
	"sort"

	"github.com/tamnguyendinh/anvien/internal/scanner"
)

const Version = "scopeir.v1"

type ScopeIR struct {
	Version               string                     `json:"version"`
	FilePath              string                     `json:"filePath"`
	FileHash              string                     `json:"fileHash,omitempty"`
	Language              scanner.Language           `json:"language,omitempty"`
	ModuleScope           string                     `json:"moduleScope"`
	Scopes                []ScopeFact                `json:"scopes"`
	Definitions           []DefinitionFact           `json:"definitions"`
	BindingLeaves         []BindingLeafFact          `json:"bindingLeaves,omitempty"`
	ExtractionDiagnostics []ExtractionDiagnosticFact `json:"extractionDiagnostics,omitempty"`
	Exports               []ExportFact               `json:"exports,omitempty"`
	ExportDiagnostics     []ExportDiagnosticFact     `json:"exportDiagnostics,omitempty"`
	Imports               []ImportFact               `json:"imports,omitempty"`
	Calls                 []CallSiteFact             `json:"calls,omitempty"`
	Accesses              []AccessFact               `json:"accesses,omitempty"`
	Heritage              []HeritageFact             `json:"heritage,omitempty"`
	TypeAnnotations       []TypeAnnotationFact       `json:"typeAnnotations,omitempty"`
	ReturnTypes           []ReturnTypeFact           `json:"returnTypes,omitempty"`
	Frameworks            []FrameworkFact            `json:"frameworks,omitempty"`
	Domains               []DomainFact               `json:"domains,omitempty"`
}

func (ir ScopeIR) Normalized() ScopeIR {
	out := ir
	if out.Version == "" {
		out.Version = Version
	}
	out.Scopes = append([]ScopeFact(nil), ir.Scopes...)
	out.Definitions = append([]DefinitionFact(nil), ir.Definitions...)
	out.BindingLeaves = append([]BindingLeafFact(nil), ir.BindingLeaves...)
	out.ExtractionDiagnostics = append([]ExtractionDiagnosticFact(nil), ir.ExtractionDiagnostics...)
	out.Exports = append([]ExportFact(nil), ir.Exports...)
	out.ExportDiagnostics = append([]ExportDiagnosticFact(nil), ir.ExportDiagnostics...)
	out.Imports = append([]ImportFact(nil), ir.Imports...)
	out.Calls = append([]CallSiteFact(nil), ir.Calls...)
	out.Accesses = append([]AccessFact(nil), ir.Accesses...)
	out.Heritage = append([]HeritageFact(nil), ir.Heritage...)
	out.TypeAnnotations = append([]TypeAnnotationFact(nil), ir.TypeAnnotations...)
	out.ReturnTypes = append([]ReturnTypeFact(nil), ir.ReturnTypes...)
	out.Frameworks = append([]FrameworkFact(nil), ir.Frameworks...)
	out.Domains = append([]DomainFact(nil), ir.Domains...)

	for index := range out.Scopes {
		out.Scopes[index].Bindings = append([]BindingFact(nil), out.Scopes[index].Bindings...)
		out.Scopes[index].OwnedDefIDs = append([]string(nil), out.Scopes[index].OwnedDefIDs...)
		out.Scopes[index].TypeBindings = append(
			[]TypeBindingFact(nil),
			out.Scopes[index].TypeBindings...,
		)
	}
	for index := range out.Definitions {
		out.Definitions[index].ParameterTypes = append([]string(nil), out.Definitions[index].ParameterTypes...)
		out.Definitions[index].Annotations = append([]string(nil), out.Definitions[index].Annotations...)
	}
	cloneBindingCollections(&out)
	cloneExportCollections(&out)
	for index := range out.Imports {
		out.Imports[index].TransitiveVia = append([]string(nil), out.Imports[index].TransitiveVia...)
	}
	for index := range out.Calls {
		out.Calls[index].ArgTypes = append([]string(nil), out.Calls[index].ArgTypes...)
	}

	return out.NormalizeInPlace()
}

func (ir ScopeIR) NormalizeInPlace() ScopeIR {
	if ir.Version == "" {
		ir.Version = Version
	}
	for index := range ir.Scopes {
		sort.Slice(ir.Scopes[index].Bindings, func(i, j int) bool {
			return compareBinding(ir.Scopes[index].Bindings[i], ir.Scopes[index].Bindings[j]) < 0
		})
		sort.Strings(ir.Scopes[index].OwnedDefIDs)
		sort.Slice(ir.Scopes[index].TypeBindings, func(i, j int) bool {
			return compareTypeBinding(ir.Scopes[index].TypeBindings[i], ir.Scopes[index].TypeBindings[j]) < 0
		})
	}
	for index := range ir.Definitions {
		sort.Strings(ir.Definitions[index].ParameterTypes)
		sort.Strings(ir.Definitions[index].Annotations)
	}
	for index := range ir.Imports {
		sort.Strings(ir.Imports[index].TransitiveVia)
	}
	for index := range ir.Exports {
		meanings := ir.Exports[index].Meanings
		sort.Slice(meanings, func(i, j int) bool {
			return meanings[i] < meanings[j]
		})
		writeIndex := 0
		for _, meaning := range meanings {
			if writeIndex > 0 && meanings[writeIndex-1] == meaning {
				continue
			}
			meanings[writeIndex] = meaning
			writeIndex++
		}
		ir.Exports[index].Meanings = meanings[:writeIndex]
	}

	sort.Slice(ir.Scopes, func(i, j int) bool { return compareScope(ir.Scopes[i], ir.Scopes[j]) < 0 })
	sort.Slice(ir.Definitions, func(i, j int) bool {
		return compareDefinition(ir.Definitions[i], ir.Definitions[j]) < 0
	})
	sort.Slice(ir.BindingLeaves, func(i, j int) bool {
		return compareBindingLeaf(ir.BindingLeaves[i], ir.BindingLeaves[j]) < 0
	})
	sort.Slice(ir.ExtractionDiagnostics, func(i, j int) bool {
		return compareExtractionDiagnostic(ir.ExtractionDiagnostics[i], ir.ExtractionDiagnostics[j]) < 0
	})
	sort.Slice(ir.Exports, func(i, j int) bool {
		return compareExport(ir.Exports[i], ir.Exports[j]) < 0
	})
	sort.Slice(ir.ExportDiagnostics, func(i, j int) bool {
		return compareExportDiagnostic(ir.ExportDiagnostics[i], ir.ExportDiagnostics[j]) < 0
	})
	sort.Slice(ir.Imports, func(i, j int) bool { return compareImport(ir.Imports[i], ir.Imports[j]) < 0 })
	sort.Slice(ir.Calls, func(i, j int) bool { return compareCall(ir.Calls[i], ir.Calls[j]) < 0 })
	sort.Slice(ir.Accesses, func(i, j int) bool { return compareAccess(ir.Accesses[i], ir.Accesses[j]) < 0 })
	sort.Slice(ir.Heritage, func(i, j int) bool { return compareHeritage(ir.Heritage[i], ir.Heritage[j]) < 0 })
	sort.Slice(ir.TypeAnnotations, func(i, j int) bool {
		return compareTypeAnnotation(ir.TypeAnnotations[i], ir.TypeAnnotations[j]) < 0
	})
	sort.Slice(ir.ReturnTypes, func(i, j int) bool {
		return compareReturnType(ir.ReturnTypes[i], ir.ReturnTypes[j]) < 0
	})
	sort.Slice(ir.Frameworks, func(i, j int) bool {
		return compareFramework(ir.Frameworks[i], ir.Frameworks[j]) < 0
	})
	sort.Slice(ir.Domains, func(i, j int) bool { return compareDomain(ir.Domains[i], ir.Domains[j]) < 0 })

	return ir
}

func (ir ScopeIR) NormalizeOwned() ScopeIR {
	if ir.Version == "" {
		ir.Version = Version
	}
	ir.Scopes = append([]ScopeFact(nil), ir.Scopes...)
	ir.Definitions = append([]DefinitionFact(nil), ir.Definitions...)
	ir.BindingLeaves = append([]BindingLeafFact(nil), ir.BindingLeaves...)
	ir.ExtractionDiagnostics = append([]ExtractionDiagnosticFact(nil), ir.ExtractionDiagnostics...)
	ir.Exports = append([]ExportFact(nil), ir.Exports...)
	ir.ExportDiagnostics = append([]ExportDiagnosticFact(nil), ir.ExportDiagnostics...)
	ir.Imports = append([]ImportFact(nil), ir.Imports...)
	ir.Calls = append([]CallSiteFact(nil), ir.Calls...)
	ir.Accesses = append([]AccessFact(nil), ir.Accesses...)
	ir.Heritage = append([]HeritageFact(nil), ir.Heritage...)
	ir.TypeAnnotations = append([]TypeAnnotationFact(nil), ir.TypeAnnotations...)
	ir.ReturnTypes = append([]ReturnTypeFact(nil), ir.ReturnTypes...)
	ir.Frameworks = append([]FrameworkFact(nil), ir.Frameworks...)
	ir.Domains = append([]DomainFact(nil), ir.Domains...)

	cloneBindingCollections(&ir)
	cloneExportCollections(&ir)

	return ir.NormalizeInPlace()
}

func cloneExportCollections(ir *ScopeIR) {
	for index := range ir.Exports {
		ir.Exports[index].Meanings = append([]ExportMeaning(nil), ir.Exports[index].Meanings...)
		ir.Exports[index].SelectionRange = cloneRange(ir.Exports[index].SelectionRange)
		ir.Exports[index].TargetRaw = cloneStringPointer(ir.Exports[index].TargetRaw)
	}
}

func cloneBindingCollections(ir *ScopeIR) {
	for index := range ir.BindingLeaves {
		ir.BindingLeaves[index].SelectionRange = cloneRange(ir.BindingLeaves[index].SelectionRange)
		ir.BindingLeaves[index].Path = cloneBindingPath(ir.BindingLeaves[index].Path)
	}
	for index := range ir.ExtractionDiagnostics {
		ir.ExtractionDiagnostics[index].Path = cloneBindingPath(ir.ExtractionDiagnostics[index].Path)
	}
}

func cloneRange(value *Range) *Range {
	if value == nil {
		return nil
	}
	cloned := *value
	return &cloned
}

func cloneStringPointer(value *string) *string {
	if value == nil {
		return nil
	}
	cloned := *value
	return &cloned
}

func cloneBindingPath(path []BindingPathSegment) []BindingPathSegment {
	cloned := append([]BindingPathSegment(nil), path...)
	for index := range cloned {
		if cloned[index].ArrayIndex == nil {
			continue
		}
		value := *cloned[index].ArrayIndex
		cloned[index].ArrayIndex = &value
	}
	return cloned
}

func (ir ScopeIR) MarshalDeterministic() ([]byte, error) {
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(ir.Normalized()); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func Unmarshal(raw []byte) (ScopeIR, error) {
	var ir ScopeIR
	if err := json.Unmarshal(raw, &ir); err != nil {
		return ScopeIR{}, err
	}
	return ir.Normalized(), nil
}
