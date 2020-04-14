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

func (p *person) updateFirstName(firstName string) {
	(*p).firstName = firstName
}

func main() {
	human := person{}
	human.firstName = "Mega"
	human.lastName = "Noob"
	// human.contact.email = "m.noob@noobs.com"
	// human.contact.mobile = 9999999999
	human.contactDetails.email = "test@test.com"
	human.contactDetails.mobile = 1234567890

	// humanPtr := &human
	// humanPtr.updateFirstName("Mr.")
	// The above method is the complete way of using pointers and sending a referance to a receiver
	// The below version works the exact same way as Go helps you use it as a shortcut

	human.updateFirstName("Mr.")
	human.printDetails()
}

//=============================================================================================
//																	Maps
//=============================================================================================

type hashMap map[int]string

func (m hashMap) print() {
	for key, value := range m {
		fmt.Println(key, ":", value)
	}
}

func main2() { //Rename to main when needed to execute

	myMap := hashMap{}
	//	myMap := make(map[int]string)
	myMap[1] = "One"
	myMap[2] = myMap[1] + "Two"
	myMap.print()
}
