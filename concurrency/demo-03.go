package main

import (
	"fmt"
	"time"
)

func print(message string, delay time.Duration, in chan string, out chan string) {
	for {
		<-in
		time.Sleep(delay * time.Second)
		fmt.Println(message)
		out <- "done"
	}
}

func main() {
	/* Hello & World shoule be printed alternatively */
	ch1 := make(chan string)
	ch2 := make(chan string)
	go print("Hello", 1, ch1, ch2) /* print with the delay of 1 sec */
	go print("World", 2, ch2, ch1) /* print with the delay of 2 sec */
	ch1 <- "start"
	var input string
	fmt.Println("Enter any key to exit")
	fmt.Scanln(&input)
}
