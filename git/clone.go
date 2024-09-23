package git

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"path/filepath"
)

func CloneRepo(repoURL string, logger *zap.Logger) error {
	repoName := filepath.Base(repoURL)
	if _, err := os.Stat(repoName); os.IsNotExist(err) {
		logger.Info("Cloning repository...", zap.String("RepoURL", repoURL))
		cmd := exec.Command("git", "clone", repoURL)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to clone repo: %w", err)
		}
		logger.Info("Repository cloned successfully")
	} else {
		logger.Info("Repository already exists locally")
	}
	return nil
}
