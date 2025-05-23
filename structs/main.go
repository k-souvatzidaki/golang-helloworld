package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zipcode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (p person) updateFirstName(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) updateFirstNamePointer(newFirstName string) {
	(*p).firstName = newFirstName
}


func main() {

	p1 := person{"Person", "One", contactInfo{"person1@gmail.com", 0000}}
	p2 := person{firstName: "Person", lastName: "Two",contact: contactInfo{"person2@gmail.com", 0000} }

	var p3 person
	p3.firstName = "Person"
	p3.lastName = "Three"
	p3.contact = contactInfo{"person3@gmail.com", 0000}

	p1.print()
	p2.print()
	p3.print()
	p3.updateFirstName("PersonNew3") // pass by value
	p3.print()

	(&p3).updateFirstNamePointer("PersonNew3")
	p3.print()
}