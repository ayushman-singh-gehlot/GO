package main
import (
	"fmt"
)

func manyInt(numbers ... int){
	fmt.Println(numbers)
}

func mix(number int ,flag bool, strings ... string){ // first two arguments are nonvariadic therefore they are compulsory
	fmt.Println(number,flag,strings)
}											// followed by first two argumnets any number of string arguments can be passed

func main(){
	fmt.Println("calling function manyInt")
	manyInt()
	manyInt(3)
	manyInt(6,7,9,0)
	fmt.Println("calling function mix")
	mix(5,true)
	mix(4,false,"a","b","ayushman")
}