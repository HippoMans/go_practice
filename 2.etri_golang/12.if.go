package main

import "fmt"

func main() {
	// 기본 if 구문
	i := 6

	if i >= 10 {
		fmt.Println("10 or more")
	} else if i >= 5 && i < 10 {
		fmt.Println("5 or more, less than 10")
	} else {
		fmt.Println("less than 5")
	}

	// for + if 구문
	for x := 1; x <= 10; x++ {
		if x%2 == 0 {
			fmt.Println(x, "is even")
		} else if x%3 == 0 {
			fmt.Println(x, "is triple")
		} else if x%5 == 0 {
			fmt.Println(x, "is quintuple")
		} else {
			fmt.Println("Unknown")
		}
	}
}
