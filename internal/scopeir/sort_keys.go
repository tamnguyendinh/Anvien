package scopeir

func deref(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func compareScope(left ScopeFact, right ScopeFact) int {
	if value := compareString(left.FilePath, right.FilePath); value != 0 {
		return value
	}
	if value := compareRange(left.Range, right.Range); value != 0 {
		return value
	}
	if value := compareString(string(left.Kind), string(right.Kind)); value != 0 {
		return value
	}
	return compareString(left.ID, right.ID)
}

func compareDefinition(left DefinitionFact, right DefinitionFact) int {
	if value := compareString(left.FilePath, right.FilePath); value != 0 {
		return value
	}
	if value := compareRange(left.Range, right.Range); value != 0 {
		return value
	}
	if value := compareString(string(left.Label), string(right.Label)); value != 0 {
		return value
	}
	if value := compareString(left.Name, right.Name); value != 0 {
		return value
	}
	return compareString(left.ID, right.ID)
}

func compareImport(left ImportFact, right ImportFact) int {
	if value := compareString(left.FilePath, right.FilePath); value != 0 {
		return value
	}
	if value := compareString(string(left.Kind), string(right.Kind)); value != 0 {
		return value
	}
	if value := compareString(left.LocalName, right.LocalName); value != 0 {
		return value
	}
	if value := compareString(left.ImportedName, right.ImportedName); value != 0 {
		return value
	}
	if value := compareString(deref(left.TargetRaw), deref(right.TargetRaw)); value != 0 {
		return value
	}
	return compareString(left.ID, right.ID)
}

func compareExport(left ExportFact, right ExportFact) int {
	if value := compareString(left.FilePath, right.FilePath); value != 0 {
		return value
	}
	if value := compareRange(left.Range, right.Range); value != 0 {
		return value
	}
	if value := compareOptionalRange(left.SelectionRange, right.SelectionRange); value != 0 {
		return value
	}
	if value := compareString(string(left.Kind), string(right.Kind)); value != 0 {
		return value
	}
	if value := compareString(left.ExportedName, right.ExportedName); value != 0 {
		return value
	}
	if value := compareString(left.LocalName, right.LocalName); value != 0 {
		return value
	}
	if value := compareString(left.LocalDefID, right.LocalDefID); value != 0 {
		return value
	}
	switch {
	case left.TargetRaw == nil && right.TargetRaw != nil:
		return -1
	case left.TargetRaw != nil && right.TargetRaw == nil:
		return 1
	case left.TargetRaw != nil && right.TargetRaw != nil:
		if value := compareString(*left.TargetRaw, *right.TargetRaw); value != 0 {
			return value
		}
	}
	if value := compareString(left.TargetExportedName, right.TargetExportedName); value != 0 {
		return value
	}
	if value := compareBool(left.TypeOnly, right.TypeOnly); value != 0 {
		return value
	}
	if value := compareExportMeanings(left.Meanings, right.Meanings); value != 0 {
		return value
	}
	if value := compareExportProvenance(left.Provenance, right.Provenance); value != 0 {
		return value
	}
	return compareString(left.FileHash, right.FileHash)
}

func compareExportDiagnostic(left ExportDiagnosticFact, right ExportDiagnosticFact) int {
	if value := compareString(left.FilePath, right.FilePath); value != 0 {
		return value
	}
	if value := compareRange(left.Range, right.Range); value != 0 {
		return value
	}
	if value := compareString(string(left.Code), string(right.Code)); value != 0 {
		return value
	}
	if value := compareString(left.NodeKind, right.NodeKind); value != 0 {
		return value
	}
	if value := compareExportProvenance(left.Provenance, right.Provenance); value != 0 {
		return value
	}
	if value := compareString(left.Reason, right.Reason); value != 0 {
		return value
	}
	return compareString(left.FileHash, right.FileHash)
}

func compareExportMeanings(left []ExportMeaning, right []ExportMeaning) int {
	switch {
	case left == nil && right != nil:
		return -1
	case left != nil && right == nil:
		return 1
	}
	limit := len(left)
	if len(right) < limit {
		limit = len(right)
	}
	for index := 0; index < limit; index++ {
		if value := compareString(string(left[index]), string(right[index])); value != 0 {
			return value
		}
	}
	return compareInt(len(left), len(right))
}

func compareExportProvenance(left ExportProvenance, right ExportProvenance) int {
	if value := compareRange(left.StatementRange, right.StatementRange); value != 0 {
		return value
	}
	return compareString(left.SiteKind, right.SiteKind)
}

func compareCall(left CallSiteFact, right CallSiteFact) int {
	if value := compareString(left.FilePath, right.FilePath); value != 0 {
		return value
	}
	if value := compareRange(left.Range, right.Range); value != 0 {
		return value
	}
	if value := compareString(left.InScope, right.InScope); value != 0 {
		return value
	}
	return compareString(left.Name, right.Name)
}

func compareAccess(left AccessFact, right AccessFact) int {
	if value := compareString(left.FilePath, right.FilePath); value != 0 {
		return value
	}
	if value := compareRange(left.Range, right.Range); value != 0 {
		return value
	}
	if value := compareString(left.InScope, right.InScope); value != 0 {
		return value
	}
	if value := compareString(string(left.Kind), string(right.Kind)); value != 0 {
		return value
	}
	return compareString(left.Name, right.Name)
}

func compareHeritage(left HeritageFact, right HeritageFact) int {
	if value := compareString(left.FilePath, right.FilePath); value != 0 {
		return value
	}
	if value := compareRange(left.Range, right.Range); value != 0 {
		return value
	}
	if value := compareString(left.InScope, right.InScope); value != 0 {
		return value
	}
	if value := compareString(string(left.Kind), string(right.Kind)); value != 0 {
		return value
	}
	return compareString(left.Name, right.Name)
}

func compareTypeAnnotation(left TypeAnnotationFact, right TypeAnnotationFact) int {
	if value := compareString(left.FilePath, right.FilePath); value != 0 {
		return value
	}
	if value := compareRange(left.Range, right.Range); value != 0 {
		return value
	}
	if value := compareString(left.InScope, right.InScope); value != 0 {
		return value
	}
	if value := compareString(left.Name, right.Name); value != 0 {
		return value
	}
	return compareString(left.Type.RawName, right.Type.RawName)
}

func compareReturnType(left ReturnTypeFact, right ReturnTypeFact) int {
	if value := compareString(left.FilePath, right.FilePath); value != 0 {
		return value
	}
	if value := compareRange(left.Range, right.Range); value != 0 {
		return value
	}
	if value := compareString(left.DefID, right.DefID); value != 0 {
		return value
	}
	return compareString(left.Type.RawName, right.Type.RawName)
}

func compareFramework(left FrameworkFact, right FrameworkFact) int {
	if value := compareString(left.FilePath, right.FilePath); value != 0 {
		return value
	}
	if value := compareRange(left.Range, right.Range); value != 0 {
		return value
	}
	if value := compareString(left.DefID, right.DefID); value != 0 {
		return value
	}
	if value := compareString(left.Framework, right.Framework); value != 0 {
		return value
	}
	return compareString(left.Reason, right.Reason)
}

func compareDomain(left DomainFact, right DomainFact) int {
	if value := compareString(left.FilePath, right.FilePath); value != 0 {
		return value
	}
	if value := compareRange(left.Range, right.Range); value != 0 {
		return value
	}
	if value := compareString(left.DefID, right.DefID); value != 0 {
		return value
	}
	if value := compareString(left.Domain, right.Domain); value != 0 {
		return value
	}
	return compareString(left.Role, right.Role)
}

func compareBinding(left BindingFact, right BindingFact) int {
	if value := compareString(left.Name, right.Name); value != 0 {
		return value
	}
	if value := compareString(left.DefID, right.DefID); value != 0 {
		return value
	}
	if value := compareString(string(left.Origin), string(right.Origin)); value != 0 {
		return value
	}
	return compareString(left.ViaID, right.ViaID)
}

func compareBindingLeaf(left BindingLeafFact, right BindingLeafFact) int {
	if value := compareString(left.FilePath, right.FilePath); value != 0 {
		return value
	}
	if value := compareRange(left.Range, right.Range); value != 0 {
		return value
	}
	if value := compareString(left.Name, right.Name); value != 0 {
		return value
	}
	if value := compareBindingPath(left.Path, right.Path); value != 0 {
		return value
	}
	if value := compareBool(left.Rest, right.Rest); value != 0 {
		return value
	}
	if value := compareBool(left.Default, right.Default); value != 0 {
		return value
	}
	if value := compareBindingProvenance(left.Provenance, right.Provenance); value != 0 {
		return value
	}
	if value := compareOptionalRange(left.SelectionRange, right.SelectionRange); value != 0 {
		return value
	}
	return compareString(left.FileHash, right.FileHash)
}

func compareExtractionDiagnostic(left ExtractionDiagnosticFact, right ExtractionDiagnosticFact) int {
	if value := compareString(left.FilePath, right.FilePath); value != 0 {
		return value
	}
	if value := compareRange(left.Range, right.Range); value != 0 {
		return value
	}
	if value := compareString(string(left.Code), string(right.Code)); value != 0 {
		return value
	}
	if value := compareString(left.NodeKind, right.NodeKind); value != 0 {
		return value
	}
	if value := compareBindingPath(left.Path, right.Path); value != 0 {
		return value
	}
	if value := compareBindingProvenance(left.Provenance, right.Provenance); value != 0 {
		return value
	}
	if value := compareString(left.Reason, right.Reason); value != 0 {
		return value
	}
	return compareString(left.FileHash, right.FileHash)
}

func compareBindingPath(left []BindingPathSegment, right []BindingPathSegment) int {
	limit := len(left)
	if len(right) < limit {
		limit = len(right)
	}
	for index := 0; index < limit; index++ {
		if value := compareString(string(left[index].Kind), string(right[index].Kind)); value != 0 {
			return value
		}
		if value := compareOptionalInt(left[index].ArrayIndex, right[index].ArrayIndex); value != 0 {
			return value
		}
		if value := compareString(left[index].PropertyName, right[index].PropertyName); value != 0 {
			return value
		}
		if value := compareString(left[index].ComputedExpression, right[index].ComputedExpression); value != 0 {
			return value
		}
		if value := compareRange(left[index].SourceRange, right[index].SourceRange); value != 0 {
			return value
		}
	}
	return compareInt(len(left), len(right))
}

func compareBindingProvenance(left BindingPatternProvenance, right BindingPatternProvenance) int {
	if value := compareString(string(left.Context), string(right.Context)); value != 0 {
		return value
	}
	if value := compareRange(left.ConstructRange, right.ConstructRange); value != 0 {
		return value
	}
	if value := compareRange(left.PatternRange, right.PatternRange); value != 0 {
		return value
	}
	return compareString(left.PatternKind, right.PatternKind)
}

func compareOptionalRange(left *Range, right *Range) int {
	switch {
	case left == nil && right == nil:
		return 0
	case left == nil:
		return -1
	case right == nil:
		return 1
	default:
		return compareRange(*left, *right)
	}
}

func compareOptionalInt(left *int, right *int) int {
	switch {
	case left == nil && right == nil:
		return 0
	case left == nil:
		return -1
	case right == nil:
		return 1
	default:
		return compareInt(*left, *right)
	}
}

func compareBool(left bool, right bool) int {
	switch {
	case left == right:
		return 0
	case !left:
		return -1
	default:
		return 1
	}
}

func compareTypeBinding(left TypeBindingFact, right TypeBindingFact) int {
	if value := compareString(left.Name, right.Name); value != 0 {
		return value
	}
	if value := compareString(left.Type.RawName, right.Type.RawName); value != 0 {
		return value
	}
	if value := compareString(left.Type.DeclaredAtScope, right.Type.DeclaredAtScope); value != 0 {
		return value
	}
	return compareString(string(left.Type.Source), string(right.Type.Source))
}

func compareRange(left Range, right Range) int {
	if value := compareInt(left.StartLine, right.StartLine); value != 0 {
		return value
	}
	if value := compareInt(left.StartCol, right.StartCol); value != 0 {
		return value
	}
	if value := compareInt(left.EndLine, right.EndLine); value != 0 {
		return value
	}
	return compareInt(left.EndCol, right.EndCol)
}

func compareString(left string, right string) int {
	switch {
	case left < right:
		return -1
	case left > right:
		return 1
	default:
		return 0
	}
}

func compareInt(left int, right int) int {
	switch {
	case left < right:
		return -1
	case left > right:
		return 1
	default:
		return 0
	}
}
