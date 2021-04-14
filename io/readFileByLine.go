package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	inputHandle, inputError := os.Open("sample.txt")
	if inputError != nil {
		log.Fatalln("Error opening file", inputError)
	}
	defer inputHandle.Close()
	inputReader := bufio.NewReader(inputHandle)
	for {
		line, err := inputReader.ReadString('\n')
		fmt.Print(line)
		if err == io.EOF {
			break
		}
	}

}
