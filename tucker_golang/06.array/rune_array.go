package main

import "fmt"

func main() {
	s := "Hello 월드"

	//ASCII code result
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], ", ")
	}
	fmt.Println()

	//korean character print ASCII code ==> ERROR
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c, ", s[i])
	}
	fmt.Println()

	//convert string to []rune ==> rune is from 1 byte to 3 byte
	rune_string := []rune(s)

	//rune code result
	for i := 0; i < len(rune_string); i++ {
		fmt.Print(rune_string[i], ", ")
	}
	fmt.Println()

	// rune character
	for i := 0; i < len(rune_string); i++ {
		fmt.Printf("%c, ", rune_string[i])
	}
	fmt.Println()
}
