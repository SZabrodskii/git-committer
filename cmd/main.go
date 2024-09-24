package main

import (
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
			service.GetRandomAnekdot,
			git.NewRepository,
			git.NewGitCommitter,
		),
		fx.Invoke(git.Run),
	)
	app.Run()
}
