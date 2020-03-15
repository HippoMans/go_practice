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

func main() {
	fmt.Println("\n*****find the mistery number game*****")
	fmt.Println("*****you have ten chance to find the number*****\n\n")

	//random value
	source := rand.NewSource(time.Now().UnixNano()) // seed(source) is made by UnixNano time
	random := rand.New(source)                      // random(value) is made by seed
	mistery := random.Intn(100)                     // mistery is from 0 to 100 (only int)

	for i := 0; i < 10; i++ {
		fmt.Print("Input the number : ")
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		number, _ := strconv.Atoi(line)

		//switch case is possible to use operator
		switch {
		case mistery > number:
			fmt.Println("Input number is less than mistery number")
		case mistery < number:
			fmt.Println("Input number is greater than mistery number")
		case mistery == number:
			fmt.Println("Congratulation~~~!!!\n you success to find the nubmer\n you win the game~~~!!")
			return
		}
	}
	fmt.Println("you failed to find number... Sorry")
}
