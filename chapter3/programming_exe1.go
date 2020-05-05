package main
import (
	"fmt"
)

func negate(val *bool) {
	*val=! *val
}

func main(){
	value:=true
	fmt.Println("before  : ",value)
	negate(&value)
	fmt.Println("after   : ",value)
}