package api

import "github.com/pt-main/tycl/shared"

// StdLib returns the config created on first launch: a single demo script and
// an empty template list.
func StdLib() (string, error) {
	config := shared.NewNilConfig()
	NewRunScript("test", `print("test script"); print(script_path("test.py")); print(get_args()[1])`)
	config.InnerArrV["scripts"] = append(config.InnerArrV["scripts"], NewScriptConfig(
		"test", "test", "[?BBK]Simple script for functions test[?RT]", "", []string{"__test"},
	))
	config.InnerArrV["templates"] = []*shared.Config{}
	conf, err := FormatConfig(config)
	if err != nil {
		return "", err
	}
	return conf, nil
}
