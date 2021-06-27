package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func Test210406(t *testing.T) {
	type test struct {
		nums    []int
		want    []int
		wantlen int
	}

	testGroup := []test{
		{[]int{1, 1, 1, 2, 2, 3}, []int{1, 1, 2, 2, 3}, 5},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, []int{0, 0, 1, 1, 2, 3, 3}, 7},
		{[]int{1, 1}, []int{1, 1}, 2},
		{[]int{1, 1, 1}, []int{1, 1}, 2},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 3},
	}

	for _, v := range testGroup {
		got := removeDuplicates(v.nums)
		if got != v.wantlen || !reflect.DeepEqual(v.want, v.nums[0:got]) {
			t.Errorf("error at test:%v,want:%v, got:%v", v.nums, v.want, got)
		}
	}
}
func Benchmark210406(b *testing.B) {
}
func Test21040601(t *testing.T) {
	var a = []int{1, 2, 3, 4, 5}
	a = append(a[0:2], a[3:]...)
	fmt.Println(a, len(a), cap(a))
}
