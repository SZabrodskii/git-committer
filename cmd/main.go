package main

import (
	"github.com/SZabrodskii/git-committer/config"
	"github.com/SZabrodskii/git-committer/git"
	"github.com/SZabrodskii/git-committer/logger"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(logger.NewLogger),
		fx.Provide(config.NewConfig),
		fx.Invoke(git.Run),
	)
	app.Run()
}
