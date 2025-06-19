package backtracking

import "fmt"

func Subsets(nums []int) [][]int {
	var result [][]int
	var subset []int

	var backtrack func(int)
	backtrack = func(start int) {
		// Append a copy of the current subset to the result
		result = append(result, append([]int{}, subset...))

		for i := start; i < len(nums); i++ {
			// Include nums[i] in the current subset
			subset = append(subset, nums[i])
			// Recur with the next index
			backtrack(i + 1)
			// Exclude nums[i] from the current subset
			subset = subset[:len(subset)-1]
		}
	}

	backtrack(0)
	return result
}

func SubsetUsingBacktracking(Array []int, ans []int, i int) {
	if i == len(Array) {
		// Print the current subset
		for _, v := range ans {
			print(v, " ")
		}
		println()
		return
	}

	// Include the current element
	ans = append(ans, Array[i])
	SubsetUsingBacktracking(Array, ans, i+1)

	//When we are backtracking, we need to remove the last element added to ans
	// This is because we want to explore the path without the current element
	// Exclude the current element
	ans = ans[:len(ans)-1]
	SubsetUsingBacktracking(Array, ans, i+1)
}

func SubsetUsingBacktrackingWithoutprint(Array []int, ans []int, i int, result *[][]int) {
	if i == len(Array) {
		// Append a copy of the current subset to the result
		*result = append(*result, append([]int{}, ans...))
		return
	}

	// Include the current element
	ans = append(ans, Array[i])
	SubsetUsingBacktrackingWithoutprint(Array, ans, i+1, result)

	// When we are backtracking, we need to remove the last element added to ans
	// This is because we want to explore the path without the current element
	// Exclude the current element
	ans = ans[:len(ans)-1]
	SubsetUsingBacktrackingWithoutprint(Array, ans, i+1, result)
}
func SubsetUsingBacktrackingWithoutDuplicate(Array []int, ans []int, i int, result *[][]int, seen map[string]bool) {
	// Use a map to track unique subsets
	// seen := make(map[string]bool)

	if i == len(Array) {
		// Convert the subset to a string to check uniqueness
		subsetKey := fmt.Sprint(ans)
		if !seen[subsetKey] {
			*result = append(*result, append([]int{}, ans...))
			seen[subsetKey] = true
		}
		return
	}

	// Include the current element
	ans = append(ans, Array[i])
	SubsetUsingBacktrackingWithoutDuplicate(Array, ans, i+1, result, seen)

	// When we are backtracking, we need to remove the last element added to ans
	// This is because we want to explore the path without the current element
	// Exclude the current element
	ans = ans[:len(ans)-1]
	SubsetUsingBacktrackingWithoutDuplicate(Array, ans, i+1, result, seen)
}
