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
	// reading complete file
	content, err := ioutil.ReadFile("file.txt") // complete path after src if file is not in same folder
	checkError(err)
	result := string(content)
	fmt.Println("content of the file is : ", result)

	// reading first few bytes from the file
	f, er := os.Open("file.txt")
	defer f.Close()
	checkError(er)
	bucket := make([]byte, 100)
	bytesRead, er := f.Read(bucket) // bucket reads first 100 bytes
	checkError(er)
	fmt.Println("content of the file(limited) is : ", string(bucket[:bytesRead]))
	fmt.Println("Bytes Read : ", bytesRead)
}

// refer to video lec 45 for more info
