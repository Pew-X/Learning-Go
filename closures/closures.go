package closures

/*
In other languages one had to explicitly mention memory in heap if the
memory had to surrive  the scope lifetime of a variable

like in c , c++ one had to allocate using new/malloc on heap so that it stays on memory
after scope is done

In Go it tries to alloc everything on stack unless it needs to
allocate anything on heap

like for example if an address of a variable ( from that function scope)
is retuned at the end of that fucntion — sayint that the variable ( and it's address must live on after  the scope)

thus it is allocated in heap

*/

import "fmt"

func Fib() func() int {
	a, b := 0, 1
	/*

	  A fucntion returning
	  a fuction
	*/

	return func() int {
		a, b = b, a+b //holds onto these as long as f in closures/main.go exists
		return b
	}
}

/*
After fib is done  in line  f, g := fib(), fib()
We think the value and reference of a and b is gone
but in reality as  the anonymus function captures
the value of a and b " closes upon a and b" their values will
live upon as long as f and g are invoked and done.

"As long as there is reference to inner/closure function"
*/

func ClosureFib() {
	f, g := Fib(), Fib()
	/* here each pair of f and g will get a new a & b for themselves each.
	Because two calls to Fib are made so soemwhere out in the heap two distinct a(s) and b(s)
	are there
	*/

	fmt.Println(f(), f(), f(), f(), f(), f(), f())
	fmt.Println(g(), g(), g(), g(), g(), g(), g())
}
