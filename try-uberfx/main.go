package main

import (
	"try-uberfx/demo1"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(demo1.BuildModule())
	app.Run()
}
