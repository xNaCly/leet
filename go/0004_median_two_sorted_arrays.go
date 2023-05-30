package main

import (
	"fmt"
	"math"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedArr := make([]int, 0)
	mergedArr = append(mergedArr, nums1...)
	mergedArr = append(mergedArr, nums2...)
	lenght := len(mergedArr)

	sort.Ints(mergedArr)

	if len(mergedArr)%2 != 0 {
		medianIndex := math.Floor(float64(lenght) / 2.0)
		return float64(mergedArr[int(medianIndex)])
	}

	medianIndexLeft := math.Floor(float64(lenght / 2.0))
	medianIndexRight := medianIndexLeft - 1
	return float64((mergedArr[int(medianIndexLeft)] + mergedArr[int(medianIndexRight)])) / 2.0
}

func main() {
	example1 := []int{1, 3}
	example2 := []int{2}
	fmt.Println(findMedianSortedArrays(example1, example2))

	example3 := []int{1, 2}
	example4 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(example3, example4))
}
