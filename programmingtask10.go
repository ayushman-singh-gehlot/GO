package main

import "fmt"

func checkType(anyType interface{}) { // any type of data can be passed here
	switch anyType.(type) {
	case string:
		fmt.Println("argument is string")
	case float64:
		fmt.Println("argument is float64")
	case float32:
		fmt.Println("argument is float32")
	case int:
		fmt.Println("argument is integer")
	default:
		fmt.Println("other than my cases")
	}

}
func main() {
	var i interface{} = 10
	storeVar, ok := i.(int)
	i = "hello"
	i = 7.8
	// var of interface type can store any type of value

	fmt.Println(storeVar, ok)
	checkType("hello")
}
