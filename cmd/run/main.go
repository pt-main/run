package main

import (
	"fmt"
	"os"

	"github.com/pt-main/run/cmd/runcli"
)

func main() {
	cli := runcli.NewCli()
	err := runcli.Process(cli, os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}
}
