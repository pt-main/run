package runtime

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/iancoleman/orderedmap"
	lccore "github.com/pt-main/lc/engine/core"
	"github.com/pt-main/run/tal/core"
	"github.com/pt-main/run/tal/generation"
	"github.com/pt-main/run/tal/lang"
	"github.com/pt-main/run/tal/lua"
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
  [?BBK]tal list tasks.lua[?RT]`,
		[]string{"task-lua-file"}, nil, true)

	p.AddCommand("run", RunHandler,
		`[?GN]Run a Tal file[?RT]

Parse and execute the given .task.lua file. All arguments after
the file name are passed directly to the Lua script via get_args()
and are not interpreted by the CLI.

[?BBE]Usage:[?RT]
  [?BBK]tal run <file> [args...][?RT]

[?BBE]Examples:[?RT]
  [?BBK]tal run tasks.lua[?RT]           # run the script (main block or default)
  [?BBK]tal run tasks.lua build[?RT]     # pass "build" as argument to the script
  [?BBK]tal run tasks.lua test -v[?RT]   # pass arguments to the script`,
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
	color.PrintlnColored("Update err: %v", Update())
	color.PrintlnColored("File creating err: %v", write("main.task.lua", []byte(`-- @
if #get_args() > 0 then
    run(get_args()[1]) 
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
		file, err := OpenF(fileName)
		if err != nil {
			return err
		}
		parsed, err := lang.Process(file)
		if err != nil {
			return err
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
	return Update()
}

func RunHandler(p *tap.Parser, s []string) error {
	ch, err := GetChanges()
	if err != nil {
		return err
	}
	if err := Update(); err != nil {
		return err
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
	ls := lua.NewTalLuaState(ch, args)
	file, err := OpenF(s[0])
	if err != nil {
		return err
	}
	processed, err := lang.Process(file)
	if err != nil {
		return fmt.Errorf(lccore.GetRealError(err))
	}
	generated, err := generation.GenerateCode(processed)
	return ls.DoString(generated)
}

func GetSavedFile() (*orderedmap.OrderedMap, error) {
	file, err := open(shared.TalFile)
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

func Update() error {
	st, err := core.SaveState(".")
	if err != nil {
		return err
	}
	file, err := core.StateAsPackCore(st)
	if err != nil {
		return err
	}
	return write(shared.TalFile, file)
}

func OpenF(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("Open: %v", err)
	}
	return string(data), nil
}

func open(file string) ([]byte, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("Open: %v", err)
	}
	return data, nil
}

func write(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Write: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.Write(data)
	if err != nil {
		return fmt.Errorf("Write: %v", err)
	}
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("Write: %v", err)
	}
	return nil
}
