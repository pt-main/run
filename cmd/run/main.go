package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pt-main/run/run/runcli"
)

func main() {
	cli, err := runcli.NewCli()
	if err != nil {
		log.Fatal("SYSTEM ERROR: CREATING CLI:\n", err)
		return
	}
	err = runcli.Process(cli, os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}
}
