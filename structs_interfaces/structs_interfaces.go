package structs_interfaces

import (
	"fmt"
	"math"
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

type Line struct {
	Begin, End Point
}

func (l Line) Distance() float64 {
	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
}

// let us create a type path that can have points

type Path []Point

func (p Path) Distance() (sum float64) {

	for i := 1; i < len(p); i++ {

		sum += Line{p[i-1], p[i]}.Distance() // we are just creating a line on the fly on two points and just storing distance away

	}

	return sum
}
