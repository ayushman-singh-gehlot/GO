package main

import "fmt"

func main(){
	slice:= [] string{"a","b"}
	fmt.Println(slice,len(slice))
	slice=append(slice,"c") // adding   			
	//calling append, it’s conventional to just assign the 
	//return value back to the same slice variable you passed to append. 
	// for explanation refer pg 267								
										
	fmt.Println(slice,len(slice))

	slice1:= [] string{"d","e"}
	slice=append(slice,slice1 ...) 			 
										
	fmt.Println(slice,len(slice))
}