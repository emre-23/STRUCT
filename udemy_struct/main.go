package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	emre := person{
		firstName: "Emre",
		lastName:  "Görmüş",
		contact: contactInfo{
			email:   "ygormus173@gmail.com",
			zipCode: 34906,
		},
	}
	fmt.Printf("%+v", emre)
}
