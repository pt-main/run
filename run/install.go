package runlib

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/mattn/go-shellwords"
	"github.com/pt-main/run/tal"
	"github.com/pt-main/tap"
)

func downloadScript(url string) (content string, fileName string, err error) {
	if strings.Contains(url, "github.com/") {
		rawURL, fname, err := parseGitHubURL(url)
		if err != nil {
			return "", "", err
		}
		url = rawURL
		fileName = fname
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

	u := rawURL
	if strings.HasPrefix(u, "https://") {
		u = strings.TrimPrefix(u, "https://")
	} else if strings.HasPrefix(u, "http://") {
		u = strings.TrimPrefix(u, "http://")
	}

	if !strings.HasPrefix(u, "github.com/") {
		return "", "", fmt.Errorf("not a GitHub URL")
	}
	u = strings.TrimPrefix(u, "github.com/")

	parts := strings.SplitN(u, "@", 2)
	if len(parts) == 2 {
		repo := parts[0]
		rest := parts[1]
		slashIdx := strings.Index(rest, "/")
		if slashIdx == -1 {
			return "", "", fmt.Errorf("missing path after ref")
		}
		ref := rest[:slashIdx]
		path := rest[slashIdx+1:]
		rawContentURL = fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", repo, ref, path)
		fileName = filepath.Base(path)
		return rawContentURL, fileName, nil
	}

	idx := strings.Index(u, "/blob/")
	if idx == -1 {
		return "", "", fmt.Errorf("invalid GitHub URL: missing '@' or '/blob/'")
	}
	repoPart := u[:idx]
	rest := u[idx+len("/blob/"):]
	slashIdx := strings.Index(rest, "/")
	if slashIdx == -1 {
		return "", "", fmt.Errorf("invalid blob URL: missing branch/path")
	}
	branch := rest[:slashIdx]
	path := rest[slashIdx+1:]
	rawContentURL = fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/%s", repoPart, branch, path)
	fileName = filepath.Base(path)
	return rawContentURL, fileName, nil
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
		ext := filepath.Ext(rawName)
		scriptName = strings.TrimSuffix(rawName, ext)
	}

	_, force := p.Flags["force"]

	// running tal isntallation file
	if rawName == "run.task.lua" {
		args := s[1:]
		_args, hasArgs := p.Flags["args"]
		if hasArgs {
			args, err = shellwords.Parse(_args)
			if err != nil {
				return fmt.Errorf("Parsing args: %v", err)
			}
		}

		err := tal.Process([]string{}, args, content)
		if err != nil {
			return err
		}
		return nil
	}

	// adding script
	return AddScript(content, rawName, scriptName, docs, force)
}
