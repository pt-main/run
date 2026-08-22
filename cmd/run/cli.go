package main

import (
	"fmt"

	"github.com/pt-main/run"
	runlib "github.com/pt-main/run/run"
	localmode "github.com/pt-main/run/run/localMode"
	"github.com/pt-main/run/tal/runtime"
	"github.com/pt-main/tap"
	"github.com/pt-main/tycl/cli"
)

func NewCli() *tap.Parser {
	p := tap.NewParser("run", `[?BE]╭─────── [?BRD]Run[?RT]
[?BE]⎬─ [?RT]Simple and powerfull script manager
[?BE]│  [?RT]By [?UE]Pt[?RT], only [?BD]humanmade[?RT].
[?BE]╰───────[?RT]

[?GN]Usage: [?RT]
  [?BYW]Type: [?BBK]run [--<lm|locamode>] <cmd> <args...>

  [?BBK]run <cmd> <args...> [?YW]- execute run command (which described in help)[?RT]
    [?BYW]Example: [?BBK]run -add script.py script1[?RT]
  
  [?BBK]run -r <script> <args...> [?YW]- run registred script[?RT]
    [?BYW]Example: [?BBK]run -r script1 --os='linux'[?RT]

  [?BBK]run <script> <args...> [?YW]- run registred script which name isn't same with cli commands[?RT]
    [?BYW]Example: [?BBK]run script1 --os='linux'[?RT]
  
  [?BBK]run -r --tagged="<tag1>;<tag2>;<...>" <args...>[?YW] - run script with any tag[?RT]
    [?BYW]Example: [?BBK]run -r --tagged="deploy" --os='linux'[?RT]`, []string{"-h", "help"}, tap.DefaultParserConfig())

	p.AddCommand("tycl", func(p *tap.Parser, s []string) error {
		args, err := runlib.ProcessShell(s[0])
		if err != nil {
			return err
		}
		if err = cli.NewCli().Parse(args); err != nil {
			return err
		}
		return nil
	}, tap.DONT_SHOW, []string{"code"}, nil, false)

	p.AddCommand("-add", runlib.AddHandler,
		`Add script (python/bash) with auto run script generation. 
Use --force flag to replace script if it's already added with same name.`,
		[]string{"path", "name"}, []string{"docs"}, false)

	p.AddCommand("-remove", runlib.RemoveHandler,
		`Remove script from global config`,
		[]string{"name"}, nil, false)

	p.AddSubcommand("tal", runtime.CreateCli())

	p.AddCommand("-list", runlib.ListHandler,
		`Show scripts list`, nil, nil, false)

	p.AddCommand("-r",
		runlib.MakeRunHandler(true), `Run scripts`,
		nil, nil, true)

	p.AddCommand(tap.DEFAULT_CMD,
		runlib.MakeRunHandler(false), ``,
		nil, nil, true)

	p.AddCommand("-tag",
		runlib.TagHahdler, `Tag script`,
		[]string{"script"}, []string{"tag"}, true)

	p.AddCommand("-rm-tag",
		runlib.RmTagHahdler, `Remove tags from script`,
		[]string{"script"}, []string{"tag"}, true)

	p.AddCommand("-version", func(p *tap.Parser, s []string) error {
		fmt.Println("run v" + run.Version + ", by Pt, Apache 2.0 licence")
		return nil
	}, `Show version and info`, nil, nil, false)

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
	}, `Set run working mode. Mode - 'true'/'false'`, nil, []string{"mode"}, false)

	return p
}
