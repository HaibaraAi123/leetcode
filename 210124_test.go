package leetcode

import (
	"testing"
)

func Test210124(t *testing.T) {
	type test struct {
		nums []int
		want int
	}

	testGroup := []test{
		{[]int{1, 3, 5, 4, 7}, 3},
		{[]int{2, 2, 2, 2, 2}, 1},
		{[]int{1, 3, 5, 7}, 4},
		{[]int{2, 1, 3}, 2},
	}

	for _, v := range testGroup {
		got := findLengthOfLCIS(v.nums)
		if got != v.want {
			t.Errorf("error at test:%v, want:%v, got %v", v.nums, v.want, got)
		}
	}
}
func Benchmark210124(b *testing.B) {

}
