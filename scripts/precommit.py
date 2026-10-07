#!/usr/bin/env python3
"""Test the staged repository in isolation before regenerating its README."""

import os
from pathlib import Path
import subprocess
import sys
import tempfile

import consolidate


def git(root, *arguments):
    result = subprocess.run(
        ["git", "-C", str(root), *arguments],
        capture_output=True, text=True, check=True,
    )
    return result.stdout


def validate_staged(root):
    root = Path(root)
    original_index = git(root, "ls-files", "--stage", "-z")
    # Git hooks export repository-local variables. Do not leak them into the
    # snapshot's tests, which may create their own temporary Git repositories.
    environment = os.environ.copy()
    for name in git(root, "rev-parse", "--local-env-vars").splitlines():
        environment.pop(name, None)
    with tempfile.TemporaryDirectory(prefix="leetcode-staged-") as temporary:
        snapshot = Path(temporary)
        git(root, "checkout-index", "--all", f"--prefix={snapshot}{os.sep}")
        if git(root, "ls-files", "--stage", "-z") != original_index:
            raise ValueError("staged files changed while creating the test snapshot; retry the commit")
        for required in ("go.mod", "cmd/leetcode/main.go", "internal/judge/judge.go", "scripts/consolidate.py"):
            if not (snapshot / required).is_file():
                raise ValueError(f"required file is not staged: {required}")
        checks = [
            ("Generate solution adapters", ["go", "run", "./cmd/leetcode", "generate"]),
            ("Go tests (solutions and shared runner)", ["go", "test", "-count=1", "-timeout=60s", "./..."]),
            ("Repository automation tests", [sys.executable, "-B", "-m", "unittest", "discover", "-s", "scripts/tests"]),
        ]
        for label, command in checks:
            print(f"pre-commit: {label} (staged files)", flush=True)
            subprocess.run(command, cwd=snapshot, env=environment, check=True)
    if git(root, "ls-files", "--stage", "-z") != original_index:
        raise ValueError("staged files changed during testing; retry the commit")


def main():
    try:
        root = Path(git(Path.cwd(), "rev-parse", "--show-toplevel").strip())
        validate_staged(root)
        changed = consolidate.consolidate(consolidate.Repository(root, staged=True))
        if changed:
            print("pre-commit: generated and staged " + ", ".join(changed))
        print("pre-commit: all checks passed")
    except subprocess.CalledProcessError as error:
        if error.stderr:
            print(error.stderr, file=sys.stderr)
        print("pre-commit: checks failed; commit stopped", file=sys.stderr)
        return 1
    except (OSError, ValueError) as error:
        print(f"pre-commit: {error}; commit stopped", file=sys.stderr)
        return 1
    return 0


if __name__ == "__main__":
    sys.exit(main())
