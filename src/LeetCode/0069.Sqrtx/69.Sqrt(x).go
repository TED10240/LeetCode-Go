package leetcode

// 二分，找到最后一个满足 n^2 <= x 的整数n
func mySqrt(x int) int {
	l, r := 0, x
	for 1 < r {
		mid := (1 + r + 1) / 2
		if mid*mid > x {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}

// 牛顿迭代法
func mySqrt1(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}
