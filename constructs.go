package main

import "fmt"

func main() {

	/* if else construct */
	if no := 31; no%2 == 0 {
		fmt.Printf("%v is even of type %T\n", no, no)
	} else {
		fmt.Printf("%v is odd of type %T\n", no, no)
	}

	//for construct
	/*
		sum := 0
		for i := 0; i < 10; i++ {
			sum += i
		}
		fmt.Println(sum)
	*/

	//for (used like a while loop)
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)

	//infinite loop
	for {
		break
	}

	//switch
	/*
		no := 51
		remainder := no % 2
		switch remainder {
		case 0:
			fmt.Printf("%v is even\n", no)
		case 1:
			fmt.Printf("%v is odd\n", no)
		}
	*/

	// switch - compared again multiple values
	/*
		score := 13
		switch score {
		case 0, 1, 2, 3:
			fmt.Println("Terrible")
		case 4, 5, 6, 7:
			fmt.Println("Not bad!")
		case 8, 9:
			fmt.Println("Good")
		case 10:
			fmt.Println("Excellect")
		default:
			fmt.Println("Unknown grade!")
		}
	*/

	//using 'falthrough' construct
	var n int
	fmt.Println("Enter the value :")
	fmt.Scanf("%d", &n)
	switch n {
	case 0:
		fmt.Println("is zero")
		fallthrough
	case 1:
		fmt.Println("is <= 1")
		fallthrough
	case 2:
		fmt.Println("is <= 2")
		fallthrough
	case 3:
		fmt.Println("is <= 3")
		fallthrough
	case 4:
		fmt.Println("is <= 4")
		fallthrough
	case 5:
		fmt.Println("is <= 5")
		fallthrough
	default:
		fmt.Println("Try again!")
	}
}
