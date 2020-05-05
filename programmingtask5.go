package main

import "fmt"

func tempFunction(arr [5] int) [5] float64{
	var arrRet [5] float64
	for i,value:=range arr{
		arrRet[i] = float64(value)/3.0
	}
	return arrRet
}

func main(){
	arrInt:=[5] int{5,7,9,14,17}
	arrFloat:=tempFunction(arrInt)
	for _,value:=range arrFloat{
		fmt.Printf("%0.2f\t",value)
	}
}