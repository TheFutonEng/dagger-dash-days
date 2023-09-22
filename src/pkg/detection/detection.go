package detection

import (
	"log"
	"os"
	"strings"
	"github.com/go-git/go-git/v5"
)

func isHelmChart(path string) bool {
	// Check for local path
	if _, err := os.Stat(path); err == nil {
		return checkLocalPathForHelm(path)
	}

	// If not a local path, try as a git repo URL
	return checkGitRepoForHelm(path)
}

func checkLocalPathForHelm(path string) bool {
	chartPath := path + "/Chart.yaml"
	if _, err := os.Stat(chartPath); err == nil {
		return true
	}
	return false
}

func checkGitRepoForHelm(repoURL string) bool {
	// Clone the repository into a temporary directory
	dir, err := os.MkdirTemp("", "tempRepo")
	if err != nil {
		log.Fatalf("Failed to create temp directory: %s", err)
	}
	defer os.RemoveAll(dir)

	_, err = git.PlainClone(dir, false, &git.CloneOptions{
		URL:      repoURL,
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatalf("Failed to clone git repo: %s", err)
	}

	return checkLocalPathForHelm(dir)
}