package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 4, 5}
	c := [...]int{1, 2, 3, 4}
	//d := [...] int {1,2,3,4,5}
	t.Log(a == b)
	t.Log(a == c)
	t.Log(b == c)
	//t.Log(a == d)   //长度不同的数组做比较,会直接报编译错误
}

//  &^ 位运算符的  按位置零 右边位是1  结果都是0, 右边位是0 保留左边的值
func TestBitClear(t *testing.T) {
	a := 7
	t.Logf("%0b %0b", a, 2)
	a = a &^ 7
	t.Logf("%b", a)

}
