package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fileContents, err := ioutil.ReadFile("sample.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(fileContents))
	err = ioutil.WriteFile("sample-copy.txt", fileContents, 0x777)
	if err != nil {
		log.Fatalln(err)
	}
}
