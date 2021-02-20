package leetcode

import (
	"reflect"
	"testing"
)

func Test210203(t *testing.T) {
	type test struct {
		nums []int
		k    int
		want []float64
	}

	testGroup := []test{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []float64{1, -1, -1, 3, 5, 6}},
	}

	for _, v := range testGroup {
		got := medianSlidingWindow(v.nums, v.k)
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("error at test:%v, want:%v, got %v", v.nums, v.want, got)
		}
	}
}
func Benchmark210203(b *testing.B) {

}
