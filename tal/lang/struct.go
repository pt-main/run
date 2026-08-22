package lang

type TalSection struct {
	Cmds map[string]string
	Code string
	Name string
}

func NewTalSection() *TalSection {
	return &TalSection{
		Cmds: make(map[string]string),
		Code: "",
		Name: "",
	}
}

type TalCode struct {
	Global *TalSection
	Main   *TalSection
	Blocks map[string]*TalSection
}

func NewTalCode() *TalCode {
	return &TalCode{
		Global: nil,
		Main:   nil,
		Blocks: make(map[string]*TalSection),
	}
}
