package runlib

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/pt-main/tap"
	"github.com/pt-main/tap/color"
	"github.com/pt-main/tycl/shared"
	"github.com/pt-main/tycl/utils"
)

func AddHandler(p *tap.Parser, s []string) error {
	ext := filepath.Ext(s[0])
	fpsplit := strings.Split(s[0], "/")
	file := fpsplit[0]
	if len(fpsplit) > 0 {
		file = fpsplit[len(fpsplit)-1]
	}
	runScript := ""
	conf, err := GetCfg()
	if err != nil {
		return err
	}
	_, force := p.Flags["force"]
	newScripts := []*shared.Config{}
	for _, script := range conf.InnerArrV["scripts"] {
		name := script.StringV["script"]
		if name == s[1] && !force {
			return fmt.Errorf("Can't add script: script already added. Use --force to replace script.")
		}
		if name != s[1] {
			newScripts = append(newScripts, script)
		}
	}
	conf.InnerArrV["scripts"] = newScripts
	script, err := utils.OpenF(s[0])
	if err != nil {
		return err
	}
	addScript := true
	switch ext {
	case ".py":
		runScript = PythonRunScriptTemplate(s[0])
	case ".sh":
		runScript = BashRunScriptTemplate(s[0])
	case ".bat":
		runScript = BatRunScriptTemplate(s[1])
	case ".lua":
		runScript = script
		addScript = false
	default:
		return fmt.Errorf("Unsupportable file extension: %v", ext)
	}
	docs := ""
	if len(s) > 2 {
		docs = s[2]
	}
	conf.InnerArrV["scripts"] = append(conf.InnerArrV["scripts"], NewScriptConfig(s[1], s[1], docs, script, ext, nil))
	if err := NewRunScript(s[1], runScript); err != nil {
		return err
	}
	if err := UpdateConfig(conf); err != nil {
		return err
	}
	if addScript {
		if err := NewScript(file, script); err != nil {
			return err
		}
	}
	return nil
}

func RemoveHandler(p *tap.Parser, s []string) error {
	cfg, err := GetCfg()
	if err != nil {
		return err
	}
	newScripts := []*shared.Config{}
	for _, script := range cfg.InnerArrV["scripts"] {
		name := script.StringV["name"]
		if name != s[0] {
			newScripts = append(newScripts, script)
		}
	}
	cfg.InnerArrV["scripts"] = newScripts
	return UpdateConfig(cfg)
}

func ListHandler(p *tap.Parser, s []string) error {
	cfg, err := GetCfg()
	if err != nil {
		return err
	}
	color.PrintlnColored("[?GN]╭─────── [?YW] Scripts [?RT]")
	linestart := "[?GN]│     [?RT]"
	for _, script := range cfg.InnerArrV["scripts"] {
		name := script.StringV["name"]
		ext := script.StringV["ext"]
		if ext != "" {
			ext = "[?BBK] (" + ext + ")"
		}
		description := script.StringV["description"]
		color.PrintColored("[?GN]⎬─ [?YW]%v%v[?RT]", name, ext)
		if description != "" {
			color.PrintlnColored(":\n"+linestart+"%v[?RT]", strings.ReplaceAll(description, "\n", "\n"+linestart))
		} else {
			fmt.Println()
		}
	}
	color.PrintlnColored("[?GN]╰───────[?RT]")
	return nil
}

func RunScript(cfg *shared.Config, name string, rArgs []string) error {
	var scriptPath string
	for _, script := range cfg.InnerArrV["scripts"] {
		scriptName := script.StringV["name"]
		scriptPath_ := script.StringV["script"]
		if scriptName == name {
			scriptPath = scriptPath_
			break
		}
	}
	if scriptPath == "" {
		return fmt.Errorf("Script is not found")
	}
	file, err := utils.OpenF(filepath.Join(ConfigDirScriptsPath(), scriptPath+".lua"))
	if err != nil {
		return err
	}
	if err := NewLuaState(rArgs).DoString(file); err != nil {
		return err
	}
	return nil
}

func MakeRunHandler(hasR bool) func(p *tap.Parser, s []string) error {
	return func(p *tap.Parser, s []string) error {
		cfg, err := GetCfg()
		if err != nil {
			return err
		}
		idx := 1
		if slices.Contains([]string{"--gm", "--globalmode",
			"--lm", "--localmode"}, p.RawArgs[0]) {
			idx += 1
		}
		if hasR {
			idx += 1
		}
		var args []string = nil
		args_, ok := p.Flags["args"]
		if ok {
			args, err = ProcessShell(args_)
			if err != nil {
				return err
			}
			if args == nil {
				args = []string{}
			}
		}
		if args == nil {
			args = p.RawArgs[idx:]
		}
		if tags_, ok := p.Flags["tagged"]; ok {
			_, parallel := p.Flags["parallel"]
			tags := strings.Split(tags_, ";")
			errs := []string{}
			goru := 0
			for _, script := range cfg.InnerArrV["scripts"] {
				scrTags := script.StringArrV["tags"]
				scriptName := script.StringV["name"]
				for _, tag := range scrTags {
					if slices.Contains(tags, tag) {
						p.Print("verbose", "Run %v: ", scriptName)
						if parallel {
							go func() {
								goru += 1
								if err := RunScript(cfg, scriptName, args); err != nil {
									p.Print("verbose", "[?RD]Err[?YW]:[RT] %v", err)
									errs = append(errs, err.Error())
								} else {
									p.Print("verbose", "[?GN]Ok[?RT]")
								}
								goru -= 1
							}()
							time.Sleep(time.Second / 500)
						} else {
							if err := RunScript(cfg, scriptName, args); err != nil {
								p.Print("verbose", "[?RD]Err[?YW]:[RT] %v", err)
								errs = append(errs, err.Error())
							} else {
								p.Print("verbose", "[?GN]Ok[?RT]")
							}
						}
					}
				}
			}
			p.Print("verbose", "[?GN]Gorutines: %v[?RT]", goru)
			for goru != 0 {
				time.Sleep(time.Second / 500)
			}
			if len(errs) == 0 {
				return nil
			}
			return fmt.Errorf(" - " + strings.Join(errs, "\n - "))
		} else {
			if len(s) < 1 {
				return fmt.Errorf("Invalid argument length: need more or equals to 1")
			}
			name := s[0]
			p.Print("verbose", "Run %v: ", name)
			return RunScript(cfg, name, args)
		}
	}
}

func RmTagHahdler(p *tap.Parser, s []string) error {
	cfg, err := GetCfg()
	if err != nil {
		return err
	}
	for _, script := range cfg.InnerArrV["scripts"] {
		scriptName := script.StringV["name"]
		if scriptName == s[0] {
			newTags := []string{}
			for _, tag := range script.StringArrV["tags"] {
				for _, rm := range s[1:] {
					if tag != rm {
						newTags = append(newTags, tag)
					}
				}
			}
			script.StringArrV["tags"] = newTags
			break
		}
	}
	return UpdateConfig(cfg)
}

func TagHahdler(p *tap.Parser, s []string) error {
	cfg, err := GetCfg()
	if err != nil {
		return err
	}
	for _, script := range cfg.InnerArrV["scripts"] {
		scriptName := script.StringV["name"]
		if scriptName == s[0] {
			script.StringArrV["tags"] = append(script.StringArrV["tags"], s[1:]...)
			break
		}
	}
	return UpdateConfig(cfg)
}
