package tsstdlib

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestEmbeddedCatalogMetadataAndDefaultMeaningLanes(t *testing.T) {
	authority := NewAuthority(".", []string{"src/main.ts"})
	metadata := authority.Metadata()
	if metadata.SchemaVersion != SchemaVersion ||
		metadata.AuthorityKind != AuthorityKind ||
		metadata.IdentityVersion != IdentityVersion ||
		metadata.TypeScriptVersion != TypeScriptVersion ||
		metadata.PackageIntegrity != TypeScriptPackageIntegrity ||
		metadata.GenerationCommand != GenerationCommand ||
		metadata.InputCount != 100 ||
		metadata.InputBytes != 3141835 ||
		metadata.SymbolCount == 0 ||
		metadata.MemberCount == 0 ||
		metadata.CatalogHash == "" ||
		metadata.CatalogArtifactHash != sha256Hex(embeddedCatalog) ||
		metadata.AuthorityHash == "" {
		t.Fatalf("unexpected catalog metadata: %#v", metadata)
	}
	profile := authority.Profile()
	if profile.Status != ProfileReady || profile.Target != "default" || profile.ConfigHash == "" || profile.ProfileHash == "" {
		t.Fatalf("unexpected default profile: %#v", profile)
	}

	requireLookupStatus(t, authority.LookupGlobal("Promise", MeaningType), LookupResolved, "")
	requireLookupStatus(t, authority.LookupGlobal("Promise", MeaningValue), LookupProfileExcluded, ReasonProfileExcludes)
	requireLookupStatus(t, authority.LookupGlobal("Math", MeaningValue), LookupResolved, "")
	requireLookupStatus(t, authority.LookupGlobal("Math", MeaningNamespace), LookupMeaningMismatch, ReasonMeaningMismatch)
	requireLookupStatus(t, authority.LookupMember("Math", MeaningValue, "max", MeaningValue), LookupResolved, "")
	requireLookupStatus(t, authority.LookupMember("Math", MeaningValue, "min", MeaningValue), LookupResolved, "")
	requireLookupStatus(t, authority.LookupMember("HTMLInputElement", MeaningType, "addEventListener", MeaningValue), LookupResolved, "")
	requireLookupStatus(t, authority.LookupGlobal("DefinitelyNotATypeScriptGlobal", MeaningValue), LookupNotFound, "")
}

func TestP6BExactCompilerVectorMatrix(t *testing.T) {
	values := func(status LookupStatus, reason Reason) []lookupCheck {
		return []lookupCheck{
			{name: "Promise", meaning: MeaningValue, status: status, reason: reason},
			{name: "Map", meaning: MeaningValue, status: status, reason: reason},
			{name: "Set", meaning: MeaningValue, status: status, reason: reason},
			{name: "Symbol", meaning: MeaningValue, status: status, reason: reason},
		}
	}
	globalTypes := []string{"Array", "Boolean", "Function", "IArguments", "Number", "Object", "RegExp", "String"}
	noLibGlobals := make([]lookupCheck, 0, len(globalTypes))
	for _, name := range globalTypes {
		noLibGlobals = append(noLibGlobals, lookupCheck{
			name: name, meaning: MeaningType, status: LookupCapabilityUnavailable, reason: ReasonDisabledByNoLib,
		})
	}
	typeAndMembers := func(status LookupStatus, reason Reason) ([]lookupCheck, []memberLookupCheck) {
		return []lookupCheck{{name: "Promise", meaning: MeaningType, status: status, reason: reason}}, []memberLookupCheck{
			{owner: "Math", ownerMeaning: MeaningValue, name: "max", meaning: MeaningValue, status: status, reason: reason},
			{owner: "Math", ownerMeaning: MeaningValue, name: "min", meaning: MeaningValue, status: status, reason: reason},
		}
	}
	defaultType, defaultMembers := typeAndMembers(LookupResolved, "")
	es5Type, es5Members := typeAndMembers(LookupResolved, "")
	noLibType, noLibMembers := typeAndMembers(LookupCapabilityUnavailable, ReasonDisabledByNoLib)

	vectors := []struct {
		name      string
		authority *Authority
		globals   []lookupCheck
		members   []memberLookupCheck
	}{
		{name: "compiler-default-es2015-values-fail", authority: NewAuthority(".", []string{"src/main.ts"}), globals: values(LookupProfileExcluded, ReasonProfileExcludes)},
		{name: "target-es2022-es2015-values-pass", authority: fixtureAuthority("es2022"), globals: values(LookupResolved, "")},
		{name: "explicit-lib-es5-es2015-values-fail", authority: fixtureAuthority("es5"), globals: values(LookupProfileExcluded, ReasonProfileExcludes)},
		{name: "explicit-lib-es2015-es2015-values-pass", authority: fixtureAuthority("jsonc"), globals: values(LookupResolved, "")},
		{name: "es5-plus-es2015-promise-split", authority: fixtureAuthority("es5-promise"), globals: []lookupCheck{
			{name: "Promise", meaning: MeaningValue, status: LookupResolved},
			{name: "Map", meaning: MeaningValue, status: LookupProfileExcluded, reason: ReasonProfileExcludes},
			{name: "Set", meaning: MeaningValue, status: LookupProfileExcluded, reason: ReasonProfileExcludes},
			{name: "Symbol", meaning: MeaningValue, status: LookupProfileExcluded, reason: ReasonProfileExcludes},
		}},
		{name: "no-lib-global-types-fail", authority: fixtureAuthority("no-lib"), globals: noLibGlobals},
		{name: "invalid-lib-compiler-option-error", authority: fixtureAuthority("invalid-lib"), globals: []lookupCheck{
			{name: "Promise", meaning: MeaningValue, status: LookupCapabilityUnavailable, reason: ReasonConfigInvalid},
			{name: "Promise", meaning: MeaningType, status: LookupCapabilityUnavailable, reason: ReasonConfigInvalid},
		}, members: []memberLookupCheck{{owner: "Math", ownerMeaning: MeaningValue, name: "max", meaning: MeaningValue, status: LookupCapabilityUnavailable, reason: ReasonConfigInvalid}}},
		{name: "compiler-default-promise-type-and-math-members-pass", authority: NewAuthority(".", []string{"src/main.ts"}), globals: defaultType, members: defaultMembers},
		{name: "explicit-lib-es5-promise-type-and-math-members-pass", authority: fixtureAuthority("es5"), globals: es5Type, members: es5Members},
		{name: "no-lib-promise-type-and-math-members-fail", authority: fixtureAuthority("no-lib"), globals: noLibType, members: noLibMembers},
	}
	if len(vectors) != 10 {
		t.Fatalf("compiler vector denominator = %d, want exactly 10", len(vectors))
	}
	for index, vector := range vectors {
		t.Run(fmt.Sprintf("vector-%02d-%s", index+1, vector.name), func(t *testing.T) {
			for _, check := range vector.globals {
				requireLookupStatus(t, vector.authority.LookupGlobal(check.name, check.meaning), check.status, check.reason)
			}
			for _, check := range vector.members {
				requireLookupStatus(t, vector.authority.LookupMember(check.owner, check.ownerMeaning, check.name, check.meaning), check.status, check.reason)
			}
		})
	}
}

func TestProfileSelectionMatchesCompilerDeclarationClosures(t *testing.T) {
	tests := []struct {
		name          string
		profile       string
		wantLibraries int
		checks        []lookupCheck
	}{
		{
			name:          "target-es2022",
			profile:       "es2022",
			wantLibraries: 63,
			checks: []lookupCheck{
				{name: "Promise", meaning: MeaningValue, status: LookupResolved},
				{name: "Map", meaning: MeaningValue, status: LookupResolved},
				{name: "Set", meaning: MeaningValue, status: LookupResolved},
				{name: "Symbol", meaning: MeaningValue, status: LookupResolved},
			},
		},
		{
			name:          "explicit-es5",
			profile:       "es5",
			wantLibraries: 3,
			checks: []lookupCheck{
				{name: "Promise", meaning: MeaningType, status: LookupResolved},
				{name: "Promise", meaning: MeaningValue, status: LookupProfileExcluded, reason: ReasonProfileExcludes},
				{name: "Math", meaning: MeaningValue, status: LookupResolved},
			},
		},
		{
			name:          "jsonc-es2015",
			profile:       "jsonc",
			wantLibraries: 13,
			checks: []lookupCheck{
				{name: "Promise", meaning: MeaningValue, status: LookupResolved},
				{name: "Map", meaning: MeaningValue, status: LookupResolved},
			},
		},
		{
			name:    "es5-plus-promise",
			profile: "es5-promise",
			checks: []lookupCheck{
				{name: "Promise", meaning: MeaningValue, status: LookupResolved},
				{name: "Map", meaning: MeaningValue, status: LookupProfileExcluded, reason: ReasonProfileExcludes},
				{name: "Set", meaning: MeaningValue, status: LookupProfileExcluded, reason: ReasonProfileExcludes},
				{name: "Symbol", meaning: MeaningValue, status: LookupProfileExcluded, reason: ReasonProfileExcludes},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			authority := fixtureAuthority(test.profile)
			profile := authority.Profile()
			if profile.Status != ProfileReady {
				t.Fatalf("profile is unavailable: %#v", profile)
			}
			if test.wantLibraries > 0 && len(profile.Libraries) != test.wantLibraries {
				t.Fatalf("library closure = %d, want %d: %#v", len(profile.Libraries), test.wantLibraries, profile.Libraries)
			}
			for _, check := range test.checks {
				requireLookupStatus(t, authority.LookupGlobal(check.name, check.meaning), check.status, check.reason)
			}
		})
	}
}

func TestUnavailableProfilesFailClosed(t *testing.T) {
	tests := []struct {
		name      string
		authority *Authority
		reason    Reason
	}{
		{name: "no-lib", authority: fixtureAuthority("no-lib"), reason: ReasonDisabledByNoLib},
		{name: "invalid-lib", authority: fixtureAuthority("invalid-lib"), reason: ReasonConfigInvalid},
		{name: "unsupported-topology", authority: fixtureAuthority("topology"), reason: ReasonConfigTopology},
		{name: "nested-config", authority: NewAuthority(".", []string{"src/main.ts", "nested/tsconfig.json"}), reason: ReasonConfigTopology},
		{name: "jsconfig", authority: NewAuthority(".", []string{"src/main.ts", "jsconfig.json"}), reason: ReasonConfigTopology},
		{name: "unreadable", authority: NewAuthority(filepath.Join("testdata", "profiles", "missing"), []string{"src/main.ts", "tsconfig.json"}), reason: ReasonConfigUnreadable},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			profile := test.authority.Profile()
			if profile.Status != ProfileUnavailable || profile.Reason != test.reason {
				t.Fatalf("profile = %#v, want unavailable/%s", profile, test.reason)
			}
			requireLookupStatus(
				t,
				test.authority.LookupGlobal("Promise", MeaningType),
				LookupCapabilityUnavailable,
				test.reason,
			)
		})
	}
}

func TestCatalogValidationRejectsHashVersionAndManifestDrift(t *testing.T) {
	if _, err := loadCatalog(embeddedCatalog); err != nil {
		t.Fatalf("embedded catalog failed validation: %v", err)
	}
	requireCatalogReason(t, append(append([]byte(nil), embeddedCatalog...), []byte("{}")...), ReasonCatalogInputManifest)

	var dto catalogDTO
	if err := json.Unmarshal(embeddedCatalog, &dto); err != nil {
		t.Fatalf("decode embedded catalog: %v", err)
	}

	hashDrift := dto
	hashDrift.Hash = "0" + hashDrift.Hash[1:]
	requireCatalogReason(t, mustMarshalCatalog(t, hashDrift), ReasonCatalogHash)

	schemaDrift := dto
	schemaDrift.Schema = "tsstdlib.catalog.v999"
	requireCatalogReason(t, mustMarshalCatalog(t, schemaDrift), ReasonCatalogSchema)

	versionDrift := dto
	versionDrift.TypeScript = "0.0.0"
	requireCatalogReason(t, mustMarshalCatalog(t, versionDrift), ReasonCatalogVersion)

	manifestDrift := dto
	manifestDrift.Inputs = append([]inputDTO(nil), dto.Inputs...)
	manifestDrift.Inputs[0].References = []int{len(dto.Inputs)}
	manifestDrift.Hash = ""
	canonical, err := marshalCatalogCanonical(manifestDrift)
	if err != nil {
		t.Fatalf("marshal canonical catalog: %v", err)
	}
	manifestDrift.Hash = sha256Hex(canonical)
	requireCatalogReason(t, mustMarshalCatalog(t, manifestDrift), ReasonCatalogInputManifest)

	identityDrift := dto
	identityDrift.Symbols = append([]symbolDTO(nil), dto.Symbols...)
	promiseIndex := catalogSymbolIndex(t, identityDrift.Symbols, "Promise")
	identityDrift.Symbols[promiseIndex].ValueID = "tsstdlib:" + strings.Repeat("0", 64)
	requireCatalogReasonContains(t, mustMarshalCatalog(t, identityDrift), ReasonCatalogInputManifest, "component drift")

	duplicateID := dto
	duplicateID.Symbols = append([]symbolDTO(nil), dto.Symbols...)
	promiseIndex = catalogSymbolIndex(t, duplicateID.Symbols, "Promise")
	duplicateID.Symbols[promiseIndex].TypeID = duplicateID.Symbols[promiseIndex].ValueID
	requireCatalogReasonContains(t, mustMarshalCatalog(t, duplicateID), ReasonCatalogInputManifest, "duplicate semantic ID")
}

func TestCatalogFailuresPreserveExactLookupReasonAndProofAbsence(t *testing.T) {
	tests := []struct {
		name       string
		fixture    string
		reason     Reason
		proofState CatalogProofState
	}{
		{name: "empty-missing", reason: ReasonCatalogMissing, proofState: CatalogProofMissing},
		{name: "schema", fixture: "schema.json", reason: ReasonCatalogSchema, proofState: CatalogProofRejected},
		{name: "version", fixture: "version.json", reason: ReasonCatalogVersion, proofState: CatalogProofRejected},
		{name: "input-manifest", fixture: "input-manifest.json", reason: ReasonCatalogInputManifest, proofState: CatalogProofRejected},
		{name: "logical-hash", fixture: "logical-hash.json", reason: ReasonCatalogHash, proofState: CatalogProofRejected},
		{name: "trailing-artifact-integrity", fixture: "trailing-artifact-integrity.json", reason: ReasonCatalogInputManifest, proofState: CatalogProofRejected},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			raw := readCatalogFailureFixture(t, test.fixture)
			authority := NewAuthorityFromCatalog(".", []string{"src/main.ts"}, raw)
			profile := authority.Profile()
			if profile.Status != ProfileUnavailable || profile.Reason != test.reason || profile.ProfileHash == "" || profile.ConfigHash == "" {
				t.Fatalf("catalog failure profile = %#v, want unavailable/%s with profile+inventory proof", profile, test.reason)
			}
			lookups := []LookupResult{
				authority.LookupGlobal("parseInt", MeaningValue),
				authority.LookupGlobal("Promise", MeaningType),
				authority.LookupMember("Math", MeaningValue, "max", MeaningValue),
			}
			for _, result := range lookups {
				if result.Status != LookupCapabilityUnavailable || result.Reason != test.reason || result.CatalogProofState != test.proofState {
					t.Fatalf("lookup %s/%s = %#v, want unavailable/%s/%s", result.Name, result.Meaning, result, test.reason, test.proofState)
				}
				if result.AuthorityKind != AuthorityKind || result.TypeScriptVersion != TypeScriptVersion ||
					result.ProfileHash == "" || result.ConfigHash == "" || result.AuthorityHash != "" || result.CatalogHash != "" {
					t.Fatalf("lookup %s/%s did not preserve the explicit rejected-catalog proof contract: %#v", result.Name, result.Meaning, result)
				}
				if test.proofState == CatalogProofMissing {
					if result.CatalogArtifactHash != "" {
						t.Fatalf("missing catalog fabricated an artifact hash: %#v", result)
					}
				} else if result.CatalogArtifactHash != sha256Hex(raw) {
					t.Fatalf("rejected catalog artifact hash = %q, want attempted bytes %q", result.CatalogArtifactHash, sha256Hex(raw))
				}
			}
		})
	}
}

func readCatalogFailureFixture(t *testing.T, name string) []byte {
	t.Helper()
	if name == "" {
		return nil
	}
	raw, err := os.ReadFile(filepath.Join("testdata", "catalog-failures", name))
	if err != nil {
		t.Fatalf("read catalog failure fixture %s: %v", name, err)
	}
	return raw
}

func TestCatalogSemanticIdentityInventoryIsUniqueAndComplete(t *testing.T) {
	catalog, err := loadCatalog(embeddedCatalog)
	if err != nil {
		t.Fatalf("load embedded catalog: %v", err)
	}
	seen := make(map[string]string)
	check := func(ownerPath [][]string, name string, meaning Meaning, refs []declarationDTO, id string) {
		t.Helper()
		if len(refs) == 0 {
			if id != "" {
				t.Fatalf("identity exists without lane: owner=%v name=%s meaning=%s id=%s", ownerPath, name, meaning, id)
			}
			return
		}
		expected, err := catalogSemanticID(catalog.dto, ownerPath, name, meaning, refs)
		if err != nil {
			t.Fatalf("derive semantic identity for %v/%s/%s: %v", ownerPath, name, meaning, err)
		}
		if id != expected {
			t.Fatalf("semantic identity drift: owner=%v name=%s meaning=%s got=%s want=%s", ownerPath, name, meaning, id, expected)
		}
		if previous, duplicate := seen[id]; duplicate {
			t.Fatalf("duplicate semantic identity %s for %s and %v/%s/%s", id, previous, ownerPath, name, meaning)
		}
		seen[id] = fmt.Sprintf("%v/%s/%s", ownerPath, name, meaning)
	}
	for _, symbol := range catalog.dto.Symbols {
		check([][]string{{"global"}}, symbol.Name, MeaningValue, symbol.Value, symbol.ValueID)
		check([][]string{{"global"}}, symbol.Name, MeaningType, symbol.Type, symbol.TypeID)
		check([][]string{{"global"}}, symbol.Name, MeaningNamespace, symbol.Namespace, symbol.NamespaceID)
		groups := []struct {
			meaning Meaning
			members []memberDTO
		}{
			{meaning: MeaningValue, members: symbol.ValueMembers},
			{meaning: MeaningType, members: symbol.TypeMembers},
			{meaning: MeaningNamespace, members: symbol.NamespaceMembers},
		}
		for _, group := range groups {
			ownerPath := [][]string{{"global", symbol.Name, string(group.meaning)}}
			for _, member := range group.members {
				check(ownerPath, member.Name, MeaningValue, member.Value, member.ValueID)
				check(ownerPath, member.Name, MeaningType, member.Type, member.TypeID)
				check(ownerPath, member.Name, MeaningNamespace, member.Namespace, member.NamespaceID)
			}
		}
	}
	if len(seen) != 14587 {
		t.Fatalf("semantic identity inventory = %d, want 14587", len(seen))
	}
}

type lookupCheck struct {
	name    string
	meaning Meaning
	status  LookupStatus
	reason  Reason
}

type memberLookupCheck struct {
	owner        string
	ownerMeaning Meaning
	name         string
	meaning      Meaning
	status       LookupStatus
	reason       Reason
}

func fixtureAuthority(name string) *Authority {
	return NewAuthority(
		filepath.Join("testdata", "profiles", name),
		[]string{"src/main.ts", "tsconfig.json"},
	)
}

func requireLookupStatus(t *testing.T, result LookupResult, status LookupStatus, reason Reason) {
	t.Helper()
	if result.Status != status || result.Reason != reason {
		t.Fatalf("lookup %q/%s = %#v, want status=%s reason=%s", result.Name, result.Meaning, result, status, reason)
	}
	if status == LookupResolved && (result.SymbolID == "" || len(result.Declarations) == 0 || result.AuthorityHash == "" || result.CatalogHash == "" || result.CatalogArtifactHash == "" || result.ProfileHash == "" || result.ConfigHash == "") {
		t.Fatalf("resolved lookup is missing provenance: %#v", result)
	}
}

func catalogSymbolIndex(t *testing.T, symbols []symbolDTO, name string) int {
	t.Helper()
	for index := range symbols {
		if symbols[index].Name == name {
			return index
		}
	}
	t.Fatalf("catalog symbol %q not found", name)
	return -1
}

func requireCatalogReason(t *testing.T, raw []byte, reason Reason) {
	t.Helper()
	_, err := loadCatalog(raw)
	if err == nil {
		t.Fatalf("catalog unexpectedly passed validation")
	}
	var typed *catalogError
	if !errors.As(err, &typed) || typed.reason != reason {
		t.Fatalf("catalog error = %v, want reason %s", err, reason)
	}
}

func requireCatalogReasonContains(t *testing.T, raw []byte, reason Reason, text string) {
	t.Helper()
	_, err := loadCatalog(raw)
	if err == nil {
		t.Fatalf("catalog unexpectedly passed validation")
	}
	var typed *catalogError
	if !errors.As(err, &typed) || typed.reason != reason || !strings.Contains(err.Error(), text) {
		t.Fatalf("catalog error = %v, want reason %s containing %q", err, reason, text)
	}
}

func mustMarshalCatalog(t *testing.T, dto catalogDTO) []byte {
	t.Helper()
	raw, err := json.Marshal(dto)
	if err != nil {
		t.Fatalf("marshal catalog: %v", err)
	}
	return raw
}

func BenchmarkLoadEmbeddedCatalog(b *testing.B) {
	b.ReportAllocs()
	for index := 0; index < b.N; index++ {
		if _, err := loadCatalog(embeddedCatalog); err != nil {
			b.Fatalf("load embedded catalog: %v", err)
		}
	}
}

func BenchmarkLookupGlobal(b *testing.B) {
	authority := NewAuthority(".", []string{"src/main.ts"})
	b.ReportAllocs()
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		result := authority.LookupGlobal("Promise", MeaningType)
		if result.Status != LookupResolved {
			b.Fatalf("lookup Promise type: %#v", result)
		}
	}
}

func BenchmarkLookupInheritedMember(b *testing.B) {
	authority := NewAuthority(".", []string{"src/main.ts"})
	b.ReportAllocs()
	b.ResetTimer()
	for index := 0; index < b.N; index++ {
		result := authority.LookupMember("HTMLInputElement", MeaningType, "addEventListener", MeaningValue)
		if result.Status != LookupResolved {
			b.Fatalf("lookup inherited member: %#v", result)
		}
	}
}
