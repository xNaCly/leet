package main

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
