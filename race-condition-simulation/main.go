package main

import (
	"log"
	"sync"
)

type DB interface {
	Store()
}

func Foo(db DB) {
	db.Store()
}

var counter = 0

var cond = sync.NewCond(&sync.Mutex{})
var waitGroupWrite = sync.WaitGroup{}
var waitGroupRead = sync.WaitGroup{}

func incr() {
	cond.L.Lock()

	currentCounter := readCurrentCounter()

	waitGroupRead.Done()
	cond.Wait()

	setCounter(currentCounter + 1)

	cond.L.Unlock()
	waitGroupWrite.Done()
}

func setCounter(newValue int) {
	log.Println("Update counter")
	counter = newValue
}

func readCurrentCounter() int {
	log.Println("Read counter")
	var currentCounter = counter
	return currentCounter
}

func main() {
	waitGroupWrite.Add(2)
	waitGroupRead.Add(2)
	go incr()
	go incr()
	waitGroupRead.Wait()
	cond.Broadcast()
	waitGroupWrite.Wait()
	log.Println(counter)
}
