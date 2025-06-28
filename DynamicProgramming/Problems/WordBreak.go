package dynamicprogramming

func WordBreak(s string, wordDict []string) bool {

	wordset := make(map[string]bool)
	for _, word := range wordDict {
		wordset[word] = true
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordset[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
