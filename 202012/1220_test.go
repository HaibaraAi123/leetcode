package leetcode

import (
	"testing"
)

func Test1220(t *testing.T) {
	type test struct {
		s    string
		want string
	}
	testGroup := []test{
		{"bcabc", "abc"},
		{"cbacdcbc", "acdb"},
		{"cdadabcc", "adbc"},
		{"cbac", "bac"},
		{"bbcaac", "bac"},
		{"ccacbaba", "acb"},
		{"mitnlruhznjfyzmtmfnstsxwktxlboxutbic", "ilrhjfyzmnstwkboxuc"},
	}

	for _, v := range testGroup {
		got := removeDuplicateLetters(v.s)
		if v.want != got {
			t.Errorf("error at %v want:%v got:%v", v.s, v.want, got)
		}
	}
}
func Benchmark1220(b *testing.B) {

}
