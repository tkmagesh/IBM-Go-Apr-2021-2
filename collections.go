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

	//slice
	/*
		dynamic size
		highlevel apis for manipulation
	*/
	fmt.Println("Slice")
	nosSlice := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("nosSlice = %v\n", nosSlice)

	nosSlice = append(nosSlice, 6, 7, 8)
	fmt.Printf("after appending, nosSlice = %v\n", nosSlice)
	fmt.Printf("Len of nosSlice = %d\n", len(nosSlice))
	fmt.Printf("Capacity of nosSlice = %d\n", cap(nosSlice))

	/*
		nosSlice[lo : hi] => return from lo to (hi - 1)
		nosSlice[lo : lo] => empty
		nosSlice[lo : lo+1] => one element (at lo)
		nosSlice[lo : ] => all the values from lo
		nosSlice[:hi] => all the values from 0 to hi-1
	*/

	fmt.Printf("nosSlice[1:4] => %v\n", nosSlice[1:4])
	fmt.Printf("nosSlice[1:1] => %v\n", nosSlice[1:1])
	fmt.Printf("nosSlice[1:2] => %v\n", nosSlice[1:2])
	fmt.Printf("nosSlice[5:] => %v\n", nosSlice[5:])
	fmt.Printf("nosSlice[:5] => %v\n", nosSlice[:5])

	fmt.Println("using range construct")
	evenNos := []int{2, 4, 6, 8, 10}
	for idx, evenNo := range evenNos {
		fmt.Printf("evenNo at %d = %d\n", idx, evenNo)
	}

	fmt.Printf("\nMap\n")

	cityRanks := map[string]int{
		"Udupi":     2,
		"Mysuru":    1,
		"Mangaluru": 3,
		"Bengaluru": 4,
	}

	//cityRanks := make(map[string]int, 200);
	fmt.Printf("City Ranks : %v\n", cityRanks)
	fmt.Println("Adding a new key/value pair")

	//adding a new key/value pair
	//cityRanks["Tumukuru"] = 5
	appendCityRank(cityRanks, "Tumukuru", 5)
	fmt.Printf("After adding Tumukuru, City Ranks : %v\n", cityRanks)

	//iteration using the range
	fmt.Println("Using range over map")
	for city, rank := range cityRanks {
		fmt.Printf("City = %s, Rank = %d\n", city, rank)
	}

	//delete a value
	delete(cityRanks, "Bengaluru")
	fmt.Println("After removing Bengaluru")
	fmt.Printf("City Ranks : %v\n", cityRanks)

	_, exists := cityRanks["Bengaluru"]
	fmt.Println("Is Bengaluru Ranked = ", exists)

}

func appendCityRank(cityRanks map[string]int, city string, rank int) {
	cityRanks[city] = rank
}
