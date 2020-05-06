package main

import "fmt"

type stuInfo struct {
	name   string
	rollNo int
}

type stuEducationInfo struct {
	clgName   string
	branch 	string
}

func main() {
	var student1 stuInfo
	var student2 stuInfo
	var student1Edu stuEducationInfo
	var student2Edu stuEducationInfo
	student1.name = "ayushman singh gehlot"
	student1.rollNo = 107
	student1Edu.clgName="III Nagpur"
	student1Edu.branch="CSE"
	student2.name = "fardin ansari"
	student2.rollNo = 108
	student2Edu.clgName="III Nagpur"
	student2Edu.branch="CSE"
	fmt.Println(student1,student1Edu)
	fmt.Println(student2,student2Edu)

}
