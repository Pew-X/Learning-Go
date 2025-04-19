package main

import (
	"fmt"
	"os"

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

	for context, val := range myMap {
		fmt.Println(context)
		fmt.Println(val)
	}

	for i, v := range arr {
		fmt.Println(i)
		fmt.Println("->")
		fmt.Println(v)
	}

	for i, v := range slice {
		fmt.Println(i)
		fmt.Println("->")
		fmt.Println(v)
	}

	//While Loop in go

	for {
		fmt.Print("Hello")
		break // Remove for inf loop or use conditionally
	}

	//Normal For Loop
	for i := 0; i < 7; i++ {
		fmt.Println(slice[i])
	}

	/// ALT for loop using range keyword
	for i := range 7 {
		fmt.Println(slice[i])
	}

	// Number of unique words in .txt file
	file, err := os.Open("cmd/loops_arrays_slices_maps/tokens.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Println(loops_arrays_slices_maps.UniqueWords(file))

}
