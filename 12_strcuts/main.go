package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	//firstName string
	//lastName string
	//city string
	//gender string
	//age int
	firstName, lastName, gender, city string
	age                               int
}

// greeting method (Value reciever)
func (p Person) greet() string {
	return "hello my name is : " + p.firstName + " " + p.lastName + " and i am  " + strconv.Itoa(p.age)
}

// checkBirthday method (pointer reciever)
func (p *Person) checkBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseName string) {
	if p.gender == "Male" {
		return
	} else {
		p.lastName = spouseName
	}
}

func main() {
	// init X love using struct
	//love := Person {
	//	firstName: "Sevda",
	//	lastName: "Ebadi",
	//	city: "Tehran",
	//	gender: "Female",
	//	age: 21,
	//}
	// alternative way
	love := Person{
		"Sevda",
		"Ebadi",
		"Tehran",
		"Female",
		21,
	}
	friend := Person{
		"Morteza",
		"Najafi",
		"Quazvin",
		"Male",
		21,
	}
	fmt.Println(love)
	// get single field
	//love.age++
	//fmt.Println(love.firstName)
	//fmt.Println(love.age)
	love.checkBirthday()
	love.getMarried("Osouli")
	friend.getMarried("Ahmadi")
	fmt.Println(love.greet())
	fmt.Println(friend.greet())
}
