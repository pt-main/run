# Changelog

All notable changes to **run** (a script and task manager) and to **tal**, the
task runner bundled into it.

Format follows [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
versioning follows [SemVer](https://semver.org/).

[Russian version](changelog-ru.md)

## [1.4.5] - 2026-09-29

Moves the project onto the current `lc` and `tycl` releases. `tap` was already
on its latest Go module and did not move.

### Changed

- **`lc` v1.5.8 -> v2.0.1.** The module path now follows semantic import
  versioning, so every import of the framework carries the `/v2` suffix:
  `github.com/pt-main/lc/engine/core` becomes
  `github.com/pt-main/lc/v2/engine/core`. Five files were updated:
  `tal/lang/lcproc.go`, `tal/lang/process.go`, `tal/lang/errFmt.go`,
  `tal/generation/generate.go` and `tal/shared/errors.go`.
- **`tycl` v1.3.8 -> v1.4.0.** The config layer gained the flat `diag` package,
  a `--json` envelope and new CLI commands. The Go API used by `run`
  (`shared.Config`, `generation.Tycl`, `format.FormatError`, `utils.OpenF`,
  `utils.WriteF`) is unchanged, so no call site moved.

### Fixed

- **The tal lexer stopped accepting `-- #depends` after the `lc` upgrade.**
  `lc` v2 matches every rule against the whole rune slice anchored at the
  current position, so `^` became a real line start again and no longer
  swallowed the newline in front of an annotation. That newline arrived as a
  separate `CODE` token, and the `block` grammar had no rule to absorb it, so
  parsing failed with
  `expected "BLOCK", got "COMMAND" (raw: "-- #depends *.txt")`.
  `tal/lang/lcproc.go` now emits newlines as a `WHITESPACE` token and passes it
  to `NewParser` as an ignorable type, which restores the previous token
  stream. Only newlines are consumed: a broader `\s+` rule also strips the
  spaces inside code lines and turns `local function` into `localfunction`.

## [1.4.0] - 2026-09-27

Script and template management, local mode, and the `run tal` subcommand.

### Added

- **Custom wrapper templates.** `run manage templ-add` registers a Go
  `text/template` body for any file extension; the output is a Lua wrapper that
  calls the original file. Bodies live in `templates/` and `config.tycl` keeps
  only `{ext, file}`, so templates can be edited with ordinary editors.
- **Tags.** `run manage tag <script> <tag...>` adds tags, a leading `!` removes
  them. `run -r --tagged="a;b"` runs every script carrying any of the tags,
  `--parallel` runs them concurrently, `--args` passes arguments to the
  scripts rather than to `run`.
- **Parallel execution from Lua.** `run_script_parallel(name, ...)` starts a
  script in a background thread, and `wait()` joins all of them before the
  wrapper exits.
- **Local mode.** `run sys localmode true` switches run to a `.run/` directory
  next to the project; `--localmode` and `--globalmode` override the mode for a
  single invocation.
- **A `run tal` subcommand** exposing the bundled task runner, with
  `run tal init`, `update`, `list` and `run`.
- `.task.lua` and `.nd.task.lua` are recognised as Task Lua and run through tal,
  the second form with dependency checking disabled.

### Changed

- Configs written by older versions, where a template body was kept inline in
  the `template` field, are migrated once on read: the body moves to
  `templates/<ext>.templ` and only `file` stays in the config.

## [1.3.3] - 2026-09-12

### Added

- `run/lua.go`: the Lua API surface used by generated wrappers, including
  `run_script`, `get_arg`, `get_args` and `script_path`.

### Changed

- The licence file was renamed from `LICENCE` to the conventional `LICENSE`.

## [1.3.2] - 2026-09-11

### Added

- **Installation from a URL** (`run/install.go`): raw file URLs, GitHub blob
  URLs and the shorter `github.com/user/repo@branch/path` form are resolved to
  a raw download, and a `.task.lua` URL is executed as an installation script
  with `--args="..."`.

## [1.3.0] - 2026-09-07

### Changed

- **The CLI was split into subcommands.** `run/handlers.go` and a single
  `cmd/run/cli.go` became `cmd/runcli/` with `cli.go` and `process.go`, so
  `manage`, `sys` and the run path are separate parsers instead of one
  dispatcher. Aliases (`scradd`, `screm`, `tladd`, `tlrem`) were introduced to
  keep the short forms working.

## [1.2.5] - 2026-08-27

### Fixed

- tal runtime and Lua binding corrections across `tal/core`, `tal/lua` and
  `tal/runtime`.

## [1.2.4] - 2026-08-26

### Changed

- The standalone tal binary moved from `tal/cmd/tal` to `cmd/tal`, so both
  commands are installed by the same `go install` pattern.

## [1.2.3] - 2026-08-26

### Fixed

- Config handling corrections in `run/files.go`, `run/handlers.go`,
  `run/stdlib.go` and `run/tycl.go`.

## [1.2.2] - 2026-08-26

### Added

- `tal/lang/errFmt.go`: readable formatting of parse errors for tal.

### Changed

- tal runtime adjustments in `tal/core/main.go` and `tal/runtime/main.go`.

## [1.2.1] - 2026-08-24

### Changed

- Binaries under `build/` were removed from the repository; they are produced
  by the build script instead of being committed.
- `merged.md` was deleted.

## [1.2.0] - 2026-08-23

### Added

- **tal, the bundled task runner.** A new `tal/` tree (`lang`, `generation`,
  `core`, `lua`, `runtime`, `shared`), a standalone `cmd/tal` binary, and
  bilingual `tal/README.md` and `tal/README-ru.md`. Tasks are plain Lua with
  annotations in comments - `-- @name` for a task, `-- @` for the main block,
  `-- @!` for the global block and `-- #depends <glob...>` for file
  dependencies - so a file stays valid Lua and incrementality is tracked with
  SHA256 hashes in `.tal.pack` rather than modification times.
- `build.json` describing the cross-compilation matrix.

## [1.1.0] - 2026-08-21

### Added

- `README-RU.md`, a Russian translation of the README.
- A `.run/` example showing the local configuration layout
  (`config.tycl` and `scripts/`).

## [1.0.0] - 2026-08-20

### Added

- First release: the `run` binary, the TYCL-backed configuration, wrapper
  generation for `.py`, `.sh`, `.bat` and `.lua`, and the built-in Lua API
  (`script_path`, `get_arg`, `get_args`).

---

## Change types

- **Added** - new functionality.
- **Changed** - changes in existing functionality.
- **Deprecated** - soon to be removed functionality.
- **Removed** - removed functionality.
- **Fixed** - bug fixes.
- **Security** - vulnerability fixes.
- **Breaking** - API breaking changes.

## Links

- [GitHub](https://github.com/pt-main/run)
- [Go Reference](https://pkg.go.dev/github.com/pt-main/run)
- [tal](tal/README.md)
