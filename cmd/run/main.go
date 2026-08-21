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
	temp := lm

	if len(args) > 1 {
		if args[1] == "--localmode" || args[1] == "--lm" {
			temp = true
		} else if args[1] == "--globalmode" || args[1] == "--gm" {
			temp = false
		}
	}
	localmode.Set(temp)

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

	if localmode.IsLocalmode() == temp {
		localmode.Set(lm)
	}
}
