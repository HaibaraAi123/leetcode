package leetcode

import (
	"testing"
)

func Test1223(t *testing.T) {
	type test struct {
		s    string
		want int
	}
	testGroup := []test{
		{"leetcode", 0},
		{"loveleetcode", 2},
	}

	for _, v := range testGroup {
		got := firstUniqChar(v.s)
		if v.want != got {
			t.Errorf("error at %v want:%v got:%v", v.s, v.want, got)
		}
	}
}
func Benchmark1223(b *testing.B) {

}
