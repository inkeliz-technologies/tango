package math32

import (
	 "github.com/EngoEngine/math"
)

// Asinh returns the inverse hyperbolic sine of x.
//
// Special cases are:
//	Asinh(±0) = ±0
//	Asinh(±Inf) = ±Inf
//	Asinh(NaN) = NaN
func Asinh(x float32) float32 {
	return math.Asinh(x)
}
