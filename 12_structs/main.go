package main

import (
	"fmt"
	"strconv"
)

// Person define
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + 
	strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMerried method (pointer reviever)
func (p *Person) getMerried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {

	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age:25}
	person2 := Person{"Bob", "Johnson", "New York", "m", 30}

	// fmt.Printf("%T\n", person1.age)
	// person1.age++
	// fmt.Println(person1.age)
	person1.hasBirthday()
	person1.getMerried("Williams")
	person2.getMerried("Thompson")

	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}