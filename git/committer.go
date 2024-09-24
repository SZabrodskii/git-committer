package git

import (
	"fmt"
	"math/rand"

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

func (g *GitCommitter) createCommitsForDay(day time.Time, commits int) error {

	for i := 0; i < commits; i++ {
		commitMessage := fmt.Sprintf("%s %d on %s", g.CommitTemplate, i+1, day.Format("2023-01-02"))
		err := g.Repo.CreateCommit(commitMessage)
		if err != nil {
			return fmt.Errorf("failed to create commit for day %s: %w", day.String(), err)
		}
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
