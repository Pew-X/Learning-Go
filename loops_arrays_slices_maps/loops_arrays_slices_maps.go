package loops_arrays_slices_maps

// Arrays - Fixed size , same type , indexable and contiguous memory

func Arrays(seed int32) ([5]int32, int32) {
	var arr [5]int32
	return arr, seed

}

func Slice(seed int32) ([]int32, int32) {
	arr := make([]int32, seed)
	return arr, seed
}

func Map() map[string]int32 {
	var myMap = map[string]int32{}
	return myMap

}
