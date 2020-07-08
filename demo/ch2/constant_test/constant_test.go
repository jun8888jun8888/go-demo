package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wedensday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wedensday)
}

func TestConstantTry1(t *testing.T) {
	a := 1
	t.Log(a&Readable, a&Writable, a&Executable)
}
