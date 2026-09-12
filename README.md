# run - script and task manager

[![Go Reference](https://pkg.go.dev/badge/github.com/pt-main/run.svg)](https://pkg.go.dev/github.com/pt-main/run)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-yellow.svg)](https://opensource.org/licenses/Apache-2.0)
[![Release](https://img.shields.io/github/v/release/pt-main/run)](https://github.com/pt-main/run/releases)

```bash
# run installation
go install github.com/pt-main/run/cmd/run@latest
# tal installation
go install github.com/pt-main/run/cmd/tal@latest
```

**run** is a tool for managing scripts, scripting any scenarios in an embedded Lua-like language with incrementality, storing scripts in global/local storage, complete independence from system and platform (works anywhere Go compiles), and with built-in ways to distribute scripts, for example via GitHub.

The project contains Task Lua (tal) inside itself - a task runner seamlessly integrated into run. More details can be read in the project [README](https://github.com/pt-main/run/blob/main/tal/README.md).

---

## Why run?

| Problem | run solves |
|----------|------------|
| **Scripts scattered across projects** | Global storage `~/run/` |
| **Need to remember paths** | One command: `run -r myscript` |
| **Different languages** | Support for Python, Bash, Batch, Lua - and easily extensible |
| **Grouping** | Tags for selective running (`--tagged`) |
| **Project scripts** | Local mode with `.run/` in the current folder |
| **Security** | TYCL config with a strict contract |
| **Compactness** | Small binary while fully platform-independent |

run gives **globality, simplicity and control** without unnecessary complexity.

## And why [Tal](https://github.com/pt-main/run/blob/main/tal/README.md)?

| Problem | tal solves |
|----------|------------|
| **Makefile is hard to read and write** | Simple DSL with comments and Lua instead of Shell |
| **Incrementality works poorly** | SHA256 hashes instead of modification time |
| **No calling tasks from each other** | Tasks can be called via a built-in function |
| **File dependencies are cumbersome** | works out of the box |

tal gives **incrementality, modernity and Lua** - all in one tool.

---

## Installation

### As a binary

Download the [release](https://github.com/pt-main/run/releases) for your OS/architecture and put it in `PATH`:

```bash
# Linux/macOS
chmod +x run-linux-amd64
sudo mv run-linux-amd64 /usr/local/bin/run

# Windows
# Just put run-windows-amd64.exe in a folder that is in PATH
```

### Via `go install`

```bash
go install github.com/pt-main/run@latest
```

**On first launch** run will create a structure in `~/run/`:
- `config.tycl` - config with the list of scripts.
- `scripts/` - Lua wrappers for launching.
- `base/` - original script files.

---


## Commands

| Command | Description | Example |
|---------|-------------|---------|
| `-add <path> <name> [docs] [--force]` | Add a script (supports `.py`, `.sh`, `.bat`, `.lua`) | `run -add script.py mypy` |
| `-remove <name>` | Remove a script | `run -remove mypy` |
| `-list` | Show the list of scripts | `run -list` |
| `-install <url> [name] [description] [--force] [--args="..."]` | Install a script from an external source, or run a tal script for installation | |
| `<name> [args...]` | Run a script (if the name does not match a command) | `run mypy arg1` |
| `-tag <name> <tags...>` | Add/remove tags. Use the `!` prefix for a tag to remove it. | `run -tag mypy deploy prod` |
| `-localmode [true/false]` | Enable/disable local mode, show the current script launch state | `run -localmode true` |
| `-r <name> [args...] [--tagged='...']` | Run a script | `run -r mypy arg1 arg2` |
| `-r --tagged="tag1;tag2;..."` | Run scripts with any of the tags | `run -r --tagged="deploy;test"` |
| `-r --tagged="..." --parallel` | Run scripts with the required tag in parallel | `run -r --tagged="deploy;build" --parallel` |
| `-r --tagged="..." --args=""` | Pass arguments to the script (if you need to avoid a conflict, for example with run flags, or not pass arguments) | `run -r --tagged="deploy;build" --args="--tagged dev"`,`run -r --tagged="deploy;build" --parallel --args` - does not pass arguments instead of passing `--parallel` |
| `-version` | Show the version of run and tal | `run -version` |

`--no_color` – flag disables colored output throughout the session.

---

## Local mode

By default run works globally (config in `~/run/`).  
Enable local mode - and run will use `.run/` in the current folder:

```bash
run -localmode true  # enable
run -localmode false # disable
run -localmode       # show state
```

This is convenient for projects: scripts are stored in the repository and do not interfere with the global config.

`--ll / --localmode / --gm / --globalmode` immediately after `run` - launch in local/global mode; after completion, restores the mode set with `run -localmode`.

---

## Language support

run automatically generates **Lua wrappers** that call the original scripts with the passed arguments.

| Extension | Language | Note |
|------------|------|------------|
| `.py` | Python | Looks for `python3`, then `python` |
| `.sh` | Bash | Executes via `bash` |
| `.bat` | Batch | Executes via `cmd /c` |
| `.lua` | Lua | Executes directly (without a wrapper) |
| `.task.lua` | Task Lua (Tal) | Executes via `run tal run` |

---

## Project structure

```
~/run/
├── config.tycl          # Config in TYCL (strict contract)
├── scripts/             # Lua wrappers for launching
│   └── myscript.lua
└── base/                # Original scripts
    └── myscript.py
```


### TYCL config

Script configuration is built on [Tycl](https://github.com/pt-main/tycl) - a typed language with the concept of contracts (fixed config formats).

Config contract -

```tycl
strict {
    scripts: objects = strict {
        name: string,        // Script name (command)
        script: string,      // Name of the wrapper file (matches the Lua script name inside run/scripts, without extension)
        description: string, // Description
        tags: strings,       // Tags
        ext: string,         // Extension (.py, .sh, .bat, .lua)
    },
}
```

The config is filled in automatically by the `run` CLI; after the first launch it looks like this -

```tycl
{
    scripts: objects = [
        {
            name: string = "test",
            script: string = "test",
            description: string = "[?BBK]Simple script for functions test[?RT]",
            ext: string = "",
            tags: strings = ["__test"],
        }
    ],
}
```

---

## Built-in Lua

Each wrapper is a Lua script that provides:

- `script_path(name)` - path to the original script.
- `get_arg(idx)` - get an argument by index.
- `get_args()` - table of all arguments.
- `run_script(name, ...)` - run another script from the wrapper.
- `run_script_parallel(name, ...)` – runs the specified script asynchronously in a background thread. Does not block execution of the current script. All arguments after the name are passed to the called script.
- `wait()` – waits for all background scripts started via `run_script_parallel` to finish. It is recommended to call it after starting parallel tasks to wait for their completion before the main script exits.
- `run_cli(args)` - run run cli with the passed arguments (as a string) in the current session.

Example:
```lua
run_script_parallel("build", "--release")
run_script_parallel("test")
wait()  -- wait for the build and tests to finish
```

---

## Examples

### Adding a script

```bash
run -add ~/projects/tools/deploy.py deploy "Deploy to production"
run -list
# ╭─────── Scripts
# ⎬─ deploy (.py):
# │     Deploy to production
# ╰───────
```

### Running

```bash
run -r deploy --env=prod
# or
run deploy --env=prod # when the script name does not conflict with run commands
```

### Tags

```bash
run -tag deploy prod utils
run -r --tagged="prod"   # will run all scripts with the prod tag
```

### Local mode

```bash
cd ~/myproject
run -localmode true
run -add script.py build
# now the script will be saved in .run/
```

or

```bash
run --localmode add script.py build
```

**Important**: for correct operation, the `--localmode` flag must be immediately after `run`.

---

By Pt, 2026 – written using `lc`, `tap`, `pack`, `tycl`.