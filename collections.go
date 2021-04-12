package main

import "fmt"

func main() {
	//Array
	var nos [10]int
	fmt.Println(len(nos))
	for i := 0; i < 10; i++ {
		nos[i] = i
	}
	fmt.Printf("nos = %v\n", nos)

	//initialization while creation
	var names [2]string = [2]string{"Magesh", "Suresh"}
	fmt.Printf("names = %v\n", names)

	//multidimensional array
	var matrix [2][3]string
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = fmt.Sprintf("row-%d*col-%d", i, j)
		}
	}
	fmt.Printf("matrix - %v\n", matrix)
}
