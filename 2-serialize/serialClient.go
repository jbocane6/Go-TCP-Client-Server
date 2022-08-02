package main

import (
	"encoding/gob"
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
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	go client()

	var input string
	fmt.Scanln(&input)
}
