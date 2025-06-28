package problems

import "fmt"

func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// Calculate prefix products
	prefix := 1
	for i := 0; i < n; i++ {
		result[i] = prefix
		prefix *= nums[i]
	}

	// Calculate suffix products and multiply with prefix products
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= suffix
		suffix *= nums[i]
	}

	return result
}

func ProductExceptSelf___(nums []int) []int {
	fmt.Println(len(nums))
	result := make([]int, len(nums))
	fmt.Println(len(result))

	for i, _ := range nums {
		res := 1
		for j := 0; j < len(nums); j++ {
			if i != j {
				res *= nums[j]
			}
		}
		result[i] = res
	}
	return result
}
