package leetcode

import (
	"reflect"
	"testing"
)

func Test210217(t *testing.T) {
	type test struct {
		nums [][]int
		r, c int
		want [][]int
	}

	testGroup := []test{
		{[][]int{{1, 2}, {3, 4}}, 1, 4, [][]int{{1, 2, 3, 4}}},
		{[][]int{{1, 2}, {3, 4}}, 2, 4, [][]int{{1, 2}, {3, 4}}},
		{[][]int{{1, 2}, {3, 4}}, 4, 1, [][]int{{1}, {2}, {3}, {4}}},
	}

	for _, v := range testGroup {
		got := matrixReshape(v.nums, v.r, v.c)
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("error at:%v %v %v, want:%v, got %v", v.nums, v.r, v.c, v.want, got)
		}
	}
}
func Benchmark210217(b *testing.B) {
}
