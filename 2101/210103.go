package leetcode

import "fmt"

//ListNode 分隔链表
type ListNode struct {
	Val  int
	Next *ListNode
}

//https://leetcode-cn.com/problems/partition-list/
func printList(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Printf("%v->", p.Val)
	}
	fmt.Printf("\n")
}
func partition(head *ListNode, x int) *ListNode {
	head1 := &ListNode{}
	tail1 := head1
	head2 := &ListNode{}
	tail2 := head2
	for p := head; p != nil; p = p.Next {
		if p.Val < x {
			tail1.Next = &ListNode{p.Val, nil}
			tail1 = tail1.Next
		} else {
			tail2.Next = &ListNode{p.Val, nil}
			tail2 = tail2.Next
		}
	}
	printList(head1)
	printList(head2)
	tail1.Next = head2.Next
	printList(head1)
	return head1.Next

}
