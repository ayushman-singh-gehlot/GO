package main

import (
	"fmt"
)

func main(){
	var ranks map[string] int
	ranks=make(map[string] int)

	ranks["gold"]=1
	ranks["silver"]=2
	ranks["bronze"]=3
	fmt.Println(ranks)


	elements:=make(map[string]string)

	elements["H"]="Hydrogen"
	elements["Li"]="Lithium"
	fmt.Println(elements)
	

	myMap:= map[string] float64{"a":1.2,"b":3.4}

	fmt.Println(myMap)


}