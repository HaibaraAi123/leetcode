package leetcode

import (
	"testing"
)

func Test210227(t *testing.T) {
	type test struct {
		s       string
		k, want int
	}

	testGroup := []test{
		{"aaabb", 3, 3},
		{"ababbc", 2, 5},
	}

	for _, v := range testGroup {
		got := longestSubstring(v.s, v.k)
		if got != v.want {
			t.Errorf("error at:%v %v, want: %v, got %v", v.s, v.k, v.want, got)
		}
	}
}
func Benchmark210227(b *testing.B) {
}
