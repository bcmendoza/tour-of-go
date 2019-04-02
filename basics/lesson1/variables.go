package lesson1

// var declares a list of variables
// As in function argument lists, the type is last
// Uninitialized variables default to the type's zero value
// 0 for numeric types
// false for the boolean type
// "" for strings
var i int

func Variables() (int, bool, string) {
	var no bool
	var empty string
	return i, no, empty // 0 false ""
}

// var can also initialize variables
// When initializing, type is inferred from the initializer
var j = 20

func Initializers() (j int, no, yes bool) {
	no, yes = false, true
	return // 0 false true
}

// Inside a function, the := short assignment statement
// can be used in place of a var declaration with implicit type.
// Outside a function, every statement begins with a keyword
// (var, func, and so on) and so the := construct is not available.
func ShortDeclarations() (int, bool, bool, string) {
	k := 3
	c, python, java := true, false, "no!"
	return k, c, python, java
}
