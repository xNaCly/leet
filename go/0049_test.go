package main

import (
	"fmt"
	"testing"
)

func Test49(t *testing.T) {
	in := [][]string{
		{"eat", "tea", "tan", "ate", "nat", "bat"},
	}
	for _, v := range in {
		fmt.Println(v, "->", groupAnagrams(v))
	}
}
