package main

import "fmt"

func main(){
	// declarations and assignment to slices
	// style 1
	var intslice [] int
	intslice=make([]int ,7)
	// intslice:=make([]int ,7) another short hand technique
	intslice[0]=0
	intslice[1]=1
	intslice[3]=3
	intslice[5]=5
	intslice[6]=6
	
	// style 2 short hand declarations 
	notes:=[] string {"sa","re","ga","ma","pa","dha","ni","sa"}


	// printing data
	for _,value:=range(intslice){
		fmt.Println(value)
	}
	for _,value:=range(notes){
		fmt.Println(value)
	}

}