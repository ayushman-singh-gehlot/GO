package main

import (
	"fmt"
	"reflect"
)

type litre float64
type gallons float64

func main() {
	var carfuel litre // declared a variable of type litre
	carfuel = 10

	busfuel := gallons(20) // short hand notation with type casting

	fmt.Println(carfuel, reflect.TypeOf(carfuel), "\n", busfuel, reflect.TypeOf(busfuel))

}
