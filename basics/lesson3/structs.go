package lesson3

// Vertex -
// A struct is a collection of fields.
type Vertex struct {
	X, Y int
}

// Struct fields are accesesed by dot-notation.
func Struct() Vertex {
	v := Vertex{1, 2}
	v.X = 4
	return v // {4, 2}
}

// StructPointer -
// Struct fields can also be accessed via struct pointer.
// Instead of using (*p).X, Go lets us omit the dereference operator.
// The special prefix `&` returns a pointer to the struct.
func StructPointer() Vertex {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	return *p // {1000000000, 2}
}

// StructLiteral -
// A struct literal denotes a newly allocated struct value
// by explicitly listing the values of its fields.
// You can list just a subset of fields by name (instead of order).
func StructLiteral() (Vertex, Vertex) {
	var (
		v1 = Vertex{X: 1} // Y: 0 is implicit
		v2 = Vertex{}     // X: 0 and Y: 0
	)
	return v1, v2
}
