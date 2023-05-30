package main

import (
	"testing"

	"github.com/xNaCly/leet/go/helper"
)

func Test0035(t *testing.T) {
	var in = map[int]int{
		5: 2,
		2: 1,
		7: 4,
	}
	for k, v := range in {
		r := searchInsert([]int{1, 3, 5, 6}, k)
		helper.AssertEquals(r, v, k)
	}
}
