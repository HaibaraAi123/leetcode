package leetcode

import (
	"fmt"
	"testing"
)

func Test210205(t *testing.T) {
	type test struct {
		s       string
		t       string
		maxCost int
		want    int
	}

	testGroup := []test{
		{"abcd", "bcdf", 3, 3},
		{"abcd", "cdef", 3, 1},
		{"abcd", "acde", 0, 1},
		{"pxezla", "loewbi", 25, 4},
	}
	printAsicii := func(s string) {
		fmt.Println()
		for _, v := range s {
			fmt.Printf("%v ", v)
		}

	}
	for _, v := range testGroup {
		got := equalSubstring(v.s, v.t, v.maxCost)
		if got != v.want {

			printAsicii(v.s)
			printAsicii(v.t)
			t.Errorf("error at test:%v, want:%v %v, got %v", v.s, v.t, v.want, got)

		}
	}
}
func Benchmark210205(b *testing.B) {

}
