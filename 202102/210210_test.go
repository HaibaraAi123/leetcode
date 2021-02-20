package leetcode

import (
	"testing"
)

func Test210210(t *testing.T) {
	type test struct {
		s1   string
		s2   string
		want bool
	}

	testGroup := []test{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
	}

	for _, v := range testGroup {
		got := checkInclusion(v.s1, v.s2)
		if got != v.want {
			t.Errorf("error at:%v %v, want:%v, got %v", v.s1, v.s2, v.want, got)
		}
	}
}
func Benchmark210210(b *testing.B) {

}
