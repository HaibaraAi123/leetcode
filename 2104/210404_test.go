package leetcode

import (
	"testing"
)

func Test210404(t *testing.T) {
	type test struct {
		answer []int
		want   int
	}

	testGroup := []test{
		{[]int{1, 1, 2}, 5},
		{[]int{10, 10, 10}, 11},
		{[]int{}, 0},
		{[]int{1}, 2},
	}

	for _, v := range testGroup {
		got := numRabbits(v.answer)
		if got != v.want {
			t.Errorf("error at test:%v,want:%v, got:%v", v.answer, v.want, got)
		}
	}
}
func Benchmark210404(b *testing.B) {
}
