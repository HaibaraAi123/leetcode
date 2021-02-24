package leetcode

import (
	"reflect"
	"testing"
)

func Test210224(t *testing.T) {
	type test struct {
		A    [][]int
		want [][]int
	}

	testGroup := []test{
		{[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}, [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
		{[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}, [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}},
	}

	for _, v := range testGroup {
		got := flipAndInvertImage(v.A)
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("error at:%v, want: %v, got %v", v.A, v.want, got)
		}
	}
}
func Benchmark210224(b *testing.B) {
}
