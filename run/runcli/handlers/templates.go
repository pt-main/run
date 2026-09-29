package runlib

import (
	"fmt"

	"github.com/pt-main/run/run/api"
	tap "github.com/pt-main/tap/go"
	"github.com/pt-main/tycl/utils"
)

func TemplateAdd(p *tap.Parser, s []string) (err error) {
	template := ""
	if source, has := p.Flags["source"]; has {
		template = source
	} else if len(s) > 1 {
		template, err = utils.OpenF(s[1])
		if err != nil {
			return
		}
	} else {
		return fmt.Errorf("Can't add template: source is not provided")
	}

	cfg, err := api.GetCfg()
	if err != nil {
		return err
	}
	_, force := p.Flags["force"]
	if err := api.AddTemplate(cfg, s[0], template, force); err != nil {
		return err
	}
	return api.UpdateConfig(cfg)
}

func TemplateRem(p *tap.Parser, s []string) error {
	cfg, err := api.GetCfg()
	if err != nil {
		return err
	}
	if err := api.RemoveTemplate(cfg, s[0]); err != nil {
		return err
	}
	return api.UpdateConfig(cfg)
}
