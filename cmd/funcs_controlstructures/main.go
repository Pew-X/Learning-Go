package main

import (
	"fmt"

	funcscontrolstructures "github.com/Pew-X/learning-go/funcs_controlstructures"
)

func main() {
	funcscontrolstructures.Printme()

	s := "Hello"
	funcscontrolstructures.Printmeparam(s)

	ans := funcscontrolstructures.IntDiv(420, 0)
	fmt.Print(ans)

	ans, rem, err := funcscontrolstructures.IntDiv_mulReturns(420, 69)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Print(ans)
	fmt.Print(rem)

}
