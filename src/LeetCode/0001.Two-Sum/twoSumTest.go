package main

import "fmt"

type question1 struct {
	para1
	ans1
}

// para 是参数
// one 代表第一个参数
type para1 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans1 struct {
	one []int
}

func main() {

	qs := []question1{
		{
			para1{[]int{2, 11, 7, 15}, 9},
			ans1{[]int{0, 1}},
		},
	}
	for _, q := range qs {
		_, p := q.ans1, q.para1
		fmt.Printf("【input】:%v   【output】:%v\n", p, twoSum(p.nums, p.target))
	}

}
