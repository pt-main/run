package lang

import (
	"github.com/pt-main/lc/engine/core"
	"github.com/pt-main/lc/parsing/stringParsing"
	"github.com/pt-main/lc/public/errors"
	"github.com/pt-main/lc/tooling/astools"
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
	c := NewTalCode()
	for _, node := range astools.GetChildren(&pn[0]) {
		chs := astools.GetChildren(&node)
		sec := NewTalSection()
		for _, ch := range chs {
			switch ch.Switch {
			case "code":
				sec.Code += ch.Raw
			case "GLOBALBLOCK", "MAINBLOCK":
				sec.Name = "__SYSBLOCK_" + ch.Switch
			case "BLOCK":
				sec.Name = ch.Metadata["name"].(string)
			case "COMMAND":
				sec.Cmds[ch.Metadata["cmd"].(string)] = ch.Metadata["args"].(string)
			default:
				return nil, core.Err(errors.ParsingError, "Unknown: %v", ch.Switch)
			}
		}
		switch sec.Name {
		case SysGlobal:
			c.Global = sec
		case SysMain:
			c.Main = sec
		default:
			c.Blocks[sec.Name] = sec
		}
	}
	return c, nil
}
