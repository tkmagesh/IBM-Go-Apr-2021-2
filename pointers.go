package main

import "fmt"

func main() {

	var no int = 10
	var addressOfNo *int = &no
	//deferencing (accessing the value using the address)
	fmt.Println("Data at addressOfNo = %d\n", *addressOfNo)

	increment(&no)
	fmt.Printf("After incrementing, no = %d\n", no)

	x, y := 10, 20
	swap(x, y)
	fmt.Printf("After swapping x, y = %d, %d\n", x, y)
}

func increment(no *int) {
	*no = *no + 1
}

func swap(x, y int) {
	temp := x
	x = y
	y = temp
}
