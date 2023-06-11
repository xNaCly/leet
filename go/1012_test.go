package main

import (
	"testing"

	"github.com/xNaCly/leet/go/helper"
)

func Test1012(t *testing.T) {
	input := map[int]int{
		// 20:   1,
		// 100:  10,
		// 1000: 262,
		6718458: 6205653,
	}
	for i, v := range input {
		helper.AssertEquals(numDupDigitsAtMostN(i), v, i)
	}
}
