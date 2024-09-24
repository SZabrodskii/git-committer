package git

import (
	"github.com/SZabrodskii/git-committer/config"
	"go.uber.org/zap"
	"path/filepath"
)

func Run(logger *zap.Logger, cfg *config.Config) {
	logger.Info("Git Auto Committer CLI initialized")

	repoPath := filepath.Base(cfg.RepoURL)

	repo := &Repository{
		URL:    cfg.RepoURL,
		Name:   repoPath,
		Logger: logger,
	}

	if err := repo.Clone(); err != nil {
		logger.Fatal("Failed to clone repo", zap.Error(err))
	}

	gitCommitter := NewGitCommitter(cfg, repo)

	gitCommitter.Commit(logger)

	logger.Info("Commits generated successfully!")
}
