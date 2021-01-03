package leetcode

import (
	"testing"
)

func Test210101(t *testing.T) {
	type test struct {
		flowerbed []int
		n         int
		want      bool
	}
	testGroup := []test{
		{[]int{1, 0, 0, 0, 1}, 1, true},
		{[]int{1, 0, 0, 0, 1}, 2, false},
		{[]int{1, 0, 0, 0, 1, 0, 0}, 2, true},
	}

	for _, v := range testGroup {
		got := canPlaceFlowers(v.flowerbed, v.n)
		if got != v.want {
			t.Errorf("error at %v want:%v, got %v", v.flowerbed, v.want, got)
		}
	}
}
func Benchmark210101(b *testing.B) {

}
