package main

import (
	"fmt"
	"reflect"
)

func typeIdentifierWithPointerReflection[T any](key string) string {
	return fmt.Sprintf("%s:%v", reflect.TypeFor[*T]().String(), key)
}

func typeIdentifierWithReflection[T any](key string) string {
	return fmt.Sprintf("%s:%v", reflect.TypeFor[T]().String(), key)
}

func typeIdentifierWithPercentT[T any](key string) string {
	var mockValue T
	return fmt.Sprintf("%T:%v", mockValue, key)
}

func typeIdentifierWithPointerPercentT[T any](key string) string {
	var mockValue *T
	return fmt.Sprintf("%T:%v", mockValue, key)
}

func typeIdentifierWithPointerPercentC[T any](key string) string {
	var mockValue *T
	return fmt.Sprintf("%c:%v", mockValue, key)
}
