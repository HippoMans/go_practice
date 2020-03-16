package main

import "fmt"

func main() {
	var what interface{}
	what = 100
	PrintWhat(what)
	what = "Hippo"
	PrintWhat(what)
}

func PrintWhat(v interface{}) {
	fmt.Println(v)
}
