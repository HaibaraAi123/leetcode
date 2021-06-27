package leetcode

//https://leetcode-cn.com/problems/rotate-list/

func rotateRight(head *ListNode, k int) *ListNode {
	length := 0
	if head == nil || head.Next == nil {
		return head
	}
	for p := head; p != nil; p = p.Next {
		length++
		if p.Next == nil {
			if k == 0 || k%length == 0 {
				return head
			}
			p.Next = head
			break
		}
	}
	if k > length {
		k %= length
	}
	k = length - k
	newhead := &ListNode{0, nil}
	for i, p := 0, head; p != nil; p, i = p.Next, i+1 {
		if k == i+1 {
			newhead.Next = p.Next
			p.Next = nil
		} else if i == length-1 {
			// p.Next = head 为什么此处接不到head 因为修改的是p.Next? p.Next.Next = head 应该是正确的
			break
		}
	}
	return newhead.Next
}
