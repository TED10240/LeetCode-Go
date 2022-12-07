package leetcode

import (
	"fmt"
	"testing"
)

type para8 struct {
	one string
}

type ans8 struct {
	one int
}

type question8 struct {
	para8
	ans8
}

func Test_Problem8(t *testing.T) {
	qs := []question8{
		{
			para8{"42"},
			ans8{42},
		},
		{
			para8{"  -42"},
			ans8{-42},
		},
		{
			para8{"4192 with words"},
			ans8{41492},
		},
	}

	for _, q := range qs {
		_, p := q.ans8, q.para8
		fmt.Printf("【input】:%v     【output】:%v\n", p.one, myAtoi(p.one))
	}

}
