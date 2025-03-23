package circulardeps

import (
	"context"
	"log"

	"github.com/firasdarwish/ore"
	"github.com/samber/do/v2"
)

type Foo1 struct {
}
type Bar1 struct {
}

func CircularDepsSamber() {
	injector := do.New()
	do.ProvideTransient(injector, func(inj do.Injector) (*Foo1, error) {
		_ = do.MustInvoke[*Bar1](inj)
		return &Foo1{}, nil
	})
	do.ProvideTransient(injector, func(inj do.Injector) (*Bar1, error) {
		_ = do.MustInvoke[*Foo1](inj)
		return &Bar1{}, nil
	})

	log.Println("Resolving...")
	foo := do.MustInvoke[*Foo1](injector)
	log.Printf("%v", foo)
}

func CircularDepsOre() {
	ore.RegisterLazyFunc(ore.Transient, func(ctx context.Context) (*Foo1, context.Context) {
		_, _ = ore.Get[*Bar1](ctx)
		return &Foo1{}, ctx
	})
	ore.RegisterLazyFunc(ore.Transient, func(ctx context.Context) (*Bar1, context.Context) {
		_, _ = ore.Get[*Foo1](ctx)
		return &Bar1{}, ctx
	})
	log.Println("Resolving...")
	foo, _ := ore.Get[*Foo1](context.Background())
	log.Printf("%v", foo)
}
