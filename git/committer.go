package git

import (
	"fmt"
	"github.com/SZabrodskii/git-committer/config"
	"github.com/SZabrodskii/git-committer/service"
	"go.uber.org/zap"
	"math/rand"
	"path/filepath"
	"time"
)

type GitCommitter struct {
	MinCommits        int
	MaxCommits        int
	Days              int
	IncludeWeekends   bool
	WeekendMinCommits int
	WeekendMaxCommits int
	RepoURL           string
	CommitTemplate    string
	Repo              *Repository
	AnekdotService    *service.AnekdotService
}

func NewGitCommitter(cfg *config.Config, repo *Repository, anekdotService *service.AnekdotService) *GitCommitter {
	return &GitCommitter{
		MinCommits:        cfg.MinCommits,
		MaxCommits:        cfg.MaxCommits,
		Days:              cfg.Days,
		IncludeWeekends:   cfg.IncludeWeekends,
		WeekendMinCommits: cfg.WeekendMinCommits,
		WeekendMaxCommits: cfg.WeekendMaxCommits,
		RepoURL:           cfg.RepoURL,
		CommitTemplate:    cfg.CommitTemplate,
		Repo:              repo,
		AnekdotService:    anekdotService,
	}
}

func (g *GitCommitter) createCommitsForDay(day time.Time, commits int) error {
	for i := 0; i < commits; i++ {
		anekdot, err := g.AnekdotService.GetRandomAnekdot()
		if err != nil {
			return fmt.Errorf("failed to get anekdot: %w", err)
		}

		fileName := fmt.Sprintf("anekdot_%s_%d", day.Format("2006-01-02"), i+1)
		filePath := filepath.Join("example", fileName)

		err = g.AnekdotService.SaveAnekdotToFile(anekdot, filePath)
		if err != nil {
			return fmt.Errorf("failed to save anekdot: %w", err)
		}

		commitMessage := fmt.Sprintf("feat: %s %d on %s", g.CommitTemplate, i+1, day.Format("2006-01-02"))
		err = g.Repo.CreateCommit(filePath, commitMessage, day)
		if err != nil {
			return fmt.Errorf("failed to create commit: %w", err)
		}
	}
	return nil
}
func (g *GitCommitter) generateCommits(dateStart, dateEnd time.Time) error {
	currentDate := dateStart

	for !currentDate.After(dateEnd) {
		var commitsForTheDay int

		if isWeekend(currentDate) {
			commitsForTheDay = GetRandomCommitCount(g.WeekendMinCommits, g.WeekendMaxCommits)
		} else {
			commitsForTheDay = GetRandomCommitCount(g.MinCommits, g.MaxCommits)
		}

		g.Repo.Logger.Info(fmt.Sprintf("Committing %d commits for %s", commitsForTheDay, currentDate.Format("2006-01-02")))

		err := g.createCommitsForDay(currentDate, commitsForTheDay)
		if err != nil {
			return err
		}

		currentDate = currentDate.AddDate(0, 0, 1)
	}
	return nil
}

func isWeekend(date time.Time) bool {
	weekday := date.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

func GetRandomCommitCount(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func (g *GitCommitter) UpdateCommitLimits(minCommits, maxCommits int) {
	g.MinCommits = minCommits
	g.MaxCommits = maxCommits
}

func (g *GitCommitter) Commit() error {
	g.Repo.Logger.Info("Starting commit generation process")

	dateStart := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	dateEnd := time.Date(2025, 7, 23, 0, 0, 0, 0, time.UTC)

	err := g.generateCommits(dateStart, dateEnd)
	if err != nil {
		g.Repo.Logger.Error("Failed to generate commits", zap.Error(err))
		return err
	}

	g.Repo.Logger.Info("Commits generated successfully")
	return nil
}
