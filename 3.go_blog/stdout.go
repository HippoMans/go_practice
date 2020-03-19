package main

import (
	"io"
	"os"
)

func main() {
	mybuffer := ""
	arg := os.Args
	count := len(arg)
	if len(arg) == 1 {
		mybuffer += "커맨드라인 인자 수를 입력하세요~~~"
	} else {
		for i := 1; i < count; i++ {
			mybuffer += arg[i]
		}
	}

	io.WriteString(os.Stdout, mybuffer)
	io.WriteString(os.Stdout, "\n")

}
