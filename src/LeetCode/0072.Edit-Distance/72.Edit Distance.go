package leetcode

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	if m == 0 || n == 0 {
		return max(m, n)
	}
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		dp[i][n] = m - i
	}
	for j := 0; j < n; j++ {
		dp[m][j] = n - j
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = 1 + dp[i+1][j]
				dp[i][j] = min(dp[i][j], 1+dp[i][j+1])
				dp[i][j] = min(dp[i][j], 1+dp[i+1][j+1])
			}
		}
	}
	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
