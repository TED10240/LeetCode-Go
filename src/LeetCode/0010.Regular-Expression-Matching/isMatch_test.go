package leetcode

import (
	"fmt"
	"testing"
)

type para10 struct {
	one string
	two string
}
type ans10 struct {
	one bool
}

type question10 struct {
	para10
	ans10
}

func Test_Problem10(t *testing.T) {

	qs := []question10{
		{
			para10{"aa", "a"},
			ans10{false},
		},
		{
			para10{"aa", "a*"},
			ans10{true},
		},
		{
			para10{"ab", ".*"},
			ans10{true},
		},
	}

	for _, q := range qs {
		_, p := q.ans10, q.para10
		fmt.Printf("【input】:%v %v     【output】:%v\n", p.one, p.two, isMatch(p.one, p.two))
	}
}
