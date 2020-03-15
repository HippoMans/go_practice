package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func change_right_number() int {
	for {
		source := rand.NewSource(time.Now().UnixNano())
		random := rand.New(source)
		number := random.Intn(10)

		if number >= 2 && number <= 9 {
			dan := number
			return dan
		}
	}
}

func main() {
	fmt.Print("Input the dan : ")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	dan, _ := strconv.Atoi(line)

	if dan < 2 || dan > 9 {
		dan = change_right_number()
	}

	fmt.Printf("\ndan : %d\n", dan)
	for i := 1; i < 10; i++ {
		fmt.Printf("%d * %d = %d\n", dan, i, dan*i)
	}
}
