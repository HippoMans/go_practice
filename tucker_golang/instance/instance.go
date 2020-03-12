package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

func main() {
	fmt.Println("************\nbasic struct")
	a := Student{"Hippo", 30, 100}
	fmt.Println(a, "\n")
	fmt.Println("value copy")
	b := a
	fmt.Println(b, "\n")
	fmt.Println("reference copy")
	var c *Student
	c = &a
	fmt.Println(c, "\n")
	b.age = 10
	c.grade = 95

	//a는 구조체 변수 => 주소값을 얻기 위해서는 &a를 사용해야 한다.
	fmt.Printf("구조체 a의 메모리 주소 : %p  ", &a)
	fmt.Println("구조체a의 값 : ", a, "\n")

	//b는 구조체 변수 => 주소값을 얻기 위해서는 &b를 사용해야 한다.
	//b는 a 구조체의 value copy를 한 것으로  새로운 메모리 공간을 만들어서 복사한 것이다.
	//b는 a에 영향을 주지도 받지도 않는 순수 독립적인 구조체이다.
	fmt.Printf("구조체 b의 메모리 주소 : %p  ", &b)
	fmt.Println("구조체b의 값 :", b, "\n")

	//c 구조체 주소값 변수 (c는 Student의 구조체의 주소값이 들어가는 변수이다.)
	//c는 a의 주소값을 받기에 reference copy를 한 것이다.
	//a의 구조체를 c가 가리키고 있음으로 a에 영향을 주고 받는다.
	fmt.Printf("구조체 c의 메모리 주소 : %p  ", c)
	fmt.Println("구조체 c의 값 : ", *c, "\n")
	fmt.Println("구조체 c의 값 : ", c, "\n")
}
