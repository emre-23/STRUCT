package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	city      string
	age       int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Selam, benim adım " + p.firstName + " Yaşadığım şehir " + p.city + " ve " + strconv.Itoa(p.age) + " yaşındayım"
}

//hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// changeCity (pointer receiver)
func (p *Person) changeCity(c string) {
	if p.city == "Ardahan" {
		return
	} else {
		p.city = c
	}
}

func main() {

	person1 := Person{"Emre", "İstanbul", 26}
	person1.hasBirthday()
	person1.changeCity("Kars")
	fmt.Println(person1.greet())
	// fmt.Println(Person(person1))
	// person1.age++
	// fmt.Println(person1.age)
}
