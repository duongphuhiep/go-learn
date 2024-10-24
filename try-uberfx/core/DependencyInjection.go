package core

import (
	"go.uber.org/fx"
)

func BuildModule() fx.Option {
	module := fx.Module("core",
		fx.Provide(
			fx.Annotate(NewThorWalletRepository, fx.As(new(IWalletRepository)), fx.ResultTags(`name:"thor"`)),
			fx.Annotate(NewLokiWalletRepository, fx.As(new(IWalletRepository)), fx.ResultTags(`name:"loki"`)),
			fx.Annotate(
				NewWalletServiceImpl,
				fx.As(new(IWalletService)),
				fx.ParamTags(`name:"thor"`),
			),
		),
		fx.Decorate(
			fx.Annotate(NewCachedWalletService, fx.As(new(IWalletService))),
		),
		fx.Invoke(Run),
	)
	err := fx.ValidateApp(module)
	if err != nil {
		panic(err)
	}
	return module
}
