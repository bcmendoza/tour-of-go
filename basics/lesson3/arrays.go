package lesson3

// Array -
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

// Slice -
// A slice is a dynamically-sized view into an array.
// The type []T is a slice with elements of type T.
// Slices are simply references to arrays, so modifying
// the elements of a slice modifies its actual array.
// The zero value of a slice is appropriately `nil`.
// In practice, slices are more commonly used than arrays.
func Slice() [4]string {
	names := [4]string{"John", "Paul", "George", "Ringo"}
	slice := names[1:] // high bounds omitted, uses default
	slice[0] = "PAUL"
	return names // [John PAUL George Ringo]
}

// SliceLiteral -
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

// SliceAttrs -
// A slice has both a length and capacity
// Length: Amt of elements. Length can be adjusted by re-slicing
// so that it includes more/less elements from the underlying array
// Capacity: Amt of elements left in the underlying array,
// counting from the first element in the slice to the end of array.
// A slice's length can be extended by re-slicing it to include
func SliceAttrs() (int, int, []int) {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[:0]                // give slice zero length
	s = s[:4]                // extend its length (since)
	s = s[1:]                // drop its first value
	return len(s), cap(s), s // 3 5 [3 5 7]
}

// MakeSlice -
// `make` creates dynamically-sized arrays
// It allocates a zeroed array and returns a slice referring to it
// Its 2nd arg specifies length, and 3rd specifies capacity
func MakeSlice() ([]int, []int, []int, []int) {
	a := make([]int, 5)    // [0 0 0 0 0]
	b := make([]int, 0, 5) // []
	c := b[:2]             // [0 0]
	d := c[2:5]            // [0 0 0 ]
	return a, b, c, d
}

// AppendToSlice -
// `append` adds new values to a slice
// If the underlying array is too small, a bigger array is allocated
// The returned slice will point to the newly allocated array
func AppendToSlice() []int {
	s := make([]int, 0) // nil
	s = append(s, 1)    // [1]
	s = append(s, 2, 3) // [1,2,3]
	return s
}
