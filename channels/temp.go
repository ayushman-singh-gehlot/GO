package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	go writeToChannel(ch)
	time.Sleep(time.Second)
	fmt.Println(len(ch))
	value1 := <-ch
	fmt.Println(value1, len(ch))
	// value2 := <-ch
	// fmt.Println(value2, len(ch))
	fmt.Println("end of main")
}

// unbuffered ch will return len(ch) as 0
// whereas ch with capacity return no of value that are not read till now

func writeToChannel(ch chan int) {
	fmt.Println("insiode go routine")
	ch <- 10
	fmt.Println("data has been written to channel")
}
