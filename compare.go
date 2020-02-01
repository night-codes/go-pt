package main

import "math"

// error threshold
const Epsilon = 0.00001

// Equal checks if float64 values are equal to a certain threshold
func Equal(a, b float64) bool {
	if math.Abs(a-b) < Epsilon {
		return true
	}
	return false
}
