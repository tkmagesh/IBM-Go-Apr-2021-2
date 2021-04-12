package main

import (
	"fmt"
	"math"
)

type Dimension interface {
	area() float32
}

type Rect struct {
	height, width float32
}

func (r Rect) area() float32 {
	return r.width * r.height
}

type Circle struct {
	radius float32
}

func (c Circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func printDimension(d Dimension) {
	fmt.Printf("Area = %v\n", d.area())
	fmt.Printf("Perimeter = %v\n", d.perimeter())
}

/*
perimeter for circle = 2 * pi * radius
perimeter for rect = 2 * height + 2 * width
*/

type MyShape struct {
}

func main() {
	rect := Rect{width: 10, height: 20}
	printDimension(rect)

	circle := Circle{radius: 15}
	printDimension(circle)
}
