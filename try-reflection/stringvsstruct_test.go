package main

import (
	"fmt"
	"testing"
)

type contextValueID struct {
	typeID
	index int
}
type typeID struct {
	pointerTypeName string
	oreKey          string
}

const (
	randomPointerTypeName = "Ad exercitation ea dolore cillum nostrud irure qui labore ex reprehenderit fugiat sunt qui."
	randomKey             = "Minim elit culpa ea velit ullamco."
	randomIndex           = 1
)

var contextString1 = generateString()
var contextString2 = generateString()
var contextStruct1 = generateStruct()
var contextStruct2 = generateStruct()

func generateStruct() contextValueID {
	typeID := typeID{randomPointerTypeName, randomKey}
	return contextValueID{typeID, randomIndex}
}

func generateString() string {
	typeID := fmt.Sprintf("%s:%v", randomPointerTypeName, randomKey)
	return fmt.Sprintln(typeID, randomIndex)
}

func Benchmark_GenerateString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = generateString()
	}
}
func Benchmark_GenerateStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = generateStruct()
	}
}
func Benchmark_CompareString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if contextString1 == contextString2 {
			_ = contextString1
		}
	}
}
func Benchmark_CompareStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if contextStruct1 == contextStruct2 {
			_ = contextStruct1
		}
	}
}
