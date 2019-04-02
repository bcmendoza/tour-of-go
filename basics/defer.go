package basics

import "fmt"

// A defer statement defers the execution of a
// function until the surrounding function returns.
func Defer() {
	defer fmt.Printf(" world\n")
	fmt.Printf("hello")
}

// Deferred function calls are pushed onto a stack.
// When a function returns, its deferred calls are
// executed in last-in-first-out order.
// This will print from 9 to 0...
func Stack() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
