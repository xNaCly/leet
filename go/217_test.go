package main

import (
	"testing"

	"github.com/xNaCly/leet/go/helper"
)

func Test217(t *testing.T) {
	input := [][]int{
		{1, 2, 3, 1},
		{1, 2, 3, 4},
		{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
	}
	expected := []bool{true, false, true}
	for i, v := range input {
		helper.AssertEquals(containsDuplicate(v), expected[i], v)
	}
}
