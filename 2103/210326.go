package leetcode

//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	m := make(map[int]bool)
	m[head.Val] = true
	for lp, rp := head, head.Next; rp != nil; {
		if _, ok := m[rp.Val]; ok {
			if rp.Next == nil {
				lp.Next = nil
				break
			}
			lp.Next = rp.Next
			rp = lp.Next
		} else {
			m[rp.Val] = true
			lp, rp = lp.Next, rp.Next
		}

	}
	return head
}
