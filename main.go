package main

import (
	"github.com/SZabrodskii/git-committer/cli"
	"github.com/SZabrodskii/git-committer/config"
	"github.com/SZabrodskii/git-committer/git"
	"github.com/SZabrodskii/git-committer/logger"
	"github.com/SZabrodskii/git-committer/service"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			logger.NewLogger,
			config.NewConfig,
			service.NewAnekdotService,
			git.NewRepository,
			git.NewGitCommitter,
			cli.NewCLIRunner,
		),
		fx.Invoke((*cli.CLIRunner).Run),
	)
	app.Run()
}
