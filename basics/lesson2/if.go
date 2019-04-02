package lesson2

import (
	"fmt"
	"math"
)

// Go's `if` statements are like its for loops;
// the expression need not be surrounded by parentheses
// but the braces are required
func Sqrt(x float64) float64 {
	if x < 0 {
		return Sqrt(-x)
	}
	return math.Sqrt(x)
}

// The `if` can be prefaced by a short statement to execute
// Variables declared by the statement are only scoped within the `if`
func Pow(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)
	}
	return limit
}
