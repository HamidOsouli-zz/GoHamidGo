package main

import (
	"fmt"
	"math"
)

// declare interface
type Shape interface {
	area() float64
}
type Circle struct {
	x, y, radius float64
}
type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) area() float64 {
	return r.width * r.height
}
func getArea(s Shape) float64 {
	return s.area()
}
func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 12}
	fmt.Printf("circle area is : %f\n", getArea(circle))
	fmt.Printf("rectangle area is : %f\n", getArea(rectangle))
}
