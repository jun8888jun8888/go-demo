package fib

import "testing"

func TestFibList(t *testing.T) {
	var a int = 1
	var b int = 1
	t.Log(a)
	for i := 0; i < 50; i++ {
		t.Log("", b)
		//tmp := a
		//a = b
		//b = tmp + a
		a, b = b, a+b
	}

}
