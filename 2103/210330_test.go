package leetcode

import (
	"testing"
)

func Test210330(t *testing.T) {
	type test struct {
		matrix [][]int
		target int
		want   bool
	}

	testGroup := []test{
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 60, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 130, false},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 1, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 4, false},
		{[][]int{{1, 3, 5}}, 2, false},
		{[][]int{{1}, {3}, {5}}, 2, false},
	}

	for _, v := range testGroup {
		got := searchMatrix(v.matrix, v.target)
		if got != v.want {
			t.Errorf("error at test:%v,target:%v want:%v, got:%v", v.matrix, v.target, v.want, got)
		}
	}
}
func Benchmark210330(b *testing.B) {

}
