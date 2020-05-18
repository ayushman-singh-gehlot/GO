package main

import (
	"fmt"
	"time"
)

func printa() {
	for i := 0; i < 50; i++ {
		fmt.Print("a")
	}
}
func printb() {
	for i := 0; i < 50; i++ {
		fmt.Print("b")
	}
}

func main() {
	go printa()
	go printb() // both function will execute simultaneously
	// VVIMP function used as goroutine cannot return any value refer pg541
	time.Sleep(time.Second)

}
