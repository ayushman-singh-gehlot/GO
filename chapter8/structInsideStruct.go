package main

import "fmt"
import "packages/studentinfo"

func main() {
	var student1 studentinfo.StuInfo
	student1.Name="ayushman singh gehlot"
	student1.RollNo=107
	student1.ClgName="IIIT Nagpur"
	student1.Branch="CSE"

	// adding address using concept of struct inside struct
	student1.HomeAddress.HouseNo="139"
	student1.HomeAddress.Colony="Mahalaxmi Nagar"
	student1.HomeAddress.City="Indore"
	student1.HomeAddress.Pincode=452010

	// adding subjects using concept of embedding structs
	student1.LabSub=[] string{ "compilers", "cryptography", "AI"}
	student1.TheorySub=[] string{ "NNDL","MDS"}

	fmt.Println(student1)

}
