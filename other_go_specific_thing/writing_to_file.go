package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// 1st method
	data := []byte("Hello from go")
	err := ioutil.WriteFile("file1.txt", data, 0644) // read about the permission
	checkError(err)
	fmt.Println("Done writing !")

	// 2nd method
	f, err := os.Create("file1.txt")
	defer f.Close()
	checkError(err)
	bytesWritten, err := f.WriteString("Hello ayushman \ni m writing inside file")
	checkError(err)
	fmt.Println("Bytes written are : ", bytesWritten)
	fmt.Println("Done writing !")

}
