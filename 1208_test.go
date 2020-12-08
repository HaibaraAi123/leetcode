package leetcode

import (
	"reflect"
	"testing"
)

func Test1208(t *testing.T) {
	type test struct {
		s    string
		want []int
	}
	testGroup := []test{
		{"123456579", []int{123, 456, 579}},
		{"11235813", []int{1, 1, 2, 3, 5, 8, 13}},
		{"112358130", []int{}},
		{"0123", []int{}},
		{"12", []int{}},
	}
	for _, v := range testGroup {
		got := splitIntoFibonacci(v.s)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("error at %v want:%v, got:%v", v.s, v.want, got)
		}
	}
}
func Benchmark1208(b *testing.B) {

}
