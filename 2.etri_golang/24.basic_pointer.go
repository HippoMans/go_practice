package main

import "fmt"

func zero(x int) {
	x = 0
}

func ten(x *int) {
	*x = 10
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x)
	ten(&x)
	fmt.Println(x)
}
