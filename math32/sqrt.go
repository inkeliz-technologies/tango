package math32

import (
	 "github.com/EngoEngine/math"
)

// Sqrt returns the square root of x.
//
// Special cases are:
//	Sqrt(+Inf) = +Inf
//	Sqrt(±0) = ±0
//	Sqrt(x < 0) = NaN
//	Sqrt(NaN) = NaN
func Sqrt(x float32) float32 {
	return math.Sqrt(x)
}
