package leetcode

import (
	"testing"
)

func Test210204(t *testing.T) {
	type test struct {
		nums []int
		k    int
		want float64
	}

	testGroup := []test{
		{[]int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{[]int{5}, 1, 5},
		{[]int{4, 2, 1, 3, 3}, 2, 3},
	}

	for _, v := range testGroup {
		got := findMaxAverage(v.nums, v.k)
		if got != v.want {
			t.Errorf("error at test:%v, want:%v, got %v", v.nums, v.want, got)
		}
	}
}
func Benchmark210204(b *testing.B) {

}
