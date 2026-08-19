# run - Script Manager

[![Go Reference](https://pkg.go.dev/badge/github.com/pt-main/run.svg)](https://pkg.go.dev/github.com/pt-main/run)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache_2.0-yellow.svg)](https://opensource.org/licenses/Apache-2.0)
[![Release](https://img.shields.io/github/v/release/pt-main/run)](https://github.com/pt-main/run/releases)

```bash
go install github.com/pt-main/run@latest
```

**run** is a script manager that lets you add, remove, and run scripts in different languages with a single command. Scripts are stored in `~/run/` and are accessible from any folder.

---

## Why run?

| Problem | run solves |
|---------|------------|
| **Scripts scattered across projects** | Global storage in `~/run/` |
| **Need to remember paths** | One command: `run r myscript` |
| **Different languages** | Supports Python, Bash, Batch, Lua — and easily extensible |
| **Grouping** | Tags for selective execution (`--tagged`) |
| **Project‑specific scripts** | Local mode with `.run/` in the current folder |
| **Safety** | Configuration in TYCL with a strict contract |
| **Compactness** | Small binary, fully platform‑independent |

run gives you **globality, simplicity, and control** without unnecessary complexity.

---

## Installation

### As a binary

Download the [release](https://github.com/pt-main/run/releases) for your OS/architecture and place it in your `PATH`:

```bash
# Linux/macOS
chmod +x run-linux-amd64
sudo mv run-linux-amd64 /usr/local/bin/run

# Windows
# Just place run-windows-amd64.exe in a folder that is in your PATH
```

### Via `go install`

```bash
go install github.com/pt-main/run@latest
```

**On first run**, run will create the following structure in `~/run/`:
- `config.tycl` – configuration (TYCL) with the script list.
- `scripts/` – Lua wrappers for running scripts.
- `base/` – original script files.

---

## Syntax

```bash
run [--<lm/localmode>] <cmd> <args...>
```

### Commands

| Command | Description | Example |
|---------|-------------|---------|
| `add <path> <name> [docs]` | Add a script (supports `.py`, `.sh`, `.bat`, `.lua`) | `run add script.py mypy` |
| `remove <name>` | Remove a script | `run remove mypy` |
| `list` | Show the list of scripts | `run list` |
| `r <name> [args...]` | Run a script | `run r mypy arg1 arg2` |
| `<name> [args...]` | Run a script (if the name does not conflict with a run command) | `run mypy arg1` |
| `tag <name> <tags...>` | Add tags | `run tag mypy deploy prod` |
| `rm-tag <name> <tags...>` | Remove tags | `run rm-tag mypy prod` |
| `localmode [true/false]` | Enable/disable local mode, show current state | `run localmode true` |
| `--tagged="tag1;tag2"` | Run scripts with any of the given tags | `run --tagged="deploy;test"` |
| `version` | Show version | `run version` |

### Flags

- `--force` with `add` – replaces an existing script with the same name.
- `--tagged="tag1;tag2"` with `r` – run by tags.
- `--ll` / `--localmode` immediately after `run` – run in local mode; after completion, restores the mode previously set by `run localmode`.

---

## Local mode

By default, run works globally (configuration in `~/run/`).  
Enable local mode and run will use `.run/` in the current folder:

```bash
run localmode true  # enable
run localmode false # disable
run localmode       # shows current state (false / true)
```

This is convenient for projects: scripts are stored in the repository and do not interfere with the global configuration.

---

## Language support

run automatically generates **Lua wrappers** that invoke the original scripts with the passed arguments.

| Extension | Language | Note |
|-----------|----------|------|
| `.py` | Python | Looks for `python3`, then `python` |
| `.sh` | Bash | Executes via `bash` |
| `.bat` | Batch (Windows) | Executes via `cmd /c` |
| `.lua` | Lua | Executed directly (no wrapper) |

Adding a new language is easy – just extend the template in `templates.go`.

---

## Project structure

```
~/run/
├── config.tycl          # Configuration in TYCL (strict contract)
├── scripts/             # Lua wrappers for execution
│   └── myscript.lua
└── base/                # Original scripts
    └── myscript.py
```

### TYCL configuration

Script configuration is built on [Tycl](https://github.com/pt-main/tycl) – a typed language with the concept of contracts (fixed configuration formats).

The configuration contract:

```tycl
strict {
    scripts: objects = strict {
        name: string,        // Script name (command)
        script: string,      // Wrapper file name (matches the Lua script name inside run/scripts, without extension)
        description: string, // Description
        tags: strings,       // Tags
        source: string,      // Original script source path
        ext: string,         // Extension (.py, .sh, .bat, .lua)
    },
}
```

The configuration is populated automatically via the `run` CLI. After the first run it looks like this:

```tycl
{
    scripts: objects = [
        {
            name: string = "test",
            script: string = "test",
            description: string = "[?BBK]Simple script for functions test[?RT]",
            source: string = "",
            ext: string = "",
            tags: strings = ["__test"],
        }
    ],
}
```

---

## Built‑in Lua

Each wrapper is a Lua script that provides:

- `script_path(name)` – path to the original script.
- `get_arg(idx)` – get an argument by index.
- `get_args()` – table of all arguments.
- `run_script(name, ...)` – run another script from the wrapper.

---

## Examples

### Adding a script

```bash
run add ~/projects/tools/deploy.py deploy "Deploy to production"
run list
# ╭─────── Scripts
# ⎬─ deploy (.py):
# │     Deploy to production
# ╰───────
```

### Running

```bash
run r deploy --env=prod
# or
run deploy --env=prod   # when the script name does not conflict with run commands
```

### Tags

```bash
run tag deploy prod utils
run r --tagged="prod"   # runs all scripts with the tag prod
```

### Local mode

```bash
cd ~/myproject
run localmode true
run add script.py build
# now the script will be saved in .run/
```

or

```bash
run --localmode add script.py build
```

**Important**: the `--localmode` flag must appear immediately after `run` to work correctly.

---

## License

Apache 2.0 – see [LICENSE](LICENSE) for details.

---

By Pt, 2026 – built with tap, tycl, and lc.
