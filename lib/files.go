package runlib

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mattn/go-shellwords"
	localmode "github.com/pt-main/run/lib/localMode"
	"github.com/pt-main/tycl/generation"
	"github.com/pt-main/tycl/shared"
	"github.com/pt-main/tycl/utils"
)

func ConfigDirPath() string {
	p, err := os.UserHomeDir()
	dir := "run"
	if localmode.IsLocalmode() {
		p, err = os.Getwd()
		dir = ".run"
	}
	if err != nil {
		panic(err)
	}
	return filepath.Join(p, dir)
}

func ConfigDirScriptsPath() string {
	return filepath.Join(ConfigDirPath(), "scripts")
}

func ConfigDirBasePath() string {
	return filepath.Join(ConfigDirPath(), "base")
}

func ConfigDirConfigPath() string {
	return filepath.Join(ConfigDirPath(), "config.tycl")
}

func ProcessShell(cmdStr string) ([]string, error) {
	args, err := shellwords.Parse(cmdStr)
	if err != nil {
		return nil, fmt.Errorf("Parse shell args: %v", err)
	}
	if len(args) == 0 {
		return nil, nil
	}
	return args, nil
}

func CheckConfigDir() (bool, error) {
	info, err := os.Stat(ConfigDirPath())
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

func NewScriptConfig(name, script, description, source, ext string, tags []string) *shared.Config {
	conf := shared.NewNilConfig()
	conf.StringV["name"] = name
	conf.StringV["script"] = script
	conf.StringV["description"] = description
	conf.StringV["source"] = source
	conf.StringV["ext"] = ext
	if tags == nil {
		tags = []string{}
	}
	conf.StringArrV["tags"] = tags
	return conf
}

func FormatConfig(config *shared.Config) (res string, err error) {
	if _, ok := config.InnerArrV["scripts"]; !ok {
		config.InnerArrV["scripts"] = make([]*shared.Config, 0)
	}
	res, err = generation.Tycl(config)
	return
}

func NewRunScript(name, content string) error {
	return utils.WriteF(filepath.Join(ConfigDirScriptsPath(), name+".lua"), content)
}

func NewScript(name, content string) error {
	return utils.WriteF(filepath.Join(ConfigDirBasePath(), name), content)
}

func UpdateConfig(config *shared.Config) error {
	conf, err := FormatConfig(config)
	if err != nil {
		return err
	}
	if err := utils.WriteF(ConfigDirConfigPath(), conf); err != nil {
		return err
	}
	return nil
}

func InstallConfigDir() error {
	if err := os.Mkdir(ConfigDirPath(), 0755); err != nil {
		return err
	}
	if err := os.Mkdir(ConfigDirScriptsPath(), 0755); err != nil {
		return err
	}
	if err := os.Mkdir(ConfigDirBasePath(), 0755); err != nil {
		return err
	}
	conf, err := StdLib()
	if err != nil {
		return err
	}
	if err := utils.WriteF(ConfigDirConfigPath(), conf); err != nil {
		return err
	}
	return nil
}
