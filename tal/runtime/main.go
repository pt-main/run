package runtime

import (
	"errors"
	"fmt"
	"strings"

	"github.com/iancoleman/orderedmap"
	lccore "github.com/pt-main/lc/engine/core"
	"github.com/pt-main/run/tal"
	"github.com/pt-main/run/tal/core"
	"github.com/pt-main/run/tal/lang"
	"github.com/pt-main/run/tal/shared"
	"github.com/pt-main/tap"
	"github.com/pt-main/tap/color"
)

func CreateCli() *tap.Parser {
	p := tap.NewParser("tal",
		`[?BE]╭─────── [?BRD]Tal[?RT] - Task Lua
[?BE]⎬─ [?RT]Incremental task runner with Lua and file-based dependencies.
[?BE]│  [?RT]By [?UE]Pt[?RT], only [?BD]humanmade[?RT].
[?BE]╰───────[?RT]`,
		[]string{"-h", "help"},
		tap.DefaultParserConfig(),
	)

	p.AddCommand("update", UpdateHandler,
		`[?GN]Update .tal.pack[?RT]

Scan the current directory, compute SHA256 hashes for all files,
and store them in a binary .tal.pack file. Run this once before
using 'tal run' to enable incremental builds.

[?BBE]Example:[?RT]
  [?BBK]tal update[?RT]`,
		nil, nil, false)

	p.AddCommand("list", ListHandler,
		`[?GN]List tasks in a Tal file[?RT]

Read a .task.lua file and display all defined tasks (including
global and main blocks). Useful for quickly checking what's
available in a project.

[?BBE]Usage:[?RT]
  [?BBK]tal list <files...>[?RT]

[?BBE]Example:[?RT]
  [?BBK]tal list main.task.lua[?RT]`,
		[]string{"task-lua-file"}, nil, true)

	p.AddCommand("run", RunHandler,
		`[?GN]Run a Tal file[?RT]

Parse and execute the given .task.lua file. All arguments after
the file name are passed directly to the Lua script via get_args()
and are not interpreted by the CLI.

[?BBE]Usage:[?RT]
  [?BBK]tal run <file> [args...] [--deps='dep1;dep2'][?RT]

[?BBE]Flag [?BBK]--deps[?RT]
  [?BBK]Used to pass dependencies for the [?GN]--#depends[?BBK] annotation. 
  [?BBK]If this flag is not provided, modified files are specified 
  [?BBK]as dependencies (requires the presence of [?YW].tal.pack[?BBK]).

[?BBE]Examples:[?RT]
  [?BBK]tal run main.task.lua[?RT]           # run the script (main block or default)
  [?BBK]tal run main.task.lua build[?RT]     # pass "build" as argument to the script
  [?BBK]tal run main.task.lua test -v[?RT]   # pass arguments to the script`,
		[]string{"task-lua-file"}, nil, true)

	p.AddCommand("init", InitHandler,
		`[?GN]Create a default Tal file[?RT]

Generate a minimal main.task.lua file with a main block that
accepts arguments. This is a quick way to start a new project.

[?BBE]Example:[?RT]
  [?BBK]tal init[?RT]`,
		nil, nil, false)

	return p
}

func InitHandler(p *tap.Parser, s []string) error {
	color.PrintlnColored("Update err: %v", core.Update())
	color.PrintlnColored("File creating err: %v", core.Write("main.task.lua", []byte(`-- @
if #get_args() > 0 then
    script(get_args()[1]) 
end`)))
	return nil
}

func center(s string, width int) string {
	runes := []rune(s)
	n := len(runes)
	if n >= width {
		return s
	}
	left := (width - n) / 2
	right := width - n - left
	return fmt.Sprintf("%*s%s%*s", left, "", s, right, "")
}

func ListHandler(p *tap.Parser, s []string) error {
	for _, fileName := range s {
		file, err := core.OpenF(fileName)
		if err != nil {
			return err
		}
		var err_ lccore.ErrorInterface
		parsed, err_ := lang.Process(file)
		if err_ != nil {
			return errors.New(lang.ErrFmt(err_))
		}
		res := []string{"[?GN]╭─────── [?RT][[?YW]" +
			center(fileName, 20) + "[?RT]] Scripts"}
		templ := "[?GN]│  [?RT]%3d: [?BGN]%v"
		idx := 0
		if parsed.Global != nil {
			idx += 1
			res = append(res, fmt.Sprintf(templ, idx, "Global"))
		}
		for script := range parsed.Blocks {
			res = append(res, fmt.Sprintf(templ, idx, script))
			idx += 1
		}
		if parsed.Main != nil {
			res = append(res, fmt.Sprintf(templ, idx, "Main"))
		}
		res = append(res, "[?GN]╰───────[?RT]")
		color.PrintlnColored(strings.Join(res, "\n"))
	}
	return nil
}

func UpdateHandler(p *tap.Parser, s []string) error {
	return core.Update()
}

func RunHandler(p *tap.Parser, s []string) (err error) {
	ch := []string{}
	if deps, hasDeps := p.Flags["deps"]; hasDeps {
		ch = strings.Split(deps, ";")
	} else {
		ch, err = GetChanges()
		if err != nil {
			return
		}
	}
	args := []string{}
	skippedName := false
	for _, arg := range p.RawArgs[1:] {
		if arg == s[0] && !skippedName {
			skippedName = true
		} else {
			args = append(args, arg)
		}
	}
	file, err := core.OpenF(s[0])
	return tal.Process(ch, args, file)
}

func GetSavedFile() (*orderedmap.OrderedMap, error) {
	file, err := core.Open(shared.TalFile)
	if err != nil {
		return nil, err
	}
	return core.PackCoreAsState(file)
}

func GetChanges() ([]string, error) {
	w, err := GetSavedFile()
	if err != nil {
		return nil, err
	}
	return core.Changes(w, ".")
}
