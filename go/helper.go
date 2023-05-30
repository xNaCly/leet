package leetcode_helpers

import "log"

func AssertEquals(a any, b any, k1 any) {
	if a != b {
		log.Panicf("Result '%v' not equal to '%v', for Key '%v' ", a, b, k1)
	}
}
