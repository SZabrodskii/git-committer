package git

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateCommit(message string) error {
	args := m.Called(message)
	return args.Error(0)
}

func TestCreateCommitsForDay(t *testing.T) {
	repo := new(MockRepository)
	commitTemplate := "Test commit"
	day := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	numCommits := 3

	repo.On("CreateCommit", "Test commit 1 on 2023-01-01").Return(nil)
	repo.On("CreateCommit", "Test commit 2 on 2023-01-01").Return(nil)
	repo.On("CreateCommit", "Test commit 3 on 2023-01-01").Return(nil)

	gitCommitter := &GitCommitter{
		Repo:           repo, // problem here
		CommitTemplate: commitTemplate,
	}

	err := gitCommitter.createCommitsForDay(day, numCommits)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestCreateCommitsForDay_Error(t *testing.T) {
	repo := new(MockRepository)
	commitTemplate := "Test commit"
	day := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	numCommits := 2

	repo.On("CreateCommit", "Test commit 1 on 2023-01-01").Return(errors.New("commit error"))

	gitCommitter := &GitCommitter{
		Repo:           repo, // problem here
		CommitTemplate: commitTemplate,
	}

	err := gitCommitter.createCommitsForDay(day, numCommits)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "commit error")
	repo.AssertExpectations(t)
}

func TestGenerateCommits_WithoutWeekends(t *testing.T) {
	repo := new(MockRepository)
	commitTemplate := "Test commit"

	repo.On("CreateCommit", mock.Anything).Return(nil)

	gitCommitter := &GitCommitter{
		MinCommits:      1,
		MaxCommits:      3,
		Days:            5,
		IncludeWeekends: false,
		Repo:            repo, // problem here
		CommitTemplate:  commitTemplate,
	}

	err := gitCommitter.generateCommits()

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}
