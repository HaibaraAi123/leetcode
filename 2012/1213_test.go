package leetcode

import "testing"

func Test1213(t *testing.T) {
	type test struct {
		nums []int
		want bool
	}
	testGroup := []test{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
	for _, v := range testGroup {
		got := containsDuplicate(v.nums)
		if got != v.want {
			t.Errorf("error at %v, want:%v, got:%v", v.nums, v.want, got)
		}
	}
}
func Benchmark1213(b *testing.B) {

}
