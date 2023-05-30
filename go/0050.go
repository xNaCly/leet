package main

import (
	"math"
)

// func myPow(x float64, n int) float64 {
// 	var neg bool
// 	r := 1.0

// 	if n == 0 {
// 		return 1.0
// 	} else if math.Abs(math.Abs(x)-1.0) < 1e-10 {
// 		if x < 0 && n < 0 {
// 			return 1.0
// 		} else if x < 0 {
// 			return -1.0
// 		}
// 		return 1.0
// 	} else if n == 1 {
// 		return x
// 	} else if n < 0 {
// 		n = int(math.Abs(float64(n)))
// 		neg = true
// 	}

// 	for i := 1; i <= n; i++ {
// 		r *= x
// 	}

// 	if neg {
// 		return 1.0 / r
// 	} else {
// 		return r
// 	}
// }

func myPow(x float64, n int) float64 {
	return math.Pow(x, float64(n))
}
