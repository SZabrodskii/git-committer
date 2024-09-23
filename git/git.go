package git

import (
	"github.com/SZabrodskii/git-committer/config"

	"fmt"
	"go.uber.org/zap"
	"math/rand"
	"path/filepath"
	"time"
)

func GenerateRandomCommitsPerDay(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func Run(logger *zap.Logger, config *config.Config) {
	logger.Info("Git Auto Committer CLI initialized")

	err := CloneRepo(config.RepoURL, logger)
	if err != nil {
		logger.Fatal("Failed to clone repo", zap.Error(err))
	}

	repoName := filepath.Base(config.RepoURL)

	for i := 0; i < config.Days; i++ {
		currentDay := time.Now().AddDate(0, 0, i)
		weekday := currentDay.Weekday()

		if !config.IncludeWeekends && (weekday == time.Saturday || weekday == time.Sunday) {
			logger.Info("Skipping weekend", zap.String("Day", currentDay.String()))
			continue
		}

		var commitsToday int
		if weekday == time.Saturday || weekday == time.Sunday {
			commitsToday = GenerateRandomCommitsPerDay(config.WeekendMinCommits, config.WeekendMaxCommits)
		} else {
			commitsToday = GenerateRandomCommitsPerDay(config.MinCommits, config.MaxCommits)
		}

		for j := 0; j < commitsToday; j++ {
			message := fmt.Sprintf("%s %d", config.CommitTemplate, j+1)
			err = CreateCommit(repoName, message, logger)
			if err != nil {
				logger.Fatal("Failed to create commit", zap.Error(err))
			}
		}

		logger.Info("Commits created for day", zap.Time("Day", currentDay), zap.Int("Commits", commitsToday))
	}
}
