package main

import "fmt"

func main() {
	x := 50
	y := 10
	// If Else
	if x < y {
		fmt.Printf("x: %d is lower than y %d", x, y)
	} else {
		fmt.Printf("x: %d is lower than y %d", y, x)
	}
	// else if
	color := "red"
	// if (color == "green") {
	// 	fmt.Printf("color is red")
	// } else if (color == "blue") {
	// 	fmt.Printf("color is blue")
	// } else {
	// 	fmt.Printf("color is not red or blue either")
	//}
	switch color {
	case "red":
		fmt.Printf("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not red or blue")
	}
}
