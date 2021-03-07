package main

import (
	"fmt"
)

type car struct {
	name     string
	topSpeed float64
}

type part struct {
	description string
	count       int
}

func main() {
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24

	fmt.Println("Name of the car:", porsche.name)
	fmt.Println("Topspeed:", porsche.topSpeed)
	fmt.Println("Description:", bolts.description)
	fmt.Println("Count:", bolts.count)
}
