package tests

import (
	"github.com/SZabrodskii/git-committer/git"
	"go.uber.org/zap"
	"testing"
)

func NewTestLogger() (*zap.Logger, error) {
	return zap.NewNop(), nil
}

func TestCreateCommit_ErrorAdding(t *testing.T) {
	logger, err := NewTestLogger()
	if err != nil {
		t.Fatalf("failed to create logger: %v", err)
	}

	err = git.CreateCommit("/invalid/path", "Test commit", logger)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}
