package git

import (
	"fmt"
	"go.uber.org/zap"
	"os/exec"
)

func CreateCommit(repoPath string, message string, logger *zap.Logger) error {
	cmd := exec.Command("git", "-C", repoPath, "add", ".")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add changes: %w", err)
	}

	cmd = exec.Command("git", "-C", repoPath, "commit", "-m", message)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create commit: %w", err)
	}

	logger.Info("Commit created", zap.String("Message", message))
	return nil
}
