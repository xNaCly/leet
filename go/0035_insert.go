package main

import "fmt"

var in = map[int]int{
	5: 2,
	2: 1,
	7: 4,
}

func binarySearch(nums []int, target int, low int, high int) int {
	if low > high {
		return low
	} else {
		mid := (low + high) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			return binarySearch(nums, target, mid+1, high)
		} else {
			return binarySearch(nums, target, low, mid-1)
		}
	}
}

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums)-1)
}

func main() {
	for k, v := range in {
		r := searchInsert([]int{1, 3, 5, 6}, k)
		if v != r {
			fmt.Printf("got %d, wanted %d\n", r, v)
		}
	}
}
