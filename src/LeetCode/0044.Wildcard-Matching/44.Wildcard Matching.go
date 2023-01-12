package leetcode

func isMath(s string, p string) bool {
	n, m := len(s), len(p)
	dp := make([][]bool, m+1)
	//p串前i个字符能否匹配s串前j个字符
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	//相当于初始化第一列
	for i := 1; i <= m; i++ {
		if p[i-1] != '*' {
			break
		}
		dp[i][0] = true
	}
	for i := 1; i <= m; i++ {
		if p[i-1] >= 'a' && p[i-1] <= 'z' {
			for j := 1; j <= n; j++ {
				if p[i-1] == s[j-1] && dp[i-1][j-1] {
					dp[i][j] = true
				}
			}
		} else if p[i-1] == '?' {
			for j := 1; j <= n; j++ {
				if dp[i-1][j-1] {
					dp[i][j] = true
				}
			}
		} else {
			idx := n + 1
			for k := 0; k <= n; k++ {
				if dp[i-1][k] {
					idx = k
					break
				}
			}
			for j := idx; j <= n; j++ {
				dp[i][j] = true
			}
		}
	}
	return dp[m][n]
}
