package api

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/pt-main/tycl"
	"github.com/pt-main/tycl/format"
	"github.com/pt-main/tycl/shared"
	"github.com/pt-main/tycl/utils"
	lua "github.com/yuin/gopher-lua"
)

// LegacyTyclContract is the contract of configs written before template bodies
// were moved out of the config into the templates dir. It is only used to
// detect such configs and migrate them.
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

// MigrateTemplates moves inline template bodies (the `template` field) into the
// templates dir and replaces them with a `file` reference. It reports whether
// the config was changed.
func MigrateTemplates(cfg *shared.Config) (bool, error) {
	templates, ok := cfg.InnerArrV["templates"]
	if !ok {
		return false, nil
	}
	changed := false
	for _, templ := range templates {
		content, has := templ.StringV["template"]
		if !has {
			continue
		}
		file := templ.StringV["file"]
		if file == "" {
			file = TemplateFileName(templ.StringV["ext"])
			if err := WriteTemplateFile(file, content); err != nil {
				return changed, err
			}
			templ.StringV["file"] = file
		}
		delete(templ.StringV, "template")
		changed = true
	}
	return changed, nil
}

func GetCfg() (*shared.Config, error) {
	file, err := utils.OpenF(ConfigDirConfigPath())
	if err != nil {
		return nil, err
	}
	cfg, errI := tycl.Process(file, TyclContract, true)
	if errI != nil {
		// a config that still keeps template bodies inline
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
		if err := UpdateConfig(cfg); err != nil {
			return nil, err
		}
	}
	return cfg, nil
}

func AddScript(conf *shared.Config, script, rawScriptName, scriptName, docs string, force bool) error {
	rawScriptName = scriptName + "_" + strings.ReplaceAll(rawScriptName, "/", "__")

	runScript := ""

	kept := []*shared.Config{}
	for _, existing := range conf.InnerArrV["scripts"] {
		name := existing.StringV["script"]
		if name == scriptName && !force {
			return fmt.Errorf("Can't add script: script already added. Use --force to replace script.")
		}
		if name != scriptName {
			kept = append(kept, existing)
		}
	}
	conf.InnerArrV["scripts"] = kept
	storeBase := true
	matched := false
	ext := filepath.Ext(rawScriptName)

	if strings.HasSuffix(rawScriptName, ".nd.task.lua") { // nd - no deps
		matched = true
		ext = ".nd.task.lua"
		runScript = TalRunScriptTemplate(rawScriptName)
	} else if strings.HasSuffix(rawScriptName, ".task.lua") {
		matched = true
		ext = ".task.lua"
		runScript = TalRunScriptTemplate(rawScriptName)
	}

	var fallbackFile string

	render := func(templ string) error {
		tpl, err := template.New(ext).Parse(templ)
		if err != nil {
			return fmt.Errorf("Add script: parsing extension template: %v", err)
		}
		var b strings.Builder
		if err := tpl.Execute(&b, map[string]string{"ext": ext, "name": rawScriptName}); err != nil {
			return fmt.Errorf("Add script: executing extension template: %v", err)
		}
		runScript = b.String()
		return nil
	}

	if !matched {
		for _, tmpl := range conf.InnerArrV["templates"] {
			tmplExt := tmpl.StringV["ext"]
			file := tmpl.StringV["file"]

			if tmplExt == "" {
				if file != "" {
					fallbackFile = file
				}
				continue
			}

			if strings.HasSuffix(rawScriptName, tmplExt) {
				templ, err := ReadTemplateFile(file)
				if err != nil {
					return err
				}
				if err := render(templ); err != nil {
					return err
				}
				matched = true
			}
		}
	}

	if !matched && fallbackFile != "" {
		templ, err := ReadTemplateFile(fallbackFile)
		if err != nil {
			return err
		}
		if err := render(templ); err != nil {
			return err
		}
		matched = true
	}

	if !matched {
		switch ext {
		case ".py":
			matched = true
			runScript = PythonRunScriptTemplate(rawScriptName)
		case ".sh":
			matched = true
			runScript = BashRunScriptTemplate(rawScriptName)
		case ".bat":
			matched = true
			runScript = BatRunScriptTemplate(rawScriptName)
		case ".lua":
			matched = true
			runScript = script
			storeBase = false
		}
	}
	if !matched {
		return fmt.Errorf("Unsupportable file extension: %v", ext)
	}

	conf.InnerArrV["scripts"] = append(conf.InnerArrV["scripts"],
		NewScriptConfig(scriptName, scriptName, docs, ext, nil))
	if err := NewRunScript(scriptName, runScript); err != nil {
		return err
	}
	if storeBase {
		file, _, _ := strings.Cut(rawScriptName, "/")
		if err := NewScript(file, script); err != nil {
			return err
		}
	}
	return nil
}

func AddTemplate(conf *shared.Config, ext, template string, force bool) error {
	kept := []*shared.Config{}
	for _, tmpl := range conf.InnerArrV["templates"] {
		if tmplExt := tmpl.StringV["ext"]; tmplExt == ext {
			if !force {
				return fmt.Errorf("Can't add template: extension duplicate and has no force flag")
			}
			continue
		}
		kept = append(kept, tmpl)
	}
	file := TemplateFileName(ext)
	if err := WriteTemplateFile(file, template); err != nil {
		return err
	}
	c := shared.NewNilConfig()
	c.StringV["ext"] = ext
	c.StringV["file"] = file
	conf.InnerArrV["templates"] = append(kept, c)
	return nil
}

func RemoveTemplate(conf *shared.Config, ext string) error {
	kept := []*shared.Config{}
	for _, tmpl := range conf.InnerArrV["templates"] {
		if tmpl.StringV["ext"] == ext {
			if err := RemoveTemplateFile(tmpl.StringV["file"]); err != nil {
				return err
			}
			continue
		}
		kept = append(kept, tmpl)
	}
	conf.InnerArrV["templates"] = kept
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
	var scriptName string
	for _, script := range cfg.InnerArrV["scripts"] {
		if script.StringV["name"] == name {
			scriptName = script.StringV["script"]
			break
		}
	}
	if scriptName == "" {
		return fmt.Errorf("Script is not found")
	}
	file, err := utils.OpenF(filepath.Join(ConfigDirScriptsPath(), scriptName+".lua"))
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

// Upconf writes the config back to disk unless err is set, so a handler can
// pass the result of a mutating call directly.
func Upconf(conf *shared.Config, err error) error {
	if err != nil {
		return err
	}
	return UpdateConfig(conf)
}
