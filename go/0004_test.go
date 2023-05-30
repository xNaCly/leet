package main

import (
	"fmt"
	"testing"
)

func Test0004(t *testing.T) {
	example1 := []int{1, 3}
	example2 := []int{2}
	fmt.Println(findMedianSortedArrays(example1, example2))

	example3 := []int{1, 2}
	example4 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(example3, example4))
}
