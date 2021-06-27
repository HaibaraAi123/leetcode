package leetcode

import (
	"fmt"
	"testing"
)

func Test210419(t *testing.T) {
	type test struct {
		nums      []int
		val, want int
	}

	testGroup := []test{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for _, v := range testGroup {
		got := removeElement(v.nums, v.val)
		if got != v.want {
			t.Errorf("error at test:%v,want:%v, got:%v", v.nums, v.want, got)
		}
	}
}
func Test21041901(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	fmt.Println(nums[0:1], nums[1:])
	println(nums)
	println(len(nums), cap(nums))
}
func Benchmark210419(b *testing.B) {
}
