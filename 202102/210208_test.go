package leetcode

import (
	"testing"
)

func Test210208(t *testing.T) {
	type test struct {
		arr  []int
		want int
	}

	testGroup := []test{
		{[]int{9, 4, 2, 10, 7, 8, 8, 1, 9}, 5},
		{[]int{4, 8, 12, 16}, 2},
		{[]int{100}, 1},
		{[]int{0, 8, 45, 88, 48, 68, 28, 55, 17, 24}, 8},
		{[]int{4, 5}, 2},
	}

	for _, v := range testGroup {
		got := maxTurbulenceSize(v.arr)
		if got != v.want {
			t.Errorf("error at test:%v, want:%v, got %v", v.arr, v.want, got)
		}
	}
}
func Benchmark210208(b *testing.B) {

}
