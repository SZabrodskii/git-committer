package git

import (
	"fmt"
	"github.com/SZabrodskii/git-committer/config"
	"github.com/SZabrodskii/git-committer/service"
	"go.uber.org/zap"
	"math/rand"
	"os"
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
}

func NewGitCommitter(cfg *config.Config, repo *Repository) *GitCommitter {
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
	}
}

func (g *GitCommitter) createCommitsForDay(day time.Time, commits int) error {
	for i := 0; i < commits; i++ {
		anekdot, err := service.GetRandomAnekdot()
		if err != nil {
			return fmt.Errorf("failed to get anekdot: %w", err)
		}

		fileName := fmt.Sprintf("anekdot_%d", i+1)
		err = SaveAnekdotToFile(anekdot, g.Repo.Name, fileName)
		if err != nil {
			return fmt.Errorf("failed to save anekdot: %w", err)
		}

		commitMessage := fmt.Sprintf("feat: %s %d on %s", g.CommitTemplate, i+1, day.Format("2006-01-02"))
		err = g.Repo.CreateCommit(commitMessage)
		if err != nil {
			return fmt.Errorf("failed to create commit: %w", err)
		}
	}
	return nil
}

func SaveAnekdotToFile(anekdot, repoName, fileName string) error {
	filePath := filepath.Join(repoName, fileName+".txt")

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(anekdot)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

func (g *GitCommitter) generateCommits() error {
	currentDate := time.Now()

	for i := 0; i < g.Days; i++ {
		var commitsForTheDay int

		if g.IncludeWeekends || (!g.IncludeWeekends && !isWeekend(currentDate)) {
			commitsForTheDay = GenerateRandomCommitsPerDay(g.MinCommits, g.MaxCommits)
		}

		if isWeekend(currentDate) {
			commitsForTheDay = GenerateRandomCommitsPerDay(g.WeekendMinCommits, g.WeekendMaxCommits)
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

func GenerateRandomCommitsPerDay(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func (g *GitCommitter) Commit(logger *zap.Logger) {
	err := g.generateCommits()
	if err != nil {
		logger.Error("Failed to generate commits", zap.Error(err))
		return
	}
	logger.Info("Commits generated successfully")
}
