package leetcode

import (
	"testing"
)

func Test210223(t *testing.T) {
	type test struct {
		customers, grumpy []int
		X, want           int
	}

	testGroup := []test{
		{[]int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3, 16},
	}

	for _, v := range testGroup {
		got := maxSatisfied(v.customers, v.grumpy, v.X)
		if got != v.want {
			t.Errorf("error at:%v %v %v, want: %v, got %v", v.customers, v.grumpy, v.X, v.want, got)
		}
	}
}
func Benchmark210223(b *testing.B) {
}
