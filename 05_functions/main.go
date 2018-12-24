package main

import "fmt"

func greeting(name string) string {
	return "hello " + name
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("class"))
	fmt.Println(getSum(5, 12))
}
