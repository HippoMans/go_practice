package main

import "fmt"

type Person struct {
	id         int
	name       string
	university University
	academy    Academy
}

type University struct {
	name    string
	address string
	grade   string
}

type Academy struct {
	name    string
	address string
	subject string
	price   int
}

func (p *Person) Input_University(name string, address string, grade string) {
	p.university.name = name
	p.university.address = address
	p.university.grade = grade
}

func (p *Person) Input_Academy(name string, address string, subject string, price int) {
	p.academy.name = name
	p.academy.address = address
	p.academy.subject = subject
	p.academy.price = price
}

func (p Person) View_Person() {
	fmt.Println("\n**********************")
	fmt.Printf("person name : %s\n", p.name)
	fmt.Printf("person university name : %s\n", p.university.name)
	fmt.Printf("person university address : %s\n", p.university.address)
	fmt.Printf("person university grade : %s\n\n", p.university.grade)
	fmt.Printf("person academy name : %s\n", p.academy.name)
	fmt.Printf("person academy address : %s\n", p.academy.address)
	fmt.Printf("person academy subject : %s\n", p.academy.subject)
	fmt.Printf("person academy price : %d\n", p.academy.price)
	fmt.Println("**********************\n")
}

func main() {
	mans := Person{id: 1000, name: "mans"}
	mans.Input_University("Pukyong", "Busan", "B+")
	mans.Input_Academy("Beijing Chinese Academy", "Busan", "Chinese speak", 200000)
	mans.View_Person()

	Hippo := Person{id: 1, name: "Hippo"}
	Hippo.Input_University("Korea", "Seoul", "A+")
	Hippo.Input_Academy("ETRI", "Seoul", "blockchain", 1000000)
	Hippo.View_Person()
}
