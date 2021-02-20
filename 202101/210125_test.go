package leetcode

import (
	"testing"
)

func Test210125(t *testing.T) {
	type test struct {
		grid []string
		want int
	}

	testGroup := []test{
		{[]string{" /", "/ "}, 2},
		{[]string{" /", "  "}, 1},
		{[]string{"\\/", "/\\"}, 4},
		{[]string{"/\\", "\\/"}, 5},
		{[]string{"//", "/ "}, 3},
	}

	for _, v := range testGroup {
		got := regionsBySlashes(v.grid)
		if got != v.want {
			t.Errorf("error at test:%v, want:%v, got %v", v.grid, v.want, got)
		}
	}
}
func Benchmark210125(b *testing.B) {

}
