package main

import "fmt"

func main() {
	a := 5
	b := &a
	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// user to read value from address
	fmt.Println(*&a)

	// change value with pointer
	*b = 10
	fmt.Println(a)

}
