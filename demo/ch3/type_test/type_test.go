package type_test

import (
	"math"
	"testing"
)

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	b = int64(a) //显示的转换,不支持隐式转换

	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
	t.Log(math.MaxInt64, math.MaxInt8)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
