package config

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	MinCommits        int    `json:"min_commits"`
	MaxCommits        int    `json:"max_commits"`
	Days              int    `json:"days"`
	IncludeWeekends   bool   `json:"include_weekends"`
	WeekendMinCommits int    `json:"weekend_min_commits"`
	WeekendMaxCommits int    `json:"weekend_max_commits"`
	RepoURL           string `json:"repo_url"`
	RepoName          string `json:"repo_name"`
	CommitTemplate    string `json:"commit_template"`
}

func NewConfig(logger *zap.Logger) (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Warn("Failed to read config file from disk, trying embedded config", zap.Error(err))

		configFile, err := GetConfigExample()
		if err != nil {
			return nil, fmt.Errorf("failed to read embedded config: %w", err)
		}

		err = viper.ReadConfig(bytes.NewReader(configFile))
		if err != nil {
			return nil, fmt.Errorf("failed to read embedded config: %w", err)
		}
	}

	config := &Config{
		MinCommits:        viper.GetInt("min_commits"),
		MaxCommits:        viper.GetInt("max_commits"),
		Days:              viper.GetInt("days"),
		IncludeWeekends:   viper.GetBool("include_weekends"),
		WeekendMinCommits: viper.GetInt("weekend_min_commits"),
		WeekendMaxCommits: viper.GetInt("weekend_max_commits"),
		RepoURL:           viper.GetString("repo_url"),
		RepoName:          viper.GetString("repo_name"),
		CommitTemplate:    viper.GetString("commit_template"),
	}
	logger.Info("Config loaded successfully", zap.String("File", viper.ConfigFileUsed()))
	return config, nil
}

func SetConfigFile(filename string) {
	viper.SetConfigFile(filename)
}
