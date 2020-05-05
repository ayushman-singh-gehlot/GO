package main

import (
	"fmt"
	"reflect"
	"bufio"
	"os"
)

func main() {
	fmt.Print("Enter a grade : ")
	inputobject:=bufio.NewReader(os.Stdin)
	inp_grade,_:=inputobject.ReadString('\n')  // error ignored using _
	fmt.Println(inp_grade)
	fmt.Println("// here  variable inputobject is of type",reflect.TypeOf(inputobject))
}
