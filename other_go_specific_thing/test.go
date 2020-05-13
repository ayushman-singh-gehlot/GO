package main

import (
	"bufio"
	"fmt"
	"os"
)

func f1(input string) {
	fmt.Println("inside function f1\nargument : ", input)
}

func main() {
	defer f1(input)
	readObject := bufio.NewReader(os.Stdin)
	input, _ := readObject.ReadString('\n')

	input = input + input
	fmt.Println("inside main")
}
