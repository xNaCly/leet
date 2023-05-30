package main

import (
	"testing"

	"github.com/xNaCly/leet/go/helper"
)

func Test0303(t *testing.T) {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	tests := map[int][]int{
		1:  {0, 2},
		-1: {2, 5},
		-3: {0, 5},
	}
	for r, ks := range tests {
		helper.AssertEquals(obj.SumRange(ks[0], ks[1]), r, ks)
	}
}
