package basics

// The type [n]T is an array of n values of type T.
// An array's length is fixed as part of its type
func Array() ([2]string, [6]int) {
	// Here's an uninitialized array
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	// Here's an array literal
	primes := [6]int{2, 3, 5, 7, 11, 13}
	return a, primes
}

// A slice is a dynamically-sized view into an array.
// The type []T is a slice with elements of type T.
// Slices are simply references to arrays, so modifying
// the elements of a slice modifies its actual array.
// In practice, slices are more commonly used than arrays.
func Slice() [4]string {
	names := [4]string{"John", "Paul", "George", "Ringo"}
	slice := names[1:] // high bounds omitted, uses default
	slice[0] = "PAUL"
	return names // [John PAUL George Ringo]
}

// A slice literal can be used to automatically create
// both an array and then a slice that references it
func SliceLiteral() []struct {
	i int
	b bool
} {
	q := []int{2, 3, 5, 7, 11, 13}
	r := []bool{true, false, true, true, false, true}
	s := []struct {
		i int
		b bool
	}{
		{q[0], r[0]},
		{q[1], r[1]},
		{q[2], r[2]},
		{q[3], r[3]},
		{q[4], r[4]},
		{q[5], r[5]},
	}
	return s
}
