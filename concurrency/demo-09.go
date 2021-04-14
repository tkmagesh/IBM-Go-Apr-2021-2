package main

import (
	"fmt"
	"time"
)

func fibonacci(ch chan int, quit <-chan time.Time) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
			time.Sleep(2 * time.Second)
		case <-quit:
			fmt.Println("Quitting")
			return
		}
	}
}

func main() {

	ch := make(chan int)
	//quit := make(chan bool)
	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()
	go fibonacci(ch, time.After(15*time.Second))
	input := ""
	fmt.Println("Hit ENTER to stop")
	fmt.Scanln(&input)

	fmt.Println("End of main")
}
