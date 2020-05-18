package main

import "fmt"

type myType string

func (t myType) sayHi() { // receiver parameter has to be user-defined type cannot be of universal type(ex int , string , bool etc)
	// reciver parameter name has isn't important
	// infact it is optional same a return value name
	fmt.Println("HI")
	fmt.Println("HI", t) // prints HI followed by receiver parameter value
}

/*
The type of the receiver parameter is the type that the method becomes associated with. But aside
from that, the receiver parameter doesn’t get special treatment from Go. You can access its contents
within the method block just like you would any other function parameter.
As with any other function, you can define additional parameters within parentheses following the
method name
*/
func main() {
	var temp myType = "from ayushman"
	temp.sayHi()

	value := myType("from another variable value") // declares value of myType and its value as "from another ..."
	value.sayHi()
}
