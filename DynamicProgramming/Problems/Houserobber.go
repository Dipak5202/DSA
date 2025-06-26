package dynamicprogramming

// You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

// Given an integer array nums representing the amount of money of each house, return the maximum amount of money you can rob tonight without alerting the police.
// O(1)-space version

// You’re given an array nums, where nums[i] is the amount of money hidden in the i-th house along a street.

// Rule: If you rob two adjacent houses, the police are alerted.

// Goal: Maximize the amount you can steal.

func HouseRobber(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// if n == 1 {
	// 	return nums[0]
	// }

	prev2 := 0       // dp[i-2]
	prev1 := nums[0] // dp[i-1]

	for i := 1; i < n; i++ {
		current := max(prev1, nums[i]+prev2)
		prev2 = prev1
		prev1 = current
	}

	return prev1
}
