package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func getResponse(ch chan *http.Response) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}
	ch <- res
}

func main() {
	ch := make(chan *http.Response)
	go getResponse(ch)
	for {
		select {
		case res := <-ch:
			io.Copy(os.Stdout, res.Body)
			return
		case <-time.After(250 * time.Millisecond):
			log.Fatalln("Request timeed out!!")
		}
	}
}
