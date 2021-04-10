//pakcage declaration
package main

//package imports
import (
	"fmt"
	"hello-world/calculator"
)

//package variables

//package init fn

func main() {
	/*
		var msg string
		msg = "Hello World"
	*/

	/*
		var msg string = "Hello World"
	*/

	/*
		var msg = "Hello World"
	*/
	msg := "Hello World"
	fmt.Println(msg)

	/*
		var x int
		var y int
		x = 10
		y = 20
	*/
	/*
		var x, y int
		x = 10
		y = 20
	*/
	/*
		var x, y int = 10, 20
	*/

	/* x, y := 10, 20 */

	var (
		x, y int = 10, 20
		s    string
	)
	s = "somthing"
	//sum := sum(x, y)
	sum := calculator.Add(x, y)

	fmt.Println(sum)
	fmt.Println(s)

	fmt.Println(greet("Magesh"))
	fmt.Println(defaultGreet)
	fmt.Println(calculator.MyVar)
}

//other functions
/* func sum(x, y int) int {
	return x + y
} */
