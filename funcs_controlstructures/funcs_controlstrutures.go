package funcscontrolstructures

import (
	"errors"
	"fmt"
)

func Printme() {
	fmt.Println("aha thats me")
}

func Printmeparam(printValue string) {
	fmt.Println(printValue)
}

func IntDiv(n int, d int) int {
	a := n / d
	return a

}

func IntDiv_mulReturns(n int, d int) (int, int, error) {
	var err error
	if d == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	a := n / d
	return a, n % d, err

}
