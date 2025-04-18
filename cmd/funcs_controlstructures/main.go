package main

import (
	"fmt"

	funcscontrolstructures "github.com/Pew-X/learning-go/funcs_controlstructures"
)

func main() {
	funcscontrolstructures.Printme()

	s := "Hello"
	funcscontrolstructures.Printmeparam(s)

	ans := funcscontrolstructures.IntDiv(420, 69)
	fmt.Print(ans)

	ans, rem := funcscontrolstructures.IntDiv_mulReturns(420, 69)
	fmt.Print(ans)
	fmt.Print(rem)
}
