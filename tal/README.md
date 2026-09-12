# tal — incremental tasker with Lua and dependencies

[![Go Reference](https://pkg.go.dev/badge/github.com/pt-main/tal.svg)](https://pkg.go.dev/github.com/pt-main/tal)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-yellow.svg)](https://opensource.org/licenses/Apache-2.0)
[![Release](https://img.shields.io/github/v/release/pt-main/tal)](https://github.com/pt-main/tal/releases)

> tal — Task Lua

```bash
go install github.com/pt-main/run/cmd/tal@latest
```

**tal** is a simple, modern task runner that uses Lua as its scripting language. It lets you describe tasks in plain Lua with annotations, track file changes, and run only what has actually changed.

---

## Why tal?

| Problem | tal solves |
|---------|------------|
| **Makefiles are hard to read and write** | Simple DSL with comments and Lua instead of Shell |
| **Incremental builds are flaky** | SHA256 hashes instead of modification times |
| **No way to call tasks from each other** | Built‑in function to invoke tasks |
| **File dependencies are verbose** | `#depends file1 file2` works out of the box |

tal gives you **incrementality, simplicity, and Lua** – all in one tool.

---

## Installation

```bash
go install github.com/pt-main/run/tal/cmd/tal@latest
```

When you install [run](https://github.com/pt-main/run), tal is also installed automatically and becomes available as `run tal ...`.

On first run, `tal update` creates `.tal.pack` – a file containing hashes of all files in the current directory. You can also initialise projects with `tal init`.

---

## Syntax

Task files are written in plain Lua with annotations in comments, without breaking Lua syntax.

### Main constructs

| Construct | Description |
|-----------|-------------|
| `-- @taskname` | Start of a task block |
| `-- @` | Main block (runs by default) |
| `-- @!` | Global block (runs before main and task registration) |
| `-- #depends <glob‑pattern...>` | Command – dependency on files (checked by their hashes). File names are written in glob format. |

Any other code is ordinary Lua. The file must start with a global/main block.

Example:

```lua
-- @!
-- Global code: variables, functions, imports
local function log(msg) print("[TASK] " .. msg) end

-- @build
-- #depends *.go
log("Building...")
os.execute("go build .")

-- @test
-- #depends test/**
log("Testing...")
os.execute("go test .")

-- @
-- Runs by default
script("build")
script("test")

update() -- updates .tal.pack, confirms that all changes are processed
```

---

## CLI Commands

```bash
tal run <file> <args> # parses <file>, executes the DSL with arguments using .tal.pack (required)
tal update            # forcibly update or initialise .tal.pack
tal init              # initialise the project (creates .tal.pack and main.task.lua)
tal list <file>       # show all scripts in the file

# for more information -
tal help
```

`tal run ...` supports the `--deps="..."` flag, which allows passing dependencies for the `-- #depends` annotation to the script. If the flag is omitted, the script operates using `.tal.pack`.

---

## How incrementality works

Incrementality is enabled by the `depends` command (`-- #depends ...`) and does not work without it.

1. `tal` scans the current directory and computes SHA256 hashes for all files.
2. Hashes are saved in `.tal.pack` (binary format, using [`pack`](https://github.com/pt-main/pack)).
3. On the next run, `tal` compares hashes, detects which files changed, and automatically updates the hashes.
4. In the generated Lua script, the `changed_list` array contains the paths of changed files.
5. The runtime checks dependencies for each task and only runs those for which at least one dependent file has changed.

You must call the `update()` function to refresh `.tal.pack`.

---

## Built‑in Lua runtime

Each task is a Lua function executed in an environment with access to these functions:

- `changed_list` – a table with paths of changed files.
- `tasker.add(deps, name, func)` – registers a task.
- `tasker.run(name)` – runs a task.
- `script(name)` – shorthand for `tasker.run`.
- `shell(string)` – shorthand for `os.execute`.
- `print_colored(string)` – coloured output (uses the colour system from <code>[tap](https://github.com/pt-main/tap).color</code>).
- `update()` – recalculates hashes and fully updates `.tal.pack`.
- `match_pattern(glob, path)` – checks whether the path matches the glob pattern (supports doublestar syntax).

When tal is used from `run cli`, an additional function becomes available – `run_cli(args_string)` – which directly calls run and parses arguments from the input string.

**Important**: You cannot use external Lua libraries (the Lua interpreter in tal is written in [Go](https://github.com/yuin/gopher-lua) and does not depend on the system or installed Lua libraries).

---

## Project structure

```
.
├── tasks.lua          # task file (DSL)
├── .tal.pack          # binary hash file (created automatically by tal update)
└── ...
```

---

## Comparison with alternatives

| Feature | tal | make | just | task |
|---------|-----|------|------|------|
| **Hash‑based incrementality** | Yes | No | No | Yes |
| **Scripting language** | Lua with annotations | Shell | Shell | Shell |
| **Calling other tasks** | Yes | Yes | No | Yes |
| **Ease of writing** | Easy | Hard | Easy | Medium |

---

## License

Apache 2.0 – see [LICENSE](LICENSE) for details.

---

By Pt, 2026 – built using `lc`, `tap`, `pack`.