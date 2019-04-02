package basics

import (
	"fmt"
	"runtime"
	"time"
)

// Go's switch only runs the selected case,
// not all the cases that follow.
// The break statement is implied.
// Also, Go's switch cases need not be constants,
// and the values involved need not be integers.
func Switch() string {
	desc := "Go runs on "
	var result string
	switch os := runtime.GOOS; os {
	case "darwin":
		result = "OS X."
	case "linux":
		result = "Linux."
	default:
		result = fmt.Sprintf("%s", os)
	}
	return desc + result
}

/* Switch cases evaluate cases from top to bottom,
stopping when a case succeeds. For example:
switch i {
case 0:
case f(): // f is not called if i == 0
}
*/
func When() string {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		return "Today."
	case today + 1:
		return "Tomorrow."
	case today + 2:
		return "In two days."
	default:
		return "Too far away."
	}
}

// Switch without a condition is the same as `switch true`
// This can be a clean way of writing long if-then-else chains
func SwitchTrue() string {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		return "Good morning!"
	case t.Hour() < 17:
		return "Good afternoon."
	default:
		return "Good evening."
	}
}
