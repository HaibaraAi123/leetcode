package leetcode

import (
	"testing"
)

func Test210401(t *testing.T) {
	type test struct {
		N, want int
	}

	testGroup := []test{
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 7},
		{10, 12},
	}

	for _, v := range testGroup {
		got := clumsy(v.N)
		if got != v.want {
			t.Errorf("error at test:%v,want:%v, got:%v", v.N, v.want, got)
		}
	}
}
func Benchmark210401(b *testing.B) {
}
