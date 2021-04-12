package main

import "fmt"

func main() {
	var x interface{}
	x = 10
	//x = "abc"
	fmt.Println(x)
	value, ok := x.(int)
	if ok {
		fmt.Printf("%v is of type int\n", value)
	} else {
		fmt.Printf("%v is not int type\n", x)
	}

	switch v := x.(type) {
	case int:
		fmt.Printf("Twice of %v is %v\n", v, v*2)
	case string:
		fmt.Printf("Len of %v is %v\n", v, len(v))
	default:
		fmt.Println("Unknown type of %v = %T", v, v)

	}
}
