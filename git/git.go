package git

import (
	"github.com/SZabrodskii/git-committer/config"
	"go.uber.org/zap"
	"path/filepath"
)

func Run(logger *zap.Logger, cfg *config.Config) {
	logger.Info("Git Auto Committer CLI initialized")

	repo := &Repository{
		URL:    cfg.RepoURL,
		Name:   filepath.Base(cfg.RepoURL),
		Logger: logger,
	}

	if err := repo.Clone(); err != nil {
		logger.Fatal("Failed to clone repo", zap.Error(err))
	}

	gitCommitter := &GitCommitter{
		MinCommits:        cfg.MinCommits,
		MaxCommits:        cfg.MaxCommits,
		Days:              cfg.Days,
		IncludeWeekends:   cfg.IncludeWeekends,
		WeekendMinCommits: cfg.WeekendMinCommits,
		WeekendMaxCommits: cfg.WeekendMaxCommits,
		RepoURL:           cfg.RepoURL,
		CommitTemplate:    cfg.CommitTemplate,
		Repo:              repo,
	}

	if err := gitCommitter.generateCommits(); err != nil {
		logger.Fatal("Failed to generate commits", zap.Error(err))
	}
}
