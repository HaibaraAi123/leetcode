package leetcode

import (
	"reflect"
	"testing"
)

func Test1206(t *testing.T) {
	type test struct {
		numRows int
		want    [][]int
	}
	testGroup := []test{
		{numRows: 1, want: [][]int{{1}}},
		{numRows: 2, want: [][]int{{1}, {1, 1}}},
		{numRows: 3, want: [][]int{{1}, {1, 1}, {1, 2, 1}}},
		{numRows: 4, want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}}},
		{numRows: 5, want: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
	}
	for _, v := range testGroup {
		got := generate(v.numRows)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("error at numRows:%v want: %v got:%v", v.numRows, v.want, got)
		}
	}
}
func Benchmark1206(b *testing.B) {

}
