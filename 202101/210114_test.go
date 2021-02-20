package leetcode

import (
	"reflect"
	"testing"
)

func Test210114(t *testing.T) {
	type test struct {
		A    []int
		want []bool
	}

	testGroup := []test{
		{[]int{0, 1, 1}, []bool{true, false, false}},
		{[]int{1, 1, 1}, []bool{false, false, false}},
		{[]int{0, 1, 1, 1, 1, 1}, []bool{true, false, false, false, true, false}},
		{[]int{1, 1, 1, 0, 1}, []bool{false, false, false, false, false}},
		{[]int{1, 0, 1, 1, 1, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 1, 0, 0, 1, 0}, []bool{false, false, true, false, false, false, false, false, false, false, true, true, true, true, true, true, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, true, false, false, true, false, false, true, true, true, true, true, true, true, false, false, true, false, false, false, false, true, true}},
	}

	for _, v := range testGroup {
		got := prefixesDivBy5(v.A)
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("error at test:%v, want:%v, got %v", v.A, v.want, got)
		}
	}
}
func Benchmark210114(b *testing.B) {

}
