# Repository conventions

## One document per problem

Use `problems/algorithms/<id>-<slug>/`, padding IDs to at least four digits. Each problem has three kinds of files:

```text
0001-two-sum/
├── problem.md
├── testcases.json
└── solutions/go/hashmap/
    └── solution.go
```

`problem.md` holds the original URL, difficulty, topics, paraphrased statement, constraints, examples, and your solution notes. There is no separate problem README, metadata file, or attempt log to maintain.

Start the document with:

```markdown
# 1. Two Sum

https://leetcode.com/problems/two-sum/

Difficulty: Easy
Topics: Array, Hash table
```

Use `Easy`, `Medium`, or `Hard`. Topics are a comma-separated list and may be omitted. After the statement, add:

```markdown
## My solutions

| Approach | Trigger | Key idea | Time | Space | Mistakes |
| --- | --- | --- | --- | --- | --- |
| [go/hashmap](solutions/go/hashmap/) |  |  |  |  |  |
```

The trigger is the clue suggesting a technique; the key idea describes how it works. Add one row per implementation. Blank notes are fine. Keep rows on one line, use `<br>` for line breaks, and escape literal pipes as `\|`. Approach links must point to existing `solutions/<language>/<approach>/` directories containing a file.

Use methodology names such as `hashmap`, `sliding-window`, and `stack`. `initial` is acceptable for blank scaffolding. Separate Go approaches use separate packages within the root module. Tools must not edit, format, move, or rename solution source files without the owner's explicit permission.

## Adding a problem

1. Create the problem directory and its `problem.md` using the structure above.
2. Add [testcases.json](testing-and-benchmarks.md) with the example inputs and expected outputs.
3. Add your implementation in `solution.go` and a row linking to it in the solution table. Match the JSON input keys to the function parameter names.
4. Run `go run ./cmd/leetcode test <problem-directory>`; test adapters are generated automatically.
5. Stage your changes and commit. The hook tests the staged repository, then refreshes the root overview if all tests pass. Alternatively, run `python3 scripts/consolidate.py` before staging to preview the overview.

Custom tests go in the same `testcases.json`; no template directory, second fixture file, or manually written test adapter is needed. Comparison defaults to exact equality; specify `Comparison: unordered` in `problem.md` when only output order may differ. See the [runner documentation](testing-and-benchmarks.md) for commands and supported signatures.

## Generated overview and hook

`scripts/consolidate.py` reads `problem.md` files and updates only the marked section of the root README. It generates counts, a problem index, and the combined solution table. It never reads or modifies solution source contents. Counts describe documented problems and solution entries, not correctness or acceptance.

Commands require Python 3 and Git:

- `python3 scripts/consolidate.py`: regenerate from working-tree documents.
- `python3 scripts/consolidate.py --check`: report stale output without writing.
- `python3 scripts/consolidate.py --staged`: regenerate from the Git index and stage only the changed root README. This is the hook's mode.
- `python3 -m unittest discover -s scripts/tests`: test generation and staged-file safeguards.

The tracked hook is `.githooks/pre-commit`, enabled for this clone. Enable it in another clone with `git config --local core.hooksPath .githooks`; hooks are not enabled automatically by cloning.

The hook runs `scripts/precommit.py`. It copies the Git index to a temporary directory, generates adapters there, runs every Go test (solution cases and shared runner tests), and runs the Python repository automation tests. A failing test, compilation error, or generation error blocks the commit. It checks that the staged file list and object hashes have not changed during validation; if they have, retry the commit. Only after successful testing does it regenerate and stage the README. The temporary directory is removed afterward, including its generated adapters. Your solution files and unstaged edits are never copied back or modified. Local hooks can be bypassed; CI enforcement remains future work.

The hook reads staged documents so partially staged notes remain partially staged. It refuses to overwrite unstaged root README edits. Stage those edits or restore the affected README before retrying. It never stages your problem documents, fixtures, solutions, or unrelated files.

Root README text outside the generated markers remains handwritten. The `--check` command can later be used in CI.

## Future categories

Other languages can share the same problem document and fixtures. SQL and system design can later use sibling categories under `problems/`, with their own conventions. They are outside the current scope.
