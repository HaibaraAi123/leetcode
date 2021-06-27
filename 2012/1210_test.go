package leetcode

import "testing"

func Test1210(t *testing.T) {
	type test struct {
		bills []int
		want  bool
	}
	testGroup := []test{
		{[]int{5, 5, 5, 10, 20}, true},
		{[]int{5, 5, 10}, true},
		{[]int{10, 10}, false},
		{[]int{5, 5, 10, 10, 20}, false},
	}
	for _, v := range testGroup {
		got := lemonadeChange(v.bills)
		if got != v.want {
			t.Errorf("error at bill:%v,want:%v,got:%v", v.bills, v.want, got)
		}
	}
}
func Benchmark1210(b *testing.B) {

}
