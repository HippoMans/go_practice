package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num1 int8 = 1
	var num2 int16 = 1
	var num3 int32 = 1
	var num4 int64 = 1

	fmt.Println("int8의 변수 크기 : ", unsafe.Sizeof(num1))  // 1 byte
	fmt.Println("int16의 변수 크기 : ", unsafe.Sizeof(num2)) // 2 byte
	fmt.Println("int32의 변수 크기 : ", unsafe.Sizeof(num3)) // 3 byte
	fmt.Println("int64의 변수 크기 : ", unsafe.Sizeof(num4)) // 4 byte

}
