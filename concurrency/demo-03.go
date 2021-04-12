package main

import "fmt"

func print(message string /*  */) {
	fmt.Println(message)
}

func main() {
	/* Hello & World shoule be printed alternatively */
	print("Hello") /* print with the delay of 1 sec */
	print("World") /* print with the delay of 2 sec */
	var input string
	fmt.Println("Enter any key to exit")
	fmt.Scanf(&input)
}
