package lang

// TalSection is a single block of a task file: its name, its `-- #cmd` lines
// and its Lua body.
type TalSection struct {
	Cmds map[string]string
	Code string
	Name string
}

func NewTalSection() *TalSection {
	return &TalSection{Cmds: make(map[string]string)}
}

// TalCode is a parsed task file: its global and main blocks plus every named
// task.
type TalCode struct {
	Global *TalSection
	Main   *TalSection
	Blocks map[string]*TalSection
}

func NewTalCode() *TalCode {
	return &TalCode{Blocks: make(map[string]*TalSection)}
}
