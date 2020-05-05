package main

import "fmt"
import "packages/studentinfo"

func showInfo( stuobj studentinfo.StuInfo){  // argument passed of user-defined type stuInfo
	fmt.Println(stuobj.Name,stuobj.RollNo,stuobj.ClgName,stuobj.Branch)
}

func defaultInfo(name string,rollNo int) studentinfo.StuInfo{ // return type of stuInfo
	var stuobj studentinfo.StuInfo
	stuobj.Name=name
	stuobj.RollNo=rollNo
	stuobj.ClgName="III Nagpur"
	stuobj.Branch="CSE"
	return stuobj
}

func changeBranch(stuobj * studentinfo.StuInfo){ // argument recives pointer of type stuInfo
	stuobj.Branch="ECE" // changing a field with help of pointer
						 // this change will reflect in student1 cause we used pass by pointer method
}
// why I didnot wrote it like this 
// *stuobj.branch="ECE"  which i generally do while modifying any primitive type variable using pointer(value at variable)
//reason-The dot notation to access fields works on struct pointers as well as the structs themselves.
// for detailed reason refer page 342

func main() {
	student1:=defaultInfo("ayushman singh gehlot",107)
	showInfo(student1)
	changeBranch(&student1) // passed address instead of value
	showInfo(student1)
}
