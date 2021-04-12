package main

import (
	"fmt"
)

func print(message string) {
	fmt.Println(message)
}

func main() {
	print("Hello")
	go print("World")
	fmt.Println("End of main!")
	//time.Sleep(4 * time.Second)
	var input string
	fmt.Scanf("%s", &input)
}
