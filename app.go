package main

import (
	"github.com/Ismael-Romero/cakelet-suite/config"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(config.NewConfig),
		fx.Invoke(config.TestConfig),
		fx.Invoke(Setup),
	)
	app.Run()
}
