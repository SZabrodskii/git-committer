package git

import (
	"go.uber.org/zap"
	"testing"
)

func NewTestLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stdout"}
	return config.Build()
}

func TestCreateCommit_ErrorAdding(t *testing.T) {
	logger, err := NewTestLogger()
	if err != nil {
		t.Fatalf("failed to create logger: %v", err)
	}

	err = CreateCommit("/invalid/path", "Test commit", logger)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}
