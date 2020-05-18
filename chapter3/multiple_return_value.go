package main
import (
	"fmt"
	"math"
)

func floatParts(number float64) (int,float64) {
	integerpart:=math.Floor(number)
	return int(integerpart),number-integerpart
}
func main(){
	intpart,remainder:=floatParts(1.9)
	fmt.Println(intpart,remainder)
}