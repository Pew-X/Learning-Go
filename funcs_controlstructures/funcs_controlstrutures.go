package funcscontrolstructures

import (
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

func IntDiv_mulReturns(n int, d int) (int, int) {
	a := n / d
	return a, n % d

}
