package generation

import (
	"fmt"
	"strings"

	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/run/tal/lang"
	"github.com/pt-main/run/tal/shared"
)

func GenerateCode(from *lang.TalCode) (string, core.ErrorInterface) {
	t := &Tasker{Functions: make([]string, 0, len(from.Blocks))}

	var res strings.Builder
	res.WriteString("-- ==== RUNTIME CODE ==== --\n")
	res.WriteString(GetRuntime())
	res.WriteString("\n")

	if from.Global != nil {
		res.WriteString("\n-- ==== GLOBAL CODE ==== --\n")
		res.WriteString(from.Global.Code)
		res.WriteString("\n")
	}

	for _, block := range from.Blocks {
		if err := t.GenerateTask(block); err != nil {
			return "", err
		}
	}

	res.WriteString("\n-- ==== TASKS DECLARATION ==== --\n")
	res.WriteString(strings.Join(t.Functions, "\n\n"))

	if from.Main != nil {
		res.WriteString("\n\n-- ==== MAIN CODE ==== --\n")
		res.WriteString(from.Main.Code)
		res.WriteString("\n")
	}

	return res.String(), nil
}

// Tasker collects the generated task declarations.
type Tasker struct {
	Functions []string
}

func (t *Tasker) GenerateTask(ts *lang.TalSection) (err core.ErrorInterface) {
	var patterns []string
	for cmd, args := range ts.Cmds {
		switch cmd {
		case "depends":
			patterns = append(patterns, strings.Fields(args)...)
		default:
			return core.Wrap(shared.GenerationError, core.Err(shared.GenerationError, "Unknown cmd"),
				"Error in '%v' cmd", cmd)
		}
	}

	quoted := make([]string, 0, len(patterns))
	for _, p := range patterns {
		quoted = append(quoted, fmt.Sprintf("%q", p))
	}

	t.Functions = append(t.Functions, fmt.Sprintf(`tasker.add({%v},
"%v", function()
%v
end)`, strings.Join(quoted, ", "), ts.Name, ts.Code))
	return nil
}
