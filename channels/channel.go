package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go writeToChannel(ch)
	// why we do not see any delay or sync in this program even though we are using go?
	value := <-ch // reading data from channel is similar to scan function
	//it waits till it get any valid data
	fmt.Println("data in our channel is : ", value)

	// temp := <-ch  // point to remember data from channel is not coppied rather it is transffered from channel
	// fmt.Println(temp)

}

func writeToChannel(ch chan int) {
	//time.Sleep(time.Second)
	fmt.Println("insiode go routine")
	ch <- 10
	fmt.Println("data has been written to channel")
}
