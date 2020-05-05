package main

import "fmt"

func func2(){
	fmt.Println("U called function 2")
}

func func1() func() {
	fmt.Println("U called function 1")
	return func2
}

func main(){
	tempFunc:=func1()
	tempFunc()	
}