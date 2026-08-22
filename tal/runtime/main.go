package runtime

import (
	"bufio"
	"fmt"
	"os"

	"github.com/iancoleman/orderedmap"
	lccore "github.com/pt-main/lc/engine/core"
	"github.com/pt-main/run/tal/core"
	"github.com/pt-main/run/tal/generation"
	"github.com/pt-main/run/tal/lang"
	"github.com/pt-main/run/tal/lua"
	"github.com/pt-main/run/tal/shared"
	"github.com/pt-main/tap"
)

func CreateCli() *tap.Parser {
	p := tap.NewParser("tal", ``, []string{"-h", "help"}, tap.DefaultParserConfig())
	p.AddCommand("update", UpdateHandler, ``, nil, nil, false)
	p.AddCommand("run", RunHandler, ``, []string{"tal-lua-file"}, nil, true)
	return p
}

func UpdateHandler(p *tap.Parser, s []string) error {
	return Update()
}

func RunHandler(p *tap.Parser, s []string) error {
	ch, err := GetChanges()
	if err != nil {
		return err
	}
	if err := Update(); err != nil {
		return err
	}
	args := []string{}
	skippedName := false
	for _, arg := range p.RawArgs[1:] {
		if arg == s[0] && !skippedName {
			skippedName = true
		} else {
			args = append(args, arg)
		}
	}
	ls := lua.NewTalLuaState(ch, args)
	file, err := OpenF(s[0])
	if err != nil {
		return err
	}
	processed, err := lang.Process(file)
	if err != nil {
		return fmt.Errorf(lccore.GetRealError(err))
	}
	generated, err := generation.GenerateCode(processed)
	return ls.DoString(generated)
}

func GetSavedFile() (*orderedmap.OrderedMap, error) {
	file, err := open(shared.TalFile)
	if err != nil {
		return nil, err
	}
	return core.PackCoreAsState(file)
}

func GetChanges() ([]string, error) {
	w, err := GetSavedFile()
	if err != nil {
		return nil, err
	}
	return core.Changes(w, ".")
}

func Update() error {
	st, err := core.SaveState(".")
	if err != nil {
		return err
	}
	file, err := core.StateAsPackCore(st)
	if err != nil {
		return err
	}
	return write(shared.TalFile, file)
}

func OpenF(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("Open: %v", err)
	}
	return string(data), nil
}

func open(file string) ([]byte, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("Open: %v", err)
	}
	return data, nil
}

func write(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Write: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.Write(data)
	if err != nil {
		return fmt.Errorf("Write: %v", err)
	}
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("Write: %v", err)
	}
	return nil
}
