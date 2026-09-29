package tal

import (
	"errors"

	"github.com/pt-main/run/tal/generation"
	"github.com/pt-main/run/tal/lang"
	"github.com/pt-main/run/tal/lua"
)

const Version = "1.1.0"

func Process(changedFiles, args []string, file string) error {
	ls := lua.NewTalLuaState(changedFiles, args)
	processed, err := lang.Process(file)
	if err != nil {
		return errors.New(lang.ErrFmt(err))
	}
	generated, genErr := generation.GenerateCode(processed)
	if genErr != nil {
		return errors.New(lang.ErrFmt(genErr))
	}
	return ls.DoString(generated)
}
