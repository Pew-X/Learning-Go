package main

import (
	"fmt"

	"github.com/Pew-X/learning-go/structs_interfaces"
)

func main() {
	var Suv structs_interfaces.Car
	Suv.Name = "ABC"
	Suv.Capacity = 69
	Suv.Car_type = "SUV"
	Suv.Owner.Name = "PW"
	fmt.Printf("Suv: %v\n", Suv)

	fmt.Println(Suv.MilesLeft(28))

}
