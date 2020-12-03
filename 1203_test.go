package leetcode

import (
	"testing"
)

func Test1203(t *testing.T) {
	type test struct {
		n    int
		want int
	}
	tests := []test{
		{n: 0, want: 0},
		{n: 1, want: 0},
		{n: 10, want: 4},
		{n: 10000, want: 1229},
		{n: 499979, want: 41537},
		{n: 999983, want: 78497},
		{n: 1500000, want: 114155},
	}
	for _, tc := range tests {
		got := countPrimes(tc.n)
		if tc.want != got {
			t.Errorf("exceped:%v, got:%v", tc.want, got)
		}
	}
}
func Benchmark1203(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countPrimes(1000000)
	}
}
