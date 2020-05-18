package main

import (
	"fmt"
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	fmt.Print("Enter a grade : ")
	inputobject:=bufio.NewReader(os.Stdin)
	inp_grade,err:=inputobject.ReadString('\n')
	if err!=nil{   	// check if the error is nil or not 
		log.Fatal(err)// if error is not nil only then it prints the error and terminate the program
	}
	inp_grade=strings.TrimSpace(inp_grade)
	grade,err:=strconv.ParseInt(inp_grade,10,64) // 10 for decimal and 64 is for int64
	if err!=nil{   	// check if the error is nil or not 
		log.Fatal(err)// if error is not nil only then it prints the error and terminate the program
	}
	var status string // scope of the status is for rest of the function
	if grade==100{
		status="perfect!!"
	}else if grade>=60{
		status="pass"
	}else{
		status="fail"
	}
	fmt.Println("A grade of",grade,"is",status)

}
