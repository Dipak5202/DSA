package dynamicprogramming

//LongestCommonSubstring calculates the length of the longest common substring
// between two strings using dynamic programming.
func LongestCommonSubstring(str1, str2 string) int {
	m := len(str1)
	n := len(str2)

	// Create a 2D array to store the lengths of common substrings
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	longest := 0
	// Fill the dp array
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				longest = max(longest, dp[i][j])
			}
		}
	}

	return longest
}
