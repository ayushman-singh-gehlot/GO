package main

import "fmt"

type stuInfo struct {  // with use of type we can declare variabe of your own type 
	name   string
	rollNo int
	clgName   string
	branch 	string
}

func showInfo( stuobj stuInfo){  // argument passed of user-defined type stuInfo
	fmt.Println(stuobj.name,stuobj.rollNo,stuobj.clgName,stuobj.branch)
}


func main() {
	var student1 stuInfo  // variable declared of type stuInfo

	student1.name = "ayushman singh gehlot"
	student1.rollNo = 107
	student1.clgName="III Nagpur"
	student1.branch="CSE"
	
	showInfo(student1)
}
