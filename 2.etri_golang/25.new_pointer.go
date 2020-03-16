package main

import "fmt"

func one(x *int) {
	*x = 1
}

func main() {
	x := new(int)
	one(x)
	fmt.Println(*x)
}
