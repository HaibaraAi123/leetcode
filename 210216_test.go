package leetcode

import (
	"testing"
)

func Test210216(t *testing.T) {
	type test struct {
		nums []int
		want int
	}

	testGroup := []test{
		{[]int{1, 4, 3, 2}, 4},
		{[]int{6, 2, 6, 5, 1, 2}, 9},
	}

	for _, v := range testGroup {
		got := arrayPairSum(v.nums)
		if got != v.want {
			t.Errorf("error at:%v, want:%v, got %v", v.nums, v.want, got)
		}
	}
}
func Benchmark210216(b *testing.B) {
}
