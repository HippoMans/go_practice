package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	//Hippo 사람
	Hippo := Person{}
	Hippo.name = "Hippo"
	Hippo.age = 30
	fmt.Println(Hippo)

	//Mans 사람
	Mans := Person{name: "Mans", age: 28}
	fmt.Println(Mans)

	//Happy 사람
	var Happy Person
	Happy = Person{name: "Happy", age: 20}
	fmt.Println(Happy)

	//Smile 사람 -> new를 사용를 사용하면 포인터를 선언하는 것과 같은 효과
	Smile := new(Person)
	Smile.name = "Smile"
	Smile.age = 99
	fmt.Println(*Smile)

	//Good 사람 -> 구조체 포인터 선언
	Good := &Person{}
	Good.name = "Good"
	Good.age = 10
	fmt.Println(*Good)
}
