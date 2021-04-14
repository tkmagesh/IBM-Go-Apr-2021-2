package main

import (
	"fmt"
	"time"
)

func writeData(ch chan int) {
	fmt.Println("\tBefore Writing 100 into the channel")
	ch <- 100
	fmt.Println("\tAfter Writing 100 into the channel")
	fmt.Println("\tBefore Writing 200 into the channel")
	ch <- 200
	fmt.Println("\tAfter Writing 200 into the channel")
	fmt.Println("\tBefore Writing 300 into the channel")
	ch <- 300
	fmt.Println("\tAfter Writing 300 into the channel")
	fmt.Println("\tBefore Writing 400 into the channel")
	ch <- 400
	fmt.Println("\tAfter Writing 400 into the channel")
}

func main() {
	fmt.Println("Starting the main function")
	ch := make(chan int, 2)
	go writeData(ch)

	fmt.Println("Before reading 100 from the channel")
	fmt.Println("Data from the channel ", <-ch)
	fmt.Println("After reading 100 from the channel")
	/*
		fmt.Println("Before reading 200 from the channel")
		fmt.Println("Data from the channel ", <-ch)
		fmt.Println("After reading 200 from the channel")
		fmt.Println("Before reading 300 from the channel")
		fmt.Println("Data from the channel ", <-ch)
		fmt.Println("After reading 300 from the channel")
		fmt.Println("Before reading 400 from the channel")
		fmt.Println("Data from the channel ", <-ch)
		fmt.Println("After reading 400 from the channel") */
	time.Sleep(3 * time.Second)
	//fmt.Println("Data from the channel ch is ", <-ch)

}
