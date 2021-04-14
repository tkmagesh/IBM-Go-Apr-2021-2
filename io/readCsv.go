package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	data := `
firstname, lastname, age
Magesh, Kuppan, 40
Suresh, Kannan, 35
Ramesh, Jayaraman, 43
`
	reader := csv.NewReader(strings.NewReader(data))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(record)
	}
}
