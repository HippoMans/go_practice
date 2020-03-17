package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {
	client, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()
	go func(c net.Conn) {
		data := make([]byte, 4096)

		for {
			n, err := c.Read(data)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(data[:n]))
			time.Sleep(2 * time.Second)
		}
	}(client)

	go func(c net.Conn) {
		i := 0
		for {
			s := "Golang " + strconv.Itoa(i)
			_, err := c.Write([]byte(s))

			if err != nil {
				fmt.Println(err)
				return
			}
			i++
			time.Sleep(2 * time.Second)
		}
	}(client)
	fmt.Scanln()
}
