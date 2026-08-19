package main

import (
	"fmt"
	"os"

	run "github.com/pt-main/run/lib"
	localmode "github.com/pt-main/run/lib/localMode"
)

func main() {
	args := os.Args
	lm := localmode.IsLocalmode()
	has := false
	if len(args) > 1 {
		if args[1] == "--localmode" || args[1] == "--lm" {
			has = true
		}
	}
	if has {
		localmode.Set(true)
	}

	ok, err := run.CheckConfigDir()
	if err != nil {
		fmt.Println("Can't check installation:", err)
		return
	}
	if !ok {
		if err := run.InstallConfigDir(); err != nil {
			fmt.Println("Can't make run dir:", err)
			return
		}
	}

	err = NewCli().Main()
	if err != nil {
		fmt.Println(err)
	}

	if has {
		localmode.Set(lm)
	}
}
