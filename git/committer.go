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

		fileName := fmt.Sprintf("anekdot_%d", i+1)
		filePath := filepath.Join(g.Repo.Name, fileName)

		err = g.AnekdotService.SaveAnekdotToFile(anekdot, g.Repo.Name, fileName)
		if err != nil {
			return fmt.Errorf("failed to save anekdot: %w", err)
		}

		commitMessage := fmt.Sprintf("feat: %s %d on %s", g.CommitTemplate, i+1, day.Format("2006-01-02"))
		err = g.Repo.CreateCommit(filePath, commitMessage)
		if err != nil {
			return fmt.Errorf("failed to create commit: %w", err)
		}
	}
	return nil
}
func (g *GitCommitter) generateCommits() error {
	currentDate := time.Now()

	for i := 0; i < g.Days; i++ {
		var commitsForTheDay int

		if g.IncludeWeekends || (!g.IncludeWeekends && !isWeekend(currentDate)) {
			commitsForTheDay = GetRandomCommitCount(g.MinCommits, g.MaxCommits)
		}

		if isWeekend(currentDate) {
			commitsForTheDay = GetRandomCommitCount(g.WeekendMinCommits, g.WeekendMaxCommits)
		}

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

func (g *GitCommitter) Commit() error {
	g.Repo.Logger.Info("Starting commit generation process")

	err := g.generateCommits()
	if err != nil {
		g.Repo.Logger.Error("Failed to generate commits", zap.Error(err))
		return err
	}

	g.Repo.Logger.Info("Commits generated successfully")
	return nil
}
