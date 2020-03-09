package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sum(number int) int {
	result := 0
	if number == 0 {
		return 0
	}
	result = sum(number-1) + number
	return result
}

func main() {
	result := 0
	fmt.Print("Input the number : ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	number, _ := strconv.Atoi(line)
	result = sum(number)
	fmt.Printf("Sum number is %d\n", result)
}
