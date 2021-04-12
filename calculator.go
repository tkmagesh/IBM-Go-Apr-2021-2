package main

import (
	"fmt"
	"os"
)

var operations map[int]func(int, int) int = map[int]func(int, int) int{
	1: func(a, b int) int {
		return a + b
	},
	2: func(a, b int) int {
		return a - b
	},
	3: func(a, b int) int {
		return a * b
	},
	4: func(a, b int) int {
		return a / b
	},
}

func main() {
	var choice int
	var a, b int
	for {
		fmt.Println("Please choose a valid option :")
		displayMenu()
		fmt.Scanf("%d", &choice)
		if choice == 5 {
			os.Exit(0)
		}
		if operation, operExists := operations[choice]; operExists {
			fmt.Println("Enter two numnber")
			fmt.Scanf("%d%d", &a, &b)
			result := operation(a, b)
			fmt.Printf("result is %v\n", result)
		}

	}
}
func displayMenu() {
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
}
