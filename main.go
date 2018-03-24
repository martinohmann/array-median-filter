package main

import (
	"array-median-filter/pkg/finder"
	"fmt"
)

func main() {
	example := [][]int{[]int{1}, []int{1, 2, 3}}

	fmt.Printf("missing length is: %d\n", finder.FindTheMissingLength(example))
}
