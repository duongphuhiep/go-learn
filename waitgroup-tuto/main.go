package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	MainWaitGroup()
}

func MainWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		doMission()
	}()

	go func() {
		defer wg.Done()
		doMission()
	}()

	wg.Wait()
	log.Print("terminated")
}

func doMission() {
	log.Print("Start mission")
	time.Sleep(1 * time.Second)
	log.Print("Mission completed")
}
