package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	fmt.Println(reflect.TypeOf(time.Now()))// time.Now() return a value of type time.Time
	var now time.Time = time.Now() // time.Time is datatype 
	year := now.Year()
	fmt.Println(now,"\n",year)

}

//for more info refer page 71