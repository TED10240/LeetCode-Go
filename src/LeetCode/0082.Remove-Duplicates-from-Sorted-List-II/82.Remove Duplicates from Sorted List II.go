package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	newHead := &ListNode{Next: head, Val: -999999}
	cur := newHead
	last := newHead
	front := head
	for front.Next != nil {
		if front.Val == cur.Val {
			front = front.Next
			continue
		} else {
			if cur.Next != front {
				last.Next = front
				if front.Next != nil && front.Next.Val != front.Val {
					last = front
				}
				cur = front
				front = front.Next
			} else {
				last = cur
				cur = cur.Next
				front = front.Next
			}
		}
	}
	if front.Val == cur.Val {
		last.Next = nil
	} else {
		if cur.Next != front {
			last.Next = front
		}
	}
	return newHead.Next

}
