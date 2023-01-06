package leetcode

import (
	"fmt"
	"testing"
)

type question38 struct {
	para38
	ans38
}

// para 是参数
// one 代表第一个参数
type para38 struct {
	num int
}

// ans 是答案
// one 代表第一个答案
type ans38 struct {
	one string
}

func Test_Problem38(t *testing.T) {

	qs := []question38{

		{
			para38{1},
			ans38{"1"},
		},

		{
			para38{4},
			ans38{"1211"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 38------------------------\n")

	for _, q := range qs {
		_, p := q.ans38, q.para38
		fmt.Printf("【input】:%v       【output】:%v\n", p, countAndSay(p.num))
	}
	fmt.Printf("\n\n\n")
}
