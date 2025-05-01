package structs_interfaces

import (
	"fmt"
	"time"
)

type Car struct {
	Name     string
	Car_type string
	Capacity int
	Owner    Owner
}

type Owner struct {
	Name string
}

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

type Response struct {
	Page  int      `json:"page"`
	Words []string `json:"words,omitempty"`
}

// method associated with a struct - has access to struct itself

func (c Car) MilesLeft(milesdone int) int {
	return c.Capacity - milesdone

}

type Point struct {
	X float64
	Y float64
}

//method with pointer reciever

func (p Point) DistanceFromOrg() float64 {
	fmt.Print(&p)

	return p.X*p.X + p.Y*p.Y
}

func (p *Point) DistanceFromOrgPtr() float64 {
	fmt.Print(&p)
	return p.X*p.X + p.Y*p.Y
}
