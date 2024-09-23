package config

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

type Config struct {
	MinCommits        int    `json:"min_commits"`
	MaxCommits        int    `json:"max_commits"`
	Days              int    `json:"days"`
	IncludeWeekends   bool   `json:"include_weekends"`
	WeekendMinCommits int    `json:"weekend_min_commits"`
	WeekendMaxCommits int    `json:"weekend_max_commits"`
	RepoURL           string `json:"repo_url"`
	CommitTemplate    string `json:"commit_template"`
}

func NewConfig(logger *zap.Logger) (*Config, error) {
	viper.SetConfigFile("config.json")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	config := &Config{
		MinCommits:        viper.GetInt("min_commits"),
		MaxCommits:        viper.GetInt("max_commits"),
		Days:              viper.GetInt("days"),
		IncludeWeekends:   viper.GetBool("include_weekends"),
		WeekendMinCommits: viper.GetInt("weekend_min_commits"),
		WeekendMaxCommits: viper.GetInt("weekend_max_commits"),
		RepoURL:           viper.GetString("repo_url"),
		CommitTemplate:    viper.GetString("commit_template"),
	}
	logger.Info("Config loaded successfully")
	return config, nil
}

func SetConfigFile(filename string) {
	viper.SetConfigFile(filename)
}

//функция для ручной загрузки конфигурации из файла в обход вайпера - не удалять

func LoadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
