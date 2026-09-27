package runcli

import (
	"fmt"

	"github.com/mattn/go-shellwords"
	"github.com/pt-main/run/run/api"
	runlib "github.com/pt-main/run/run/runcli/handlers"
	luaruntime "github.com/pt-main/run/tal/lua"
	"github.com/pt-main/run/tal/runtime"
	tap "github.com/pt-main/tap/go"
	lua "github.com/yuin/gopher-lua"
)

func NewCli() (*tap.Parser, error) {
	var lp *tap.Parser
	conf := tap.DefaultParserConfig()
	conf.BuiltinVerboseDebug = true
	p := tap.NewParser("run", `[?BE]╭─────── [?BRD]Run[?RT]
[?BE]⎬─ [?RT]Simple and powerful script manager
[?BE]│  [?RT]By [?UE]Pt[?RT], only [?BD]humanmade[?RT].
[?BE]╰───────[?RT]`, []string{"-h", "-help"}, conf)

	p.AddSubcommand("tal", runtime.CreateCli())

	p.AddCommand("-r",
		runlib.MakeRunHandler(true),
		`[?GN]Run a script by name.[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run -r <name> [args...] [--args="..."][[?RT]
  [?BBK]run -r --tagged="..." [--parallel] [--args="..."][?RT]
[?YW]Flags:[?RT]
  [?GN]--tagged="tag1;tag2;..."[?RT]  Run all scripts with any of the specified tags (semicolon-separated)
  [?GN]--parallel[?RT]                Run tagged scripts in parallel
  [?GN]--args="..."[?RT]              Explicitly pass arguments to script (useful when args conflict with run flags)
  [?GN]--args[?RT]                    Pass no arguments (instead of passing run flags)
[?YW]Examples:[?RT]
  [?BBK]run -r deploy --env=prod[?RT]
  [?BBK]run -r --tagged="deploy;test" --parallel[?RT]
  [?BBK]run -r --tagged="build" --args="--verbose"[?RT]`,
		nil, nil, true)

	p.AddCommand(tap.DEFAULT_CMD,
		runlib.MakeRunHandler(false),
		`[?GN]Run a script by name (when name doesn't conflict with run commands).[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run <script> [args...][?RT]
  [?BBK]run <cmd> [args...][?RT]
[?YW]Examples:[?RT]
  [?BBK]run deploy --env=prod[?RT]
  [?BBK]run mypy arg1 arg2[?RT]`,
		nil, nil, true)

	m, err := NewManage()
	if err != nil {
		return nil, err
	}
	p.AddSubcommand("manage", m)

	p.AddSubcommand("sys", NewSys())

	runcli := func(L *lua.LState) int {
		input := L.OptString(1, "")
		if input == "" {
			L.Push(lua.LString("missing command string"))
			return 2
		}
		parsed, err := shellwords.Parse(input)
		if err != nil {
			fmt.Println("Parsing cli args:", err)
			L.Push(lua.LString(err.Error()))
			return 2
		}
		if lp == nil {
			lp, err = NewCli()
			if err != nil {
				L.Push(lua.LString(err.Error()))
				return 2
			}
		}
		if err := Process(lp, parsed); err != nil {
			fmt.Println("Run cli err:", err)
			L.Push(lua.LString(err.Error()))
			return 2
		}
		return 0
	}

	luaruntime.RegisterLuaFunc("run_cli", func(changedFiles, args []string) lua.LGFunction {
		return runcli
	})

	api.RegisterLuaFunc("run_cli", runcli)

	return p, nil
}
