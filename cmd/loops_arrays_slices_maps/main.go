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

}
