package main

import "fmt"

type number int

func (n number) add(otherNumber int) {
	fmt.Printf("%d + %d = %d\n", n, otherNumber, int(n)+otherNumber)
}

func (n number) subtract(otherNumber int) {
	fmt.Printf("%d - %d = %d\n", n, otherNumber, int(n)-otherNumber)
}

func main() {
	ten := number(10)
	ten.add(4)
	four := number(4)
	four.subtract(2)
}
