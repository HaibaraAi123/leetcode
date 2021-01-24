package leetcode

import (
	"testing"
)

func Test210123(t *testing.T) {
	type test struct {
		n           int
		connections [][]int
		want        int
	}

	testGroup := []test{
		{4, [][]int{{0, 1}, {0, 2}, {1, 2}}, 1},
		{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}, {1, 3}}, 2},
		{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 2}}, -1},
		{5, [][]int{{0, 1}, {0, 2}, {3, 4}, {2, 3}}, 0},
	}

	for _, v := range testGroup {
		got := makeConnected(v.n, v.connections)
		if v.want != got {
			t.Errorf("error at test:%v, %v, want:%v, got %v", v.n, v.connections, v.want, got)
		}
	}
}
func Benchmark210123(b *testing.B) {

}
