package git

//
//import (
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//type TestRepository struct {
//	Commits []string
//}
//
//func (r *TestRepository) CreateCommit(message string) error {
//	r.Commits = append(r.Commits, message)
//	return nil
//}
//
//func TestGitCommitter_GenerateCommits(t *testing.T) {
//	repo := &TestRepository{}
//	committer := &GitCommitter{
//		MinCommits:        1,
//		MaxCommits:        3,
//		Days:              3,
//		IncludeWeekends:   true,
//		WeekendMinCommits: 2,
//		WeekendMaxCommits: 5,
//		Repo:              repo,
//		CommitTemplate:    "Test commit",
//	}
//
//	err := committer.generateCommits()
//	assert.NoError(t, err)
//
//	assert.True(t, len(repo.Commits) >= 3)
//	assert.True(t, len(repo.Commits) <= 9)
//
//	// Проверка содержимого коммитов
//	for _, commit := range repo.Commits {
//		assert.Contains(t, commit, "Test commit")
//	}
//}
//
//func TestGitCommitter_GenerateCommits_WithoutWeekends(t *testing.T) {
//	repo := &TestRepository{}
//	committer := &GitCommitter{
//		MinCommits:        1,
//		MaxCommits:        3,
//		Days:              3,
//		IncludeWeekends:   false,
//		WeekendMinCommits: 2,
//		WeekendMaxCommits: 5,
//		Repo:              repo,
//		CommitTemplate:    "Test commit",
//	}
//
//	err := committer.generateCommits()
//	assert.NoError(t, err)
//
//	assert.True(t, len(repo.Commits) >= 3)
//	assert.True(t, len(repo.Commits) <= 9)
//
//	for _, commit := range repo.Commits {
//		assert.Contains(t, commit, "Test commit")
//	}
//}
