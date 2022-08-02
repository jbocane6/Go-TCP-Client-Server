package main

import (
	"fmt"
	"net"
)

func server() {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c)
	}
}

func handleClient(c net.Conn) {
	b := make([]byte, 100)
	bs, err := c.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Message: ", string(b[:bs]))
		fmt.Println("Bytes: ", bs)
	}

}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
