package main

import "fmt"

func main(){
	// declarations and intialization
	var int_arr =[5] int {1,2,3,4,5}
	//int_arr =[5] int {1,2,3,4,5} //short hand declaration
	var str_arr [3] string
	str_arr[0]="ayushman"
	str_arr[1]="hello"

	// printing the data
	for i:=0;i<5;i++ {  //traditional way
		fmt.Println(int_arr[i])
	}

	for index,value:= range int_arr{   // for ...range way provided by go for safe looping
									   // now lets say if index is not needed use _ instead as place holder(we have done while dealing with error)
		fmt.Println(index,"\t",value)
	} 

	fmt.Println(str_arr[1],str_arr[0])
	fmt.Println(int_arr)
}