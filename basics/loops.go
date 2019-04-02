package basics

// Go only has one looping construct: `for`
// Uses basic structure (init, condition, post), but without parens
// The init statement will often be short var declaration
func ForLoop() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

// The init and post statements are optional
// Which converts it into a while loop (still using `for`)
func WhileFor() int {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	return sum
}
