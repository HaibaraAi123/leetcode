package leetcode

import (
	"testing"
)

func Test210206(t *testing.T) {
	type test struct {
		cardPoints []int
		k          int
		want       int
	}

	testGroup := []test{
		{[]int{1, 2, 3, 4, 5, 6, 1}, 3, 12},
		{[]int{2, 2, 2}, 2, 4},
		{[]int{9, 7, 7, 9, 7, 7, 9}, 7, 55},
		{[]int{1, 1000, 1}, 1, 1},
		{[]int{1, 79, 80, 1, 1, 1, 200, 1}, 3, 202},
	}

	for _, v := range testGroup {
		got := maxScore(v.cardPoints, v.k)
		if got != v.want {
			t.Errorf("error at test:%v, want:%v %v, got %v", v.cardPoints, v.k, v.want, got)
		}
	}
}
func Benchmark210206(b *testing.B) {

}
