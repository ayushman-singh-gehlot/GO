package main
import (
	"fmt"
)

func swap(no1 *int,no2 *int){
	temp:=*no1
	*no1=*no2
	*no2=temp
}

func main(){
	num1,num2:=5,9
	fmt.Println("before swapping : ",num1,num2)
	swap(&num1,&num2)
	fmt.Println("after  swapping : ",num1,num2)
}