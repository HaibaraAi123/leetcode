package leetcode

import (
	"reflect"
	"testing"
)

func Test210315(t *testing.T) {
	type test struct {
		matrix [][]int
		want   []int
	}

	testGroup := []test{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}

	for _, v := range testGroup {
		got := spiralOrder(v.matrix)
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("error at test:%v, want:%v, got %v", v.matrix, v.want, got)
		}
	}
}
func Benchmark210315(b *testing.B) {

}
