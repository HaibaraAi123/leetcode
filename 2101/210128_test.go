package leetcode

import (
	"testing"
)

func Test210128(t *testing.T) {
	type test struct {
		nums []int
		want int
	}

	testGroup := []test{
		{[]int{1, 7, 3, 6, 5, 6}, 3},
		{[]int{1, 2, 3}, -1},
		{[]int{2, 1, -1}, 0},
		{[]int{0, 0, 0, 0, 1}, 4},
	}

	for _, v := range testGroup {
		got := pivotIndex(v.nums)
		if got != v.want {
			t.Errorf("error at test:%v, want:%v, got %v", v.nums, v.want, got)
		}
	}
}
func Benchmark210128(b *testing.B) {

}
