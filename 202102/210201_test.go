package leetcode

import (
	"reflect"
	"testing"
)

func Test210201(t *testing.T) {
	type test struct {
		A    []int
		B    []int
		want []int
	}

	testGroup := []test{
		{[]int{1, 1}, []int{2, 2}, []int{1, 2}},
		{[]int{1, 2}, []int{2, 3}, []int{1, 2}},
		{[]int{2}, []int{1, 3}, []int{2, 3}},
		{[]int{1, 2, 5}, []int{2, 4}, []int{5, 4}},
	}

	for _, v := range testGroup {
		got := fairCandySwap(v.A, v.B)
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("error at test:%v, want:%v, got %v", v.A, v.want, got)
		}
	}
}
func Benchmark210201(b *testing.B) {

}
