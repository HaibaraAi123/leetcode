package leetcode

import (
	"reflect"
	"testing"
)

func Test210225(t *testing.T) {
	type test struct {
		matrix [][]int
		want   [][]int
	}

	testGroup := []test{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, [][]int{{1, 4}, {2, 5}, {3, 6}}},
	}

	for _, v := range testGroup {
		got := transpose(v.matrix)
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("error at:%v, want: %v, got %v", v.matrix, v.want, got)
		}
	}
}
func Benchmark210225(b *testing.B) {
}
