package main

import (
	"testing"

	"github.com/xNaCly/leet/go/helper"
)

func Test242(t *testing.T) {
	inputs := [][]string{
		{"anagram", "nagaram"},
		{"rat", "car"},
		{"aacc", "ccac"},
	}
	expected := []bool{true, false, false}
	for i, input := range inputs {
		helper.AssertEquals(isAnagram(input[0], input[1]), expected[i], input)
	}
}
