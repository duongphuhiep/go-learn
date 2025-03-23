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
	log.Println("// ctx 1")
	runOre(context.Background())

	log.Println("// ctx 2 with mock")
	ore.RegisterLazyFunc(ore.Singleton, func(ctx context.Context) (demo1.H, context.Context) {
		return demo1.NewHm(), ctx
	})
	runOre(context.Background())
	log.Println("// ctx 3 with mock")
	runOre(context.Background())
}

func runDo(rootScope do.Injector, scopeId string) {
	scope1 := demo1.NewScopeSlow(rootScope, scopeId)
	a1 := do.MustInvoke[*demo1.A](scope1)
	log.Println(scopeId + ".a1=" + a1.ToString())
	a2 := do.MustInvoke[*demo1.A](scope1)
	log.Println(scopeId + ".a2=" + a2.ToString())
}
func Trydo() {
	injector := demo1.BuildFastContainer()
	log.Println("//scope1")
	runDo(injector, "scope1")

	log.Println("//scope2 with mock")
	do.Override(injector, func(inj do.Injector) (demo1.H, error) {
		return demo1.NewHm(), nil
	})
	runDo(injector, "scope2")
	injector.Shutdown()
}

func main() {
	//circulardeps.CircularDepsOre()
	log.Println("Ore *******")
	Tryore()
	// demo1.ResetCounter()
	// log.Println("Do *******")
	// Trydo()
}
