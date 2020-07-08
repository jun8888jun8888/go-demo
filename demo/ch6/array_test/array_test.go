package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	t.Log(arr[1], arr[2])
	arr1 := [4]int{1, 2, 3}
	t.Log(arr1[3])
	arr3 := [...]string{"a", "1"}
	t.Log(arr3[0], arr3[1])

}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]string{"a", "1"}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for _, e := range arr3 {
		t.Log(e)
	}

}

func TestArraySection(t *testing.T) {
	arr3 := [...]string{"a", "1", "2"}
	t.Log(arr3[1:])
}

func TestMutlArray(t *testing.T) {
	arr4 := [3][2]float32{{1, 2}, {3, 4}, {5, 6}}
	for _, e := range arr4 {
		t.Log(e)
	}
}

//数组是可以比较的 相同类型,相同元素,相同顺序
//切片  连续的存储空间 数组切片
