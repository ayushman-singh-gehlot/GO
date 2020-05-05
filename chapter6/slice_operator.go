package main

import "fmt"

func main(){
	var underlyingarray = [7] int {0,1,2,3,4,5,6} //indexs from 0 to 6

	slice1:= underlyingarray[2:5]  
	fmt.Println(slice1)						// 2-4

	slice2:= underlyingarray[3:]
	fmt.Println(slice2)						//3-6

	slice3:= underlyingarray[:3]
	fmt.Println(slice3)						//0-2

	// if elements of underlying arrays are changed then slice elements also gets changed
	underlyingarray[0]=100 // changes in array
	underlyingarray[1]=101

	slice4:=underlyingarray[0:3]
	fmt.Println(slice4)

}