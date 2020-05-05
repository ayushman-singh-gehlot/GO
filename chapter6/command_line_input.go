package main

import(
	"fmt"
	"os"
	"log"
	"strconv"
)

func main(){
	arguments:=os.Args
	var sum float64

	for _,value:=range arguments{
 
		number,err:=strconv.ParseFloat(value,64) 	// Q->how does redeclarations work here  ??
													// ANS->because after every iteration old varibales are destroyed 
													// and new variable with name number is created
													// thats why it is not redeclarations and simply declarations
		if err!=nil{
			log.Fatal(err)
		}
		sum=sum + number
	}
	sampleCount:=float64(len(arguments))
	fmt.Printf("average beef required is %0.2f",sum/sampleCount)
}