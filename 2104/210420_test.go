package leetcode

import (
	"testing"
)

func Test210420(t *testing.T) {
	type test struct {
		haystack, needle string
		want             int
	}

	testGroup := []test{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
		{"", "", 0},
	}

	for _, v := range testGroup {
		got := strStr(v.haystack, v.needle)
		if got != v.want {
			t.Errorf("error at test:%v %v,want:%v, got:%v", v.haystack, v.needle, v.want, got)
		}
	}
}
func Test21042001(t *testing.T) {

}
func Benchmark210420(b *testing.B) {
}
