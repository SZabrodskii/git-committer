package git

import (
	"fmt"
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
}

func (g *GitCommitter) createCommitsForDay(day time.Time, commits int) error {
	repoPath := filepath.Base(g.RepoURL)

	for i := 0; i < commits; i++ {
		commitMessage := fmt.Sprintf("%s %d on %s", g.CommitTemplate, i+1, day.Format("2023-01-02"))
		err := CreateCommit(repoPath, commitMessage, nil)
		if err != nil {
			return fmt.Errorf("failed to create commit for day %s: %w", day.String(), err)
		}
	}
	return nil
}

func (g *GitCommitter) generateCommits() error {
	currentDate := time.Now()

	for i := 0; i < g.Days; i++ {
		if g.IncludeWeekends || (!g.IncludeWeekends && !isWeekend(currentDate)) {
			commitsForTheDay := GenerateRandomCommitsPerDay(g.MinCommits, g.MaxCommits)
			err := g.createCommitsForDay(currentDate, commitsForTheDay)
			if err != nil {
				return err
			}
		}
		currentDate = currentDate.AddDate(0, 0, 1)
	}
	return nil
}

func isWeekend(date time.Time) bool {
	weekday := date.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}
