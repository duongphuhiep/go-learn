package main

import (
	"try-uberfx/core"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(core.BuildCoreModule())
	app.Run()
}
