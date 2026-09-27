package runcli

import (
	runlib "github.com/pt-main/run/run/runcli/handlers"
	tap "github.com/pt-main/tap/go"
)

func NewManage() (p *tap.Parser, err error) {
	manageP := tap.NewParser("manage", "[?GN]Manage run data.[?RT]", []string{"help", "-help", "-h"}, tap.DefaultParserConfig())

	manageP.AddCommand("script-add", runlib.AddHandler,
		`[?GN]Add a local script to the configuration.[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run manage script-add <path> <name> [description] [--force][?RT]
[?BBK]Supported extensions:[?RT] .py, .sh, .bat, .lua, .task.lua
[?YW]Flags:[?RT]
  [?GN]--force[?RT]    Replace existing script with the same name
[?YW]Examples:[?RT]
  [?BBK]run manage script-add ./deploy.py deploy "Deploy script"[?RT]
  [?BBK]run manage scradd ./build.sh build --force[?RT]`,
		[]string{"path", "name"}, []string{"description"}, false)
	if err = manageP.AddAlias("scradd", "script-add"); err != nil {
		return
	}

	manageP.AddCommand("script-remove", runlib.RemoveHandler,
		`[?GN]Remove a script from the configuration.[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run manage script-remove <name>[?RT]
[?YW]Example:[?RT]
  [?BBK]run manage script-remove myscript[?RT]`,
		[]string{"name"}, nil, false)
	if err = manageP.AddAlias("screm", "script-remove"); err != nil {
		return
	}

	manageP.AddCommand("templ-add", runlib.TemplateAdd,
		`[?GN]Add a custom template for a file extension.[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run manage templ-add <ext> --source="..." [--force][?RT]
  [?BBK]run manage templ-add <ext> <file> [--force][?RT]
[?YW]Flags:[?RT]
  [?GN]--source="..."[?RT]  Template content as string (instead of file)
  [?GN]--force[?RT]         Overwrite existing template for this extension
[?YW]Examples:[?RT]
  [?BBK]run manage templ-add ".go" templ.txt --force[?RT]
  [?BBK]run manage templ-add ".go" --source="..."[?RT]
[?YW]Note:[?RT]
  [?BBK]Template body is saved to the templates/ dir, the config only keeps the file reference.[?RT]`,
		[]string{"ext"}, []string{"file"}, false)
	if err = manageP.AddAlias("tladd", "templ-add"); err != nil {
		return
	}

	manageP.AddCommand("templ-remove", runlib.TemplateRem,
		`[?GN]Remove a template for an extension.[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run manage templ-remove <ext>[?RT]
[?YW]Example:[?RT]
  [?BBK]run manage templ-rem ".go"[?RT]
[?YW]Note:[?RT]
  [?BBK]The template file in the templates/ dir is removed too.[?RT]`,
		[]string{"ext"}, nil, false)
	if err = manageP.AddAlias("tlrem", "templ-remove"); err != nil {
		return
	}

	manageP.AddCommand("tag",
		runlib.TagHahdler,
		`[?GN]Add or remove tags for a script.[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run manage tag <script> <tag1> <tag2> ... [?RT]
  Prefix tag with '!' to remove it.
[?YW]Examples:[?RT]
  [?BBK]run manage tag deploy prod staging[?RT]
  [?BBK]run manage tag deploy !prod[?RT]`,
		[]string{"script"}, []string{"tag"}, true)

	manageP.AddCommand("list", runlib.ListHandler,
		`[?GN]List all registered scripts.[?RT]
[?BBK]Shows script names, descriptions, and tags.[?RT]
[?YW]Example:[?RT]
  [?BBK]run manage list[?RT]`,
		nil, nil, false)

	manageP.AddCommand("install", runlib.InstallHandler,
		`[?GN]Download and install a script from any URL.[?RT]
[?BBK]Supported URLs:[?RT]
  - Raw file URLs [?BBK](https://raw.githubusercontent.com/...)[?RT]
  - GitHub blob URLs [?BBK](github.com/user/repo/blob/branch/path/script.py)[?RT]
  - GitHub simpler URLs [?BBK](github.com/user/repo@branch/path/script.py)[?RT]
  - Running installation script [?BBK](github.com/user/repo@branch/path/run.task.lua)[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run manage install <url> [name] [description] [--force] [--args="..."][?RT]
[?YW]Flags:[?RT]
  [?GN]--force[?RT]        Replace existing script with the same name
  [?GN]--args="..."[?RT]   Pass args to run script (for run.task.lua)
[?YW]Examples:[?RT]
  [?BBK]run manage install https://raw.githubusercontent.com/user/repo/main/deploy.py[?RT]
  [?BBK]run manage install github.com/user/repo@branch/script.py myscript[?RT]
  [?BBK]run manage install github.com/user/repo@branch/run.task.lua[?RT]`,
		[]string{"url"}, []string{"name", "description"}, false)

	return manageP, nil
}
