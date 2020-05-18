package main

import (
	"fmt"
	"reflect"
)

// type conversion at page 58
func main() {
	length := 2
	width := 4.6
	fmt.Println("length is of type ", reflect.TypeOf(length), "\n", "width is of type ", reflect.TypeOf(width))
	// we cannot perfrom any comparison and arithmetic operation on variable of diff datatypes
	//fmt.Println("area is ",length*width) // should through error
	//length=float64(length) // cannont assign float number to int type variable
	area := float64(length) * width
	fmt.Println("area is ", area)
	fmt.Println("rounding area to smallest  integer", int(area))
}
