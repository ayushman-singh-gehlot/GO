package main

import (
	"fmt"
)

func main(){

	myMap:= map[string] float64{"a":1.2,"b":3.4,"c":0}
	fmt.Println(myMap)
	value,ok:=myMap["a"]
	fmt.Println(value,ok)
	value,ok=myMap["b"]
	fmt.Println(value,ok)
	value,ok=myMap["c"]
	fmt.Println(value,ok)
	value,ok=myMap["d"]
	fmt.Println(value,ok) // by using ok variable we can put a
						// criteria for not using unassigned value


}