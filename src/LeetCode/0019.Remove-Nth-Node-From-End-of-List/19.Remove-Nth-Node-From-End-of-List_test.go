package leetcode

import (
	"fmt"
	"testing"
)

type question19 struct {
	para19
	ans19
}

// para 是参数
// one 代表第一个参数
type para19 struct {
	one []int
	n   int
}

// ans 是答案
// one 代表第一个答案
type ans19 struct {
	one []int
}

func Test_Problem19(t *testing.T) {

	qs := []question19{

		{
			para19{[]int{1}, 3},
			ans19{[]int{1}},
		},

		{
			para19{[]int{1, 2}, 2},
			ans19{[]int{2}},
		},

		{
			para19{[]int{1}, 1},
			ans19{[]int{}},
		},

		{
			para19{[]int{1, 2, 3, 4, 5}, 10},
			ans19{[]int{1, 2, 3, 4, 5}},
		},

		{
			para19{[]int{1, 2, 3, 4, 5}, 1},
			ans19{[]int{1, 2, 3, 4}},
		},

		{
			para19{[]int{1, 2, 3, 4, 5}, 2},
			ans19{[]int{1, 2, 3, 5}},
		},

		{
			para19{[]int{1, 2, 3, 4, 5}, 3},
			ans19{[]int{1, 2, 4, 5}},
		},
		{
			para19{[]int{1, 2, 3, 4, 5}, 4},
			ans19{[]int{1, 3, 4, 5}},
		},

		{
			para19{[]int{1, 2, 3, 4, 5}, 5},
			ans19{[]int{2, 3, 4, 5}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 19------------------------\n")

	for _, q := range qs {
		_, p := q.ans19, q.para19
		fmt.Printf("【input】:%v       【output】:%v\n", p, List2Ints(removeNthFromEnd(Ints2List(p.one), p.n)))
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
