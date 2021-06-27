package leetcode

import (
	"testing"
)

func Test1215(t *testing.T) {
	type test struct {
		N    int
		want int
	}
	testGroup := []test{
		{10, 9},
		{1234, 1234},
		{332, 299},
		{565, 559},
		{121, 119},
	}
	for _, v := range testGroup {
		got := monotoneIncreasingDigits(v.N)
		if v.want != got {
			t.Errorf("error at %v want:%v got:%v", v.N, v.want, got)
		}
	}
}
func Benchmark1215(b *testing.B) {

}
