package generation

import (
	"fmt"
	"strings"

	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/run/tal/lang"
	"github.com/pt-main/run/tal/shared"
)

func GenerateCode(from *lang.TalCode) (string, core.ErrorInterface) {
	t := &Tasker{
		Functions: make([]string, 0),
	}
	res := ""
	res += "-- ==== RUNTIME CODE ==== --\n"
	res += GetRuntime() + "\n"
	if from.Global != nil {
		res += "\n-- ==== GLOBAL CODE ==== --\n"
		res += from.Global.Code + "\n"
	}
	for _, task := range from.Blocks {
		err := t.GenerateTask(task)
		if err != nil {
			return "", err
		}
	}
	res += "\n-- ==== TASKS DECLARATION ==== --\n"
	res += strings.Join(t.Functions, "\n\n")
	if from.Main != nil {
		res += "\n\n-- ==== MAIN CODE ==== --\n"
		res += from.Main.Code + "\n"
	}
	return res, nil
}

type Tasker struct {
	Functions []string
}

func (t *Tasker) GenerateTask(ts *lang.TalSection) (err core.ErrorInterface) {
	res := ""
	var patterns []string
	for cmd, args := range ts.Cmds {
		switch cmd {
		case "depends":
			parts := strings.Fields(args)
			patterns = append(patterns, parts...)
			continue
		default:
			err = core.Err(shared.GenerationError, "Unknown cmd")
		}
		return core.Wrap(shared.GenerationError, err, "Error in '%v' cmd", cmd)
	}

	patternsLua := "{"
	for i, p := range patterns {
		if i > 0 {
			patternsLua += ", "
		}
		patternsLua += fmt.Sprintf("%q", p)
	}
	patternsLua += "}"

	res += fmt.Sprintf(`tasker.add(%v, 
"%v", function()
%v
end)`, patternsLua, ts.Name, ts.Code)
	t.Functions = append(t.Functions, res)
	return nil
}
