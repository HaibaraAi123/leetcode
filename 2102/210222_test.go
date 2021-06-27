package leetcode

import (
	"testing"
)

func Test210222(t *testing.T) {
	type test struct {
		matrix [][]int
		want   bool
	}

	testGroup := []test{
		{[][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}, true},
		{[][]int{{1, 2}, {2, 2}}, false},
	}

	for _, v := range testGroup {
		got := isToeplitzMatrix(v.matrix)
		if got != v.want {
			t.Errorf("error at:%v, want: %v, got %v", v.matrix, v.want, got)
		}
	}
}
func Benchmark210222(b *testing.B) {
}
