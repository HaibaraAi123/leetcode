package leetcode

import "testing"

func Test1216(t *testing.T) {
	type test struct {
		pattern string
		s       string
		want    bool
	}
	testGroup := []test{
		{"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		{"aaaa", "dog cat cat dog", false},
		{"abba", "dog dog dog dog", false},
	}
	for _, v := range testGroup {
		got := wordPattern(v.pattern, v.s)
		if v.want != got {
			t.Errorf("error at %v, %v want:%v got:%v", v.pattern, v.s, v.want, got)
		}
	}
}
func Benchmark1216(b *testing.B) {

}
