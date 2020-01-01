package math32

import (
	 "github.com/EngoEngine/math"
)

// Clamp returns f clamped to [low, high]
func Clamp(f, low, high float32) float32 {
	return math.Clamp(f, low, high)
}
