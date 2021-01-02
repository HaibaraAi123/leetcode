package leetcode

import (
	"testing"
)

func Test1230(t *testing.T) {
	type test struct {
		stones []int
		want   int
	}
	testGroup := []test{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{3, 7, 2}, 2},
	}

	for _, v := range testGroup {
		got := lastStoneWeight(v.stones)
		if got != v.want {
			t.Errorf("error at %v want:%v, got %v", v.stones, v.want, got)
		}
	}
}
func Benchmark1230(b *testing.B) {

}
