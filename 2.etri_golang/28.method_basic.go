package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	height uint64
	width  uint64
}

func (c *Circle) circleArea() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) rectangleArea() uint64 {
	return r.height * r.width
}

func (r *Rectangle) rectanglePerimeter() uint64 {
	return (r.height * 2) + (r.width * 2)
}

func main() {
	// circle 구조체 사용
	circle := Circle{0, 0, 10}
	fmt.Println("circle area is ", circle.circleArea())

	// rectangle 구조체 사용
	rectangle := Rectangle{20, 30}
	fmt.Println("rectangle area is ", rectangle.rectangleArea())
	fmt.Println("rectangle perimeter is ", rectangle.rectanglePerimeter())
}
