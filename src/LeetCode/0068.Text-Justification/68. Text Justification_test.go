package leetcode

import (
	"fmt"
	"testing"
)

type question68 struct {
	para68
	ans68
}

// para 是参数
// one 代表第一个参数
type para68 struct {
	a []string
	b int
}

// ans 是答案
// one 代表第一个答案
type ans68 struct {
	one []string
}

func Test_Problem68(t *testing.T) {

	qs := []question68{

		{
			para68{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16},
			ans68{},
		},
		{
			para68{[]string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16},
			ans68{},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 67------------------------\n")

	for _, q := range qs {
		_, p := q.ans68, q.para68
		fmt.Printf("【input】:%v       【output】:%v\n", p, fullJustify(p.a, p.b))
	}
	fmt.Printf("\n\n\n")
}
