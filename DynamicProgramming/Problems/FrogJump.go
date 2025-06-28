package dynamicprogramming

import "math"

// Frog Jump Problem
// Given an array of heights, the frog can jump to the next stone or the stone after that.
// The goal is to find the minimum cost to reach the last stone.

func FrogJumpTabulation(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		Step1 := dp[i-1] + int(math.Abs(float64(heights[i]-heights[i-1])))
		Step2 := math.MaxInt32
		if i > 1 {
			Step2 = dp[i-2] + int(math.Abs(float64(heights[i]-heights[i-2])))
		}
		dp[i] = min(Step1, Step2)
	}
	return dp[n-1]
}

func FrogJumpTabulationWithSpaceOptimization(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	prev1, prev2 := 0, 0
	for i := 1; i < n; i++ {
		Step1 := prev1 + int(math.Abs(float64(heights[i]-heights[i-1])))
		Step2 := math.MaxInt32
		if i > 1 {
			Step2 = prev2 + int(math.Abs(float64(heights[i]-heights[i-2])))
		}
		current := min(Step1, Step2)
		prev2 = prev1
		prev1 = current
	}
	return prev1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
