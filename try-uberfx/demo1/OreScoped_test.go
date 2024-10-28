package demo1

import (
	"context"
	"testing"

	"github.com/firasdarwish/ore"
)

type T1 struct {
	s1 *S1
}

type S1 struct {
}

func TestScoped(t *testing.T) {
	ore.RegisterLazyFunc(ore.Transient, func(ctx context.Context) (*T1, context.Context) {
		s1, _ := ore.Get[*S1](ctx)
		return &T1{
			s1: s1,
		}, ctx
	})
	ore.RegisterLazyFunc(ore.Scoped, func(ctx context.Context) (*S1, context.Context) {
		return &S1{}, ctx
	})

	t1a, ctx := ore.Get[*T1](context.Background())
	t1b, _ := ore.Get[*T1](ctx)
	if t1a.s1 != t1b.s1 {
		t.FailNow()
	}
}
