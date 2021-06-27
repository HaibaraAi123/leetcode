package leetcode

import "testing"

func Test1218(t *testing.T) {
	type test struct {
		s    string
		t    string
		want byte
	}
	testGroup := []test{
		{"abcd", "abcde", 'e'},
		{"", "y", 'y'},
		{"a", "aa", 'a'},
		{"ae", "aea", 'a'},
	}
	for _, v := range testGroup {
		got := findTheDifference(v.s, v.t)
		if v.want != got {
			t.Errorf("error at %v, %v want:%v got:%v", v.s, v.t, v.want, got)
		}
	}
}
func Benchmark1218(b *testing.B) {

}
