package localmode

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pt-main/tycl/utils"
)

func ConfigLocalmodePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(home, "run.localmode")
}

func CheckConfigLocalmode() bool {
	if _, err := os.Stat(ConfigLocalmodePath()); err != nil {
		return false
	}
	return true
}

func Install() {
	if !CheckConfigLocalmode() {
		utils.WriteF(ConfigLocalmodePath(), "false")
	}
}

func IsLocalmode() bool {
	Install()
	file, err := utils.OpenF(ConfigLocalmodePath())
	if err != nil {
		return false
	}
	return strings.TrimSpace(file) == "true"
}

func Set(local bool) {
	Install()
	content := "false"
	if local {
		content = "true"
	}
	if err := utils.WriteF(ConfigLocalmodePath(), content); err != nil {
		panic(err)
	}
}
