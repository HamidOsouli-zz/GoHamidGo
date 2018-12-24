package main

import "fmt"

var name = "test"

func main() {
	//var name = "hamid"
	var age int32 = 37
	var isCool = true
	isCool = false
	// different declare of variable name
	name := "hamid"
	name, email := "hamid", "google.com"
	var size float32 = 1.32
	fmt.Println(name, age, isCool, size)
	fmt.Printf("%T\n%T", name, email)
}
