package leetcode

import (
	"fmt"
	"testing"
)

type question24 struct {
	para24
	ans24
}

// para 是参数
// one 代表第一个参数
type para24 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans24 struct {
	one []int
}

func Test_Problem24(t *testing.T) {

	qs := []question24{

		{
			para24{[]int{}},
			ans24{[]int{}},
		},

		{
			para24{[]int{1}},
			ans24{[]int{1}},
		},

		{
			para24{[]int{1, 2, 3, 4}},
			ans24{[]int{2, 1, 4, 3}},
		},

		{
			para24{[]int{1, 2, 3, 4, 5}},
			ans24{[]int{2, 1, 4, 3, 5}},
		},

		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 24------------------------\n")

	for _, q := range qs {
		_, p := q.ans24, q.para24
		fmt.Printf("【input】:%v       【output】:%v\n", p, List2Ints(swapPairs(Ints2List(p.one))))
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
