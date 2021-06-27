package leetcode

import (
	"testing"
)

func Test210104(t *testing.T) {
	type test struct {
		n    int
		want int
	}

	testGroup := []test{
		{2, 1},
		{3, 2},
		{4, 3},
	}

	for _, v := range testGroup {
		got := fib(v.n)
		if v.want != got {
			t.Errorf("error at test:f(%v), want:%v, got %v", v.n, v.want, got)
		}
	}
}
func Benchmark210104(b *testing.B) {

}
