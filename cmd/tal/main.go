package main

import (
	"fmt"

	"github.com/pt-main/run/tal/runtime"
)

func main() {
	p := runtime.CreateCli()
	if err := p.Main(); err != nil {
		fmt.Println(err)
	}
}
