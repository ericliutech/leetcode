from pathlib import Path
import shutil
import subprocess
import sys
import tempfile
import unittest


sys.path.insert(0, str(Path(__file__).resolve().parents[1]))
import consolidate as generator


class ConsolidateTests(unittest.TestCase):
    def setUp(self):
        self.temporary = tempfile.TemporaryDirectory()
        self.addCleanup(self.temporary.cleanup)
        self.root = Path(self.temporary.name)
        self.git("init", "-q")
        self.write(
            "README.md",
            f"# My solutions\n\nHandwritten introduction.\n\n{generator.START}\n"
            f"{generator.END}\n\nHandwritten footer.\n",
        )
        self.folder = "problems/algorithms/0001-two-sum"
        self.notes = (
            "# 1. Two Sum\n\nhttps://leetcode.com/problems/two-sum/\n\n"
            "Difficulty: Easy\nTopics: Array, Hash table\n\n"
            "## Examples\n\n| Input | Output |\n| --- | --- |\n| Example | Result |\n\n"
            "## My solutions\n\n"
            "| Approach | Trigger | Key idea | Time | Space | Mistakes |\n"
            "| --- | --- | --- | --- | --- | --- |\n"
            "| [go/hashmap](solutions/go/hashmap/) | Pair sum | STAGED NOTE | O(n) | O(n) | a \\| b |\n"
        )
        self.write(f"{self.folder}/problem.md", self.notes)
        self.source_path = f"{self.folder}/solutions/go/hashmap/solution.go"
        self.write(self.source_path, "Solution contents are opaque to this tool.\n")

    def write(self, path, content):
        destination = self.root / path
        destination.parent.mkdir(parents=True, exist_ok=True)
        destination.write_text(content, encoding="utf-8")

    def git(self, *args):
        result = subprocess.run(
            ["git", "-C", str(self.root), *args],
            capture_output=True, text=True,
        )
        self.assertEqual(result.returncode, 0, result.stderr)
        return result.stdout

    def generate(self, staged=False, check=False):
        return generator.consolidate(generator.Repository(self.root, staged=staged), check=check)

    def test_working_generation_is_idempotent_and_preserves_handwritten_text(self):
        source = (self.root / self.source_path).read_bytes()
        self.assertEqual(self.generate(), ["README.md"])
        readme = (self.root / "README.md").read_text()
        self.assertIn("Handwritten introduction.", readme)
        self.assertIn("Handwritten footer.", readme)
        overview = readme
        self.assertIn("O(n)", overview)
        self.assertIn(r"a \| b", overview)
        self.assertIn(f"{self.folder}/solutions/go/hashmap/", overview)
        self.assertIn(f"{self.folder}/problem.md", overview)
        self.assertFalse((self.root / "docs/attempts.md").exists())
        self.assertEqual(self.generate(), [])
        self.assertEqual((self.root / self.source_path).read_bytes(), source)

    def test_check_rejects_stale_output_without_writing(self):
        before = (self.root / "README.md").read_bytes()
        with self.assertRaisesRegex(ValueError, "stale"):
            self.generate(check=True)
        self.assertEqual((self.root / "README.md").read_bytes(), before)
        self.generate()
        self.assertEqual(self.generate(check=True), [])

    def test_staged_generation_excludes_unstaged_notes_and_preserves_sources(self):
        self.git("add", ".")
        self.write(f"{self.folder}/problem.md", self.notes.replace("STAGED NOTE", "UNSTAGED NOTE"))
        self.write(self.source_path, "Unstaged solution changes must survive.\n")
        source = (self.root / self.source_path).read_bytes()
        self.generate(staged=True)
        overview = self.git("show", ":README.md")
        self.assertIn("STAGED NOTE", overview)
        self.assertNotIn("UNSTAGED NOTE", overview)
        self.assertIn("STAGED NOTE", self.git("show", f":{self.folder}/problem.md"))
        self.assertIn("UNSTAGED NOTE", (self.root / self.folder / "problem.md").read_text())
        self.assertEqual((self.root / self.source_path).read_bytes(), source)
        self.assertEqual(self.git("show", f":{self.source_path}"), "Solution contents are opaque to this tool.\n")
        self.assertEqual(self.generate(staged=True), [])

    def test_unstaged_generated_changes_block_all_writes(self):
        self.git("add", ".")
        self.write("README.md", (self.root / "README.md").read_text() + "Handwritten unstaged change\n")
        root_before = (self.root / "README.md").read_bytes()
        index_before = self.git("show", ":README.md")
        with self.assertRaisesRegex(ValueError, "unstaged changes in README.md"):
            self.generate(staged=True)
        self.assertEqual((self.root / "README.md").read_bytes(), root_before)
        self.assertEqual(self.git("show", ":README.md"), index_before)

    def test_missing_solution_link_is_rejected_without_writes(self):
        self.write(f"{self.folder}/problem.md", self.notes.replace("solutions/go/hashmap/", "solutions/go/missing/"))
        before = (self.root / "README.md").read_bytes()
        with self.assertRaisesRegex(ValueError, "No solution files"):
            self.generate()
        self.assertEqual((self.root / "README.md").read_bytes(), before)

    def test_duplicate_approach_rows_are_rejected(self):
        self.write(f"{self.folder}/problem.md", self.notes + self.notes.splitlines()[-1] + "\n")
        with self.assertRaisesRegex(ValueError, "Duplicate approach"):
            self.generate()

    def test_no_staged_documentation_is_a_noop(self):
        self.assertEqual(self.generate(staged=True), [])
        self.assertNotIn("## Solution notes", (self.root / "README.md").read_text())

    def test_missing_difficulty_is_rejected_without_writes(self):
        self.write(f"{self.folder}/problem.md", self.notes.replace("Difficulty: Easy", "Difficulty: Unknown"))
        before = (self.root / "README.md").read_bytes()
        with self.assertRaisesRegex(ValueError, "Difficulty"):
            self.generate()
        self.assertEqual((self.root / "README.md").read_bytes(), before)

    def test_topics_can_be_omitted(self):
        self.write(f"{self.folder}/problem.md", self.notes.replace("Topics: Array, Hash table\n", ""))
        self.generate()
        self.assertIn("| Easy |  | go |", (self.root / "README.md").read_text())

    def test_hook_refreshes_documents_during_an_isolated_commit(self):
        self.write("scripts/.keep", "")
        shutil.copyfile(generator.__file__, self.root / "scripts/consolidate.py")
        self.git("add", ".")
        hook_directory = Path(generator.__file__).resolve().parent.parent / ".githooks"
        self.git(
            "-c", f"core.hooksPath={hook_directory}",
            "-c", "user.name=Generator tests", "-c", "user.email=tests@example.invalid",
            "commit", "-q", "-m", "Test generated documentation hook",
        )
        self.assertIn("STAGED NOTE", self.git("show", "HEAD:README.md"))
        self.assertIn("| Documented problems | 1 |", self.git("show", "HEAD:README.md"))
        self.assertEqual(self.git("status", "--porcelain"), "")


if __name__ == "__main__":
    unittest.main()
