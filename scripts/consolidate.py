#!/usr/bin/env python3
"""Generate the problem index and approach overview without reading solution code."""

import argparse
from pathlib import Path, PurePosixPath
import re
import subprocess
import sys


START = "<!-- generated:progress:start -->"
END = "<!-- generated:progress:end -->"
HEADERS = ["Approach", "Trigger", "Key idea", "Time", "Space", "Mistakes"]
PROBLEM_FILE = re.compile(r"problems/algorithms/\d{4,}-[^/]+/problem\.md")


class Repository:
    def __init__(self, root, staged=False):
        self.root = Path(root)
        self.staged = staged
        if staged:
            self.paths = set(self.git("ls-files", "-z").split("\0")) - {""}
        else:
            self.paths = {
                path.relative_to(self.root).as_posix()
                for base in (self.root / "problems",)
                for path in base.rglob("*")
                if path.is_file()
            }
            self.paths.update(
                name for name in ("README.md",)
                if (self.root / name).is_file()
            )

    def git(self, *args):
        result = subprocess.run(
            ["git", "-C", str(self.root), *args],
            capture_output=True, text=True, check=False,
        )
        if result.returncode:
            raise ValueError(result.stderr.strip() or "Git command failed")
        return result.stdout

    def read(self, path):
        if path not in self.paths:
            raise ValueError(f"Missing source file: {path}")
        if self.staged:
            return self.git("show", f":{path}")
        return (self.root / path).read_text(encoding="utf-8")


def cells(line):
    line = line.strip()
    if not (line.startswith("|") and line.endswith("|")):
        raise ValueError("Notes tables must start and end each row with a pipe")
    return [cell.strip() for cell in re.split(r"(?<!\\)\|", line[1:-1])]


def read_notes(repository, path):
    lines = repository.read(path).splitlines()
    header = next(
        (index for index, line in enumerate(lines)
         if line.strip().startswith("|") and cells(line) == HEADERS),
        None,
    )
    if header is None or header + 1 >= len(lines):
        raise ValueError(f"Missing six-column solution table in {path}")
    separator = cells(lines[header + 1])
    if len(separator) != len(HEADERS) or not all(re.fullmatch(r":?-{3,}:?", cell) for cell in separator):
        raise ValueError(f"Invalid table separator in {path}")
    rows = []
    targets = set()
    folder = PurePosixPath(path).parent
    for line in lines[header + 2:]:
        if not line.strip().startswith("|"):
            break
        row = cells(line)
        if len(row) != len(HEADERS):
            raise ValueError(f"Expected six cells in {path}; escape literal pipes as \\|")
        link = re.fullmatch(
            r"\[([^\]]+)\]\((solutions/[a-z0-9-]+/[a-z0-9-]+)/?\)", row[0]
        )
        if not link:
            raise ValueError(f"Approach must link to solutions/<language>/<approach> in {path}")
        label, target = link.groups()
        if target in targets:
            raise ValueError(f"Duplicate approach link in {path}: {target}")
        targets.add(target)
        prefix = f"{folder}/{target}/"
        if not any(source.startswith(prefix) for source in repository.paths):
            raise ValueError(f"No solution files found at {prefix}")
        rows.append((label, target, row[1:]))
    return rows


def render(repository):
    problem_paths = sorted(path for path in repository.paths if PROBLEM_FILE.fullmatch(path))
    if repository.staged and not problem_paths and "README.md" not in repository.paths:
        return {}
    problems = []
    for path in problem_paths:
        folder = PurePosixPath(path).parent
        statement = repository.read(path)
        heading = re.search(r"^# (\d+)\. (.+)$", statement, re.MULTILINE)
        if not heading:
            raise ValueError(f"Expected '# <ID>. <Title>' in {folder}/problem.md")
        identifier, title = heading.groups()
        if int(identifier) != int(folder.name.split("-", 1)[0]):
            raise ValueError(f"Problem ID does not match directory: {folder}")
        difficulty = re.search(r"^Difficulty: (Easy|Medium|Hard)\s*$", statement, re.MULTILINE)
        if not difficulty:
            raise ValueError(f"Expected 'Difficulty: Easy', 'Medium', or 'Hard' in {path}")
        topics = re.search(r"^Topics:[ \t]*(.*)$", statement, re.MULTILINE)
        facts = {"difficulty": difficulty.group(1), "topics": topics.group(1).strip() if topics else ""}
        problems.append((int(identifier), title, folder, facts, read_notes(repository, path)))
    problems.sort(key=lambda problem: problem[0])

    def escape(value):
        return str(value).replace("|", r"\|").replace("\n", " ")

    progress = [
        "## Progress", "", "| Measure | Count |", "| --- | ---: |",
        f"| Documented problems | {len(problems)} |",
        f"| Solution entries | {sum(len(problem[4]) for problem in problems)} |",
        "", "Counts reflect documentation, not correctness or LeetCode acceptance.",
        "", "## Problem index", "",
        "| ID | Problem | Difficulty | Topics | Languages |",
        "| --- | --- | --- | --- | --- |",
    ]
    overview = [
        "", "## Solution notes", "",
        "| Problem | Approach | Trigger | Key idea | Time | Space | Mistakes |",
        "| --- | --- | --- | --- | --- | --- | --- |",
    ]
    for identifier, title, folder, metadata, rows in problems:
        languages = ", ".join(sorted({target.split("/")[1] for _, target, _ in rows}))
        topics = escape(metadata["topics"])
        progress.append(
            f"| {identifier} | [{escape(title)}]({folder}/problem.md) | "
            f"{metadata['difficulty']} | {topics} | {languages} |"
        )
        for label, target, notes in rows:
            columns = [
                f"[{identifier}. {escape(title)}]({folder}/problem.md)",
                f"[{escape(label)}]({folder}/{target}/)", *notes,
            ]
            overview.append("| " + " | ".join(columns) + " |")
    readme = repository.read("README.md")
    if readme.count(START) != 1 or readme.count(END) != 1:
        raise ValueError("Root README must contain exactly one generated progress marker pair")
    before, remainder = readme.split(START, 1)
    _, after = remainder.split(END, 1)
    return {
        "README.md": before + START + "\n\n" + "\n".join(progress + overview) + "\n\n" + END + after,
    }


def consolidate(repository, check=False):
    outputs = render(repository)
    changes = {
        path: content for path, content in outputs.items()
        if path not in repository.paths or repository.read(path) != content
    }
    if check:
        if changes:
            raise ValueError("Generated documents are stale: " + ", ".join(changes))
        return []
    if repository.staged:
        # Check every output before making any changes, including unchanged outputs.
        for path in outputs:
            destination = repository.root / path
            staged = repository.read(path) if path in repository.paths else None
            working = destination.read_text(encoding="utf-8") if destination.exists() else None
            if working != staged:
                raise ValueError(
                    f"Refusing to overwrite unstaged changes in {path}. "
                    "Stage that file or restore it before committing."
                )
    for path, content in changes.items():
        destination = repository.root / path
        destination.parent.mkdir(parents=True, exist_ok=True)
        destination.write_text(content, encoding="utf-8")
    if repository.staged and changes:
        repository.git("add", "--", *changes)
    return list(changes)


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--staged", action="store_true", help="Read staged files and stage generated outputs")
    parser.add_argument("--check", action="store_true", help="Check generated output without writing files")
    args = parser.parse_args()
    root = Path(__file__).resolve().parent.parent
    try:
        changed = consolidate(Repository(root, staged=args.staged), check=args.check)
    except (ValueError, OSError) as error:
        print(f"consolidate: {error}", file=sys.stderr)
        return 1
    print("Updated: " + ", ".join(changed) if changed else "Generated documents are current.")
    return 0


if __name__ == "__main__":
    sys.exit(main())
