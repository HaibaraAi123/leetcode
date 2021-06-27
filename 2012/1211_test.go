package leetcode

import "testing"

func Test1211(t *testing.T) {
	type test struct {
		senate string
		want   string
	}
	testGroup := []test{
		{"RD", "Radiant"},
		{"RDD", "Dire"},
	}
	for _, v := range testGroup {
		got := predictPartyVictory(v.senate)
		if got != v.want {
			t.Errorf("error at %v, want:%v, got:%v", v.senate, v.want, got)
		}
	}
}
func Benchmark1211(b *testing.B) {

}
