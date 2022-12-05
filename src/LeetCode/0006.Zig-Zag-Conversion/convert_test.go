package leetcode

import (
	"fmt"
	"testing"
)

type question6 struct {
	para6
	ans6
}

type para6 struct {
	s       string
	numRows int
}

type ans6 struct {
	one string
}

func Test_Problem6(t *testing.T) {
	qs := []question6{
		{
			para6{"PAYPALISHIRING", 3},
			ans6{"PAHNAPLSIIGYIR"},
		},
	}

	for _, q := range qs {
		_, p := q.ans6, q.para6
		fmt.Printf("【input】:%v    【output】 :%v\n", p, convert(p.s, p.numRows))
	}
}
