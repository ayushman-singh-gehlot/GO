package main

import "fmt"

func f1() {
	defer f3()
	fmt.Println("inside function f1")
}
func f2(num int) {
	fmt.Println("inside function f2. number :", num)
}
func f3() {
	fmt.Println("inside function f3")
}

func main() {
	// with use of defer it follows LIFO(stack approach)
	number := 10
	defer f1()
	defer f2(number)
	number = 20
	fmt.Println("inside main. number :", number)
}

// in this code first main will be called
// var number will be declared
//it will push f1 on stack
//then f2 on stack with argument value as 10
//then it will print inside main
//call f2(10)
//print("inside f2")
//call f1()
//push f3()
//print("inside f1")
// call f3()
//print("inside f3")
// for more clarification refer to go lec 41
