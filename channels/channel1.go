package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	ch := make(chan int, 1) // buffered channel with 1 capacity
	go writeToChannel(ch)
	go readFromChannel(ch)
	wg.Wait()

}

func writeToChannel(ch chan int) {
	fmt.Println("starting to write")
	ch <- 10 // 1st this data will be put to channel
	ch <- 8  // 3rd this data will be put to channel
	fmt.Println("data has been written to channel")
	wg.Done()
}

func readFromChannel(ch chan int) {
	fmt.Println("staring to read")
	time.Sleep(time.Second)
	value := <-ch // 2nd data would be read from channel
	fmt.Println("data in our channel is : ", value)
	wg.Done()
}
