package leetcode

import (
	"testing"
)

func Test210202(t *testing.T) {
	type test struct {
		s    string
		k    int
		want int
	}

	testGroup := []test{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
	}

	for _, v := range testGroup {
		got := characterReplacement(v.s, v.k)
		if got != v.want {
			t.Errorf("error at test:%v, want:%v, got %v", v.s, v.want, got)
		}
	}
}
func Benchmark210202(b *testing.B) {

}
