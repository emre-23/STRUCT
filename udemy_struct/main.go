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
	contactInfo
}

func main() {
	emre := person{
		firstName: "Emre",
		lastName:  "Görmüş",
		contactInfo: contactInfo{
			email:   "ygormus173@gmail.com",
			zipCode: 34906,
		},
	}
	emrePointer := &emre
	emrePointer.updateName("Yunus")
	emrePointer.updateLastName("Alkılıç")
	emre.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson *person) updateLastName(newLastName string) {
	(*pointerToPerson).lastName = newLastName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
