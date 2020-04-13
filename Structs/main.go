package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	// contact   contactDetails // This line can be shortened as below
	contactDetails
}

type contactDetails struct {
	mobile int
	email  string
}

func (p person) printDetails() {
	fmt.Println("First Name:", p.firstName)
	fmt.Println("Last Name:", p.lastName)
	fmt.Println("Email:", p.contactDetails.email)
	fmt.Println("Mobile:", p.contactDetails.mobile)

}

func main() {
	human := person{}
	human.firstName = "Mega"
	human.lastName = "Noob"
	// human.contact.email = "m.noob@noobs.com"
	// human.contact.mobile = 9999999999
	human.contactDetails.email = "test@test.com"
	human.contactDetails.mobile = 1234567890
	// fmt.Println(human)
	// fmt.Printf("%+v", human)

	human.printDetails()
}
