package leetcode

import (
	"fmt"
	"testing"
)

type question82 struct {
	para82
	ans82
}

// para 是参数
// one 代表第一个参数
type para82 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans82 struct {
	one []int
}

func Test_Problem82(t *testing.T) {

	qs := []question82{

		{
			para82{[]int{1, 1, 2, 2, 3, 4, 4, 4}},
			ans82{[]int{3}},
		},

		{
			para82{[]int{1, 1, 1, 1, 1, 1}},
			ans82{[]int{}},
		},

		{
			para82{[]int{1, 1, 1, 2, 3}},
			ans82{[]int{2, 3}},
		},

		{
			para82{[]int{1}},
			ans82{[]int{1}},
		},

		{
			para82{[]int{}},
			ans82{[]int{}},
		},

		{
			para82{[]int{1, 2, 2, 2, 2}},
			ans82{[]int{1}},
		},

		{
			para82{[]int{1, 1, 2, 3, 3, 4, 5, 5, 6}},
			ans82{[]int{2, 4, 6}},
		},

		{
			para82{[]int{1, 1, 2, 3, 3, 4, 5, 6}},
			ans82{[]int{2, 4, 5, 6}},
		},

		{
			para82{[]int{0, 1, 2, 2, 3, 4}},
			ans82{[]int{0, 1, 2, 2, 3, 4}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 82------------------------\n")

	for _, q := range qs {
		_, p := q.ans82, q.para82
		fmt.Printf("【input】:%v 【output】:%v\n", p, List2Ints(deleteDuplicates(Ints2List(p.one))))
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
