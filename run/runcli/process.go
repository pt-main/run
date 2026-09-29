package runcli

import (
	"fmt"

	"github.com/pt-main/run/run/api"
	localmode "github.com/pt-main/run/run/api/localMode"
	tap "github.com/pt-main/tap/go"
)

// Process parses args, honouring a mode flag that overrides the stored
// localmode only for this invocation.
func Process(cli *tap.Parser, args []string) error {
	stored := localmode.IsLocalmode()
	current := stored

	if len(args) > 0 {
		switch args[0] {
		case "--localmode", "--lm":
			current = true
		case "--globalmode", "--gm":
			current = false
		}
	}
	localmode.Set(current)

	defer func() {
		if localmode.IsLocalmode() == current {
			localmode.Set(stored)
		}
	}()

	ok, err := api.CheckConfigDir()
	if err != nil {
		return fmt.Errorf("Can't check installation: %v", err)
	}
	if !ok {
		if err := api.InstallConfigDir(); err != nil {
			return fmt.Errorf("Can't make run dir: %v", err)
		}
	}

	return cli.Parse(args)
}
