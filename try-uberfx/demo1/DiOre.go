package demo1

import (
	"context"

	"github.com/firasdarwish/ore"
)

func (this *A) New(ctx context.Context) *A {
	b, ctx := ore.Get[*B](ctx)
	c, _ := ore.Get[*C](ctx)
	return NewA(b, c)
}
func (this *B) New(ctx context.Context) *B {
	d, ctx := ore.Get[*D](ctx)
	e, _ := ore.Get[*E](ctx)
	return NewB(d, e)
}
func (this *C) New(ctx context.Context) *C {
	return NewC()
}
func (this *D) New(ctx context.Context) *D {
	f, ctx := ore.Get[*F](ctx)
	h, _ := ore.Get[*H](ctx)
	return NewD(f, h)
}
func (this *E) New(ctx context.Context) *E {
	gs, _ := ore.GetList[G](ctx)
	return NewE(gs)
}
func (this *F) New(ctx context.Context) *F {
	return NewF()
}
func (this *Ga) New(ctx context.Context) *Ga {
	return NewGa()
}
func (this *Gb) New(ctx context.Context) G {
	return NewGb()
}
func (this *Gc) New(ctx context.Context) G {
	return NewGc()
}
func (this *DGa) New(ctx context.Context) G {
	ga, _ := ore.Get[*Ga](ctx)
	return NewDGa(ga)
}
func (this *H) New(ctx context.Context) *H {
	return NewH()
}

func RegisterDependenciesToOre_UseFunc() {
	ore.RegisterLazyFunc(ore.Transient, func(ctx context.Context) *A {
		b, ctx := ore.Get[*B](ctx)
		c, _ := ore.Get[*C](ctx)
		return NewA(b, c)
	})
	ore.RegisterLazyFunc(ore.Transient, func(ctx context.Context) *B {
		d, ctx := ore.Get[*D](ctx)
		e, _ := ore.Get[*E](ctx)
		return NewB(d, e)
	})
	ore.RegisterLazyFunc(ore.Scoped, func(ctx context.Context) *C {
		return NewC()
	})
	ore.RegisterLazyFunc(ore.Transient, func(ctx context.Context) *D {
		f, ctx := ore.Get[*F](ctx)
		h, _ := ore.Get[*H](ctx)
		return NewD(f, h)
	})
	ore.RegisterLazyFunc(ore.Scoped, func(ctx context.Context) *E {
		gs, _ := ore.GetList[G](ctx)
		return NewE(gs)
	})
	ore.RegisterLazyFunc(ore.Transient, func(ctx context.Context) *F {
		return NewF()
	})
	ore.RegisterLazyFunc(ore.Singleton, func(ctx context.Context) *Ga {
		return NewGa()
	})
	ore.RegisterLazyFunc(ore.Scoped, func(ctx context.Context) G {
		return NewGb()
	})
	ore.RegisterLazyFunc(ore.Singleton, func(ctx context.Context) G {
		return NewGc()
	})
	ore.RegisterLazyFunc(ore.Singleton, func(ctx context.Context) G {
		ga, _ := ore.Get[*Ga](ctx)
		return NewDGa(ga)
	})
	ore.RegisterLazyFunc(ore.Singleton, func(ctx context.Context) *H {
		return NewH()
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
