# run - script and task manager

[![Go Reference](https://pkg.go.dev/badge/github.com/pt-main/run.svg)](https://pkg.go.dev/run)
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

> Russian version of this document: [README-RU.md](README-RU.md).

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
go install github.com/pt-main/run/cmd/run@latest
go install github.com/pt-main/run/cmd/tal@latest   # optional, standalone tal
```

**On first launch** run will create a structure in `~/run/`:
- `config.tycl` - config with the list of scripts.
- `scripts/` - Lua wrappers for launching.
- `templates/` - bodies of custom wrapper templates.
- `base/` - original script files.

---

## Commands

The CLI consists of the root parser `run` and three subcommands: `manage` (script and
template management), `sys` (system operations) and `tal` (the bundled task runner).

### Running scripts

| Command | Description | Example |
|---------|-------------|---------|
| `-r <name> [args...]` | Run a script by name (explicit form) | `run -r mypy arg1 arg2` |
| `<name> [args...]` | Run a script (when the name does not conflict with run commands) | `run deploy --env=prod` |
| `-r --tagged="tag1;tag2;..."` | Run all scripts carrying any of the given tags | `run -r --tagged="deploy;test"` |
| `-r --tagged="..." --parallel` | Run the tagged scripts in parallel | `run -r --tagged="deploy;build" --parallel` |
| `-r --tagged="..." --args="..."` | Pass arguments to the script (use it if arguments conflict with run flags) | `run -r --tagged="deploy" --args="--tagged dev"` |
| `-r --tagged="..." --args` | Pass **no** arguments, instead of passing run flags to the script | `run -r --tagged="deploy" --parallel --args` |

### Managing data: `run manage`

| Command | Description | Example |
|---------|-------------|---------|
| `run manage script-add <path> <name> [docs] [--force]` | Add a script (supports `.py`, `.sh`, `.bat`, `.lua`, `.task.lua`, `.nd.task.lua`) | `run manage script-add ./deploy.py deploy "Deploy to production"` |
| `run manage script-remove <name>` | Remove a script from the config | `run manage script-remove mypy` |
| `run manage list` | Show the list of registered scripts | `run manage list` |
| `run manage tag <name> <tags...>` | Add/remove tags. Prefix a tag with `!` to remove it | `run manage tag mypy deploy !prod dev` |
| `run manage install <url> [name] [description] [--force] [--args="..."]` | Install a script from an external source, or run a tal installation script | `run manage install github.com/user/repo@main/deploy.py` |
| `run manage templ-add <ext> [file] [--source="..."] [--force]` | Add a wrapper template for an extension | `run manage templ-add ".go" templ.txt` |
| `run manage templ-remove <ext>` | Remove a template | `run manage templ-remove ".go"` |

Aliases: `scradd` = `script-add`, `screm` = `script-remove`, `tladd` = `templ-add`, `tlrem` = `templ-remove`.

For details: `run manage help`.

### System operations: `run sys`

| Command | Description | Example |
|---------|-------------|---------|
| `run sys version` | Show the versions of run and tal | `run sys version` (alias: `run sys -v`) |
| `run sys localmode` | Show the current mode and the config path | `run sys localmode` |
| `run sys localmode true` | Enable local mode | `run sys localmode true` |
| `run sys localmode false` | Disable local mode | `run sys localmode false` |

Alias: `-lm` = `localmode`.

### Bundled task runner: `run tal`

All tal commands are available through `run tal ...`. See the [tal README](https://github.com/pt-main/run/blob/main/tal/README.md).

```bash
run tal init
run tal update
run tal list main.task.lua
run tal run main.task.lua build
```

### Built-in [tap](https://github.com/pt-main/tap) flags

- `--verbose` - verbose output.
- `--debug` - debug output.
- `-h`, `-help`, `help` - help.
- `--no_color` - disable colored output for the session.

---

## Local mode

By default run works globally (config in `~/run/`).
Enable local mode - and run will use `.run/` in the current folder:

```bash
run sys localmode true   # enable
run sys localmode false  # disable
run sys localmode        # show state
```

This is convenient for projects: scripts are stored in the repository and do not interfere with the global config.

The flags `--lm`, `--localmode`, `--gm`, `--globalmode` - placed immediately after `run` - switch the mode
only for the current invocation, and afterwards the mode set with `run sys localmode` is restored.

```bash
run --localmode manage list                  # show local scripts
run --globalmode -r deploy                   # run a global script
run --lm manage install github.com/pt-main/run-scripts@main/sysfetch.lua
```

**Important**: the `--localmode` / `--globalmode` flag must come immediately after `run`.

---

## Language support

run automatically generates **Lua wrappers** that call the original scripts with the passed arguments.

Built in:

| Extension | Language | Note |
|------------|------|------------|
| `.py` | Python | Looks for `python3`, then `python` |
| `.sh` | Bash | Executes via `bash` |
| `.bat` | Batch | Executes via `cmd /c` |
| `.lua` | Lua | Executes directly (without a wrapper) |
| `.task.lua` | Task Lua (Tal) | Runs the file as a tal task file (with dependencies) |
| `.nd.task.lua` | Task Lua (Tal) | Same, but with dependency checking disabled (`nd` - no deps) |

Any other extension is available through a [custom wrapper template](#custom-wrapper-templates).

---

## Custom wrapper templates

For extensions that are not built in, you can add your own **wrapper template**. A template is a
[Go `text/template`](https://pkg.go.dev/text/template) file, and its output becomes the Lua script
saved to `scripts/<name>.lua`.

The template body is **not stored in the config**: it lives in the `templates/` directory, and
`config.tycl` keeps only a file reference:

```tycl
templates: objects = [
    {
        ext: string = ".rb",
        file: string = "rb.templ",   // file inside templates/
    }
],
```

### Available variables

| Variable | Value |
|------------|----------|
| `{{.ext}}` | File extension, e.g. `.rb`, `.go` |
| `{{.name}}` | Internal name of the file inside `base/` (used in `script_path(...)`) |

The wrapper gets its generated name via `script_path("{{.name}}")` - this is how it finds the original
script in `base/`.

### Template selection order

When `run manage script-add` runs, the template is chosen like this:

1. If the file name ends with `.nd.task.lua` or `.task.lua` - the built-in tal template is used.
2. Otherwise a **custom** template is looked up, whose `ext` is a suffix of the file name
   (e.g. `.rb` for `script.rb`).
3. If nothing matched - the template with an **empty** `ext` (fallback) is taken.
4. If there is no fallback either - the built-in templates for `.py`, `.sh`, `.bat`, `.lua`.
5. If nothing fits - the error `Unsupportable file extension`.

### Managing templates

```bash
# from a file (the content is copied to templates/<ext>.templ)
run manage templ-add ".rb" ruby-template.lua

# from a string
run manage templ-add ".rb" --source='...'

# replace an existing one
run manage templ-add ".rb" ruby-template.lua --force

# remove (the file in templates/ is removed too)
run manage templ-remove ".rb"     # or: run manage tlrem ".rb"
```

> `file` may also be an absolute or relative path - if it is absolute, it is read as is, otherwise it
> is looked up in the `templates/` directory.

### Example: a template for Ruby

File `ruby-template.lua`:

```lua
-- === CONFIGURATION ===
local script_file = script_path("{{.name}}")
local args = get_args()
-- =====================

local function escape(arg)
    if arg:match("[ \t\"']") then
        return '"' .. arg:gsub('"', '\\"') .. '"'
    end
    return arg
end

local cmd = "ruby " .. escape(script_file)
for _, a in ipairs(args) do
    cmd = cmd .. " " .. escape(a)
end

local result = os.execute(cmd)
os.exit(result or 0)
```

Add it and use it:

```bash
run manage templ-add ".rb" ruby-template.lua
run manage script-add ./my_tool.rb mytool "My Ruby tool"
run -r mytool arg1 arg2
```

### Important notes

- A template is a **Go template**, not Lua. `{{.ext}}` and `{{.name}}` are substituted before the file
  is written to `scripts/`.
- The template body is the Lua code of the future wrapper. It must terminate correctly (`os.exit(...)`).
- One `ext` = one template. Use `--force` to replace an existing one.
- Template bodies live in `templates/`, and the `templates` field of `config.tycl` only keeps
  `{ext, file}` - the config stays readable and templates can be edited and highlighted by ordinary
  editors.
- Old configs, where the template body was kept inline in the `template` field, are migrated
  automatically: the body is moved to `templates/<ext>.templ` and only `file` is left in the config.
  The migration happens once, when the config is read.

---

## Project structure

```
~/run/
├── config.tycl          # Config in TYCL (strict contract)
├── scripts/             # Lua wrappers for launching
│   └── myscript.lua
├── templates/           # Bodies of custom wrapper templates
│   └── rb.templ
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
        ext: string,         // Extension (.py, .sh, .bat, .lua, .task.lua)
    },
    templates: objects = strict {
        ext: string,         // File extension
        file: string,        // Template body file inside the templates/ dir
    },
}
```

The config is filled in automatically by the `run` CLI; after the first launch it looks like this -

```tycl
{
    templates: objects = [],
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
- `run_script_parallel(name, ...)` - runs the specified script asynchronously in a background thread. Does not block execution of the current script. All arguments after the name are passed to the called script.
- `wait()` - waits for all background scripts started via `run_script_parallel` to finish. It is recommended to call it after starting parallel tasks to wait for their completion before the main script exits.
- `run_cli(args)` - run the run cli with the passed arguments (as a string) in the current session.

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
run manage script-add ~/projects/tools/deploy.py deploy "Deploy to production"
run manage list
# ╭─────── Scripts
# ⎬─ deploy (.py):
# │     Deploy to production
# ╰───────
```

Alias:

```bash
run manage scradd ~/projects/tools/deploy.py deploy "Deploy to production"
```

### Running

```bash
run -r deploy --env=prod
# or
run deploy --env=prod   # when the script name does not conflict with run commands
```

### Tags

```bash
run manage tag deploy prod utils
run -r --tagged="prod"    # will run all scripts with the prod tag
```

Removing a tag:

```bash
run manage tag deploy !utils
```

### Installing from GitHub

```bash
# a plain script file
run manage install github.com/user/repo@main/deploy.py deploy "Prod deploy"

# a tal installation script
run manage install github.com/user/repo@main/run.task.lua --args="--version 1.2.3"
```

### Local mode

```bash
cd ~/myproject
run sys localmode true
run manage script-add script.py build
# now the script will be saved in .run/
```

or for a single invocation:

```bash
run --localmode manage script-add script.py build
```

### Version

```bash
run sys version
# or
run sys -v
```

---

## License

Apache 2.0 - details in [LICENSE](LICENSE).

---

By Pt, 2026 - written using `lc`, `tap`, `pack`, `tycl`.
