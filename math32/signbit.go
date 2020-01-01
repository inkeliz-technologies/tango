package math32

import (
	 "github.com/EngoEngine/math"
)

// Signbit returns true if x is negative or negative zero.
func Signbit(x float32) bool {
	return math.Signbit(x)
}
