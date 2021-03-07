package main

import (
	"fmt"
)

type student struct {
	name  string
	grade float64
}

func printInfo(s student) {

	fmt.Println("Name:", s.name)
	fmt.Printf("Grade: %0.1f\n", s.grade)
}
func main() {
	var s student
	s.name = "Emre"
	s.grade = 92.3
	printInfo(s)
}
