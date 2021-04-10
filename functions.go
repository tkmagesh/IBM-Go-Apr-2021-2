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

func calculate(x, y int, operation func(int, int) int) int {
	//validate x & y
	return operation(x, y)
}

func getAdder() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func getAdderFor(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	subtract := func(x, y int) int {
		return x - y
	}
	//fmt.Println(add(100, 200))
	fmt.Println(subtract(100, 200))
	//fmt.Println(calculate(10, 20, add))

	/*
		add := func(x, y int) int {
			return x + y
		}
		fmt.Println(calculate(10, 20, add))
	*/

	fmt.Println(calculate(10, 20, func(x, y int) int {
		return x + y
	}))

	//functions as return values
	adder := getAdder()
	fmt.Println(adder(100, 200))

	adderFor100 := getAdderFor(100)
	fmt.Println(adderFor100(200))

	quotient, _, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(quotient)

}
