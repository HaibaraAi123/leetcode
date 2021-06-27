package leetcode

import (
	"testing"
)

func Test210115(t *testing.T) {
	type test struct {
		stones [][]int
		want   int
	}

	testGroup := []test{
		{[][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}, 5},
		{[][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}, 3},
		{[][]int{{0, 0}}, 0},
	}

	for _, v := range testGroup {
		got := removeStones(v.stones)
		if got != v.want {
			t.Errorf("error at test:%v, want:%v, got %v", v.stones, v.want, got)
		}
	}
}
func Benchmark210115(b *testing.B) {

}
