package main

import (
	"fmt"
	"strconv"
)

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

	/*
		var jsonBlob = []byte(`[
				{"Name": "Platypus", "Order": "Monotremata", "Forest" : { "Country" : "India"}},
				{"Name": "Quoll",    "Order": "Dasyuromorphia", "Forest" : { "Country" : "Brazil"} }
			]`)
		var animals []map[string]interface{}
		err := json.Unmarshal(jsonBlob, &animals)
		if err == nil {
			for _, animal := range animals {
				fmt.Println(animal["Name"], animal["Order"], animal["Forest"])
			}
		} else {
			fmt.Println("error = ", err)
		}
	*/

	fmt.Println("Sum of 10,20,30,40 = ", sum(10, 20, 30, 40))
}

/*
Use cases:
sum(10,20) => 30
sum(10,20,30,40,50) => 150
sum(10) => 10
sum() => 0
sum(10,20,30,"40") => 100
sum(10, "20", 30 , "abc") => 60
sum(10,20, []int{30,40,50}) => 150
sum(10,20, []interface{}{10,40,"50"}) => 150
*/

func sum(nos ...interface{}) int {
	result := 0
	for _, no := range nos {
		switch no.(type) {
		case int:
			result += no.(int)
		case string:
			if val, err := strconv.Atoi(no.(string)); err == nil {
				result += val
			}
		case []int:
			noList := no.([]int)
			intfList := make([]interface{}, len(noList))
			for idx, x := range noList {
				intfList[idx] = x
			}
			result += sum(intfList...)
		case []interface{}:
			result += sum(no.([]interface{})...)
		}

	}
	return result
}
