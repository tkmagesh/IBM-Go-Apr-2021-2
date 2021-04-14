package main

import "fmt"

func fibonacci(ch chan int) {
	count := 10
	x, y := 0, 1
	for i := 0; i < count; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	count := 10
	ch := make(chan int, count)
	go fibonacci(ch)
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("End of main")
}
