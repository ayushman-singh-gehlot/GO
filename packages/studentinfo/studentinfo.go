package studentinfo


type StuInfo struct {  // need to capitalize stuinfo to StuInfo
	Name   string      // rememeber need to capitalize all filed name also in order to export
	RollNo int
	ClgName   string
	Branch 	string
	HomeAddress Address  // used HomeAddress as type Address(type struct)
	Subject              // embedding struct using concept of anonymous struct 
}

type Address struct{
	HouseNo string
	Colony  string
	City    string
	Pincode int
}

type Subject struct{
	LabSub [] string
	TheorySub [] string
}