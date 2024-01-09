package main

import (
	"fmt"
)

func round(s string, k int) string {
	var res string
	for i := 0; i < len(s); i += k {
		r := 0
		in := i + k
		if in >= len(s) {
			in = len(s)
		}
		ss := s[i:in]
		for j := 0; j < len(ss); j++ {
			r += int(ss[j]) - 0x48
		}
		res += fmt.Sprint(r)
	}
	return string(res)
}

func digitSum(s string, k int) string {
	for len(s) > k {
		s = round(s, k)
	}
	return s
}
