package leetcode

import "testing"

func Test1204(t *testing.T) {
	type test struct {
		nums []int
		want bool
	}
	tests := []test{
		{nums: []int{1, 2, 4}, want: false},
		{nums: []int{1, 2}, want: false},
		{nums: []int{1, 2, 3, 3, 4, 5}, want: true},
		{nums: []int{1, 2, 3, 3, 4, 4, 5, 5}, want: true},
		{nums: []int{1, 2, 3, 4, 4, 5}, want: false},
	}
	for i, v := range tests {
		got := isPossible(v.nums)
		if v.want != got {
			t.Errorf("error at %v want:%v got:%v", i, v.want, got)
		}
	}
}
func Benchmark1204(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isPossible([]int{1, 2, 3, 4, 5, 6, 7})
	}
}
