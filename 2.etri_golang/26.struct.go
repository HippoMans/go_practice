package main

import "fmt"

type Rectangle struct {
	x uint64
	y uint64
}

//사각형 둘레
func RectanglePerimeter(r *Rectangle) uint64 {
	return (2 * r.x) + (2 * r.y)
}

//사각형 넓이
func RectangleArea(r *Rectangle) uint64 {
	return r.x * r.y
}

func main() {
	rec := Rectangle{x: 10, y: 20}
	fmt.Println("사각형의 둘레 : ", RectanglePerimeter(&rec))
	fmt.Println("사각형의 넓이 : ", RectangleArea(&rec))
}
