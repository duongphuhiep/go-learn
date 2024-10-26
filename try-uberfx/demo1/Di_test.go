package demo1

import (
	"context"
	"testing"

	"github.com/firasdarwish/ore"
	"github.com/samber/do/v2"
)

func Benchmark_Do_SlowInjector(b *testing.B) {
	slowInjector := BuildSlowContainerWithAutoInjection()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		do.MustInvoke[*A](slowInjector)
	}
}
func Benchmark_Do_FastInjector(b *testing.B) {
	fastInjector := BuildFastContainer()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		do.MustInvoke[*A](fastInjector)
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
