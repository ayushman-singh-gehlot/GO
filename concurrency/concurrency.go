package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	temp := runtime.GOMAXPROCS(2) //The argument passed inside GOMAXPROCS will decide
	//the number of CPU/thread for that go program it returns previous setting(i.e generally 8) to var temp
	fmt.Println(temp)
	wg.Add(2) // add two go routines if u make it one it will wait for only one function to get complete
	// and then it will terminate main
	go delayedIteration1()
	go delayedIteration2() // created two seprate go routines(thread)
	//fmt.Scanln() // this was just called to create delay if we remove it main will exit first
	wg.Wait()
}

func delayedIteration1() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("In 1 : Second : ", i)
	}
	wg.Done()
}
func delayedIteration2() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("In 2 : Second : ", i)
	}
	wg.Done()
}
