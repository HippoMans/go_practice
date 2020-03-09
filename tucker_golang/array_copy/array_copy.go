package main

import "fmt"

func main() {
	arr := [5]int{}
	fmt.Print("input the array number : ")
	for i := 0; i < len(arr); i++ {
		fmt.Scanf("%d", &arr[i])
	}

	clone := [5]int{}

	for i := 0; i < 5; i++ {
		clone[i] = arr[i]
	}

	//array result
	fmt.Print("array : ")
	fmt.Println(clone)

	fmt.Print("array result : ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", clone[i])
	}
	fmt.Println()
}
