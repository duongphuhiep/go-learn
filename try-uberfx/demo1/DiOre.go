package demo1

import (
	"context"

	"github.com/firasdarwish/ore"
)

func (this *A) New(ctx context.Context) (*A, context.Context) {
	b, ctx := ore.Get[*B](ctx)
	c, ctx := ore.Get[*C](ctx)
	return NewA(b, c), ctx
}
func (this *B) New(ctx context.Context) (*B, context.Context) {
	d, ctx := ore.Get[*D](ctx)
	e, ctx := ore.Get[*E](ctx)
	return NewB(d, e), ctx
}
func (this *C) New(ctx context.Context) (*C, context.Context) {
	return NewC(), ctx
}
func (this *D) New(ctx context.Context) (*D, context.Context) {
	f, ctx := ore.Get[*F](ctx)
	h, ctx := ore.Get[*H](ctx)
	return NewD(f, h), ctx
}
func (this *E) New(ctx context.Context) (*E, context.Context) {
	gs, ctx := ore.GetList[G](ctx)
	return NewE(gs), ctx
}
func (this *F) New(ctx context.Context) (*F, context.Context) {
	return NewF(), ctx
}
func (this *Ga) New(ctx context.Context) (*Ga, context.Context) {
	return NewGa(), ctx
}
func (this *Gb) New(ctx context.Context) (G, context.Context) {
	return NewGb(), ctx
}
func (this *Gc) New(ctx context.Context) (G, context.Context) {
	return NewGc(), ctx
}
func (this *DGa) New(ctx context.Context) (G, context.Context) {
	ga, ctx := ore.Get[*Ga](ctx)
	return NewDGa(ga), ctx
}
func (this *H) New(ctx context.Context) (*H, context.Context) {
	return NewH(), ctx
}

func RegisterDependenciesToOre_UseFunc() {
	ore.RegisterLazyFunc(ore.Transient, func(ctx context.Context) (*A, context.Context) {
		b, ctx := ore.Get[*B](ctx)
		c, ctx := ore.Get[*C](ctx)
		return NewA(b, c), ctx
	})
	ore.RegisterLazyFunc(ore.Transient, func(ctx context.Context) (*B, context.Context) {
		d, ctx := ore.Get[*D](ctx)
		e, ctx := ore.Get[*E](ctx)
		return NewB(d, e), ctx
	})
	ore.RegisterLazyFunc(ore.Scoped, func(ctx context.Context) (*C, context.Context) {
		return NewC(), ctx
	})
	ore.RegisterLazyFunc(ore.Transient, func(ctx context.Context) (*D, context.Context) {
		f, ctx := ore.Get[*F](ctx)
		h, ctx := ore.Get[*H](ctx)
		return NewD(f, h), ctx
	})
	ore.RegisterLazyFunc(ore.Scoped, func(ctx context.Context) (*E, context.Context) {
		gs, ctx := ore.GetList[G](ctx)
		return NewE(gs), ctx
	})
	ore.RegisterLazyFunc(ore.Transient, func(ctx context.Context) (*F, context.Context) {
		return NewF(), ctx
	})
	ore.RegisterLazyFunc(ore.Singleton, func(ctx context.Context) (*Ga, context.Context) {
		return NewGa(), ctx
	})
	ore.RegisterLazyFunc(ore.Scoped, func(ctx context.Context) (G, context.Context) {
		return NewGb(), ctx
	})
	ore.RegisterLazyFunc(ore.Singleton, func(ctx context.Context) (G, context.Context) {
		return NewGc(), ctx
	})
	ore.RegisterLazyFunc(ore.Singleton, func(ctx context.Context) (G, context.Context) {
		ga, ctx := ore.Get[*Ga](ctx)
		return NewDGa(ga), ctx
	})
	ore.RegisterLazyFunc(ore.Singleton, func(ctx context.Context) (*H, context.Context) {
		return NewH(), ctx
	})
}

func RegisterDependenciesToOre_UseCreator() {
	ore.RegisterLazyCreator(ore.Transient, &A{})
	ore.RegisterLazyCreator(ore.Transient, &B{})
	ore.RegisterLazyCreator(ore.Scoped, &C{})
	ore.RegisterLazyCreator(ore.Transient, &D{})
	ore.RegisterLazyCreator(ore.Scoped, &E{})
	ore.RegisterLazyCreator(ore.Transient, &F{})
	ore.RegisterLazyCreator(ore.Singleton, &Ga{})
	ore.RegisterLazyCreator(ore.Scoped, &Gb{})
	ore.RegisterLazyCreator(ore.Singleton, &Gc{})
	ore.RegisterLazyCreator(ore.Singleton, &DGa{})
	ore.RegisterLazyCreator(ore.Singleton, &H{})
}
