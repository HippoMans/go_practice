package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var file *os.File
	file = os.Stdin

	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		fmt.Println(">", scan.Text())
	}
}
