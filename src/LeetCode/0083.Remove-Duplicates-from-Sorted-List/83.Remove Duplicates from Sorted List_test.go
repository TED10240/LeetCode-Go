package leetcode

import (
	"fmt"
	"structures"
	"testing"
)

type question83 struct {
	para83
	ans83
}

// para 是参数
// one 代表第一个参数
type para83 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans83 struct {
	one []int
}

func Test_Problem83(t *testing.T) {

	qs := []question83{

		{
			para83{[]int{1, 1, 2}},
			ans83{[]int{1, 2}},
		},

		{
			para83{[]int{1, 1, 2, 2, 3, 3, 3}},
			ans83{[]int{1, 2, 3}},
		},

		{
			para83{[]int{1, 1, 1, 1, 1, 1, 1, 1}},
			ans83{[]int{1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 83------------------------\n")

	for _, q := range qs {
		_, p := q.ans83, q.para83
		fmt.Printf("【input】:%v       【output】:%v\n", p, structures.List2Ints(deleteDuplicates(structures.Ints2List(p.one))))
	}
	fmt.Printf("\n\n\n")
}

func List2Ints(head *structures.ListNode) []int {
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
func Ints2List(nums []int) *structures.ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &structures.ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &structures.ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}
