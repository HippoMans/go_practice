package main

import "fmt"

type Student struct {
	name  string
	age   int
	grade int
}

//value type copy
func (t Student) SetName(newName string) {
	fmt.Printf("SetName 내부의 Student의 메모리 주소 : %p\n", &t)
	t.name = newName
}

//reference type copy
func (t *Student) SetPointName(newName string) {
	fmt.Printf("SetPointName 내부의 Student의 메모리 주소 : %p\n", t)
	t.name = newName
}

func (t *Student) SetPointAge(newAge int) {
	fmt.Printf("SetPointAge 내부의 Student의 메모리 주소 : %p\n", t)
	t.age = newAge
}

func (t *Student) PrintStudent() {
	fmt.Printf("t의 메모리 구조 : %p\n", t)
	fmt.Println(t)
}

func main() {
	// a를 인스턴스라고 한다.
	// 객체에 대하여 생명주기를 부여하여 사라질 때까지 주체성을 띄게 하는 대상
	a := Student{"Hippo", 30, 99}
	fmt.Printf("main 내부에서 Student를 인스턴스한 a의 메모리 주소 : %p\n\n", &a)
	a.SetName("Mans")
	fmt.Println(a) //이 때 원하는 값으로 변경되지 않는 것을 확인할 수 있다. 이유는 주소

	a.SetPointName("HippoMans")
	a.PrintStudent()
	fmt.Println(a) //주소의 값이 겹쳐져야 원하는 결과를 얻을 수 있다.

}
