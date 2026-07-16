#!/usr/bin/env python3
from __future__ import annotations

import importlib.util
import pathlib
import sys
import unittest


SCRIPT_PATH = pathlib.Path(__file__).with_name("check-tree-sitter-upgrade-readiness.py")
SPEC = importlib.util.spec_from_file_location("tree_sitter_drift", SCRIPT_PATH)
assert SPEC and SPEC.loader
tree_sitter_drift = importlib.util.module_from_spec(SPEC)
sys.modules[SPEC.name] = tree_sitter_drift
SPEC.loader.exec_module(tree_sitter_drift)


class TreeSitterDriftTests(unittest.TestCase):
    def test_parse_go_list_stream_extracts_updates(self) -> None:
        raw = """
{"Path":"github.com/tamnguyendinh/anvien"}
{"Path":"github.com/tree-sitter/go-tree-sitter","Version":"v0.25.0"}
{"Path":"github.com/UserNobody14/tree-sitter-dart","Version":"v0.0.0-20260508020638-507c5546dc73","Update":{"Version":"v0.0.0-20260707040301-be07cf7118d3"},"Indirect":true}
"""
        modules = tree_sitter_drift.parse_go_list_json_stream(raw)

        self.assertEqual(len(modules), 3)
        self.assertEqual(modules[1].path, "github.com/tree-sitter/go-tree-sitter")
        self.assertEqual(modules[2].update, "v0.0.0-20260707040301-be07cf7118d3")
        self.assertTrue(modules[2].indirect)

    def test_core_ahead_is_informational_when_no_go_update_exists(self) -> None:
        modules = [
            tree_sitter_drift.ModuleInfo("github.com/tree-sitter/go-tree-sitter", "v0.25.0"),
            tree_sitter_drift.ModuleInfo("github.com/tree-sitter/tree-sitter-go", "v0.25.0"),
        ]

        report = tree_sitter_drift.build_report(modules, "v0.26.11")
        statuses = {row.module.path: row.status for row in report.rows}

        self.assertEqual(statuses["github.com/tree-sitter/go-tree-sitter"], tree_sitter_drift.STATUS_CORE_AHEAD)
        self.assertEqual(statuses["github.com/tree-sitter/tree-sitter-go"], tree_sitter_drift.STATUS_UP_TO_DATE)
        self.assertEqual(report.exit_code, 0)
        self.assertIn("informational", report.markdown)

    def test_grammar_update_is_actionable(self) -> None:
        modules = [
            tree_sitter_drift.ModuleInfo("github.com/tree-sitter/go-tree-sitter", "v0.25.0"),
            tree_sitter_drift.ModuleInfo(
                "github.com/UserNobody14/tree-sitter-dart",
                "v0.0.0-20260508020638-507c5546dc73",
                "v0.0.0-20260707040301-be07cf7118d3",
                True,
            ),
        ]

        report = tree_sitter_drift.build_report(modules, "v0.26.11")
        statuses = {row.module.path: row.status for row in report.rows}

        self.assertEqual(statuses["github.com/UserNobody14/tree-sitter-dart"], tree_sitter_drift.STATUS_GRAMMAR_UPDATE)
        self.assertEqual(report.exit_code, 1)
        self.assertIn("GRAMMAR_UPDATE_AVAILABLE", report.markdown)

    def test_missing_upstream_fetch_is_visible_failure(self) -> None:
        modules = [tree_sitter_drift.ModuleInfo("github.com/tree-sitter/go-tree-sitter", "v0.25.0")]

        report = tree_sitter_drift.build_report(modules, None)

        self.assertEqual(report.rows[0].status, tree_sitter_drift.STATUS_UNKNOWN_FETCH)
        self.assertEqual(report.exit_code, 1)
        self.assertIn("unknown", report.markdown)

    def test_no_npm_package_dependencies_are_required(self) -> None:
        modules = [tree_sitter_drift.ModuleInfo("github.com/tree-sitter/go-tree-sitter", "v0.25.0")]

        report = tree_sitter_drift.build_report(modules, "v0.25.0")

        self.assertIn("github.com/tree-sitter/go-tree-sitter", report.markdown)
        self.assertNotIn("dependencies", report.markdown)


if __name__ == "__main__":
    unittest.main()
