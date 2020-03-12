package main

import "fmt"

func main() {
	var a int
	var p *int

	p = &a
	a = 3

	fmt.Println("print a value : ", a)
	fmt.Println("print a address : ", &a)
	fmt.Println("print p value is a address : ", p)
	fmt.Println("print p value is p point a : ", *p)
	fmt.Println("print p address : ", &p)
}
