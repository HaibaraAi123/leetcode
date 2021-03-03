package leetcode

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	countBits(2)
	countBits(5)
	countBits(10)
}
func Test210303(t *testing.T) {
	type test struct {
		num  int
		want []int
	}

	testGroup := []test{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
	}

	for _, v := range testGroup {
		got := countBits(v.num)
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("error at test:%v, want:%v, got %v", v.num, v.want, got)
		}
	}
}
func Benchmark210303(b *testing.B) {

}
