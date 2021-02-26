package leetcode

import (
	"reflect"
	"testing"
)

func Test210226(t *testing.T) {
	type test struct {
		words, puzzles []string
		want           []int
	}

	testGroup := []test{
		{[]string{"aaaa", "asas", "able", "ability", "actt", "actor", "access"}, []string{"aboveyz", "abrodyz", "abslute", "absoryz", "actresz", "gaswxyz"}, []int{1, 1, 3, 2, 4, 0}},
		{[]string{"apple", "pleas", "please"}, []string{"aelwxyz", "aelpxyz", "aelpsxy", "saelpxy", "xaelpsy"}, []int{0, 1, 3, 2, 0}},
	}

	for _, v := range testGroup {
		got := findNumOfValidWords(v.words, v.puzzles)
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("error at:%v %v, want: %v, got %v", v.words, v.puzzles, v.want, got)
		}
	}
}
func Benchmark210226(b *testing.B) {
}
