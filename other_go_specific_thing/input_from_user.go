package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// taking input using bufio package
	fmt.Println("type something")
	readerObj := bufio.NewReader(os.Stdin)
	inp, _ := readerObj.ReadString('\n')
	fmt.Println("input is  : ", inp)

	//taking input using scan method of fmt package
	fmt.Println("type something")
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("input is  : ", number)

}
