package git

import (
	"github.com/SZabrodskii/git-committer/config"
	"github.com/SZabrodskii/git-committer/service"
	"go.uber.org/zap"
	"path/filepath"
)

func Run(logger *zap.Logger, cfg *config.Config, anekdotService *service.AnekdotService) {
	logger.Info("Git Auto Committer CLI initialized")

	repoPath := filepath.Base(cfg.RepoURL)

	repo := &Repository{
		URL:    cfg.RepoURL,
		Name:   repoPath,
		Logger: logger,
	}

	if err := repo.Init(); err != nil {
		logger.Fatal("Failed to initialize repo", zap.Error(err))
	}

	gitCommitter := NewGitCommitter(cfg, repo)

	if err := gitCommitter.generateCommits(anekdotService); err != nil {
		logger.Fatal("Failed to generate commits", zap.Error(err))
	}

	logger.Info("Commits generated successfully!")
}
