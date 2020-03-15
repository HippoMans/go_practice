package main

import "fmt"

func main() {

	// 기본 for 구문(1)
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	// 기본 for 구문(2)
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	// for-break 구문
	for x := 1; x <= 3; x++ {
		for y := 1; y <= 5; y++ {
			if y == 2 {
				fmt.Println("stop")
				break
			}
		}
		fmt.Println("Good")
	}

	// for-continue 구문
	for a := 0; a <= 2; a++ {
		for b := 0; b <= 2; b++ {
			if a == 1 && b == 1 {
				continue
				fmt.Println("It is not printed")
			}
		}
		fmt.Println("Good")
	}

	// lable + for-break 구문
Loop:
	for t := 0; t <= 2; t++ {
		for v := 0; v <= 2; v++ {
			if v == 1 {
				break Loop
			}
			fmt.Println(t, v)
		}
	}
	fmt.Println("Loop stop")
}
