package api

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/tycl"
	"github.com/pt-main/tycl/format"
	"github.com/pt-main/tycl/shared"
	"github.com/pt-main/tycl/utils"
	lua "github.com/yuin/gopher-lua"
)

// LegacyTyclContract is the old contract where templates were stored inline
// in the config. It is only used to detect and migrate old configs.
const LegacyTyclContract = `
flexible {
	scripts: objects = flexible {
		name: string,
		script: string,
		description: string,
		tags: strings,
		ext: string,
	},
	templates: objects = flexible {
		ext: string,
		template: string,
	},
}
`

// MigrateTemplates moves inline templates (old `template` field) into the
// templates dir and replaces them with a `file` reference. Returns true when
// the config has been changed.
func MigrateTemplates(cfg *shared.Config) (bool, error) {
	if _, ok := cfg.InnerArrV["templates"]; !ok {
		return false, nil
	}
	changed := false
	templates := []*shared.Config{}
	for _, templ := range cfg.InnerArrV["templates"] {
		content, has := templ.StringV["template"]
		if !has {
			templates = append(templates, templ)
			continue
		}
		ext := templ.StringV["ext"]
		file, hasFile := templ.StringV["file"]
		if hasFile && file != "" {
			// already migrated, drop the legacy field
			delete(templ.StringV, "template")
			templates = append(templates, templ)
			changed = true
			continue
		}
		file = TemplateFileName(ext)
		if err := WriteTemplateFile(file, content); err != nil {
			return changed, err
		}
		delete(templ.StringV, "template")
		templ.StringV["file"] = file
		templates = append(templates, templ)
		changed = true
	}
	cfg.InnerArrV["templates"] = templates
	return changed, nil
}

func GetCfg() (*shared.Config, error) {
	file, err := utils.OpenF(ConfigDirConfigPath())
	if err != nil {
		return nil, err
	}
	var errI core.ErrorInterface
	cfg, errI := tycl.Process(file, TyclContract, true)
	if errI != nil {
		// old configs keep template contents inline, try to migrate them
		var legacyErrI core.ErrorInterface
		legacyCfg, legacyErrI := tycl.Process(file, LegacyTyclContract, true)
		if legacyErrI != nil {
			return cfg, errors.New(format.FormatError(errI))
		}
		migrated, mErr := MigrateTemplates(legacyCfg)
		if mErr != nil {
			return legacyCfg, mErr
		}
		if !migrated {
			return cfg, errors.New(format.FormatError(errI))
		}
		if err := UpdateConfig(legacyCfg); err != nil {
			return legacyCfg, err
		}
		return legacyCfg, nil
	}
	if _, ok := cfg.InnerArrV["templates"]; !ok {
		cfg.InnerArrV["templates"] = []*shared.Config{}
		err = UpdateConfig(cfg)
		if err != nil {
			return nil, err
		}
	}
	return cfg, nil
}

func AddScript(conf *shared.Config, script, rawScriptName, scriptName, docs string, force bool) error {
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
		runScript = TalRunScriptTemplate(rawScriptName)
	} else if strings.HasSuffix(rawScriptName, ".task.lua") {
		processed = true
		ext = ".task.lua"
		runScript = TalRunScriptTemplate(rawScriptName)
	}

	var fallbackFile *string

	parseTempl := func(templ string) error {
		tpl, err := template.New(ext).Parse(templ)
		if err != nil {
			return fmt.Errorf("Add script: parsing extension template: %v", err)
		}
		var b strings.Builder
		err = tpl.Execute(&b, map[string]string{"ext": ext, "name": rawScriptName})
		if err != nil {
			return fmt.Errorf("Add script: executing extension template: %v", err)
		}
		runScript = b.String()
		return nil
	}

	if !processed {
		templs := conf.InnerArrV["templates"]
		for _, cfg := range templs {
			ext := cfg.StringV["ext"]
			file := cfg.StringV["file"]

			if ext == "" {
				if file != "" {
					f := file
					fallbackFile = &f
				}
			}

			if strings.HasSuffix(rawScriptName, ext) && ext != "" {
				templ, err := ReadTemplateFile(file)
				if err != nil {
					return err
				}
				if err := parseTempl(templ); err != nil {
					return err
				}
				processed = true
			}
		}
	}

	if fallbackFile != nil && !processed {
		templ, err := ReadTemplateFile(*fallbackFile)
		if err != nil {
			return err
		}
		if err := parseTempl(templ); err != nil {
			return err
		}
		processed = true
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
	if addScript {
		fpsplit := strings.Split(rawScriptName, "/")
		file := fpsplit[0]
		if err := NewScript(file, script); err != nil {
			return err
		}
	}
	return nil
}

func AddTemplate(conf *shared.Config, ext, template string, force bool) error {
	templates := conf.InnerArrV["templates"]
	templatesA := []*shared.Config{}
	for _, templ := range templates {
		tExt := templ.StringV["ext"]
		if tExt == ext && !force {
			return fmt.Errorf("Can't add template: extension duplicate and has no force flag")
		} else if tExt != ext {
			templatesA = append(templatesA, templ)
		}
	}
	file := TemplateFileName(ext)
	if err := WriteTemplateFile(file, template); err != nil {
		return err
	}
	c := shared.NewNilConfig()
	c.StringV["ext"] = ext
	c.StringV["file"] = file
	templatesA = append(templatesA, c)
	conf.InnerArrV["templates"] = templatesA
	return nil
}

func RemoveTemplate(conf *shared.Config, ext string) error {
	templates := conf.InnerArrV["templates"]
	templatesA := []*shared.Config{}
	for _, templ := range templates {
		tExt := templ.StringV["ext"]
		if tExt == ext {
			if err := RemoveTemplateFile(templ.StringV["file"]); err != nil {
				return err
			}
			continue
		}
		templatesA = append(templatesA, templ)
	}
	conf.InnerArrV["templates"] = templatesA
	return nil
}

// RunScript executes the wrapper of the given script inside a fresh Lua state.
//
// Every generated wrapper ends with os.exit(...). gopher-lua implements os.exit
// as a real os.Exit, so running two scripts in a single process (tagged runs,
// run_script from another wrapper) would kill the process after the first one.
// osTableWithExit replaces os.exit with a recorded code, so the wrapper returns
// normally and the caller decides what to do with the exit code.
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
	code := 0
	L := NewLuaState(rArgs)
	L.SetGlobal("os", osTableWithExit(L, &code))
	if err := L.DoString(file); err != nil {
		return err
	}
	if code != 0 {
		return fmt.Errorf("Script %q exited with code %d", name, code)
	}
	return nil
}

// osTableWithExit copies the standard os table, replacing exit with a version
// that records the requested exit code instead of terminating the process.
func osTableWithExit(L *lua.LState, code *int) *lua.LTable {
	os, ok := L.GetGlobal("os").(*lua.LTable)
	if !ok {
		return L.NewTable()
	}
	res := L.NewTable()
	os.ForEach(func(k, v lua.LValue) {
		res.RawSetString(k.String(), v)
	})
	res.RawSetString("exit", L.NewFunction(func(L *lua.LState) int {
		*code = L.OptInt(1, 0)
		return 0
	}))
	return res
}

func Upconf(conf *shared.Config, err error) error {
	if err != nil {
		return err
	}
	if err := UpdateConfig(conf); err != nil {
		return err
	}
	return nil
}
