package leetcode

import (
	"testing"
)

func Test210215(t *testing.T) {
	type test struct {
		nums []int
		want int
	}

	testGroup := []test{
		{[]int{1, 1, 0, 1, 1, 1}, 3},
		{[]int{1}, 1},
	}

	for _, v := range testGroup {
		got := findMaxConsecutiveOnes(v.nums)
		if got != v.want {
			t.Errorf("error at:%v, want:%v, got %v", v.nums, v.want, got)
		}
	}
}
func Benchmark210215(b *testing.B) {

}
