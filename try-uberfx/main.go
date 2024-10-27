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

func tryDoOnNewScope(rootScope do.Injector, scopeId string) {
	scope1 := demo1.NewScopeSlow(rootScope, scopeId)
	a1 := do.MustInvoke[*demo1.A](scope1)
	log.Println(scopeId + ".a1=" + a1.ToString())
	a2 := do.MustInvoke[*demo1.A](scope1)
	log.Println(scopeId + ".a2=" + a2.ToString())
}
func Trydo() {
	injector := demo1.BuildFastContainer()
	tryDoOnNewScope(injector, "scope1")
	tryDoOnNewScope(injector, "scope2")
	injector.Shutdown()
}

func main() {
	log.Println("Ore *******")
	Tryore()
	demo1.ResetCounter()
	log.Println("Do *******")
	Trydo()
}
