package main

import "fmt"

type KeyStringer any

func child(key ...KeyStringer) {
	fmt.Println("key in child (is not the same as parent)")
	for _, k := range key {
		fmt.Printf("%s\n", k)
	}
}

func parent(key ...KeyStringer) {
	fmt.Println("key in parent")
	for _, k := range key {
		fmt.Printf("%s\n", k)
	}
	//forward key to child
	child(key...)
}

func main() {
	parent("a", "b", "c")
}
