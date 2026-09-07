package runcli

import (
	"fmt"

	"github.com/mattn/go-shellwords"
	"github.com/pt-main/run"
	runlib "github.com/pt-main/run/run"
	localmode "github.com/pt-main/run/run/localMode"
	luaruntime "github.com/pt-main/run/tal/lua"
	"github.com/pt-main/run/tal/runtime"
	"github.com/pt-main/tap"
	lua "github.com/yuin/gopher-lua"
)

func NewCli() *tap.Parser {
	var lp *tap.Parser
	conf := tap.DefaultParserConfig()
	conf.BuiltinVerboseDebug = true
	p := tap.NewParser("run", `[?BE]╭─────── [?BRD]Run[?RT]
[?BE]⎬─ [?RT]Simple and powerful script manager
[?BE]│  [?RT]By [?UE]Pt[?RT], only [?BD]humanmade[?RT].
[?BE]╰───────[?RT]

[?GN]Usage: [?RT]
  [?BYW]Type: [?BBK]run [--<lm|localmode>] [--no_color] <cmd> <args...>

  [?BBK]run <cmd> <args...> [?YW]- execute run command (described in help)[?RT]
    [?BYW]Example: [?BBK]run -add script.py script1[?RT]

  [?BBK]run -r <script> <args...> [?YW]- run registered script[?RT]
    [?BYW]Example: [?BBK]run -r script1 --os='linux'[?RT]

  [?BBK]run <script> <args...> [?YW]- run registered script if name doesn't conflict with commands[?RT]
    [?BYW]Example: [?BBK]run script1 --os='linux'[?RT]

  [?BBK]run -r --tagged="<tag1>;<tag2>;<...>" <args...>[?YW] - run scripts with any tag[?RT]
    [?BYW]Example: [?BBK]run -r --tagged="deploy" --os='linux'[?RT]`, []string{"-h", "-help"}, conf)

	p.AddSubcommand("tal", runtime.CreateCli())

	p.AddCommand("-add", runlib.AddHandler,
		`[?GN]Add a local script to the configuration.[?RT]
[?BBK]Supported extensions:[?RT] .py, .sh, .bat, .lua, .task.lua
[?YW]Flags:[?RT]
  [?GN]--force[?RT]    Replace existing script with the same name
[?YW]Examples:[?RT]
  [?BBK]run -add ./deploy.py deploy "Deploy script"[?RT]
  [?BBK]run -add ./build.sh build --force[?RT]`,
		[]string{"path", "name"}, []string{"description"}, false)

	p.AddCommand("-install", runlib.InstallHandler,
		`[?GN]Download and install a script from a URL.[?RT]
[?BBK]Supported URLs:[?RT]
  - Raw file URLs ([?BBK]https://raw.githubusercontent.com/...)[?RT]
  - GitHub blob URLs ([?BBK]github.com/user/repo/blob/branch/path/script.py)[?RT]
  - GitHub simpler URLs ([?BBK]github.com/user/repo@branch/path/script.py)[?RT]
[?YW]Flags:[?RT]
  [?GN]--force[?RT]    Replace existing script with the same name
[?YW]Examples:[?RT]
  [?BBK]run -install https://raw.githubusercontent.com/user/repo/main/deploy.py[?RT]
  [?BBK]run -install github.com/user/repo@branch/script.py myscript[?RT]`,
		[]string{"url", "name"}, []string{"description"}, false)

	p.AddCommand("-remove", runlib.RemoveHandler,
		`[?GN]Remove a script from the configuration.[?RT]
[?YW]Example:[?RT]
  [?BBK]run -remove myscript[?RT]`,
		[]string{"name"}, nil, false)

	p.AddCommand("-list", runlib.ListHandler,
		`[?GN]List all registered scripts.[?RT]
[?BBK]Shows script names, descriptions, and tags.[?RT]`,
		nil, nil, false)

	p.AddCommand("-r",
		runlib.MakeRunHandler(true),
		`[?GN]Run a script by name.[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run -r <name> [args...][?RT]
[?YW]Flags:[?RT]
  [?GN]--tagged="tag1;tag2;..."[?RT]    Run all scripts with any of the specified tags
  [?GN]--parallel[?RT]                   Run tagged scripts in parallel
  [?GN]--args="..."[?RT]                 Explicitly pass arguments (useful when arguments conflict with run flags)
[?YW]Examples:[?RT]
  [?BBK]run -r deploy --env=prod[?RT]
  [?BBK]run -r --tagged="deploy;test" --parallel[?RT]
  [?BBK]run -r --tagged="build" --args="--verbose"[?RT]`,
		nil, nil, true)

	p.AddCommand(tap.DEFAULT_CMD,
		runlib.MakeRunHandler(false),
		`[?GN]Run a script by name (shortcut when name doesn't conflict with commands).[?RT]

[?BBK]Usage:[?RT]
  [?BBK]run <name> [args...][?RT]

[?YW]Example:[?RT]
  [?BBK]run deploy --env=prod[?RT]`,
		nil, nil, true)

	p.AddCommand("-tag",
		runlib.TagHahdler,
		`[?GN]Add tags to a script.[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run -tag <name> <tags...>[?RT]
[?YW]Example:[?RT]
  [?BBK]run -tag deploy prod staging[?RT]`,
		[]string{"script"}, []string{"tag"}, true)

	p.AddCommand("-rm-tag",
		runlib.RmTagHahdler,
		`[?GN]Remove tags from a script.[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run -rm-tag <name> <tags...>[?RT]
[?YW]Example:[?RT]
  [?BBK]run -rm-tag deploy staging[?RT]`,
		[]string{"script"}, []string{"tag"}, true)

	p.AddCommand("-version", func(p *tap.Parser, s []string) error {
		fmt.Println("run v" + run.Version + ", by Pt, Apache 2.0 licence")
		return nil
	},
		`[?GN]Show version and license information.[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run -version[?RT]
[?YW]Example:[?RT]
  [?BBK]run -version[?RT]`, nil, nil, false)

	p.AddCommand("-localmode", func(p *tap.Parser, s []string) error {
		if len(s) == 0 {
			fmt.Println("localmode:", localmode.IsLocalmode(), "| path:", runlib.ConfigDirPath())
			return nil
		}
		switch s[0] {
		case "true":
			localmode.Set(true)
		case "false":
			localmode.Set(false)
		default:
			return fmt.Errorf("Invalid argument")
		}
		return nil
	},
		`[?GN]Set or show the current working mode (global/local).[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run -localmode[?RT]           Show current mode and config path
  [?BBK]run -localmode true[?RT]      Enable local mode (use .run/ in current directory)
  [?BBK]run -localmode false[?RT]     Disable local mode (use ~/run/)
[?YW]Examples:[?RT]
  [?BBK]run -localmode[?RT]
  [?BBK]run -localmode true[?RT]`,
		nil, []string{"mode"}, false)

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
			lp = NewCli()
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

	runlib.RegisterLuaFunc("run_cli", runcli)

	return p
}
