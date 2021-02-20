package leetcode

import (
	"reflect"
	"testing"
)

func Test210110(t *testing.T) {
	type test struct {
		nums []int
		want []string
	}

	testGroup := []test{
		{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
		{[]int{}, []string{}},
		{[]int{-1}, []string{"-1"}},
		{[]int{0}, []string{"0"}},
	}

	for _, v := range testGroup {
		got := summaryRanges(v.nums)
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("error at test:%v, want:%v, got %v", v.nums, v.want, got)
		}
	}
}
func Benchmark210110(b *testing.B) {

}
