package runcli

import (
	"fmt"

	run "github.com/pt-main/run/run"
	localmode "github.com/pt-main/run/run/localMode"
	"github.com/pt-main/tap"
)

func Process(cli *tap.Parser, args []string) error {
	lm := localmode.IsLocalmode()
	temp := lm

	if len(args) > 0 {
		if args[0] == "--localmode" || args[0] == "--lm" {
			temp = true
		} else if args[0] == "--globalmode" || args[0] == "--gm" {
			temp = false
		}
	}
	localmode.Set(temp)

	ok, err := run.CheckConfigDir()
	if err != nil {
		return fmt.Errorf("Can't check installation: %v", err)
	}
	if !ok {
		if err := run.InstallConfigDir(); err != nil {
			return fmt.Errorf("Can't make run dir: %v", err)
		}
	}

	err = cli.Parse(args)
	if err != nil {
		return err
	}

	if localmode.IsLocalmode() == temp {
		localmode.Set(lm)
	}
	return nil
}
