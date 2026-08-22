# tal – incremental tasker with Lua and dependencies

[![Go Reference](https://pkg.go.dev/badge/github.com/pt-main/tal.svg)](https://pkg.go.dev/github.com/pt-main/tal)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-yellow.svg)](https://opensource.org/licenses/Apache-2.0)
[![Release](https://img.shields.io/github/v/release/pt-main/tal)](https://github.com/pt-main/tal/releases)

> tal - task lua

```bash
go install github.com/pt-main/tal/cmd/tal@latest
```

**tal** is a simple, modern task runner that uses Lua as its scripting language. It lets you describe tasks in plain Lua with annotations, track file changes, and run only what has actually changed.

---

## Why tal?

| Problem | tal solves |
|---------|------------|
| **Makefiles are hard to read and write** | Simple DSL with comments and Lua instead of Shell |
| **Incrementality is broken** | SHA256 hashes instead of modification times |
| **No way to call tasks from each other** | Tasks can be called via built‑in functions |
| **File dependencies are cumbersome** | `#depends file1 file2` works out of the box |

tal gives you **incrementality, simplicity, and Lua** – all in one tool.

---

## Installation

### As a binary

Download the [release](https://github.com/pt-main/tal/releases) for your OS/architecture, rename it, and put it in your `PATH`:

```bash
# Linux/macOS
chmod +x tal
sudo mv tal /usr/local/bin/tal

# Windows
# Place tal.exe in a folder that is in your PATH
```

### Via `go install`

```bash
go install github.com/pt-main/tal/cmd/tal@latest
```

On the first run, `tal update` creates `.tal.pack` – a file containing hashes of all files in the current directory.

---

## Syntax

The task file is written in plain Lua with annotations in comments, without breaking the Lua syntax.

### Basic constructs

| Construct | Description |
|-----------|-------------|
| `-- @taskname` | Start of a task block |
| `-- @` | Main block (run by default) |
| `-- @!` | Global block (runs before main and task registration) |
| `-- #depends <glob-pattern...>` | Command – file dependency (checked by their hashes). File names are written as glob patterns |

Any other code is treated as normal Lua code. The file must start with a normal or global block.

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
run("build")
run("test")
```

---

## CLI Commands

```bash
tal run <args>    # parses tasks.tal, executes the DSL with arguments
tal update        # force‑update or initialise .tal.pack
```

`tal update` is mandatory when running tal for the first time in a directory.

---

## How incrementality works

Incrementality is enabled by the `depends` command (`-- #depends ...`) and does not work without it.

1. `tal` scans the current directory and computes SHA256 hashes for all files.
2. Hashes are stored in `.tal.pack` (binary format, using [`pack`](https://github.com/pt-main/pack)).
3. On the next run, `tal` compares hashes to detect changed files, and automatically updates the hashes.
4. In the generated Lua script, the `changed_list` array contains the paths of changed files.
5. The runtime checks each task's dependencies and only runs those where at least one dependent file has changed.

---

## Built‑in Lua runtime

Each task is a Lua function that executes in an environment with access to:

- `changed_list` – a table with paths of changed files.
- `tasker.add(deps, name, func)` – registers a task.
- `tasker.run(name)` – executes a task.
- `run(name)` – shorthand for `tasker.run`.

**Important**: You cannot use external Lua libraries (the Lua interpreter in tal is written in Go using [gopher‑lua](https://github.com/yuin/gopher-lua) and does not depend on system‑installed Lua libraries).

---

## Project structure

```
.
├── tasks.lua          # task file (DSL)
├── .tal.pack          # binary file with hashes (created automatically by tal update)
└── ...
```

---

## Comparison with alternatives

| Feature | tal | make | just | task |
|---------|-----|------|------|------|
| **Incrementality by hashes** | Yes | No | No | Yes |
| **Scripting language** | Lua with annotations | Shell | Shell | Shell |
| **Calling other tasks** | Yes | Yes | No | Yes |
| **Ease of writing** | Easy | Hard | Easy | Medium |

---

## License

Apache 2.0 – see [LICENSE](LICENSE) for details.

---

By Pt, 2026 – written using `lc`, `tap`, `pack`.