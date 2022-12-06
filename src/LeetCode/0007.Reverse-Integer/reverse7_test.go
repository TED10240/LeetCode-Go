package leetcode

import (
	"fmt"
	"testing"
)

type question7 struct {
	para7
	ans7
}

type para7 struct {
	one int
}

type ans7 struct {
	one int
}

func Test_Problem7(t *testing.T) {
	qs := []question7{
		{
			para7{321},
			ans7{123},
		},
		{
			para7{-123},
			ans7{-321},
		},
		{
			para7{120},
			ans7{21},
		},
		{
			para7{1534236469},
			ans7{0},
		},
	}

	for _, q := range qs {
		_, p := q.ans7, q.para7
		fmt.Printf("【input】:%v    【output】:%v\n", p.one, reverse7(p.one))

	}
}
