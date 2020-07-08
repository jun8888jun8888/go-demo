package main

import "testing"

func TestSliceInit(t *testing.T) { //切片是一个共享的存储结构
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1, 2, 3, 4, 5)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	s2 = append(s2, 1, 2, 3)
	t.Log(s2)
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "seq", "oct", "nov", "dec"}
	q2 := year[3:6]
	t.Log(q2, len(q2), cap(q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "unknow"
	t.Log(q2)
	t.Log(year)
}

//slice 不能比较
func TestSliceEqual(t *testing.T) {
	q1 := []int{1, 2, 3}
	q2 := []int{1, 2, 3}
	if q1 == q2 {
		t.Log("equal")
	}
}
