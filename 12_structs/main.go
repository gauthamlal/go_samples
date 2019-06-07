package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (Value receiver)
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old"
}

// hasBirthday method (Pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (Pointer receiver)
func (p *Person) getMarried(spouseLastname string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastname
	}
}

func main() {
	// Init person using struct

	person1 := Person{firstName: "Jon", lastName: "Moxley", city: "Vegas", gender: "m", age: 33}

	person2 := Person{firstName: "Renee", lastName: "Young", city: "Toronto", gender: "f", age: 30}

	// Alternative
	// person2 := Person{"Cody", "Rhodes", "Atlanta", "m", 33}

	fmt.Println(person1)
	// fmt.Println(person2)

	// fmt.Println(person1.firstName)
	// person1.age++
	// fmt.Println(person1)

	person1.hasBirthday()
	person1.getMarried("Grey")
	fmt.Println(person1.greet())
	fmt.Println(person2)
	person2.getMarried("Moxley")
	fmt.Println(person2.greet())

}
