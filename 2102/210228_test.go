package leetcode

import (
	"testing"
)

func Test210228(t *testing.T) {
	type test struct {
		A    []int
		want bool
	}

	testGroup := []test{
		{[]int{1, 2, 2, 3}, true},
		{[]int{6, 5, 4, 4}, true},
		{[]int{1, 3, 2}, false},
		{[]int{1, 2, 4, 5}, true},
		{[]int{1, 1, 1}, true},
	}

	for _, v := range testGroup {
		got := isMonotonic(v.A)
		if got != v.want {
			t.Errorf("error at:%v, want: %v, got %v", v.A, v.want, got)
		}
	}
}
func Benchmark210228(b *testing.B) {
}
