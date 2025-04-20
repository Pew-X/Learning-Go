package funcscontrolstructures

import (
	"errors"
	"fmt"
)

/*
Parameter passing by value :
numbers
bool
arrays
structs


By reference
things passed by pointer &x
strings ( but strings are immutable)

slices - sent as reference always
maps - sent by reference

(A MAP AND A SLICE IS ALWAYS A REFERENCCE TOUNDERLYING DATA - DESCRIPTOR)

channels
*/

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

func A() {
	fmt.Println("A invoked")
	B() // fun thing to have infinite loop
}

func B() {
	fmt.Println("B invoked")
	A() // fun thing to have infinite loop
}

/*
defer keyword means HERE IS SOMETHING I NEED TO DO WHEN
I LEAVE THE FUCNTION NO MATTER WHAT.
 for ex :
 - close a file we opened
 - close a socket.HTTP req we made
 - unlock a mutex we lcoekd
 - make sure something is saved before we are  done


 Multiple Defers in one function will be done in opposite manner

 ie fucn main() {
 defer a()
 dever b()  will happen as b() -> a()

 Thus defer is a stack sort of LIFO thing

 }
*/
