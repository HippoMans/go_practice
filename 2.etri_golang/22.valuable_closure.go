package main

import "fmt"

func makeEven() func() uint {
	i := uint(0)
	return func() (value uint) {
		value = i
		i += 2
		return value
	}
}

func outadd(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func main() {
	fmt.Println("makeeven함수")
	even := makeEven()
	fmt.Println(even())
	fmt.Println(even())
	fmt.Println(even())

	fmt.Println("outadd함수")
	add := outadd(100)
	fmt.Println(add(5))
	fmt.Println(add(200))
	fmt.Println(add(300))
}
