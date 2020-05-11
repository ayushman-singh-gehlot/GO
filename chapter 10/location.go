package main

import (
	"fmt"
	"log"
	"packages/geo"
)

func main() {
	location := geo.Landmark{}
	err := location.SetName("Indore")
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLatitude(43)
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLongitude(100)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(location.Name())
	fmt.Println(location.Latitude())
	fmt.Println(location.Longitude())
}
