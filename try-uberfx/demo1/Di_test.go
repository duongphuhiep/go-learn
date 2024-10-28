package demo1

import (
	"context"
	"strconv"
	"testing"

	"github.com/firasdarwish/ore"
	"github.com/samber/do/v2"
)

func Benchmark_Do_SlowInjector(b *testing.B) {
	slowInjector := BuildSlowContainer()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		scope := NewScopeSlow(slowInjector, strconv.Itoa(n))
		do.MustInvoke[*A](scope)
		scope.Shutdown()
	}
}
func Benchmark_Do_FastInjector(b *testing.B) {
	fastInjector := BuildFastContainer()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		scope := NewScopeFast(fastInjector, strconv.Itoa(n))
		do.MustInvoke[*A](scope)
		scope.Shutdown()
	}
}

func Benchmark_Ore(b *testing.B) {
	RegisterDependenciesToOre_UseCreator()
	ctx := context.Background()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ore.Get[*A](ctx)
	}
}
