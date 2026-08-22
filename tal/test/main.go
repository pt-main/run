package main

import (
	"fmt"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/run/tal/generation"
	"github.com/pt-main/run/tal/lang"
)

func main() {
	fmt.Println("start")
	code := `
-- @build
-- #depends main.go test.go
print("start building...")

-- @
run("build")
`
	proc, err := lang.Process(code)
	fmt.Println(core.GetRealError(err), proc)

	res, err := generation.GenerateCode(proc)
	fmt.Println(core.GetRealError(err), res)
}
