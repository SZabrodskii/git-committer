package config

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"os"
)

type Config struct {
	MinCommits        int    `json:"min_commits" validate:"required,gte=0,lte=10"`
	MaxCommits        int    `json:"max_commits" validate:"required,gte=0,lte=10"`
	Days              int    `json:"days" validate:"required,gte=1"`
	IncludeWeekends   bool   `json:"include_weekends"`
	WeekendMinCommits int    `json:"weekend_min_commits" validate:"gte=0"`
	WeekendMaxCommits int    `json:"weekend_max_commits" validate:"gte=0"`
	RepoURL           string `json:"repo_url" validate:"required,url"`
	RepoName          string `json:"repo_name" validate:"required"`
	CommitTemplate    string `json:"commit_template" validate:"required"`
}

func NewConfig() (*Config, error) {
	cfg, err := os.ReadFile("config.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = json.Unmarshal(cfg, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	err = validator.New().Struct(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to validate config: %w", err)
	}

	return &config, nil
}

func GenerateConfig() error {
	config, err := os.Create("config.json")
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer config.Close()

	config.WriteString(`{
		"min_commits": 3,
			"max_commits": 7,
			"days": 30,
			"include_weekends": true,
			"weekend_min_commits": 1,
			"weekend_max_commits": 3,
			"repo_url": "https://github.com/SZabrodskii/git-committer.git",
			"repo_name": "git-committer",
			"commit_template": "Add anekdot"
	}`)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}
