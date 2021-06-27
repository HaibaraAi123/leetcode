package leetcode

import "testing"

func Test1209(t *testing.T) {
	type test struct {
		m, n int
		want int
	}
	testGroup := []test{
		{3, 7, 28},
		{3, 2, 3},
		{7, 3, 28},
		{3, 3, 6},
	}
	for _, v := range testGroup {
		got := uniquePaths(v.m, v.n)
		if got != v.want {
			t.Errorf("error at m=%v,n=%v, want:%v, got:%v", v.m, v.n, v.want, got)
		}
	}
}
func Benchmark1209(b *testing.B) {

}
