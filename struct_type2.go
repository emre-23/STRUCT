package main

import (
	"fmt"
)

type subcscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	var subcscriber1 subcscriber
	subcscriber1.name = "Emre"
	fmt.Println("Name:", subcscriber1.name)
}
