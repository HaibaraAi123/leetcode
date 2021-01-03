package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func list2slice(head *ListNode) []int {
	result := []int{}
	for p := head; p != nil; p = p.Next {
		result = append(result, p.Val)
	}
	return result
}
func slice2list(result []int) *ListNode {
	head := &ListNode{}
	tail := head
	for _, v := range result {
		tail.Next = &ListNode{v, nil}
		tail = tail.Next
	}
	return head.Next
}
func Test210103(t *testing.T) {
	type test struct {
		head []int
		x    int
		want []int
	}

	testGroup := []test{
		{[]int{1, 4, 3, 2, 5, 2}, 3, []int{1, 2, 2, 4, 3, 5}},
	}

	for i, v := range testGroup {
		fmt.Println("--Test", i, "--")
		got := list2slice(partition(slice2list(v.head), v.x))
		if !reflect.DeepEqual(v.want, got) {
			t.Errorf("error at test:%v, want:%v, got %v", i, v.want, got)
		}
	}
}
func Benchmark210103(b *testing.B) {

}
