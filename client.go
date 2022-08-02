package main

import (
	"fmt"
	"net"
)

func client() {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	msg := "Hello World"
	fmt.Println(msg)
	c.Write([]byte(msg))
	c.Close()
}

func main() {
	go client()

	var input string
	fmt.Scanln(&input)
}
