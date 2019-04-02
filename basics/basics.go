package basics

import (
	"fmt"

	"./lesson1"
	"./lesson2"
	"./lesson3"
)

// Lesson1 - Packages, Variables, Functions
func Lesson1() {
	fmt.Println(lesson1.Packages())
	fmt.Println(lesson1.Add(42, 13))
	fmt.Println(lesson1.Swap("hello", "world"))
	fmt.Println(lesson1.Split(17))
	fmt.Println(lesson1.Variables())
	fmt.Println(lesson1.Initializers())
	fmt.Println(lesson1.ShortDeclarations())
	fmt.Println(lesson1.Types())
	fmt.Println(lesson1.Conversions())
	fmt.Println(lesson1.Constants())
	fmt.Println(lesson1.NumericConstants())
}

// Lesson2 - Flow Control Statements
func Lesson2() {
	fmt.Println(lesson2.ForLoop())
	fmt.Println(lesson2.WhileFor())
	fmt.Println(lesson2.Sqrt(-4))
	fmt.Println(lesson2.Pow(3, 3, 26))
	fmt.Println(lesson2.Switch())
	fmt.Println(lesson2.When())
	fmt.Println(lesson2.SwitchTrue())
	lesson2.Defer()
	lesson2.Stack()
}

// Lesson3 - Pointers and More Types
func Lesson3() {
	fmt.Println(lesson3.UsePointer())
	fmt.Println(lesson3.Struct())
	fmt.Println(lesson3.StructPointer())
	fmt.Println(lesson3.StructLiteral())
	fmt.Println(lesson3.Array())
	fmt.Println(lesson3.Slice())
	fmt.Println(lesson3.SliceLiteral())
}

// Exercises -
func Exercises() {
	// fmt.Println(NewtonSqrt1(128))
	// fmt.Println(NewtonSqrt2(128))
}
