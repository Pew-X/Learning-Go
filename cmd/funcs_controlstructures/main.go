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

	ans, rem, err := funcscontrolstructures.IntDiv_mulReturns(420, 69)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Println(ans)
	fmt.Println(rem)

	// pre decleration of variables — anothe way fo writing if statements
	if answer, remainder, err := funcscontrolstructures.IntDiv_mulReturns(420, 69); err != nil {
		fmt.Print("bro wtf")
	} else {
		fmt.Println(answer)
		fmt.Println(remainder)
	}

	//Defer implementation

	defer funcscontrolstructures.A()
	defer funcscontrolstructures.B()
}
