package main

import (
	"fmt"
	"sync"
	"time"
)

var no int
var mutex *sync.Mutex = &sync.Mutex{}

func f1(wg *sync.WaitGroup) {
	fmt.Println("Starting f1")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Completing f1")
	mutex.Lock()
	no = no * 10
	mutex.Unlock()
	wg.Done()
}

func f2(wg *sync.WaitGroup) {
	fmt.Println("Starting f2")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("Completing f2")
	wg.Done()
}

func f3(wg *sync.WaitGroup) {
	fmt.Println("Starting f3")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Completing f3")
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go f1(wg)
	go f2(wg)
	go f3(wg)
	wg.Wait()
	fmt.Println("Done...!")
}
