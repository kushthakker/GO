package main

import "fmt"

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}

func main() {
	// alex := person{firstName: "John", lastName: "Lennon"}
	// fmt.Println(alex)
	var alex person
	fmt.Printf("%+v", alex)
	alex.firstName = "Alex"
	alex.lastName = "Lennon"
	fmt.Printf("%+v", alex)
	jim := person{
		firstName: "John",
		lastName: "Lennon",
		contact : contactInfo{
			email : "kush@gmailcom",
			zipCode: 560090,
		},
	}
	fmt.Printf("%+v", jim)
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy")
	fmt.Printf("%+v", jim, "---updated---")
}