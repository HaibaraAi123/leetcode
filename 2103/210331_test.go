package leetcode

import (
	"reflect"
	"testing"
)

func Test210331(t *testing.T) {
	type test struct {
		nums []int
		want [][]int
	}

	testGroup := []test{
		{[]int{1, 2, 2}, [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}}},
		{[]int{0}, [][]int{{}, {0}}},
	}

	for _, v := range testGroup {
		got := subsetsWithDup(v.nums)
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("error at test:%v,want:%v, got:%v", v.nums, v.want, got)
		}
	}
}
func Benchmark210331(b *testing.B) {

}
