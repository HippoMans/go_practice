package main

import "fmt"

func main() {

	// map의 key와 value값을 일일이 대입
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrongen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])
	fmt.Println(elements["O"])

	// map의 기본적인 사용
	a := map[string]int{"Hippo": 10000, "Mans": 9999}
	var b *map[string]int
	b = &a
	c := *b
	fmt.Println("a map : ", a["Hippo"])
	fmt.Printf("a adress : %p\n", &a)
	fmt.Printf("b adress : %p\n", b)
	fmt.Println("c map : ", c["Mans"])
	fmt.Printf("c address %p\n", &c)
}
