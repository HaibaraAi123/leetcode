package leetcode

import (
	"reflect"
	"testing"
)

func Test1224(t *testing.T) {
	type test struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}
	testGroup := []test{}

	for _, v := range testGroup {
		merge(v.nums1, v.m, v.nums2, v.n)
		if !reflect.DeepEqual(v.want, v.nums1) {
			t.Errorf("error at %v want:%v", v.nums1, v.want)
		}
	}
}
func Benchmark1224(b *testing.B) {

}
