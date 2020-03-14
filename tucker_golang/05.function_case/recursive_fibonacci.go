package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fibonacci(number int) int {
	result := 0
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 1
	}
	result = fibonacci(number-1) + fibonacci(number-2)
	return result
}

func main() {
	result := 0
	fmt.Print("Input order number : ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	number, _ := strconv.Atoi(line)

	for i := 1; i <= number; i++ {
		result = fibonacci(i)
		fmt.Printf("%d ", result)
	}
	fmt.Println()
}
