package lang

import (
	"github.com/pt-main/lc/v2/engine/core"
	"github.com/pt-main/lc/v2/parsing/stringParsing"
	"github.com/pt-main/lc/v2/public/errors"
	"github.com/pt-main/lc/v2/tooling/astools"
)

const (
	SysGlobal = "__SYSBLOCK_GLOBALBLOCK"
	SysMain   = "__SYSBLOCK_MAINBLOCK"
)

func Process(code string) (*TalCode, core.ErrorInterface) {
	p := NewParser()
	pn, err := p.Parse(code)
	if err != nil {
		return nil, core.Wrap(errors.ParsingError, err, "%v", p.String())
	}
	return ProcessTalLang(pn)
}

func ProcessTalLang(pn []stringParsing.ParsedNode) (*TalCode, core.ErrorInterface) {
	code := NewTalCode()
	for _, node := range astools.GetChildren(&pn[0]) {
		section := NewTalSection()
		for _, child := range astools.GetChildren(&node) {
			switch child.Switch {
			case "code":
				section.Code += child.Raw
			case "GLOBALBLOCK", "MAINBLOCK":
				section.Name = "__SYSBLOCK_" + child.Switch
			case "BLOCK":
				section.Name = child.Metadata["name"].(string)
			case "COMMAND":
				section.Cmds[child.Metadata["cmd"].(string)] = child.Metadata["args"].(string)
			default:
				return nil, core.Err(errors.ParsingError, "Unknown: %v", child.Switch)
			}
		}
		switch section.Name {
		case SysGlobal:
			code.Global = section
		case SysMain:
			code.Main = section
		default:
			code.Blocks[section.Name] = section
		}
	}
	return code, nil
}
