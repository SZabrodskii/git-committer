package cli

import (
	"fmt"
	"github.com/SZabrodskii/git-committer/config"
	"github.com/SZabrodskii/git-committer/git"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

const version = "1.0.0"

type CLIRunner struct {
	Logger    *zap.Logger
	Committer *git.GitCommitter
}

func NewCLIRunner(logger *zap.Logger, committer *git.GitCommitter) *CLIRunner {
	return &CLIRunner{Logger: logger, Committer: committer}
}

func (cli *CLIRunner) Run() {
	rootCmd := &cobra.Command{
		Use:   "git-committer",
		Short: "Git Committer CLI",
		Long:  "A CLI tool to automate git commits based on specified rules.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Please provide a valid command. Use --help to see available commands.")
			_ = cmd.Help()
		},
	}

	rootCmd.Version = version
	rootCmd.SetVersionTemplate(fmt.Sprintf("Git Committer version: %s\n", version))

	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate something (commits or config)",
	}

	commitCmd := &cobra.Command{
		Use:   "commit",
		Short: "Generate commits",
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := config.NewConfig()
			if err != nil {
				cli.Logger.Fatal("Failed to load config", zap.Error(err))
			}
			dateStart, _ := cmd.Flags().GetString("dateStart")
			dateEnd, _ := cmd.Flags().GetString("dateEnd")
			minPerDay, _ := cmd.Flags().GetInt("minPerDay")
			maxPerDay, _ := cmd.Flags().GetInt("maxPerDay")

			cli.Logger.Info("Loaded config", zap.Any("config", cfg))
			cli.Logger.Info(fmt.Sprintf("Start date: %s, End date: %s, Min commits per day: %d, Max commits per day: %d", dateStart, dateEnd, minPerDay, maxPerDay))

			cli.Committer.UpdateCommitLimits(minPerDay, maxPerDay)

			if err := cli.Committer.Commit(); err != nil {
				cli.Logger.Error("Failed to generate commits", zap.Error(err))
			}
		},
	}

	commitCmd.Flags().String("dateStart", "", "Start date (YYYY-MM-DD)")
	commitCmd.Flags().String("dateEnd", "", "End date (YYYY-MM-DD)")
	commitCmd.Flags().Int("minPerDay", 3, "Min commits per day")
	commitCmd.Flags().Int("maxPerDay", 7, "Max commits per day")

	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Generate config file",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cli.generateConfig(); err != nil {
				cli.Logger.Error("Failed to generate config file", zap.Error(err))
			}
		},
	}

	generateCmd.AddCommand(commitCmd)
	generateCmd.AddCommand(configCmd)

	rootCmd.AddCommand(generateCmd)

	if err := rootCmd.Execute(); err != nil {
		cli.Logger.Fatal("Command execution failed", zap.Error(err))
	}
}

func (cli *CLIRunner) generateConfig() error {
	_, err := os.Stat("config.json")
	if err == nil {
		return fmt.Errorf("config file already exists")
	}
	err = config.GenerateConfig()
	if err != nil {
		return fmt.Errorf("failed to generate config: %w", err)
	}
	return nil
}
