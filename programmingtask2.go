package main

import (
	"fmt"
)

func sum(num1,num2 float64) float64{
	return num1+num2
}

func mul(num1,num2 float64) float64{
	return num1*num2
}

func sub(num1,num2 float64) float64{
	return num1-num2
}

func div(num1,num2 float64) float64{
	return num1/num2
}

func mathOperations(funcName func(no1,no2 float64) float64) float64{    //passing function as argument to another fucntion
	return funcName(7.6,5.4)
}


func main(){
	fmt.Printf("%0.2f + %0.2f = %0.2f\n",7.6,5.4,mathOperations(sum))
	fmt.Printf("%0.2f + %0.2f = %0.2f\n",7.6,5.4,mathOperations(sub))
	fmt.Printf("%0.2f + %0.2f = %0.2f\n",7.6,5.4,mathOperations(mul))
	fmt.Printf("%0.2f + %0.2f = %0.2f\n",7.6,5.4,mathOperations(div))

}