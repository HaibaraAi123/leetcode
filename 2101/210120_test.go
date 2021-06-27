package leetcode

import (
	"testing"
)

func Test210120(t *testing.T) {
	type test struct {
		nums []int
		want int
	}

	testGroup := []test{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4}, 24},
		{[]int{-1, -2, -3}, -6},
	}

	for _, v := range testGroup {
		got := maximumProduct(v.nums)
		if got != v.want {
			t.Errorf("error at test:%v, want:%v, got %v", v.nums, v.want, got)
		}
	}
}
func Benchmark210120(b *testing.B) {

}
