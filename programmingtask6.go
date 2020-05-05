package main

import "fmt" 
func manyInt(numbers ... int) {
	fmt.Println(numbers)
	println(numbers)
}

func main(){
	manyInt(2,3,4,5)	
}