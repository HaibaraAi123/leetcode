package leetcode

import (
	"reflect"
	"testing"
)

func Test1214(t *testing.T) {
	type test struct {
		strs []string
		want [][]string
	}
	testGroup := []test{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}}},
		{[]string{"bdddddddddd", "bbbbbbbbbbc"}, [][]string{{"bbbbbbbbbbc"}, {"bdddddddddd"}}},
	}
	for _, v := range testGroup {
		got := groupAnagrams(v.strs)
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("error at %v\nwant:%v\ngot:%v", v.strs, v.want, got)
		}
	}
}
func Benchmark1214(b *testing.B) {

}
