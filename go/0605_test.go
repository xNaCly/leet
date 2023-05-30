package main

import (
	"testing"

	"github.com/xNaCly/leet/go/helper"
)

func Test0605(t *testing.T) {
	input := [][]int{
		{0, 0, 1, 0, 0},
		{1, 0, 0, 0, 0, 1},
		{0, 0, 1, 0, 1},
		{1, 0, 0, 0, 1},
		{1, 0, 0, 0, 1},
	}
	inputN := []int{1, 2, 1, 1, 2}
	output := []bool{true, false, true, true, false}
	for i, v := range input {
		helper.AssertEquals(canPlaceFlowers(v, inputN[i]), output[i], v)
	}
}
