package main

import (
	"fmt"
	"math"
)

/*
type Dimension interface {
	area() float32
	perimeter() float32
}
*/

type Ariable interface {
	area() float32
}

type Perimeterable interface {
	perimeter() float32
}

type Dimension interface {
	Ariable
	Perimeterable
}

type Rect struct {
	height, width float32
}

func (r Rect) area() float32 {
	return r.width * r.height
}

func (r Rect) perimeter() float32 {
	return 2*r.height + 2*r.width
}

func (r Rect) String() string {
	return fmt.Sprintf("Type = Area, height = %v, width = %v, area = %v, perimeter = %v", r.height, r.width, r.area(), r.perimeter())
}

type Circle struct {
	radius float32
}

func (c Circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float32 {
	return 2 * math.Pi * c.radius
}

/*
type MyFunc func(int, int) int

func (myFunc MyFunc) execute(x int, y int) int {
	return myFunc(x, y)
}
*/

func printArea(x Ariable) {
	//print the area of the given object if implemented
	fmt.Printf("Area = %v\n", x.area())
}

func printPerimeter(p Perimeterable) {
	fmt.Printf("Perimeter = %v\n", p.perimeter())
}

func printDimension(d Dimension) {
	/*
		fmt.Printf("Area = %v\n", d.area())
		fmt.Printf("Perimeter = %v\n", d.perimeter())
	*/
	printArea(d)
	printPerimeter(d)
}

/*
perimeter for circle = 2 * pi * radius
perimeter for rect = 2 * height + 2 * width
*/

type MyShape struct {
}

func main() {
	rect := Rect{width: 10, height: 20}
	fmt.Printf("%v\n", rect)
	//printDimension(rect)
	printArea(rect)
	circle := Circle{radius: 15}
	//printDimension(circle)
	printArea(circle)

	/* var fn MyFunc = func(x int, y int) int {
		result := x + y
		fmt.Printf("fn is invoked with x = %d, y=%d and returning result = %d\n", x, y, result)
		return result
	}
	fn.execute(10, 20) */
}
