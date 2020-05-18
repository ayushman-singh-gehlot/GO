package main

import "fmt"

const myConstant = 30       //untyped constant
const otherCon = "ayushman" //int constant
const (
	cat = iota
	dog // automatically assign incremented value iota
	camel
	horse
)

func main() {
	//myConstant=40// we cannot modify constants
	fmt.Println(myConstant)
	fmt.Println(otherCon)
	fmt.Println(cat, dog, camel, horse)

}
