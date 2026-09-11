package tal

import (
	"errors"

	lccore "github.com/pt-main/lc/engine/core"
	"github.com/pt-main/run/tal/generation"
	"github.com/pt-main/run/tal/lang"
	"github.com/pt-main/run/tal/lua"
)

const Version = "1.1.0"

func Process(changedFiles, args []string, file string) error {
	ls := lua.NewTalLuaState(changedFiles, args)
	var err_ lccore.ErrorInterface
	processed, err_ := lang.Process(file)
	if err_ != nil {
		return errors.New(lang.ErrFmt(err_))
	}
	generated, err := generation.GenerateCode(processed)
	if err != nil {
		return errors.New(lang.ErrFmt(err))
	}
	return ls.DoString(generated)
}
