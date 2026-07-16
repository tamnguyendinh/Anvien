#!/usr/bin/env python3
"""Monitor Tree-sitter Go module drift for Anvien.

Anvien's parser stack is currently owned by Go modules in go.mod.  This
checker reports drift for those modules and separately records when the
upstream Tree-sitter core release is ahead of the latest Go binding tag.

Exit code:
  0 - no actionable Go module drift was found
  1 - actionable module drift, metadata fetch failure, or checker failure
"""

from __future__ import annotations

from dataclasses import dataclass
import json
import os
import pathlib
import re
import subprocess
import sys
import urllib.error
import urllib.request

REPO_ROOT = pathlib.Path(__file__).resolve().parents[2]
TREE_SITTER_RELEASE_API = "https://api.github.com/repos/tree-sitter/tree-sitter/releases/latest"
TREE_SITTER_GO_BINDING = "github.com/tree-sitter/go-tree-sitter"

TREE_SITTER_MODULE_PREFIXES = (
    "github.com/tree-sitter/",
    "github.com/tree-sitter-grammars/",
)

TREE_SITTER_EXACT_MODULES = {
    "github.com/UserNobody14/tree-sitter-dart",
    "github.com/flamingoosesoftwareinc/tree-sitter-swift",
    "github.com/smacker/go-tree-sitter",
}

STATUS_UP_TO_DATE = "UP_TO_DATE"
STATUS_GO_UPDATE = "GO_MODULE_UPDATE_AVAILABLE"
STATUS_CORE_AHEAD = "UPSTREAM_CORE_AHEAD_GO_BINDING"
STATUS_GRAMMAR_UPDATE = "GRAMMAR_UPDATE_AVAILABLE"
STATUS_UNKNOWN_FETCH = "UNKNOWN_FETCH_FAILED"

ACTIONABLE_STATUSES = {
    STATUS_GO_UPDATE,
    STATUS_GRAMMAR_UPDATE,
    STATUS_UNKNOWN_FETCH,
}


@dataclass(frozen=True)
class ModuleInfo:
    path: str
    version: str
    update: str = ""
    indirect: bool = False


@dataclass(frozen=True)
class ModuleRow:
    module: ModuleInfo
    kind: str
    status: str
    note: str


@dataclass(frozen=True)
class Report:
    markdown: str
    exit_code: int
    rows: tuple[ModuleRow, ...]


def run_command(args: list[str], cwd: pathlib.Path = REPO_ROOT) -> str:
    result = subprocess.run(
        args,
        cwd=str(cwd),
        text=True,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
        check=False,
    )
    if result.returncode != 0:
        stderr = result.stderr.strip()
        raise RuntimeError(f"{' '.join(args)} failed with exit {result.returncode}: {stderr}")
    return result.stdout


def parse_go_list_json_stream(raw: str) -> list[ModuleInfo]:
    decoder = json.JSONDecoder()
    modules: list[ModuleInfo] = []
    index = 0
    while index < len(raw):
        while index < len(raw) and raw[index].isspace():
            index += 1
        if index >= len(raw):
            break
        obj, index = decoder.raw_decode(raw, index)
        update = obj.get("Update") or {}
        modules.append(
            ModuleInfo(
                path=str(obj.get("Path", "")),
                version=str(obj.get("Version", "")),
                update=str(update.get("Version", "")),
                indirect=bool(obj.get("Indirect", False)),
            )
        )
    return modules


def load_go_modules(cwd: pathlib.Path = REPO_ROOT) -> list[ModuleInfo]:
    return parse_go_list_json_stream(run_command(["go", "list", "-m", "-u", "-json", "all"], cwd))


def is_tree_sitter_module(path: str) -> bool:
    return path in TREE_SITTER_EXACT_MODULES or any(path.startswith(prefix) for prefix in TREE_SITTER_MODULE_PREFIXES)


def module_kind(path: str) -> str:
    if path == TREE_SITTER_GO_BINDING:
        return "go-binding"
    name = path.rsplit("/", 1)[-1]
    if name.startswith("tree-sitter-"):
        return "grammar"
    if path.endswith("/go-tree-sitter"):
        return "support"
    return "support"


def parse_semver(value: str) -> tuple[int, int, int] | None:
    match = re.search(r"v?(\d+)\.(\d+)\.(\d+)", value)
    if not match:
        return None
    return int(match.group(1)), int(match.group(2)), int(match.group(3))


def fetch_json(url: str, timeout: int = 8) -> dict | None:
    headers = {
        "Accept": "application/vnd.github+json",
        "User-Agent": "anvien-tree-sitter-drift-check",
    }
    token = os.environ.get("GITHUB_TOKEN")
    if token:
        headers["Authorization"] = f"Bearer {token}"
    try:
        req = urllib.request.Request(url, headers=headers)
        with urllib.request.urlopen(req, timeout=timeout) as resp:
            return json.loads(resp.read().decode("utf-8"))
    except (urllib.error.URLError, urllib.error.HTTPError, json.JSONDecodeError, TimeoutError):
        return None


def latest_tree_sitter_core_tag(fetcher=fetch_json) -> str | None:
    payload = fetcher(TREE_SITTER_RELEASE_API)
    if not payload:
        return None
    tag = payload.get("tag_name")
    return str(tag) if tag else None


def classify_module(module: ModuleInfo, upstream_core_tag: str | None) -> ModuleRow:
    kind = module_kind(module.path)
    if module.update:
        status = STATUS_GRAMMAR_UPDATE if kind == "grammar" else STATUS_GO_UPDATE
        return ModuleRow(module, kind, status, f"update available: {module.update}")

    if module.path == TREE_SITTER_GO_BINDING:
        if upstream_core_tag is None:
            return ModuleRow(module, kind, STATUS_UNKNOWN_FETCH, "could not fetch upstream tree-sitter latest release")
        current = parse_semver(module.version)
        upstream = parse_semver(upstream_core_tag)
        if current and upstream and upstream[:2] > current[:2]:
            return ModuleRow(module, kind, STATUS_CORE_AHEAD, f"upstream core {upstream_core_tag}; no Go module update published")

    return ModuleRow(module, kind, STATUS_UP_TO_DATE, "latest Go module version")


def build_report(modules: list[ModuleInfo], upstream_core_tag: str | None) -> Report:
    tree_sitter_modules = sorted(
        (module for module in modules if is_tree_sitter_module(module.path)),
        key=lambda item: item.path.lower(),
    )
    if not tree_sitter_modules:
        raise RuntimeError("no Tree-sitter Go modules found in go list output")

    rows = tuple(classify_module(module, upstream_core_tag) for module in tree_sitter_modules)
    actionable = [row for row in rows if row.status in ACTIONABLE_STATUSES]
    core_ahead = [row for row in rows if row.status == STATUS_CORE_AHEAD]

    lines: list[str] = [
        "# Tree-sitter Go module drift",
        "",
        f"- Source of truth: `go list -m -u -json all` from `{REPO_ROOT}`",
        f"- Tree-sitter modules found: **{len(rows)}**",
        f"- Actionable drift rows: **{len(actionable)}**",
        f"- Upstream Tree-sitter core latest: `{upstream_core_tag or 'unknown'}`",
        "",
        "## Module Drift",
        "",
        "| Module | Current | Available Update | Kind | Status | Note |",
        "|---|---:|---:|---|---|---|",
    ]

    for row in rows:
        update = row.module.update or "-"
        lines.append(
            f"| `{row.module.path}` | `{row.module.version}` | `{update}` | {row.kind} | {row.status} | {row.note} |"
        )

    lines.extend(
        [
            "",
            "## Upstream Core Delta",
            "",
        ]
    )
    if core_ahead:
        lines.append(
            "Tree-sitter core is newer than the current Go binding module. "
            "This is informational until a matching Go module update is published."
        )
    elif upstream_core_tag is None:
        lines.append("Could not fetch upstream Tree-sitter core latest release.")
    else:
        lines.append("No upstream-core-vs-Go-binding lag was detected.")

    lines.extend(
        [
            "",
            "## Summary",
            "",
        ]
    )
    if actionable:
        lines.append(f"**{len(actionable)} actionable Tree-sitter module drift item(s):**")
        lines.append("")
        for row in actionable:
            lines.append(f"- `{row.module.path}`: {row.status} ({row.note})")
        exit_code = 1
    else:
        lines.append("No actionable Tree-sitter Go module drift was found.")
        exit_code = 0

    return Report("\n".join(lines) + "\n", exit_code, rows)


def main() -> int:
    try:
        modules = load_go_modules(REPO_ROOT)
        upstream_core_tag = latest_tree_sitter_core_tag()
        report = build_report(modules, upstream_core_tag)
    except Exception as exc:
        print("# Tree-sitter Go module drift")
        print()
        print("## Checker Failure")
        print()
        print(f"{type(exc).__name__}: {exc}")
        return 1

    print(report.markdown, end="")
    return report.exit_code


if __name__ == "__main__":
    sys.exit(main())
