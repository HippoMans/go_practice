package main

import "fmt"

/*
	1. 1에서 100까지 출력
	2. 3의 배수는 Fizz 출력
	3. 5의 배수는 Buzz 출력
	4. 3과 5의 공배수는 FizzBuzz 출력
*/

func main() {
	fmt.Println("*******************\nFizzBuzz Game\n")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
