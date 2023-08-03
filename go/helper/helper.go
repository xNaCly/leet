package helper

import (
	"log"
)

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

func AssertEqualsGeneric[T comparable](a T, b T) {
	if a != b {
		log.Panicf("assertion failed, value were not equal: %v != %v", a, b)
	}
}

func AssertFalse(a bool) {
	if a {
		log.Panic("assertion failed, value was true")
	}
}

func EqualsFloat(a float64, b float64) {
	margin := 1e-5
	sub1 := a - b
	sub2 := b - a
	if !(sub1 < margin && sub2 < margin) {
		log.Panicf("'%f' is not equal to '%f' within the margin of %f, absolut subtraction resulted in: '%f', '%f'", a, b, margin, sub1, sub2)
	}
}
