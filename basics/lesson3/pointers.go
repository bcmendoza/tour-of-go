package lesson3

var integer = 10

// GeneratePointer -
// A pointer holds the memory address of a value.
// The & operator generates a pointer to its value.
// This function returns the type *int, which is
// a pointer to some int value.
func GeneratePointer() *int {
	pointer := &integer
	return pointer
}

// UsePointer -
// The * operator denotes the pointer's underlying value.
// The pointer can be reassigned (i.e. `*pointer = 20`).
// And it can also be read (i.e. `*pointer`)
// This is called `dereferencing` or `indirecting`
func UsePointer() int {
	pointer := GeneratePointer()
	*pointer += 5
	return *pointer // 15
	// integer's new value is 15
}
