package demo1

import (
	"github.com/samber/do/v2"
)

func BuildSlowContainer() do.Injector {
	injector := do.New()
	do.ProvideTransient(injector, func(inj do.Injector) (*D, error) {
		resu := do.MustInvokeStruct[D](inj)
		resu.id = generateId("D")
		return resu, nil
	})
	do.ProvideTransient(injector, func(inj do.Injector) (*F, error) {
		resu := do.MustInvokeStruct[F](inj)
		resu.id = generateId("F")
		return resu, nil
	})
	do.Provide(injector, func(inj do.Injector) (H, error) {
		resu := do.MustInvokeStruct[Hr](inj)
		resu.id = generateId("H")
		return resu, nil
	})
	do.Provide(injector, func(inj do.Injector) (*Ga, error) {
		resu := do.MustInvokeStruct[Ga](inj)
		resu.id = generateId("Ga")
		return resu, nil
	})
	do.Provide(injector, func(inj do.Injector) (*Gc, error) {
		resu := do.MustInvokeStruct[Gc](inj)
		resu.id = generateId("Gc")
		return resu, nil
	})
	do.Provide(injector, func(inj do.Injector) (*DGa, error) {
		resu := do.MustInvokeStruct[DGa](inj)
		resu.id = generateId("DGa")
		return resu, nil
	})
	return injector
}

func BuildFastContainer() do.Injector {
	injector := do.New()
	do.ProvideTransient(injector, func(inj do.Injector) (*D, error) {
		return NewD(do.MustInvoke[*F](inj), do.MustInvoke[H](inj)), nil
	})
	do.ProvideTransient(injector, func(inj do.Injector) (*F, error) {
		return NewF(), nil
	})
	do.Provide(injector, func(inj do.Injector) (H, error) {
		return NewHr(), nil
	})
	do.Provide(injector, func(inj do.Injector) (*Ga, error) {
		return NewGa(), nil
	})
	do.Provide(injector, func(inj do.Injector) (*Gc, error) {
		return NewGc(), nil
	})
	do.Provide(injector, func(inj do.Injector) (*DGa, error) {
		return NewDGa(do.MustInvoke[*Ga](inj)), nil
	})
	return injector
}

func NewScopeFast(rootScope do.Injector, scopeId string) *do.Scope {
	scope := rootScope.Scope(scopeId)
	do.ProvideTransient(scope, func(inj do.Injector) (*A, error) {
		return NewA(do.MustInvoke[*B](inj), do.MustInvoke[*C](inj)), nil
	})
	do.ProvideTransient(scope, func(inj do.Injector) (*B, error) {
		return NewB(do.MustInvoke[*D](inj), do.MustInvoke[*E](inj)), nil
	})
	do.Provide(scope, func(inj do.Injector) (*C, error) {
		return NewC(), nil
	})
	do.Provide(scope, func(inj do.Injector) (*E, error) {
		gs := []G{
			do.MustInvoke[*DGa](inj),
			do.MustInvoke[*Gb](inj),
			do.MustInvoke[*Gc](inj),
		}
		return NewE(gs), nil
	})
	do.Provide(scope, func(inj do.Injector) (*Gb, error) {
		return NewGb(), nil
	})
	return scope
}

func NewScopeSlow(rootScope do.Injector, scopeId string) *do.Scope {
	scope := rootScope.Scope(scopeId)
	do.ProvideTransient(scope, func(inj do.Injector) (*A, error) {
		resu := do.MustInvokeStruct[A](inj)
		resu.id = generateId("A")
		return resu, nil
	})
	do.ProvideTransient(scope, func(inj do.Injector) (*B, error) {
		resu := do.MustInvokeStruct[B](inj)
		resu.id = generateId("B")
		return resu, nil
	})
	do.Provide(scope, func(inj do.Injector) (*C, error) {
		resu := do.MustInvokeStruct[C](inj)
		resu.id = generateId("C")
		return resu, nil
	})
	do.Provide(scope, func(inj do.Injector) (*E, error) {
		gs := []G{
			do.MustInvoke[*DGa](inj),
			do.MustInvoke[*Gb](inj),
			do.MustInvoke[*Gc](inj),
		}
		return NewE(gs), nil
	})
	do.Provide(scope, func(inj do.Injector) (*Gb, error) {
		resu := do.MustInvokeStruct[Gb](inj)
		resu.id = generateId("Gb")
		return resu, nil
	})
	return scope
}
