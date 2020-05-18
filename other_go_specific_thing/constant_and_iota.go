package main

import "fmt"

type byteSize float64

const (
	_           = iota
	KB byteSize = 1 << (10 * iota) // left shift operation search on internet
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Println(KB, MB, GB)
	var fileSize byteSize = 50000000.0
	fmt.Printf("file size %0.2fKB", fileSize/KB)
}
