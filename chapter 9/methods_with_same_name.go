package main

import "fmt"

type mililitre float64
type litre float64
type gallon float64

func (l litre) toGallons() gallon {
	return gallon(l * 0.264)
}

func (m mililitre) toGallons() gallon {
	return gallon(m * 0.000264)
}

//unlike function methods can have same name
//with restriction that receiver parametres type should be different
func main() {
	soda := litre(1)
	water := mililitre(5000)
	fmt.Printf("soda in gallons : %0.3f\n", soda.toGallons())
	fmt.Printf("water in gallons : %0.3f", water.toGallons())
}
