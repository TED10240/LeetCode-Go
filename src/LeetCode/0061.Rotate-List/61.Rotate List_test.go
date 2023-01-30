package leetcode

import (
	"fmt"
	"testing"
)

type question61 struct {
	para61
	ans61
}

// para 是参数
// one 代表第一个参数
type para61 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans61 struct {
	one []int
}

func Test_Problem61(t *testing.T) {

	qs := []question61{

		{
			para61{[]int{1, 2, 3, 4, 5}, 2},
			ans61{[]int{4, 5, 1, 2, 3}},
		},

		{
			para61{[]int{1, 2, 3, 4, 5}, 3},
			ans61{[]int{4, 5, 1, 2, 3}},
		},

		{
			para61{[]int{0, 1, 2}, 4},
			ans61{[]int{2, 0, 1}},
		},

		{
			para61{[]int{1, 1, 1, 2}, 3},
			ans61{[]int{1, 1, 2, 1}},
		},

		{
			para61{[]int{1}, 10},
			ans61{[]int{1}},
		},

		{
			para61{[]int{}, 100},
			ans61{[]int{}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 61------------------------\n")

	for _, q := range qs {
		_, p := q.ans61, q.para61
		fmt.Printf("【input】:%v       【output】:%v\n", p, List2Ints(rotateRight(Ints2List(p.one), p.k)))
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
