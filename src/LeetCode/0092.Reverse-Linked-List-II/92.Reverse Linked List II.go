package leetcode

import "structures"

func reverseBetween(head *structures.ListNode, m, n int) *structures.ListNode {
	if head == nil || m >= n {
		return head
	}
	newHead := &structures.ListNode{Val: 0, Next: head}
	pre := newHead
	for count := 0; pre.Next != nil && count < m-1; count++ {
		pre = pre.Next
	}
	if pre.Next == nil {
		return head
	}
	cur := pre.Next
	for i := 0; i < n-m; i++ {
		tmp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = tmp
	}
	return newHead.Next

}
