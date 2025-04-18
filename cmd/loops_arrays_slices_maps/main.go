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

	//Slices

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

}
