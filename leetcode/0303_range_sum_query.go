package main

import "log"

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

func main() {
// [-2, 0, 3, -5, 2, -1]
    obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
    log.Println(obj.SumRange(0, 2))
    log.Println(obj.SumRange(2, 5))
    log.Println(obj.SumRange(0, 5))
}
