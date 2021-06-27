package leetcode

import (
	"testing"
)

func Test210412(t *testing.T) {
	type test struct {
		nums []int
		want string
	}

	testGroup := []test{
		{[]int{10, 2}, "210"},
		{[]int{3, 30, 34, 5, 9}, "9534330"},
		{[]int{1}, "1"},
		{[]int{10}, "10"},
		{[]int{9, 95, 8, 80}, "995880"},
		{[]int{4, 45, 3, 30}, "454330"},
	}

	for _, v := range testGroup {
		got := largestNumber(v.nums)
		if got != v.want {
			t.Errorf("error at test:%v,want:%v, got:%v", v.nums, v.want, got)
		}
	}
}
func Benchmark210412(b *testing.B) {
}
