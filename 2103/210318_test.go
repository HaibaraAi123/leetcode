package leetcode

import (
	"reflect"
	"testing"
)

func Test210318(t *testing.T) {
	type test struct {
		list        []int
		left, right int
		want        []int
	}
	InitList := func(a []int) *ListNode {
		head := &ListNode{}
		tail := head
		for _, v := range a {
			tmp := &ListNode{Val: v, Next: nil}
			tail.Next = tmp
			tail = tail.Next
		}
		return head.Next
	}
	PrintList := func(head *ListNode) []int {
		res := []int{}
		for tmp := head; ; tmp = tmp.Next {
			res = append(res, tmp.Val)
			if tmp.Next == nil {
				break
			}
		}
		return res
	}
	testGroup := []test{
		{[]int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5}},
	}

	for i, v := range testGroup {
		got := PrintList(reverseBetween(InitList(v.list), v.left, v.right))
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("error at test:%v, want:%v, got:%v", i, v.want, got)
		}
	}
}
func Benchmark210318(b *testing.B) {

}
