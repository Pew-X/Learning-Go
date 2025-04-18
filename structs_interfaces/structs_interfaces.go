package structs_interfaces

type Car struct {
	Name     string
	Car_type string
	Capacity int
	Owner    Owner
}

type Owner struct {
	Name string
}

// method associated with a struct - has access to struct itself

func (c Car) MilesLeft(milesdone int) int {
	return c.Capacity - milesdone

}
