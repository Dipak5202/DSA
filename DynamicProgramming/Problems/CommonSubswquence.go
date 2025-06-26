package dynamicprogramming

// CommonSubsequence calculates the length of the longest common subsequence (LCS)
// between two strings using dynamic programming.
func CommonSubsequence(str1, str2 string) int {
	m := len(str1)
	n := len(str2)

	// Create a 2D array to store the lengths of LCS
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Fill the dp array
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
