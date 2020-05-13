package main

import "fmt"

func main() {
	// panic
	fmt.Println("hello")
	panic("panicking: i don't know what to do!!")
	fmt.Println("end of main")
}
