package leetcode

//https://leetcode-cn.com/problems/reverse-linked-list-ii/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	ltail := head
	for i := 1; i < left; i++ {
		if i < left-1 {
			ltail = ltail.Next
		} else {
			break
		}
	}
	rhead := ltail.Next
	mid := ltail.Next
	rhead = rhead.Next
	for i := left + 1; i <= right; i++ {
		ltail.Next = rhead
		rhead.Next = mid
		rhead = rhead.Next
	}
	mid.Next = rhead

	return head
}
