package api

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mattn/go-shellwords"
	localmode "github.com/pt-main/run/run/api/localMode"
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
		panic("Config dir path finding:" + err.Error())
	}
	return filepath.Join(p, dir)
}

func ConfigDirScriptsPath() string {
	return filepath.Join(ConfigDirPath(), "scripts")
}

func ConfigDirBasePath() string {
	return filepath.Join(ConfigDirPath(), "base")
}

func ConfigDirTemplatesPath() string {
	return filepath.Join(ConfigDirPath(), "templates")
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

func NewScriptConfig(name, script, description, ext string, tags []string) *shared.Config {
	conf := shared.NewNilConfig()
	conf.StringV["name"] = name
	conf.StringV["script"] = script
	conf.StringV["description"] = description
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

// TemplateFileName builds a file name for a template of the given extension.
// It is used when a template is stored as a file instead of being inlined
// into the config.
func TemplateFileName(ext string) string {
	name := strings.TrimPrefix(ext, ".")
	if name == "" {
		return "default.templ"
	}
	name = strings.Map(func(r rune) rune {
		switch r {
		case '/', '\\', ':', '*', '?', '"', '<', '>', '|', 0:
			return '_'
		}
		return r
	}, name)
	if name == "" {
		return "default.templ"
	}
	return name + ".templ"
}

// TemplatePath resolves a template `file` value into an absolute path.
// Absolute paths are returned as is, relative ones are resolved against
// the templates dir.
func TemplatePath(file string) string {
	if file == "" {
		return ""
	}
	if filepath.IsAbs(file) {
		return file
	}
	return filepath.Join(ConfigDirTemplatesPath(), file)
}

func ReadTemplateFile(file string) (string, error) {
	path := TemplatePath(file)
	if path == "" {
		return "", fmt.Errorf("Can't read template: file is not provided")
	}
	res, err := utils.OpenF(path)
	if err != nil {
		return "", fmt.Errorf("Can't read template file %q: %v", path, err)
	}
	return res, nil
}

func WriteTemplateFile(file, content string) error {
	path := TemplatePath(file)
	if path == "" {
		return fmt.Errorf("Can't write template: file is not provided")
	}
	if dir := filepath.Dir(path); dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return utils.WriteF(path, content)
}

func RemoveTemplateFile(file string) error {
	path := TemplatePath(file)
	if path == "" {
		return nil
	}
	if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
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
	if err := os.Mkdir(ConfigDirTemplatesPath(), 0755); err != nil {
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
