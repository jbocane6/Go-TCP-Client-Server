package main

import (
	"fmt"
	"net"
)

func client() {
	net.Dial("tcp", ":9999")
}

func main() {
	go client()

	var input string
	fmt.Scanln(&input)
}
