package leetcode

import "testing"

func Test1217(t *testing.T) {
	type test struct {
		nums   []int
		target int
		want   int
	}
	testGroup := []test{
		{[]int{1, 3, 5, 6}, 1, 0},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1, 3, 5, 6, 7, 9}, 9, 5},
		{[]int{1, 2, 4, 5, 10, 11}, 11, 5},
	}
	for _, v := range testGroup {
		got := searchInsert(v.nums, v.target)
		if v.want != got {
			t.Errorf("error at %v, %v want:%v got:%v", v.nums, v.target, v.want, got)
		}
	}
}
func Benchmark1217(b *testing.B) {

}
