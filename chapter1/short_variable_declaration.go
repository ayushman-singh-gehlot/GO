package main

import "fmt"

func main() {
	rollNo := 107
	name := "ayushman singh gehlot"
	cgpa := 6.01
	// this type of declaration can only be used when u know the value of variable before hand

	// rollNo,name,cgpa:=107,"ayushman singh gehlot",6.01
	// declaring all variable(even of diff types ) in same line
	// refer page no 55 for more info
	fmt.Println(rollNo, "\t", name, "\t", cgpa)
}
