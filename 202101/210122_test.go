package leetcode

import (
	"reflect"
	"testing"
)

func Test210122(t *testing.T) {
	type test struct {
		A    []int
		K    int
		want []int
	}

	testGroup := []test{
		{[]int{1, 2, 0, 0}, 34, []int{1, 2, 3, 4}},
		{[]int{2, 7, 4}, 181, []int{4, 5, 5}},
		{[]int{2, 1, 5}, 806, []int{1, 0, 2, 1}},
		{[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1, []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
		{[]int{0}, 0, []int{0}},
	}

	for _, v := range testGroup {
		got := addToArrayForm(v.A, v.K)
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("error at test:%v, %v, want:%v, got %v", v.A, v.K, v.want, got)
		}
	}
}
func Benchmark210122(b *testing.B) {

}
