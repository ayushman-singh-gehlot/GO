package main

import (
	"fmt"
	"chapter4/calc"
)

func main(){
	numbers:=[] float64{6.7, 8.9, 4.3, 2.9, 10.9}
	fmt.Println(numbers)
	max_num:=calc.Maximum(numbers ...)  // when we use ...(ellipsis) we pass elements of slice one by one two function 
										// generally used while passing arguments in variadic function
	fmt.Println(max_num)
	
}