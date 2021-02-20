package leetcode

import "testing"

func Test1212(t *testing.T) {
	type test struct {
		nums []int
		want int
	}
	testGroup := []test{
		{[]int{1, 7, 4, 9, 2, 5}, 6},
		{[]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}, 7},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},
	}
	for _, v := range testGroup {
		got := wiggleMaxLength(v.nums)
		if got != v.want {
			t.Errorf("error at %v, want:%v, got:%v", v.nums, v.want, got)
		}
	}
}
func Benchmark1212(b *testing.B) {

}
