package leetcode

import (
	"testing"
)

func Test1221(t *testing.T) {
	type test struct {
		cost []int
		want int
	}
	testGroup := []test{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}

	for _, v := range testGroup {
		got := minCostClimbingStairs(v.cost)
		if v.want != got {
			t.Errorf("error at %v want:%v got:%v", v.cost, v.want, got)
		}
	}
}
func Benchmark1221(b *testing.B) {

}
