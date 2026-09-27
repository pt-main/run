# tal - incremental task runner with Lua and dependencies

[![Go Reference](https://pkg.go.dev/badge/github.com/pt-main/tal.svg)](https://pkg.go.dev/github.com/pt-main/tal)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-yellow.svg)](https://opensource.org/licenses/Apache-2.0)
[![Release](https://img.shields.io/github/v/release/pt-main/tal)](https://github.com/pt-main/tal/releases)

> tal - Task Lua

```bash
go install github.com/pt-main/run/cmd/tal@latest
```

**tal** is a simple modern task runner with Lua as its scripting language. It lets you describe tasks in plain Lua with annotations, track file changes, and run only what has actually changed.

---

## Why tal?

| Problem | tal solves |
|----------|------------|
| **Makefiles are hard to read and write** | Simple DSL with comments and Lua instead of Shell |
| **Incrementality works poorly** | SHA256 hashes instead of modification time |
| **No calling tasks from one another** | Tasks can be called via a built-in function |
| **File dependencies are cumbersome** | `-- #depends file1 file2` works out of the box |

tal gives you **incrementality, simplicity, and Lua** - all in one tool.

---

## Installation

```bash
go install github.com/pt-main/run/cmd/tal@latest
```

Also, when [run](https://github.com/pt-main/run) is installed, tal is installed automatically and is available as `run tal ...`.

On first run, `tal update` will create `.tal.pack` - a file with hashes of files in the current directory. You can initialize a project (create `.tal.pack` and `main.task.lua`) with the `tal init` command.

---

## Syntax

The task file is written in plain Lua with annotations in comments, without breaking the syntax.

### Basic constructs

| Construct | Description |
|-------------|-------------|
| `-- @taskname` | Start of a task block |
| `-- @` | Main block (runs by default) |
| `-- @!` | Global block (executed before main and task registration) |
| `-- #depends <glob...>` | File dependency command (checked by hashes). Paths are specified in glob format |

Any other code is considered regular Lua code. The file must start with a global or regular block.

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

update() -- updates .tal.pack, confirming that all changes have been processed
```

---

## CLI commands

| Command | Description | Example |
|---------|-------------|---------|
| `tal run <file> [args...] [--deps="..."]` | Parses `<file>` and executes the DSL with arguments, using `.tal.pack` (required if `--deps` is not passed) | `tal run main.task.lua build` |
| `tal update` | Update or force-initialize `.tal.pack` | `tal update` |
| `tal init` | Initialize a project (creates `.tal.pack` and `main.task.lua`) | `tal init` |
| `tal list <file> [file...]` | Show all tasks in the file (including Global and Main) | `tal list main.task.lua` |

To get more information:

```bash
tal help
```

### The `--deps` flag

`tal run ...` supports the `--deps="dep1;dep2"` flag, which allows passing dependencies to the script for the `-- #depends` annotation (separator is `;`). If the flag is absent, the script works with `.tal.pack`.

```bash
tal run main.task.lua build --deps="main.go;go.mod"
```

Arguments after the file name are passed directly to Lua via `get_args()` and are not interpreted by the CLI.

---

## How incrementality works

Incrementality is enabled by the `depends` command (`-- #depends ...`) and does not work without it.

1. `tal` scans the current directory and computes SHA256 for all files.
2. Hashes are saved in `.tal.pack` (binary format, uses [`pack`](https://github.com/pt-main/pack)).
3. On the next run, `tal` compares hashes, determines which files changed, and automatically updates the hashes.
4. In the generated Lua script, the `changed_list` array contains paths to changed files.
5. The runtime checks each task's dependencies and executes only those for which at least one dependent file has changed.

To update `.tal.pack`, you must use the `update()` function.

---

## Built-in Lua runtime

Each task is a Lua function that runs in an environment with access to the functions:

- `changed_list` - table with paths of changed files (relative to the current directory).
- `get_args()` - table of arguments passed to `tal run`.
- `script(name)` - executing a script.
- `shell(string)` - shorthand for `os.execute`.
- `print_colored(string)` - colored output (uses the color system from [`tap`](https://github.com/pt-main/tap).color).
- `update()` - recalculates hashes and completely updates `.tal.pack`.
- `match_pattern(glob, path)` - checks whether path `path` matches the `glob` pattern (doublestar syntax is supported).

When tal is used from `run cli`, an additional function becomes available - `run_cli(args_string)`, which directly calls run and parses arguments from the input string.

**Important**: you cannot use external Lua libraries (the Lua interpreter in tal is written in [Go](https://github.com/yuin/gopher-lua) and does not depend on the system or installed Lua libraries).

---

## Project structure

```
.
├── main.task.lua      # task file (DSL)
├── .tal.pack          # binary file with hashes (created automatically when tal runs)
└── ...
```

---

## Comparison with alternatives

| Feature | tal | make | just | task |
|-------------|-----|------|------|------|
| **Hash-based incrementality** | Yes | No | No | Yes |
| **Scripting language** | Lua with annotations | Shell | Shell | Shell |
| **Calling other tasks** | Yes | Yes | No | Yes |
| **Ease of writing** | Easy | Hard | Easy | Medium |

---

## License

Apache 2.0 - details in [LICENSE](LICENSE).

---

By Pt, 2026 - written using `lc`, `tap`, `pack`.