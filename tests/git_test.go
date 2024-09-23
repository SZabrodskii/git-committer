package tests

import (
	"github.com/SZabrodskii/git-committer/git"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"os"
	"os/exec"
	"testing"
)

func TestCloneRepo(t *testing.T) {
	logger, _ := zap.NewProduction()
	testRepoURL := "https://github.com/SZabrodskii/git-committer.git"

	if _, err := os.Stat("git-committer"); !os.IsNotExist(err) {
		os.RemoveAll("git-committer")
	}

	err := git.CloneRepo(testRepoURL, logger)
	assert.NoError(t, err)

	if _, err := os.Stat("git-committer"); os.IsNotExist(err) {
		t.Error("Expected repository to be cloned, but it was not")
	}

	os.RemoveAll("git-committer")
}

func TestCreateCommit(t *testing.T) {
	logger, _ := zap.NewProduction()
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

	message := "Initial commit"
	err := git.CreateCommit(repoPath, message, logger)
	assert.NoError(t, err)

	cmd = exec.Command("git", "-C", repoPath, "log")
	output, err := cmd.Output()
	assert.NoError(t, err)
	assert.Contains(t, string(output), message)
}
