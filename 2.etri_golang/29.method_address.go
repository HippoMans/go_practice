package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) HappyNewYear() int {
	p.age++
	return p.age
}

func (p *Person) ChangeName(newName string) string {
	p.name = newName
	return p.name
}

func main() {
	Hippo := Person{"Hippo", 29}
	fmt.Printf("새해 복 많이 받으세요. 새해에 당신의 나이는 %d살입니다.\n", Hippo.HappyNewYear())

	//개명
	var newName string
	fmt.Print("개명하려는 이름을 입려하세요 : ")
	fmt.Scanf("%s", &newName)
	fmt.Printf("%s 이름으로 변경하였습니다.\n", Hippo.ChangeName(newName))
}
