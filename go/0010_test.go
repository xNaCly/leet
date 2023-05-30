package main

import (
	"testing"

	"github.com/xNaCly/leet/go/helper"
)

func Test0010(t *testing.T) {
	in := [][]string{{"aa", "a"}, {"aa", "a*"}, {"ab", ".*"}}
	out := []bool{false, true, true}
	for i := 0; i < len(in); i++ {
		f := isMatch(in[i][0], in[i][1])
		helper.AssertEquals(f, out[i], in[i])
	}
}
