package main

import "fmt"

func main() {
	var num1 byte = 10   // 10 진수
	var num2 byte = 0x32 // 16 진수
	var num3 byte = 'a'  // 문자

	/* byte에 문제 발생
	var num4 byte = "a"    // 문자열은 컴파일 에러 발생
	var num5 byte = 'ab' // 2개의 문자는 컴파일 에러 발생
	var num6 byte = "ab"   // 문자열은 컴파일 에러 발생
	*/
	fmt.Println("10의 byte : ", num1, "\n0x32의 byte : ", num2, "\n문자의 byte : ", num3)
	// fmt.Println("\"a\"의 byte : ", num4, "'ab'의 byte : ", num5, "\"ab\"의 byte : ", num3)
}
