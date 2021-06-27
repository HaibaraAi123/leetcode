package leetcode

import (
	"reflect"
	"testing"
)

func Test210105(t *testing.T) {
	type test struct {
		s    string
		want [][]int
	}

	testGroup := []test{
		{"abbxxxxzzy", [][]int{{3, 6}}},
		{"abc", [][]int{}},
		{"abcdddeeeeaabbbcd", [][]int{{3, 5}, {6, 9}, {12, 14}}},
	}

	for _, v := range testGroup {
		got := largeGroupPositions(v.s)
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("error at test:%v, want:%v, got %v", v.s, v.want, got)
		}
	}
}
func Benchmark210105(b *testing.B) {

}
