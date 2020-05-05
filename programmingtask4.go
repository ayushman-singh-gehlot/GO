package main

import "fmt"

func division(no1,no2 int) float64{
	return float64(no1/no2)
}

func func1() func(no1,no2 int) float64 {
	fmt.Println("U called function which returns division function")
	return division
}

func main(){
	tempFunc:=func1()
	fmt.Println(tempFunc(5,2))	
} 