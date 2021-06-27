package leetcode

import (
	"testing"
)

func Test210221(t *testing.T) {
	type test struct {
		nums        []int
		limit, want int
	}

	testGroup := []test{
		{[]int{8, 2, 4, 7}, 4, 2},
		{[]int{10, 1, 2, 4, 7, 2}, 5, 4},
		{[]int{4, 2, 2, 2, 4, 4, 2, 2}, 0, 3},
	}

	for _, v := range testGroup {
		got := longestSubarray(v.nums, v.limit)
		if got != v.want {
			t.Errorf("error at:%v %v, want: %v, got %v", v.nums, v.limit, v.want, got)
		}
	}
}
func Benchmark210221(b *testing.B) {
}
