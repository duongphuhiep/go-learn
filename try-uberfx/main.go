package main

import (
	"context"
	"log"
	"try-uberfx/demo1"

	"github.com/firasdarwish/ore"
	"github.com/samber/do/v2"
)

func tryore() {
	demo1.RegisterDependenciesToOre_UseCreator()

	ctx := context.Background()
	a1, _ := ore.Get[*demo1.A](ctx)
	log.Println(a1.ToString())
	a2, _ := ore.Get[*demo1.A](ctx)
	log.Println(a2.ToString())

	//injector.Shutdown()
}

func trydo() {
	injector := demo1.BuildSlowContainerWithAutoInjection()
	a1 := do.MustInvoke[*demo1.A](injector)
	log.Println(a1.ToString())
	a2 := do.MustInvoke[*demo1.A](injector)
	log.Println(a2.ToString())

	injector.Shutdown()
}

func main() {
	log.Println("Ore *******")
	tryore()
	demo1.ResetCounter()
	log.Println("Do *******")
	trydo()
}
