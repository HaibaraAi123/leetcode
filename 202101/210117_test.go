package leetcode

import (
	"testing"
)

func Test210117(t *testing.T) {
	type test struct {
		coordinates [][]int
		want        bool
	}

	testGroup := []test{
		{[][]int{{1, 2}, {3, 4}, {2, 3}, {4, 5}, {5, 6}, {6, 7}}, true},
		{[][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}, false},
	}

	for _, v := range testGroup {
		got := checkStraightLine(v.coordinates)
		if got != v.want {
			t.Errorf("error at test:%v, want:%v, got %v", v.coordinates, v.want, got)
		}
	}
}
func Benchmark210117(b *testing.B) {

}
