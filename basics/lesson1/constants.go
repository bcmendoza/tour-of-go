package lesson1

// Constants are declared like variables, but with `const`
// They cannot be declared using the := syntax
const Pi = 3.14

func Constants() string {
	const World = "世界"
	return World
}

// Numeric constants are high-precision values
// An untyped constant takes the type needed by its context
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func NumericConstants() (int, float64, float64) {
	return needInt(Small), needFloat(Small), needFloat(Big)
}
