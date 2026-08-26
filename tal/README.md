# tal - Incremental Tasker with Lua and Dependencies

[![Go Reference](https://pkg.go.dev/badge/github.com/pt-main/tal.svg)](https://pkg.go.dev/github.com/pt-main/tal)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-yellow.svg)](https://opensource.org/licenses/Apache-2.0)
[![Release](https://img.shields.io/github/v/release/pt-main/tal)](https://github.com/pt-main/tal/releases)

> tal – Task Lua

```bash
go install github.com/pt-main/run/tal/cmd/tal@latest
```

**tal** is a simple, modern task runner that uses Lua as its scripting language. It lets you describe tasks in plain Lua with annotations, track file changes, and run only what has actually changed.

---

## Why tal?

| Problem | How tal solves it |
|---------|-------------------|
| **Makefiles are hard to read and write** | A simple DSL with comments and Lua instead of Shell |
| **Incremental builds work poorly** | SHA256 hashes instead of modification timestamps |
| **No way to call tasks from one another** | Call tasks via a built‑in function |
| **Cumbersome file dependencies** | `#depends file1 file2` works out of the box |

tal gives **incrementality, simplicity, and Lua** – all in one tool.

---

## Installation

```bash
go install github.com/pt-main/run/tal/cmd/tal@latest
```

When you install [run](https://github.com/pt-main/run), tal is installed automatically and becomes available as `run tal ...`.

On first run, `tal update` creates `.tal.pack` – a file containing hashes of all files in the current directory. You can also initialize projects with `tal init`.

---

## Syntax

The task file is written in plain Lua with annotations in comments, without breaking the syntax.

### Basic constructs

| Construct | Description |
|-----------|-------------|
| `-- @taskname` | Start of a task block |
| `-- @` | Main block (runs by default) |
| `-- @!` | Global block (runs before main and task registration) |
| `-- #depends <glob-name...>` | Command – declares file dependencies (checked by their hashes). File names are written in glob format |

Any other code is treated as plain Lua. The file must start with a global or main block.

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
```

---

## CLI Commands

```bash
tal run <file> <args> # Parses <file>, executes the DSL with arguments using .tal.pack (required)
tal update            # Force update or initialize .tal.pack
tal init              # Initialize a project (creates .tal.pack and main.task.lua)
tal list <file>       # List all scripts in the file

# For more information –
tal help
```

`tal update` is mandatory on first use of tal in a directory.

---

## How Incrementality Works

Incrementality is enabled by the `depends` command (`-- #depends ...`) and is disabled if absent.

1. `tal` scans the current directory and computes SHA256 hashes for all files.
2. The hashes are stored in `.tal.pack` (binary format, uses [`pack`](https://github.com/pt-main/pack)).
3. On the next run, `tal` compares hashes, detects which files have changed, and automatically updates the hashes.
4. In the generated Lua script, the array `changed_list` contains paths to the changed files.
5. The runtime checks each task's dependencies and executes only those tasks for which at least one dependent file has changed.

---

## Built-in Lua Runtime

Each task is a Lua function executed in an environment that provides access to the following functions:

- `changed_list` – a table with paths of changed files.
- `tasker.add(deps, name, func)` – registers a task.
- `tasker.run(name)` – executes a task.
- `script(name)` – shorthand for `tasker.run`.
- `shell(string)` – shorthand for `os.execute`.
- `print_colored(string)` – coloured output (uses the colour system from [`tap`](https://github.com/pt-main/tap).color).

When tal is used from `run`, an additional function becomes available – `run(args_string)`, which directly calls run and parses arguments from the input string.

**Important**: You cannot use external Lua libraries (the Lua interpreter in tal is written in [Go](https://github.com/yuin/gopher-lua) and does not depend on the system or installed Lua libraries).

---

## Project Structure

```
.
├── tasks.lua          # Task file (DSL)
├── .tal.pack          # Binary hash file (automatically created by tal update)
└── ...
```

---

## Comparison with Alternatives

| Feature | tal | make | just | task |
|---------|-----|------|------|------|
| **Hash‑based incrementality** | Yes | No | No | Yes |
| **Scripting language** | Lua with annotations | Shell | Shell | Shell |
| **Call other tasks** | Yes | Yes | No | Yes |
| **Ease of writing** | Easy | Hard | Easy | Medium |

---

## License

Apache 2.0 – see [LICENSE](LICENSE) for details.

---

By Pt, 2026 – built with `lc`, `tap`, `pack`.
