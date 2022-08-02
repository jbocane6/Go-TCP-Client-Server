package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Person struct {
	Name  string
	Email []string
}

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
	var person Person
	err := gob.NewDecoder(c).Decode(&person)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Message: ", person)
	}

}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
