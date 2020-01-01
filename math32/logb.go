package math32

import (
	 "github.com/EngoEngine/math"
)

// Logb returns the binary exponent of x.
//
// Special cases are:
//	Logb(±Inf) = +Inf
//	Logb(0) = -Inf
//	Logb(NaN) = NaN
func Logb(x float32) float32 {
	return math.Logb(x)
}

// Ilogb returns the binary exponent of x as an integer.
//
// Special cases are:
//	Ilogb(±Inf) = MaxInt32
//	Ilogb(0) = MinInt32
//	Ilogb(NaN) = MaxInt32
func Ilogb(x float32) int {
	return math.Ilogb(x)
}
