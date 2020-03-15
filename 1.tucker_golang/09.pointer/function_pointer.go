package main

import "fmt"

func Increase(value *int) {
	*value += 1
}

func Decrease(value *int) {
	*value -= 1
}

func main() {
	var a int
	var number int
	a = 1
	fmt.Println("\n*****************")
	fmt.Printf("a value is %d\n\n", a)
	fmt.Print("Increase number : ")
	fmt.Scanf("%d", &number)
	for i := 0; i < number; i++ {
		Increase(&a)
	}
	fmt.Printf("Increase a value is %d\n\n", a)
	fmt.Println("*****************")
	fmt.Print("Decrease number : ")
	fmt.Scanf("%d", &number)
	for i := 0; i < number; i++ {
		Decrease(&a)
	}
	fmt.Printf("Increase a value is %d\n\n", a)
	fmt.Print("*****************\n")
}
