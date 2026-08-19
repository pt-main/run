package runlib

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"

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
	switch ext {
	case ".py":
		runScript = PythonRunScriptTemplate(s[0])
	case ".sh":
		runScript = BashRunScriptTemplate(s[0])
	case ".bat":
		runScript = BatRunScriptTemplate(s[1])
	case ".lua":
		runScript = LuaRunScriptTemplate(s[1])
	default:
		return fmt.Errorf("Unsupportable file extension: %v", ext)
	}
	docs := ""
	if len(s) > 2 {
		docs = s[2]
	}
	script, err := utils.OpenF(s[0])
	if err != nil {
		return err
	}
	conf.InnerArrV["scripts"] = append(conf.InnerArrV["scripts"], NewScriptConfig(s[1], s[1], docs, script, ext, nil))
	if err := NewRunScript(s[1], runScript); err != nil {
		return err
	}
	if err := UpdateConfig(conf); err != nil {
		return err
	}
	if err := NewScript(file, script); err != nil {
		return err
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
	if err := NewLuaState(rArgs[1:]).DoString(file); err != nil {
		return err
	}
	return nil
}

func RunHandler(p *tap.Parser, s []string) error {
	cfg, err := GetCfg()
	if err != nil {
		return err
	}
	if tags_, ok := p.Flags["tagged"]; ok {
		tags := strings.Split(tags_, ";")
		for _, script := range cfg.InnerArrV["scripts"] {
			scrTags := script.StringArrV["tags"]
			scriptName := script.StringV["name"]
			for _, tag := range scrTags {
				if slices.Contains(tags, tag) {
					if err := RunScript(cfg, scriptName, p.RawArgs[1:]); err != nil {
						return err
					}
				}
			}
		}
		return nil
	} else {
		if len(s) < 1 {
			return fmt.Errorf("Invalid argument length: need more or equals to 1")
		}
		name := s[0]
		return RunScript(cfg, name, p.RawArgs)
	}
}

func TagHahdler(p *tap.Parser, s []string) error {
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

func RmTagHahdler(p *tap.Parser, s []string) error {
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
