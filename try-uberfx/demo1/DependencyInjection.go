package demo1

import (
	"log"

	"go.uber.org/fx"
)

func BuildModule() fx.Option {
	module := fx.Module("demo1",
		fx.Provide(
			//func(d *D, e *E) *B { return NewB(d, e) },
			NewB,
			NewC,
			//
			NewD,
			fx.Annotate(NewE, fx.ParamTags(`group:"G"`)),
			//func() *F { return NewF() },
			NewF,
			NewGa,
			fx.Annotate(NewGb, fx.As(new(G)), fx.ResultTags(`group:"G"`)),
			fx.Annotate(NewGc, fx.As(new(G)), fx.ResultTags(`group:"G"`)),
			fx.Annotate(NewDGa, fx.As(new(G)), fx.ResultTags(`group:"G"`)),
			NewH,
		),
		fx.Invoke(func(b *B, c *C) {
			log.Println("Run 1 ***")
			a1 := NewA(b, c)
			log.Println(a1.ToString())
			log.Println("Run 2 ***")
			a2 := NewA(b, c)
			log.Println(a2.ToString())
		}),
	)
	err := fx.ValidateApp(module)
	if err != nil {
		panic(err)
	}
	return module
}
