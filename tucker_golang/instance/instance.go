package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func main() {
	a := Student{"Hippo", 30, 100}

	fmt.Println("value copy")
	b := a

	fmt.Println()
	fmt.Println("reference copy")
	var c *Student
	c = &a

	b.age = 10
	c.gradei = 95
}
