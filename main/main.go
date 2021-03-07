package main

import (
	"fmt"

	"github.com/headfirstgo/geo"
	"github.com/headfirstgo/magazine"
)

func main() {
	subscriber := magazine.Subscriber{Name: "Emre"}
	subscriber.Address.Street = "Hidayet"
	fmt.Println(subscriber.Address.Street)
	fmt.Println(subscriber.Name)

	location := geo.Landmark{}
	location.Name = "The Googleplex"
	location.Latitude = 37.42
	location.Longitude = -122.08
	fmt.Println(location)

}
