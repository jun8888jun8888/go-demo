package map_test

import (
	"testing"
)

func TestMapInit(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	t.Log(m1[2])
	t.Log("m1 length ", len(m1))

	m2 := map[string]int{}
	m2["ak"] = 16
	t.Log("m2 length ", len(m2))

	m3 := make(map[int]int, 10)
	t.Log("len m3", len(m3))
}

func TestAccessNotExistingKey0(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	if v, err := m1[3]; err {
		t.Log(v)
	} else {
		t.Log("key 3 not exist", err)
	}
}

func TestTravl(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 2, 3: 3}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
