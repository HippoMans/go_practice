package main

import (
	"fmt"
	"os"
)

func main() {
	mybuffer := ""
	arg := os.Args
	if len(arg) == 1 {
		mybuffer += "커맨드라인 인자 수를 입력하세요~~~"
	} else {
		mybuffer += arg[1:]
	}

	io.WriteString(os.Stdout, mybuffer)
	io.WriteString(os.Stdout, "\n")

}
