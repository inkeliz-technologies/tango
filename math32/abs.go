package math32

import (
	 "github.com/EngoEngine/math"
)

// Abs returns the absolute value of x.
//
// Special cases are:
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
func Abs(x float32) float32 {
	return math.Abs(x)
}
