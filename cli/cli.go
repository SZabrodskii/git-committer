package main

import (
	"fmt"
	"github.com/SZabrodskii/git-committer/config"
	"github.com/SZabrodskii/git-committer/git"
	"github.com/SZabrodskii/git-committer/service"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

type CLIRunner struct {
	Logger *zap.Logger
}

func NewCLIRunner(logger *zap.Logger) *CLIRunner {
	return &CLIRunner{Logger: logger}
}

func (cli *CLIRunner) Run() {
	rootCmd := &cobra.Command{
		Use:   "git-committer",
		Short: "Git Committer CLI",
		Long:  `CLI for generating git commits.`,
		Run: func(cmd *cobra.Command, args []string) {
			cli.Logger.Info("Default behavior executed.")
		},
	}

	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Генерация коммитов",
		Run: func(cmd *cobra.Command, args []string) {
			cli.Logger.Info("Commits generation started...")
			generateCommit(cli.Logger)
		},
	}

	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().String("dateStart", "", "Date of start of generation (YYYY-MM-DD)")
	generateCmd.Flags().String("dateEnd", "", "Date of end of generation (YYYY-MM-DD)")
	generateCmd.Flags().Int("maxPerDay", 10, "Max number of commits per day")
	generateCmd.Flags().Int("minPerDay", 1, "Min number of commits per day")

	viper.BindPFlag("dateStart", generateCmd.Flags().Lookup("dateStart"))
	viper.BindPFlag("dateEnd", generateCmd.Flags().Lookup("dateEnd"))
	viper.BindPFlag("maxPerDay", generateCmd.Flags().Lookup("maxPerDay"))
	viper.BindPFlag("minPerDay", generateCmd.Flags().Lookup("minPerDay"))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func generateCommit(logger *zap.Logger) {
	cfg, err := config.NewConfig(logger)
	if err != nil {
		logger.Fatal("Error loading config", zap.Error(err))
	}

	repo := git.NewRepository(cfg, logger)
	anekdotService := service.NewAnekdotService()

	gitCommitter := git.NewGitCommitter(cfg, repo, anekdotService)

	err = gitCommitter.Commit()
	if err != nil {
		logger.Fatal("Error generating commits", zap.Error(err))
	}

	logger.Info("Commits generated successfully.")
}
