package leetcode

import (
	"testing"
)

func Test210304(t *testing.T) {
	type test struct {
		envelopes [][]int
		want      int
	}

	testGroup := []test{}

	for _, v := range testGroup {
		got := maxEnvelopes(v.envelopes)
		if got != v.want {
			t.Errorf("error at test:%v, want:%v, got %v", v.envelopes, v.want, got)
		}
	}
}
func Benchmark210304(b *testing.B) {

}
