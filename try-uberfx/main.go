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
	debugScope := do.ExplainInjector(scope1)
	println(debugScope.String())

	debugService, found := do.ExplainService[*demo1.E](scope1)
	if found {
		println(debugService.String())
	} else {
		println("service not found")
	}

	services := scope1.ListProvidedServices()
	println(services)
}
func Trydo() do.Injector {
	injector := demo1.BuildFastContainer()
	tryDoOnNewScope(injector, "scope1")
	injector.Shutdown()
	return injector
}

func main() {
	Trydo()
}
