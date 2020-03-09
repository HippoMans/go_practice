package main

import "fmt"

func main() {
	s := "Hello World"

	//ASCII code result
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], ", ")
	}

	//character result
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c, ", s[i])
	}
	fmt.Println()
}
