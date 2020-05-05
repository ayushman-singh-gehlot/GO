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

	value,bool:=ranks["bronze"]
	fmt.Println(value,bool)
	delete(ranks,"bronze")     // delete key-value pair of key="bronze"
	value,bool=ranks["bronze"]
	fmt.Println(value,bool)

}