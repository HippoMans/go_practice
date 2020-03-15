package main

import "fmt"

func main() {
	// 5개의 int 타입으로 구성된 배열
	var x [5]int
	x[4] = 100
	fmt.Println("x : ", x)

	//배열의 값을 모두 더하여 평균구하는 방법
	var y [5]float64
	y[0] = 98
	y[1] = 87
	y[2] = 93
	y[3] = 90
	y[4] = 94

	var total float64 = 0
	for i := 0; i <= 4; i++ {
		total += y[i]
	}
	fmt.Println("배열의 총합 : ", total)
	fmt.Println("집합의 평균 : ", total/5)

	//간단하게 배열 생성한 후 모든 배열출력
	z := [5]int{3, 6, 9, 13, 16}
	fmt.Println("z의 배열 : ", z)
}
