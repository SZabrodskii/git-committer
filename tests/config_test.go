package tests

import (
	"github.com/SZabrodskii/git-committer/config"
	"github.com/spf13/viper"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"os"
	"testing"
)

func TestLoadConfigSuccess(t *testing.T) {
	if err := os.WriteFile("config.json", []byte(`{
		"min_commits_per_day": 1,
		"max_commits_per_day": 5,
		"days": 30,
		"include_weekends": false,
		"weekend_commits": {
			"min_commits_per_day": 1,
			"max_commits_per_day": 2
		},
		"repo_url": "https://github.com/SZabrodskii/git-committer.git",
		"commit_template": "feat: auto commit"
	}`), 0644); err != nil {
		t.Fatal(err)
	}
	defer os.Remove("config.json")

	logger, _ := zap.NewProduction()
	cfg, err := config.NewConfig(logger)

	assert.NoError(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, 1, cfg.MinCommitsPerDay)
	assert.Equal(t, 5, cfg.MaxCommitsPerDay)
	assert.Equal(t, 30, cfg.Days)
	assert.Equal(t, false, cfg.IncludeWeekends)
	assert.Equal(t, "https://github.com/SZabrodskii/git-committer.git", cfg.RepoURL)
	assert.Equal(t, "feat: auto commit", cfg.CommitTemplate)
}

func TestLoadConfigFileNotFound(t *testing.T) {
	logger, _ := zap.NewProduction()
	_, err := config.NewConfig(logger)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read config file")
}

func TestSetConfigFile(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test_config.json")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	config.SetConfigFile(tempFile.Name())

	if viper.ConfigFileUsed() != tempFile.Name() {
		t.Errorf("expected config file to be %s, got %s", tempFile.Name(), viper.ConfigFileUsed())
	}
}
