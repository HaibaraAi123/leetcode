package leetcode

import (
	"testing"
)

func Test210219(t *testing.T) {
	type test struct {
		A       []int
		K, want int
	}

	testGroup := []test{
		{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
	}

	for _, v := range testGroup {
		got := longestOnes(v.A, v.K)
		if got != v.want {
			t.Errorf("error at:%v %v, want:%v, got %v", v.A, v.K, v.want, got)
		}
	}
}
func Benchmark210219(b *testing.B) {
}
