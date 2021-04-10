package main

import (
	"errors"
	"fmt"
)

func add(x, y int) int {
	return x + y
}

/*
func divide(x, y int) (int, int) {
	return x / y, x % y
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = errors.New("Divisor cannot be zero!")
		return
	}
	quotient = x / y
	remainder = x % y
	return
}

func main() {
	subtract := func(x, y int) int {
		return x - y
	}
	//fmt.Println(add(100, 200))
	fmt.Println(subtract(100, 200))
	quotient, _, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(quotient)
}
