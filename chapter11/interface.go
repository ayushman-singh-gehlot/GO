package main

import (
	"fmt"
)

type pen string // what type can be assigned?

type pencil int // what type can be assigned?

func (p pen) write() {
	fmt.Println("writing through pen")
}

func (p pencil) write() {
	fmt.Println("writing through pencil")
}

type colorFullPen string

func (p colorFullPen) write() {
	fmt.Println("writing through colorFullpen")
}
func (p colorFullPen) writeRed() {
	fmt.Println("writing through colorFullpen using red ink")
}

type writingDevice interface {
	write()
}

func writer(device writingDevice) { // function parameter with interface type
	device.write()
	//device.writeRed()     // cannot call function writeRed because type writingDevice has no method named writeRed
}

func main() {
	var device writingDevice
	device = pen("pen")
	device.write()
	device = pencil(9)
	device.write()

	var anotherDevice writingDevice
	anotherDevice = pencil(10)
	writer(anotherDevice)

	writer(pen("pen"))

	writer(colorFullPen("red"))
}
