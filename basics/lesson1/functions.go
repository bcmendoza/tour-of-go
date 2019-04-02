package lesson1

// Add -
// In this function example, add takes two params of type int
// Notice the type comes after the variable names
// When two or more consecutive named function parameters share a type,
// you can omit the type from all but the last.
func Add(x, y int) int {
	return x + y
}

// Swap -
// Functions can return any number of results
func Swap(x, y string) (string, string) {
	return y, x
}

// Split -
// Return values can be named in the return type declaration
// So the function will return those values by calling return
// This is a "naked" return, only for short functions by convention
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
