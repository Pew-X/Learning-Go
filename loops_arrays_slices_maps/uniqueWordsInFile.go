package loops_arrays_slices_maps

import (
	"bufio"
	"fmt"
	"os"
	_ "sort"
)

func UniqueWords(input *os.File) int {
	scan := bufio.NewScanner(input)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(len(words), "unique words")
	return len(words)
}
