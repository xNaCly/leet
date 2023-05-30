package helper

import "log"

func AssertEquals(a any, b any, k1 any) {
	if a != b {
		log.Panicf("Result '%v' not equal to '%v', for Key '%v' ", a, b, k1)
	}
}

func AssertTrue(a bool) {
	if !a {
		log.Panic("assertion failed, value was false")
	}
}

func AssertFalse(a bool) {
	if a {
		log.Panic("assertion failed, value was true")
	}
}
