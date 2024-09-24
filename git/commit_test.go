package git

import (
	"github.com/stretchr/testify/assert"
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

	repo := &Repository{
		Name:   "/invalid/path",
		Logger: logger,
	}

	err = repo.CreateCommit("Test commit")
	assert.Error(t, err, "expected an error, got nil")
}
