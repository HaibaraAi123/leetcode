package leetcode

import (
	"testing"
)

func Test210220(t *testing.T) {
	type test struct {
		nums []int
		want int
	}

	testGroup := []test{
		{[]int{1, 2, 2, 3, 1}, 2},
		{[]int{1, 2, 2, 3, 1, 4, 2}, 6},
	}

	for _, v := range testGroup {
		got := findShortestSubArray(v.nums)
		if got != v.want {
			t.Errorf("error at:%v, want:%v, got %v", v.nums, v.want, got)
		}
	}
}
func Benchmark210220(b *testing.B) {
}
