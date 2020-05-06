package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int8
	address   []string
}

func createPerson(fName, lName string, age int8, add []string) *person {
	var perObj person
	perObj.firstName = fName
	perObj.lastName = lName
	perObj.age = age
	perObj.address = add
	return &perObj
}

func printInfo(perobj *person) {
	fmt.Println(perobj.firstName)
	fmt.Println(perobj.lastName)
	fmt.Println(perobj.age)
	for _, add := range perobj.address {
		fmt.Println(add)
	}
}
func main() {

	person1 := createPerson("pravin N S ", "gehlot", 56, []string{"R-139 Mahalaxmi Nagar", "605 rishab regency"})
	printInfo(person1) // person1 itself is a pointer to struct so no need to put "&" in front of it

	person2 := person{
		firstName: "ayushman",
		lastName:  "gehlot",
		age:       20,
		address: []string{
			"R-139 Mahalaxmi Nagar",
			"605 rishab regency",
		},
	}
	printInfo(&person2)
}
