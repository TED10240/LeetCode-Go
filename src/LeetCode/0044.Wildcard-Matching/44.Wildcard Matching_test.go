package leetcode

import (
	"fmt"
	"testing"
)

type question44 struct {
	para44
	ans44
}

// para 是参数
// one 代表第一个参数
type para44 struct {
	one string
	two string
}

// ans 是答案
// one 代表第一个答案
type ans44 struct {
	one bool
}

func Test_Problem44(t *testing.T) {

	qs := []question44{

		{
			para44{"aa", "a"},
			ans44{false},
		},
		{
			para44{"aa", "*"},
			ans44{true},
		},
		{
			para44{"cb", "?a"},
			ans44{false},
		},
		{
			para44{"abceb", "*a*b"},
			ans44{true},
		},
		{
			para44{"abceb", "*a*b"},
			ans44{true},
		},
		{
			para44{"acdcb", "a*c?b"},
			ans44{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 20------------------------\n")

	for _, q := range qs {
		_, p := q.ans44, q.para44
		fmt.Printf("【input】:%v       【output】:%v\n", p, isMath(p.one, p.two))
	}
}
