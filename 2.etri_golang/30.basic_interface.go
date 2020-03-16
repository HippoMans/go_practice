package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

//Rect 타입에 대한 Shape 인터페이스 구현
func (r Rect) area() float64      { return r.width * r.height }
func (r Rect) perimeter() float64 { return 2 * (r.width + r.height) }

//Circle 타입에 대한 Shape 인터페이스 구현
func (c Circle) area() float64      { return math.Pi * c.radius * c.radius }
func (c Circle) perimeter() float64 { return 2 * math.Pi * c.radius }

func main() {
	r := Rect{width: 10.3, height: 4.2}
	c := Circle{radius: 3.3}
	showPerimeter(r, c)
}

func showPerimeter(shapes ...Shape) {
	for _, s := range shapes {
		function := s.perimeter()
		fmt.Println(function)
	}
}
