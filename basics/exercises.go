package basics

import (
	"math"
)

/* Newton's Method: Sqrt
Find the square root of some given x float64 using Newton's method
Using some estimate z, keep applying Newton's method to it
until you reach an answer that is as close to the actual sqrt as possible.
A decent starting guess is 1, no matter the input.
To begin, repeat the calculation 10 times and print each z along the way.
See how close you get to the answer for various values of x (1, 2, 3, ...)
and how quickly the guess improves. */
func NewtonSqrt1(x float64) float64 {
	z := 1.0
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

/* Next, change the loop to stop when the value has stopped changing
(or only changes by a very small amount).
See if that's more or fewer than 10 iterations.
Try other initial guesses for z, like x, or x/2.
*/
func NewtonSqrt2(x float64) (float64, int) {
	z := x / 2
	i := 1
	var prev float64
	for {
		prev = z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prev) < 0.00001 {
			return z, i
		}
		i = i + 1
	}
}
