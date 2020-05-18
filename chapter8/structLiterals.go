package main

import "fmt"

type stuInfo struct {  // with use of type we can declare variabe of your own type 
	name   string
	rollNo int
	clgName   string
	branch 	string
}

func main() {
	student1:=stuInfo{
		name : "ayushman singh gehlot",
		rollNo : 107,
		clgName : "III Nagpur",
		branch : "CSE",
	}
	fmt.Println(student1)
}
