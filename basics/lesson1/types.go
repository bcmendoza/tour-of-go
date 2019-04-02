package lesson1

/*
Go's basic types are:
- bool
- string
- int  int8  int16  int32  int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte (alias for uint8)
- rune (alias for int32, represents a Unicode code point)
- float32 float64
- complex64 complex128

The int, uint, and uintptr types are usually
32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
When you need an integer value you should use int unless you have a
specific reason to use a sized or unsigned integer type.
*/

import (
	"fmt"
	"math"
	"math/cmplx"
)

// The example shows variables of several types
// Note that var declarations may be "factored" into blocks,
// as with import statements.
var (
	toBe   bool       = false
	maxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// Types -
func Types() (string, string, string) {
	b := fmt.Sprintf("%T %v\n", toBe, toBe)
	u := fmt.Sprintf("%T %v\n", maxInt, maxInt)
	c := fmt.Sprintf("%T %v", z, z)
	return b, u, c
}

// Conversions -
// The expression T(v) converts the value v to the type T.
// Assignment between items of different type requires explicit conversion
func Conversions() (int, int, uint) {
	x, y := 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	return x, y, z // 3, 4, 5
}
