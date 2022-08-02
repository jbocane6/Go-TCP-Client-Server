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

func client(person Person) {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(person)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	person := Person{
		Name: "Juan Camilo",
		Email: []string{
			"jbocane6@gmail.com",
			"3311@holbertonstudents.com",
		},
	}
	go client(person)

	var input string
	fmt.Scanln(&input)
}
