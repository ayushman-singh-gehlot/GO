package main

import "fmt"

type square struct {
	side float64
}

func (s *square) perimeter() float64 {
	return 4.0 * s.side
}
func (s *square) area() float64 {
	return s.side * s.side
}

type shape interface {
	perimeter() float64
	area() float64
}

func display(sh shape) {
	fmt.Println("perimeter is : ", sh.perimeter())
	fmt.Println("area is : ", sh.area())
}

func main() {
	var sq shape
	square1 := &square{side: 20} // & is necessary when ur methods are using pointer receiver parameter
	sq = square1
	display(sq)
}
