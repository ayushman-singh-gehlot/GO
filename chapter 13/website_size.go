package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	go responseSize("https://golang.org/")
	go responseSize("https://golang.org/doc")
	go responseSize("https://google.com/")
	time.Sleep(5 * time.Second)

}

func responseSize(url string) {
	fmt.Println("getting url : ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("url:- ", url, "\t\t\tsize:- ", len(body))
}
