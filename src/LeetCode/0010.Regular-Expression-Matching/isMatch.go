package leetcode

func isMatch(s string, p string) bool {
	n, m := len(s), len(p)
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}
	// 初始状态初始化
	dp[0][0] = true
	for j := 2; j <= m; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	// 中间状态转移
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else {
				if p[j-1] >= 'a' && p[j-1] <= 'z' {
					dp[i][j] = false
				} else if p[j-1] == '*' { // 是*
					if p[j-2] != s[i-1] && p[j-2] != '.' {
						dp[i][j] = dp[i][j-2]
					} else {
						dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j]
					}
				}
			}
		}
	}
	return dp[n][m]
}
