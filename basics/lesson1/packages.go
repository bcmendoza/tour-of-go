// Package lesson1 ...
// Every Go program is made up of packages.
// Programs start running in package main.
package lesson1

// This package groups imports into a parenthesized, "factored" import statement.
// The package name is the same as the last element of the import path.
import (
	"math/rand"
	"strconv"
)

// Packages -
// A name is exported if it begins with a capital letter.
// When importing a package, you can only refer to its exported names.
func Packages() string {
	return "My favorite number is " + strconv.Itoa(rand.Intn(10))
}
