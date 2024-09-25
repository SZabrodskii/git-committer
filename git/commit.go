package git

import (
	"fmt"
	"github.com/SZabrodskii/git-committer/config"
	"go.uber.org/zap"
	"os"
	"os/exec"
)

type Repository struct {
	URL    string
	Name   string
	Logger *zap.Logger
}

func NewRepository(config *config.Config, logger *zap.Logger) *Repository {
	repoName := "git-committer"
	return &Repository{
		URL:    config.RepoURL,
		Name:   repoName,
		Logger: logger,
	}
}

func (r *Repository) Init() error {
	if _, err := os.Stat(r.Name); os.IsNotExist(err) {
		r.Logger.Info("Initializing new Git repository...", zap.String("RepoName", r.Name))

		err := os.Mkdir(r.Name, 0755)
		if err != nil {
			return fmt.Errorf("failed to create repository directory: %w", err)
		}

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
