package main

import "fmt"

var mystruct struct {
	name   string
	rollNo int
}

func main() {
	mystruct.name = "ayushman singh gehlot"
	mystruct.rollNo = 107
	fmt.Println(mystruct.name, mystruct.rollNo)

}
