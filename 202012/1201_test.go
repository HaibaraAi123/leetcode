package leetcode

import "testing"

func Benchmark1201(b *testing.B) {
	var nums = []int{1, 2, 3, 4, 5, 7, 7}
	for i := 0; i < 1; i++ {
		_ = searchRange(nums, 4)
	}
}
