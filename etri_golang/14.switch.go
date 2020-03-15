package main

import "fmt"

func main() {
	i := 5

	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("Unknown")
	}
}
