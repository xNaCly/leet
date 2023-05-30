package main

type NumArray struct {
	items []int
}

func Constructor(nums []int) NumArray {
	return NumArray{items: nums[:]}
}

func (this *NumArray) SumRange(left int, right int) int {
	res := 0
	for i := left; i <= right; i++ {
		res += this.items[i]
	}
	return res
}
