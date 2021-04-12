package main

import (
	"fmt"
	"time"
)

func writeData(ch chan int) {
	fmt.Println("Writing data into the channel")
	ch <- 100
	fmt.Println("Writing subsequent data into the channel")
	ch <- 200
	fmt.Println("Writing subsequent data into the channel")
	ch <- 300
	fmt.Println("Writing subsequent data into the channel")
	ch <- 400
}

func main() {
	fmt.Println("Starting the main function")
	ch := make(chan int)
	//go writeData(ch)
	fmt.Println("attempting to read from the channel")
	data := <-ch
	fmt.Println("Data from the channel ch is ", data)
	time.Sleep(3 * time.Second)
	fmt.Println("Data from the channel ch is ", <-ch)

}
