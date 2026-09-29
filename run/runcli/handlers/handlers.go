package runlib

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"sync"

	"github.com/pt-main/run/run/api"
	tap "github.com/pt-main/tap/go"
	"github.com/pt-main/tap/go/color"
	"github.com/pt-main/tycl/shared"
	"github.com/pt-main/tycl/utils"
)

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
	conf, err := api.GetCfg()
	if err != nil {
		return err
	}
	return api.Upconf(conf, api.AddScript(conf, script, s[0], s[1], docs, force))
}

func RemoveHandler(p *tap.Parser, s []string) error {
	cfg, err := api.GetCfg()
	if err != nil {
		return err
	}
	newScripts := []*shared.Config{}
	for _, script := range cfg.InnerArrV["scripts"] {
		name := script.StringV["name"]
		if name != s[0] {
			newScripts = append(newScripts, script)
			continue
		}
		if err := api.RemoveRunScript(script.StringV["script"]); err != nil {
			return err
		}
	}
	cfg.InnerArrV["scripts"] = newScripts
	return api.UpdateConfig(cfg)
}

func ListHandler(p *tap.Parser, s []string) error {
	cfg, err := api.GetCfg()
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

func MakeRunHandler(hasRawArgs bool) func(p *tap.Parser, s []string) error {
	return func(p *tap.Parser, s []string) error {
		cfg, err := api.GetCfg()
		if err != nil {
			return err
		}
		idx := 0
		if len(s) > 0 {
			idx += 1
		}
		if len(p.RawArgs) > 0 && slices.Contains([]string{"--gm", "--globalmode",
			"--lm", "--localmode"}, p.RawArgs[0]) {
			idx += 1
		}
		if hasRawArgs {
			idx += 1
		}
		var args []string = nil
		args_, ok := p.Flags["args"]
		if ok {
			args, err = api.ProcessShell(args_)
			if err != nil {
				return err
			}
			if args == nil {
				args = []string{}
			}
		}
		if args == nil && len(p.RawArgs) > 0 {
			args = p.RawArgs[idx:]
		}
		if tags_, ok := p.Flags["tagged"]; ok {
			_, parallel := p.Flags["parallel"]
			tags := strings.Split(tags_, ";")
			errs := []string{}
			var errsMu sync.Mutex
			var wg sync.WaitGroup

			runOne := func(name string) error {
				if err := api.RunScript(cfg, name, args); err != nil {
					p.Print("verbose", "[?RD]Err[?YW]:[RT] %v", err)
					return err
				}
				p.Print("verbose", "[?GN]Ok[?RT]")
				return nil
			}

			for _, script := range cfg.InnerArrV["scripts"] {
				for _, tag := range script.StringArrV["tags"] {
					if !slices.Contains(tags, tag) {
						continue
					}
					name := script.StringV["name"]
					p.Print("verbose", "Run %v: ", name)
					if parallel {
						wg.Add(1)
						go func() {
							defer wg.Done()
							if err := runOne(name); err != nil {
								errsMu.Lock()
								errs = append(errs, err.Error())
								errsMu.Unlock()
							}
						}()
					} else if err := runOne(name); err != nil {
						errs = append(errs, err.Error())
					}
				}
			}

			wg.Wait()
			if len(errs) == 0 {
				return nil
			}
			return errors.New(" - " + strings.Join(errs, "\n - "))
		}

		if len(s) < 1 {
			return fmt.Errorf("Invalid argument length: need more or equals to 1")
		}
		name := s[0]
		p.Print("verbose", "Run %v: ", name)
		return api.RunScript(cfg, name, args)
	}
}

func TagHahdler(p *tap.Parser, s []string) error {
	cfg, err := api.GetCfg()
	if err != nil {
		return err
	}
	addTags := []string{}
	rmTags := []string{}
	for _, tag := range s[1:] {
		if strings.HasPrefix(tag, "!") {
			rmTags = append(rmTags, tag[1:])
		} else {
			addTags = append(addTags, tag)
		}
	}
	for _, script := range cfg.InnerArrV["scripts"] {
		scriptName := script.StringV["name"]
		if scriptName == s[0] {
			tags := append(script.StringArrV["tags"], addTags...)
			newTags := []string{}
			for _, tag := range tags {
				if !slices.Contains(rmTags, tag) {
					newTags = append(newTags, tag)
				}
			}
			script.StringArrV["tags"] = newTags
			break
		}
	}
	return api.UpdateConfig(cfg)
}
