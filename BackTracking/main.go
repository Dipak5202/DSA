package main

import (
	backtracking "DSA/BackTracking/subset"
	"fmt"
)

func main() {
	// Example usage of the SubsetUsingBacktracking function
	Array := []int{1, 2, 3}
	ans := []int{}
	backtracking.SubsetUsingBacktracking(Array, ans, 0)

	// Example usage of the SubsetUsingBacktrackingWithoutprint function
	result := [][]int{}
	backtracking.SubsetUsingBacktrackingWithoutprint(Array, ans, 0, &result)
	// Print the result
	for _, subset := range result {
		for _, v := range subset {
			print(v, " ")
		}
		println()
	}

	// Example usage of the SubsetUsingBacktrackingWithoutDuplicate function
	Array2 := []int{1, 2, 2}
	// Use a map to track unique subsets
	seen := make(map[string]bool)
	resultWithoutDuplicate := [][]int{}
	backtracking.SubsetUsingBacktrackingWithoutDuplicate(Array2, ans, 0, &resultWithoutDuplicate, seen)
	// Print the result without duplicates
	for _, subset := range resultWithoutDuplicate {
		fmt.Println(subset)
	}
}
