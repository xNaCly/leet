package main

import (
	"fmt"
)

// INFO: this works, but exceeds the time limit :(
func numDupDigitsAtMostN(n int) int {
	if !(n >= 1 || n >= 10e9) {
		panic("impossible")
	}
	res := 0

	for i := 0; i <= n; i++ {
		m := make(map[rune]struct{})
		isOk := false

		for _, r := range fmt.Sprint(i) {
			if _, ok := m[r]; !ok {
				m[r] = struct{}{}
			} else {
				isOk = true
				break
			}
		}
		if isOk {
			res++
		}
	}

	return res
}
