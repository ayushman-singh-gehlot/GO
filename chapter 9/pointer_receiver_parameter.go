package main

import "fmt"

type number int

func (n *number) double() {
	*n = *n * 2
}
func main() {
	no := number(4)
	no.double()
	fmt.Println(no)
}
