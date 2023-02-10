package leetcode

import (
	"fmt"
	"testing"
)

type question72 struct {
	para72
	ans72
}

// para 是参数
// one 代表第一个参数
type para72 struct {
	s1 string
	s2 string
}

// ans 是答案
// one 代表第一个答案
type ans72 struct {
	one int
}

func Test_Problem72(t *testing.T) {

	qs := []question72{

		{
			para72{"horse", "ros"},
			ans72{3},
		},

		{
			para72{"intention", "execution"},
			ans72{5},
		},
	}
	fmt.Printf("------------------------Leetcode Problem 72------------------------\n")

	for _, q := range qs {
		_, p := q.ans72, q.para72
		fmt.Printf("【input】:%v       【output】:%v\n", p, minDistance(p.s1, p.s2))
	}
	fmt.Printf("\n\n\n")
}
