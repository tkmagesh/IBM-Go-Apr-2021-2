package main

import "fmt"

//func readFile() (words []string, err error) {
//open file
//defer fs.close()
//read file
/*
	if error
		err = error
		return
*/
//processing file contents
/*
	if error
		err = error
		return
*/

//return filecontents
//}

func fn() (ret int) {
	fmt.Println("fn initiated")
	defer func() {
		ret = 300
		fmt.Println("deferred function f1 invoked")
	}()
	defer func() {
		fmt.Println("deferred function f2 invoked")
	}()
	defer func() {
		fmt.Println("deferred function f3 invoked")
	}()
	ret = 100
	fmt.Println("fn completed")
	return
}

func main() {
	fmt.Println(fn())
}
