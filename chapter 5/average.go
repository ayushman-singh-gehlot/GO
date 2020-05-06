package main

import (
	"fmt"
	"os"
	"log"
	"strconv"
	"bufio"
)

func main(){
	file,err:=os.Open("data.txt")
	if err != nil{
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var numbers [] float64
	//var numbers [3] float64  changes made after studying slices
	for scanner.Scan(){
		temp,err:=strconv.ParseFloat(scanner.Text(),64)
		if err != nil{
			log.Fatal(err)
		}
		numbers=append(numbers,temp)
	}
	err=file.Close()
	if err != nil{
		log.Fatal(err)
	}
	if scanner.Err() != nil{
		log.Fatal(scanner.Err())
	}


	var sum float64 // intialized with value 0

	for _,value:= range numbers{   // for ...range way provided by go for safe looping
									   // now lets say if index is not needed use _ instead as place holder(we have done while dealing with error)
		sum=sum+value
	}
	sample:=float64(len(numbers))
	fmt.Println("average beef required is : ",(sum/sample)) 

}