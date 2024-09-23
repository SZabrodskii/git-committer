package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	MinCommitsPerDay int
	MaxCommitsPerDay int
	Days             int
	IncludeWeekends  bool
	WeekendCommits   struct {
		MinCommitsPerDay int
		MaxCommitsPerDay int
	}
	RepoURL        string
	CommitTemplate string
}

func NewConfig(logger *zap.Logger) (*Config, error) {
	viper.SetConfigFile("config.json")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	config := &Config{
		MinCommitsPerDay: viper.GetInt("min_commits_per_day"),
		MaxCommitsPerDay: viper.GetInt("max_commits_per_day"),
		Days:             viper.GetInt("days"),
		IncludeWeekends:  viper.GetBool("include_weekends"),
		RepoURL:          viper.GetString("repo_url"),
		CommitTemplate:   viper.GetString("commit_template"),
	}
	logger.Info("Config loaded successfully")
	return config, nil
}

func SetConfigFile(filename string) {
	viper.SetConfigFile(filename)
}
