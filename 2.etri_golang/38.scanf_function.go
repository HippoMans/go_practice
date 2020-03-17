package main

import "fmt"

func main() {
	fmt.Println("scan() 함수 사용")
	var string1, string2 string
	num1, _ := fmt.Scan(&string1, &string2)
	//입력한 개수가 num1에 들어간다.
	fmt.Println("Input Number : ", num1)
	fmt.Println(string1, string2)

	fmt.Println("\nscanln() 함수 사용")
	var stringstring1, stringstring2 string
	string_num, _ := fmt.Scanln(&stringstring1, &stringstring2)
	fmt.Println("Input Number : ", string_num)
	fmt.Println(stringstring1, stringstring2)

	fmt.Println("\nscanf() 함수 사용")
	var s1, s2 string
	n1, _ := fmt.Scanf("%s %s", &s1, &s2)
	fmt.Println("Input Number : ", n1)
	fmt.Println(s1, s2)
}
