package git

import (
	"go.uber.org/zap"
)

func Run(logger *zap.Logger, gitCommitter *GitCommitter) {
	logger.Info("Git Auto Committer CLI initialized")

	if err := gitCommitter.Repo.Init(); err != nil {
		logger.Fatal("Failed to initialize repo", zap.Error(err))
	}

	if err := gitCommitter.generateCommits(); err != nil {
		logger.Fatal("Failed to generate commits", zap.Error(err))
	}

	logger.Info("Commits generated successfully!")
}
