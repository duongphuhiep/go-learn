package main

import (
	"context"
	"log"
	"try-uberfx/demo1"

	"github.com/firasdarwish/ore"
	"github.com/samber/do/v2"
)

func runOre(ctx context.Context) {
	a1, ctx := ore.Get[*demo1.A](ctx)
	log.Println(a1.ToString())
	a2, _ := ore.Get[*demo1.A](ctx)
	log.Println(a2.ToString())
}

func Tryore() {
	demo1.RegisterDependenciesToOre_UseFunc()
	log.Println("// root ctx")
	runOre(context.Background())
	// log.Println("// other ctx")
	// runOre(context.Background())
}

func Trydo() {
	injector := demo1.BuildSlowContainerWithAutoInjection()
	a1 := do.MustInvoke[*demo1.A](injector)
	log.Println(a1.ToString())
	a2 := do.MustInvoke[*demo1.A](injector)
	log.Println(a2.ToString())

	injector.Shutdown()
}

func main() {
	log.Println("Ore *******")
	Tryore()
	demo1.ResetCounter()
	log.Println("Do *******")
	Trydo()
}
