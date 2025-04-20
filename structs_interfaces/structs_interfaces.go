package structs_interfaces

import "time"

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
