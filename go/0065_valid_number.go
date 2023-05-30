package main

import (
	"leetcode/leetcode_helpers"
)

func isNumber(s string) bool {
	var v bool
	return v
}

func main() {
	ints := map[string]bool{
		"2":            true,
		"0089":         true,
		"-0.1":         true,
		"+3.14":        true,
		"4.":           true,
		"-.9":          true,
		"2e10":         true,
		"-90E3":        true,
		"3e+7":         true,
		"+6e-1":        true,
		"53.5e93":      true,
		"-123.456e789": true,
		"0e1":          true,
		"abc":          false,
		"1a":           false,
		"1e":           false,
		"e3":           false,
		"99e2.5":       false,
		"--6":          false,
		"-+3":          false,
		"95a54e53":     false,
		"Infinity":     false,
		"inf":          false,
		"NaN":          false,
		"+Inf":         false,
		"-inf":         false,
		"7e69e":        false,
	}

	for k, v := range ints {
		leetcode_helpers.AssertEquals(v, isNumber(k), k)
	}
}
