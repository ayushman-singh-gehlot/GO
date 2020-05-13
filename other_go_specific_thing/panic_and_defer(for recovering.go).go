package main

import "fmt"

func rec(){
	details:=recover()
	fmt.Println("Details : "details)
}

func main() {
	// panic
	defer rec()
	fmt.Println("hello")
	panic("panicking: i don't know what to do!!")
	fmt.Println("end of main")
}

// as soon as we encounter panic in main
// it leaves all the code in the main after panic statement
// it checks whether we have any thing on stack
// found rec() function call
//call rec
// recovered from panic 
// printed details for the reason to panic
