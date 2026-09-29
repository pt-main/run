package runcli

import (
	"fmt"

	"github.com/pt-main/run"
	"github.com/pt-main/run/run/api"
	localmode "github.com/pt-main/run/run/api/localMode"
	"github.com/pt-main/run/tal"
	tap "github.com/pt-main/tap/go"
)

func NewSys() *tap.Parser {
	sysP := tap.NewParser("sys", "[?GN]Run systems.[?RT]", []string{"help", "-help", "-h"}, tap.DefaultParserConfig())

	sysP.AddCommand("version", func(p *tap.Parser, s []string) error {
		fmt.Println("run v" + run.Version)
		fmt.Println("tal v" + tal.Version)
		fmt.Println("humanmade, by Pt, Apache 2.0 licence")
		return nil
	},
		`[?GN]Show version and license information.[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run sys version[?RT]`,
		nil, nil, false)
	if err := sysP.AddAlias("-v", "version"); err != nil {
		panic("SYSTEM ERROR: CREATING CLI: " + err.Error())
	}

	sysP.AddCommand("localmode", func(p *tap.Parser, s []string) error {
		if len(s) == 0 {
			fmt.Println("localmode:", localmode.IsLocalmode(), "| path:", api.ConfigDirPath())
			return nil
		}
		switch s[0] {
		case "true":
			localmode.Set(true)
		case "false":
			localmode.Set(false)
		default:
			return fmt.Errorf("Invalid argument")
		}
		return nil
	}, `[?GN]Set or show the current working mode (global/local).[?RT]
[?BBK]Usage:[?RT]
  [?BBK]run sys localmode[?RT]           Show current mode and config path
  [?BBK]run sys localmode true[?RT]      Enable local mode (use .run/ in current directory)
  [?BBK]run sys localmode false[?RT]     Disable local mode (use ~/run/)
[?YW]Examples:[?RT]
  [?BBK]run sys localmode[?RT]
  [?BBK]run sys localmode true[?RT]`,
		nil, []string{"mode"}, false)
	if err := sysP.AddAlias("-lm", "localmode"); err != nil {
		panic("SYSTEM ERROR: CREATING CLI: " + err.Error())
	}

	return sysP
}
