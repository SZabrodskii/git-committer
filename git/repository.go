package git

import (
	"fmt"
	"github.com/SZabrodskii/git-committer/config"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"time"
)

type Repository struct {
	URL    string
	Name   string
	Logger *zap.Logger
}

func NewRepository(config *config.Config, logger *zap.Logger) *Repository {
	return &Repository{
		URL:    config.RepoURL,
		Name:   config.RepoName,
		Logger: logger,
	}
}

func (r *Repository) Init() error {
	if _, err := os.Stat(r.Name); os.IsNotExist(err) {
		r.Logger.Info("Initializing new Git repository...", zap.String("RepoName", r.Name))

		cmd := exec.Command("git", "init", r.Name)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to initialize repo: %w", err)
		}

		cmd = exec.Command("git", "-C", r.Name, "remote", "add", "origin", r.URL)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to add remote origin: %w", err)
		}

		r.Logger.Info("Git repository initialized successfully")
	} else {
		r.Logger.Info("Repository directory already exists locally")
	}
	return nil
}

func (r *Repository) CreateCommit(filePath, message string, date time.Time) error {
	r.Logger.Info("Attempting to add file", zap.String("filePath", filePath))

	cmd := exec.Command("git", "-C", r.Name, "add", filePath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to add file %s: %w\nOutput: %s", filePath, err, string(output))
	}

	dateStr := date.Format(time.RFC3339)
	cmd = exec.Command("git", "-C", r.Name, "commit", "--date", dateStr, "-m", message)
	cmd.Env = os.Environ()

	output, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create commit: %w\nOutput: %s", err, string(output))
	}

	r.Logger.Info("Commit created", zap.String("Message", message))
	return nil
}
