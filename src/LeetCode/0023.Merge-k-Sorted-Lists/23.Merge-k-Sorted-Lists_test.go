package leetcode

import (
	"fmt"
	"testing"
)

type question23 struct {
	para23
	ans23
}

// para 是参数
// one 代表第一个参数
type para23 struct {
	one [][]int
}

// ans 是答案
// one 代表第一个答案
type ans23 struct {
	one []int
}

func Test_Problem23(t *testing.T) {

	qs := []question23{

		{
			para23{[][]int{}},
			ans23{[]int{}},
		},

		{
			para23{[][]int{
				{1},
				{1},
			}},
			ans23{[]int{1, 1}},
		},

		{
			para23{[][]int{
				{1, 2, 3, 4},
				{1, 2, 3, 4},
			}},
			ans23{[]int{1, 1, 2, 2, 3, 3, 4, 4}},
		},

		{
			para23{[][]int{
				{1, 2, 3, 4, 5},
				{1, 2, 3, 4, 5},
			}},
			ans23{[]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}},
		},

		{
			para23{[][]int{
				{1},
				{9, 9, 9, 9, 9},
			}},
			ans23{[]int{1, 9, 9, 9, 9, 9}},
		},

		{
			para23{[][]int{
				{9, 9, 9, 9, 9},
				{1},
			}},
			ans23{[]int{1, 9, 9, 9, 9, 9}},
		},

		{
			para23{[][]int{
				{2, 3, 4},
				{4, 5, 6},
			}},
			ans23{[]int{2, 3, 4, 4, 5, 6}},
		},

		{
			para23{[][]int{
				{1, 3, 8},
				{1, 7},
			}},
			ans23{[]int{1, 1, 3, 7, 8}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 23------------------------\n")

	for _, q := range qs {
		var ls []*ListNode
		for _, qq := range q.para23.one {
			ls = append(ls, Ints2List(qq))
		}
		fmt.Printf("【input】:%v       【output】:%v\n", q.para23.one, List2Ints(mergeKLists(ls)))
	}
	fmt.Printf("\n\n\n")
}

func List2Ints(head *ListNode) []int {
	// 链条深度限制，链条深度超出此限制，会 panic
	limit := 100

	times := 0

	res := []int{}
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// Ints2List convert []int to List
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}
