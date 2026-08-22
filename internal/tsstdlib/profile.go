package tsstdlib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const rootConfigName = "tsconfig.json"

func selectProfile(catalog *catalogIndex, repoRoot string, inventory []string) Profile {
	configPaths, unsupported := declarationConfigInventory(inventory)
	if len(unsupported) > 0 {
		return unavailableProfile(ReasonConfigTopology, "", inventoryHash(append(configPaths, unsupported...)))
	}
	if len(configPaths) == 0 {
		return readyProfile(catalog, "default", nil, "", sha256Hex([]byte("absent")))
	}
	if len(configPaths) != 1 || !strings.EqualFold(configPaths[0], rootConfigName) {
		return unavailableProfile(ReasonConfigTopology, "", inventoryHash(configPaths))
	}

	configPath := configPaths[0]
	absolutePath, ok := repositoryPath(repoRoot, configPath)
	if !ok {
		return unavailableProfile(ReasonConfigTopology, configPath, inventoryHash(configPaths))
	}
	raw, err := os.ReadFile(absolutePath)
	if err != nil {
		return unavailableProfile(ReasonConfigUnreadable, configPath, inventoryHash(configPaths))
	}
	configHash := sha256Hex(raw)
	cleaned, err := stripJSONC(raw)
	if err != nil {
		return unavailableProfile(ReasonConfigInvalid, configPath, configHash)
	}
	var root map[string]json.RawMessage
	if err := json.Unmarshal(cleaned, &root); err != nil || root == nil {
		return unavailableProfile(ReasonConfigInvalid, configPath, configHash)
	}
	for _, key := range []string{"extends", "references", "files", "include", "exclude"} {
		if _, present := root[key]; present {
			return unavailableProfile(ReasonConfigTopology, configPath, configHash)
		}
	}

	compilerOptions := map[string]json.RawMessage{}
	if rawOptions, present := root["compilerOptions"]; present {
		if err := json.Unmarshal(rawOptions, &compilerOptions); err != nil || compilerOptions == nil {
			return unavailableProfile(ReasonConfigInvalid, configPath, configHash)
		}
	}

	target := "default"
	if rawTarget, present := compilerOptions["target"]; present {
		var value string
		if err := json.Unmarshal(rawTarget, &value); err != nil {
			return unavailableProfile(ReasonConfigInvalid, configPath, configHash)
		}
		target = strings.ToLower(strings.TrimSpace(value))
		if _, ok := catalog.dto.Targets[target]; !ok {
			return unavailableProfile(ReasonConfigInvalid, configPath, configHash)
		}
	}

	noLib := false
	if rawNoLib, present := compilerOptions["noLib"]; present {
		if err := json.Unmarshal(rawNoLib, &noLib); err != nil {
			return unavailableProfile(ReasonConfigInvalid, configPath, configHash)
		}
	}

	var roots []int
	rawLib, explicitLib := compilerOptions["lib"]
	if explicitLib {
		if noLib {
			return unavailableProfile(ReasonConfigInvalid, configPath, configHash)
		}
		var values []string
		if err := json.Unmarshal(rawLib, &values); err != nil || values == nil {
			return unavailableProfile(ReasonConfigInvalid, configPath, configHash)
		}
		seen := make(map[int]struct{}, len(values))
		for _, value := range values {
			alias := normalizeLibName(value)
			rootIndex, ok := catalog.dto.Aliases[alias]
			if !ok {
				return unavailableProfile(ReasonConfigInvalid, configPath, configHash)
			}
			if _, duplicate := seen[rootIndex]; duplicate {
				continue
			}
			seen[rootIndex] = struct{}{}
			roots = append(roots, rootIndex)
		}
		sort.Ints(roots)
	}
	if noLib {
		return unavailableProfile(ReasonDisabledByNoLib, configPath, configHash)
	}
	return readyProfile(catalog, target, roots, configPath, configHash)
}

func readyProfile(catalog *catalogIndex, target string, explicitRoots []int, configPath string, configHash string) Profile {
	roots := explicitRoots
	if roots == nil {
		root, ok := catalog.dto.Targets[target]
		if !ok {
			return unavailableProfile(ReasonConfigInvalid, configPath, configHash)
		}
		roots = []int{root}
	}
	librarySet := catalog.libraryClosure(roots)
	libraries := make([]string, 0, len(librarySet))
	for index := range librarySet {
		libraries = append(libraries, catalog.dto.Inputs[index].Path)
	}
	sort.Strings(libraries)
	profile := Profile{
		Status:     ProfileReady,
		Target:     target,
		Libraries:  libraries,
		ConfigPath: configPath,
		ConfigHash: configHash,
		librarySet: librarySet,
	}
	profile.ProfileHash = profileDigest(catalog.dto.Hash, profile)
	return profile
}

func unavailableProfile(reason Reason, configPath string, configHash string) Profile {
	profile := Profile{
		Status:     ProfileUnavailable,
		Reason:     reason,
		ConfigPath: configPath,
		ConfigHash: configHash,
	}
	profile.ProfileHash = profileDigest("", profile)
	return profile
}

func (catalog *catalogIndex) libraryClosure(roots []int) map[int]struct{} {
	closure := make(map[int]struct{}, len(roots))
	queue := append([]int(nil), roots...)
	for len(queue) > 0 {
		index := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if index < 0 || index >= len(catalog.dto.Inputs) {
			continue
		}
		if _, seen := closure[index]; seen {
			continue
		}
		closure[index] = struct{}{}
		queue = append(queue, catalog.dto.Inputs[index].References...)
	}
	return closure
}

func declarationConfigInventory(inventory []string) (tsconfigs []string, unsupported []string) {
	seenTSConfig := make(map[string]struct{})
	seenUnsupported := make(map[string]struct{})
	for _, value := range inventory {
		cleaned := filepath.ToSlash(filepath.Clean(strings.TrimSpace(value)))
		cleaned = strings.TrimPrefix(cleaned, "./")
		if cleaned == "." || cleaned == "" {
			continue
		}
		base := strings.ToLower(filepath.Base(cleaned))
		switch base {
		case rootConfigName:
			key := strings.ToLower(cleaned)
			if _, duplicate := seenTSConfig[key]; duplicate {
				continue
			}
			seenTSConfig[key] = struct{}{}
			tsconfigs = append(tsconfigs, cleaned)
		case "jsconfig.json":
			key := strings.ToLower(cleaned)
			if _, duplicate := seenUnsupported[key]; duplicate {
				continue
			}
			seenUnsupported[key] = struct{}{}
			unsupported = append(unsupported, cleaned)
		}
	}
	sort.Strings(tsconfigs)
	sort.Strings(unsupported)
	return tsconfigs, unsupported
}

func repositoryPath(repoRoot string, relative string) (string, bool) {
	root, err := filepath.Abs(repoRoot)
	if err != nil {
		return "", false
	}
	candidate, err := filepath.Abs(filepath.Join(root, filepath.FromSlash(relative)))
	if err != nil {
		return "", false
	}
	rel, err := filepath.Rel(root, candidate)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", false
	}
	return candidate, true
}

func normalizeLibName(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	value = strings.TrimPrefix(value, "lib.")
	value = strings.TrimSuffix(value, ".d.ts")
	return value
}

func inventoryHash(values []string) string {
	cloned := append([]string(nil), values...)
	sort.Strings(cloned)
	return sha256Hex([]byte(strings.Join(cloned, "\n")))
}

func profileDigest(catalogHash string, profile Profile) string {
	parts := []string{
		catalogHash,
		string(profile.Status),
		string(profile.Reason),
		profile.Target,
		profile.ConfigPath,
		profile.ConfigHash,
		strings.Join(profile.Libraries, "\n"),
	}
	return sha256Hex([]byte(strings.Join(parts, "\x00")))
}

func stripJSONC(raw []byte) ([]byte, error) {
	withoutComments := make([]byte, 0, len(raw))
	inString := false
	escaped := false
	for index := 0; index < len(raw); index++ {
		current := raw[index]
		if inString {
			withoutComments = append(withoutComments, current)
			switch {
			case escaped:
				escaped = false
			case current == '\\':
				escaped = true
			case current == '"':
				inString = false
			}
			continue
		}
		if current == '"' {
			inString = true
			withoutComments = append(withoutComments, current)
			continue
		}
		if current == '/' && index+1 < len(raw) {
			next := raw[index+1]
			switch next {
			case '/':
				withoutComments = append(withoutComments, ' ', ' ')
				index += 2
				for ; index < len(raw); index++ {
					if raw[index] == '\n' || raw[index] == '\r' {
						withoutComments = append(withoutComments, raw[index])
						break
					}
					withoutComments = append(withoutComments, ' ')
				}
				continue
			case '*':
				withoutComments = append(withoutComments, ' ', ' ')
				index += 2
				closed := false
				for ; index < len(raw); index++ {
					if raw[index] == '*' && index+1 < len(raw) && raw[index+1] == '/' {
						withoutComments = append(withoutComments, ' ', ' ')
						index++
						closed = true
						break
					}
					if raw[index] == '\n' || raw[index] == '\r' {
						withoutComments = append(withoutComments, raw[index])
					} else {
						withoutComments = append(withoutComments, ' ')
					}
				}
				if !closed {
					return nil, fmt.Errorf("unterminated JSONC block comment")
				}
				continue
			}
		}
		withoutComments = append(withoutComments, current)
	}
	if inString {
		return nil, fmt.Errorf("unterminated JSON string")
	}

	withoutTrailingCommas := make([]byte, 0, len(withoutComments))
	inString = false
	escaped = false
	for index := 0; index < len(withoutComments); index++ {
		current := withoutComments[index]
		if inString {
			withoutTrailingCommas = append(withoutTrailingCommas, current)
			switch {
			case escaped:
				escaped = false
			case current == '\\':
				escaped = true
			case current == '"':
				inString = false
			}
			continue
		}
		if current == '"' {
			inString = true
			withoutTrailingCommas = append(withoutTrailingCommas, current)
			continue
		}
		if current == ',' {
			next := index + 1
			for next < len(withoutComments) && bytes.ContainsRune([]byte(" \t\r\n"), rune(withoutComments[next])) {
				next++
			}
			if next < len(withoutComments) && (withoutComments[next] == '}' || withoutComments[next] == ']') {
				withoutTrailingCommas = append(withoutTrailingCommas, ' ')
				continue
			}
		}
		withoutTrailingCommas = append(withoutTrailingCommas, current)
	}
	return withoutTrailingCommas, nil
}
