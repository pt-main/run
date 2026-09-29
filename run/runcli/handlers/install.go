package runlib

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/mattn/go-shellwords"
	"github.com/pt-main/run/run/api"
	"github.com/pt-main/run/tal"
	tap "github.com/pt-main/tap/go"
)

func downloadScript(url string) (content string, fileName string, err error) {
	if strings.Contains(url, "github.com/") {
		url, fileName, err = parseGitHubURL(url)
		if err != nil {
			return "", "", err
		}
	} else {
		fileName = filepath.Base(url)
		if idx := strings.Index(fileName, "?"); idx != -1 {
			fileName = fileName[:idx]
		}
	}

	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("HTTP error: %s", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}
	return string(data), fileName, nil
}

func parseGitHubURL(rawURL string) (rawContentURL string, fileName string, err error) {
	url := strings.TrimPrefix(strings.TrimPrefix(rawURL, "https://"), "http://")

	if !strings.HasPrefix(url, "github.com/") {
		return "", "", fmt.Errorf("not a GitHub URL")
	}
	url = strings.TrimPrefix(url, "github.com/")

	if repo, rest, found := strings.Cut(url, "@"); found {
		ref, path, found := strings.Cut(rest, "/")
		if !found {
			return "", "", fmt.Errorf("missing path after ref")
		}
		return rawFileURL(repo, ref, path), filepath.Base(path), nil
	}

	repoPart, rest, found := strings.Cut(url, "/blob/")
	if !found {
		return "", "", fmt.Errorf("invalid GitHub URL: missing '@' or '/blob/'")
	}
	branch, path, found := strings.Cut(rest, "/")
	if !found {
		return "", "", fmt.Errorf("invalid blob URL: missing branch/path")
	}
	return rawFileURL(repoPart, branch, path), filepath.Base(path), nil
}

func rawFileURL(repo, ref, path string) string {
	return fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", repo, ref, path)
}

func InstallHandler(p *tap.Parser, s []string) error {
	if len(s) < 1 {
		return fmt.Errorf("need at least URL")
	}

	url := s[0]
	scriptName := ""
	docs := ""

	if len(s) > 1 {
		scriptName = s[1]
	}
	if len(s) > 2 {
		docs = s[2]
	}

	content, rawName, err := downloadScript(url)
	if err != nil {
		return err
	}

	if scriptName == "" {
		scriptName = strings.TrimSuffix(rawName, filepath.Ext(rawName))
	}

	if rawName == "run.task.lua" {
		return runInstaller(p, s[1:], content)
	}

	_, force := p.Flags["force"]

	cfg, err := api.GetCfg()
	if err != nil {
		return err
	}

	return api.Upconf(cfg, api.AddScript(cfg, content, rawName, scriptName, docs, force))
}

func runInstaller(p *tap.Parser, args []string, content string) error {
	if scriptArgs, ok := p.Flags["args"]; ok {
		parsed, err := shellwords.Parse(scriptArgs)
		if err != nil {
			return fmt.Errorf("Parsing args: %v", err)
		}
		args = parsed
	}
	return tal.Process([]string{}, args, content)
}
