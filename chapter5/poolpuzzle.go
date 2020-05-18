package main

import "fmt"

func main(){
	numbers:=[6] int{3,16,-2,10,23,12}

	for i,number:= range numbers{   // for ...range way provided by go for safe looping
		if number>=10 && number<=20{
			fmt.Println(i,number)
		}									   // now lets say if index is not needed use _ instead as place holder(we have done while dealing with error)
	}

}