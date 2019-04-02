package lesson2

import "fmt"

// Defer -
// A defer statement defers the execution of a
// function until the surrounding function returns.
func Defer() {
	defer fmt.Printf(" world\n")
	fmt.Printf("hello")
}

// Stack -
// Deferred function calls are pushed onto a stack.
// When a function returns, its deferred calls are
// executed in last-in-first-out order.
// This will print from 9 to 0...
func Stack() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

// PanicRecover -
// `panic` stops the ordinary flow of control and calls deferred fns
// `panic` receives a value that can be passed to `recover`
// `recover` regains control of the panicking goroutine
// `recover` is only used inside deferred functions
func PanicRecover() {
	panicker()
	fmt.Println("Recovered!")
}

func panicker() {
	defer func() {
		fmt.Println("Now recovering...")
		msg := recover()
		fmt.Println(msg)
	}()
	panic("I am panicking!!!")
}
