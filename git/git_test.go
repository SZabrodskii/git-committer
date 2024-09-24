package git

import (
	"github.com/stretchr/testify/assert"
	"os"
	"os/exec"
	"testing"
)

func TestCloneRepo(t *testing.T) {
	logger, _ := NewTestLogger()
	testRepoURL := "https://github.com/SZabrodskii/git-committer.git"
	repoName := "git-committer"

	if _, err := os.Stat(repoName); !os.IsNotExist(err) {
		os.RemoveAll(repoName)
	}

	repo := &Repository{
		URL:    testRepoURL,
		Name:   repoName,
		Logger: logger,
	}

	err := repo.Clone()
	assert.NoError(t, err)

	if _, err := os.Stat(repoName); os.IsNotExist(err) {
		t.Error("Expected repository to be cloned, but it was not")
	}

	os.RemoveAll(repoName)
}

func TestCreateCommit(t *testing.T) {
	logger, _ := NewTestLogger()
	repoPath := "./test-repo"

	if err := os.Mkdir(repoPath, 0755); err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(repoPath)

	cmd := exec.Command("git", "-C", repoPath, "init")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	filePath := repoPath + "/test.txt"
	if err := os.WriteFile(filePath, []byte("test content"), 0644); err != nil {
		t.Fatal(err)
	}

	cmd = exec.Command("git", "-C", repoPath, "add", ".")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}

	repo := &Repository{
		Name:   repoPath,
		Logger: logger,
	}

	message := "Initial commit"
	err := repo.CreateCommit(message)
	assert.NoError(t, err)

	cmd = exec.Command("git", "-C", repoPath, "log")
	output, err := cmd.Output()
	assert.NoError(t, err)
	assert.Contains(t, string(output), message)
}
