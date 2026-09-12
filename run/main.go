package runlib

import (
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/tycl"
	"github.com/pt-main/tycl/format"
	"github.com/pt-main/tycl/shared"
	"github.com/pt-main/tycl/utils"
)

func GetCfg() (*shared.Config, error) {
	file, err := utils.OpenF(ConfigDirConfigPath())
	if err != nil {
		return nil, err
	}
	var errI core.ErrorInterface
	cfg, errI := tycl.Process(file, tyclContract, true)
	if errI != nil {
		return cfg, fmt.Errorf(format.FormatError(errI))
	}
	return cfg, nil
}

func AddScript(script, rawScriptName, scriptName, docs string, force bool) error {
	conf, err := GetCfg()
	if err != nil {
		return err
	}

	rawScriptName = scriptName + "_" + strings.ReplaceAll(rawScriptName, "/", "__")

	runScript := ""

	newScripts := []*shared.Config{}
	for _, script := range conf.InnerArrV["scripts"] {
		name := script.StringV["script"]
		if name == scriptName && !force {
			return fmt.Errorf("Can't add script: script already added. Use --force to replace script.")
		}
		if name != scriptName {
			newScripts = append(newScripts, script)
		}
	}
	conf.InnerArrV["scripts"] = newScripts
	addScript := true

	processed := false
	ext := filepath.Ext(rawScriptName)

	if strings.HasSuffix(rawScriptName, ".nd.task.lua") { // nd - no deps
		processed = true
		ext = ".nd.task.lua"
		runScript = TalRunScriptTemplate(rawScriptName, true)
	} else if strings.HasSuffix(rawScriptName, ".task.lua") {
		processed = true
		ext = ".task.lua"
		runScript = TalRunScriptTemplate(rawScriptName, true)
	}

	if !processed {
		templs := conf.InnerArrV["templates"]
		for _, cfg := range templs {
			ext := cfg.StringV["ext"]
			templ := cfg.StringV["template"]
			if strings.HasSuffix(rawScriptName, ext) {
				tpl, err := template.New(ext).Parse(templ)
				if err != nil {
					return fmt.Errorf("Add script: parsing extension template: %v", err)
				}
				var b strings.Builder
				err = tpl.Execute(&b, map[string]string{})
				if err != nil {
					return fmt.Errorf("Add script: executing extension template: %v", err)
				}
				runScript = b.String()
				processed = true
			}
		}
	}

	if !processed {
		switch ext {
		case ".py":
			processed = true
			runScript = PythonRunScriptTemplate(rawScriptName)
		case ".sh":
			processed = true
			runScript = BashRunScriptTemplate(rawScriptName)
		case ".bat":
			processed = true
			runScript = BatRunScriptTemplate(rawScriptName)
		case ".lua":
			processed = true
			runScript = script
			addScript = false
		}
	}
	if !processed {
		return fmt.Errorf("Unsupportable file extension: %v", ext)
	}
	conf.InnerArrV["scripts"] = append(conf.InnerArrV["scripts"], NewScriptConfig(scriptName, scriptName, docs, ext, nil))
	if err := NewRunScript(scriptName, runScript); err != nil {
		return err
	}
	if err := UpdateConfig(conf); err != nil {
		return err
	}
	if addScript {
		fpsplit := strings.Split(rawScriptName, "/")
		file := fpsplit[0]
		if err := NewScript(file, script); err != nil {
			return err
		}
	}
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
