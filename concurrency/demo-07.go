package main

import (
	"fmt"
	"time"
)

func print(s string, ch chan string, delay time.Duration) {
	fmt.Println(s)
	time.Sleep(delay * time.Second)
	ch <- s + " - done!"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go print("Magesh", ch1, 7)
	go print("Suresh", ch2, 3)
	go print("Ramesh", ch3, 5)

	/*
		fmt.Println(<-ch1)
		fmt.Println(<-ch2)
		fmt.Println(<-ch3) */

	time.Sleep(6 * time.Second)

	for {
		select {
		case result1 := <-ch1:
			fmt.Println(result1)
		case result2 := <-ch2:
			fmt.Println(result2)
		case result3 := <-ch3:
			fmt.Println(result3)
		default:
			break
		}
	}

	input := ""
	fmt.Scanln(&input)
}
