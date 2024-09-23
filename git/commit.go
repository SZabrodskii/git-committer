package git

import (
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/exec"
)

type Repository struct {
	URL    string
	Name   string
	Logger *zap.Logger
}

func (r *Repository) Clone() error {
	if _, err := os.Stat(r.Name); os.IsNotExist(err) {
		r.Logger.Info("Cloning repository...", zap.String("RepoURL", r.URL))
		cmd := exec.Command("git", "clone", r.URL)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to clone repo: %w", err)
		}
		r.Logger.Info("Repository cloned successfully")
	} else {
		r.Logger.Info("Repository already exists locally")
	}
	return nil
}

func (r *Repository) CreateCommit(message string) error {
	cmd := exec.Command("git", "-C", r.Name, "add", ".")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to add changes: %w", err)
	}

	cmd = exec.Command("git", "-C", r.Name, "commit", "-m", message)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create commit: %w", err)
	}

	r.Logger.Info("Commit created", zap.String("Message", message))
	return nil
}
