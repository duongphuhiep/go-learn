package main

import (
	"context"
	"log"
)

type serviceResolver interface {
	resolve() any
}
type serviceDescriptor[T any] struct {
	concrete T
}

func (this *serviceDescriptor[T]) load() T {
	return this.resolve().(T)
}
func (this *serviceDescriptor[T]) resolve() any {
	return this.concrete
}

type numeric interface {
	uint
}

type ICounterGeneric[T numeric] interface {
	Add(number T)
	GetCount() T
}

var _ ICounterGeneric[uint] = (*SimpleCounterUint)(nil)

type SimpleCounterUint struct {
	counter uint
}

func (this *SimpleCounterUint) Add(number uint) {
	this.counter += number
}

func (this *SimpleCounterUint) GetCount() uint {
	return this.counter
}

type someId1 string
type someId2 string

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, someId1("toto"), "hello")
	ctx = context.WithValue(ctx, someId2("toto"), "world")

	log.Printf("%s, %s", ctx.Value(someId1("toto")), ctx.Value(someId2("toto")))
}
