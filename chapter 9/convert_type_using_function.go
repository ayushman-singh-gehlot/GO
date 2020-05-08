package main

import "fmt"

type litre float64
type gallons float64

func toGallons(l litre) gallons {
	return gallons(l * 0.264)
}

func toLitre(g gallons) litre {
	return litre(g * 3.785)
}

func main() {
	var carfuel litre // declared a variable of type litre
	var busfuel gallons
	carfuel = toLitre(gallons(5))
	busfuel = toGallons(litre(100))
	fmt.Printf("car fuel in litres : %0.2f\n", carfuel)
	fmt.Printf("bus fuel in gallons : %0.2f", busfuel)
}
