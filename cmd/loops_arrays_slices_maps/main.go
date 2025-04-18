package main

import (
	"fmt"

	"github.com/Pew-X/learning-go/loops_arrays_slices_maps"
)

func main() {
	arr, seed := loops_arrays_slices_maps.Arrays(5)
	arr[0] = seed
	fmt.Println(arr[0])
	fmt.Println(seed)
	fmt.Println(&arr[0])

	addr := &arr[0]

	//Slices -------------------------------------

	slice, seed := loops_arrays_slices_maps.Slice(7)
	slice[0] = seed
	slice[1] = *addr

	fmt.Println(slice[0])
	fmt.Println(seed)
	fmt.Println(&slice[0])

	// second element of size will be the the elemet that is at addr of arr[0]
	fmt.Println(slice[1])
	fmt.Println(addr)
	fmt.Println(&slice[1])

	new_slice := append(slice, 4)
	fmt.Println(new_slice[1])
	fmt.Println(addr)

	//Maps {k : v} ---------------------------------------------------------

	var myMap map[string]int32 = loops_arrays_slices_maps.Map()
	// Or if you donno data type use :=

	myMap["piyush"] = 22
	myMap["year"] = 02
	fmt.Println(len(myMap))

}
