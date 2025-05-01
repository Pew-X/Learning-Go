package main

import (
	// "encoding/json"
	"fmt"

	"github.com/Pew-X/learning-go/structs_interfaces"
)

func main() {
	// var Suv structs_interfaces.Car
	// Suv.Name = "ABC"
	// Suv.Capacity = 69
	// Suv.Car_type = "SUV"
	// Suv.Owner.Name = "PW"
	// fmt.Printf("Suv: %v\n", Suv)

	// fmt.Println(Suv.MilesLeft(28))

	// // Some json stuff for struct tags

	// r := structs_interfaces.Response{Page: 1, Words: []string{"UP", "AWAY", "IT", "GOES"}}
	// j, _ := json.Marshal(r)

	// fmt.Println(string(j))

	// fmt.Println(r)

	// var r2 structs_interfaces.Response

	// json.Unmarshal(j, &r2)
	// fmt.Println(r2)

	// // pointer in method recievers

	// p := structs_interfaces.Point{X: 15, Y: 22}

	// fmt.Print(&p)

	// p.DistanceFromOrg()

	// p.DistanceFromOrgPtr()

	//Line arithmetic

	side := structs_interfaces.Line{Begin: structs_interfaces.Point{X: 1, Y: 2}, End: structs_interfaces.Point{X: 4, Y: 6}}

	fmt.Println(side.Distance())

	//Path via points distance methods
	Perimeter := structs_interfaces.Path{{X: 1, Y: 1}, {X: 5, Y: 1}, {X: 5, Y: 4}, {X: 1, Y: 1}}

	fmt.Println(Perimeter.Distance())

}
