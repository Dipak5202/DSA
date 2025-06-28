package main

import (
	dynamicprogramming "DSA/DynamicProgramming/Problems"
	"fmt"
)

func main() {
	println("Dynamic Programming Problems")

	WordBreak()
	// HouseRobber()
	// LongestCommonSubstring()
	// LongestCommonSubsequence()
	// FrogJump()
	// ClimbingStairs()
	// Fibonacci()
}

// word break
func WordBreak() {
	println("Word Break Problem")
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println("Can the word be segmented?", dynamicprogramming.WordBreak(s, wordDict))
}

// House Robber
func HouseRobber() {
	println("House Robber Problem")
	nums := []int{2, 17, 9, 3, 5, 4}
	fmt.Println("Max Amount Robbed:", dynamicprogramming.HouseRobber(nums))
}

// longest common substring
func LongestCommonSubstring() {
	println("Longest Common Substring Problem")
	str1 := "abcde"
	str2 := "abfce"
	fmt.Println("Longest Common Substring:", dynamicprogramming.LongestCommonSubstring(str1, str2))
}

// longest common subsequence
func LongestCommonSubsequence() {
	println("Longest Common Subsequence Problem")
	str1 := "abcde"
	str2 := "ace"
	fmt.Println("Longest Common Subsequence:", dynamicprogramming.CommonSubsequence(str1, str2))
}

// FrogJump()
func FrogJump() {
	println("Frog Jump Problem")
	heights := []int{10, 30, 40, 20}
	fmt.Println("Frog Jump (Tabulation):", dynamicprogramming.FrogJumpTabulation(heights))
	fmt.Println("Frog Jump (Tabulation with Space Optimization):", dynamicprogramming.FrogJumpTabulationWithSpaceOptimization(heights))

}

func Fibonacci() {
	println("Fibonacci Problem")
	n := 5

	fmt.Println("Fibonacci (Recursive):", dynamicprogramming.Fibonacci(n))
	fmt.Println("Fibonacci (Memoization):", dynamicprogramming.FibonacciMemoization(n))
	fmt.Println("Fibonacci (Tabulation):", dynamicprogramming.FibonacciTabulation(n))
	fmt.Println("Fibonacci (Space Optimized):", dynamicprogramming.FibonacciSpaceOptimized(n))
}

func ClimbingStairs() {
	println("Climbing Stairs Problem")
	n := 4

	fmt.Println("Climbing Stairs (Recursive):", dynamicprogramming.ClimbingStairs(n))
	fmt.Println("Climbing Stairs (Memoization):", dynamicprogramming.ClimbingStairsMemoization(n))
	fmt.Println("Climbing Stairs (Tabulation):", dynamicprogramming.ClimbingStairsTabulation(n))
	fmt.Println("Climbing Stairs (Space Optimized):", dynamicprogramming.ClimbingStairsSpaceOptimized(n))
}
