package leetcode

import (
	"testing"
)

func Test1225(t *testing.T) {
	type test struct {
		g    []int
		s    []int
		want int
	}
	testGroup := []test{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
	}

	for _, v := range testGroup {
		got := findContentChildren(v.g, v.s)
		if got != v.want {
			t.Errorf("error at %v, %v want:%v, got %v", v.g, v.s, v.want, got)
		}
	}
}
func Benchmark1225(b *testing.B) {

}
