package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	var alex person
	alex.firstName = "Emre"
	alex.lastName = "Görmüş"
	fmt.Println(alex)
	fmt.Printf("%+v", alex)
	fmt.Println("Hello")
}
