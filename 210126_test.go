package leetcode

import (
	"testing"
)

func Test210126(t *testing.T) {
	type test struct {
		dominoes [][]int
		want     int
	}

	testGroup := []test{
		{[][]int{{1, 2}, {3, 4}, {2, 1}, {5, 6}}, 1},
		{[][]int{{1, 2}, {1, 2}, {1, 1}, {1, 2}, {2, 2}}, 3},
	}

	for _, v := range testGroup {
		got := numEquivDominoPairs(v.dominoes)
		if got != v.want {
			t.Errorf("error at test:%v, want:%v, got %v", v.dominoes, v.want, got)
		}
	}
}
func Benchmark210126(b *testing.B) {

}
