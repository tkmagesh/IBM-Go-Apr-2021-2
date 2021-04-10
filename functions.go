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

	quotient, _, err := divide(10, 1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(quotient)
	count := func() func() int {
		counter := 0
		return func() int {
			counter += 1
			return counter
		}
	}()
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())

	up, down := func() (func() int, func() int) {
		counter := 0
		up := func() int {
			counter += 1
			return counter
		}
		down := func() int {
			counter -= 1
			return counter
		}
		return up, down
	}()

	fmt.Println(up())
	fmt.Println(up())
	fmt.Println(up())

	fmt.Println(down())
	fmt.Println(down())
	fmt.Println(down())
	fmt.Println(down())
	fmt.Println(down())

	evenNos := getNos(func(no int) bool {
		return no%2 == 0
	})
	fmt.Println("Even nos")
	fmt.Println(evenNos())
	fmt.Println(evenNos())
	fmt.Println(evenNos())
	fmt.Println(evenNos())

	isPrime := func(n int) bool {
		if n == 1 {
			return false
		}
		for i := 2; i <= (n / 2); i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	primeNos := getNos(isPrime)
	fmt.Println("Prime Nos :")
	fmt.Println(primeNos())
	fmt.Println(primeNos())
	fmt.Println(primeNos())
	fmt.Println(primeNos())
	fmt.Println(primeNos())
	fmt.Println(primeNos())

}

func getCounter() func() int {
	counter := 0
	return func() int {
		counter += 1
		return counter
	}
}

/*
up() => 1
up() => 2
up() => 3

down() => 2
down() => 1
down() => 0
down() => -1

*/

func getNos(predicate func(int) bool) func() int {
	counter := 0
	return func() int {
		for {
			counter += 1
			if predicate(counter) {
				return counter
			}
		}

	}
}

/*
var evenNos = genNos()
evenNos() //=> 2
evenNos() //=> 4
evenNos() //=> 4

var oddNos = getNos()
oddNos() //=> 1
oddNos() //=> 3
oddNos() //=> 5

var nosBy3 = getNos()
nosBy3() //=> 3
nosBy3() //=> 6
nosBy3() //=> 9  */
