package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Input two Number : ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString(' ')
	line = strings.TrimSpace(line)

	n1, _ := strconv.Atoi(line)

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.Atoi(line)

	fmt.Printf("Input number are %d, %d\n", n1, n2)

	fmt.Print("Input operator : ")
	reader = bufio.NewReader(os.Stdin)
	line, _ = reader.ReadString('\n')
	operator := strings.TrimSpace(line)

	if operator == "+" {
		fmt.Printf("%d + %d = %d\n", n1, n2, n1+n2)
	} else if operator == "-" {
		if n1 > n2 {
			fmt.Printf("%d - %d = %d\n", n1, n2, n1-n2)
		} else {
			fmt.Printf("%d - %d = %d\n", n2, n1, n2-n1)
		}
	} else if operator == "*" {
		fmt.Printf("%d * %d = %d\n", n1, n2, n1*n2)
	} else if operator == "/" {
		if n2 == 0 {
			fmt.Println("divide is not possible.")
			return
		} else {
			fmt.Printf("%d / %d = %d\n", n1, n2, n1/n2)
		}
	} else {
		fmt.Println("Operator is not possible")
	}
}
