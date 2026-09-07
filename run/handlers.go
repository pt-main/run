package runlib

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/pt-main/tap"
	"github.com/pt-main/tap/color"
	"github.com/pt-main/tycl/shared"
	"github.com/pt-main/tycl/utils"
)

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

	ext := filepath.Ext(rawScriptName)
	if strings.HasSuffix(rawScriptName, ".task.lua") {
		ext = ".task.lua"
		runScript = TalRunScriptTemplate(rawScriptName)
	}
	switch ext {
	case ".py":
		runScript = PythonRunScriptTemplate(rawScriptName)
	case ".sh":
		runScript = BashRunScriptTemplate(rawScriptName)
	case ".bat":
		runScript = BatRunScriptTemplate(rawScriptName)
	case ".task.lua":
		runScript = TalRunScriptTemplate(rawScriptName)
	case ".lua":
		runScript = script
		addScript = false
	default:
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

func AddHandler(p *tap.Parser, s []string) error {
	_, force := p.Flags["force"]

	script, err := utils.OpenF(s[0])
	if err != nil {
		return err
	}
	docs := ""
	if len(s) > 2 {
		docs = s[2]
	}
	return AddScript(script, s[0], s[1], docs, force)
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

func downloadScript(url string) (content string, fileName string, err error) {
	if strings.Contains(url, "github.com/") {
		rawURL, fname, err := parseGitHubURL(url)
		if err != nil {
			return "", "", err
		}
		url = rawURL
		fileName = fname
	} else {

		fileName = filepath.Base(url)
		if idx := strings.Index(fileName, "?"); idx != -1 {
			fileName = fileName[:idx]
		}
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("HTTP error: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}
	return string(data), fileName, nil
}

func parseGitHubURL(rawURL string) (rawContentURL string, fileName string, err error) {

	u := rawURL
	if strings.HasPrefix(u, "https://") {
		u = strings.TrimPrefix(u, "https://")
	} else if strings.HasPrefix(u, "http://") {
		u = strings.TrimPrefix(u, "http://")
	}

	if !strings.HasPrefix(u, "github.com/") {
		return "", "", fmt.Errorf("not a GitHub URL")
	}
	u = strings.TrimPrefix(u, "github.com/")

	parts := strings.SplitN(u, "@", 2)
	if len(parts) == 2 {
		repo := parts[0]
		rest := parts[1]
		slashIdx := strings.Index(rest, "/")
		if slashIdx == -1 {
			return "", "", fmt.Errorf("missing path after ref")
		}
		ref := rest[:slashIdx]
		path := rest[slashIdx+1:]
		rawContentURL = fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", repo, ref, path)
		fileName = filepath.Base(path)
		return rawContentURL, fileName, nil
	}

	idx := strings.Index(u, "/blob/")
	if idx == -1 {
		return "", "", fmt.Errorf("invalid GitHub URL: missing '@' or '/blob/'")
	}
	repoPart := u[:idx]
	rest := u[idx+len("/blob/"):]
	slashIdx := strings.Index(rest, "/")
	if slashIdx == -1 {
		return "", "", fmt.Errorf("invalid blob URL: missing branch/path")
	}
	branch := rest[:slashIdx]
	path := rest[slashIdx+1:]
	rawContentURL = fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", repoPart, branch, path)
	fileName = filepath.Base(path)
	return rawContentURL, fileName, nil
}

func InstallHandler(p *tap.Parser, s []string) error {
	if len(s) < 1 {
		return fmt.Errorf("need at least URL")
	}

	url := s[0]
	scriptName := ""
	docs := ""

	if len(s) > 1 {
		scriptName = s[1]
	}
	if len(s) > 2 {
		docs = s[2]
	}

	content, rawName, err := downloadScript(url)
	if err != nil {
		return err
	}

	if scriptName == "" {
		ext := filepath.Ext(rawName)
		scriptName = strings.TrimSuffix(rawName, ext)
	}

	_, force := p.Flags["force"]
	return AddScript(content, rawName, scriptName, docs, force)
}
