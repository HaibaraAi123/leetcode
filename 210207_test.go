package leetcode

import (
	"testing"
)

func Test210207(t *testing.T) {
	type test struct {
		nums []int
		want bool
	}

	testGroup := []test{
		{[]int{4, 2, 3}, true},
		{[]int{4, 2, 1}, false},
		{[]int{3, 4, 2, 3}, false},
	}

	for _, v := range testGroup {
		got := checkPossibility(v.nums)
		if got != v.want {
			t.Errorf("error at test:%v, want:%v, got %v", v.nums, v.want, got)
		}
	}
}
func Benchmark210207(b *testing.B) {

}
