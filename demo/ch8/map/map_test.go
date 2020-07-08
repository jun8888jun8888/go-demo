package _map

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}

	t.Log(m[1](ReturnValue()), m[2](3), m[3](4))
}

func ReturnValue() int {
	return 1
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[0] = true
	if mySet[1] {
		t.Log("0 is existing")
	} else {
		t.Log("0 is not existing")
	}
	delete(mySet, 1)
}
