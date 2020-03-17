package main

import (
	"fmt"
	"net"
)

func requestHandler(c net.Conn) {
	data := make([]byte, 4096)
	for {
		n, err := c.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(data[:n]))

		_, err = c.Write(data[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()
		go requestHandler(conn)
	}
}
