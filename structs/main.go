package structs

import (
	"fmt"
	"math"
)

// Rectangle shape
type Rectangle struct {
	Height float64
	Width  float64
}

// Circle shape
type Circle struct {
	Radius float64
}

// Shape is an interface with area function
type Shape interface {
	area() float64
}

func (r Rectangle) area() float64 {
	return r.Height * r.Width
}

func (c Circle) area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func main() {
	// declare s1/s2 as Shape inteface
	var s1 Shape
	var s2 Shape

	// init r as Rectangle shape
	r := Rectangle{
		12,
		2,
	}
	// using interface
	s1 = r

	// init c as Circle shape
	c := Circle{
		100,
	}
	// using interface
	s2 = c

	// using interface
	fmt.Println(s1.area())
	fmt.Println(s2.area())

	fmt.Println("===============")

	// using method
	fmt.Println(r.area())
	fmt.Println(c.area())

	fmt.Println("===============")

	fmt.Println("main function")
}
