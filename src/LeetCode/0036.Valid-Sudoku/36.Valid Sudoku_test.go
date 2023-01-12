package leetcode

import (
	"fmt"
	"testing"
)

type question36 struct {
	para36
	ans36
}

// para 是参数
// one 代表第一个参数
type para36 struct {
	s [][]byte
}

// ans 是答案
// one 代表第一个答案
type ans36 struct {
	s bool
}

func Test_Problem36(t *testing.T) {

	qs := []question36{

		{
			para36{[][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}},
			ans36{true},
		},

		{
			para36{[][]byte{
				{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}},
			ans36{false},
		},

		{
			para36{[][]byte{
				{'.', '8', '7', '6', '5', '4', '3', '2', '1'},
				{'2', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'3', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'4', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'5', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'6', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'8', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'9', '.', '.', '.', '.', '.', '.', '.', '.'}}},
			ans36{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 36------------------------\n")

	for _, q := range qs {
		_, p := q.ans36, q.para36
		fmt.Printf("【input】:%v       【output】:%v\n", p, isValidSudoku(p.s))
	}
	fmt.Printf("\n\n\n")
}