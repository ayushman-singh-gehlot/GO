package main

import (
	"chapter4/greet"
	"chapter4/greet/deutsch"			// very important u need to provide path from src dir src->temp1->temp2->greet
										// temp1/temp2/greet
)
func main(){
	greet.Hello()
	greet.Hi()
	deutsch.Hallo()
	deutsch.GutenTag()
}