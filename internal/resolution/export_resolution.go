package resolution

import (
	"sort"
	"strings"

	"github.com/tamnguyendinh/anvien/internal/scopeir"
)

type exportResolutionOutcome string

const (
	exportResolutionTerminal        exportResolutionOutcome = "terminal"
	exportResolutionAmbiguity       exportResolutionOutcome = "ambiguity"
	exportResolutionCycle           exportResolutionOutcome = "cycle"
	exportResolutionMissing         exportResolutionOutcome = "missing"
	exportResolutionMeaningMismatch exportResolutionOutcome = "meaning-mismatch"
)

type exportResolutionTerminalKind string

const (
	exportResolutionDefinitionTerminal exportResolutionTerminalKind = "definition"
	exportResolutionNamespaceTerminal  exportResolutionTerminalKind = "namespace"
)

type exportResolutionHopKind string

const (
	exportResolutionExportHop exportResolutionHopKind = "export"
	exportResolutionMemberHop exportResolutionHopKind = "member"
)

// exportResolutionHop is one deterministic, source-backed step in an export
// proof. Export hops retain the accepted Child 04 fact verbatim. Member hops
// identify the repository definition selected after a namespace or exported
// owner has been resolved.
type exportResolutionHop struct {
	Kind             exportResolutionHopKind
	FilePath         string
	RequestedName    string
	Fact             scopeir.ExportFact
	TargetFile       string
	MemberOwnerDefID string
	MemberName       string
	TerminalDefID    string
}

type exportResolutionProof struct {
	Outcome       exportResolutionOutcome
	TerminalKind  exportResolutionTerminalKind
	Terminal      defRef
	NamespaceFile string
	Meanings      []scopeir.ExportMeaning
	Hops          []exportResolutionHop
	FailureFile   string
	FailureName   string
}

// exportResolutionCandidate is one semantic terminal. Proofs intentionally
// retain every deterministic path to that terminal so multiple paths to the
// same Symbol can be deduplicated without discarding provenance.
type exportResolutionCandidate struct {
	Kind          exportResolutionTerminalKind
	Terminal      defRef
	NamespaceFile string
	Meanings      []scopeir.ExportMeaning
	Proofs        []exportResolutionProof
}

// exportResolutionResult is immutable by ownership: every nested slice and
// fact is copied before it is returned. A terminal result has exactly one
// candidate; ambiguity retains every distinct candidate and selects none.
type exportResolutionResult struct {
	Outcome           exportResolutionOutcome
	Request           scopeir.ImportFact
	RequestedName     string
	RequestedMeanings []scopeir.ExportMeaning
	TargetFiles       []string
	Candidates        []exportResolutionCandidate
	Failures          []exportResolutionProof
}

func (result exportResolutionResult) definition() (defRef, bool) {
	if result.Outcome != exportResolutionTerminal || len(result.Candidates) != 1 {
		return defRef{}, false
	}
	candidate := result.Candidates[0]
	if candidate.Kind != exportResolutionDefinitionTerminal {
		return defRef{}, false
	}
	return cloneExportResolutionDefRef(candidate.Terminal), true
}

func (result exportResolutionResult) allProofs() []exportResolutionProof {
	proofs := make([]exportResolutionProof, 0, len(result.Failures)+len(result.Candidates))
	for _, candidate := range result.Candidates {
		for _, proof := range candidate.Proofs {
			proofs = append(proofs, cloneExportResolutionProof(proof))
		}
	}
	for _, failure := range result.Failures {
		proofs = append(proofs, cloneExportResolutionProof(failure))
	}
	sortExportResolutionProofs(proofs)
	return proofs
}

// resolveImportExport resolves one source-written semantic import against the
// already-built P5-B tables. desiredMeanings is the use-site lane when one is
// known; nil means the import's complete requested allowed-set.
func (w *workspace) resolveImportExport(item scopeir.ImportFact, targetFiles []string, desiredMeanings []scopeir.ExportMeaning) exportResolutionResult {
	requestedMeanings := canonicalExportMeanings(item.RequestedMeanings)
	if len(desiredMeanings) > 0 {
		requestedMeanings = intersectExportMeanings(requestedMeanings, desiredMeanings)
	}
	requestedName := strings.TrimSpace(item.ImportedName)
	result := w.resolveExportRequest(item, targetFiles, requestedName, requestedMeanings)
	if len(requestedMeanings) == 0 {
		result.Outcome = exportResolutionMeaningMismatch
		result.Failures = []exportResolutionProof{{
			Outcome:     exportResolutionMeaningMismatch,
			FailureFile: cleanPath(item.FilePath),
			FailureName: requestedName,
		}}
		result.Candidates = nil
	}
	return result
}

func (w *workspace) resolveExportRequest(item scopeir.ImportFact, targetFiles []string, requestedName string, requestedMeanings []scopeir.ExportMeaning) exportResolutionResult {
	requestedName = strings.TrimSpace(requestedName)
	requestedMeanings = canonicalExportMeanings(requestedMeanings)
	targetFiles = cloneSortedExportTablePaths(targetFiles)

	proofs := []exportResolutionProof{}
	if requestedName == "" || len(targetFiles) == 0 {
		proofs = append(proofs, exportResolutionProof{
			Outcome:     exportResolutionMissing,
			FailureFile: firstExportResolutionPath(targetFiles),
			FailureName: requestedName,
		})
	} else if len(requestedMeanings) == 0 {
		proofs = append(proofs, exportResolutionProof{
			Outcome:     exportResolutionMeaningMismatch,
			FailureFile: firstExportResolutionPath(targetFiles),
			FailureName: requestedName,
		})
	} else {
		resolver := exportTraversal{
			workspace: w,
			active:    make(map[exportTraversalKey]struct{}),
		}
		for _, targetFile := range targetFiles {
			proofs = append(proofs, resolver.resolve(targetFile, requestedName, requestedMeanings)...)
		}
	}

	result := aggregateExportResolution(proofs)
	result.Request = cloneExportResolutionImportFact(item)
	result.RequestedName = requestedName
	result.RequestedMeanings = append([]scopeir.ExportMeaning(nil), requestedMeanings...)
	result.TargetFiles = append([]string(nil), targetFiles...)
	return result
}

// resolveSemanticImportedMember owns the proof-bearing semantic half of
// resolveImportedMember. The bool reports whether a current P5-A semantic
// import claimed the receiver; callers must not fall back to physical scans
// when it is true and the result is unresolved.
func (w *workspace) resolveSemanticImportedMember(receiver string, name string, startScope string, labels []scopeir.NodeLabel) (exportResolutionResult, bool) {
	sourceFile := w.scopeFilePath(startScope)
	if sourceFile == "" {
		return exportResolutionResult{}, false
	}
	key := importReceiverKey{filePath: sourceFile, localName: receiver}
	proofs := []exportResolutionProof{}
	var request scopeir.ImportFact
	targetFiles := []string{}
	handled := false
	ownerAmbiguous := false

	for _, importIndex := range w.importsByReceiver[key] {
		item := w.imports[importIndex]
		if !isSemanticExportImport(item.Fact) {
			continue
		}
		if !handled {
			request = cloneExportResolutionImportFact(item.Fact)
		}
		handled = true
		targetFiles = append(targetFiles, item.TargetFiles...)
		if item.LinkStatus == "unresolved" || len(item.TargetFiles) == 0 {
			proofs = append(proofs, exportResolutionProof{
				Outcome:     exportResolutionMissing,
				FailureFile: sourceFile,
				FailureName: name,
			})
			continue
		}

		if item.Fact.Kind == scopeir.ImportNamespace {
			if item.Fact.TypeOnly {
				proofs = append(proofs, exportResolutionProof{
					Outcome:     exportResolutionMeaningMismatch,
					FailureFile: sourceFile,
					FailureName: name,
				})
				continue
			}
			proofs = append(proofs, w.resolveExportRequest(
				item.Fact,
				item.TargetFiles,
				name,
				[]scopeir.ExportMeaning{scopeir.ExportMeaningValue},
			).allProofs()...)
			continue
		}

		ownerResult := w.resolveImportExport(
			item.Fact,
			item.TargetFiles,
			[]scopeir.ExportMeaning{scopeir.ExportMeaningValue, scopeir.ExportMeaningNamespace},
		)
		if ownerResult.Outcome == exportResolutionAmbiguity {
			ownerAmbiguous = true
		}
		for _, ownerFailure := range ownerResult.Failures {
			proofs = append(proofs, cloneExportResolutionProof(ownerFailure))
		}
		for _, ownerCandidate := range ownerResult.Candidates {
			for _, ownerProof := range ownerCandidate.Proofs {
				switch ownerProof.TerminalKind {
				case exportResolutionNamespaceTerminal:
					memberProofs := w.resolveExportRequest(
						item.Fact,
						[]string{ownerProof.NamespaceFile},
						name,
						[]scopeir.ExportMeaning{scopeir.ExportMeaningValue},
					).allProofs()
					for _, memberProof := range memberProofs {
						memberProof.Hops = joinExportResolutionHops(ownerProof.Hops, memberProof.Hops)
						proofs = append(proofs, memberProof)
					}
				case exportResolutionDefinitionTerminal:
					member, ok := w.resolveOwnedMember(ownerProof.Terminal.Fact.ID, name, labels)
					if !ok {
						proofs = append(proofs, missingOwnedMemberProof(ownerProof, name))
						continue
					}
					memberProof := cloneExportResolutionProof(ownerProof)
					memberProof.Terminal = cloneExportResolutionDefRef(member)
					memberProof.TerminalKind = exportResolutionDefinitionTerminal
					memberProof.Meanings = []scopeir.ExportMeaning{scopeir.ExportMeaningValue}
					memberProof.Hops = append(memberProof.Hops, exportResolutionHop{
						Kind:             exportResolutionMemberHop,
						FilePath:         cleanPath(member.Fact.FilePath),
						RequestedName:    name,
						MemberOwnerDefID: ownerProof.Terminal.Fact.ID,
						MemberName:       name,
						TerminalDefID:    member.Fact.ID,
					})
					proofs = append(proofs, memberProof)
				}
			}
		}
	}

	if !handled {
		return exportResolutionResult{}, false
	}
	result := aggregateExportResolution(proofs)
	if ownerAmbiguous {
		result.Outcome = exportResolutionAmbiguity
	}
	result.Request = request
	result.RequestedName = name
	result.RequestedMeanings = []scopeir.ExportMeaning{scopeir.ExportMeaningValue}
	result.TargetFiles = cloneSortedExportTablePaths(targetFiles)
	return result, true
}

func missingOwnedMemberProof(ownerProof exportResolutionProof, name string) exportResolutionProof {
	owner := cloneExportResolutionDefRef(ownerProof.Terminal)
	proof := cloneExportResolutionProof(ownerProof)
	proof.Outcome = exportResolutionMissing
	proof.TerminalKind = ""
	proof.Terminal = defRef{}
	proof.NamespaceFile = ""
	proof.Meanings = nil
	proof.FailureFile = cleanPath(owner.Fact.FilePath)
	proof.FailureName = name
	proof.Hops = append(proof.Hops, exportResolutionHop{
		Kind:             exportResolutionMemberHop,
		FilePath:         cleanPath(owner.Fact.FilePath),
		RequestedName:    name,
		MemberOwnerDefID: owner.Fact.ID,
		MemberName:       name,
	})
	return proof
}

// explicitImportCallState distinguishes a resolved import binding from an
// inner lexical shadow and proves whether the explicit import can supply a
// value-lane call target. It is consumed only by resolveCall.
func (w *workspace) explicitImportCallState(name string, startScope string, labels []scopeir.NodeLabel, target *defRef) (bool, bool) {
	sourceFile := w.scopeFilePath(startScope)
	if sourceFile == "" {
		return false, true
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return false, true
	}

	claimed := false
	allowed := make(map[string]struct{})
	importTargets := make(map[string]struct{})
	for index := range w.imports {
		item := &w.imports[index]
		if cleanPath(item.Fact.FilePath) != sourceFile || item.Fact.LocalName != name || !isSemanticExportImport(item.Fact) {
			continue
		}
		claimed = true
		if item.TargetDef != nil {
			importTargets[item.TargetDef.GraphID] = struct{}{}
		}
		result := w.resolveImportExport(
			item.Fact,
			item.TargetFiles,
			[]scopeir.ExportMeaning{scopeir.ExportMeaningValue},
		)
		if resolved, ok := result.definition(); ok && isAnyLabel(resolved.Fact.Label, labels) {
			allowed[resolved.GraphID] = struct{}{}
		}
	}
	if !claimed {
		return false, true
	}
	if target == nil {
		return true, false
	}
	if _, ok := allowed[target.GraphID]; ok {
		return true, true
	}
	if _, ok := importTargets[target.GraphID]; ok {
		return true, false
	}
	// A different resolved definition is an inner/local shadow, not the import.
	return false, true
}

func isSemanticExportImport(item scopeir.ImportFact) bool {
	return len(item.RequestedMeanings) > 0
}

type exportTraversalKey struct {
	filePath string
	name     string
	meanings string
}

type exportTraversal struct {
	workspace *workspace
	active    map[exportTraversalKey]struct{}
}

func (resolver *exportTraversal) resolve(filePath string, name string, requestedMeanings []scopeir.ExportMeaning) []exportResolutionProof {
	filePath = cleanExportTablePath(filePath)
	name = strings.TrimSpace(name)
	requestedMeanings = canonicalExportMeanings(requestedMeanings)
	key := exportTraversalKey{filePath: filePath, name: name, meanings: exportMeaningKey(requestedMeanings)}
	if _, active := resolver.active[key]; active {
		return []exportResolutionProof{{
			Outcome:     exportResolutionCycle,
			FailureFile: filePath,
			FailureName: name,
		}}
	}
	resolver.active[key] = struct{}{}
	defer delete(resolver.active, key)

	table, ok := resolver.workspace.exportTables[filePath]
	if !ok {
		return []exportResolutionProof{{
			Outcome:     exportResolutionMissing,
			FailureFile: filePath,
			FailureName: name,
		}}
	}

	if entries := table.Explicit[name]; len(entries) > 0 {
		proofs := make([]exportResolutionProof, 0, len(entries))
		for _, entry := range entries {
			proofs = append(proofs, resolver.resolveExplicit(filePath, name, requestedMeanings, entry)...)
		}
		return proofs
	}

	if name == "default" {
		return []exportResolutionProof{{
			Outcome:     exportResolutionMissing,
			FailureFile: filePath,
			FailureName: name,
		}}
	}
	if len(table.StarAdjacency) == 0 {
		return []exportResolutionProof{{
			Outcome:     exportResolutionMissing,
			FailureFile: filePath,
			FailureName: name,
		}}
	}

	proofs := []exportResolutionProof{}
	for _, adjacency := range table.StarAdjacency {
		narrowed := intersectExportMeanings(requestedMeanings, adjacency.Fact.Meanings)
		if len(narrowed) == 0 {
			proofs = append(proofs, exportResolutionProof{
				Outcome:     exportResolutionMeaningMismatch,
				Hops:        []exportResolutionHop{newExportResolutionHop(filePath, name, adjacency.Fact, "")},
				FailureFile: filePath,
				FailureName: name,
			})
			continue
		}
		if len(adjacency.TargetFiles) == 0 {
			proofs = append(proofs, exportResolutionProof{
				Outcome:     exportResolutionMissing,
				Hops:        []exportResolutionHop{newExportResolutionHop(filePath, name, adjacency.Fact, "")},
				FailureFile: filePath,
				FailureName: name,
			})
			continue
		}
		for _, targetFile := range adjacency.TargetFiles {
			hop := newExportResolutionHop(filePath, name, adjacency.Fact, targetFile)
			children := resolver.resolve(targetFile, name, narrowed)
			proofs = append(proofs, prependExportResolutionHop(hop, children)...)
		}
	}
	return proofs
}

func (resolver *exportTraversal) resolveExplicit(filePath string, name string, requestedMeanings []scopeir.ExportMeaning, entry exportTableEntry) []exportResolutionProof {
	narrowed := intersectExportMeanings(requestedMeanings, entry.Fact.Meanings)
	if len(narrowed) == 0 {
		return []exportResolutionProof{{
			Outcome:     exportResolutionMeaningMismatch,
			Hops:        []exportResolutionHop{newExportResolutionHop(filePath, name, entry.Fact, "")},
			FailureFile: filePath,
			FailureName: name,
		}}
	}

	switch entry.Fact.Kind {
	case scopeir.ExportReexport:
		targetName := strings.TrimSpace(entry.Fact.TargetExportedName)
		if targetName == "" || len(entry.TargetFiles) == 0 {
			return []exportResolutionProof{{
				Outcome:     exportResolutionMissing,
				Hops:        []exportResolutionHop{newExportResolutionHop(filePath, name, entry.Fact, "")},
				FailureFile: filePath,
				FailureName: firstNonEmpty(targetName, name),
			}}
		}
		proofs := []exportResolutionProof{}
		for _, targetFile := range entry.TargetFiles {
			hop := newExportResolutionHop(filePath, name, entry.Fact, targetFile)
			children := resolver.resolve(targetFile, targetName, narrowed)
			proofs = append(proofs, prependExportResolutionHop(hop, children)...)
		}
		return proofs

	case scopeir.ExportNamespace:
		if len(entry.TargetFiles) == 0 {
			return []exportResolutionProof{{
				Outcome:     exportResolutionMissing,
				Hops:        []exportResolutionHop{newExportResolutionHop(filePath, name, entry.Fact, "")},
				FailureFile: filePath,
				FailureName: name,
			}}
		}
		proofs := make([]exportResolutionProof, 0, len(entry.TargetFiles))
		for _, targetFile := range entry.TargetFiles {
			proofs = append(proofs, exportResolutionProof{
				Outcome:       exportResolutionTerminal,
				TerminalKind:  exportResolutionNamespaceTerminal,
				NamespaceFile: cleanExportTablePath(targetFile),
				Meanings:      append([]scopeir.ExportMeaning(nil), narrowed...),
				Hops:          []exportResolutionHop{newExportResolutionHop(filePath, name, entry.Fact, targetFile)},
			})
		}
		return proofs

	default:
		def, ok := resolver.workspace.defsByID[entry.Fact.LocalDefID]
		if entry.Fact.LocalDefID == "" || !ok {
			return []exportResolutionProof{{
				Outcome:     exportResolutionMissing,
				Hops:        []exportResolutionHop{newExportResolutionHop(filePath, name, entry.Fact, "")},
				FailureFile: filePath,
				FailureName: name,
			}}
		}
		return []exportResolutionProof{{
			Outcome:      exportResolutionTerminal,
			TerminalKind: exportResolutionDefinitionTerminal,
			Terminal:     cloneExportResolutionDefRef(def),
			Meanings:     append([]scopeir.ExportMeaning(nil), narrowed...),
			Hops:         []exportResolutionHop{newExportResolutionHop(filePath, name, entry.Fact, "")},
		}}
	}
}

func aggregateExportResolution(proofs []exportResolutionProof) exportResolutionResult {
	proofs = append([]exportResolutionProof(nil), proofs...)
	sortExportResolutionProofs(proofs)
	grouped := make(map[string]*exportResolutionCandidate)
	failures := []exportResolutionProof{}
	for _, proof := range proofs {
		proof = cloneExportResolutionProof(proof)
		if proof.Outcome != exportResolutionTerminal {
			failures = append(failures, proof)
			continue
		}
		key := exportResolutionCandidateKey(proof)
		candidate := grouped[key]
		if candidate == nil {
			candidate = &exportResolutionCandidate{
				Kind:          proof.TerminalKind,
				Terminal:      cloneExportResolutionDefRef(proof.Terminal),
				NamespaceFile: proof.NamespaceFile,
			}
			grouped[key] = candidate
		}
		candidate.Meanings = unionExportMeanings(candidate.Meanings, proof.Meanings)
		candidate.Proofs = append(candidate.Proofs, proof)
	}

	keys := make([]string, 0, len(grouped))
	for key := range grouped {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	candidates := make([]exportResolutionCandidate, 0, len(keys))
	for _, key := range keys {
		candidate := *grouped[key]
		sortExportResolutionProofs(candidate.Proofs)
		candidates = append(candidates, candidate)
	}

	result := exportResolutionResult{Candidates: candidates, Failures: failures}
	switch len(candidates) {
	case 0:
		result.Outcome = unresolvedExportResolutionOutcome(failures)
	case 1:
		result.Outcome = exportResolutionTerminal
	default:
		result.Outcome = exportResolutionAmbiguity
	}
	return result
}

func unresolvedExportResolutionOutcome(failures []exportResolutionProof) exportResolutionOutcome {
	if len(failures) == 0 {
		return exportResolutionMissing
	}
	allMeaningMismatch := true
	for _, failure := range failures {
		if failure.Outcome == exportResolutionCycle {
			return exportResolutionCycle
		}
		if failure.Outcome != exportResolutionMeaningMismatch {
			allMeaningMismatch = false
		}
	}
	if allMeaningMismatch {
		return exportResolutionMeaningMismatch
	}
	return exportResolutionMissing
}

func newExportResolutionHop(filePath string, requestedName string, fact scopeir.ExportFact, targetFile string) exportResolutionHop {
	return exportResolutionHop{
		Kind:          exportResolutionExportHop,
		FilePath:      cleanExportTablePath(filePath),
		RequestedName: strings.TrimSpace(requestedName),
		Fact:          cloneExportTableFact(fact),
		TargetFile:    cleanExportTablePath(targetFile),
	}
}

func prependExportResolutionHop(hop exportResolutionHop, proofs []exportResolutionProof) []exportResolutionProof {
	out := make([]exportResolutionProof, 0, len(proofs))
	for _, proof := range proofs {
		proof = cloneExportResolutionProof(proof)
		proof.Hops = append([]exportResolutionHop{cloneExportResolutionHop(hop)}, proof.Hops...)
		out = append(out, proof)
	}
	return out
}

func joinExportResolutionHops(prefix []exportResolutionHop, suffix []exportResolutionHop) []exportResolutionHop {
	out := make([]exportResolutionHop, 0, len(prefix)+len(suffix))
	for _, hop := range prefix {
		out = append(out, cloneExportResolutionHop(hop))
	}
	for _, hop := range suffix {
		out = append(out, cloneExportResolutionHop(hop))
	}
	return out
}

func cloneExportResolutionProof(proof exportResolutionProof) exportResolutionProof {
	cloned := proof
	cloned.Terminal = cloneExportResolutionDefRef(proof.Terminal)
	cloned.Meanings = append([]scopeir.ExportMeaning(nil), proof.Meanings...)
	cloned.Hops = make([]exportResolutionHop, 0, len(proof.Hops))
	for _, hop := range proof.Hops {
		cloned.Hops = append(cloned.Hops, cloneExportResolutionHop(hop))
	}
	return cloned
}

func cloneExportResolutionDefRef(ref defRef) defRef {
	cloned := ref
	cloned.Fact.ParameterTypes = append([]string(nil), ref.Fact.ParameterTypes...)
	cloned.Fact.Annotations = append([]string(nil), ref.Fact.Annotations...)
	if ref.Fact.SelectionRange != nil {
		value := *ref.Fact.SelectionRange
		cloned.Fact.SelectionRange = &value
	}
	if ref.Fact.ParameterCount != nil {
		value := *ref.Fact.ParameterCount
		cloned.Fact.ParameterCount = &value
	}
	if ref.Fact.RequiredParameterCount != nil {
		value := *ref.Fact.RequiredParameterCount
		cloned.Fact.RequiredParameterCount = &value
	}
	cloned.Fact.Static = cloneExportResolutionBool(ref.Fact.Static)
	cloned.Fact.Readonly = cloneExportResolutionBool(ref.Fact.Readonly)
	cloned.Fact.Abstract = cloneExportResolutionBool(ref.Fact.Abstract)
	cloned.Fact.Final = cloneExportResolutionBool(ref.Fact.Final)
	cloned.Fact.Virtual = cloneExportResolutionBool(ref.Fact.Virtual)
	cloned.Fact.Override = cloneExportResolutionBool(ref.Fact.Override)
	cloned.Fact.Async = cloneExportResolutionBool(ref.Fact.Async)
	cloned.Fact.Partial = cloneExportResolutionBool(ref.Fact.Partial)
	return cloned
}

func cloneExportResolutionBool(value *bool) *bool {
	if value == nil {
		return nil
	}
	cloned := *value
	return &cloned
}

func cloneExportResolutionHop(hop exportResolutionHop) exportResolutionHop {
	cloned := hop
	cloned.Fact = cloneExportTableFact(hop.Fact)
	return cloned
}

func cloneExportResolutionImportFact(item scopeir.ImportFact) scopeir.ImportFact {
	cloned := item
	cloned.RequestedMeanings = append([]scopeir.ExportMeaning(nil), item.RequestedMeanings...)
	cloned.TransitiveVia = append([]string(nil), item.TransitiveVia...)
	if item.TargetRaw != nil {
		targetRaw := *item.TargetRaw
		cloned.TargetRaw = &targetRaw
	}
	if item.TargetFile != nil {
		targetFile := *item.TargetFile
		cloned.TargetFile = &targetFile
	}
	return cloned
}

func canonicalExportMeanings(values []scopeir.ExportMeaning) []scopeir.ExportMeaning {
	if len(values) == 0 {
		return nil
	}
	out := append([]scopeir.ExportMeaning(nil), values...)
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	writeIndex := 0
	for _, value := range out {
		if value == "" || (writeIndex > 0 && out[writeIndex-1] == value) {
			continue
		}
		out[writeIndex] = value
		writeIndex++
	}
	return out[:writeIndex]
}

func intersectExportMeanings(left []scopeir.ExportMeaning, right []scopeir.ExportMeaning) []scopeir.ExportMeaning {
	left = canonicalExportMeanings(left)
	right = canonicalExportMeanings(right)
	if len(left) == 0 || len(right) == 0 {
		return nil
	}
	out := make([]scopeir.ExportMeaning, 0, len(left))
	leftIndex := 0
	rightIndex := 0
	for leftIndex < len(left) && rightIndex < len(right) {
		switch {
		case left[leftIndex] < right[rightIndex]:
			leftIndex++
		case right[rightIndex] < left[leftIndex]:
			rightIndex++
		default:
			out = append(out, left[leftIndex])
			leftIndex++
			rightIndex++
		}
	}
	return out
}

func unionExportMeanings(left []scopeir.ExportMeaning, right []scopeir.ExportMeaning) []scopeir.ExportMeaning {
	return canonicalExportMeanings(append(append([]scopeir.ExportMeaning(nil), left...), right...))
}

func exportMeaningKey(values []scopeir.ExportMeaning) string {
	values = canonicalExportMeanings(values)
	parts := make([]string, 0, len(values))
	for _, value := range values {
		parts = append(parts, string(value))
	}
	return strings.Join(parts, ",")
}

func exportResolutionCandidateKey(proof exportResolutionProof) string {
	if proof.TerminalKind == exportResolutionNamespaceTerminal {
		return string(proof.TerminalKind) + "\x00" + cleanExportTablePath(proof.NamespaceFile)
	}
	return string(proof.TerminalKind) + "\x00" + proof.Terminal.GraphID
}

func sortExportResolutionProofs(proofs []exportResolutionProof) {
	sort.SliceStable(proofs, func(i, j int) bool {
		return exportResolutionProofKey(proofs[i]) < exportResolutionProofKey(proofs[j])
	})
}

func exportResolutionProofKey(proof exportResolutionProof) string {
	parts := []string{
		string(proof.Outcome),
		string(proof.TerminalKind),
		proof.Terminal.GraphID,
		cleanExportTablePath(proof.NamespaceFile),
		exportMeaningKey(proof.Meanings),
		cleanExportTablePath(proof.FailureFile),
		proof.FailureName,
	}
	for _, hop := range proof.Hops {
		parts = append(parts,
			string(hop.Kind),
			hop.FilePath,
			hop.RequestedName,
			exportTableFactKey(hop.Fact),
			hop.TargetFile,
			hop.MemberOwnerDefID,
			hop.MemberName,
			hop.TerminalDefID,
		)
	}
	return strings.Join(parts, "\x00")
}

func firstExportResolutionPath(values []string) string {
	values = cloneSortedExportTablePaths(values)
	if len(values) == 0 {
		return ""
	}
	return values[0]
}
