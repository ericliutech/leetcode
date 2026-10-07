import os
from pathlib import Path
import shutil
import subprocess
import tempfile
import unittest


SOURCE_ROOT = Path(__file__).resolve().parents[2]


class PrecommitTests(unittest.TestCase):
    def setUp(self):
        self.temporary = tempfile.TemporaryDirectory()
        self.addCleanup(self.temporary.cleanup)
        self.root = Path(self.temporary.name)
        self.environment = os.environ.copy()
        self.environment.update({"GIT_CONFIG_NOSYSTEM": "1", "GIT_CONFIG_GLOBAL": os.devnull})
        self.git("init", "-q")
        # Copy the real runner and its tests. Exclude this integration test file
        # from the fixture to avoid recursively testing hooks inside hooks.
        paths = [
            "go.mod", "scripts/consolidate.py", "scripts/precommit.py",
            "scripts/tests/test_consolidate.py", ".githooks/pre-commit",
        ]
        paths.extend(path.relative_to(SOURCE_ROOT).as_posix() for directory in ("cmd/leetcode", "internal/judge")
                     for path in (SOURCE_ROOT / directory).glob("*.go"))
        for relative in paths:
            destination = self.root / relative
            destination.parent.mkdir(parents=True, exist_ok=True)
            shutil.copyfile(SOURCE_ROOT / relative, destination)
        (self.root / ".githooks/pre-commit").chmod(0o755)
        self.write("README.md", "# Test repository\n\n<!-- generated:progress:start -->\n<!-- generated:progress:end -->\n")
        self.folder = "problems/algorithms/0099-example"
        self.solution = f"{self.folder}/solutions/go/example/solution.go"
        self.write(self.solution, "package solution\nfunc value(x int) int { return x }\n")
        self.write(f"{self.folder}/problem.md", (
            "# 99. Example\n\nDifficulty: Easy\nTopics: Testing\n\n## My solutions\n\n"
            "| Approach | Trigger | Key idea | Time | Space | Mistakes |\n"
            "| --- | --- | --- | --- | --- | --- |\n"
            "| [go/example](solutions/go/example/) |  |  |  |  |  |\n"
        ))
        self.cases = f"{self.folder}/testcases.json"
        self.write(self.cases, '[{"id":1,"input":{"x":7},"expected":7}]\n')

    def write(self, path, content):
        destination = self.root / path
        destination.parent.mkdir(parents=True, exist_ok=True)
        destination.write_text(content)

    def command(self, args):
        return subprocess.run(args, cwd=self.root, env=self.environment, capture_output=True, text=True)

    def git(self, *arguments):
        result = self.command(["git", *arguments])
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        return result.stdout

    def commit(self):
        return self.command([
            "git", "-c", "core.hooksPath=.githooks", "-c", "user.name=Hook tests",
            "-c", "user.email=tests@example.invalid", "commit", "-q", "-m", "Test staged hook",
        ])

    def test_success_uses_staged_source_and_preserves_unstaged_changes(self):
        self.git("add", ".")
        original = (self.root / self.solution).read_text()
        self.write(self.solution, original.replace("return x", "return -1"))
        working_source = (self.root / self.solution).read_bytes()
        result = self.commit()
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        self.assertEqual(self.git("show", f"HEAD:{self.solution}"), original)
        self.assertEqual((self.root / self.solution).read_bytes(), working_source)
        self.assertIn("| Documented problems | 1 |", self.git("show", "HEAD:README.md"))
        self.assertFalse(list((self.root / "problems").rglob("zz_generated_test.go")))

    def test_failing_staged_case_blocks_commit_despite_unstaged_fix(self):
        good = (self.root / self.cases).read_text()
        self.write(self.cases, good.replace('"expected":7', '"expected":99'))
        self.git("add", ".")
        self.write(self.cases, good)
        before_index = self.git("ls-files", "--stage")
        before_readme = (self.root / "README.md").read_bytes()
        result = self.commit()
        self.assertNotEqual(result.returncode, 0)
        self.assertIn("commit stopped", result.stderr)
        self.assertIn("want 99", result.stdout + result.stderr)
        self.assertEqual(self.git("ls-files", "--stage"), before_index)
        self.assertEqual((self.root / "README.md").read_bytes(), before_readme)
        self.assertEqual((self.root / self.cases).read_text(), good)
        self.assertNotEqual(self.command(["git", "rev-parse", "--verify", "HEAD"]).returncode, 0)

    def test_failing_shared_runner_test_blocks_commit(self):
        self.write("internal/judge/failure_test.go", 'package judge\nimport "testing"\nfunc TestDeliberateFailure(t *testing.T) { t.Fatal("runner gate works") }\n')
        self.git("add", ".")
        result = self.commit()
        self.assertNotEqual(result.returncode, 0)
        self.assertIn("runner gate works", result.stdout + result.stderr)
        self.assertIn("commit stopped", result.stderr)
        self.assertNotEqual(self.command(["git", "rev-parse", "--verify", "HEAD"]).returncode, 0)


if __name__ == "__main__":
    unittest.main()
