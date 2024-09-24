package config

import (
	"github.com/spf13/viper"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"os"
	"testing"
)

func TestLoadConfigSuccess(t *testing.T) {
	if err := os.WriteFile("config.json", []byte(`{
	"min_commits": 1,
	"max_commits": 5,
	"days": 30,
	"include_weekends": false,
	"weekend_min_commits": 1,
	"weekend_max_commits": 2,
	"repo_url": "https://github.com/SZabrodskii/git-committer.git",
	"commit_template": "feat: auto commit"
}`), 0644); err != nil {
		t.Fatal(err)
	}
	defer os.Remove("config.json")

	logger, _ := zap.NewProduction()
	cfg, err := NewConfig(logger)

	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, 1, cfg.MinCommits)
	assert.Equal(t, 5, cfg.MaxCommits)
	assert.Equal(t, 30, cfg.Days)
	assert.Equal(t, false, cfg.IncludeWeekends)
	assert.Equal(t, 1, cfg.WeekendMinCommits)
	assert.Equal(t, 2, cfg.WeekendMaxCommits)
	assert.Equal(t, "https://github.com/SZabrodskii/git-committer.git", cfg.RepoURL)
	assert.Equal(t, "feat: auto commit", cfg.CommitTemplate)
}

func TestLoadConfigFileNotFound(t *testing.T) {
	logger, _ := zap.NewProduction()
	_, err := NewConfig(logger)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "config file not found")
}

func TestSetConfigFile(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test_config.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	SetConfigFile(tempFile.Name())

	if viper.ConfigFileUsed() != tempFile.Name() {
		t.Errorf("expected config file to be %s, got %s", tempFile.Name(), viper.ConfigFileUsed())
	}
}
