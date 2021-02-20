package leetcode

import (
	"reflect"
	"testing"
)

func Test1207(t *testing.T) {
	type test struct {
		strs []string
		want string
	}
	testGroup := []test{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"dog"}, "dog"},
		{[]string{}, ""},
	}
	for _, v := range testGroup {
		got := longestCommonPrefix(v.strs)
		if !reflect.DeepEqual(got, v.want) {
			t.Errorf("error at %v want:%v got:%v", v.strs, v.want, got)
		}
	}
}
func Benchmark1207(b *testing.B) {

}
