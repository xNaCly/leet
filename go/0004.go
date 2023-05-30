package main

import (
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
